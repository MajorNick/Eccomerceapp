FROM golang:latest

WORKDIR /ecommerce
COPY . /ecommerce

RUN go build -o main main.go

CMD ["./main"]