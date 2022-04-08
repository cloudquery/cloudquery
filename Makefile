## install the latest version of CQ
.PHONY: install-cq
install-cq:
	@if [[ "$(OS)" != "Darwin" && "$(OS)" != "Linux" && "$(OS)" != "Windows" ]]; then echo "\n Invalid OS set. Valid Options are Darwin, Linux and Windows. Example invocation is:\n make OS=Linux ARCH=arm64 install-cq \n For more information go to  https://docs.cloudquery.io/docs/getting-started \n"; exit 1; fi
	@if [[ "$(ARCH)" != "x86_64" && "$(ARCH)" != "arm64" ]]; then echo "\n Invalid ARCH set. Valid options are x86_64 and arm64. Example invocation is:\n make OS=Linux ARCH=arm64 install-cq \n For more information go to  https://docs.cloudquery.io/docs/getting-started \n"; exit 1; fi
	curl -L https://github.com/cloudquery/cloudquery/releases/latest/download/cloudquery_${OS}_${ARCH} -o cloudquery
	chmod a+x cloudquery

# start a timescale db running in a local container
.PHONY: ts-start
ts-start:
	docker run -p 5433:5432 -e POSTGRES_PASSWORD=pass -d timescale/timescaledb:latest-pg14

# stop the timescale db running in a local container
.PHONY: ts-stop
ts-stop:
	docker stop $(docker ps -q --filter ancestor=timescale/timescaledb:latest-pg14)

# start a running docker container
.PHONY: start-pg
start-pg:
	docker run -p 5432:5432 -e POSTGRES_PASSWORD=pass -d postgres

# stop a running docker container
.PHONY: stop-pg
stop-pg:
	docker stop $$(docker ps -q --filter ancestor=postgres:latest)

# connect to pg via cli
.PHONY: pg-connect
pg-connect:
	psql -h localhost -p 5432 -U postgres -d postgres

# build the cq aws provider
.PHONY: build
build:
	go build -o cq-provider

# build and run the cq provider
.PHONY: run
run: build
	CQ_PROVIDER_DEBUG=1 CQ_REATTACH_PROVIDERS=.cq_reattach ./cq-provider

# Run a fetch command
.PHONY: fetch
fetch:
	CQ_PROVIDER_DEBUG=1 CQ_REATTACH_PROVIDERS=.cq_reattach ./cloudquery fetch --dsn "postgres://postgres:pass@localhost:5432/postgres?sslmode=disable" -v

# Generate mocks for mock/unit testing 
.PHONY: generate-mocks
generate-mocks:
	go generate ./client/services/...

# Test unit
.PHONY: test-unit
test-unit:
	go test -timeout 3m ./...

# Run an integration tests
.PHONY: test-integration
test-integration:
	@if [[ "$(tableName)" == "" ]]; then go test -run=TestIntegration -timeout 3m -tags=integration ./...; else go test -run="TestIntegration/$(tableName)" -timeout 3m -tags=integration ./...; fi
