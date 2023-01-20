# GCP Source Plugin Contribution Guide

Thanks for contributing to CloudQuery! You are awesome. This document serves as a guide for adding new services and resources to the GCP source plugin.

There are two main steps to adding a new GCP resource:

1. [Add a code generation recipe](#1-add-a-code-generation-recipe)
2. [Writing the resolver function to fetch the resource using the AWS SDK](#2-setting-up-the-resource)

## 1. Add a Code Generation Recipe

Every supported GCP service has a recipe file under [codegen/recipes](codegen/recipes).

In the following examples, we will use the fictional `MyService` GCP service with `MyResource` resource as an example. We recommend taking a look at a few examples in [codegen/recipes](codegen/recipes) first, as these steps will make more sense with some examples to reference.

If you are adding a service that needs a new recipe, see [Add a New Recipe File](#add-a-new-recipe-file). Otherwise, if the GCP service is already supported but is missing resource(s), you may skip to [Add a Resource to a Recipe](#add-a-resource-to-a-recipe).

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
2. `SubService`: This will be the final part of the table name, e.g. `gcp_myservice_subservice`
3. `Struct`: This should be a pointer to the struct that will be synced to the destination. CloudQuery's plugin-sdk code generation will read the fields of this struct and convert it to a `Table` instance with appropriate column types.


### Choosing a Multiplexer

In the GCP plugin there are three types of multiplexers. Every top level resource needs to use multiplexer:

1. `ProjectMultiplex` (_default_): This is the most basic of multiplexers in that it will resolve the resource in each project that is being synced. 
2. `ProjectMultiplexEnabledServices(serviceDNS string)`:  This multiplexer will only attempt to sync a resource if that project has the service enabled otherwise the resource will be skipped for that specific projectID. On top of this the user must also enable the feature via `enabled_services_only: true` in the spec. In order to use this multiplexer you must specify a valid `resource.ServiceDNS`
3. `client.OrgMultiplex`: For resources that are unique across an entire Organization. In order to use this multiplexer you have to explicitly set the multiplexer `client.OrgMultiplex
``` go
var OrgMultiplex = "client.OrgMultiplex"
Resource{
	Multiplex:  &OrgMultiplex   
}
```

#### All Available Resource Fields

All available Resource fields can be seen in [base.go](codegen/recipes/base.go). See the documentation for each field for an explanation of what it does.

#### Common Fields

If all the resources share the same value for a field (as is often the case for `Service` and `Multiplex`), our convention is to reduce boilerplate by setting these properties in a loop after defining the resources slice, e.g.

```go
// set default values
for _, r := range resources {
    r.Service = "myservice"
}
```

### Run Code Generation

With the recipe file added and some resources defined, you are ready to run `codegen`. Inside the [`codegen`](codegen) directory, run:

```shell
go run main.go
```

This will update all resources and generate a new directory for your service under [resources/services](resources/services).

## 2. Setting up the resource

By following the steps outlined above, you should now have generated a `myservice` directory under `resources/services`, containing a file called `myresource.go` (these names are examples, your actual filenames will differ). We will now set up the resource. This involves two steps: refining the `codegen` recipe, and writing one or more resolver functions.

1. Open the generated `myservice/myresource.go` and inspect the `schema.Table` that is being returned. Does it contain the appropriate columns for the resource? Does it have a primary key? If something looks off, return to the recipe for this resource (under [codegen/recipes](codegen/recipes)) and make adjustments. Then re-run code generation as described in [Run Code Generation](#run-code-generation). Repeat this process until the Table looks right.
2. Your generated `Table` will reference a `Resolver` function that needs to be implemented. Sometimes this resolver can also be generated, and sometimes it will need to be written by hand.
3. Finally, implement a mock test in `myresource_mock_test.go`. This can also often be generated, but not always.

We recommend looking at other resources similar to yours to get an idea of what needs to be done in this step.  

### Implementing Resolver Functions

A few important things to note when adding functions that call the GCP API:

- If possible, always use an API call that allows you to fetch many resources at once
- Take pagination into account. Ensure you fetch **all** the resources.
- Columns may also have their own resolver functions (not covered in this guide). This may be used for simple transformations or when additional calls can help add further context to the table.

## General Tips

- Keep transformations to a minimum. As far as possible, we aim to deliver an accurate reflection of what the API provides.
- We generally only unroll structs one level deep. Nested structs should be transformed into JSON columns.
- If you get stuck or need help, feel free to reach out on [Discord](https://www.cloudquery.io/discord). We are a friendly community and would love to help!
