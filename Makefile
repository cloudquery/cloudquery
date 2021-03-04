PACKAGE_NAME          := github.com/cloudquery/cloudquery

.PHONY: build
build: config
	@docker run \
		--rm -v "${PWD}":/var/task \
		-w /var/task \
		-e GOOS=linux \
		-e GOARCH=amd64 \
		-e CGO_ENABLED=1 \
		lambci/lambda:build-go1.x \
		go build -v -o bin/cloudquery

.PHONY: proto
proto:
	@protoc --go_out=. --go-grpc_out=. ./sdk/proto/plugin.proto

.PHONY: config
config:
	@go run main.go gen config aws --force --path ./bin/config.yml

.PHONY: plan
plan:
	@cd deploy/aws/terraform && terraform init && terraform plan

.PHONY: apply
apply:
	@cd deploy/aws/terraform && terraform init && terraform apply