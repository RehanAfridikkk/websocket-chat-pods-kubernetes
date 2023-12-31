# Use the official Golang image as the base image
FROM golang:latest as builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .


RUN go build -o main .



EXPOSE 8080
CMD ["./main"]
