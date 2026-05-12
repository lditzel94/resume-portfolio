# resume-portfolio — cross-platform dev/run targets
#
# Usage:
#   make dev   → live-reload via air (auto-detects Windows vs Unix)
#   make run   → plain go run (any OS)
#   make build → build a binary into ./tmp/
#   make tidy  → go mod tidy

ifeq ($(OS),Windows_NT)
	AIR_CONFIG := .air/windows.toml
	BIN := ./tmp/main.exe
else
	AIR_CONFIG := .air/linux.toml
	BIN := ./tmp/main
endif

.PHONY: dev run build tidy clean

dev:
	air -c $(AIR_CONFIG)

run:
	go run main.go

build:
	go build -o $(BIN) .

tidy:
	go mod tidy

clean:
	rm -rf tmp
