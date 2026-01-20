BINARY_NAME=cafe
BUILD_PATH=bin/$(BINARY_NAME)
MAIN_PATH=$(BINARY_NAME)/main.go
COMPOSE_DEV=toolchain/docker-compose.dev.yml
TAILWIND=toolchain/tailwind

.PHONY: setup clean tidy build run dev test postgres stop reset css watch all

setup:
	@echo "Setting up environment..."
	@go mod download
	@go mod tidy
	@./scripts/tailwind.setup.sh
	@./scripts/htmx.setup.sh
	@echo "Environment setup complete."

clean:
	@echo "Cleaning up..."
	@rm -rf bin
	@rm -rf tmp
	@echo "Cleanup complete."

tidy:
	@echo "Tidying modules..."
	@go mod tidy
	@echo "Modules tidied."

build:
	@echo "Building..."
	@go build -o $(BUILD_PATH) $(MAIN_PATH) || true
	@echo "Build complete."

run:
	@if [ ! -f $(BUILD_PATH) ]; then echo "Binary not found. Building binary..."; $(MAKE) -s build; fi
	@echo "Running..."
	@$(BUILD_PATH) || true

dev: css
	@echo "Running in development mode..."
	@go run $(MAIN_PATH) || true

test:
	@echo "Running tests..."
	@go test -v ./... || true

postgres:
	@echo "Starting database..."
	@docker-compose -f $(COMPOSE_DEV) up -d

stop:
	@echo "Stopping database..."
	@docker-compose -f $(COMPOSE_DEV) down

reset:
	@echo "Resetting database..."
	@./scripts/reset.db.sh

css:
	@if [ ! -f $(TAILWIND) ]; then echo "Tailwind not found. Installing..."; ./scripts/install.tailwind.sh; fi
	@echo "Building CSS..."
	@$(TAILWIND) -i static/css/tailwind.css -o static/css/style.css --minify

watch:
	@if [ ! -f $(TAILWIND) ]; then echo "Tailwind not found. Installing..."; ./scripts/install.tailwind.sh; fi
	@echo "Watching CSS..."
	@$(TAILWIND) -i static/css/tailwind.css -o static/css/style.css --watch

all: setup clean build run

.SILENT: