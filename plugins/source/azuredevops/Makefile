# Test unit
.PHONY: test
test:
	go test -race -timeout 3m ./...

.PHONY: lint
lint:
	golangci-lint run --config ../../.golangci.yml

.PHONY: gen-docs
gen-docs:
	rm -rf ./docs/tables/*
	go run main.go doc ./docs/tables
	sed 's_(\(.*\))_(https://github.com/cloudquery/cloudquery/blob/main/plugins/source/azuredevops/docs/tables/\1)_' docs/tables/README.md > ../../../website/pages/docs/plugins/sources/azuredevops/tables.md

.PHONY: gen
gen: gen-docs