create:
	protoc --go_out=./services ./proto/booksapp.proto 
	protoc --go-grpc_out=./services ./proto/booksapp.proto  
	protoc -I . --grpc-gateway_out ./services ./proto/booksapp.proto	
	protoc --go_out=./services ./proto/customersapp.proto 
	protoc --go-grpc_out=./services ./proto/customersapp.proto  
	protoc -I . --grpc-gateway_out ./services ./proto/customersapp.proto	

db-run:
	docker-compose up
go-run:
	go run server/server.go

