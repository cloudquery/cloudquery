```yaml copy
kind: destination
spec:
  name: "azblob"
  path: "cloudquery/azblob"
  registry: "cloudquery"
  version: "VERSION_DESTINATION_AZBLOB"
  spec:
    storage_account: "cqdestinationazblob"
    container: "test"
    path: "path/to/files"

    format: "csv" # options: parquet, json, csv
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
