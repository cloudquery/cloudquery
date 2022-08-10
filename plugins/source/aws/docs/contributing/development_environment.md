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
2. *Integration Tests* - run against real AWS APIs and uses test environment defined with `terraform` under [`terraform/service_name/`](../../terraform)

#### Unit Tests

Unit Tests don't require any credentials or internet access

```bash
make test-unit # This runs go test ./...
```

Unit tests include:
- Specific resources tests. You can find those next to each resource, in the [`resources/services`](../../resources/services) folder.
- DB migration tests. You can find the code for these tests [here](../../resources/provider/provider_test.go).
- Client tests. You can find those in the [`client`](../../client) folder.

#### Integration Tests

These are documented in the Adding a new resource guide. See [here](./adding_a_new_resource.md#integration-tests) for more information.

### Pre-commit Hooks

This repository provides a pre-commit hook to check your code before committing it. To use it:

 1. [Install pre-commit](https://pre-commit.com/) (on Mac OS: `brew install pre-commit`)
 2. [Install golangci-lint](https://golangci-lint.run/usage/install/#local-installation) (on Mac OS: `brew install golangci-lint`)
 3. In the root of this repo, run:
    ```shell
    make install-hooks
    ```

Now your code will be statically checked for errors prior to commit. 