##DEFINE
GO_LDFLAGS=-ldflags " -w"
VERSION=1.0.0
BIN_PATH=./build/output/${VERSION}
BASE_NAME=insur-box

WORKSPACE="/go/src"

clean: 
	@rm -rf ${BIN_PATH}/*

build: build-imagecode

build-imagecode: 
	go build ${GO_LDFLAGS} -o ${BIN_PATH}/${BASE_NAME}-imagecode src/main.go

build-docker:
	docker run --rm -v `PWD`:${WORKSPACE}/${BASE_NAME} -w ${WORKSPACE}/${BASE_NAME} -it pujielan/golang:1.10.3 go build ${GO_LDFLAGS} -o ${BIN_PATH}/${BASE_NAME}-imagecode src/main.go 

build-img:
	docker build -t pujielan/imgcode:v1.0.0 .