# Generate mocks for mock/unit testing 
.PHONY: generate-mocks
generate-mocks:
	go generate ./client/services/...

# Test unit
.PHONY: test
test:
	go test -timeout 3m ./...

# Install tools
.PHONY: install-tools
install-tools:
	@echo Installing tools from tools/tool.go
	@cat tools/tool.go | grep _ | awk -F'"' '{print $$2}' | xargs -tI % go install %

.PHONY: gen-docs
gen-docs:
	rm -rf ./docs/tables/*
	go run main.go doc ./docs/tables
	sed 's_(\(.*\))_(https://github.com/cloudquery/cloudquery/blob/main/plugins/source/heroku/docs/tables/\1)_' docs/tables/README.md > ../../../website/pages/docs/plugins/sources/heroku/tables.md

.PHONY: lint
lint:
	golangci-lint run --config ../../.golangci.yml

.PHONY: gen-code
gen-code:
	go run codegen/main.go

# All gen targets
.PHONY: gen
gen: gen-code gen-docs