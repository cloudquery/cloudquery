.PHONY: run-test
run-test:
# we clean the cache to avoid scenarios when we change something in the db and we want to retest without noticing nothing run
	go clean -testcache
	go test -race -timeout 3m ./...

.PHONY: test
test: setup-gremlin run-test teardown-gremlin

.PHONY: setup-gremlin
setup-gremlin:
	docker compose up --wait
	sleep 10

.PHONY: teardown-gremlin
teardown-gremlin:
	docker compose down --volumes --remove-orphans

.PHONY: gen-coverage-report
gen-coverage-report:
	go test -timeout 3m -coverprofile=coverage.out.tmp ./...
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

.PHONY: coverage
coverage: setup-gremlin gen-coverage-report teardown-gremlin

.PHONY: lint
lint:
	golangci-lint run --config ../../.golangci.yml

.PHONY: gen-spec-schema
gen-spec-schema:
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
