# Test unit
.PHONY: test
test:
	go test -race -timeout 3m ./...

.PHONY: lint
lint:
	golangci-lint run --config ../../.golangci.yml

.PHONY: gen-docs
gen-docs:
	rm -rf ../../../website/tables/azuredevops
	go run main.go doc ../../../website/tables/azuredevops
	sed -i.bak -e 's_(\(.*\).md)_(tables/\1)_' ../../../website/tables/azuredevops/README.md
	mv ../../../website/tables/azuredevops/README.md ../../../website/pages/docs/plugins/sources/azuredevops/tables.md
	sed -i.bak -e 's_(\(.*\).md)_(\1)_' ../../../website/tables/azuredevops/*.md
	rm -rf ../../../website/tables/azuredevops/*.bak

.PHONY: gen
gen: gen-docs