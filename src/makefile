export BINARY_NAME = "palindrome-service"

# global command
.PHONY: all
all: dependencies build test

.PHONY: dependencies
dependencies:
	go mod download

.PHONY: build
build:
	go build -o ../bin/${BINARY_NAME}


.PHONY: start
start:
	../bin/${BINARY_NAME}


packages = \
	./data\
	./server\

.PHONY: test
test:
	@$(foreach package,$(packages), \
    		set -e; \
    		go test $(package);)