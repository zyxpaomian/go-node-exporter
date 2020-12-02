.PHONY: compile run

PROJECTNAME=$(shell basename "$(PWD)")

# Go related variables.
GOBASE=$(shell pwd)
GOPATH :=
ifeq ($(OS),Windows_NT)
	GOPATH = $(CURDIR)/vendor;$(CURDIR)
else
	GOPATH = $(CURDIR)/vendor:$(CURDIR)
endif
export GOPATH

GOBIN=$(GOBASE)/bin
GOFILES=$(wildcard *.go)

# Redirect error output to a file, so we can show it in development mode.
STDERR=/tmp/.$(PROJECTNAME)-stderr.txt

# PID file will store the server process id when it's running on development mode
PID=/tmp/.$(PROJECTNAME).pid

# Make is verbose in Linux. Make it silent.
MAKEFLAGS += --silent

tidy:
	go env -w GO111MODULE=on; go mod tidy

vendor:
	go env -w GO111MODULE=on; go mod vendor

compile:
	@for GOOS in linux windows;do\
		for GOARCH in amd64; do\
			GOOS=$${GOOS} GOARCH=$${GOARCH} go build -v -o  bin/$(PROJECTNAME).$${GOOS}-$${GOARCH} -mod=vendor; \
		done ;\
	done

run:
	go run main.go
