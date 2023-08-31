test:
	pytest .

fmt:
	black .

fmt-check:
	black --check .

build-docker:
	docker build -t typeform:latest .

gen-docs: build-docker
	@command -v cloudquery >/dev/null 2>&1 || { \
		echo "Error: 'cloudquery' command not found. Please install it before running gen-docs."; \
		echo "You can install it by following the instructions at: https://www.cloudquery.io/docs/quickstart"; \
		exit 1; \
	}
	rm -rf ../../../website/tables/typeform
	cloudquery tables --format markdown --output-dir ../../../website/tables testdata/config.yml
	sed -i.bak -e 's_(\(.*\).md)_(tables/\1)_' ../../../website/tables/typeform/README.md
	mv ../../../website/tables/typeform/README.md ../../../website/pages/docs/plugins/sources/typeform/tables.md
	sed -i.bak -e 's_(\(.*\).md)_(\1)_' ../../../website/tables/typeform/*.md
	rm -rf ../../../website/tables/typeform/*.bak

gen: gen-docs
