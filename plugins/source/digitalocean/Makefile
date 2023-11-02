# Generate mocks for mock/unit testing 
.PHONY: generate-mocks
generate-mocks:
	go install github.com/golang/mock/mockgen@v1.6.0
	go generate ./client/services.go

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
	rm -rf docs/tables
	rm -rf ../../../website/tables/digitalocean
	cloudquery tables --format markdown --output-dir ../../../website/tables test/config.yml
	cp -r ../../../website/tables/digitalocean docs/tables
	rm ../../../website/tables/digitalocean/README.md
	sed -i.bak -e 's_(\(.*\).md)_(\1)_' ../../../website/tables/digitalocean/*.md
	rm -rf ../../../website/tables/digitalocean/*.bak

# All gen targets
.PHONY: gen
gen: generate-mocks gen-docs