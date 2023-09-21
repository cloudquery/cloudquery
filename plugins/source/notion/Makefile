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
	rm -rf ../../../website/tables/notion
	cloudquery tables --format markdown --output-dir ../../../website/tables test/config.yml
	sed 's_(\(.*\))_(../../../../../website/tables/notion/\1)_' ../../../website/tables/notion/README.md > ./docs/tables/README.md
	sed -i.bak -e 's_(\(.*\).md)_(tables/\1)_' ../../../website/tables/notion/README.md
	mkdir --parents ../../../website/pages/docs/plugins/sources/notion/
	mv ../../../website/tables/notion/README.md ../../../website/pages/docs/plugins/sources/notion/tables.md
	sed -i.bak -e 's_(\(.*\).md)_(\1)_' ../../../website/tables/notion/*.md
	rm -rf ../../../website/tables/notion/*.bak

# All gen targets
.PHONY: gen
gen: gen-docs
