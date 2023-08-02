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
	rm -rf ../../../website/tables/jira
	go run main.go doc ../../../website/tables/jira
	sed 's_(\(.*\))_(../../../../../website/tables/jira/\1)_' ../../../website/tables/jira/README.md > ./docs/tables/README.md
	sed -i.bak -e 's_(\(.*\).md)_(tables/\1)_' ../../../website/tables/jira/README.md
	mv ../../../website/tables/jira/README.md ../../../website/pages/docs/plugins/sources/jira/tables.md
	sed -i.bak -e 's_(\(.*\).md)_(\1)_' ../../../website/tables/jira/*.md
	rm -rf ../../../website/tables/jira/*.bak

# All gen targets
.PHONY: gen
gen: gen-docs