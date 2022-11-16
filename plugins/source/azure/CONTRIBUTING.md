# Azure Source Plugin Contribution Guide

Thanks for contributing to CloudQuery! You are awesome. This document serves as a guide for adding new services and resources to the Azure source plugin.

There are two steps to adding a new Azure resource:

1. [Generate interfaces for the Azure SDK function(s) that fetch the resource](#1-generate-interfaces-for-the-azure-sdk-functions-that-fetch-the-resource)
2. [Add a code generation recipe](#2-add-a-code-generation-recipe)

As a prerequisite, ensure that API calls to list/describe the desired resource are present in [the old Azure client](https://github.com/Azure/azure-sdk-for-go#client-previous-versions) or [the new one](https://github.com/Azure/azure-sdk-for-go#management-new-releases), and make note of:

- to which Azure service the resource belongs
- the schema of the returned object(s)

## 1. Generate Interfaces for the Azure SDK Function(s) that Fetch the Resource

### Before you Start

Inside the Azure plugin directory, run:

```shell
make install-tools
```

This will install `mockgen` and any other tools necessary to complete the process.

### Generate the Service Interface

1. Check in [`client/services`](client/services) that the service you need has client and interfaces defined. If it does, you can skip to [2. Add a Code Generation Recipe](#2-add-a-code-generation-recipe).
2. If the service does not exist, create a new file in [`client/services`](client/services) named `service_name.go`.
3. Add the relevant Azure SDK import to the top of the file.
4. Create the new service client and relevant interfaces. It's best to look at an existing service, for the old SDK example see [here](client/services/authorization.go), and for the new SDK  see [here](client/services/subscriptions.go).
5. Ensure the new service has a `//go:generate mockgen` (see examples from above) and run `make generate` to generate the mocks.

## 2. Add a Code Generation Recipe

Every supported Azure service has a recipe file under [`codegen/recipes`](codegen/recipes). For example, all Web resources are listed in [`codegen/recipes/web.go`](codegen/recipes/web.go).

In the following example, we will use the fictional `MyService` Azure service with `MyResource` resource as an example. We recommend taking a look at a few examples in [codegen/recipes](codegen/recipes) first, as these steps will make more sense with some examples to reference.

If you are adding a service that needs a new recipe, see [Add a New Recipe File](#add-a-new-recipe-file). Otherwise, if the Azure service is already supported but is missing resource(s), you may skip to [Add a Resource to a Recipe](#add-a-resource-to-a-recipe).

### Add a New Recipe File

The process to follow for adding a new recipe is:

1. Add a new file under [`codegen/recipes`](codegen/recipes) called `myservice.go` under a package named `recipes`.
2. Inside the new file, add a function called `MyServiceResources()` that returns `[]Resource`.
3. Call the function from [`codegen/recipes/base.go`](codegen/recipes/base.go#L119) by adding `resources = append(resources, MyServiceResources()...)`
4. Define the list of resources to be generated and return it inside this function. See [Add a Resource to a Recipe](#add-a-resource-to-a-recipe) for more details.

### Add a Resource to a Recipe

`MyServiceResources()` should return an array of `Resource` instances. Many of the fields or a `Resource` can be auto filled by using the `byTemplates` helper `struct`, and calling `generateResources(byTemplatesInstance)`.
`byTemplates` is a mapping of templates to apply on an array of resource definitions. For example:

```go
func Batch() []Resource {
	var resourcesByTemplates = []byTemplates{
		{
			templates: []template{
				{
					source:            "resource_list.go.tpl",
					destinationSuffix: ".go",
					imports:           []string{},
				},
				{
					source:            "resource_list_mock_test.go.tpl",
					destinationSuffix: "_mock_test.go",
					imports:           []string{"github.com/Azure/azure-sdk-for-go/services/batch/mgmt/2021-06-01/batch"},
				},
			},
			definitions: []resourceDefinition{
				{
					azureStruct:  &batch.Account{},
					listFunction: "List",
				},
			},
		},
	}

	return generateResources(resourcesByTemplates)
}
```

Will auto fill the required fields for the `Resource` `struct` for each template, based on the `azureStruct` and `listFunction` fields.
`listFunction` defaults to `ListAll`, making the `azureStruct` the only required field.

#### All Available Resource Definition Fields

All available resource definition fields can be seen in [`codegen/recipes/base.go`](codegen/recipes/base.go#L60). 
**For top level tables figuring out the `listFunction` and [templates](codegen/templates/) to use is usually enough.**
For relations you'll probably need to define the `listFunctionArgs`, `listFunctionArgsInit` and `relations` too. See an example [here](codegen/recipes/postgressql.go).

### Run Code Generation

With the recipe file added and some resources defined, you are ready to run code generation. Run:

```shell
make gen-code
```

This will update all resources and generate a new directory for your service under [resources/services](resources/services).
It should create the table, resolver and mock tests for the resource.
Once everything is generated, you might need to update the recipe to handle any compilation errors by updating `listFunction` for example.

## General Tips

- Keep transformations to a minimum. As far as possible, we aim to deliver an accurate reflection of what the Azure API provides.
- We generally only unroll structs one level deep. Nested structs should be transformed into JSON columns.
- It's recommended to split each resource addition into a separate PR. This makes it easier to review and merge.
- If the Azure SDK API is not consistent, it's recommended to wrap it inside [`client/services`](client/services) to simplify the code generation process. See example [here](client/services/web.go#L61).
- Before submitting a pull request, run `make gen-docs` to generate documentation for the table. Include these generated files in the pull request.
- If you get stuck or need help, feel free to reach out on [Discord](https://www.cloudquery.io/discord). We are a friendly community and would love to help!
