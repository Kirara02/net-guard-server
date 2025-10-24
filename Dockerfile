FROM golang:1.25-alpine AS builder

WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./
RUN go mod tidy

# Copy the rest of the source code
COPY . .

# Build the Go app
RUN go build -o main .

# Use a minimal image for running
FROM alpine:latest
WORKDIR /app

# Copy binary from builder
COPY --from=builder /app/main .


EXPOSE 8080

CMD ["./main"]