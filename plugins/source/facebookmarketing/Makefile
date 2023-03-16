.PHONY: test
test:
	go test -timeout 3m ./...

.PHONY: lint
lint:
	golangci-lint run --config ../../.golangci.yml

.PHONY: gen-docs
gen-docs:
	go run main.go doc --format json docs
	rm docs/__tables.json

	rm -rf ../../../website/tables/facebookmarketing
	go run main.go doc ../../../website/tables/facebookmarketing
	sed -i.bak -e 's_(\(.*\).md)_(tables/\1)_' ../../../website/tables/facebookmarketing/README.md
	mv ../../../website/tables/facebookmarketing/README.md ../../../website/pages/docs/plugins/sources/facebookmarketing/tables.md
	sed -i.bak -e 's_(\(.*\).md)_(\1)_' ../../../website/tables/facebookmarketing/*.md
	rm -rf ../../../website/tables/facebookmarketing/*.bak

# All gen targets
.PHONY: gen
gen: gen-docs