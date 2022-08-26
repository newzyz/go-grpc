package main

import (
	"context"
	"log"

	// gtrace "github.com/moxiaomomo/grpc-jaeger"
	pb "github.com/newzyz/go-grpc/client/fileserve"
	"github.com/newzyz/go-grpc/client/storage"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
)

// func rpcCli(dialOpts []grpc.DialOption) {
// 	conn, err := grpc.Dial("127.0.0.1:3000", dialOpts...)
// 	if err != nil {
// 		fmt.Printf("grpc connect failed, err:%+v\n", err)
// 		return
// 	}
// 	defer conn.Close()
// 	client := pb.NewBookClient(conn, storage.New("./downloaded/"))
// 	// TODO: do some rpc-call
// 	// ...
// 	// client.GetBooks(context.Background(), 1, 4)
// 	client.GetBookId(context.Background(), 1)
// 	////stream
// 	// flag.Parse()
// 	// if flag.NArg() == 0 {
// 	// 	log.Fatalln("Missing file path")
// 	// }
// 	// log.Println(flag.Args())
// 	// name, err := client.UploadBook(context.Background(), flag.Arg(0))
// 	// if err != nil {
// 	// 	log.Fatalln(err)
// 	// }
// 	// filename, err := client.DownloadBook(context.Background(), "a4313a37-3d47-4c14-90a7-cf872a2699b9.png")
// 	// log.Println(filename)
// }

func main() {
	// dialOpts := []grpc.DialOption{grpc.WithUnaryInterceptor(otelgrpc.UnaryClientInterceptor()),
	// 	grpc.WithStreamInterceptor(otelgrpc.StreamClientInterceptor()),
	// 	grpc.WithInsecure()}
	// // Catch user input.
	// flag.Parse()
	// if flag.NArg() == 0 {
	// 	log.Fatalln("Missing file path")
	// }
	// log.Println(flag.Args())

	// tracer, _, err := gtrace.NewJaegerTracer("testCli", "127.0.0.1:6831")
	// if err != nil {
	// 	fmt.Printf("new tracer err: %+v\n", err)
	// 	os.Exit(-1)
	// }

	// if tracer != nil {
	// 	dialOpts = append(dialOpts, gtrace.DialOption(tracer))
	// }
	// do rpc-call with dialOpts
	// rpcCli(dialOpts)

	// Initialise gRPC connection.
	conn, err := grpc.Dial("127.0.0.1:3000", grpc.WithUnaryInterceptor(otelgrpc.UnaryClientInterceptor()),
		grpc.WithStreamInterceptor(otelgrpc.StreamClientInterceptor()),
		grpc.WithInsecure())
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()
	// TODO: do some rpc-call
	client := pb.NewBookClient(conn, storage.New("./downloaded/"))

	client.GetBookId(context.Background(), 1)
	// client := pb.NewBookClient(conn, storage.New("./downloaded/"))

	// client.GetBookId(context.Background(), 1)

	// // Start uploading the file. Error if failed, otherwise echo download URL.
	// client := pb.NewBookClient(conn, storage.New("./downloaded/"))
	// // client2 := pb.NewCustomerClient(conn, storage.New("./downloaded/"))
	// name, err := client.UploadBook(context.Background(), flag.Arg(0))
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// // name2, err := client2.UploadCustomer(context.Background(), flag.Arg(0))
	// // if err != nil {
	// // 	log.Fatalln(err)
	// // }
	// log.Println(name)
	// // log.Println(name2)
	// // Start downloading the file. Error if failed
	// filename, err := client.DownloadBook(context.Background(), "8d70b6be-8d46-4c01-af20-3b09d1aabf9a.png")
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// // filename2, err := client2.DownloadCustomer(context.Background(), "d29a0602-c794-4e1f-9187-ea25c1ebf4c4.png")
	// // if err != nil {
	// // 	log.Fatalln(err)
	// // }
	// log.Printf("Download Success file book name: %s", filename)
	// // log.Printf("Download Success file customer name: %s", filename2)

}
