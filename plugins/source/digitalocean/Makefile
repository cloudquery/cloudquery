# Generate mocks for mock/unit testing 
.PHONY: generate-mocks
generate-mocks:
	go generate ./client/services/...

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

.PHONY: gen-docs
gen-docs:
	rm -rf ../../../website/tables/digitalocean
	go run main.go doc ../../../website/tables/digitalocean
	sed -i.bak -e 's_(\(.*\).md)_(tables/\1)_' ../../../website/tables/digitalocean/README.md
	mv ../../../website/tables/digitalocean/README.md ../../../website/pages/docs/plugins/sources/digitalocean/tables.md
	sed -i.bak -e 's_(\(.*\).md)_(\1)_' ../../../website/tables/digitalocean/*.md
	rm -rf ../../../website/tables/digitalocean/*.bak

# All gen targets
.PHONY: gen
gen: gen-docs