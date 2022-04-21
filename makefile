prefix?=$(GOPATH)/bin
root?=$(shell pwd)
name?=$(notdir $(root))
srcdir?=$(root)
builddir?=$(root)/bin

ldflags?=-ldflags "-s -w"

mansrc?=$(root)/README.md

ifeq ($(shell uname),Linux)
	mandir?=/usr/local/man/man1
endif
ifeq ($(shell uname),Darwin)
	mandir?=/usr/local/share/man/man1
endif
ifeq ($(OS),Windows_NT)
	ext=.exe
endif

.PHONY: tidy build man install pre post

default: all

all: $(prefix) pre tidy install post

pre:
	mkdir -p $(builddir)
post:
	rm -rf $(builddir)

tidy: $(shell which go) $($(root)/go.mod)
	go mod tidy

build: $(prefix) pre $(builddir)
	go build $(ldflags) -o $(builddir)/$(name)$(ext) $(srcdir)
man: $(shell which pandoc) $(mandir) $(builddir)
	@pandoc $(mansrc) -s -t man -o $(builddir)/$(name).1

install: build man
	mv $(builddir)/$(name)$(ext) $(prefix)/
	mv $(builddir)/$(name).1 $(mandir)/
uninstall:
	rm -f $(prefix)/$(name)$(ext)
	rm -f $(mandir)/$(name).1
