current_dir=$(shell basename $(CURDIR))

test:
	pytest .

fmt:
	black .

fmt-check:
	black --check .

build-docker:
	docker build -t sqlite-python:latest .

gen-docs: build-docker
	@command -v cloudquery >/dev/null 2>&1 || { \
		echo "Error: 'cloudquery' command not found. Please install it before running gen-docs."; \
		echo "You can install it by following the instructions at: https://www.cloudquery.io/docs/quickstart"; \
		exit 1; \
	}
	rm -rf ../../../website/pages/docs/plugins/sources/$(current_dir)/tables
	rm -rf docs/tables
	mkdir -p docs/tables
	cloudquery tables --format markdown --output-dir ../../../website/pages/docs/plugins/sources/$(current_dir) testdata/config.yml
	cp -r ../../../website/pages/docs/plugins/sources/$(current_dir)/tables/ docs/tables
	sed -i.bak -e 's_(\(.*\).md)_(tables/\1)_' ../../../website/pages/docs/plugins/sources/$(current_dir)/tables/README.md
	mv ../../../website/pages/docs/plugins/sources/$(current_dir)/tables/README.md ../../../website/pages/docs/plugins/sources/$(current_dir)/tables.md
	sed -i.bak -e 's_(\(.*\).md)_(\1)_' ../../../website/pages/docs/plugins/sources/$(current_dir)/tables/*.md
	rm -rf ../../../website/pages/docs/plugins/sources/$(current_dir)/tables/*.bak

gen: gen-docs
