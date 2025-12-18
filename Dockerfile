## --- Stage 1: Build ---
FROM golang:1.24-alpine AS builder

# Install git if needed for private modules
RUN apk add --no-cache git

WORKDIR /app

# 1. Copy dependency files first (for better caching)
COPY go.mod go.sum* ./
RUN go mod download

# 2. Copy the rest of the source code
COPY . .

# 3. Build the binary
# We target the directory where your main.go is located.
# CGO_ENABLED=0 ensures a static binary for Alpine/Scratch.
RUN CGO_ENABLED=0 GOOS=linux go build -o wordpair ./cmd/wordpair

## --- Stage 2: Final Image ---
FROM alpine:latest

# Install CA certificates for HTTPS requests
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy only the compiled binary from the builder stage
COPY --from=builder /app/wordpair .

# Run the binary
ENTRYPOINT ["./wordpair"]