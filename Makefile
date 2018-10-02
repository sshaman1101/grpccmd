OUT="protoc-gen-grpccmd"

all: build install

build:
	@echo "+ $@"
	go build -tags 'nocgo' -o ${OUT} cmd/protoc-gen-grpccmd/main.go

install:
	@echo "+ $@"
	cp ${OUT} ${GOPATH}/bin/
