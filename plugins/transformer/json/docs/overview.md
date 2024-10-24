This CloudQuery transformer plugin provides basic JSON flattening capabilities:

- JSON fields on source table will be flattened into separate columns in the destination, while still preserving the original JSON fields.
- Only JSON fields with schema metadata will be flattened. This metadata is present for selected CloudQuery sources.
- Only single-level JSON object fields will be flattened. JSON arrays will not be flattened.
- Multi-level JSON objects will not be flattened recursively.

## Configuration

First, add the transformer to your destination. For example, this will add a json transformer to a PostgreSQL destination:

```yaml copy
kind: destination
spec:
  name: "postgresql"
  path: "cloudquery/postgresql"
  registry: "cloudquery"
  version: "v8.0.7"
  write_mode: "overwrite-delete-stale"
  migrate_mode: forced # optional
  transformers:
    - "json"

  spec:
    connection_string: "postgresql://your.user:your.password@localhost:5432/db_name"
```

The `migrate_mode: forced` setting might make sense if you plan on modifying the schema from a previous sync.

Then, add your transformer spec. Here's an example that transforms the XKCD source table:

:configuration
