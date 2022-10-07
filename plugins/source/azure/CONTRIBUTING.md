## CloudQuery Azure Source Plugin Contributing Guidelines

## Adding a new resource

Some information can be found in the [docs for developing a new provider](https://docs.cloudquery.io/developers/developing-new-provider).

As a prerequisite, ensure that API calls to list/describe the desired resource are present in
[github.com/Azure/go-autorest/autorest](https://pkg.go.dev/github.com/Azure/go-autorest/autorest), and make note of:

- to which Azure service the resource belongs
- the schema of the returned object(s)

## Setting up the service

If the service to which the resource belongs has not been used before in cq-provider-azure, there are a few steps that need to be done to configure it.

1. Create the service interface in [client/services](../../client/services)
2. Add the new service interface name to the go:generate comment in [client/services/services.go](../../client/services/services.go).
3. Add the service to the `Services` struct in [client/services/services.go](../../client/services/services.go)
4. Init the service in the `InitServices` function in [client/services/services.go](../../client/services/services.go)
5. Run `make generate-mocks` to create a mock for your new service. This will update [client/services/mocks](../../client/services/mocks) automatically.

## Setting up the resource

### Skeleton

1. Create a file under `resources/` that follows the pattern of `resources/<service>/<resource_name>`.
1. In that file, create a function that returns a `*schema.Table`
1. In [resources/provider.go](./resources/provider.go), add a mapping between the function you just created and the name of the resource that will be used in the config YAML file.
1. Add a test file at `resources/<service>/<resource>_test.go`. Follow other examples to create a test for the resource.
1. Run `go run docs/docs.go` to generate the documentation for the new resource

### Implementation

Now that the skeleton has been set up, you can start to actually implement the resource. This consists of two parts:

1. Defining the schema
1. Implementing resolver functions

#### Defining the schema

It is recommended that you look at a few existing resources as examples and also read through the comments on the source code for the [Table](https://github.com/cloudquery/cq-provider-sdk/blob/main/provider/schema/table.go) and [Column](https://github.com/cloudquery/cq-provider-sdk/blob/main/provider/schema/column.go) implementations for details.

For most fields, the SDK can directly resolve them into a `Column`s for you, so all you need to do is specify the `Name` and the `Type`.

For complex fields or fields that require further API calls, you can defined your own `Resolver` for the `Column`.

#### Implementing Resolver Functions

A few important things to note when adding functions that call the Azure API:

- If possible, always use an API call that allows you to fetch many resources at once
- Take pagination into account. Ensure you fetch **all** of the resources
