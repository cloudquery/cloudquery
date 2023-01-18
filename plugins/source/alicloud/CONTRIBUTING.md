# Alibaba Cloud Source Plugin Contribution Guide

Thanks for contributing to CloudQuery! You are awesome. This document serves as a guide for adding new services and resources to the Alibaba Cloud source plugin.

There are two main steps to adding a new Alibaba Cloud resource:

1. [Generate interfaces for the Alibaba Cloud SDK function(s) that fetch the resource](#step-1-generate-interfaces-for-the-alibaba-cloud-sdk-functions-that-fetch-the-resource)
2. [Add a new table](#step-2-add-a-new-table)

## Step 1. Generate Interfaces for the Alibaba Cloud SDK Function(s) that Fetch the Resource

### Generate the Service Interface (if it doesn't exist)

1. Check in [client/services.go](client/services.go) that the service you need has an interface defined. If it does, you can skip to [2. Add a New Table](#2-add-a-new-table). If not, read on to learn how to generate the interface.
2. Inside [codegen/main.go](codegen/main.go), add the client for the Alibaba Cloud SDK you need to the `clients` slice. You may need to run `go get` to add the dependency first.
3. Run `make gen-mocks`. This takes a few seconds, but it should add the interface and mocks for your client to the [client/services](client/services) and [client/mocks](client/mocks) directories.
4. Add an entry for the service to the `Services` struct in [client/services.go](client/services.go) and instantiate it inside `initServices`. 

### Step 2. Add a New Table

The process to follow for adding a new table is:

1. Add a new directory matching the Alibaba Cloud service name under [resources/services](resources/services) (e.g. `resources/services/newservice`), if one doesn't exist already
2. Create a new file under the new directory with the name of the resource (e.g. `resources/services/newservice/myresource.go`) and add a function that returns `*schema.Table`. The easiest is to copy-paste an existing table as a starting point ([`ECS Instances`](resources/services/ecs/instances.go) is a good example).
3. **Important**: Add a call to the new function to the list of tables in [plugin.go](resources/plugin/plugin.go). Otherwise, the new table will not be included in the plugin.  
4. Update all the fields, taking special care to ensure that the `transformers.TransformWithStruct()` call in the `Resolver` function has the correct struct type (e.g. `transformers.TransformWithStruct(&types.MyResource{})`). **Pro Tip**: You can run `make gen-docs` at any point to create documentation for your table under [docs/tables](docs/tables). This is a good way to check that the columns and their types match your expectations.
5. Implement the resolver function. This should have the signature: 
   ```go
   func fetchMyResource(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
       // TODO: implement this
   }
   ```
   
   The easiest is to copy-paste an existing resolver as a starting point. (Again, [`ECS Instances`](resources/services/ecs/instances.go) is a good example.)
   
   You may use a type assertion on `meta` to obtain a reference to your interface functions, e.g.:
   ```go
   svc := meta.(*client.Client).Services().MyService
   ```
   
   With this in hand, complete the resolver function to fetch all resources. After resources are retrieved, send them to the `res` channel for the SDK to deliver to all destinations.
6. Implement a mock test in `myresource_mock_test.go`. We will not describe this in detail here; look at a few examples for similar resources to get you started.

We highly recommend looking at other resources similar to yours to get an idea of what needs to be done in each step.  

### Implementing Resolver Functions

A few important things to note when adding functions that call the Alibaba Cloud API:

- If possible, always use an API call that allows you to fetch many resources at once
- Take pagination into account. Ensure you fetch **all** the resources.
- Columns may also have their own resolver functions (not covered in this guide). This may be used for simple transformations or when additional calls can help add further context to the table.

## General Tips

- Keep transformations to a minimum. As far as possible, we aim to deliver an accurate reflection of what the Alibaba Cloud API provides.
- We generally only unroll structs one level deep. Nested structs should be transformed into JSON columns. 
- Make sure the resource has a `tags` JSON column (if possible). Sometimes this requires additional SDK calls.
- Before submitting a pull request, run `make gen-docs` to generate documentation for the table. Include these generated files in the pull request.
- If you get stuck or need help, feel free to reach out on [Discord](https://www.cloudquery.io/discord). We are a friendly community and would love to help!
