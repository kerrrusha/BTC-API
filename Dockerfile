FROM golang:latest

LABEL version="1.0"

WORKDIR /app

COPY go.mod .

RUN go mod download

COPY . .

ENV PORT 8000

RUN go build

CMD ["./BTC-API"]
