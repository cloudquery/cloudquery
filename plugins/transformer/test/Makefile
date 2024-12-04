.PHONY: test
test:
	go test -race -timeout 3m ./...

.PHONY: lint
lint:
	golangci-lint run --config ../../.golangci.yml

.PHONY: gen-licenses
gen-licenses:
ifndef CI
	go install github.com/google/go-licenses@v1.6.0
	go run github.com/cloudquery/licenser@v0.2.0 report .
endif

# All gen targets
.PHONY: gen
gen: gen-licenses