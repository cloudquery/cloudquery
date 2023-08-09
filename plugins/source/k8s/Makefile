# Generate mocks for mock/unit testing 
.PHONY: gen-mocks
gen-mocks: install-tools
	go generate ./...

.PHONY: build
build:
	go build

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
gen-docs: build
	@command -v cloudquery >/dev/null 2>&1 || { \
		echo "Error: 'cloudquery' command not found. Please install it before running gen-docs."; \
		echo "You can install it by following the instructions at: https://www.cloudquery.io/docs/quickstart"; \
		exit 1; \
	}
	cloudquery tables --format json --output-dir ./docs test/config.yml
	go run scripts/policy_docs/main.go k8s policies ../../../website/pages/docs/plugins/sources/k8s/policies.md docs/k8s/__tables.json
	rm docs/k8s/__tables.json
	
	rm -rf ../../../website/tables/k8s
	cloudquery tables --format markdown --output-dir ../../../website/tables test/config.yml
	go run scripts/example_queries/main.go policies/queries ../../../website/tables/k8s
	sed 's_(\(.*\))_(../../../../../website/tables/k8s/\1)_' ../../../website/tables/k8s/README.md > ./docs/tables/README.md
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
