.PHONY: test
test:
	go test -timeout 3m ./...

# .PHONY: build
# build:
# 	go build -o notion

.PHONY: lint
lint:
	@golangci-lint run --timeout 10m

.PHONY: gen-docs
# gen-docs:
# 	rm -rf ./docs/tables/*
# 	go run main.go doc ./docs/tables

gen-docs:
	@command -v cloudquery >/dev/null 2>&1 || { \
		echo "Error: 'cloudquery' command not found. Please install it before running gen-docs."; \
		echo "You can install it by following the instructions at: https://www.cloudquery.io/docs/quickstart"; \
		exit 1; \
	}
	rm -rf ../../../website/tables/notion
	cloudquery tables --format markdown --output-dir ../../../website/tables test/config.yml
# All gen targets
.PHONY: gen
gen: gen-docs
