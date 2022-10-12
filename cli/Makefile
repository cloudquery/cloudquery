.PHONY: lint
lint:
	@golangci-lint run --timeout 10m --verbose

.PHONY: test
test:
	go test ./...


.PHONY: gen-docs
gen-docs:
	rm -rf ../website/pages/docs/reference/cli/*.md
	go run main.go doc ../website/pages/docs/reference/cli/

