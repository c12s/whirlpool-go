FROM golang:1.20-alpine

WORKDIR /app

COPY . .

RUN go mod init test_rest
RUN go build -o main.go .

EXPOSE 8080

ENTRYPOINT ["./dummy-server"]
