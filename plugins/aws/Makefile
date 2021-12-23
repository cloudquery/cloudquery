export CQ_PROVIDER_DEBUG=1
export CQ_REATTACH_PROVIDERS=.cq_reattach

# install the latest version of CQ
install-cq:
	@if [[ "$(OS)" != "Darwin" && "$(OS)" != "Linux" && "$(OS)" != "Windows" ]]; then echo "\n Invalid OS set. Valid Options are Darwin, Linux and Windows. Example invocation is:\n make OS=Linux ARCH=arm64 install-cq \n For more information go to  https://docs.cloudquery.io/docs/getting-started \n"; exit 1; fi
	@if [[ "$(ARCH)" != "x86_64" && "$(ARCH)" != "arm64" ]]; then echo "\n Invalid ARCH set. Valid options are x86_64 and arm64. Example invocation is:\n make OS=Linux ARCH=arm64 install-cq \n For more information go to  https://docs.cloudquery.io/docs/getting-started \n"; exit 1; fi
	curl -L https://github.com/cloudquery/cloudquery/releases/latest/download/cloudquery_${OS}_${ARCH} -o cloudquery
	chmod a+x cloudquery


# build the cq aws provider
build:
	go build -o cq-provider-aws -ldflags="-X 'github.com/cloudquery/cq-provider-aws/resources.Version=LocalDevelopmentMakefile'"

# build and run the cq aws provider
run: build
	./cq-provider-aws

# start a running docker container
start-pg:
	docker run -p 5432:5432 -e POSTGRES_PASSWORD=pass -d postgres

# stop a running docker container
stop-pg:
	docker stop $$(docker ps -q --filter ancestor=postgres:latest)

# connect to pg via cli
pg-connect:
	psql -h localhost -p 5432 -U postgres -d postgres



# Run an integration test
# you can pass in a specific test to run by specifying the testName:
# make testName=TestIntegrationElasticbeanstal$ e2e-test
e2e-test:
	@if [[ "$(testName)" == "" ]]; then echo "\n testName must be set \n"; exit 1; fi
	INTEGRATION_TESTS=1 TF_VAR_PREFIX=cq-testing TF_APPLY_RESOURCES=0 TF_VAR_SUFFIX=integration go test -timeout 180m -v -run ^$(testName)  github.com/cloudquery/cq-provider-aws/resources/integration_tests


e2e-test-read-only:
	@if [[ "$(testName)" == "" ]]; then echo "\n testName must be set \n"; exit 1; fi
	@if [[ "$(TF_VAR_PREFIX)" == "" ]]; then echo "\n TF_VAR_PREFIX must be set \n"; exit 1; fi
	@if [[ "$(TF_VAR_SUFFIX)" == "" ]]; then echo "\n TF_VAR_SUFFIX must be set \n"; exit 1; fi
	INTEGRATION_TESTS=1TF_APPLY_RESOURCES=0 go test -timeout 180m -v -run ^$(testName)  github.com/cloudquery/cq-provider-aws/resources/integration_tests

e2e-test-with-apply:
	@if [[ "$(testName)" == "" ]]; then echo "\n testName must be set \n"; exit 1; fi
	INTEGRATION_TESTS=1 TF_VAR_PREFIX=cq-testing TF_APPLY_RESOURCES=1 TF_VAR_SUFFIX=integration go test -timeout 180m -v -run ^$(testName)  github.com/cloudquery/cq-provider-aws/resources/integration_tests


# Generate mocks for mock/unit testing 
create-mocks:
	go install github.com/golang/mock/mockgen
	$(shell PATH=$$PATH:$$(go env GOPATH)/bin && go generate client/services.go)

# Run a fetch command
fetch:
	./cloudquery fetch --dsn "postgres://postgres:pass@localhost:5432/postgres?sslmode=disable" -v
