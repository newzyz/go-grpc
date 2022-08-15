package customer

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
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	_ "github.com/lib/pq"

	pb "github.com/newzyz/go-grpc/services/customersapp"
)

const (
	// Docker IP
	host     = "localhost"
	port     = 5435
	user     = "root"
	password = "root"
	dbname   = "postgres"
)

type Server struct {
	pb.UnimplementedCustomerServer
	storage storage.Manager
}

func NewServer(storage storage.Manager) Server {
	return Server{
		storage: storage,
	}
}
func (s Server) Upload(stream pb.Customer_UploadServer) error {
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
func (s Server) Download(req *pb.DownloadRequest, responseStream pb.Customer_DownloadServer) error {
	bufferSize := 64 * 1024 //64KiB
	file, err := os.Open("./server/tmp2/" + req.GetName())
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

func (s Server) GetCustomers(ctx context.Context, in *pb.GetCustomersReq) (*pb.GetCustomersResp, error) {
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
	var resCustomers pb.GetCustomersResp
	for rows.Next() {
		var customer pb.CustomerInfo
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

func (s Server) GetCustomer(ctx context.Context, in *pb.Id) (*pb.CustomerInfo, error) {
	db := OpenConnection()
	defer db.Close()

	log.Printf("Received Customer id: %v", in.GetId())
	customerId := in.GetId()
	row := db.QueryRow("SELECT id,firstname,lastname,age FROM customer WHERE id = $1", customerId)
	var customer pb.CustomerInfo
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

func (s Server) CreateCustomer(ctx context.Context, in *pb.CustomerInfo) (*pb.Id, error) {
	db := OpenConnection()
	defer db.Close()

	var customerId int64
	err := db.QueryRow("INSERT INTO customer(firstname,lastname,age) VALUES ($1, $2, $3) RETURNING id;", in.GetFirstname(), in.GetLastname(), in.GetAge()).Scan(&customerId)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	createdCustomerId := &pb.Id{Id: customerId}
	log.Printf("Received Created CustomerInfo: %v a Created Customer Id: %v", in, createdCustomerId.GetId())

	return createdCustomerId, nil
}

func (s Server) UpdateCustomer(ctx context.Context, in *pb.CustomerInfo) (*pb.Status, error) {

	db := OpenConnection()
	res, err := db.Exec("UPDATE customer SET firstname = $1, lastname = $2, age = $3 WHERE id=$4;", in.GetFirstname(), in.GetLastname(), in.GetAge(), in.GetId())
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
		return &status, errors.New("the selected customer id is not existed")
	}
	status.Value = 1
	return &status, nil
}

func (s Server) DeleteCustomer(ctx context.Context, in *pb.Id) (*pb.Status, error) {
	log.Printf("Received Customer id: %v", in.GetId())

	deletedCustomerId := in.GetId()
	db := OpenConnection()
	defer db.Close()

	status := pb.Status{}
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
