test:
	pytest .

build-docker:
	docker build -t square:latest .

fmt:
	black . --exclude=cloudquery

fmt-check:
	black --check . --exclude=cloudquery

gen-docs: build-docker
	@command -v cloudquery >/dev/null 2>&1 || { \
		echo "Error: 'cloudquery' command not found. Please install it before running gen-docs."; \
		echo "You can install it by following the instructions at: https://www.cloudquery.io/docs/quickstart"; \
		exit 1; \
	}
	rm -rf ../../../website/tables/square
	cloudquery tables --format markdown --output-dir ../../../website/tables testdata/config.yml
	sed -i.bak -e 's_(\(.*\).md)_(tables/\1)_' ../../../website/tables/square/README.md
	mv ../../../website/tables/square/README.md ../../../website/pages/docs/plugins/sources/square/tables.md
	sed -i.bak -e 's_(\(.*\).md)_(\1)_' ../../../website/tables/square/*.md
	rm -rf ../../../website/tables/square/*.bak

gen: gen-docs
