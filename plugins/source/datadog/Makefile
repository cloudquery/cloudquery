# Generate mocks for mock/unit testing
.PHONY: generate-mocks
generate-mocks:
	go install github.com/golang/mock/mockgen@v1.7.0-rc.1
	go generate ./client/services/...

# Test unit
.PHONY: test
test:
	go test -race -timeout 3m ./...

.PHONY: gen-docs
gen-docs:
	rm -rf ./docs/tables/*
	go run main.go doc ./docs/tables
	sed 's_(\(.*\))_(https://github.com/cloudquery/cloudquery/blob/main/plugins/source/datadog/docs/tables/\1)_' docs/tables/README.md > ../../../website/pages/docs/plugins/sources/datadog/tables.md

.PHONY: lint
lint:
	golangci-lint run --config ../../.golangci.yml 

# All gen targets
.PHONY: gen
gen: generate-mocks gen-docs