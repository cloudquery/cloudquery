# Test unit
.PHONY: test
test:
	go test -timeout 3m ./...

.PHONY: lint
lint:
	golangci-lint run --config ../../.golangci.yml

.PHONY: gen-code
gen-code:
	go run codegen/*.go

.PHONY: gen-docs
gen-docs:
	rm -rf ../../../website/tables/hubspot
	go run main.go doc ../../../website/tables/hubspot
	sed 's_(\(.*\))_(../../../../../website/tables/hubspot/\1)_' ../../../website/tables/hubspot/README.md > ./docs/tables/README.md
	sed -i.bak -e 's_(\(.*\).md)_(tables/\1)_' ../../../website/tables/hubspot/README.md
	mv ../../../website/tables/hubspot/README.md ../../../website/pages/docs/plugins/sources/hubspot/tables.md
	sed -i.bak -e 's_(\(.*\).md)_(\1)_' ../../../website/tables/hubspot/*.md
	rm -rf ../../../website/tables/hubspot/*.bak

# All gen targets
.PHONY: gen
gen: gen-docs