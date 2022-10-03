# Test unit
.PHONY: test-unit
test-unit:
	go test -timeout 3m ./...

.PHONY: lint
lint:
	@golangci-lint run --timeout 10m --verbose
