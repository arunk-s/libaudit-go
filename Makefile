GOPATH := $(GOPATH)
GOROOT := $(GOROOT)
GO := GO15VENDOREXPERIMENT=1 $(GOROOT)/bin/go
test:
	sudo -E GOROOT=$(GOROOT) GOPATH=$(GOPATH) $(GO) test -v .
