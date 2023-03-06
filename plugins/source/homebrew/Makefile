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

.PHONY: gen-docs
gen-docs:
	rm -rf ../../../website/tables/homebrew
	go run main.go doc ../../../website/tables/homebrew
	cp ../../../website/tables/homebrew/README.md ./docs/tables/README.md
	sed -i.bak -e 's_(\(.*\).md)_(tables/\1)_' ../../../website/tables/homebrew/README.md
	mv ../../../website/tables/homebrew/README.md ../../../website/pages/docs/plugins/sources/homebrew/tables.md
	sed -i.bak -e 's_(\(.*\).md)_(\1)_' ../../../website/tables/homebrew/*.md
	rm -rf ../../../website/tables/homebrew/*.bak

# All gen targets
.PHONY: gen
gen: gen-docs
