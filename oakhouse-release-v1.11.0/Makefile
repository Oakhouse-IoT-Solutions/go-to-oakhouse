.PHONY: help build run test clean docker-build docker-run

# Variables
APP_NAME=oakhouse-release-v1.11.0
DOCKER_IMAGE=$(APP_NAME):latest

# Default target
help:
	@echo "Available commands:"
	@echo "  build        Build the application"
	@echo "  run          Run the application"
	@echo "  test         Run tests"
	@echo "  clean        Clean build artifacts"
	@echo "  docker-build Build Docker image"
	@echo "  docker-run   Run Docker container"
	@echo "  dev          Start development server"


build:
	@echo "Building $(APP_NAME)..."
	go build -o bin/$(APP_NAME) cmd/main.go

run: build
	@echo "Running $(APP_NAME)..."
	./bin/$(APP_NAME)

dev:
	@echo "Starting development server..."
	air

test:
	@echo "Running tests..."
	go test -v ./...

test-coverage:
	@echo "Running tests with coverage..."
	go test -v -cover ./...

clean:
	@echo "Cleaning..."
	rm -rf bin/
	docker-build:
	@echo "Building Docker image..."
	docker build -t $(DOCKER_IMAGE) .

docker-run:
	@echo "Running Docker container..."
	docker run -p 8080:8080 $(DOCKER_IMAGE)

docker-compose-up:
	@echo "Starting services with Docker Compose..."
	docker-compose up -d

docker-compose-down:
	@echo "Stopping services..."
	docker-compose down



install-deps:
	@echo "Installing dependencies..."
	go mod tidy
	go mod download

format:
	@echo "Formatting code..."
	go fmt ./...

lint:
	@echo "Running linter..."
	golangci-lint run
