FROM golang:1.18.3-alpine as build_base
# RUN go env -w GO111MODULE=on

# RUN mkdir /api
WORKDIR /app

# ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on

RUN go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest
RUN go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

# download the required Go dependencies
COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

RUN ls

RUN go build -o /app/server/cmd /app/server/

FROM alpine:latest  

COPY --from=build_base /app/server/cmd /app/server/cmd

EXPOSE 8080
CMD [ "/app/server/cmd" ]
