# Test unit
.PHONY: test
test:
	go test -race -timeout 3m ./...

.PHONY: lint
lint:
	golangci-lint run --config ../../.golangci.yml


.PHONY: gen-docs
gen-docs:
	rm -rf ../../../website/tables/salesforce
	go run main.go doc ../../../website/tables/salesforce
	sed -i.bak -e 's_(\(.*\).md)_(tables/\1)_' ../../../website/tables/salesforce/README.md
	mv ../../../website/tables/salesforce/README.md ../../../website/pages/docs/plugins/sources/salesforce/tables.md
	sed -i.bak -e 's_(\(.*\).md)_(\1)_' ../../../website/tables/salesforce/*.md
	rm -rf ../../../website/tables/salesforce/*.bak

# All gen targets
.PHONY: gen
gen: gen-docs
