# Datadog Source Plugin Contribution Guide

Thanks for contributing to CloudQuery! You are awesome. This document serves as a guide for adding new services and resources to the Datadog source plugin.

## Prerequisites

 - A working Go installation (1.19+). See [here](https://go.dev/doc/install) for instructions.

## Adding a new resource

### Step 1. Add the service to the client

When adding support for new APIs, you may need to add a new service to the client. To do this:
 1. Open `codegen/services/clients.go` and add the client from the Datadog SDK to the list of clients
 2. Run `make gen-mocks` to generate the service interface and mocks

### Step 2. Add the resource to the plugin

To add a new resource, you will need to add a new directory under the `resources/services` directory, if a directory for the service does not already exist. Next, add a Go file for the resource. If you are adding a Datadog resource called `Bar` as part of the `Foo` service, name the directory `foo` and call the file `bar.go`. Inside `bar.go`, create a function that returns a `Table`:

```go filename="resources/services/foo/bar.go"
package users

import (
	"context"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	"github.com/cloudquery/cloudquery/plugins/source/datadog/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func Bar() *schema.Table {
	return &schema.Table{
		// name of the table
		Name:      "datadog_foo_bar",
		// the resolver function is responsible for fetching data from the API
		Resolver:  fetchBar,
		// columns will be automatically created from the given struct
		Transform: client.TransformWithStruct(&datadogV2.Bar{}),
		// define additional columns here, or override the default columns
		Columns: []schema.Column{},
	}
}
```

Next, implement the resolver function:

```go filename="resources/services/foo/bar.go"
func fetchBar(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	ctx = c.BuildContextV2(ctx)
	resp, _, err := c.DDServices.FooAPI.ListBars(ctx)
	if err != nil {
		return err
	}
	res <- resp.GetData()
	return nil
}
```

Finally, add another file called `bar_mock_test.go` in the same directory. This file will test the resolver function using a mock. The easiest is to copy a test from another resource and modify it to suit the new resource.

## General Tips

- Keep transformations to a minimum. As far as possible, we aim to deliver an accurate reflection of what the Datadog API provides.
- We generally only unroll structs one level deep. Nested structs should be transformed into JSON columns.
- It's recommended to split each resource addition into a separate PR. This makes it easier to review and merge.
- Before submitting a pull request, run `make gen-docs` to generate documentation for the table. Include these generated files in the pull request.
- If you get stuck or need help, feel free to reach out on [Discord](https://www.cloudquery.io/discord). We are a friendly community and would love to help!
