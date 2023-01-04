# Fastly Source Plugin Contribution Guide

Thanks for contributing to CloudQuery! You are awesome. This document serves as a guide for adding new services and resources to the Fastly source plugin.

There are two main steps to adding a new resource:

1. [Add a code generation recipe](#1-add-a-code-generation-recipe)
2. [Write the resolver function to fetch the resource using the AWS SDK](#2-setting-up-the-resource)

## 1. Add a Code Generation Recipe

Every supported Fastly service has a recipe file under [codegen/recipes](codegen/recipes). For example, all service API resources are listed in [codegen/recipes/services.go](codegen/recipes/services.go). 

In the following examples, we will use the fictional `MyService` service with `MyResource` resource as an example. We recommend taking a look at a few examples in [codegen/recipes](codegen/recipes) first, as these steps will make more sense with some examples to reference. 

If you are adding a service that needs a new recipe, see [Add a New Recipe File](#add-a-new-recipe-file). Otherwise, if the service is already supported but is missing resource(s), you may skip to [Add a Resource to a Recipe](#add-a-resource-to-a-recipe).

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

 1. `Service`: This will be the directory name under [resources/services](resources/services) where the resource will be generated, and the first part of the table name, by default.
 2. `DataStruct`: This should be a pointer to the struct that will be synced to the destination. CloudQuery's plugin-sdk code generation will read the fields of this struct and convert it to a `Table` instance with appropriate column types. If `TableName` is omitted, the name of the struct will be used to infer the table name.
 3. `Description`: A link to the API documentation for the resource, as well as (optionally) any other relevant information about the table.

#### All Available Resource Fields

All available Resource fields can be seen in [base.go](codegen/recipes/base.go). These closely map to the `Table` fields [defined by the plugin-sdk](https://github.com/cloudquery/plugin-sdk/blob/main/schema/table.go), which you may use as a further reference about what each field does.

#### Common Fields

If all the resources share the same value for a field (as is often the case for `Service`), our convention is to reduce boilerplate by setting these properties in a loop after defining the resources slice, e.g.

```go
// set default values
for _, r := range resources {
    r.Service = "myservice"
}
```

### Run Code Generation

With the recipe file added and some resources defined, you are ready to run code generation. Run:

```shell
make gen
```

This will update all resources and generate a new directory for your service under [resources/services](resources/services).

## 2. Setting up the resource

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
   c := meta.(*client.Client)
   ```
   With this in hand, complete the resolver function to fetch all resources. After resources are retrieved, send them to the `res` channel for the SDK to deliver to all destinations.
3. Implement a mock test in `myresource_mock_test.go`. We will not describe this in detail here; look at a few examples for similar resources to get you started.

We recommend looking at other resources similar to yours to get an idea of what needs to be done in this step.  

### Implementing Resolver Functions

A few important things to note when adding functions that call the Fastly API:

- If possible, always use an API call that allows you to fetch many resources at once
- Take pagination into account. Use `Paginator`s if the Fastly service supports it. Ensure you fetch **all** the resources.
- Columns may also have their own resolver functions (not covered in this guide). This may be used for simple transformations or when additional calls can help add further context to the table.

## General Tips

- Keep transformations to a minimum. As far as possible, we aim to deliver an accurate reflection of what the API provides.
- We generally only unroll structs one level deep. Nested structs should be transformed into JSON columns. 
- Primary keys are used to upsert rows. Make sure the primary key uniquely defines an entry, otherwise some rows may be overwritten by accident. 
- Before submitting a pull request, run `make gen` to generate code and documentation for the table. Include these generated files in the pull request.
- If you get stuck or need help, feel free to reach out on [Discord](https://www.cloudquery.io/discord). We are a friendly community and would love to help!
