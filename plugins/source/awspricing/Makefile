.PHONY: test
test:
	go test -race -timeout 3m ./...

.PHONY: lint
lint:
	golangci-lint run --config ../../.golangci.yml 

.PHONY: gen-docs
gen-docs:
	rm -rf ../../../website/tables/awspricing
	go run main.go doc ../../../website/tables/awspricing
	sed 's_(\(.*\))_(../../../../../website/tables/awspricing/\1)_' ../../../website/tables/awspricing/README.md > ./docs/tables/README.md
	sed -i.bak -e 's_(\(.*\).md)_(tables/\1)_' ../../../website/tables/awspricing/README.md
	mv ../../../website/tables/awspricing/README.md ../../../website/pages/docs/plugins/sources/awspricing/tables.md
	sed -i.bak -e 's_(\(.*\).md)_(\1)_' ../../../website/tables/awspricing/*.md
	rm -rf ../../../website/tables/awspricing/*.bak

# All gen targets
.PHONY: gen
gen: gen-docs