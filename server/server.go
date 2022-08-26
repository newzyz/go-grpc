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
	"strconv"

	"github.com/google/uuid"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	_ "github.com/lib/pq"

	// "go.opentelemetry.io/otel"
	// "go.opentelemetry.io/otel/exporters/trace/jaeger"
	// "go.opentelemetry.io/otel/propagation"
	// "go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	// "go.opentelemetry.io/otel/semconv"

	// "go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"

	// gtrace "github.com/moxiaomomo/grpc-jaeger"
	"github.com/newzyz/go-grpc/server/book"
	"github.com/newzyz/go-grpc/server/customer"
	"github.com/newzyz/go-grpc/server/storage"
	pb "github.com/newzyz/go-grpc/services/booksapp"
	pb2 "github.com/newzyz/go-grpc/services/customersapp"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/propagation"
	"google.golang.org/grpc"
)

const (
	grpc_port         = ":3000"
	grpc_gateway_port = ":3001"
)

func main() {

	exp, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint("http://localhost:14268/api/traces")))
	if err != nil {
		log.Fatal(err)
	}
	bsp := sdktrace.NewBatchSpanProcessor(exp)
	tp := sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		// sdktrace.WithResource(resource.NewWithAttributes(
		// 	string(semconv.ServiceNameKey),
		// )),
		sdktrace.WithSpanProcessor(bsp),
	)
	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(
		propagation.NewCompositeTextMapPropagator(
			propagation.TraceContext{},
			propagation.Baggage{},
		),
	)

	bookSrv := book.NewServer(storage.New("./server/tmp/"))
	customerSrv := customer.NewServer(storage.New("./server/tmp2/"))
	go func() {
		//mux
		mux := runtime.NewServeMux()
		mux.HandlePath("POST", "/uploadFileHttp", handleBinaryFileUpload)
		mux.HandlePath("GET", "/downloadFileHttp/{filename}", handleDownload)
		//Run concurrent
		pb.RegisterBookHandlerServer(context.Background(), mux, bookSrv)
		pb2.RegisterCustomerHandlerServer(context.Background(), mux, customerSrv)

		//http server
		log.Fatalln(http.ListenAndServe(grpc_gateway_port, mux))
	}()
	// var servOpts []grpc.ServerOption
	// tracer, _, err := gtrace.NewJaegerTracer("testSrv", "127.0.0.1:6831")
	// if err != nil {
	// 	fmt.Printf("new tracer err: %+v\n", err)
	// 	os.Exit(-1)
	// }
	// if tracer != nil {
	// 	servOpts = append(servOpts, gtrace.ServerOption(tracer))
	// }

	// s := grpc.NewServer(servOpts...)
	s := grpc.NewServer(
		grpc.UnaryInterceptor(otelgrpc.UnaryServerInterceptor()),
		grpc.StreamInterceptor(otelgrpc.StreamServerInterceptor()),
	)

	pb.RegisterBookServer(s, bookSrv)
	pb2.RegisterCustomerServer(s, customerSrv)

	lis, err := net.Listen("tcp", grpc_port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Printf("GRPC server listening at %v", lis.Addr())
	log.Printf("GRPC gateway HTTP listening at %v", grpc_gateway_port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to server %v", err)
	}
}

func handleDownload(w http.ResponseWriter, r *http.Request, params map[string]string) {
	filename := params["filename"]
	filePath := "./server/tmpHttp/" + filename

	w.Header().Set("Content-Disposition", "attachment; filename="+strconv.Quote(filename))
	w.Header().Set("Content-Type", "application/octet-stream")
	http.ServeFile(w, r, filePath)
}

func handleBinaryFileUpload(w http.ResponseWriter, r *http.Request, params map[string]string) {
	// Maximum upload of 10 MB files
	// r.ParseMultipartForm(10 << 20)
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
	// Copy the uploaded file to the created file on the filesystem
	if _, err := io.Copy(fo, f); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Successfully Uploaded File\n")
	// // close fo on exit and check for its returned error
	// defer func() {
	// 	if err := fo.Close(); err != nil {
	// 		panic(err)
	// 	}
	// }()

	// // make a buffer to keep chunks that are read
	// buf := make([]byte, 1024)
	// for {
	// 	// read a chunk
	// 	n, err := f.Read(buf)
	// 	if err != nil && err != io.EOF {
	// 		panic(err)
	// 	}
	// 	if n == 0 {
	// 		break
	// 	}

	// 	// write a chunk
	// 	if _, err := fo.Write(buf[:n]); err != nil {
	// 		panic(err)
	// 	}
	// }
	// //
	// // Now do something with the io.Reader in `f`, i.e. read it into a buffer or stream it to a gRPC client side stream.
	// // Also `header` will contain the filename, size etc of the original file.
	// //
}
