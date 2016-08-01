GO := GO15VENDOREXPERIMENT=1 go
GOPATH := $(GOPATH)
GOROOT := $(GOROOT)
test:
	sudo -s
	export GOPATH=$(GOPATH)
	export GOROOT=$(GOROOT)
	export PATH=$(PATH):$(GOROOT)/bin
	$(GO) test -v github.com/mozilla/libaudit-go
