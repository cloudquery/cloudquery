# Testing

Testing providers is key factor in ensuring the quality of your provider and developer experience, so future contributors can easily test their changes.

CloudQuery SDK provides a utility test function to help you test each resource and verify it's output.

## Requirements

The testing framework provided by CloudQuery SDK are essentially end-to-end (integration) tests. This means you will need a real test account and real resources in place. It is suggested to create those test resources with any IaC tool (Terraform, Pulumi, etc.).

## How Resource Tests work

There are two main goals to CloudQuery testing framework:

- Check that the extract(fetch) functionality is defined correctly and works without errors against real infrastructures.
- Check the schema definition and all column resolvers are defined correctly, the data exists after a successful fetch.

### Test flow execution

- The resource table and all it's children tables are deleted from PostgreSQL to work with a clean state.
- CloudQuery SDK fetches the resource, transforms and loads it into PostgreSQL.
- For each resource table (and it's children) CloudQuery reads all the rows and checks that
  - Each table has at least one row available unless `IgnoreInTests` is defined in the `Table`.
  - Each column has data which means the default or custom column resolver worked (unless skipped with `IgnoreInTests` on the `Column`).

## ResourceTest and ResourceTestCase

Each resource can be tested end-to-end with `ResourceTest` function. Here is an example:

```go
func AWSTestHelper(t *testing.T) {
	cfg := `
	aws_debug = false
	`

	providertest.TestResource(t, providertest.ResourceTestCase{
		Provider: &provider.Provider{
			Name:      "aws_test_provider",
			Version:   "development",
			Configure: Configure,
			Config: func() provider.Config {
				return &Config{}
			},
			ResourceMap: map[string]*schema.Table{
				"test_resource": MyResource(),
			},
		},
		Table:         table,
		Config:        cfg,
	})
}
```

## Useful Resources

- CloudQuery official [cq-provider-aws](https://github.com/cloudquery/cq-provider-aws/blob/main/client/testing.go) as an example.
