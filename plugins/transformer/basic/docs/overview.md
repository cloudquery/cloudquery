This CloudQuery transformer plugin provides basic transformation capabilities:

- Removing columns
- Removing all columns except an allowlist (`remove_columns_except`)
- Adding literal string columns
- Adding a column with the timestamp that the record was processed by the transformer
- Obfuscating string columns
- Obfuscating all obfuscatable columns except an allowlist (`obfuscate_columns_except`)
- Customizing the redacted value via the optional `redaction` block
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

## Allowlist transformations

`obfuscate_columns_except` is the inverse of `obfuscate_columns`: `columns` is a keep-list, and every _other_ obfuscatable column (Arrow `STRING`, `BINARY` or JSON) is obfuscated as a whole column. Columns that are not obfuscatable (integers, booleans, timestamps, etc.) and columns named in the keep-list are left untouched. An empty `columns` list obfuscates every obfuscatable column. Unlike `obfuscate_columns`, this kind does not support JSON-path (`col.a.b`) partial obfuscation — keep-list entries are plain column names, and JSON columns are obfuscated whole.

```yaml copy
- kind: obfuscate_columns_except
  tables: ["example"]
  columns: ["id", "created_at"] # every other obfuscatable column is obfuscated
```

`remove_columns_except` is the inverse of `remove_columns`: `columns` is a keep-list and every _other_ column is removed (columns of any type). CloudQuery system columns (any column whose name starts with `_cq`, such as `_cq_id`, `_cq_source_name`, `_cq_sync_time` and `_cq_parent_id`) are always preserved, whether or not they are listed in `columns`.

```yaml copy
- kind: remove_columns_except
  tables: ["example"]
  columns: ["id", "name"] # keeps id, name and all _cq* columns; removes everything else
```

## Customizing the redacted value

The obfuscate transformations (`obfuscate_columns`, `obfuscate_columns_except` and `obfuscate_sensitive_columns`) accept an optional `redaction` block to control the redacted value. When omitted, the output is unchanged: plaintext columns become `Redacted by CloudQuery | <sha256>` and whole JSON columns become `{"redacted_by_cloudquery":"<sha256>"}`.

The `redaction` block requires both a `plaintext` and a `json` sub-block:

```yaml copy
- kind: obfuscate_columns
  tables: ["example"]
  columns: ["secret"]
  redaction:
    plaintext:
      message: "REDACTED"   # prefix used for STRING/BINARY columns and JSON-path values
      include_hash: true    # when true, append " <sha256>"; when false, emit the message only
    json:
      key: "redacted_by_cloudquery" # object key for whole-JSON columns
      message: "REDACTED"           # value used when include_hash is false
      include_hash: true            # when true, the value is the sha256; when false, the message
```

- `plaintext` applies to `STRING` and `BINARY` columns and to JSON-path (`col.a.b`) values. With `include_hash: true` the value is `<message> <sha256>`; with `include_hash: false` it is `<message>` verbatim.
- `json` applies to whole JSON columns. The value stored under `key` is the `<sha256>` when `include_hash: true`, or `<message>` when `include_hash: false`.
- Setting `plaintext: {message: "Redacted by CloudQuery |", include_hash: true}` reproduces the default plaintext output exactly.

Note: transformations are applied sequentially. If you rename tables, the table matcher configuration of subsequent transformations will need to be updated to the new names.
Note: escape syntax is [SJSON sytax](https://github.com/tidwall/sjson?tab=readme-ov-file#path-syntax).




Edge Cases and limitations for `drop_rows` transformation:
- Only non-list columns are supported
- To drop rows with `nil` values, configure `value: null`, `value: ~` or drop the `value` configuration altogether
- To drop rows based on a JSON value, use the compacted version of the JSON. For example, if you want to drop rows where a JSON column `tags` has a value of `{"foo": "bar"}`, you should specify the value as `{"foo":"bar"}` without any whitespace.