.PHONY: lint
lint:
	golangci-lint run --timeout 10m

.PHONY: test
test:
	go test -v -race -timeout 15m ./...


.PHONY: gen-docs
gen-docs:
	rm -rf docs/reference/*.md
	mkdir -p docs/reference
	go run main.go doc docs/reference

.PHONY: gen-spec-schema
gen-spec-schema:
	go run internal/specs/v0/gen/main.go

.PHONY: gen-licenses
gen-licenses:
ifndef CI
	go install github.com/google/go-licenses@v1.6.0
	go run github.com/cloudquery/licenser@v0.2.0 report .
endif

# All gen targets
.PHONY: gen
gen: gen-licenses

.PHONY: gen
gen: gen-docs gen-spec-schema gen-licenses
