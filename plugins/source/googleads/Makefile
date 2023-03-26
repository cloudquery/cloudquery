# Test unit
.PHONY: test
test:
	go test -race -timeout 3m ./...

.PHONY: lint
lint:
	golangci-lint run --config ../../.golangci.yml

.PHONY: gen-docs
gen-docs:
	rm -rf ../../../website/tables/googleads
	go run main.go doc ../../../website/tables/googleads
	sed 's_(\(.*\))_(../../../../../website/tables/googleads/\1)_' ../../../website/tables/googleads/README.md > ./docs/tables/README.md
	sed -i.bak -e 's_(\(.*\).md)_(tables/\1)_' ../../../website/tables/googleads/README.md
	mv ../../../website/tables/googleads/README.md ../../../website/pages/docs/plugins/sources/googleads/tables.md
	sed -i.bak -e 's_(\(.*\).md)_(\1)_' ../../../website/tables/googleads/*.md
	rm -rf ../../../website/tables/googleads/*.bak

# All gen targets
.PHONY: gen
gen: gen-docs
