package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"

	_ "github.com/lib/pq"
	pb "github.com/newzyz/go-grpc/services/booksapp"
	pb2 "github.com/newzyz/go-grpc/services/customersapp"
	"google.golang.org/grpc"
)

const (
	grpc_port         = ":3000"
	grpc_gateway_port = ":3001"
)
const (
	// Docker IP
	host     = "172.26.0.1"
	port     = 5435
	user     = "root"
	password = "root"
	dbname   = "postgres"
)

type bookServer struct {
	pb.UnimplementedBookServer
}
type customerServer struct {
	pb2.UnimplementedCustomerServer
}

func main() {
	go func() {
		//mux
		mux := runtime.NewServeMux()

		//Run paralle
		pb.RegisterBookHandlerServer(context.Background(), mux, &bookServer{})
		pb2.RegisterCustomerHandlerServer(context.Background(), mux, &customerServer{})
		//http server
		log.Fatalln(http.ListenAndServe(grpc_gateway_port, mux))
	}()

	lis, err := net.Listen("tcp", grpc_port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	pb.RegisterBookServer(s, &bookServer{})
	pb2.RegisterCustomerServer(s, &customerServer{})

	log.Printf("server listening at %v", lis.Addr())
	log.Printf("GRPC gateway port: %v", grpc_gateway_port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to server %v", err)
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

func (s *bookServer) GetBooks(ctx context.Context, in *pb.GetBooksReq) (*pb.GetBooksResp, error) {
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

func (s *bookServer) GetBook(ctx context.Context, in *pb.Id) (*pb.BookInfo, error) {
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

func (s *bookServer) CreateBook(ctx context.Context, in *pb.BookInfo) (*pb.Id, error) {
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

func (s *bookServer) UpdateBook(ctx context.Context, in *pb.BookInfo) (*pb.Status, error) {

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
func (s *bookServer) DeleteBook(ctx context.Context, in *pb.Id) (*pb.Status, error) {
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

//Customer
func (s *customerServer) GetCustomers(ctx context.Context, in *pb2.GetCustomersReq) (*pb2.GetCustomersResp, error) {
	log.Printf("Received Page: %v Per_page: %v", in.GetPage(), in.GetPerPage())

	per_page := in.GetPerPage()
	page := (in.GetPage() - 1) * per_page
	if per_page < 1 || page < 0 {
		return nil, errors.New("page or per_page be integer and must equal or greater than 1")
	}
	db := OpenConnection()
	rows, err := db.Query("SELECT id,firstname,lastname,age FROM customer ORDER BY id OFFSET $1 ROWS FETCH NEXT $2 ROWS ONLY;", page, per_page)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()
	defer db.Close()
	var resCustomers pb2.GetCustomersResp
	for rows.Next() {
		var customer pb2.CustomerInfo
		if err := rows.Scan(&customer.Id, &customer.Firstname, &customer.Lastname, &customer.Age); err != nil {
			log.Println(err)
			return nil, err
		}
		resCustomers.Customers = append(resCustomers.Customers, &customer)
	}
	if err = rows.Err(); err != nil {
		return &resCustomers, err
	}
	return &resCustomers, nil
}

func (s *customerServer) GetCustomer(ctx context.Context, in *pb2.Id) (*pb2.CustomerInfo, error) {
	db := OpenConnection()
	defer db.Close()

	log.Printf("Received Customer id: %v", in.GetId())
	customerId := in.GetId()
	row := db.QueryRow("SELECT id,firstname,lastname,age FROM customer WHERE id = $1", customerId)
	var customer pb2.CustomerInfo
	switch err := row.Scan(&customer.Id, &customer.Firstname, &customer.Lastname, &customer.Age); err {
	case sql.ErrNoRows:
		fmt.Println("No row was returned!")
		return nil, errors.New("no row was returned")
	case nil:
		return &customer, nil
	default:
		log.Fatal(err)
		return nil, err
	}
}

func (s *customerServer) CreateCustomer(ctx context.Context, in *pb2.CustomerInfo) (*pb2.Id, error) {
	db := OpenConnection()
	defer db.Close()

	var customerId int64
	err := db.QueryRow("INSERT INTO customer(firstname,lastname,age) VALUES ($1, $2, $3) RETURNING id;", in.GetFirstname(), in.GetLastname(), in.GetAge()).Scan(&customerId)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	createdCustomerId := &pb2.Id{Id: customerId}
	log.Printf("Received Created CustomerInfo: %v a Created Customer Id: %v", in, createdCustomerId.GetId())

	return createdCustomerId, nil
}

func (s *customerServer) UpdateCustomer(ctx context.Context, in *pb2.CustomerInfo) (*pb2.Status, error) {

	db := OpenConnection()
	res, err := db.Exec("UPDATE customer SET firstname = $1, lastname = $2, age = $3 WHERE id=$4;", in.GetFirstname(), in.GetLastname(), in.GetAge(), in.GetId())
	if err != nil {
		log.Fatal(err)
	}
	status := pb2.Status{}

	count, err := res.RowsAffected()
	if err != nil {
		status.Value = 0
		return &status, err
	}
	if count == 0 {
		status.Value = 1
		return &status, errors.New("the selected customer id is not existed")
	}
	status.Value = 1
	return &status, nil
}

func (s *customerServer) DeleteCustomer(ctx context.Context, in *pb2.Id) (*pb2.Status, error) {
	log.Printf("Received Customer id: %v", in.GetId())

	deletedCustomerId := in.GetId()
	db := OpenConnection()
	defer db.Close()

	status := pb2.Status{}
	res, err := db.Exec("DELETE FROM customer WHERE id = $1", deletedCustomerId)
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
