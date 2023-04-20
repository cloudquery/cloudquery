.PHONY: test
test:
	go test -tags=assert -race -timeout 3m ./...

.PHONY: lint
lint:
	golangci-lint run --config ../../.golangci.yml
