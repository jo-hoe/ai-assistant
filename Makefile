include help.mk

# get root dir
ROOT_DIR := $(dir $(realpath $(lastword $(MAKEFILE_LIST))))
UI_NAME := ai-assistant.exe

.DEFAULT_GOAL := start

.PHONY: update
update: ## pulls git repo
	@git -C ${ROOT_DIR} pull

.PHONY: test
test: ## test service
	@go test ${ROOT_DIR}...

.PHONY: build-ui
build-ui: ## builds the ui
ifeq ($(DETECT_OS),Windows)
	go build -o ${UI_NAME} -ldflags "-H windowsgui" ${ROOT_DIR}main.go
else
	go build -o ${UI_NAME} ${ROOT_DIR}main.go
endif

.PHONY: start
start: build-ui # start app
	${ROOT_DIR}${UI_NAME}