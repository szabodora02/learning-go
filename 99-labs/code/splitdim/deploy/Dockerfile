# Builder stage
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Másold a go.mod, go.sum és a vendor könyvtárat
COPY go.mod go.sum ./
COPY vendor ./vendor

# Másold a forráskódot
COPY . .

# Build a binary using vendor
RUN go build -mod=vendor -o /splitdim main.go

# Minimal runtime image
FROM alpine:latest
WORKDIR /root/

# Copy binary and static files
COPY --from=builder /splitdim .
COPY --from=builder /app/static ./static

# Entrypoint
ENTRYPOINT ["./splitdim"]
