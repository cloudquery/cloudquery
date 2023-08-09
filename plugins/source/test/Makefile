# Test unit
.PHONY: test
test:
	go test -race -timeout 3m ./...

.PHONY: build
build:
	go build

# Install tools
.PHONY: install-tools
install-tools:
	@echo Installing tools from tools/tool.go
	@cat tools/tool.go | grep _ | awk -F'"' '{print $$2}' | xargs -tI % go install %

.PHONY: lint
lint:
	golangci-lint run --config ../../.golangci.yml

.PHONY: gen-docs
gen-docs: build
	rm -rf ./docs/tables/*
	cloudquery tables --output-dir docs/tables testdata

# All gen targets
.PHONY: gen
gen: gen-docs
