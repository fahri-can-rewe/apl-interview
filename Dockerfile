## --- Stage 1: Build ---
FROM golang:1.24-alpine AS builder
RUN apk add --no-cache git ca-certificates
WORKDIR /app

# 1) Cache deps
COPY go.mod go.sum* ./
RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    go mod download

# 2) Copy source
COPY . .

# 3) Build static binary, smaller output
RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    CGO_ENABLED=0 GOOS=linux go build -trimpath -ldflags "-s -w" -o wordpair ./cmd/wordpair

## --- Stage 2: Final Image ---
FROM alpine:3.20
RUN apk --no-cache add ca-certificates && adduser -D -u 10001 appuser
WORKDIR /home/appuser
COPY --from=builder /app/wordpair ./wordpair
USER appuser
ENTRYPOINT ["./wordpair"]