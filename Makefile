GOCMD=go
GOTEST=$(GOCMD) test
GORACE=$(GOTEST) -race

all: test race

.PHONY: test
test:
	$(GOTEST) -v ./...

.PHONY: race
race:
	$(GORACE) -v ./...
