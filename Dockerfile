FROM golang:latest

WORKDIR /eccomerce
COPY . /eccomerce

RUN go build -o main main.go

CMD ["./main"]