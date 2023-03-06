# Unit tests
.PHONY: test
test:
	go test -race -timeout 3m ./...

.PHONY: lint
lint:
	golangci-lint run --config ../../.golangci.yml

.PHONY: gen-docs
gen-docs:
	rm -rf ../../../website/tables/snyk
	go run main.go doc ../../../website/tables/snyk
	cp ../../../website/tables/snyk/README.md ./docs/tables/README.md
	sed -i.bak -e 's_(\(.*\).md)_(tables/\1)_' ../../../website/tables/snyk/README.md
	mv ../../../website/tables/snyk/README.md ../../../website/pages/docs/plugins/sources/snyk/tables.md
	sed -i.bak -e 's_(\(.*\).md)_(\1)_' ../../../website/tables/snyk/*.md
	rm -rf ../../../website/tables/snyk/*.bak

# All gen targets
.PHONY: gen
gen: gen-docs
