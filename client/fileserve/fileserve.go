package fileserve

import (
	"context"
	"io"
	"log"
	"os"
	"time"

	"google.golang.org/grpc"

	"github.com/gabriel-vasile/mimetype"
	"github.com/newzyz/go-grpc/client/storage"
	pb "github.com/newzyz/go-grpc/services/booksapp"
)

type Client struct {
	client  pb.BookClient
	storage storage.Manager
}

func NewClient(conn grpc.ClientConnInterface, storage storage.Manager) Client {
	return Client{
		client: pb.NewBookClient(conn), storage: storage,
	}
}

func (c Client) UploadBook(ctx context.Context, file string) (string, error) {
	ctx, cancel := context.WithDeadline(ctx, time.Now().Add(10*time.Second))
	defer cancel()

	stream, err := c.client.Upload(ctx)
	if err != nil {
		return "", err
	}
	mtype, err := mimetype.DetectFile(file)
	if err != nil {
		return "", err
	}
	fil, err := os.Open(file)
	if err != nil {
		return "", err
	}
	// Maximum 1KB size per stream.
	buf := make([]byte, 1024)

	for {
		num, err := fil.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			return "", err
		}

		if err := stream.Send(&pb.UploadRequest{Chunk: buf[:num], Mime: mtype.Extension()}); err != nil {
			return "", err
		}
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		return "", err
	}

	return res.GetName(), nil
}

func (c Client) DownloadBook(ctx context.Context, file string) (string, error) {
	ctx, cancel := context.WithDeadline(ctx, time.Now().Add(10*time.Second))
	defer cancel()

	fileStreamResponse, err := c.client.Download(ctx, &pb.DownloadRequest{
		Name: file,
	})
	if err != nil {
		log.Println("error downloading:", err)
		return "", err
	}

	fileStore := storage.NewFile(file)
	for {
		chunkResponse, err := fileStreamResponse.Recv()
		if err == io.EOF {
			if err := c.storage.Store(fileStore); err != nil {
				return "", err
			}
			log.Println("received all chunks")
			return file, nil
		}
		if err != nil {
			log.Println("err receiving chunk:", err)
			return "", err
		}
		// log.Printf("got new chunk with data: %s \n", chunkResponse.Chunk)
		if err := fileStore.Write(chunkResponse.Chunk); err != nil {
			return "", err
		}
	}
}
