.PHONY: build
build:
	go build

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
	rm -rf docs/tables
	rm -rf ../../../website/tables/awspricing
	cloudquery tables --format markdown --output-dir ../../../website/tables test/config.yml
	cp -r ../../../website/tables/awspricing docs/tables
	rm ../../../website/tables/awspricing/README.md
	sed -i.bak -e 's_(\(.*\).md)_(\1)_' ../../../website/tables/awspricing/*.md
	rm -rf ../../../website/tables/awspricing/*.bak

# All gen targets
.PHONY: gen
gen: gen-docs