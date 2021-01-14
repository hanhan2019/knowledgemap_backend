
GOPATH:=$(shell go env GOPATH)

.PHONY: proto
proto:
	protoc -I=${GOPATH}/src/github.com/gogo/protobuf/protobuf --proto_path=${GOPATH}/src:. --micro_out=. --gofast_out=. api/*.proto
	python ${GOPATH}/src/knowledgemap_backend/remove_tag.py ./api/*.pb.go
.PHONY: build
build: proto
	go build -o knowledge cmd/main.go
build-linux:
	GOOS=linux GOARCH=amd64 go build -o knowledge-linux cmd/main.go


image: build-linux
	 docker build -t knowledge:latest .
clean:
	rm -f knowledge knowledge-linux
