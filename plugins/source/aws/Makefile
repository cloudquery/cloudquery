# Test unit
.PHONY: test
test:
	go test -race -timeout 3m ./...

# Install tools
.PHONY: install-tools
install-tools:
	@echo Installing tools from tools/tool.go
	@cat tools/tool.go | grep _ | awk -F'"' '{print $$2}' | xargs -tI % go install %

.PHONY: lint
lint:
	golangci-lint run --config ../../.golangci.yml

.PHONY: gen-mocks
gen-mocks:
	go install github.com/golang/mock/mockgen@v1.6.0
	rm -rf ./client/mocks/*
	go generate ./client/...

.PHONY: gen-docs
gen-docs:
	rm -rf ./docs/tables/*
	go run main.go doc ./docs/tables
	sed 's_(\(.*\))_(https://github.com/cloudquery/cloudquery/blob/main/plugins/source/aws/docs/tables/\1)_' docs/tables/README.md > ../../../website/pages/docs/plugins/sources/aws/tables.md
	go run main.go doc --format json docs
	go run scripts/policy_docs/main.go aws policies ../../../website/pages/docs/plugins/sources/aws/policies.md docs/__tables.json
	rm docs/__tables.json

# All gen targets
.PHONY: gen
gen: gen-mocks gen-docs
