.PHONY: build publish stop run package frontend bindir api apackage abuild
BINARY_FOLDER_NAME = bin
BINARY_NAME = SupergoonDiscordBot
PACKAGE_FILE_NAME = SupergoonDiscordBot
DOCKER_IMAGE_OWNER = enf3rno
DOCKER_IMAGE_NAME = supergoon-discord-bot
DOCKER_IMAGE_VERSION = 5
DOCKER_IMAGE_FULL = $(DOCKER_IMAGE_OWNER)/$(DOCKER_IMAGE_NAME):$(DOCKER_IMAGE_VERSION)

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
package: clean bindir
	@tar -czvf $(BINARY_FOLDER_NAME)/$(PACKAGE_FILE_NAME).tgz `find . -name "*.go"`
docker: package
	@docker image build -f ./Dockerfile -t $(DOCKER_IMAGE_FULL) .
publish:
	@docker login
	@docker image push $(DOCKER_IMAGE_FULL)
