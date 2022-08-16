# Development Environment Setup

## Requirements
 * [Go](https://go.dev/doc/install) 1.18+ (to build the plugins)
 * [Terraform](https://www.terraform.io/downloads) (to run integration tests) (optional)

## Quick Start

### Building

Clone the repository:

```bash
git clone https://github.com/cloudquery/cloudquery
```

Build the CLI and all plugins:

```
make build
```

### Running plugins in debug mode

1. Execute `make run` from the chosen plugin directory under [../plugins/source](../plugins/source) (e.g.  [../plugins/source/aws](../plugins/source/aws)).
2. Note the `CQ_REATTACH_PROVIDERS` value.
3. Open another terminal and run `CQ_REATTACH_PROVIDERS=[VALUE_FROM_PREV] ./bin/cloudquery fetch` 

> **Important**: Make sure the authentication variables are exported in the provider process and not in CloudQuery process.

See [docs](https://docs.cloudquery.io/docs/developers/debugging) for more details.

### Testing

The provider has two types of tests:

1. *Unit Tests* - run locally without any credentials and use mocking to return data from AWS APIs.
2. *Integration Tests* - run against real APIs and uses test environment defined with `terraform` under `terraform/service_name/`

#### Unit Tests

Unit Tests don't require any credentials or internet access

```bash
make test-unit # This runs go test ./...
```

Unit tests for plugins include:
- Specific resource tests. You can find those next to each resource, in the `resources/services` folder under the plugin directory.
- DB migration tests. You can find the code for these tests in `provider_test.go`.
- Client tests. You can find those in the `client` folder.

#### Integration Tests

These are documented in the Adding a new resource guide. See:
 - [Adding a new Resource (AWS)](../plugins/source/aws/docs/contributing/adding_a_new_resource.md#integration-tests)
 - [Adding a new Resource (Azure)](../plugins/source/azure/docs/contributing/adding_a_new_resource.md#integration-tests)
 - [Adding a new Resource (GCP)](../plugins/source/gcp/docs/contributing/adding_a_new_resource.md#integration-tests)
