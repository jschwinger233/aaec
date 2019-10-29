test:
	golangci-lint run

build: test
	GO111MODULE=on go mod download
	GO111MODULE=on go build
