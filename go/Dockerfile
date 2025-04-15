# go/Dockerfile
FROM golang:1.23-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

WORKDIR /app/go
RUN go build -o main main.go

EXPOSE 50051
CMD ["./main"]
