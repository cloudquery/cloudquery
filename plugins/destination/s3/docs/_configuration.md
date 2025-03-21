This example uses the parquet format, to create parquet files in `s3://bucket_name/path/to/files`, with each table placed in its own directory.

The (top level) spec section is described in the [Destination Spec Reference](/docs/reference/destination-spec).

```yaml copy
kind: destination
spec:
  name: "s3"
  path: "cloudquery/s3"
  registry: "cloudquery"
  version: "VERSION_DESTINATION_S3"
  write_mode: "append"
  # Learn more about the configuration options at https://cql.ink/s3_destination
  spec:
    bucket: "bucket_name"
    region: "region-name" # Example: us-east-1
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
    # athena: false # <- set this to true for Athena compatibility
    # write_empty_objects_for_empty_tables: false # <- set this to true if using with the CloudQuery Compliance policies
    # test_write: true # tests the ability to write to the bucket before processing the data
    # endpoint: "" # Endpoint to use for S3 API calls.
    # endpoint_skip_tls_verify # Disable TLS verification if using an untrusted certificate
    # use_path_style: false
    # batch_size: 10000 # 10K entries
    # batch_size_bytes: 52428800 # 50 MiB
    # batch_timeout: 30s # 30 seconds
    # max_retries: 3 # 3 retries
    # max_backoff: 30 # 30 seconds
    # part_size: 5242880 # 5 MiB
    # aws_debug: true
```

It is also possible to use `{{YEAR}}`, `{{MONTH}}`, `{{DAY}}` and `{{HOUR}}` in the path to create a directory structure based on the current time. For example:

```yaml
path: "path/to/files/{{TABLE}}/dt={{YEAR}}-{{MONTH}}-{{DAY}}/{{UUID}}.parquet"
```

Other supported formats are `json` and `csv`.

Note that the S3 plugin only supports `append` `write_mode`. The (top level) spec section is described in the [Destination Spec Reference](/docs/reference/destination-spec).
