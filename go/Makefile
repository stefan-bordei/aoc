GOCMD=go
GOTEST=$(GOCMD) test
BINARY_NAME=goAOC

build:
	mkdir -p build/bin
	$(GOCMD) build -o build/bin/$(BINARY_NAME)

run:
	$(GOCMD) run ./cmd/main.go $(day) $(year)
