# Test unit
.PHONY: test
test:
	go test -race -timeout 3m ./...

.PHONY: lint
lint:
	golangci-lint run --config ../../.golangci.yml


.PHONY: gen-docs
gen-docs:
	rm -rf ../../../website/tables/plausible
	go run main.go doc ../../../website/tables/plausible
	sed -i.bak -e 's_(\(.*\).md)_(tables/\1)_' ../../../website/tables/plausible/README.md
	mv ../../../website/tables/plausible/README.md ../../../website/pages/docs/plugins/sources/plausible/tables.md
	sed -i.bak -e 's_(\(.*\).md)_(\1)_' ../../../website/tables/plausible/*.md
	rm -rf ../../../website/tables/plausible/*.bak

# All gen targets
.PHONY: gen
gen: gen-docs
