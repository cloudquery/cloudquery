# Test unit
.PHONY: test
test:
	go test -race -timeout 3m ./...

# Install tools
.PHONY: install-tools
install-tools:
	@echo Installing tools from tools/tool.go
	@cat tools/tool.go | grep _ | awk -F'"' '{print $$2}' | xargs -tI % go install %

.PHONY: lint
lint:
	golangci-lint run --config ../../.golangci.yml

.PHONY: lint-fix
lint-fix:
	golangci-lint run --fix --config ../../.golangci.yml

.PHONY: gen-mocks
gen-mocks: install-tools
	rm -rf ./client/mocks/*
	go generate ./client/...

.PHONY: gen-docs
gen-docs:
	rm -rf ../../../website/tables/hackernews
	go run main.go doc ../../../website/tables/hackernews
	sed -i.bak -e 's_(\(.*\).md)_(tables/\1)_' ../../../website/tables/hackernews/README.md
	mv ../../../website/tables/hackernews/README.md ../../../website/pages/docs/plugins/sources/hackernews/tables.md
	sed -i.bak -e 's_(\(.*\).md)_(\1)_' ../../../website/tables/hackernews/*.md
	rm -rf ../../../website/tables/hackernews/*.bak

# All gen targets
.PHONY: gen
gen: gen-mocks gen-docs
