FROM golang:alpine as builder

WORKDIR /go/src/app
COPY . .
RUN go build -o bin/server main.go
CMD ["./bin/server"]