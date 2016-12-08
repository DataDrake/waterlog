.RECIPEPREFIX != ps

include Makefile.waterlog

GOCC     = go

GOPATH   = $(shell pwd)/build
GOBIN    = build/bin
GOSRC    = build/src
PROJROOT = $(GOSRC)/github.com/DataDrake

DESTDIR ?=
PREFIX  ?= /usr
BINDIR   = $(PREFIX)/bin

all: build

build: setup
    @$(call stage,BUILD)
    $(GOCC) install -v ./build/src/github.com/DataDrake/waterlog/...
    @$(call pass,BUILD)

setup:
    @$(call stage,SETUP)
    @$(call task,Setting up GOPATH)
    @mkdir -p $(GOPATH)
    @$(call task,Setting up src/)
    @mkdir -p $(GOSRC)
    @$(call task,Setting up project root)
    @mkdir -p $(PROJROOT)
    @$(call task,Setting up symlinks)
    @if [ ! -d $(PROJROOT)/waterlog ]; then ln -s $(shell pwd) $(PROJROOT)/waterlog; fi
    @$(call pass,SETUP)

validate: golint-setup
    @$(call stage,FORMAT)
    @$(GOCC) fmt -x ./...
    @$(call pass,FORMAT)
    @$(call stage,VET)
    @$(GOCC) vet -x ./...
    @$(call pass,VET)
    @$(call stage,LINT)
    @$(GOBIN)/golint -set_exit_status ./...
    @$(call pass,LINT)

golint-setup:
    @if [ ! -e $(GOBIN)/golint ]; then \
        printf "Installing golint...\n"; \
        $(GOCC) get -u github.com/golang/lint/golint; \
        rm -rf $(GOPATH)/src/golang.org $(GOPATH)/src/github.com/golang $(GOPATH)/pkg; \
    fi

install:
    @$(call stage,INSTALL)
    install -D -m 00755 $(GOBIN)/waterlog $(DESTDIR)$(BINDIR)/waterlog
    @$(call pass,INSTALL)

uninstall:
    @$(call stage,UNINSTALL)
    rm -f $(DESTDIR)$(BINDIR)/waterlog
    @$(call pass,UNINSTALL)

clean:
    @$(call stage,CLEAN)
    @$(call task,Removing symlinks)
    @unlink $(PROJROOT)/waterlog
    @$(call task,Removing build directory)
    @rm -r build
    @$(call pass,CLEAN)
