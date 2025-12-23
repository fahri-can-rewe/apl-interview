### apl-interview — Word Pair Anagram Checker (CLI)

A small Go CLI that fetches a pair of words from a remote API and prints whether they are anagrams of each other.

---

### Features

- Fetches a random word pair from a public API
- Validates input (alphabetic only, same length)
- Determines if the words are anagrams
- Simple flags for configuration (`-apiBaseUrl`)
- Containerized with a small Alpine-based image

---

### Requirements

- Go (recommended: Go 1.24+)
- Docker (optional, for containerized runs)

---

### Project Structure

```
.
├─ cmd/wordpair/                 # CLI entrypoint (composition root)
│  ├─ main.go
│  ├─ helpers.go                 # flags, endpoint construction, API client wiring
│  ├─ helpers_test.go
│  └─ helpers_integration_test.go
├─ anagram/                      # domain: validation and anagram strategies
│  ├─ checker.go
│  ├─ checker_test.go
│  ├─ frequency_mapper.go
│  ├─ frequency_mapper_test.go
│  ├─ validation.go
│  └─ validation_test.go
├─ httpclient/                   # minimal HTTP API client
│  ├─ api.go
│  └─ api_integration_test.go
├─ Dockerfile                    # multi-stage build (Alpine)
├─ go.mod
├─ LICENSE
├─ requirements.md
└─ README.md
```

---

### How it works

- Default API base URL: `https://interview.sowula.at`
- Endpoint used: `<base>/word-pair`
- The CLI:
    - parses `-apiBaseUrl` (optional)
    - builds the final endpoint
    - fetches the word pair with a 5s timeout
    - validates words (alphabetic only, same length)
    - prints the words and whether they are anagrams

Example output:

```
Word 1: listen
Word 2: silent
Are Anagrams: true
```

---

### Local development

- Run directly:
  ```bash
  go run ./cmd/wordpair
  # or with a custom API base:
  go run ./cmd/wordpair -apiBaseUrl https://interview.sowula.at
  ```

- Build a local binary:
  ```bash
  go build -o bin/wordpair ./cmd/wordpair
  ./bin/wordpair
  ```

---

### Testing

Run unit and integration tests:

```bash
go test ./...
```

Run with coverage:

```bash
go test ./... -cover
```


---

### Configuration

- Flags:
    - `-apiBaseUrl` (string) — Base URL of the Word-Pair API. Default: `https://interview.sowula.at`

Examples:

```bash
./wordpair -apiBaseUrl https://interview.sowula.at
```

---

### Docker

A production-friendly, multi-stage Dockerfile is included.

- Build the image:
  ```bash
  docker build -t apl-interview:latest .
  ```

- Run with defaults:
  ```bash
  docker run --rm apl-interview:latest
  ```

- Run with a custom API base URL:
  ```bash
  docker run --rm apl-interview:latest -apiBaseUrl https://interview.sowula.at
  ```

- Multi-arch push (optional, requires Buildx):
  ```bash
  docker buildx build \
    --platform linux/amd64,linux/arm64 \
    -t yourrepo/apl-interview:latest \
    --push .
  ```

Notes:

- The final image is based on `alpine` with CA certificates for HTTPS.
- The binary is built with `CGO_ENABLED=0` for portability.

---

### Troubleshooting

- `bad --apiBaseUrl` error: Ensure the value includes a scheme and host (e.g., `https://example.com`).
- Networking in Docker: Your network must allow outbound HTTPS to the API’s domain.
- Timeouts: The HTTP client uses a 5s timeout; intermittent network issues can cause failures.

---

### Contributing

- Format and lint as per standard Go tooling.
- Add tests for new logic (`go test ./...`).
- Keep the container image minimal and non-root where possible.

---
