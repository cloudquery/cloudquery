# Build all
.PHONY: build
build:
	./scripts/build.sh

.PHONY: build
build-changed:
	./scripts/build.sh -c

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
	@cat tools/tool.go | grep _ | awk -F'"' '{print $$2}' | xargs -tI % go install %

# Install pre-commit hooks. This requires pre-commit to be installed (https://pre-commit.com/)
.PHONY: install-hooks
install-hooks:
	pre-commit install
