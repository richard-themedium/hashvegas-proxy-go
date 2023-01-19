
PROJECT_NAME ?= bc-backend

all: build


.PHONY: build
build:
	@echo "Building $(PROJECT_NAME)"
	GOOS=$(GOOS) GOARCH=$(GOARCH) go build -o $(PROJECT_NAME)

