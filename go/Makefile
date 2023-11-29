DEFAULT_YEAR=22
GOCMD=go
GOTEST=$(GOCMD) test ./cmd/year$(DEFAULT_YEAR)/...
BINARY_NAME=goAOC

build:
	mkdir -p build/bin
	$(GOCMD) build -o build/bin/$(BINARY_NAME)

run:
	$(GOCMD) run ./cmd/main.go $(day) $(year)

test:
	$(GOTEST)
