# default target
.PHONY: help
help:
	@echo "usage: make <target>"
	@echo
	@echo "  where <target> is one of the following"
	@echo
	@echo "    clean       to delete a compiled binary (if any)"
	@echo "    get         to fetch all package dependencies"
	@echo "    test        to run all tests"
	@echo "    build       to compile binary for local machine architecture"
	@echo "    all         to run all targets"
	@echo
	@echo "    help        to show this text"

.PHONY: clean
clean:
	rm -f ukeapi

.PHONY: get
get:
	go mod tidy

.PHONY: test
test:
	go test

.PHONY: build
build:
	go build

.PHONY: all
all: clean get test build