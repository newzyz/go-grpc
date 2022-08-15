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
	client := pb.NewClient(conn, storage.New("./downloaded/"))
	name, err := client.Upload(context.Background(), flag.Arg(0))
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(name)

	// Start downloading the file. Error if failed
	fname, err := client.Download(context.Background(), "some-unique-name.png")
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("Download Success file name: %s", fname)

}
