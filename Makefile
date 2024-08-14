include help.mk

# get root dir
ROOT_DIR := $(dir $(realpath $(lastword $(MAKEFILE_LIST))))

.DEFAULT_GOAL := build-ui

.PHONY: update
update: ## pulls git repo
	@git -C ${ROOT_DIR} pull

.PHONY: test
test: ## test service
	@go test ${ROOT_DIR}...

.PHONY: build-ui
build-ui: ## builds the ui
ifeq ($(DETECT_OS),Windows)
	go build -ldflags "-H windowsgui" ${ROOT_DIR}app/ui/main.go
else
	go build ${ROOT_DIR}app/ui/main.go
endif