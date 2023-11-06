PROJECT_DIR=$(shell realpath .)
PROJECT_NAME=$(shell basename "$(PROJECT_DIR)")
DIST_DIR="$(PROJECT_DIR)/dist"
MAIN_FILE="main.go"

default: build

build: build-clean build-lin build-win
# build-mac

build-clean:
	rm -rfv $(DIST_DIR)/*
	-mkdir $(DIST_DIR)

# Linux
build-lin: build-lin-amd64
# build-lin-i386 build-lin-arm64 build-lin-arm

build-lin-i386:
	GOOS=linux GOARCH=386 go build -o $(DIST_DIR)/$(PROJECT_NAME)-linux-i386 $(MAIN_FILE)

build-lin-amd64:
	-GOOS=linux GOARCH=amd64 go build -o $(DIST_DIR)/$(PROJECT_NAME)-linux-amd64 $(MAIN_FILE)

build-lin-arm:
	GOOS=linux GOARCH=arm go build -o $(DIST_DIR)/$(PROJECT_NAME)-linux-arm $(MAIN_FILE)

build-lin-arm64:
	GOOS=linux GOARCH=arm64 go build -o $(DIST_DIR)/$(PROJECT_NAME)-linux-arm64 $(MAIN_FILE)

# Windows
build-win: build-win-i386 build-win-amd64 build-win-arm build-win-arm64

build-win-i386:
	GOOS=windows GOARCH=386 go build -o $(DIST_DIR)/$(PROJECT_NAME)-windows-i386.exe $(MAIN_FILE)

build-win-amd64:
	GOOS=windows GOARCH=amd64 go build -o $(DIST_DIR)/$(PROJECT_NAME)-windows-amd64.exe $(MAIN_FILE)

build-win-arm:
	GOOS=windows GOARCH=arm go build -o $(DIST_DIR)/$(PROJECT_NAME)-windows-arm.exe $(MAIN_FILE)

build-win-arm64:
	GOOS=windows GOARCH=arm64 go build -o $(DIST_DIR)/$(PROJECT_NAME)-windows-arm64.exe $(MAIN_FILE)

# Mac
build-mac:
	build-mac-amd64 build-mac-arm64

build-mac-amd64:
	GOOS=darwin GOARCH=amd64 go build -o $(DIST_DIR)/$(PROJECT_NAME)-darwin-amd64 $(MAIN_FILE)

build-mac-arm64:
	GOOS=darwin GOARCH=arm64 go build -o $(DIST_DIR)/$(PROJECT_NAME)-darwin-arm64 $(MAIN_FILE)


# ios/amd64
# ios/arm64

# android/386
# android/amd64
# android/arm
# android/arm64
