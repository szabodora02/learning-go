# Start from the official Golang image
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Build the Go app
RUN go build -o /splitdim main.go

# Use a minimal base image for runtime
FROM alpine:latest

WORKDIR /root/

# Copy the binary from the builder stage
COPY --from=builder /splitdim .

# Copy static files (for GUI)
COPY --from=builder /app/static ./static

# Run the binary
ENTRYPOINT ["./splitdim"]
