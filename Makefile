SHELL 		:=$(shell which bash)
.SHELLFLAGS :=-c

.PHONY: build-dependencies
build-dependencies:
	echo ">> Building dependencies..."
	go get -u github.com/rakyll/gotest; \
	go get -d -v ./...;

### TEST
.PHONY: test
test: build-dependencies
	gotest -v ./... -count=1;

