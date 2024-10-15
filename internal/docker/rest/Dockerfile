FROM golang:1.20-alpine

WORKDIR /app

COPY . .

RUN go mod init test_rest
RUN GOOS=linux go build -o main .

RUN chmod +x main

RUN ls -la

EXPOSE 8080

ENTRYPOINT ./main
