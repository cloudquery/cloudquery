.PHONY: build
build:
	go build

.PHONY: test-unit
test-unit:
	go test ./...

.PHONY: test-integration
test-integration:
	go test -run=TestIntegration -tags=integration ./...

.PHONY: generate-mocks
generate-mocks:
	go generate ./client/services/...

# This will only run terraform apply for directories that were changes.
# Mainly used in CI
.PHONY: terraform-apply-change-only
terraform-apply-change-only:
	for d in terraform/* ; \
	do \
		echo $d ; \
	done;
