ROOT_DIR:=$(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))

BINARY=tcp_ping
VERSION=0.1.1
BUILD=`git rev-parse HEAD`
PLATFORMS=darwin linux windows
ARCHITECTURES=386 amd64

# Setup linker flags option for build that interoperate with variable names in src code
LDFLAGS=-ldflags "-X main.Version=${VERSION} -X main.Build=${BUILD}"

default: build

all: clean build_all 

build:
	go build ${LDFLAGS} -o ${BINARY}

build_all:
	$(foreach GOOS, $(PLATFORMS),\
	$(foreach GOARCH, $(ARCHITECTURES), $(shell export GOOS=$(GOOS); export GOARCH=$(GOARCH); go build -v -o $(BINARY)-$(GOOS)-$(GOARCH))))

# Remove only what we've created
clean:
	find ${ROOT_DIR} -name '${BINARY}[-?][a-zA-Z0-9]*[-?][a-zA-Z0-9]*' -delete

tar:
	mv tcp_ping-windows-386 tcp_ping-windows-386.exe
	mv tcp_ping-windows-amd64 tcp_ping-windows-amd64.exe
	tar -czvf tcp_ping-windows-386.tar.gz tcp_ping-windows-386.exe index.html example.txt
	tar -czvf tcp_ping-windows-amd64.tar.gz tcp_ping-windows-amd64.exe index.html example.txt
	tar -czvf tcp_ping-darwin-386.tar.gz tcp_ping-darwin-386 index.html example.txt
	tar -czvf tcp_ping-darwin-amd64.tar.gz tcp_ping-darwin-amd64 index.html example.txt
	tar -czvf tcp_ping-linux-386.tar.gz tcp_ping-linux-386 index.html example.txt
	tar -czvf tcp_ping-linux-amd64.tar.gz tcp_ping-linux-amd64 index.html example.txt
	
tar_clean:
	rm -rf tcp_ping-windows-386.tar.gz 
	rm -rf tcp_ping-windows-amd64.tar.gz 
	rm -rf tcp_ping-darwin-386.tar.gz  
	rm -rf tcp_ping-darwin-amd64.tar.gz 
	rm -rf tcp_ping-linux-386.tar.gz
	rm -rf tcp_ping-linux-amd64.tar.gz 

.PHONY: check clean build_all all
