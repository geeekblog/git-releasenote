#VERSION ?= $(shell git describe)
VERSION ?= 0.1.0
GO_VERSION ?= $(shell go version)
BUILD_TIME ?= $(shell date "+%F %T")
.PHONY: build
test:

lint:

build:
	go build -ldflags "-X 'git-releasenote/cmd/sub_cmd/version.appVersion=$(VERSION)' \
	-X 'git-releasenote/cmd/sub_cmd/version.buildTime=$(BUILD_TIME)' \
	-X 'git-releasenote/cmd/sub_cmd/version.goVersion=$(GO_VERSION)'" \
	-o bin/git-releasenote git-releasenote/cmd/git-releasenote

release:
	make build
	cp -R config bin
