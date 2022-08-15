package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"path/filepath"

	"github.com/google/uuid"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"

	_ "github.com/lib/pq"
	"github.com/newzyz/go-grpc/server/book"
	"github.com/newzyz/go-grpc/server/customer"
	"github.com/newzyz/go-grpc/server/storage"
	pb "github.com/newzyz/go-grpc/services/booksapp"
	pb2 "github.com/newzyz/go-grpc/services/customersapp"

	"google.golang.org/grpc"
)

const (
	grpc_port         = ":3000"
	grpc_gateway_port = ":3001"
)

func main() {
	bookSrv := book.NewServer(storage.New("./server/tmp/"))
	customerSrv := customer.NewServer(storage.New("./server/tmp2/"))
	go func() {
		//mux
		mux := runtime.NewServeMux()
		mux.HandlePath("POST", "/uploadFileHttp", handleBinaryFileUpload)

		//Run concurrent
		pb.RegisterBookHandlerServer(context.Background(), mux, bookSrv)
		pb2.RegisterCustomerHandlerServer(context.Background(), mux, customerSrv)

		//http server
		log.Fatalln(http.ListenAndServe(grpc_gateway_port, mux))
	}()

	lis, err := net.Listen("tcp", grpc_port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	pb.RegisterBookServer(s, bookSrv)
	pb2.RegisterCustomerServer(s, customerSrv)

	log.Printf("GRPC server listening at %v", lis.Addr())
	log.Printf("GRPC gateway HTTP listening at %v", grpc_gateway_port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to server %v", err)
	}
}
func handleBinaryFileUpload(w http.ResponseWriter, r *http.Request, params map[string]string) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to parse form: %s", err.Error()), http.StatusBadRequest)
		return
	}

	f, header, err := r.FormFile("attachment")
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to get file 'attachment': %s", err.Error()), http.StatusBadRequest)
		return
	}
	defer f.Close()
	// open output file
	fmt.Println(header.Filename)
	fileExtension := filepath.Ext(header.Filename)
	filename := uuid.New().String() + fileExtension
	fo, err := os.Create("./server/tmpHttp/" + filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	// close fo on exit and check for its returned error
	defer func() {
		if err := fo.Close(); err != nil {
			panic(err)
		}
	}()

	// make a buffer to keep chunks that are read
	buf := make([]byte, 1024)
	for {
		// read a chunk
		n, err := f.Read(buf)
		if err != nil && err != io.EOF {
			panic(err)
		}
		if n == 0 {
			break
		}

		// write a chunk
		if _, err := fo.Write(buf[:n]); err != nil {
			panic(err)
		}
	}
	//
	// Now do something with the io.Reader in `f`, i.e. read it into a buffer or stream it to a gRPC client side stream.
	// Also `header` will contain the filename, size etc of the original file.
	//
}
