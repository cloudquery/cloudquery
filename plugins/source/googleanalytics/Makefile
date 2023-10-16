.PHONY: test
test:
	go test -race -timeout 3m ./...

.PHONY: lint
lint:
	golangci-lint run --config ../../.golangci.yml


.PHONY: gen-docs
gen-docs:
	echo "skipping docs generation for googleanalytics source plugin"

.PHONY: gen-spec-schema
gen-spec-schema:
	go run client/schemagen/main.go

# All gen targets
.PHONY: gen
gen: gen-spec-schema gen-docs
