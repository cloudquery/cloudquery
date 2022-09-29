# Generate mocks for mock/unit testing 
.PHONY: generate-mocks
generate-mocks:
	go generate ./client/services/...

# Test unit
.PHONY: test-unit
test-unit:
	go test -timeout 3m ./...

# Install pre-commit hooks. This requires pre-commit to be installed (https://pre-commit.com/)
.PHONY: install-hooks
install-hooks:
	pre-commit install
