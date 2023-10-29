# Beware Makefile requires using Tabs

# Fill in your default here, setting from command line will override
TARGET_FOLDER ?=$(shell pwd)
REPO_ROOT ?=$(shell git rev-parse --show-toplevel)


EXE_NAME ?= "imageConverter"
WIN_EXE_NAME ?= "$(EXE_NAME).exe"

build: go.mod
	go build $(REPO_ROOT)

build_for_windows:
	GOOS=windows GOARCH=386 go build -o $(WIN_EXE_NAME) $(REPO_ROOT)

build_for_windows_vista:
	GOOS=windows GOARCH=386 CGO_ENABLED=1 CC=i686-w64-mingw32-gcc go build -o $(WIN_EXE_NAME) $(REPO_ROOT)

run:
	go run $(REPO_ROOT)

tests:
	cd $(REPO_ROOT)/test && go test