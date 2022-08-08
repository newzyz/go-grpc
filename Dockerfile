FROM golang:1.18.3-alpine as build_base
ENV GO111MODULE=on
WORKDIR /api

RUN go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway
RUN go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . ./

RUN find /api/services/customersapp

# RUN go mod edit -replace github.com/newzyz/go-grpc/services/booksapp=$GOPATH/src/api/services/booksapp
# RUN go mod edit -replace github.com/newzyz/go-grpc/services/customersapp=$GOPATH/src/api/services/customersapp

RUN go build -o apiserver /api/server/server.go
# RUN CGO_ENABLED=0 GOOS=linux go build -mod=vendor -o server ./server/server.go 
# RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o api ./server/server.go
FROM alpine:latest  

COPY --from=build_base api/apiserver apicmd
# COPY --from=builder apiserver /api/apiserver
# COPY --from=builder api/apiserver .
EXPOSE 3000
CMD [ "/apicmd" ]