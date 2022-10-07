.PHONY: lint
lint:
	@golangci-lint run --timeout 10m --verbose

.PHONY: test
test:
	go test ./...


.PHONY: doc
doc:
	rm -rf ../website/pages/docs/reference/cli/*.md
	go run main.go doc ../website/pages/docs/reference/cli/

