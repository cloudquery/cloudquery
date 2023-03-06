# Generate mocks for mock/unit testing
.PHONY: gen-mocks
gen-mocks:
	go install github.com/golang/mock/mockgen@v1.6.0
	rm -rf ./client/mocks/*
	go generate ./client/...

# Test unit
.PHONY: test
test:
	go test -race -timeout 3m ./...

# Install tools
.PHONY: install-tools
install-tools:
	@echo Installing tools from tools/tool.go
	@cat tools/tool.go | grep _ | awk -F'"' '{print $$2}' | xargs -tI % go install %

.PHONY: gen-docs
gen-docs:
	rm -rf ../../../website/tables/github
	go run main.go doc ../../../website/tables/github
	sed -i.bak -e 's_(\(.*\).md)_(tables/\1)_' ../../../website/tables/github/README.md
	mv ../../../website/tables/github/README.md ../../../website/pages/docs/plugins/sources/github/tables.md
	sed -i.bak -e 's_(\(.*\).md)_(\1)_' ../../../website/tables/github/*.md
	rm -rf ../../../website/tables/github/*.bak

.PHONY: lint
lint:
	golangci-lint run --config ../../.golangci.yml 

# All gen targets
.PHONY: gen
gen: gen-mocks gen-docs