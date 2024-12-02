PACKAGE_NAME          := github.com/cloudquery/cloudquery/plugins/destination/sqlite/v2

.PHONY: test
test:
# we clean the cache to avoid scenarios when we change something in the db and we want to retest without noticing nothing run
	go clean -testcache
	go test -race -timeout 3m ./...

.PHONY: lint
lint:
	golangci-lint run --config ../../.golangci.yml

# ghcr.io/cloudquery/golang-cross:latest \

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
		--clean --skip-validate --snapshot --skip-publish

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
