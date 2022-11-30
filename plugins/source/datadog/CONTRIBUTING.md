# Datadog Source Plugin Contribution Guide

Thanks for contributing to CloudQuery! You are awesome. This document serves as a guide for adding new services and resources to the Datadog source plugin.

There are two steps to adding a new Datadog resource:

1. [Generate interfaces for the Datadog SDK function(s) that fetch the resource](#1-generate-interfaces-for-the-datadog-sdk-functions-that-fetch-the-resource)
2. [Add a code generation recipe](#2-add-a-code-generation-recipe)


## 1. Generate Interfaces for the Datadog SDK Function(s) that Fetch the Resource

### Before you Start

Inside the Datadog plugin directory, run:

```shell
make install-tools
```

This will install `mockgen` and any other tools necessary to complete the process.

### Generate the Service Interface

1. Check in [`client/services`](client/services) that the service you need has client and interfaces defined. If it does, you can skip to [2. Add a Code Generation Recipe](#2-add-a-code-generation-recipe).
2. If the service does not exist, add an instance of service you want to add to [`codegen/services/clients.go`](codegen/services/clients.go).
3. Add the relevant Datadog SDK import to the top of the file.
4. [Run Code Generation](#run-code-generation) to generate the service interfaces.
5. Ensure the new service has a `//go:generate mockgen` (see examples from above) and run `make generate` to generate the mocks.

## 2. Add a Code Generation Recipe

Every supported Datadog service has a recipe file under [`codegen/recipes`](codegen/recipes). For example, all Users resources are listed in [`codegen/recipes/users.go`](codegen/recipes/users.go).

In the following example, we will use the fictional `MyService` Datadog service with `MyResource` resource as an example. We recommend taking a look at a few examples in [codegen/recipes](codegen/recipes) first, as these steps will make more sense with some examples to reference.

If you are adding a service that needs a new recipe, see [Add a New Recipe File](#add-a-new-recipe-file). Otherwise, if the Datadog service is already supported but is missing resource(s), you may skip to [Add a Resource to a Recipe](#add-a-resource-to-a-recipe).

### Add a New Recipe File

The process to follow for adding a new recipe is:

1. Add a new file under [`codegen/recipes`](codegen/recipes) called `myservice.go` under a package named `recipes`.
2. Inside the new file, add a function called `MyServiceResources()` that returns `[]Resource`.
3. Add my service to all resources by adding `recipes.MyServiceResources,` to [`codegen/main.go`](codegen/main.go#L17)
4. Define the list of resources to be generated. See [Add a Resource to a Recipe](#add-a-resource-to-a-recipe) for more details.

### Add a Resource to a Recipe

`MyServiceResources()` should return an array of `Resource` instances. Like on example below

```go
func MyServiceResources() []*Resource {
	resources := []*Resource{
		{
            SubService: "my_service",
            Multiplex:  "client.AccountMultiplex",
            Struct:     new(datadogV2.User),
            SkipFields: []string{"Id"},
            ExtraColumns: append(defaultAccountColumns, codegen.ColumnDefinition{
                    Name:     "id",
                    Type:     schema.TypeString,
                    Resolver: `schema.PathResolver("Id")`,
                    Options:  schema.ColumnCreationOptions{PrimaryKey: true},
                }),
            Relations: []string{"MyServiceAttachments()", "MyServicePermissions()"},
		},
        {
            SubService:   "my_service_attachments",
            Struct:       new(datadogV2.Permission),
            ExtraColumns: defaultAccountColumns,
        },
        {
            SubService:   "my_service_permissions",
            Struct:       new(datadogV2.User),
            ExtraColumns: defaultAccountColumns,
        },
    }
    
    // set default values
    for _, r := range resources {
        r.Service = "my_service"
    }
    return resources
}
```

### Run Code Generation

With the recipe file added and some resources defined, you are ready to run code generation. Run:

```shell
make gen-code
```

This will update all resources and generate a new directory for your service under [resources/services](resources/services).
It should create the table files for your resources. 
For each table fetch and mock test files should be added. for example for  `my_service.go` the  `my_service_fetch.go` and `my_service_mock_test.go` files should be created

## General Tips

- Keep transformations to a minimum. As far as possible, we aim to deliver an accurate reflection of what the Datadog API provides.
- We generally only unroll structs one level deep. Nested structs should be transformed into JSON columns.
- It's recommended to split each resource addition into a separate PR. This makes it easier to review and merge.
- Before submitting a pull request, run `make gen-docs` to generate documentation for the table. Include these generated files in the pull request.
- If you get stuck or need help, feel free to reach out on [Discord](https://www.cloudquery.io/discord). We are a friendly community and would love to help!
