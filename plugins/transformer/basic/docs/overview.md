This CloudQuery transformer plugin provides basic transformation capabilities:

- Removing columns
- Adding literal string columns
- Adding a column with the timestamp that the record was processed by the transformer
- Obfuscating string columns
- Obfuscating every column except an allowlist (opt-in redaction)
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

## Opt-in redaction with `obfuscate_columns_except`

`obfuscate_columns_except` is the inverse of `obfuscate_columns`: instead of listing the columns to redact, you list the columns (and JSON sub-paths) to **keep**, and everything else is redacted. This is useful for compliance-sensitive syncs where only an explicitly approved set of fields may leave unredacted.

For each matched table, `columns` is the allowlist. Every other field is handled as follows:

- CloudQuery internal columns (those prefixed with `_cq_`) always pass through untouched.
- A bare column name (e.g. `status_phase`) keeps that whole column.
- A dotted entry (e.g. `spec_containers.#.image`) keeps that JSON sub-path and redacts every other leaf inside that JSON column, preserving its structure. Array leaves are hashed individually.
- A non-allowlisted column is obfuscated when it is a string, JSON, or binary column, and **dropped** when its type cannot be redacted into itself (numbers, timestamps, lists, structs, ...).

```yaml copy
kind: transformer
spec:
  name: "basic"
  path: "cloudquery/basic"
  registry: "cloudquery"
  spec:
    transformations:
      - kind: obfuscate_columns_except
        tables: ["k8s_core_pods"]
        columns:
          - status_phase
          - context
          - namespace
          - "spec_containers.#.image"
          - "spec_init_containers.#.image"
          - "status_container_statuses.#.state.terminated.finishedAt"
```

**Path syntax:** allowlist paths use the same gjson syntax as `obfuscate_columns`. If a path segment traverses a JSON **array**, you must write `#` for that segment. For example, to keep the image of every container, the path is `spec_containers.#.image` (not `spec_containers.image`). A path that doesn't match — a missing `#`, a typo in a nested segment — does **not** error; it simply causes that field to be redacted instead of kept, so double-check nested paths against your data.

> Note: `obfuscate_columns_except` only redacts fields within the tables it is pointed at; it does not disable other tables. Restrict which tables are synced with the source plugin's `tables` / `skip_tables` options so that no unapproved table is ever written unredacted.

Two more things to keep in mind:

- **Primary keys.** A primary-key column that is not on the allowlist is still redacted. With the default `include_sha: true` each value hashes to a distinct, stable marker, so upserts keep working. With `include_sha: false` all redacted values become identical, which would collapse rows that share a primary key — so the transformer **refuses to run** if a non-allowlisted primary-key column would be redacted with `include_sha: false`. Either allowlist the primary-key column or keep `include_sha: true`. A primary-key column whose type cannot be redacted in place (e.g. an integer) would be dropped, which also breaks row identity, so that is refused too — allowlist such columns.
- **Structure is preserved.** Redaction replaces leaf *values* only; JSON object keys and array shape remain visible. Non-allowlisted values inside a kept JSON column are redacted, but the surrounding structure (key names) is not.

### Hiding the SHA hash

By default all obfuscation transformations append the SHA-256 hash of the redacted value (`Redacted by CloudQuery | <sha>`) so that distinct values remain distinguishable. Set `include_sha: false` on any obfuscation transformation to omit the hash and emit a bare `Redacted by CloudQuery` marker instead. This applies to `obfuscate_columns`, `obfuscate_sensitive_columns`, and `obfuscate_columns_except`.

```yaml copy
      - kind: obfuscate_columns
        tables: ["k8s_core_pods"]
        include_sha: false
        columns: ["annotations"]
```

Because `include_sha: false` makes every redacted value identical, avoid it on a column that is (or is part of) a primary key — `obfuscate_columns_except` will refuse to run rather than silently collapse rows (see the primary-keys note above).

Note: transformations are applied sequentially. If you rename tables, the table matcher configuration of subsequent transformations will need to be updated to the new names.
Note: escape syntax is [SJSON sytax](https://github.com/tidwall/sjson?tab=readme-ov-file#path-syntax).




Edge Cases and limitations for `drop_rows` transformation:
- Only non-list columns are supported
- To drop rows with `nil` values, configure `value: null`, `value: ~` or drop the `value` configuration altogether
- To drop rows based on a JSON value, use the compacted version of the JSON. For example, if you want to drop rows where a JSON column `tags` has a value of `{"foo": "bar"}`, you should specify the value as `{"foo":"bar"}` without any whitespace.