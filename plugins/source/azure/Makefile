# Test unit
.PHONY: test
test:
	go test -race -timeout 3m ./...

.PHONY: build
build:
	go build -o azure

.PHONY: gen-docs
gen-docs: build
	@command -v cloudquery >/dev/null 2>&1 || { \
		echo "Error: 'cloudquery' command not found. Please install it before running gen-docs."; \
		echo "You can install it by following the instructions at: https://www.cloudquery.io/docs/quickstart"; \
		exit 1; \
	}
	cloudquery tables --format json --output-dir ./docs test/config.yml
	go run scripts/policy_docs/main.go azure policies ../../../website/pages/docs/plugins/sources/azure/policies.md docs/azure/__tables.json
	rm docs/azure/__tables.json

	rm -rf ../../../website/tables/azure
	cloudquery tables --format markdown --output-dir ../../../website/tables test/config.yml
	go run scripts/example_queries/main.go policies/queries ../../../website/tables/azure
	sed 's_(\(.*\))_(../../../../../website/tables/azure/\1)_' ../../../website/tables/azure/README.md > ./docs/tables/README.md
	sed -i.bak -e 's_(\(.*\).md)_(tables/\1)_' ../../../website/tables/azure/README.md
	mv ../../../website/tables/azure/README.md ../../../website/pages/docs/plugins/sources/azure/tables.md
	sed -i.bak -e 's_(\(.*\).md)_(\1)_' ../../../website/tables/azure/*.md
	rm -rf ../../../website/tables/azure/*.bak

.PHONY: lint
lint:
	golangci-lint run --config ../../.golangci.yml 

.PHONY: gen
gen: gen-docs
	echo "Done"
