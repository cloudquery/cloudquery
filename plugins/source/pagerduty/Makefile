# Test unit
.PHONY: test
test:
	go test -race -timeout 3m ./...

.PHONY: lint
lint:
	golangci-lint run --config ../../.golangci.yml

.PHONY: gen-mocks
gen-mocks:
	go install github.com/golang/mock/mockgen@v1.6.0
	rm -rf ./client/mocks/*
	go generate ./client/...

.PHONY: gen-docs
gen-docs:
	rm -rf ../../../website/tables/pagerduty
	go run main.go doc ../../../website/tables/pagerduty
	sed -i.bak -e 's_(\(.*\).md)_(tables/\1)_' ../../../website/tables/pagerduty/README.md
	mv ../../../website/tables/pagerduty/README.md ../../../website/pages/docs/plugins/sources/pagerduty/tables.md
	sed -i.bak -e 's_(\(.*\).md)_(\1)_' ../../../website/tables/pagerduty/*.md
	rm -rf ../../../website/tables/pagerduty/*.bak

# All gen targets
.PHONY: gen
gen: gen-mocks gen-docs
