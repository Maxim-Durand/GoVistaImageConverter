# Beware Makefile requires using Tabs

# Fill in your default here, setting from command line will override
TARGET_FOLDER ?=$(shell pwd)
REPO_ROOT ?=$(shell git rev-parse --show-toplevel)

EXE_NAME = "imageConverter.exe"

build: go.mod
	go build $(TARGET_FOLDER)

build_for_windows_vista:
	GOOS=windows GOARCH=386 CGO_ENABLED=1 CC=i686-w64-mingw32-gcc go build -o $(EXE_NAME) $(TARGET_FOLDER)
	
run:
	go run $(TARGET_FOLDER)
