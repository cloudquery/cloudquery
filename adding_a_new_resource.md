# Adding a new resource

Some information can be found in the [docs for developing a new provider](https://docs.cloudquery.io/developers/developing-new-provider).

As a prerequisite, in [aws-sdk-go-v2](https://pkg.go.dev/github.com/aws/aws-sdk-go-v2) ensure API calls exist to list/describe the desired resource, and make note of:

   - to which aws service the resource belongs
   - the schema of the returned object(s)

## Setting up the service

If the service to which the resource belongs has not been used before in cq-provider-aws, there are a few steps that need to be done to configure it.

1. Create the service interface in [client/services.go](./client/services.go)
   * Don't forget to add the new service interface name to the go:generate comment.
1. Add the service to the `Services` struct in the [client/client.go](./client/client.go)
1. Init the service in the `initServices` function in [client/client.go](./client/client.go)
1. Run `go generate client/services.go` to create a mock for your new service. This will update [client/mocks/services.go](./client/mocks/services.go) automatically

## Setting up the resource

### Skeleton

1. In [client/services.go](./client/services.go), update the service interface and add the method(s) that you will be using to fetch the data from the aws sdk.
1. Run `go generate client/services.go` to create a mock for your new methods. This will update [client/mocks/services.go](./client/mocks/services.go) automatically
1. Create a file under `resources/` that follows the pattern of `<service>_<resource>`.
1. In that file, create a function that returns a `*schema.Table`
1. In [resources/provider.go](./resources/provider.go), add a mapping between the function you just created and the name of the resource that will be used in the config yml file.
1. Add a test in [clients/mocks/resources_test.go](./client/mocks/resources_test.go) and the corresponding test implementation in [clients/mocks/builders_test.go](./client/mocks/builders_test.go) for the resource following the existing examples.

### Implementation

Now that the skeleton has been set up, you can start to actually implement the resource. This consists of two parts: 

1. Defining the schema
1. Implementing resolver functions

#### Defining the schema

It is recommended that you look at a few existing resources as examples and also read through the comments on the source code for the [Table](https://github.com/cloudquery/cq-provider-sdk/blob/main/provider/schema/table.go) and [Column](https://github.com/cloudquery/cq-provider-sdk/blob/main/provider/schema/column.go) implementations for details.

For noncomplex fields, the SDK can directly resolve them into `Column`s for you, so all you need to do is specify the `Name` and the `Type`.

For complex fields or fields that require further API calls, you can defined your own `Resolver` for the `Column`.

#### Implementing Resolver Functions

A few important things to note when adding functions that call the AWS API:

- If possible, always use an API call that allows you to fetch many resources at once
- Take pagination into account. Ensure you fetch **all** of the resources
