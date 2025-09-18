FROM golang:1.17 AS builder

WORKDIR /app

# Copy go.mod and go.sum first
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Start a new stage from scratch (for a minimal final image)
FROM alpine:latest

WORKDIR /root/

# Copy the binary from the builder
COPY --from=builder /app/main .

EXPOSE 8080

CMD ["./main"]
