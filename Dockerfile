FROM golang:1.23

WORKDIR /app

COPY . .

RUN go build -o math math.go

CMD ["./math"]
