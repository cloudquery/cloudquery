import { Callout } from 'nextra-theme-docs'

# Column Resolvers

These types of resolvers are called for each row received in the `TableResolver` data fetch.

```go
type ColumnResolver func(ctx context.Context, meta ClientMeta, resource *Resource, c Column) error
```

A ColumnResolver works by extracting data for the given `Column` from the given `Resource`, and setting it in the `Resource` using `resource.Set()`:

```go
func (r *Resource) Set(key string, value interface{}) error {
```

They usually go like:

```go
func resolveDynamodbTableKeySchema(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(*types.TableDescription)

	value := marshalKeySchema(r.KeySchema)

	return resource.Set(c.Name, value)
}
```

This way, the value for the column `c` is extracted from the resource, marshalled, and set.

## Utility Resolvers

Utility Resolvers provide solutions on converting data from one data structure to another, ready to be saved in tables. These are always of type `ColumnResolver`.

Some resolvers convert one type to another (parsing date fields, IP addresses and so on) and some look up data inside inner structs of a resource, or from the parent.

Examples:

```go
// Few examples of look-up helper resolvers:

// PathResolver resolves a field in the Resource.Item
func PathResolver(path string) ColumnResolver
// ParentResourceFieldResolver resolves a field from the parent's resource, the value is expected to be set, if name isn't set the field will be set to null
func ParentResourceFieldResolver(name string) ColumnResolver
// ParentPathResolver resolves a field from the parent
func ParentPathResolver(path string) ColumnResolver
// ParentIdResolver resolves the cq_id from the parent
func ParentIdResolver(_ context.Context, _ ClientMeta, r *Resource, c Column) error

// Few examples of type converting resolvers

// IntResolver tries to cast value into int
func IntResolver(path string) ColumnResolver
// DateResolver resolves the different date formats (ISODate - 2011-10-05T14:48:00.000Z is default) into *time.Time
func DateResolver(path string, rfcs ...string) ColumnResolver
// IPAddressResolver resolves the ip string value and returns net.IP
func IPAddressResolver(path string) ColumnResolver
```

Notice how most examples return an _inline function_ built to resolve according to the given parameters (e.g. `PathResolver` will return a function that resolves for the given `path`) except for the `ParentIdResolver`, which is a static `ColumnResolver`.

<Callout type="info">

Discover more ready-made utility resolvers [in the repository](https://github.com/cloudquery/cq-provider-sdk/blob/main/provider/schema/resolvers.go).

</Callout>
