
GOPATH:=$(shell go env GOPATH)

.PHONY: proto build build-linux image
proto:
	protoc -I=${GOPATH}/src/github.com/gogo/protobuf/protobuf --proto_path=${GOPATH}/src:. --micro_out=. --gofast_out=. api/*.proto
	python ${GOPATH}/src/knowledgemap_backend/remove_tag.py ./api/*.pb.go
.PHONY: build
build: proto
	go build -o passport cmd/main.go
build-linux: proto
	GOOS=linux GOARCH=amd64 go build -o passport-linux cmd/main.go
image: build-linux
	 docker build -t passport:latest .
clean:
	rm -f passport passport-linux
