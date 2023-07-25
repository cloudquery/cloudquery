# Generate mocks for mock/unit testing 
.PHONY: generate-mocks
generate-mocks:
	go install github.com/golang/mock/mockgen@v1.6.0
	go generate ./client/services.go

# Test unit
.PHONY: test
test:
	go test -race -timeout 3m ./...

.PHONY: lint
lint:
	golangci-lint run --config ../../.golangci.yml

.PHONY: gen-docs
gen-docs:
	rm -rf ../../../website/tables/digitalocean
	go run main.go doc ../../../website/tables/digitalocean
	sed 's_(\(.*\))_(../../../../../website/tables/digitalocean/\1)_' ../../../website/tables/digitalocean/README.md > ./docs/tables/README.md
	sed -i.bak -e 's_(\(.*\).md)_(tables/\1)_' ../../../website/tables/digitalocean/README.md
	mv ../../../website/tables/digitalocean/README.md ../../../website/pages/docs/plugins/sources/digitalocean/tables.md
	sed -i.bak -e 's_(\(.*\).md)_(\1)_' ../../../website/tables/digitalocean/*.md
	rm -rf ../../../website/tables/digitalocean/*.bak

# All gen targets
.PHONY: gen
gen: generate-mocks gen-docs