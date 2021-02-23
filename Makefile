NAME=really
BINDIR=bin
VERSION=$(shell git describe --tags || echo "unknownversion")
LDFLAGS="-s -w -X github.com/Bpazy/really.buildVer=${VERSION}"
GOBUILD=go build -ldflags=${LDFLAGS}
CMDPATH=./cmd/really

all: linux-amd64 darwin-amd64 windows-amd64 # Most used

init:
	export GOPROXY=https://mirrors.aliyun.com/goproxy/

darwin-amd64: init
	GOARCH=amd64 GOOS=darwin $(GOBUILD) -o $(BINDIR)/$(NAME)-$@-$(VERSION) $(CMDPATH)

linux-amd64: init
	GOARCH=amd64 GOOS=linux $(GOBUILD) -o $(BINDIR)/$(NAME)-$@-$(VERSION) $(CMDPATH)

windows-amd64: init
	GOARCH=amd64 GOOS=windows $(GOBUILD) -o $(BINDIR)/$(NAME)-$@-$(VERSION).exe $(CMDPATH)

install: init
	go install -ldflags=${LDFLAGS} $(CMDPATH)

clean:
	rm $(BINDIR)/*
