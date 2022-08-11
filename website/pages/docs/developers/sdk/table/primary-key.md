import { Callout } from 'nextra-theme-docs'

# Primary Key

Every top-level table should have a primary key, ideally consisting of the `account_id` and `id` (or `arn`) of some kind. If the provider supports multiple service regions (and multiple entities with the same `id` can exist in different regions) `region` column should also be included.

Primary keys are defined in the `schema.Table` definition:

```go
func DemoResource() *schema.Table {
    return &schema.Table{
// ...
		Options:      schema.TableCreationOptions{PrimaryKeys: []string{"account_id", "id"}},
// ...
```

## Default Columns

For each table, these columns are automatically added:

| Column Name     | Description                  |
| --------------- | ---------------------------- |
| `cq_id`         | The identifier for relations |
| `cq_meta`       | Holds CQ internal metadata   |
| `cq_fetch_date` | Timestamp for the fetch      |

<Callout type="info">

If no PKs are defined, the `cq_id` column becomes the default Primary Key.

</Callout>

## Relation Tables

The `Relations` key in every `schema.Table` is used to define new tables related to the main table. Every non top-level table (defined in `Relations`) should (probably) have a `<parent>_cq_id` column:

```go
    {
        Name:        "parent_cq_id",
        Description: "Unique CloudQuery ID of aws_ec2_internet_gateways table (FK)",
        Type:        schema.TypeUUID,
        Resolver:    schema.ParentIdResolver,
    },
```

with a suitable `Type`, and the resolver should be set to `schema.ParentIdResolver`.
