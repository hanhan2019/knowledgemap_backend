cd /Users/firewinggames/code/gopath/src/knowledgemap_backend/microservices/knowledgemap/class
make build-linux
cd /Users/firewinggames/code/gopath/src/knowledgemap_backend/microservices/knowledgemap/passport
make build-linux
cd /Users/firewinggames/code/gopath/src/knowledgemap_backend/microservices/knowledgemap/user
make build-linux
cd /Users/firewinggames/code/gopath/src/knowledgemap_backend/microservices/knowledgemap/question
make build-linux
cd /Users/firewinggames/code/gopath/src/knowledgemap_backend/microservices/knowledgemap/knowledgemap
make build-linux
cd /Users/firewinggames/code/gopath/src/knowledgemap_backend/agent/knowledgemap
env GOOS=linux GOARCH=amd64 go build -o agent-linux cmd/knowledgemap.go