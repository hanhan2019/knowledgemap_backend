
GOPATH:=$(shell go env GOPATH)


.PHONY: proto
proto:
	protoc -I=${GOPATH}/src/github.com/gogo/protobuf/protobuf --proto_path=${GOPATH}/src:. --micro_out=. --gofast_out=. api/{{ .ServiceName}}.proto

.PHONY: build
build: proto
	go build -o {{ .ServiceName }} cmd/main.go
