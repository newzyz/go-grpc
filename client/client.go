package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/newzyz/go-grpc/server/services/booksapp"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial(":3000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect %v", err)
	}
	defer conn.Close()

	client := pb.NewBookClient(conn)

	// runGetBooks(client)
	// runGetBook(client, 1)
	// runCreateBook(client, "Lord Of the Ring 3", "Horror", "Marta")
	// runDeleteBook(client, 19)
	runUpdateBook(client, 1, "Lord Of the Ring 1", "Fantasy", "Jame Bone")
}

func runGetBooks(client pb.BookClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	req := &pb.Empty{}
	log.Printf("Run Get Books")

	stream, err := client.GetBooks(ctx, req)
	if err != nil {
		log.Fatalf("%v.GetBooks(_) = _,%v", client, err)

	}
	for {
		row, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("%v.Books(_) = _,%v", client, err)

		}
		log.Printf("BookInfo: %v", row)
	}
}

func runGetBook(client pb.BookClient, bookId int64) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	req := &pb.Id{Id: bookId}
	log.Printf("Get Book at id:%v", bookId)

	res, err := client.GetBook(ctx, req)
	if err != nil {
		log.Fatalf("%v.GetBook(_) = _,%v", client, err)

	}
	log.Printf("BookInfo: %v", res)
}

func runCreateBook(client pb.BookClient, title string, genre string, author string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	req := &pb.BookInfo{Id: 1, Title: title, Genre: genre, Author: author}
	// log.Printf("Create Book at id:%s", bookId)

	res, err := client.CreateBook(ctx, req)
	if err != nil {
		log.Fatalf("%v.CreateBook(_) = _,%v", client, err)

	}
	log.Printf("Book Created ID: %v", res)
}
func runUpdateBook(client pb.BookClient, bookId int64, title string, genre string, author string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	req := &pb.BookInfo{Id: bookId, Title: title, Genre: genre, Author: author}
	// log.Printf("Create Book at id:%s", bookId)

	res, err := client.UpdateBook(ctx, req)
	if err != nil {
		log.Fatalf("%v.UpdateBook(_) = _,%v", client, err)

	}
	if res.GetValue() == 0 {
		log.Printf("Update Failed: Status %v", res.GetValue())
	} else if res.GetValue() == 2 {
		log.Printf("Update ID is not Existed: Status %v", res.GetValue())
	} else {
		log.Printf("Updated Successfully: Status %v", res.GetValue())
	}
}

func runDeleteBook(client pb.BookClient, bookId int64) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	req := &pb.Id{Id: bookId}
	log.Printf("Delete Book at id:%v", bookId)

	res, err := client.DeleteBook(ctx, req)
	if err != nil {
		log.Fatalf("%v.DeleteBook(_) = _,%v", client, err)

	}
	if res.GetValue() == 0 {
		log.Printf("Delete Failed: Status %v", res.GetValue())
	} else if res.GetValue() == 2 {
		log.Printf("Delete ID is not Existed: Status %v", res.GetValue())
	} else {
		log.Printf("Deleted Successfully: Status %v", res.GetValue())
	}
}
