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

.PHONY: gen-mocks
gen-mocks:
	go install github.com/golang/mock/mockgen@v1.6.0
	rm -rf ./client/mocks/*
	go generate ./client/...

.PHONY: gen-docs
gen-docs: build
	@command -v cloudquery >/dev/null 2>&1 || { \
		echo "Error: 'cloudquery' command not found. Please install it before running gen-docs."; \
		echo "You can install it by following the instructions at: https://www.cloudquery.io/docs/quickstart"; \
		exit 1; \
	}
	rm -rf ../../../website/tables/pagerduty
	cloudquery tables --format markdown --output-dir ../../../website/tables test/config.yml
	sed 's_(\(.*\))_(../../../../../website/tables/pagerduty/\1)_' ../../../website/tables/pagerduty/README.md > ./docs/tables/README.md
	sed -i.bak -e 's_(\(.*\).md)_(tables/\1)_' ../../../website/tables/pagerduty/README.md
	mv ../../../website/tables/pagerduty/README.md ../../../website/pages/docs/plugins/sources/pagerduty/tables.md
	sed -i.bak -e 's_(\(.*\).md)_(\1)_' ../../../website/tables/pagerduty/*.md
	rm -rf ../../../website/tables/pagerduty/*.bak

# All gen targets
.PHONY: gen
gen: gen-mocks gen-docs
