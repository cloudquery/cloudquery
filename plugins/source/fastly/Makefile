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
	rm -rf ./client/services/*
	go run codegen/main.go
	rm -rf ./client/mocks/*
	go generate ./client/...

.PHONY: gen-docs
gen-docs:
	rm -rf ../../../website/tables/fastly
	go run main.go doc ../../../website/tables/fastly
	cp ../../../website/tables/fastly/README.md ./docs/tables/README.md
	sed -i.bak -e 's_(\(.*\).md)_(tables/\1)_' ../../../website/tables/fastly/README.md
	mv ../../../website/tables/fastly/README.md ../../../website/pages/docs/plugins/sources/fastly/tables.md
	sed -i.bak -e 's_(\(.*\).md)_(\1)_' ../../../website/tables/fastly/*.md
	rm -rf ../../../website/tables/fastly/*.bak

# All gen targets
.PHONY: gen
gen: gen-mocks gen-docs
