## Build
# FROM golang:1.18.3-alpine AS builder

# ENV GO111MODULE=on
# WORKDIR /
# ADD google /google
# ADD proto /proto
# ADD server /server
# ADD go.mod ./
# ADD go.sum ./
# RUN find /go.mod
# RUN find server/services/booksapp
# RUN go mod download

# RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
# RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
# RUN go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway
# RUN go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2
# RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc
# RUN go install google.golang.org/protobuf/cmd/protoc-gen-go


## Deploy
# RUN go build -o buildserver /server/server.go


# CMD [ "/buildserver" ]
FROM golang:1.18.3-alpine as build_base
# ENV GO111MODULE=on
WORKDIR /api

COPY google ./google
COPY proto ./proto
COPY server ./server
COPY services ./services
COPY go.mod ./
COPY go.sum ./
COPY tool.go ./
RUN find /api/services/customersapp
RUN go mod download

RUN go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway
RUN go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go

RUN go mod edit -replace github.com/newzyz/go-grpc/services/booksapp=$GOPATH/src/api/services/booksapp
RUN go mod edit -replace github.com/newzyz/go-grpc/services/customersapp=$GOPATH/src/api/services/customersapp

RUN go build -o apiserver /api/server/server.go
# RUN CGO_ENABLED=0 GOOS=linux go build -mod=vendor -o server ./server/server.go 
# RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o api ./server/server.go
FROM alpine:latest  

COPY --from=build_base api/apiserver apicmd
# COPY --from=builder apiserver /api/apiserver
# COPY --from=builder api/apiserver .
EXPOSE 3000
CMD [ "/apicmd" ]