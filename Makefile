SHELL := /bin/bash
CMD_PKG := github.com/Xuanwo/bard/cmd/bard
VERSION := $(shell cat ./constants/version.go | grep "Version\ =" | sed -e s/^.*\ //g | sed -e s/\"//g)
GO_BUILD_OPTION := -trimpath -tags netgo

.PHONY: all check format vet lint build install uninstall release clean test

help:
	@echo "Please use \`make <target>\` where <target> is one of"
	@echo "  check      to format, vet and lint "
	@echo "  build      to create bin directory and build bard"
	@echo "  install    to install bard to /usr/local/bin/bard"
	@echo "  uninstall  to uninstall bard"
	@echo "  release    to release bard"
	@echo "  clean      to clean build and test files"
	@echo "  test       to run test"

check: format vet lint

format:
	@echo "go fmt"
	@go fmt ./...
	@echo "ok"

vet:
	@echo "go vet"
	@go vet ./...
	@echo "ok"

lint:
	@echo "golint"
	@golint ./...
	@echo "ok"

build: tidy check
	@echo "build bard"
	@mkdir -p ./bin
	@go build ${GO_BUILD_OPTION} -race -o ./bin/bard ${CMD_PKG}
	@echo "ok"

install: build
	@echo "install bard to GOPATH"
	@cp ./bin/bard ${GOPATH}/bin/bard
	@echo "ok"

release:
	@echo "release bard"
	@-rm ./release/*
	@mkdir -p ./release

	@echo "build for linux"
	@GOOS=linux GOARCH=amd64 go build ${GO_BUILD_OPTION} -o ./bin/linux/bard_v${VERSION}_linux_amd64 ${CMD_PKG}
	@tar -C ./bin/linux/ -czf ./release/bard_v${VERSION}_linux_amd64.tar.gz bard_v${VERSION}_linux_amd64

	@echo "build for macOS"
	@GOOS=darwin GOARCH=amd64 go build ${GO_BUILD_OPTION} -o ./bin/macos/bard_v${VERSION}_macos_amd64 ${CMD_PKG}
	@tar -C ./bin/macos/ -czf ./release/bard_v${VERSION}_macos_amd64.tar.gz bard_v${VERSION}_macos_amd64

	@echo "build for windows"
	@GOOS=windows GOARCH=amd64 go build ${GO_BUILD_OPTION} -o ./bin/windows/bard_v${VERSION}_windows_amd64.exe ${CMD_PKG}
	@tar -C ./bin/windows/ -czf ./release/bard_v${VERSION}_windows_amd64.tar.gz bard_v${VERSION}_windows_amd64.exe

	@echo "ok"

clean:
	@rm -rf ./bin
	@rm -rf ./release
	@rm -rf ./coverage

test:
	@echo "run test"
	@go test -race -coverprofile=coverage.txt -covermode=atomic -v ./...
	@go tool cover -html="coverage.txt" -o "coverage.html"
	@echo "ok"

tidy:
	@echo "Tidy and check the go mod files"
	@go mod tidy
	@go mod verify
	@echo "Done"
