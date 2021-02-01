VERSION ?= "newest"
GO_VERSION ?= $(shell go version)
BUILD_TIME ?= $(shell date "+%F %T")

export GO111MODULE := on
.PHONY: clean test lint build_darwin build_win build_linux release release_darwin release_linux release_win copy_template

clean:
	@rm -rf bin/git-releasenote

test:
	@go test ./...

lint:
	@golangci-lint run

build_darwin:
	@GOOS=darwin GOARCH=amd64 \
	go build -ldflags "-X 'git-releasenote/cmd/sub_cmd/version.appVersion=$(VERSION)' \
	-X 'git-releasenote/cmd/sub_cmd/version.buildTime=$(BUILD_TIME)' \
	-X 'git-releasenote/cmd/sub_cmd/version.goVersion=$(GO_VERSION)'" \
	-o bin/git-releasenote/git-releasenote git-releasenote/cmd/git-releasenote

build_win:
	@GOOS=windows GOARCH=amd64 \
	go build -ldflags "-X 'git-releasenote/cmd/sub_cmd/version.appVersion=$(VERSION)' \
	-X 'git-releasenote/cmd/sub_cmd/version.buildTime=$(BUILD_TIME)' \
	-X 'git-releasenote/cmd/sub_cmd/version.goVersion=$(GO_VERSION)'" \
	-o bin/git-releasenote/git-releasenote.exe git-releasenote/cmd/git-releasenote

build_linux:
	@GOOS=linux GOARCH=amd64 \
	go build -ldflags "-X 'git-releasenote/cmd/sub_cmd/version.appVersion=$(VERSION)' \
	-X 'git-releasenote/cmd/sub_cmd/version.buildTime=$(BUILD_TIME)' \
	-X 'git-releasenote/cmd/sub_cmd/version.goVersion=$(GO_VERSION)'" \
	-o bin/git-releasenote/git-releasenote git-releasenote/cmd/git-releasenote

release_darwin: build_darwin copy_template
	@OS=darwin make package
	@make clean

release_win: build_win copy_template
	@OS=windows make package
	@make clean

release_linux: build_linux copy_template
	@OS=linux make package
	@make clean

package:
	@cd bin && tar czf ./git-releasenote-$(OS)-$(VERSION).tar.gz ./git-releasenote

copy_template:
	@cp -R config bin/git-releasenote

release: release_darwin release_win release_linux
