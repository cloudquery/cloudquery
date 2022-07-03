# Adding a new resource

Some information can be found in the [docs for developing a new provider](https://docs.cloudquery.io/developers/developing-new-provider).

As a prerequisite, in [aws-sdk-go-v2](https://pkg.go.dev/github.com/aws/aws-sdk-go-v2) ensure API calls exist to list/describe the desired resource, and make note of:

   - to which aws service the resource belongs
   - the schema of the returned object(s)

## Setting up the service

If the service to which the resource belongs has not been used before in cq-provider-aws, there are a few steps that need to be done to configure it.

1. Create the service interface in [client/services.go](../../client/services.go)
   * Don't forget to add the new service interface name to the go:generate comment.
1. Add the service to the `Services` struct in the [client/client.go](../../client/client.go)
1. Init the service in the `initServices` function in [client/client.go](../../client/client.go)
1. Run `go generate client/services.go` to create a mock for your new service. This will update [client/mocks/mock_<service>.go](../../client/mocks) automatically

> If you get an error about not being able to find `mockgen`, run `make install-tools` to install it. If it still fails, run `export PATH=${PATH}:`go env GOPATH`/bin` in you shell to set up your `PATH` environment properly

> You might need to update an existing AWS client by running `go get github.com/aws/aws-sdk-go-v2/service/<service-name>@latest` and then `go mod tidy`

## Setting up the resource

### Skeleton

1. In [client/services.go](../../client/services.go), update the service interface and add the method(s) that you will be using to fetch the data from the aws sdk.
1. Run `go generate client/services.go` to create a mock for your new methods. This will update [client/mocks/mock_<service>.go](../../client/mocks) automatically.
1. Create a file under [resources/services/<service>](../../resources/services) that follows the pattern of `<resource>.go`.
1. In that file, create a function that returns a `*schema.Table`.
1. In [resources/provider.go](../../resources/provider/provider.go), add a mapping between the function you just created and the name of the resource that will be used in the config yml file.
1. Add a test file at [resources/services/<service>/<resource>_mock_test.go](../../resources/services). Follow other examples to create a test for the resource.
1. Run `go run docs/docs.go` to generate the documentation for the new resource.

### Implementation

Now that the skeleton has been set up, you can start to actually implement the resource. This consists of two parts: 

1. Defining the schema
1. Implementing resolver functions

#### Defining the schema

It is recommended that you look at a few existing resources as examples and also read through the comments on the source code for the [Table](https://github.com/cloudquery/cq-provider-sdk/blob/main/provider/schema/table.go) and [Column](https://github.com/cloudquery/cq-provider-sdk/blob/main/provider/schema/column.go) implementations for details.

For simple fields, the SDK can directly resolve them into `Column`s for you, so all you need to do is specify the `Name` and the `Type`.

For complex fields or fields that require further API calls, you can define your own `Resolver` for the `Column`.

#### Implementing Resolver Functions

A few important things to note when adding functions that call the AWS API:

- If possible, always use an API call that allows you to fetch many resources at once
- Take pagination into account. Ensure you fetch **all** of the resources

### Integration Tests

To prepare your environment for running integration tests:
```bash
# Start Postgres in a Docker container
docker run -p 5432:5432 -e POSTGRES_PASSWORD=pass -d  postgres:13.3

# Login with AWS. AWS_ACCESS_KEY_ID and AWS_SECRET_ACCESS_KEY can also be used here, if you wish.
# See all options at https://hub.cloudquery.io/providers/cloudquery/aws/latest
export AWS_PROFILE={Your AWS profile}
```

To run an integration test for a specific table:

```bash
go test -run="TestIntegration/ROOT_TABLE_NAME" -tags=integration ./...
# For example
go test -run="TestIntegration/aws_lambda_functions" -tags=integration ./...
```

> Note: You can override the Postgres database URL used for integration tests by specifying a DATABASE_URL environment variable, for example:
> 
> ```
> export DATABASE_URL="host=localhost user=postgres password=pass DB.name=postgres port=5432"
> ```  

To run all integration tests:

```bash
go test -run=TestIntegration -tags=integration ./...
```

>**Important** When adding a single resource, it's more common to only run the test for a specific table. You'll need to ensure your resource has the relevant Terraform service deployed. 

#### Adding new Terraform Files Guidelines

Terraform files are organized under the [`terraform`](../../terraform/) folder, and each service has its own folder.
Under each service folder, we organize files into 3 folders:
- `local`: When testing locally run the Terraform CLI from here
- `modules/tests`: Terraform resource and module definitions go here
- `prod`: This folder is used for our CI testing. See relevant scripts [here](../../scripts/). **Not to be used locally**

>Each service has its own Terraform to follow best practices. It allows creating a test environment for each service, and avoids slowdowns and memory issues if we would have had a single Terraform file for all services.

There are a few good rules of thumb to follow when creating new terraform resources that will be served as testing infrastructure:
* If possible make all resources private.
* Make sure to replace built-in plain text passwords with `random_password` generator
* For every compute/db try to use the smallest size to keep the cost low
* If autoscaling option is present, always turn it off

If you want to apply the Terraform locally first before pushing it to CI and applying there, use:

```bash
cd terraform/YOUR_SERVICE_NAME/local
terraform init
# Replace AB with your own initials so multiple team members can work on the same account without conflicting resources
terraform apply -var="prefix=AB"
go test -run="TestIntegration/ROOT_TABLE_NAME" -tags=integration ./...
```
