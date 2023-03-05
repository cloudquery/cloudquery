# Generate mocks for mock/unit testing 
.PHONY: gen-mocks
gen-mocks:
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
	rm -rf ../../../website/tables/gandi
	go run main.go doc ../../../website/tables/gandi
	sed -i.bak -e 's_(\(.*\).md)_(tables/\1)_' ../../../website/tables/gandi/README.md
	mv ../../../website/tables/gandi/README.md ../../../website/pages/docs/plugins/sources/gandi/tables.md
	sed -i.bak -e 's_(\(.*\).md)_(\1)_' ../../../website/tables/gandi/*.md
	rm -rf ../../../website/tables/gandi/*.bak

.PHONY: lint
lint:
	golangci-lint run --config ../../.golangci.yml 

# All gen targets
.PHONY: gen
gen: gen-docs
