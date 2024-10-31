FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o booking_api cmd/booking_api/main.go

FROM alpine:latest

# Set the working directory inside the container
WORKDIR /app

COPY --from=builder /app/booking_api .

EXPOSE 8080

CMD ["./booking_api"]