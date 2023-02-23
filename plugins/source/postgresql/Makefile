# Unit tests
.PHONY: test
test:
	go test -race -timeout 3m ./...

.PHONY: lint
lint:
	golangci-lint run --config ../../.golangci.yml

.PHONY: gen-docs
gen-docs:
	echo "skipping docs generation for postgresql source plugin"

# All gen targets
.PHONY: gen
gen: gen-docs
