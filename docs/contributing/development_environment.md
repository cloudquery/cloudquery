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

The provider has two types of test:

1. *Unit Tests* - Those tests can run locally without any credentials and use mocking to return data from AWS APIs.
1. *Integration Tests* - Those tests run against real AWS APIs and uses test environment defined with `terraform` under `terraform/service_name/`

#### Unit Tests

Unit Tests don't require any credentials or internet access

```bash
make test-unit # This runs go test ./...
```

#### Integration Tests

Theses tests will test the provider against real AWS resources defined with terraform under `terraform/service_name`.

If you are a team-member and you have read-access to our test environment

```bash
// Authenticate with AWS
make test-integration
```

If you are developing a new resource and want to setup a test environment it is advised only to setup the needed subset of the test account.

```
// Authenticate with AWS account
cd terraform/service/local
# Change AZ to two character prefix (For example, John Doe: JD)
terraform apply -var="prefix=AZ"
cd ../../..

# TABLE_NAME will be the root table you are working on so it won't run all other tests (as you don't wont to spin-up the whole environment)
make tableName=TABLE_NAME test-integration

# Don't forget to destroy your test resources
cd terraform/service/local
terraform destroy
```
