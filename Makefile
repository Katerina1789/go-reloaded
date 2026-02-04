.PHONY: all build test audit clean run help

all: build

build:
	@go build -o go-reloaded .

test:
	@go test ./testfiles/... -v

audit:
	@go run . sample.txt result.txt --audit

clean:
	@rm -f go-reloaded sample.txt result.txt

run:
	@go run . input.txt output.txt

help:
	@echo "Available targets:"
	@echo "  build  - Build the binary"
	@echo "  test   - Run unit tests"
	@echo "  audit  - Run audit mode tests"
	@echo "  clean  - Remove build artifacts"
	@echo "  run    - Run with input.txt and output.txt"
	@echo "  help   - Show this help message"
