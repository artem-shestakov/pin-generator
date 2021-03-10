GO111MODULE=off
swagger_install:
	go get -u github.com/go-swagger/go-swagger/cmd/swagger
swagger:
	swagger generate spec -o ./api/swagger.yml --scan-models

build:
	go build -v ./cmd/apiserver

.DEFAULT_GOAL := build