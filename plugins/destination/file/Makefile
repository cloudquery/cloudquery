.PHONY: test
test:
	go test -timeout 3m ./...

.PHONY: lint
lint:
	@golangci-lint run --config ../../.golangci.yml --timeout 10m --verbose
