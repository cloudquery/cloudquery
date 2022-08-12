PACKAGE_NAME  := github.com/cloudquery/cloudquery

.PHONY: build
build:
	@docker run \
		--rm -v "${PWD}":/var/task \
		-w /var/task \
		-e GOOS=linux \
		-e GOARCH=amd64 \
		-e CGO_ENABLED=1 \
		lambci/lambda:build-go1.x \
		go build -v -o bin/cloudquery

.PHONY: lint
lint:
	@golangci-lint run --timeout 10m --verbose

.PHONY: test-sanity
test-sanity:
	./scripts/test-sanity.sh

.PHONY: test-unit
test-unit:
	go test ./...
