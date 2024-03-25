FROM golang:latest AS builder

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -v -a -installsuffix cgo -o finance ./cmd/server

EXPOSE 8080

CMD ["./finance"]