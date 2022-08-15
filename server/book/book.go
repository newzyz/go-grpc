package book

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/google/uuid"
	"github.com/newzyz/go-grpc/server/storage"

	_ "github.com/lib/pq"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/newzyz/go-grpc/services/booksapp"
)

const (
	// Docker IP IF IN DOCKER
	host     = "localhost"
	port     = 5435
	user     = "root"
	password = "root"
	dbname   = "postgres"
)

type Server struct {
	pb.UnimplementedBookServer
	storage storage.Manager
}

func NewServer(storage storage.Manager) Server {
	return Server{
		storage: storage,
	}
}
func OpenConnection() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return db
}

func (s Server) Download(req *pb.DownloadRequest, responseStream pb.Book_DownloadServer) error {
	bufferSize := 64 * 1024 //64KiB
	file, err := os.Open("./server/tmp/" + req.GetName())
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer file.Close()
	buff := make([]byte, bufferSize)
	for {
		bytesRead, err := file.Read(buff)
		if err != nil {
			if err != io.EOF {
				fmt.Println(err)
			}
			break
		}
		resp := &pb.DownloadResponse{
			Chunk: buff[:bytesRead],
		}
		err = responseStream.Send(resp)
		if err != nil {
			log.Println("error while sending chunk:", err)
			return err
		}
	}
	return nil
}

func (s Server) Upload(stream pb.Book_UploadServer) error {
	//unique name
	name := (uuid.New()).String()
	initial := true
	file := storage.NewFile(name)
	for {
		req, err := stream.Recv()
		if initial {
			//GET FILE TYPE AND CONCAT TO FILE NAME
			name += req.GetMime()
			file = storage.NewFile(name)
			initial = false
		}
		if err == io.EOF {
			if err := s.storage.Store(file); err != nil {
				return status.Error(codes.Internal, err.Error())
			}

			return stream.SendAndClose(&pb.UploadResponse{Name: name})
		}
		if err != nil {
			return status.Error(codes.Internal, err.Error())
		}

		if err := file.Write(req.GetChunk()); err != nil {
			return status.Error(codes.Internal, err.Error())
		}
	}
}

func (s Server) GetBooks(ctx context.Context, in *pb.GetBooksReq) (*pb.GetBooksResp, error) {
	log.Printf("Received Page: %v Per_page: %v", in.GetPage(), in.GetPerPage())

	per_page := in.GetPerPage()
	page := (in.GetPage() - 1) * per_page
	if per_page < 1 || page < 0 {
		return nil, errors.New("page or per_page be integer and must equal or greater than 1")
	}
	db := OpenConnection()
	rows, err := db.Query("SELECT id,title,genre,author FROM book ORDER BY id OFFSET $1 ROWS FETCH NEXT $2 ROWS ONLY;", page, per_page)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()
	defer db.Close()
	var resBooks pb.GetBooksResp
	for rows.Next() {
		var book pb.BookInfo
		if err := rows.Scan(&book.Id, &book.Title, &book.Genre, &book.Author); err != nil {
			log.Println(err)
			return nil, err
		}
		resBooks.Books = append(resBooks.Books, &book)
	}
	if err = rows.Err(); err != nil {
		return &resBooks, err
	}
	return &resBooks, nil
}

func (s Server) GetBook(ctx context.Context, in *pb.Id) (*pb.BookInfo, error) {
	db := OpenConnection()
	defer db.Close()
	// Convert If type is not fitted, Downgrading type may cause data loss
	// log.Printf("%v", reflect.TypeOf(in.GetId()))
	// bookId, err := strconv.ParseInt(in.GetId(), 0, 64)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Printf("%v", reflect.TypeOf(bookId))
	log.Printf("Received Book id: %v", in.GetId())
	bookId := in.GetId()
	row := db.QueryRow("SELECT id,title,genre,author FROM book WHERE id = $1", bookId)
	var book pb.BookInfo
	switch err := row.Scan(&book.Id, &book.Title, &book.Genre, &book.Author); err {
	case sql.ErrNoRows:
		fmt.Println("No row was returned!")
		return nil, errors.New("no row was returned")
	case nil:
		return &book, nil
	default:
		log.Fatal(err)
		return nil, err
	}
}

func (s Server) CreateBook(ctx context.Context, in *pb.BookInfo) (*pb.Id, error) {
	db := OpenConnection()
	defer db.Close()

	var BookId int64
	err := db.QueryRow("INSERT INTO book(title,genre,author) VALUES ($1, $2, $3) RETURNING id;", in.GetTitle(), in.GetGenre(), in.GetAuthor()).Scan(&BookId)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	createdBookId := &pb.Id{Id: BookId}
	log.Printf("Received Created BookInfo: %v a Created Book Id: %v", in, createdBookId.GetId())

	return createdBookId, nil
}

func (s Server) UpdateBook(ctx context.Context, in *pb.BookInfo) (*pb.Status, error) {

	db := OpenConnection()
	res, err := db.Exec("UPDATE book SET title = $1, genre = $2, author = $3 WHERE id=$4;", in.GetTitle(), in.GetGenre(), in.GetAuthor(), in.GetId())
	if err != nil {
		log.Fatal(err)
	}
	status := pb.Status{}

	count, err := res.RowsAffected()
	if err != nil {
		status.Value = 0
		return &status, err
	}
	if count == 0 {
		status.Value = 1
		return &status, errors.New("the selected book id is not existed")
	}
	status.Value = 1
	return &status, nil
}
func (s Server) DeleteBook(ctx context.Context, in *pb.Id) (*pb.Status, error) {
	log.Printf("Received Book id: %v", in.GetId())

	deletedBookId := in.GetId()
	db := OpenConnection()
	defer db.Close()

	status := pb.Status{}
	res, err := db.Exec("DELETE FROM book WHERE id = $1", deletedBookId)
	if err != nil {
		status.Value = 0
		return &status, err
	}
	count, err := res.RowsAffected()
	if err != nil {
		status.Value = 0
		return &status, err
	}

	if count == 0 {
		status.Value = 2
		return &status, nil
	}
	status.Value = 1
	return &status, nil
}
