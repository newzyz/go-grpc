package main

import (
	"context"
	"flag"
	"log"

	pb "github.com/newzyz/go-grpc/client/fileserve"
	"github.com/newzyz/go-grpc/client/storage"
	"google.golang.org/grpc"
)

func main() {
	// Catch user input.
	flag.Parse()
	if flag.NArg() == 0 {
		log.Fatalln("Missing file path")
	}
	log.Println(flag.Args())
	// Initialise gRPC connection.
	conn, err := grpc.Dial(":3000", grpc.WithInsecure())
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	// Start uploading the file. Error if failed, otherwise echo download URL.
	client := pb.NewBookClient(conn, storage.New("./downloaded/"))
	// client2 := pb.NewCustomerClient(conn, storage.New("./downloaded/"))
	name, err := client.UploadBook(context.Background(), flag.Arg(0))
	if err != nil {
		log.Fatalln(err)
	}
	// name2, err := client2.UploadCustomer(context.Background(), flag.Arg(0))
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	log.Println(name)
	// log.Println(name2)
	// Start downloading the file. Error if failed
	filename, err := client.DownloadBook(context.Background(), "8d70b6be-8d46-4c01-af20-3b09d1aabf9a.png")
	if err != nil {
		log.Fatalln(err)
	}
	// filename2, err := client2.DownloadCustomer(context.Background(), "d29a0602-c794-4e1f-9187-ea25c1ebf4c4.png")
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	log.Printf("Download Success file book name: %s", filename)
	// log.Printf("Download Success file customer name: %s", filename2)

}
