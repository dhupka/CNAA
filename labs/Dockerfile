FROM golang:1.15

WORKDIR /labs/
COPY go.mod .
COPY go.sum .
RUN go mod download 
COPY lab7/grpc-go ./lab7/grpc-go
COPY lab7/movieapi ./lab7/movieapi
COPY lab7/movieserver ./lab7/movieserver
COPY lab7/movieclient ./lab7/movieclient


