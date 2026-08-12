.PHONY: build run clean docker publish

DOCKER_IMAGE_OWNER = enf3rno
DOCKER_IMAGE_NAME = supergoon-discord-bot
DOCKER_IMAGE_VERSION = 18
DOCKER_IMAGE_FULL = $(DOCKER_IMAGE_OWNER)/$(DOCKER_IMAGE_NAME):$(DOCKER_IMAGE_VERSION)

all: test run

test:
	@pyright .

run:
	@python main.py


clean:
	@find . -type d -name __pycache__ -exec rm -rf {} + 2>/dev/null; true

docker:
	@docker image build --platform linux/amd64 -f ./Dockerfile -t $(DOCKER_IMAGE_FULL) .

publish:
	@docker login
	@docker image push $(DOCKER_IMAGE_FULL)

update: test clean docker publish

