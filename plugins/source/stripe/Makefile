# Test unit
.PHONY: test
test: gen-fixtures
	wget -nc https://raw.githubusercontent.com/stripe/openapi/master/openapi/spec3.json -P ./resources/testdata
	go test -race -timeout 3m ./...

.PHONY: gen-docs
gen-docs:
	rm -rf ../../../website/tables/stripe
	go run main.go doc ../../../website/tables/stripe
	sed -i.bak -e 's_(\(.*\).md)_(tables/\1)_' ../../../website/tables/stripe/README.md
	mv ../../../website/tables/stripe/README.md ../../../website/pages/docs/plugins/sources/stripe/tables.md
	sed -i.bak -e 's_(\(.*\).md)_(\1)_' ../../../website/tables/stripe/*.md
	rm -rf ../../../website/tables/stripe/*.bak

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
