PACKAGE_NAME  := github.com/cloudquery/cloudquery

.PHONY: build
build:
	@docker run \
		-e GOOS=linux \
		-e GOARCH=amd64 \
		lambci/lambda:build-go1.x \
		go build -v -o bin/cloudquery

.PHONY: plan
plan:
	@cd deploy/aws/terraform && terraform init && terraform plan

.PHONY: apply
apply:
	@cd deploy/aws/terraform && terraform init && terraform apply