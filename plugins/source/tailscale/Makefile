# Unit tests
.PHONY: test
test:
	go test -race -timeout 3m ./...

.PHONY: lint
lint:
	golangci-lint run --config ../../.golangci.yml

.PHONY: gen-docs
gen-docs:
	rm -rf ../../../website/tables/tailscale
	go run main.go doc ../../../website/tables/tailscale
	sed 's_(\(.*\))_(../../../../../website/tables/tailscale/\1)_' ../../../website/tables/tailscale/README.md > ./docs/tables/README.md
	sed -i.bak -e 's_(\(.*\).md)_(tables/\1)_' ../../../website/tables/tailscale/README.md
	mv ../../../website/tables/tailscale/README.md ../../../website/pages/docs/plugins/sources/tailscale/tables.md
	sed -i.bak -e 's_(\(.*\).md)_(\1)_' ../../../website/tables/tailscale/*.md
	rm -rf ../../../website/tables/tailscale/*.bak

# All gen targets
.PHONY: gen
gen: gen-docs
