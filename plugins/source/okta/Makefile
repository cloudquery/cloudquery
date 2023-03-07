# Test unit
.PHONY: test
test:
	go test -race -timeout 3m ./...

.PHONY: lint
lint:
	golangci-lint run --config ../../.golangci.yml 

.PHONY: gen-docs
gen-docs:
	rm -rf ../../../website/tables/okta
	go run main.go doc ../../../website/tables/okta
	cp ../../../website/tables/okta/README.md ./docs/tables/README.md
	sed -i.bak -e 's_(\(.*\).md)_(tables/\1)_' ../../../website/tables/okta/README.md
	mv ../../../website/tables/okta/README.md ../../../website/pages/docs/plugins/sources/okta/tables.md
	sed -i.bak -e 's_(\(.*\).md)_(\1)_' ../../../website/tables/okta/*.md
	rm -rf ../../../website/tables/okta/*.bak

# All gen targets
.PHONY: gen
gen: gen-docs
