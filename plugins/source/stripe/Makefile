.PHONY: build
build:
	go build

# Test unit
.PHONY: test
test: gen-fixtures
	stripeSDKTag=$$(go list -m github.com/stripe/stripe-go/v74 | awk '{print $$2}') ;\
	stripeOpenAPITag=$$(curl --silent https://raw.githubusercontent.com/stripe/stripe-go/$$stripeSDKTag/OPENAPI_VERSION) ;\
	curl --silent https://raw.githubusercontent.com/stripe/openapi/$$stripeOpenAPITag/openapi/spec3.json > ./resources/testdata/spec3.json ;\
	go test -race -timeout 3m ./... ;\

.PHONY: gen-docs
gen-docs: build
	@command -v cloudquery >/dev/null 2>&1 || { \
		echo "Error: 'cloudquery' command not found. Please install it before running gen-docs."; \
		echo "You can install it by following the instructions at: https://www.cloudquery.io/docs/quickstart"; \
		exit 1; \
	}
	rm -rf ../../../website/tables/stripe
	cloudquery tables --format markdown --output-dir ../../../website/tables test/config.yml
	sed 's_(\(.*\))_(../../../../../website/tables/stripe/\1)_' ../../../website/tables/stripe/README.md > ./docs/tables/README.md
	sed -i.bak -e 's_(\(.*\).md)_(tables/\1)_' ../../../website/tables/stripe/README.md
	mv ../../../website/tables/stripe/README.md ../../../website/pages/docs/plugins/sources/stripe/tables.md
	sed -i.bak -e 's_(\(.*\).md)_(\1)_' ../../../website/tables/stripe/*.md
	rm -rf ../../../website/tables/stripe/*.bak

.PHONY: lint
lint:
	golangci-lint run --config ../../.golangci.yml 

.PHONY: gen-fixtures
gen-fixtures:
	rm -vf ./resources/testdata/fixtures_gen.json
	go run codegen/fixtures/main.go

.PHONY: gen-code
gen-code:
	go run codegen/main.go

# All gen targets
.PHONY: gen
gen: gen-fixtures gen-docs
