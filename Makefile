.PHONY: pull docs

IMAGE = grafana/docs-base:latest
CONTENT_PATH = /hugo/content/docs/writers-toolkit/
PORT = 3002:3002

pull:
	docker pull $(IMAGE)

docs:
	docker run -v $(shell pwd)/docs/sources:$(CONTENT_PATH):Z -p $(PORT) --rm -it $(IMAGE) /bin/bash -c "make server"