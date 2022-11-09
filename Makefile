# Build all
.PHONY: build
build:
	./scripts/build.sh

.PHONY: build-changed
build-changed:
	./scripts/build.sh -c

.PHONY: build-cli
build-cli:
	(cd cli && go build -o ../bin/cloudquery .)

# Test unit
.PHONY: test-unit
test-unit:
	./scripts/test-unit.sh

# Test unit (only changed files)
test-unit-changed:
	./scripts/test-unit.sh -c

# Install tools
.PHONY: install-tools
install-tools:
	@echo Installing tools from tools/tool.go
	@cat tools/tool.go | grep _ | awk -F'"' '{print $$2}' | xargs sh -c 'for arg do echo "Installing $$arg"; go get "$$arg"; go install "$$arg"; done' _

# Install pre-commit hooks. This requires pre-commit to be installed (https://pre-commit.com/)
.PHONY: install-hooks
install-hooks:
	pre-commit install

# Regenerate docs for all plugins
.PHONY: update-docs
update-docs:
	./scripts/update-docs.sh

# Update plugin-sdk for all plugins and cli
.PHONY: update-plugin-sdk
update-plugin-sdk:
	./scripts/update-plugin-sdk.sh


.PHONY: list
list:
	@grep '^[^#[:space:].].*$$' Makefile | sed 's/\(.*\):/"\1",/' | tr -d '\n' | sed 's/,$$//g'
