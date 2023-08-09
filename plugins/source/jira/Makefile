.PHONY: build
build:
	go build

# Test unit
.PHONY: test
test:
	go test -timeout 3m ./...

.PHONY: lint
lint:
	golangci-lint run --config ../../.golangci.yml

.PHONY: gen-code
gen-code:
	go run codegen/*.go

.PHONY: gen-docs
gen-docs: build
	@command -v cloudquery >/dev/null 2>&1 || { \
		echo "Error: 'cloudquery' command not found. Please install it before running gen-docs."; \
		echo "You can install it by following the instructions at: https://www.cloudquery.io/docs/quickstart"; \
		exit 1; \
	}
	rm -rf ../../../website/tables/jira
	cloudquery tables --format markdown --output-dir ../../../website/tables test/config.yml
	sed -i.bak -e 's_(\(.*\).md)_(tables/\1)_' ../../../website/tables/jira/README.md
	mv ../../../website/tables/jira/README.md ../../../website/pages/docs/plugins/sources/jira/tables.md
	sed -i.bak -e 's_(\(.*\).md)_(\1)_' ../../../website/tables/jira/*.md
	rm -rf ../../../website/tables/jira/*.bak

# All gen targets
.PHONY: gen
gen: gen-docs