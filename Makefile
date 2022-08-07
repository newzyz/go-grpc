create:
	protoc --go_out=./server ./proto/*.proto 
	protoc --go-grpc_out=./server ./proto/*.proto  
	protoc -I . --grpc-gateway_out ./server ./proto/*.proto	

clean:
	rmdir server/services

run:
	go run server/server.go
