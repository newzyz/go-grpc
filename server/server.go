package main

import (
	"context"
	"log"
	"net"
	"net/http"

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
