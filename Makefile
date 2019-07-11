GOCMD=go
GOTEST=$(GOCMD) test -v
GORACE=$(GOTEST) -race

all: test race

.PHONY: test
test:
	$(GOTEST) -cover ./...

.PHONY: race
race:
	$(GORACE) ./...
