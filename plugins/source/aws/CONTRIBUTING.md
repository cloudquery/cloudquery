# AWS Source Plugin Contribution Guide

Thanks for contributing to CloudQuery! You are awesome. This document serves as a guide for adding new services and resources to the AWS source plugin.

There are two main steps to adding a new AWS resource:

1. [Generate interfaces for the AWS SDK function(s) that fetch the resource](#step-1-generate-interfaces-for-the-aws-sdk-functions-that-fetch-the-resource)
2. [Add a new table](#step-2-add-a-new-table)

As a prerequisite, in [aws-sdk-go-v2](https://pkg.go.dev/github.com/aws/aws-sdk-go-v2) ensure API calls exist to list/describe the desired resource, and make note of:

- to which aws service the resource belongs
- the schema of the returned object(s)

## Step 1. Generate Interfaces for the AWS SDK Function(s) that Fetch the Resource

### Generate the Service Interface (if it doesn't exist)

1. Check in [client/services.go](client/services.go) that the service you need has an interface defined. If it does, you can skip to [Step 2](#step-2-add-a-new-table). If not, read on to learn how to generate the interface.
2. Inside [codegen/services/clients.go](codegen/services/clients.go), add the client for the AWS SDK you need to the `clients` slice. You may need to run `go get github.com/aws/aws-sdk-go-v2/service/<service-name>` (e.g. `go get github.com/aws/aws-sdk-go-v2/service/dynamodb`) to add the dependency first.
3. Run `make gen-mocks`. This takes a few seconds, but it should add the interface for your client to [client/services.go](client/services.go) and create a mock for it that will be used in unit tests later.

### Step 2. Add a New Table

The process to follow for adding a new table is:

1. Add a new directory matching the AWS service name under [resources/services](resources/services) (e.g. `resources/services/newservice`), if one doesn't exist already
2. Create a new file under the new directory with the name of the resource (e.g. `resources/services/newservice/myresource.go`) and add a function that returns `*schema.Table`. The easiest is to copy-paste an existing table as a starting point ([`Kinesis`](resources/services/kinesis/kinesis.go) is a good example).
3. **Important**: Add a call to the new function to the list of tables in [tables.go](resources/plugin/tables.go). Otherwise, the new table will not be included in the plugin.  
4. Update all the fields, taking special care to ensure that the `transformers.TransformWithStruct()` call in the `Resolver` function has the correct struct type (e.g. `transformers.TransformWithStruct(&types.MyResource{})`)
5. Implement the resolver function. This should have the signature: 
   ```go
   func fetchMyResource(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
       // TODO: implement this
   }
   ```
   
   The easiest is to copy-paste an existing resolver as a starting point. (Again, [`Kinesis`](resources/services/kinesis/streams_fetch.go) is a good example.)
   
   You may use a type assertion on `meta` to obtain a reference to your interface functions, e.g.:
   ```go
   svc := meta.(*client.Client).Services().MyService
   ```
   
   With this in hand, complete the resolver function to fetch all resources. After resources are retrieved, send them to the `res` channel for the SDK to deliver to all destinations.
6. Implement a mock test in `myresource_mock_test.go`. We will not describe this in detail here; look at a few examples for similar resources to get you started.

We highly recommend looking at other resources similar to yours to get an idea of what needs to be done in each step.  

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
