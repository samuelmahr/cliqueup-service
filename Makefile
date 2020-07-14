GOFILES = $(shell find . -name '*.go')

default: build

workdir:
	mkdir -p workdir

build: workdir/cliqueup

build-native: $(GOFILES)
	go build -o workdir/cliqueup-contacts .

workdir/cliqueup: $(GOFILES)
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o workdir/cliqueup .