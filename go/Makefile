CURRENT_YEAR=23
GOCMD=go
GOTEST=$(GOCMD) test ./cmd/year$(CURRENT_YEAR)/...
BINARY_NAME=goAOC

build:
	mkdir -p build/bin
	$(GOCMD) build -o build/bin/$(BINARY_NAME)

run:
	$(GOCMD) run ./cmd/main.go $(day) $(year)

test:
	$(GOTEST)

testday:
	$(GOCMD) test ./cmd/year$(CURRENT_YEAR)/day$(day)/
