```yaml copy
kind: transformer
spec:
  name: "basic"
  path: "cloudquery/basic"
  version: VERSION_TRANSFORMER_BASIC
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
      - kind: change_table_names
        tables: ["*"]
        new_table_name_template: "cq_sync_{{.OldName}}"
```