# Generate mocks for mock/unit testing 
.PHONY: gen-mocks
gen-mocks: install-tools
	go generate ./...

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
	go run main.go doc --format json docs
	go run scripts/policy_docs/main.go k8s policies ../../../website/pages/docs/plugins/sources/k8s/policies.md docs/__tables.json
	rm docs/__tables.json
	
	rm -rf ../../../website/tables/k8s
	go run main.go doc ../../../website/tables/k8s
	sed -i.bak -e 's_(\(.*\).md)_(tables/\1)_' ../../../website/tables/k8s/README.md
	mv ../../../website/tables/k8s/README.md ../../../website/pages/docs/plugins/sources/k8s/tables.md
	sed -i.bak -e 's_(\(.*\).md)_(\1)_' ../../../website/tables/k8s/*.md
	rm -rf ../../../website/tables/k8s/*.bak

.PHONY: lint
lint:
	golangci-lint run --config ../../.golangci.yml 

# All gen targets
.PHONY: gen
gen: gen-mocks gen-docs
