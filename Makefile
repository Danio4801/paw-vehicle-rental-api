.PHONY: help run build test clean

help:
	@echo "Komendy:"
	@echo "  make run    - Uruchom aplikację lokalnie"
	@echo "  make build  - Zbuduj aplikację"
	@echo "  make test   - Uruchom testy"
	@echo "  make clean  - Wyczyść temp files"

run:
	go run cmd/api/main.go

build:
	go build -o bin/api cmd/api/main.go

test:
	go test ./...

clean:
	rm -rf bin/
	go clean