# Test unit
.PHONY: test
test: gen-fixtures
	wget -nc https://raw.githubusercontent.com/stripe/openapi/master/openapi/spec3.json -P ./resources/testdata
	go test -race -timeout 3m ./...

.PHONY: gen-docs
gen-docs:
	rm -rf ./docs/tables/*
	go run main.go doc ./docs/tables
	sed 's_(\(.*\))_(https://github.com/cloudquery/cloudquery/blob/main/plugins/source/stripe/docs/tables/\1)_' docs/tables/README.md > ../../../website/pages/docs/plugins/sources/stripe/tables.md

.PHONY: lint
lint:
	golangci-lint run --config ../../.golangci.yml 

.PHONY: gen-fixtures
gen-fixtures:
	rm -vf ./resources/testdata/fixtures_gen.json
	go run codegen/fixtures/main.go

.PHONY: gen-code
gen-code:
	go run codegen/main.go

# All gen targets
.PHONY: gen
gen: gen-fixtures gen-docs
