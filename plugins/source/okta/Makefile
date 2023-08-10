.PHONY: build
build:
	go build

# Test unit
.PHONY: test
test:
	go test -race -timeout 3m ./...

.PHONY: lint
lint:
	golangci-lint run --config ../../.golangci.yml 

.PHONY: gen-docs
gen-docs: build
	@command -v cloudquery >/dev/null 2>&1 || { \
		echo "Error: 'cloudquery' command not found. Please install it before running gen-docs."; \
		echo "You can install it by following the instructions at: https://www.cloudquery.io/docs/quickstart"; \
		exit 1; \
	}
	rm -rf ../../../website/tables/okta
	cloudquery tables --format markdown --output-dir ../../../website/tables test/config.yml
	sed 's_(\(.*\))_(../../../../../website/tables/okta/\1)_' ../../../website/tables/okta/README.md > ./docs/tables/README.md
	sed -i.bak -e 's_(\(.*\).md)_(tables/\1)_' ../../../website/tables/okta/README.md
	mv ../../../website/tables/okta/README.md ../../../website/pages/docs/plugins/sources/okta/tables.md
	sed -i.bak -e 's_(\(.*\).md)_(\1)_' ../../../website/tables/okta/*.md
	rm -rf ../../../website/tables/okta/*.bak

# All gen targets
.PHONY: gen
gen: gen-docs
