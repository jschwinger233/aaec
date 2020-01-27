test:
	golangci-lint run --build-tags golangci

deps:
	GO111MODULE=on go mod download
	GO111MODULE=on go mod vendor

build: test deps
	GO111MODULE=on go build
