This CloudQuery transformer plugin provides basic JSON flattening capabilities:

- JSON fields on source table will be flattened into separate columns in the destination, while still preserving the original JSON fields.
- Only JSON fields with schema metadata will be flattened. This metadata is present for selected CloudQuery sources.
- Only single-level JSON object fields will be flattened. JSON arrays will not be flattened.
- Multi-level JSON objects will not be flattened recursively.

# Example transformation:

If before applying the transformer your sync is producing this table in the destination:

<!-- vale off -->
| account_id   | region        | details                                                                                  |
|--------------|---------------|------------------------------------------------------------------------------------------|
| 012345678901 | us-west-1     | {"field_1": "2021-06-10 07:30:00", "field_2": "value_2", "field_3": true, "field_4": 4}  |
| 012345678901 | ca-central-1  | {"field_1": "2021-07-25 12:40:00", "field_2": "value_4", "field_3": false, "field_4": 5} |
| 012345678901 | sa-east-1     | {"field_1": "2021-08-15 15:10:00", "field_2": "value_6", "field_3": null, "field_4": 6}  |
<!-- vale on -->

After applying the transformer, you will get this table (`details` column content is omitted for brevity):

<!-- vale off -->
| account_id   | region        | details        | field_1             | field_2 | field_3 | field_4 |
|--------------|---------------|----------------|---------------------|---------|---------|---------|
| 012345678901 | us-west-1     | …(unchanged)   | 2021-06-10 07:30:00 | value_2 | true    | 4       |
| 012345678901 | ca-central-1  | …(unchanged)   | 2021-07-25 12:40:00 | value_4 | false   | 5       |
| 012345678901 | sa-east-1     | …(unchanged)   | 2021-08-15 15:10:00 | value_6 | null    | 6       |
<!-- vale on -->

Extra fields are typed based on the `TypeSchema` metadata:

- `field_1` is typed as `timestamp`
- `field_2` is typed as `string`
- `field_3` is typed as `boolean`
- `field_4` is typed as `int64`

## Configuration

First, add the transformer to your destination. For example, this will add a jsonflattener transformer to a PostgreSQL destination:

```yaml copy
kind: destination
spec:
  name: "postgresql"
  path: "cloudquery/postgresql"
  registry: "cloudquery"
  version: "VERSION_DESTINATION_POSTGRESQL"
  write_mode: "overwrite-delete-stale"
  migrate_mode: forced # optional
  transformers:
    - "jsonflattener"

  spec:
    connection_string: "postgresql://your.user:your.password@localhost:5432/db_name"
```

The `migrate_mode: forced` setting might make sense if you plan on modifying the schema from a previous sync.

Then, add your transformer spec:

:configuration
