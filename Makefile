VERSION := $(shell echo $(shell git describe --tags))
COMMIT := $(shell git log -1 --format='%H')
APP_NAME := "bcnad"
UNAME_OS := $(shell uname -s)
UNAME_ARCH := $(shell uname -m)
name := "bcnad"

all: build

build:
	$(eval override VERSION = $(shell git describe --tags))
	$(eval override DATE = $(shell date))
	$(eval override HOST = $(shell hostname))
	$(eval override COMMIT = $(shell git log -1 --format='%H'))
	GO111MODULE=on go build -o bin/bcnad ./cmd/bcnad 

win64:
	GOOS=windows GOARCH=amd64 GO111MODULE=on  go build -o bin/main-windows-amd64  ./cmd/bcnad

win32:
	GOOS=windows GOARCH=386 GO111MODULE=on  go build -o bin/main-windows-386  ./cmd/bcnad

mac:
	GOOS=darwin GOARCH=amd64 GO111MODULE=on go build -o bin/main-darwin-amd64 ./cmd/bcnad
