# default target
.PHONY: help
help:
	@echo "usage: make <target>"
	@echo
	@echo "  where <target> is one of the following"
	@echo
	@echo "    init        to initialize the module (one-off action)"
	@echo "    clean       to delete a compiled binary (if any)"
	@echo "    get         to fetch all package dependencies"
	@echo "    test        to run all tests"
	@echo "    build       to compile binary for linux amd64 architecture"
	@echo "    all         to run all targets but init"
	@echo
	@echo "    help        to show this text"

.PHONY: clean
clean:
	rm -f ukeapi

.PHONY: init
init:
	go mod init

.PHONY: get
get:
	go mod tidy

.PHONY: test
test:
	go test

.PHONY: build
build:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build

.PHONY: all
all: clean get test build