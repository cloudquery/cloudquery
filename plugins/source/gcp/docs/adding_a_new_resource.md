# Adding a new resource

Some information can be found in the [docs for developing a new provider](https://docs.cloudquery.io/developers/developing-new-provider).

As a prerequisite, in [google.golang.org/api](https://pkg.go.dev/google.golang.org/api) ensure API calls exist to list/describe the desired resource, and make note of:

   - to which GCP service the resource belongs
   - the schema of the returned object(s)

## Setting up the service

If the service to which the resource belongs has not been used before in cq-provider-gcp, there are a few steps that need to be done to configure it.

1. Create the service interface in [client/services.go](../../client/services.go)
2. Add the service to the `Services` struct in the [client/services.go](../../client/services.go)
3. Init the service in the `initServices` function in [client/services.go](../../client/services.go)

## Setting up the resource

### Skeleton

1. Create a file under `resources/` that follows the pattern of `resources/<service>/<resource_name>`.
2. In that file, create a function that returns a `*schema.Table`
3. In [resources/provider.go](./resources/provider.go), add a mapping between the function you just created and the name of the resource that will be used in the config yml file.
4. Add a test file at `resources/<service>/<resource>_test.go`. Follow other examples to create a test for the resource.
5. Run `go run docs/docs.go` to generate the documentation for the new resource

### Implementation

Now that the skeleton has been set up, you can start to actually implement the resource. This consists of two parts: 

1. Defining the schema
2. Implementing resolver functions

#### Defining the schema

It is recommended that you look at a few existing resources as examples and also read through the comments on the source code for the [Table](https://github.com/cloudquery/plugin-sdk/blob/main/provider/schema/table.go) and [Column](https://github.com/cloudquery/plugin-sdk/blob/main/provider/schema/column.go) implementations for details.

For noncomplex fields, the SDK can directly resolve them into `Column`s for you, so all you need to do is specify the `Name` and the `Type`.

For complex fields or fields that require further API calls, you can defined your own `Resolver` for the `Column`.

#### Implementing Resolver Functions

A few important things to note when adding functions that call the GCP API:

- If possible, always use an API call that allows you to fetch many resources at once
- Take pagination into account. Ensure you fetch **all** of the resources

#### Integration Tests

To Run Integration tests please run

```bash
docker run -p 5432:5432 -e POSTGRES_PASSWORD=pass -d  postgres:13.3
gcloud auth application-default login
# To run against our test environment (same as in the CI), run: gcloud config set project cq-provider-gcp
# team members have write access to cq-playground (to apply new terraform files) gcloud config set project cq-playground
# Otherwise just use your development project via gcloud config set project YOUR_PROJECT
go test -run=TestIntegration -tags=integration ./...
```

To Run integration test for specific table:

```bash
go test -run="TestIntegration/ROOT_TABLE_NAME" -tags=integration ./...
# For example
go test -run="TestIntegration/gcp_compute_instances" -tags=integration ./...
```

## Adding new Terraform File Guidelines

There a few good rule of thumb to follow when creating new terraform resources that will be served as testing infrastructure:
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
```
