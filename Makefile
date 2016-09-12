all: fmt build test

fmt:
	go fmt ./...

build:
	go build .

testdeps:
	go get github.com/stretchr/testify/assert

test: testdeps
	go test ./...
