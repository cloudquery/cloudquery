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
	go run main.go doc --format json docs
	go run scripts/policy_docs/main.go gcp policies ../../../website/pages/docs/plugins/sources/gcp/policies.md docs/__tables.json
	rm docs/__tables.json

	rm -rf ../../../website/tables/gcp
	go run main.go doc ../../../website/tables/gcp
	cp ../../../website/tables/gcp/README.md ./docs/tables/README.md
	sed -i.bak -e 's_(\(.*\).md)_(tables/\1)_' ../../../website/tables/gcp/README.md
	mv ../../../website/tables/gcp/README.md ../../../website/pages/docs/plugins/sources/gcp/tables.md
	sed -i.bak -e 's_(\(.*\).md)_(\1)_' ../../../website/tables/gcp/*.md
	rm -rf ../../../website/tables/gcp/*.bak

.PHONY: gen-services
gen-services:
	go run codegen/main.go

# All gen targets
.PHONY: gen
gen: gen-docs
