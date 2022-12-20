# AWS Source Plugin Contribution Guide

Thanks for contributing to CloudQuery! You are awesome. This document serves as a guide for adding new services and resources to the AWS source plugin.

There are three main steps to adding a new AWS resource:

1. [Generate interfaces for the AWS SDK function(s) that fetch the resource](#1-generate-interfaces-for-the-aws-sdk-functions-that-fetch-the-resource)
2. [Add a code generation recipe](#2-add-a-code-generation-recipe)
3. [Write the resolver function to fetch the resource using the AWS SDK](#3-setting-up-the-resource)

As a prerequisite, in [aws-sdk-go-v2](https://pkg.go.dev/github.com/aws/aws-sdk-go-v2) ensure API calls exist to list/describe the desired resource, and make note of:

- to which aws service the resource belongs
- the schema of the returned object(s)

## 1. Generate Interfaces for the AWS SDK Function(s) that Fetch the Resource

### Before you Start

Inside the root of this repository, run:

```shell
make install-tools
```

This will install `mockgen` and any other tools necessary to complete the process.

### Generate the Service Interface

1. Check in [client/services.go](client/services.go) that the service you need has an interface defined. If it does, you can skip to [2. Add a Code Generation Recipe](#2-add-a-code-generation-recipe). If not, read on to learn how to generate the interface.
2. Inside [codegen/services/clients.go](codegen/services/clients.go), add the client for the AWS SDK you need to the `clients` slice. You may need to run `go get github.com/aws/aws-sdk-go-v2/service/<service-name>` (e.g. `go get github.com/aws/aws-sdk-go-v2/service/dynamodb`) to add the dependency first.
3. Run `make gen-mocks`. This should add the interface for your client to [client/services.go](client/services.go) and create a mock for it that will be used in unit tests later.

## 2. Add a Code Generation Recipe

Every supported AWS service has a recipe file under [codegen/recipes](codegen/recipes). For example, all ECS resources are listed in [codegen/recipes/ecs.go](codegen/recipes/ecs.go). 

In the following examples, we will use the fictional `MyService` AWS service with `MyResource` resource as an example. We recommend taking a look at a few examples in [codegen/recipes](codegen/recipes) first, as these steps will make more sense with some examples to reference. 

If you are adding a service that needs a new recipe, see [Add a New Recipe File](#add-a-new-recipe-file). Otherwise, if the AWS service is already supported but is missing resource(s), you may skip to [Add a Resource to a Recipe](#add-a-resource-to-a-recipe).

### Add a New Recipe File

The process to follow for adding a new recipe is:

1. Add a new file under [codegen/recipes](codegen/recipes) called `myservice.go`.
2. Inside the new file, add a function called `MyServiceResources()` that returns `[]*Resource`.
3. Call the function from [codegen/main.go](codegen/main.go) by adding
   `resources = append(resources, recipes.MyServiceResources()...)`
4. Define the list of resources to be generated and return it inside this function. See
   [Add a Resource to a Recipe](#add-a-resource-to-a-recipe) for more details.

### Add a Resource to a Recipe

`MyServiceResources()` should return a slice of `*Resource` instances. Each resource should, at a minimum, have the following fields defined:

 1. `Service`: This will become the table prefix, and will usually be the same as the filename you chose for the recipe.
 2. `SubService`: This will be the final part of the table name, e.g. `aws_myservice_subservice`
 3. `Multiplex`: Most AWS services have resources defined per account and region. In such cases, `client.ServiceAccountRegionMultiplexer("my-service")` is usually the correct multiplexer to use. Look in [client/data/partition_service_region.json](client/data/partition_service_region.json) for the correct service name to use.
 4. `Struct`: This should be a pointer to the struct that will be synced to the destination. CloudQuery's plugin-sdk code generation will read the fields of this struct and convert it to a `Table` instance with appropriate column types.
    
    Deciding which struct to use takes some practice. To find the right struct, explore the return types of the SDK functions provided by the [AWS Go SDK](https://pkg.go.dev/github.com/aws/aws-sdk-go-v2) (and that you created interface functions for in Step 1). Often there will be a `Get` or `Describe` call that returns a `GetResourceOutput` (or similar) struct. Sometimes this Output struct can be used directly. Other times, the Output struct will reference an inner type, which should then be used for defining the Resource.

#### All Available Resource Fields

All available Resource fields can be seen in [base.go](codegen/recipes/base.go). These closely map to the `Table` fields [defined by the plugin-sdk](https://github.com/cloudquery/plugin-sdk/blob/main/schema/table.go), which you may use as a further reference about what each field does.

#### Common Fields

If all the resources share the same value for a field (as is often the case for `Service` and `Multiplex`), our convention is to reduce boilerplate by setting these properties in a loop after defining the resources slice, e.g.

```go
// set default values
for _, r := range resources {
    r.Service = "myservice"
    r.Multiplex = `client.ServiceAccountRegionMultiplexer("my-service")`
}
```

### Run Code Generation

With the recipe file added and some resources defined, you are ready to run code generation. Run:

```shell
make gen-code
```

This will update all resources and generate a new directory for your service under [resources/services](resources/services).

## 3. Setting up the resource

By following the steps outlined above, you should now have generated a `myservice` directory under `resources/services`, containing a file called `myresource.go` (these names are examples, your actual filenames will differ). We will now set up the resource. This involves two steps: refining the `codegen` recipe, and writing one or more resolver functions.

1. Open the generated `myservice/myresource.go` and inspect the `schema.Table` that is being returned. Does it contain the appropriate columns for the resource? Does it have a primary key? If something looks off, return to the recipe for this resource (under [codegen/recipes](codegen/recipes)) and make adjustments. Then re-run code generation as described in [Run Code Generation](#run-code-generation). Repeat this process until the Table looks right.
2. Your generated `Table` will reference a `Resolver` function that needs to be implemented. The generated file should never be edited directly, so create a new file called `myresource_fetch.go` and add a resolver function with the following signature:
   ```go
   func fetchMyResource(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
       // TODO: implement this
   }
   ```
   You may use a type assertion on `meta` to obtain a reference to your interface functions, e.g.:
   ```go
   svc := meta.(*client.Client).Services().MyService
   ```
   With this in hand, complete the resolver function to fetch all resources. After resources are retrieved, send them to the `res` channel for the SDK to deliver to all destinations.
3. Implement a mock test in `myresource_mock_test.go`. We will not describe this in detail here; look at a few examples for similar resources to get you started.

We recommend looking at other resources similar to yours to get an idea of what needs to be done in this step.  

### Implementing Resolver Functions

A few important things to note when adding functions that call the AWS API:

- If possible, always use an API call that allows you to fetch many resources at once
- Take pagination into account. Use `Paginator`s if the AWS service supports it. Ensure you fetch **all** the resources.
- Columns may also have their own resolver functions (not covered in this guide). This may be used for simple transformations or when additional calls can help add further context to the table.
- Many resources require a `List` call, followed by a `Describe` call. Look for examples using `PreResourceResolver` to see the canonical way of implementing this. (In short: the table resolver function will call `List`, while the `PreResourceResolver`, called once per resource, will call `Describe`)


## General Tips

- Keep transformations to a minimum. As far as possible, we aim to deliver an accurate reflection of what the AWS API provides.
- We generally only unroll structs one level deep. Nested structs should be transformed into JSON columns. 
- For consistency, make sure the resource has an `ARN` stored in a column named `arn`. Sometimes this means using the AWS SDK to generate an ARN for the resource.
- Make sure the resource has a `tags` JSON column (if possible). Sometimes this requires additional SDK calls. Sometimes the column needs to be renamed from `tag_list` to `tags` (and converted to a map). There are custom `ResolveTags` and `ResolveTagFields` resolvers to help with this. It's not always possible, but we try to keep the `tags` column consistent across AWS resources.
- Before submitting a pull request, run `make gen-docs` to generate documentation for the table. Include these generated files in the pull request.
- If you get stuck or need help, feel free to reach out on [Discord](https://www.cloudquery.io/discord). We are a friendly community and would love to help!
