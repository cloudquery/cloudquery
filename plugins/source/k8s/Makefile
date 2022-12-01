# Generate mocks for mock/unit testing 
.PHONY: gen-mocks
gen-mocks: install-tools
	go generate ./...

# Test unit
.PHONY: test
test:
	go test -timeout 3m ./...

# Install pre-commit hooks. This requires pre-commit to be installed (https://pre-commit.com/)
.PHONY: install-hooks
install-hooks:
	pre-commit install

# Install tools
.PHONY: install-tools
install-tools:
	@echo Installing tools from tools/tool.go
	@cat tools/tool.go | grep _ | awk -F'"' '{print $$2}' | xargs -tI % go install %

.PHONY: gen-docs
gen-docs:
	rm -rf ./docs/tables/*
	go run main.go doc ./docs/tables
	sed 's_(\(.*\))_(https://github.com/cloudquery/cloudquery/blob/main/plugins/source/k8s/docs/tables/\1)_' docs/tables/README.md > ../../../website/pages/docs/plugins/sources/k8s/tables.md

.PHONY: lint
lint:
	golangci-lint run --config ../../.golangci.yml 

.PHONY: gen-code
gen-code:
	go run codegen/main.go

# All gen targets
.PHONY: gen
gen: gen-code gen-mocks gen-docs
