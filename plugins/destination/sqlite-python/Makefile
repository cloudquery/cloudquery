current_dir=$(shell basename $(CURDIR))

test:
	pytest .

fmt:
	black .

fmt-check:
	black --check .

build-docker:
	docker build -t sqlite-python:latest .
