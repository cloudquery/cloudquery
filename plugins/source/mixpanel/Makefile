# Test unit
.PHONY: test
test:
	go test -race -timeout 3m ./...

.PHONY: gen-docs
gen-docs:
	rm -rf ../../../website/tables/mixpanel
	go run main.go doc ../../../website/tables/mixpanel
	sed -i.bak -e 's_(\(.*\).md)_(tables/\1)_' ../../../website/tables/mixpanel/README.md
	mv ../../../website/tables/mixpanel/README.md ../../../website/pages/docs/plugins/sources/mixpanel/tables.md
	sed -i.bak -e 's_(\(.*\).md)_(\1)_' ../../../website/tables/mixpanel/*.md
	rm -rf ../../../website/tables/mixpanel/*.bak

.PHONY: lint
lint:
	golangci-lint run --config ../../.golangci.yml 

# All gen targets
.PHONY: gen
gen: gen-docs
