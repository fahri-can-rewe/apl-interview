# Justfile for apl-interview Go project

# Default recipe: show help
default: help

# Variables
BIN := "wordpair"
CMD_DIR := "cmd/wordpair"
MAIN := "{{CMD_DIR}}/main.go"
DOCKER_IMAGE := "apl-interview:local"

# Determine GOOS/GOARCH from environment when provided (empty by default)

# Helper: print section headers
_header msg:
    @printf "\n==> %s\n" "{{msg}}"

# Helper: ensure Docker daemon is reachable
_docker_available:
    if ! docker info >/dev/null 2>&1; then \
      echo "Docker is not running. Start Docker Desktop and retry."; \
      exit 1; \
    fi

help:
    @echo "Available recipes:"
    @echo "  test           # Run all unit tests with race detector and coverage"
    @echo "  build          # Build the CLI (./bin/{{BIN}})"
    @echo "  build-release  # Build statically linked release binary (./dist/{{BIN}})"
    @echo "  run ARGS=...   # Run the CLI with optional ARGS (passed after --)"
    @echo "  fmt            # go fmt on all modules"
    @echo "  lint           # go vet"
    @echo "  tidy           # go mod tidy"
    @echo "  clean          # Remove build artifacts"
    @echo "  docker-build   # Build Docker image {{DOCKER_IMAGE}}"
    @echo "  docker-run     # Run the Docker image"

# Run unit tests
test:
    just _header "Running tests"
    go test ./... -race -coverprofile=coverage.out -covermode=atomic

# Format code
fmt:
    just _header "Formatting"
    go fmt ./...

# Lint with go vet
lint:
    just _header "go vet"
    go vet ./...

# Tidy modules
tidy:
    just _header "go mod tidy"
    go mod tidy

# Build binary into ./bin
build:
    just _header "Building"
    mkdir -p bin
    GOOS={{env_var_or_default("GOOS", "")}} GOARCH={{env_var_or_default("GOARCH", "")}} go build -o bin/{{BIN}} {{MAIN}}

# Build statically linked release binary into ./dist
build-release:
    just _header "Building release"
    mkdir -p dist
    CGO_ENABLED=0 GOOS={{env_var_or_default("GOOS", "")}} GOARCH={{env_var_or_default("GOARCH", "")}} go build -trimpath -ldflags "-s -w" -o dist/{{BIN}} {{MAIN}}

# Run the CLI from the module path. Pass CLI args via ARGS variable, e.g.: just run ARGS="-apiBaseUrl https://interview.sowula.at"
run ARGS="":
    just _header "Running {{BIN}}"
    go run ./{{CMD_DIR}} -- {{ARGS}}

# Clean build artifacts
clean:
    just _header "Cleaning"
    rm -rf bin dist coverage.out

# Docker build using provided Dockerfile
docker-build: _docker_available
    just _header "Docker build"
    docker build -t {{DOCKER_IMAGE}} .

# Docker run: forward API_BASE_URL env if provided
docker-run: _docker_available
    just _header "Docker run"
    if [ -n "${API_BASE_URL:-}" ]; then \
      docker run --rm -e API_BASE_URL="$API_BASE_URL" {{DOCKER_IMAGE}}; \
    else \
      docker run --rm {{DOCKER_IMAGE}}; \
    fi
