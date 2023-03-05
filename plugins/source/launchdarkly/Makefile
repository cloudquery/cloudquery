# Test unit
.PHONY: test
test:
	go test -race -timeout 3m ./...

.PHONY: gen-docs
gen-docs:
	rm -rf ../../../website/tables/launchdarkly
	go run main.go doc ../../../website/tables/launchdarkly
	sed -i.bak -e 's_(\(.*\).md)_(tables/\1)_' ../../../website/tables/launchdarkly/README.md
	mv ../../../website/tables/launchdarkly/README.md ../../../website/pages/docs/plugins/sources/launchdarkly/tables.md
	sed -i.bak -e 's_(\(.*\).md)_(\1)_' ../../../website/tables/launchdarkly/*.md
	rm -rf ../../../website/tables/launchdarkly/*.bak

.PHONY: lint
lint:
	golangci-lint run --config ../../.golangci.yml 

# All gen targets
.PHONY: gen
gen: gen-docs
