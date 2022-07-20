# Adding a new resource

Some information can be found in the [docs for developing a new provider](https://docs.cloudquery.io/developers/developing-new-provider).

As a prerequisite, ensure that API calls to list/describe the desired resource are present in
[github.com/Azure/go-autorest/autorest](https://pkg.go.dev/github.com/Azure/go-autorest/autorest), and make note of:

   - to which Azure service the resource belongs
   - the schema of the returned object(s)

## Setting up the service

If the service to which the resource belongs has not been used before in cq-provider-azure, there are a few steps that need to be done to configure it.

1. Create the service interface in [client/services/services.go](./client/services/services.go)
   * Don't forget to add the new service interface name to the go:generate comment.
1. Add the service to the `Services` struct in the [client/services/services.go](./client/services/services.go)
1. Init the service in the `InitServices` function in [client/services/services.go](./client/services.go)
1. Run `make generate-mocks` to create a mock for your new service. This will update [client/mocks/services.go](./client/mocks/services.go) automatically.

## Setting up the resource

### Skeleton

1. Create a file under `resources/` that follows the pattern of `resources/<service>/<resource_name>`.
1. In that file, create a function that returns a `*schema.Table`
1. In [resources/provider.go](./resources/provider.go), add a mapping between the function you just created and the name of the resource that will be used in the config yml file.
1. Add a test file at `resources/<service>/<resource>_test.go`. Follow other examples to create a test for the resource.
1. Run `go run docs/docs.go` to generate the documentation for the new resource

### Implementation

Now that the skeleton has been set up, you can start to actually implement the resource. This consists of two parts: 

1. Defining the schema
1. Implementing resolver functions

#### Defining the schema

It is recommended that you look at a few existing resources as examples and also read through the comments on the source code for the [Table](https://github.com/cloudquery/cq-provider-sdk/blob/main/provider/schema/table.go) and [Column](https://github.com/cloudquery/cq-provider-sdk/blob/main/provider/schema/column.go) implementations for details.

For noncomplex fields, the SDK can directly resolve them into `Column`s for you, so all you need to do is specify the `Name` and the `Type`.

For complex fields or fields that require further API calls, you can defined your own `Resolver` for the `Column`.

#### Implementing Resolver Functions

A few important things to note when adding functions that call the Azure API:

- If possible, always use an API call that allows you to fetch many resources at once
- Take pagination into account. Ensure you fetch **all** of the resources

#### Integration Tests

To Run Integration tests please run

```bash
docker run -p 5432:5432 -e POSTGRES_PASSWORD=pass -d  postgres:13.3
cloud auth application-default login
# If you are team member and have access it is advised to set subscription to our test environment
az account set --subscription 78f26f10-0e60-4293-8a7e-122584ccb40d
# or set to your test subscription where you have write access to be able to apply terraform files.
go test -run=TestIntegration -tags=integration ./...
```

To Run integration test for specific table:

```bash
go test -run="TestIntegration/ROOT_TABLE_NAME" -tags=integration ./...
# For example
go test -run="TestIntegration/azure_sql_managed_instances" -tags=integration ./...
```

## Adding new Terraform File Guidelines

There a few good rule of thumb to follow when creating new terraform resources that will be served as testing infrastructure:
* For every resource create it's own resource_group.
* Use `location = "East US"`.
* If possible make all resources private.
* Make sure to replace built-in plain text passwords with `random_password` generator
* For every compute/db try to use the smallest size to keep the cost low
* If autoscaling option is present, always turn it off

If you want to apply the terraform locally first before pushing it to CI and applying there use:

```
cd terraform/YOUR_SERVICE_NAME/local
# Use AB as your initial so you can have multiple team members working on the same account without conflicting resources
terraform apply -var="prefix=AB"
go test -run="TestIntegration/ROOT_TABLE_NAME" -tags=integration ./...
# Don't forget to run destroy if you are in a playground account
```