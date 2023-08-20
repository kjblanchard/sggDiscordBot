.PHONY: build publish stop run package frontend bindir api apackage abuild
BINARY_FOLDER_NAME = bin
BINARY_NAME = SupergoonDiscordBot

all: build run

build: bindir
	@go build -o $(BINARY_FOLDER_NAME)/$(BINARY_NAME)
bindir:
		mkdir -p $(BINARY_FOLDER_NAME)
clean:
	@rm -rf `find . -type d -name bin`
rebuild: clean bindir build
run:
	@./$(BINARY_FOLDER_NAME)/$(BINARY_NAME)
