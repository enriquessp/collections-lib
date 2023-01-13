BIN=collections-lib

NEXT_VERSION=$(shell git describe --tags --always | awk -F "-" '{print $$1}' | awk -F "." '{print $$1"."$$2"."($$3 + 1)}')
VERSION := $(shell git describe --tags --always | awk -F "-" '{print $$1}')

OS := $(if $(GOOS),$(GOOS),$(shell go env GOOS))
ARCH := $(if $(GOARCH),$(GOARCH),$(shell go env GOARCH))

all: build

deps:
	go get ./...

build:
	go generate
	mkdir -p bin/
	GOARCH="${ARCH}" GOOS="${OS}" GO111MODULE=on go build -o bin/${BIN} *.go

test: build
	go vet ./...
	go test ./...
	
test-bench:
	go test -run=XXX -bench=. ./...

TEST_NAME="Test_CommissionInfo"
run-debug-test:
	dlv test --build-flags='github.com/enriquessp/collections-lib' -- -test.run ^${TEST_NAME}$

clean:
	rm -rf bin generated

release: build
	git tag -a ${NEXT_VERSION} -m "creating new tag"
	git push origin ${NEXT_VERSION}
	git push

