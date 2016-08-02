GOPATH := $(GOPATH)
GOROOT := $(GOROOT)
GO := /usr/local/go/bin/go
	sudo -E GOROOT=$(GOROOT) GOPATH=$(GOPATH) $(GO) test -v .
