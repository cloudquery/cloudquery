.PHONY: test
test:
	go test -race -timeout 3m ./...

.PHONY: lint
lint:
	golangci-lint run --config ../../.golangci.yml

.PHONY: gen-spec-schema
gen-spec-schema:
	# required for loading comments from filetypes
	go mod vendor
	go run client/spec/gen/main.go

# All gen targets
.PHONY: gen
gen: gen-spec-schema
