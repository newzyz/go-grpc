FROM golang:1.18.3-alpine as build_base
ENV GO111MODULE=on
WORKDIR /api

RUN go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest
RUN go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . ./

RUN find /api/services/customersapp

RUN go build -o apiserver /api/server/server.go

FROM alpine:latest  

COPY --from=build_base api/apiserver apicmd

# EXPOSE 3001
CMD [ "/apicmd" ]