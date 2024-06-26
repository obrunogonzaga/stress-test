FROM golang:1.22 AS builder
LABEL authors="bruno"

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o loadtester main.go

FROM debian:latest

WORKDIR /root/

COPY --from=builder /app/loadtester .

ENTRYPOINT ["./loadtester"]

CMD ["load", "--url=http://google.com", "--requests=1000", "--concurrency=10"]