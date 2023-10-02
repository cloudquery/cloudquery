.PHONY: test
test:
	go test -race -timeout 3m ./...

.PHONY: build
build:
	go build -o gcp

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
	@command -v cloudquery >/dev/null 2>&1 || { \
		echo "Error: 'cloudquery' command not found. Please install it before running gen-docs."; \
		echo "You can install it by following the instructions at: https://www.cloudquery.io/docs/quickstart"; \
		exit 1; \
	}
	cloudquery tables --format json --output-dir ./docs test/config.yml
	go run scripts/policy_docs/main.go gcp policies ../../../website/pages/docs/plugins/sources/gcp/policies.md docs/gcp/__tables.json
	rm docs/gcp/__tables.json

	rm -rf ../../../website/tables/gcp
	cloudquery tables --format markdown --output-dir ../../../website/tables test/config.yml
	go run scripts/example_queries/main.go policies/queries ../../../website/tables/gcp
	sed 's_(\(.*\))_(../../../../../website/tables/gcp/\1)_' ../../../website/tables/gcp/README.md > ./docs/tables/README.md
	sed -i.bak -e 's_(\(.*\).md)_(tables/\1)_' ../../../website/tables/gcp/README.md
	mv ../../../website/tables/gcp/README.md ../../../website/pages/docs/plugins/sources/gcp/tables.md
	sed -i.bak -e 's_(\(.*\).md)_(\1)_' ../../../website/tables/gcp/*.md
	rm -rf ../../../website/tables/gcp/*.bak

.PHONY: gen-services
gen-services:
	go run codegen/main.go

.PHONY: gen-spec-schema
gen-spec-schema:
	go run client/spec/gen/main.go

# All gen targets
.PHONY: gen
gen: gen-spec-schema gen-docs
