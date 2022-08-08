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

.PHONY: init
init:
	@cd deploy/aws/terraform && terraform init

.PHONY: plan
plan:
	@cd deploy/aws/terraform && terraform init && terraform plan

.PHONY: apply
apply:
ifeq (,$(wildcard ./bin/cloudquery))
	echo "Run \"make build\" before deploy."
else
	@cd deploy/aws/terraform && terraform init && terraform apply
endif

.PHONY: destroy
destroy:
	@cd deploy/aws/terraform && terraform destroy

.PHONY: lint
lint:
	@golangci-lint run --timeout 10m --verbose

.PHONY: test-sanity
test-sanity:
	./scripts/test-sanity.sh

.PHONY: test-unit
test-unit:
	go test ./...