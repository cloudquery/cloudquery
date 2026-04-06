current_dir=$(shell basename $(CURDIR))

test:
	uv run pytest .

fmt:
	uv run black .

fmt-check:
	uv run black --check .

build-docker:
	docker build -t sqlite-python:latest .

coverage:
	# TODO: add coverage
	exit 0
