This CloudQuery transformer plugin provides basic transformation capabilities:

- Removing columns
- Adding literal string columns
- Adding a column with the timestamp that the record was processed by the transformer
- Obfuscating string columns
- Renaming tables using a name template (use `{{.OldName}}` to refer to the original name, see example below)
- Normalizing column values to all-upper/lowercase
- Dropping rows based on column values

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

JSON is supported for removing paths and obfuscating string values, as well as lower/uppercasing field values. Array indexes are supported in both cases. For example, with a JSON column named `tags`:
```json
{"foo":{"bar":["a","b","c"]},"hello":"world","kubectl.kubernetes.io/last-applied-configuration":"secrets"}
```

You can obfuscate `"a"` and remove `"b"`, `"world"`, and `"secrets"` with:
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
        columns: ["tags.hello", "tags.foo.bar.1", "tags.kubectl\\.kubernetes\\.io\\/last-applied-configuration"]
```
To obfuscate nested JSON arrays like: column `example_column` with value: `{"top_foo":[{"foo": "baz0"},{"foo": "baz1"},{"foo": "baz2"}]}` you can use the following syntax:

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
        columns: ["example_column.top_foo.#.foo"]
```

Note: Obfuscating JSON arrays using `#.foo` syntax will cause all `foo` values to be replaced with the same obfuscated value `{"top_foo":[{"foo": "Redacted by CloudQuery | XXX"},{"foo": "Redacted by CloudQuery | XXX"},{"foo": "Redacted by CloudQuery | XXX"}]}`.

You can also use the `obfuscate_sensitive_columns` transformation to automatically obfuscate all columns marked by the source plugin as `sensitive` and possibly containing secret information.

Note: transformations are applied sequentially. If you rename tables, the table matcher configuration of subsequent transformations will need to be updated to the new names.
Note: escape syntax is [SJSON sytax](https://github.com/tidwall/sjson?tab=readme-ov-file#path-syntax).




Edge Cases and limitations for `drop_rows` transformation:
- Only non-list columns are supported
- To drop rows with `nil` values, configure `value: null`, `value: ~` or drop the `value` configuration altogether
- To drop rows based on a JSON value, use the compacted version of the JSON. For example, if you want to drop rows where a JSON column `tags` has a value of `{"foo": "bar"}`, you should specify the value as `{"foo":"bar"}` without any whitespace.