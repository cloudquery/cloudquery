# Development Environment Setup

## Requirements
 * [Go](https://go.dev/doc/install) 1.17+ (to build the provider)
 * [Terraform](https://www.terraform.io/downloads) (to run integration tests)

## Quick Start

### Building

Clone the repository:

```bash
git clone https://github.com/cloudquery/cq-provider-aws
```

Build the provider:

```
make build
```

### Running the provider in debug mode

1. Download [CloudQuery](https://github.com/cloudquery/cloudquery) latest version.
1. Execute `make run` and note of the `CQ_REATTACH_PROVIDERS` value.
1. Open another terminal and run `CQ_REATTACH_PROVIDERS=[VALUE_FROM_PREV] ./cloudquery fetch` 

> Make sure the authentication variables are exported in the provider process and not in cloudquery process.

See [docs](https://docs.cloudquery.io/docs/developers/debugging) for more details.

### Testing

The provider has two types of tests:

1. *Unit Tests* - run locally without any credentials and use mocking to return data from AWS APIs.
1. *Integration Tests* - run against real AWS APIs and uses test environment defined with `terraform` under `terraform/service_name/`

#### Unit Tests

Unit Tests don't require any credentials or internet access

```bash
make test-unit # This runs go test ./...
```
