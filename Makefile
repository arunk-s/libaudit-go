GOPATH := $(GOPATH)
GOROOT := $(GOROOT)
GO := $(GOROOT)/bin/go
test:
	sudo -E GOROOT=$(GOROOT) GOPATH=$(GOPATH) $(GO) test -v .
