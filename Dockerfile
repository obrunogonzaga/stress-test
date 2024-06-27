FROM golang:1.22 AS builder
LABEL authors="bruno"

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o loadtester main.go

FROM debian:latest

WORKDIR /root/

COPY --from=builder /app/loadtester .

ENTRYPOINT ["./loadtester"]

CMD ["load", "--url=http://google.com", "--requests=1000", "--concurrency=10"]