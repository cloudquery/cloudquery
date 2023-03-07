# Test unit
.PHONY: test
test:
	go test -race -timeout 3m ./...

.PHONY: lint
lint:
	golangci-lint run --config ../../.golangci.yml

.PHONY: gen-code
gen-code:
	go run codegen/*.go

.PHONY: gen-docs
gen-docs:
	rm -rf ../../../website/tables/oracle
	go run main.go doc ../../../website/tables/oracle
	cp ../../../website/tables/oracle/README.md ./docs/tables/README.md
	sed -i.bak -e 's_(\(.*\).md)_(tables/\1)_' ../../../website/tables/oracle/README.md
	mv ../../../website/tables/oracle/README.md ../../../website/pages/docs/plugins/sources/oracle/tables.md
	sed -i.bak -e 's_(\(.*\).md)_(\1)_' ../../../website/tables/oracle/*.md
	rm -rf ../../../website/tables/oracle/*.bak

# All gen targets
.PHONY: gen
gen: gen-docs