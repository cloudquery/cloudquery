.PHONY: lint
lint:
	golangci-lint run --timeout 10m

.PHONY: test
test:
	go test -v -race -timeout 20m ./...

.PHONY: coverage
coverage:
	go test -v -timeout 20m -coverprofile=coverage.out.tmp ./... || true
	cat coverage.out.tmp | grep -vE "MockGen|codegen|mocks" > coverage.out
	rm coverage.out.tmp
	echo "| File | Function | Coverage |" > coverage.md
	echo "| --- | --- | --- |" >> coverage.md
	go tool cover -func=coverage.out | tail -n +2 | while read line; do \
		file=$$(echo $$line | awk '{print $$1}'); \
		func=$$(echo $$line | awk '{print $$2}'); \
		cov=$$(echo $$line | awk '{print $$3}'); \
		printf "| %s | %s | %s |\\n" "$$file" "$$func" "$$cov" >> coverage.md; \
	done
	rm coverage.out


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
