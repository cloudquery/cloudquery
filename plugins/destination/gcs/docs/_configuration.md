This example configures a GCS destination, to create CSV files in `gcs://bucket_name/path/to/files`.

```yaml copy
kind: destination
spec:
  name: "gcs"
  path: "cloudquery/gcs"
  registry: "cloudquery"
  version: "VERSION_DESTINATION_GCS"
  write_mode: "append"
  spec:
    bucket: "bucket_name"
    path: "path/to/files/{{TABLE}}/{{UUID}}.{{FORMAT}}"
    format: "parquet" # options: parquet, json, csv
    format_spec:
      # CSV specific parameters:
      # delimiter: ","
      # skip_header: false
      # Parquet specific parameters:
      # version: "v2Latest"
      # root_repetition: "repeated"
      # max_row_group_length: 134217728 # 128 * 1024 * 1024

    # Optional parameters
    # compression: "" # options: gzip
    # no_rotate: false
    # batch_size: 10000
    # batch_size_bytes: 52428800 # 50 MiB
    # batch_timeout: 30s
```

Note that the GCS plugin only supports `append` `write_mode`. The (top level) spec section is described in the [Destination Spec Reference](/docs/reference/destination-spec).
