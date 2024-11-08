This example configures the file destination, to create CSV files in  `./cq_csv_output`. You can also choose `json` or `parquet` as the output format.

```yaml copy
kind: destination
spec:
  name: "file"
  path: "cloudquery/file"
  registry: "cloudquery"
  version: "VERSION_DESTINATION_FILE"
  write_mode: "append"
  # Learn more about the configuration options at https://cql.ink/file_destination
  spec:
    path: "path/to/files/{{TABLE}}/{{UUID}}.{{FORMAT}}"
    format: "parquet" # options: parquet, json, csv
    # Optional parameters
    # format_spec:
      # CSV specific parameters:
      # delimiter: ","
      # skip_header: false
      # Parquet specific parameters:
      # version: "v2Latest"
      # root_repetition: "repeated"
      # max_row_group_length: 134217728 # 128 * 1024 * 1024
    # compression: "" # options: gzip
    # no_rotate: false
    # batch_size: 10000
    # batch_size_bytes: 52428800 # 50 MiB
    # batch_timeout: 30s
```

Note that the file plugin only supports `append` `write_mode`. The (top level) spec section is described in the [Destination Spec Reference](/docs/reference/destination-spec).
