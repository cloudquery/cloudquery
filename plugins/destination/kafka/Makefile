.PHONY: test
test:
# we clean the cache to avoid scenarios when we change something in the db and we want to retest without noticing nothing run
	go clean -testcache
	go test -race -timeout 3m ./...

.PHONY: lint
lint:
	golangci-lint run --config ../../.golangci.yml

.PHONY: gen-spec-schema
gen-spec-schema:
	# required for loading comments from filetypes
	go mod vendor
	go run client/spec/gen/main.go

.PHONY: gen-licenses
gen-licenses:
ifndef CI
	go install github.com/google/go-licenses@v1.6.0
	go run github.com/cloudquery/licenser@v0.2.0 report .
endif

# All gen targets
.PHONY: gen
gen: gen-spec-schema gen-licenses
