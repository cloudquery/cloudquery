This CloudQuery transformer plugin provides basic transformation capabilities:

- Removing columns
- Adding literal string columns
- Obfuscating string columns

## Configuration

First, add the transformer to your destination. For example, this will add a basic transformer to a PostgreSQL destination:

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
    - "basic"

  spec:
    connection_string: "postgresql://your.user:your.password@localhost:5432/db_name"
```

The `migrate_mode: forced` setting might make sense if you plan on modifying the schema from a previous sync.

Then, add your transformer spec. Here's an example that transforms the XKCD source table:

```yaml copy
---
kind: transformer
spec:
  name: "basic"
  path: "localhost:7777" # TODO change this when it's published
  registry: "grpc" # TODO change this when it's published
  spec:
    transformations:
      - kind: obfuscate_columns
        tables: ["xkcd_comics"]
        columns: ["safe_title", "title"]
      - kind: remove_columns
        tables: ["xkcd_comics"]
        columns: ["transcript", "news"]
      - kind: add_column
        tables: ["xkcd_comics"]
        name: "source"
        value: "xkcd"
```