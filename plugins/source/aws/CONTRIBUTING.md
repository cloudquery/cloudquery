# AWS Source Plugin Contribution Guide

There are two main steps to adding a new AWS resource:
 1. Add a code generation recipe 
 2. Writing the resolver function to fetch the resource using the AWS SDK

## Add a Code Generation Recipe

Every supported AWS service has a recipe file under [codegen/recipes](codegen/recipes). For example, all ECS resources
are listed in [codegen/recipes/ecs.go](codegen/recipes/ecs.go). 

In the following examples, we will use the fictional `MyService` AWS service as an example. We recommend 
taking a look at a few examples in [codegen/recipes](codegen/recipes) first, as these steps will make more sense with 
some examples in mind. 

If you are adding a service that needs a new recipe, see [Add a New Recipe File](#add-a-new-recipe-file). Otherwise, if
the AWS service is already supported but is missing resource(s), you may skip to [Add a Resource to a Recipe](#add-a-resource-to-a-recipe).

### Add a New Recipe File

Below is the process to follow for adding a new recipe. 

1. Add a new file under [codegen/recipes](codegen/recipes) called `myservice.go`.
2. Inside the new file, add a function called `MyServiceResources()` that returns `[]*Resource`.
3. Call the function from [codegen/main.go](codegen/main.go) by adding
   `resources = append(resources, recipes.MyServiceResources()...)`
4. Define the list of resources to be generated and return it inside this function. See
   [Add a Resource to a Recipe](#add-a-resource-to-a-recipe) for more details.

### Add a Resource to a Recipe

`MyServiceResources()` should return a slice of `*Resource` instances. Each resource should, at a minimum, have the 
following fields defined:
 1. `Service`: This will become the table prefix, and will usually be the same as the filename you chose for the recipe.
 2. `SubService`: This will be the final part of the table name, e.g. `aws_myservice_subservice`
 3. `Multiplex`: Most AWS services have resources defined per account and region. In such cases, 
     `client.ServiceAccountRegionMultiplexer("my-service")` is usually the correct multiplexer to use. Look in
     [client/data/partition_service_region.json](client/data/partition_service_region.json) for the correct service name
     to use.
 4. `Struct`

If all the resources share the same value for a field, as is usually the case for `Service` and `Multiplex`, our 
convention is to set these properties in a loop after defining the slice, e.g.

```
// set default values
for _, r := range resources {
    r.Service = "myservice"
    r.Multiplex = `client.ServiceAccountRegionMultiplexer("my-service")`
}
```

### Run Code Generation

With the recipe file added and some resources defined, you are ready to run codegen. Inside [codegen](codegen), run:

```shell
go run main.go
```

This will update all resources and generate a new directory for your service under [resources/services](resources/services).

## Setting up the service

## Setting up the resource

### Implementation

#### Defining the schema

#### Implementing Resolver Functions

A few important things to note when adding functions that call the AWS API:

- If possible, always use an API call that allows you to fetch many resources at once
- Take pagination into account. Ensure you fetch **all** the resources. Prefer using Paginators if the resource supports it.
