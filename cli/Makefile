PACKAGE_NAME  := github.com/cloudquery/cloudquery

.PHONY: lint
lint:
	@golangci-lint run --timeout 10m --verbose

.PHONY: test-unit
test-unit:
	go test ./...