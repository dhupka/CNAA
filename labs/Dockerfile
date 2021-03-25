FROM golang:1.15-alpine AS build

WORKDIR /labs/
COPY go.mod .
COPY go.sum .
RUN go mod download 
WORKDIR /labs/
COPY lab7/mongoserver.go /labs/
RUN CGO_ENABLED=0 go build -o /bin/webserver

FROM scratch
COPY --from=build /bin/webserver /bin/webserver
ENTRYPOINT ["/bin/webserver"]

