.PHONY: build run test clean docker-build docker-run

# Build the application
build:
	go build -o bin/monkey cmd/main.go

# Run the application
run:
	go run cmd/main.go

# Run tests
test:
	go test -v ./...

# Clean build artifacts
clean:
	rm -rf bin/

# Build Docker image
docker-build:
	docker build -t monkey .

# Run with Docker Compose
docker-run:
	docker-compose up --build

# Run with Docker Compose in background
docker-up:
	docker-compose up -d --build

# Stop Docker Compose
docker-down:
	docker-compose down

# Install dependencies
deps:
	go mod download
	go mod tidy

# Format code
fmt:
	go fmt ./...

# Lint code
lint:
	golangci-lint run

# Generate code
generate:
	go generate ./...

# Database migration (add your migration commands here)
migrate-up:
	# Add your migration up commands

migrate-down:
	# Add your migration down commands
