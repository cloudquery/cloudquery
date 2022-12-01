PACKAGE_NAME          := github.com/cloudquery/cloudquery/plugins/destination/snowflake

.PHONY: test
test:
# we clean the cache to avoid scenarios when we change something in the db and we want to retest without noticing nothing run
	go clean -testcache
	go test -timeout 3m ./...

.PHONY: lint
lint:
	@golangci-lint run --config ../../.golangci.yml --timeout 10m --verbose

.PHONY: release-dry-run
release-dry-run:
	@docker run \
		--privileged \
		-e CGO_ENABLED=1 \
		-e GORELEASER_KEY="$(GORELEASER_KEY)" \
		-v /var/run/docker.sock:/var/run/docker.sock \
		-v `pwd`:/go/src/$(PACKAGE_NAME) \
		-w /go/src/$(PACKAGE_NAME) \
		ghcr.io/cloudquery/golang-cross:v10.0.0 \
		--rm-dist --skip-validate --snapshot --skip-publish