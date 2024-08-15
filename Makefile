include help.mk

# get root dir
ROOT_DIR := $(dir $(realpath $(lastword $(MAKEFILE_LIST))))
DOCKER_IMAGE_NAME := ai-assistant
UI_NAME := ${DOCKER_IMAGE_NAME}.exe

.DEFAULT_GOAL := start

.PHONY: update
update: ## pulls git repo
	@git -C ${ROOT_DIR} pull

.PHONY: test
test: ## run go test
	@go test ${ROOT_DIR}...

.PHONY: build-ui
build-ui: ## builds the ui
ifeq ($(DETECT_OS),Windows)
	go build -o ${UI_NAME} -ldflags "-H windowsgui" ${ROOT_DIR}main.go
else
	go build -o ${UI_NAME} ${ROOT_DIR}main.go
endif

.PHONY: start
start: build-ui ## start app
	${ROOT_DIR}${UI_NAME}

.PHONY: docker-build
docker-build:
	docker build -f ${ROOT_DIR}Dockerfile -t ${DOCKER_IMAGE_NAME} .

.PHONY: docker-test
docker-test: docker-build ## runs tests in docker
	docker run -v "${ROOT_DIR}:/app" --rm ${DOCKER_IMAGE_NAME} go test -v ./... -covermode=count -coverprofile=coverage.out

.PHONY: docker-build-ui-linux
docker-build-ui-linux: docker-build ## build a linux binary in docker
	docker run -v "${ROOT_DIR}:/app" --rm ${DOCKER_IMAGE_NAME} make build-ui
