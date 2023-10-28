# Beware Makefile requires using Tabs

# Fill in your default here, setting from command line will override
TARGET_FOLDER ?=$(shell pwd)
REPO_ROOT ?=$(shell git rev-parse --show-toplevel)

build: go.mod
	go build ./cmd/image_converter

run:
	go run cmd/image_converter/main.go
