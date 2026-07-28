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

For each matched table, `columns` is the allowlist. In every mode, CloudQuery internal columns (those prefixed with `_cq_`) pass through untouched, a bare column name (e.g. `status_phase`) keeps that whole column, and a dotted entry (e.g. `spec_containers.#.image`) keeps that JSON sub-path.

How every *other* field is handled is controlled by `unmatched`:

- `drop` (default) — non-allowlisted fields are removed entirely: non-allowlisted top-level columns are dropped from the schema, and non-allowlisted leaves inside a kept JSON column are removed. Smallest output.
- `collapse` — each non-allowlisted object or array is replaced by a single redaction marker; top-level non-allowlisted columns are obfuscated to one marker. Keeps the shape without a marker per leaf.
- `redact` — every non-allowlisted leaf value is replaced by its own marker, preserving the full JSON structure (array leaves hashed individually). Most verbose; a top-level column whose type cannot be redacted (numbers, timestamps, lists, structs) is dropped.

The default `drop` is recommended for allowlist/compliance syncs — non-approved data is simply absent. On a real Kubernetes pods table, `drop` cut the redacted `spec_containers` column from ~900 KB to ~10 KB and the table from 59 columns to 11 (`collapse` was the middle ground at ~200 KB).

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

- **Primary keys.** A non-allowlisted primary-key column would lose the value that identifies each row and break upserts at the destination, so the transformer **refuses to run** in that case. This happens under `unmatched: drop` (the column is removed), under `redact`/`collapse` with `include_sha: false` (every row collapses to an identical marker), and for any primary-key type that cannot be redacted in place. Allowlist your primary-key column(s) — for the Kubernetes tables that is `uid`. Under `redact`/`collapse` with the default `include_sha: true`, a non-allowlisted primary key is redacted to a distinct, stable hash and upserts still work.
- **Structure & keys.** Under `redact` and `collapse`, JSON object keys and array shape remain visible (only leaf values are redacted). Under `drop`, non-allowlisted keys and columns are removed entirely.
- **Strongly-typed destinations.** A non-allowlisted **top-level** column is only ever redacted in place if it is a string, JSON, or binary column; a boolean, numeric, timestamp, list, or struct column is **dropped** instead, in every mode. The emitted Arrow schema therefore never claims a type that the data no longer has. Inside a **JSON** column the story is different: redacted nested leaves become strings, so if the source published nested type information (`TypeSchema`), that metadata can describe a type the redacted value no longer matches — a nested `timestamp` leaf under `redact`, or a whole nested object/array replaced by one marker under `collapse`. Destinations that consume `TypeSchema` are expected to fall back to storing the value as a string when it does not match (the MongoDB destination does this). Use `unmatched: drop` to avoid the mismatch entirely, since non-allowlisted nested fields are then absent rather than retyped.

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