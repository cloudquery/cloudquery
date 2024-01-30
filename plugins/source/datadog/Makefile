# Generate mocks for mock/unit testing
.PHONY: gen-mocks
gen-mocks:
	go install github.com/golang/mock/mockgen@v1.7.0-rc.1
	go run codegen/main.go
	go generate ./client/services/...

.PHONY: build
build:
	go build

# Test unit
.PHONY: test
test:
	go test -race -timeout 3m ./...

.PHONY: gen-docs
gen-docs: build
	@command -v cloudquery >/dev/null 2>&1 || { \
		echo "Error: 'cloudquery' command not found. Please install it before running gen-docs."; \
		echo "You can install it by following the instructions at: https://www.cloudquery.io/docs/quickstart"; \
		exit 1; \
	}
	rm -rf docs/tables
	cloudquery tables --format markdown --output-dir docs test/config.yml
	mv docs/$(shell basename $(CURDIR)) docs/tables

.PHONY: lint
lint:
	golangci-lint run --config ../../.golangci.yml 

.PHONY: gen-spec-schema
gen-spec-schema:
	go run client/spec/gen/main.go

# All gen targets
.PHONY: gen
gen: gen-mocks gen-spec-schema gen-docs
