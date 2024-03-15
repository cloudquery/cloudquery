.PHONY: lint
lint:
	golangci-lint run --timeout 10m

.PHONY: test
test:
	go test -v -race ./...


.PHONY: gen-docs
gen-docs:
	rm -rf ../website/pages/docs/reference/cli/*.md
	go run main.go doc ../website/pages/docs/reference/cli/

.PHONY: gen-spec-schema
gen-spec-schema:
	go run internal/specs/v0/gen/main.go

.PHONY: gen
gen: gen-docs gen-spec-schema