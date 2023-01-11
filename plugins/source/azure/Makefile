# Test unit
.PHONY: test
test:
	go test -race -timeout 3m ./...

.PHONY: gen-docs
gen-docs:
	rm -rf ./docs/tables/*
	go run main.go doc ./docs/tables
	sed 's_(\(.*\))_(https://github.com/cloudquery/cloudquery/blob/main/plugins/source/azure/docs/tables/\1)_' docs/tables/README.md > ../../../website/pages/docs/plugins/sources/azure/tables.md

.PHONY: lint
lint:
	golangci-lint run --config ../../.golangci.yml 

.PHONY: gen
gen: gen-docs
	echo "Done"
