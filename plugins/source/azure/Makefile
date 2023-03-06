# Test unit
.PHONY: test
test:
	go test -race -timeout 3m ./...

.PHONY: gen-docs
gen-docs:
	rm -rf ../../../website/tables/azure
	go run main.go doc ../../../website/tables/azure
	cp ../../../website/tables/azure/README.md ./docs/tables/README.md
	sed -i.bak -e 's_(\(.*\).md)_(tables/\1)_' ../../../website/tables/azure/README.md
	mv ../../../website/tables/azure/README.md ../../../website/pages/docs/plugins/sources/azure/tables.md
	sed -i.bak -e 's_(\(.*\).md)_(\1)_' ../../../website/tables/azure/*.md
	rm -rf ../../../website/tables/azure/*.bak

.PHONY: lint
lint:
	golangci-lint run --config ../../.golangci.yml 

.PHONY: gen
gen: gen-docs
	echo "Done"
