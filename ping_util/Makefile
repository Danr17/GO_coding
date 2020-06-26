ROOT_DIR:=$(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))

BINARY=ping_util
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
	$(foreach GOARCH, $(ARCHITECTURES), $(shell export GOOS=$(GOOS); export GOARCH=$(GOARCH); go build -v -o releases/$(BINARY)-$(GOOS)-$(GOARCH))))

# Remove only what we've created
clean:
	find ${ROOT_DIR}/releases/ -name '${BINARY}[-?][a-zA-Z0-9]*[-?][a-zA-Z0-9]*' -delete

tar:
	mv releases/ping_util-windows-386 releases/ping_util-windows-386.exe
	mv releases/ping_util-windows-amd64 releases/ping_util-windows-amd64.exe
	cp index.html releases/index.html
	cp example.txt releases/example.txt
	tar -czvf releases/ping_util-windows-386.tar.gz releases/ping_util-windows-386.exe releases/index.html releases/example.txt
	tar -czvf releases/ping_util-windows-amd64.tar.gz releases/ping_util-windows-amd64.exe releases/index.html releases/example.txt
	tar -czvf releases/ping_util-darwin-386.tar.gz releases/ping_util-darwin-386 releases/index.html releases/example.txt
	tar -czvf releases/ping_util-darwin-amd64.tar.gz releases/ping_util-darwin-amd64 releases/index.html releases/example.txt
	tar -czvf releases/ping_util-linux-386.tar.gz releases/ping_util-linux-386 releases/index.html releases/example.txt
	tar -czvf releases/ping_util-linux-amd64.tar.gz releases/ping_util-linux-amd64 releases/index.html releases/example.txt
	
tar_clean:
	rm -rf releases/ping_util-windows-386.tar.gz 
	rm -rf releases/ping_util-windows-amd64.tar.gz 
	rm -rf releases/ping_util-darwin-386.tar.gz  
	rm -rf releases/ping_util-darwin-amd64.tar.gz 
	rm -rf releases/ping_util-linux-386.tar.gz
	rm -rf releases/ping_util-linux-amd64.tar.gz 

.PHONY: check clean build_all all
