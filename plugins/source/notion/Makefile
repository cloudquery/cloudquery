.PHONY: test
test:
	go test -timeout 3m ./...

.PHONY: build
build:
	go build -o notion

.PHONY: lint
lint:
	@golangci-lint run --timeout 10m

PHONY: gen-docs
gen-docs: build
	@command -v cloudquery >/dev/null 2>&1 || { \
		echo "Error: 'cloudquery' command not found. Please install it before running gen-docs."; \
		echo "You can install it by following the instructions at: https://www.cloudquery.io/docs/quickstart"; \
		exit 1; \
	}
	rm -rf docs/tables
	rm -rf ../../../website/tables/notion
	cloudquery tables --format markdown --output-dir ../../../website/tables test/config.yml
	cp -r ../../../website/tables/notion docs/tables
	rm ../../../website/tables/notion/README.md
	sed -i.bak -e 's_(\(.*\).md)_(\1)_' ../../../website/tables/notion/*.md
	rm -rf ../../../website/tables/notion/*.bak

# All gen targets
.PHONY: gen
gen: gen-docs
