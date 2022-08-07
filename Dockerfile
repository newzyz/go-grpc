FROM golang:1.18.4 AS Production
RUN mkdir /api
WORKDIR /api
COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY server ./
COPY proto ./
COPY google ./

RUN go build -o /docker-gs-ping

CMD [ "/docker-gs-ping" ]