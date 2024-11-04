This CloudQuery transformer plugin provides basic transformation capabilities:

- Removing columns
- Adding literal string columns
- Adding a column with the timestamp that the record was processed by the transformer
- Obfuscating string columns
- Renaming tables using a name template (use `{{.OldName}}` to refer to the original name, see example below)

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

:configuration

JSON is supported for removing paths and obfuscating string values. Array indexes are supported in both cases. For example, with a JSON column named `tags`:
```json
{"foo":{"bar":["a","b","c"]},"hello":"world"}
```

You can obfuscate `"a"` and remove `"b"` with:
```yaml copy
kind: transformer
spec:
  name: "basic"
  path: "cloudquery/basic"
  registry: "cloudquery"
  spec:
    transformations:
      - kind: obfuscate_columns
        tables: ["example"]
        columns: ["tags.foo.bar.0"]
      - kind: remove_columns
        tables: ["example"]
        columns: ["tags.hello", "tags.foo.bar.1"]
```

Note: transformations are applied sequentially. If you rename tables, the table matcher configuration of subsequent transformations will need to be updated to the new names.