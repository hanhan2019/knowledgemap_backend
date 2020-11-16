
GOPATH:=$(shell go env GOPATH)

.PHONY: proto build build-linux image
proto:
	protoc -I=${GOPATH}/src/github.com/gogo/protobuf/protobuf --proto_path=${GOPATH}/src:. --micro_out=. --gofast_out=. api/*.proto
.PHONY: build
build: proto
	go build -o question cmd/main.go
build-linux: proto
	GOOS=linux GOARCH=amd64 go build -o question-linux cmd/main.go
image: build-linux
	 docker build -t question:latest .
clean:
	rm -f question question-linux
