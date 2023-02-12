.PHONY: test
test:
# we clean the cache to avoid scenarios when we change something in the db and we want to retest without noticing nothing run
	go clean -testcache
	go test -race -timeout 1m ./...

.PHONY: lint
lint:
	golangci-lint run --config ../../.golangci.yml
