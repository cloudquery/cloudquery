---
name: S3
stage: GA
title: S3 Destination Plugin
description: CloudQuery S3 destination plugin documentation
---
# S3 Destination Plugin

:badge

This destination plugin lets you sync data from a CloudQuery source to remote S3 storage in various formats such as CSV, JSON and Parquet.

This is useful in various use-cases, especially in data lakes where you can query the data direct from Athena or load it to various data warehouses such as BigQuery, RedShift, Snowflake and others.

## Example

:configuration

The S3 destination utilizes batching, and supports `batch_size`, `batch_size_bytes` and `batch_timeout` options (see below).

## S3 Spec

This is the (nested) spec used by the CSV destination Plugin.

- `bucket` (`string`) (required)

  Bucket where to sync the files.

- `region` (`string`) (required)

  Region where bucket is located.

- `credentials` ([credentials](#credentials)) (optional)

  Optional parameters to enable non-default credentials, to authenticate with the S3 API


- `path` (`string`) (required)

  Path to where the files will be uploaded in the above bucket, for example `path/to/files/{{TABLE}}/{{UUID}}.parquet`.

  The path supports the following placeholder variables:

  - `{{TABLE}}` will be replaced with the table name
  - `{{TABLE_HYPHEN}}` will be replaced with the table name with hyphens instead of underscores.
  - `{{SYNC_ID}}` will be replaced with the unique identifier of the sync. This value is a UUID and is randomly generated for each sync.
  - `{{FORMAT}}` will be replaced with the file format, such as `csv`, `json` or `parquet`. If compression is enabled, the format will be `csv.gz`, `json.gz` etc.
  - `{{UUID}}` will be replaced with a random UUID to uniquely identify each file
  - `{{YEAR}}` will be replaced with the current year in `YYYY` format
  - `{{MONTH}}` will be replaced with the current month in `MM` format
  - `{{DAY}}` will be replaced with the current day in `DD` format
  - `{{HOUR}}` will be replaced with the current hour in `HH` format
  - `{{MINUTE}}` will be replaced with the current minute in `mm` format

  **Note** that timestamps are in `UTC` and will be the current time at the time the file is written, not when the sync started.

- `format` (`string`) (required)

  Format of the output file. Supported values are `csv`, `json` and `parquet`.

- `format_spec` ([format_spec](#format_spec)) (optional)

  Optional parameters to change the format of the file.

- `server_side_encryption_configuration` ([server_side_encryption_configuration](#server_side_encryption_configuration)) (optional)

  Optional parameters to enable server-side encryption.

- `compression` (`string`) (optional) (default: `""`)

  Compression algorithm to use. Supported values are `""` or `gzip`. Not supported for `parquet` format.

- `no_rotate` (`boolean`) (optional) (default: `false`)

  If set to `true`, the plugin will write to one file per table.
  Otherwise, for every batch a new file will be created with a different `.<UUID>` suffix.

- `athena` (`boolean`) (optional) (default: `false`)

  When `athena` is set to `true`, the S3 plugin will sanitize keys in JSON columns to be compatible with the Hive Metastore / Athena.
  This allows tables to be created with a Glue Crawler and then queried via Athena, without changes to the table schema.

- `write_empty_objects_for_empty_tables` (`boolean`) (optional) (default: `false`)

  By default only tables with resources are persisted to objects during the sync. If you'd like to persist empty objects for empty tables enable this option. Useful when using CloudQuery Compliance policies to ensure all tables have their schema populated by a query engine like Athena

- `test_write` (`boolean`) (optional) (default: `true`)

  Ensure write access to the given bucket and path by writing a test object on each sync.
  If you are sure that the bucket and path are writable, you can set this to `false` to skip the test.

- `endpoint` (`string`) (optional) (default: `""`)

  Endpoint to use for S3 API calls. This is useful for S3-compatible storage services such as MinIO.
  **Note**: if you want to use path-style addressing, i.e., `https://s3.amazonaws.com/BUCKET/KEY`, `use_path_style` should be enabled, too.

- `acl` (`string`) (optional) (default: `""`)

  Canned ACL to apply to the object. Supported values are `private`, `public-read`, `public-read-write`, `authenticated-read`, `aws-exec-read`, `bucket-owner-read`, `bucket-owner-full-control`.

- `endpoint_skip_tls_verify` (`boolean`) (optional) (default: `false`)
  
  Disable TLS verification for requests to your S3 endpoint.

  This option is intended to be used when using a custom endpoint using the `endpoint` option.

- `use_path_style` (`boolean`) (optional) (default: `false`)

  Allows to use path-style addressing in the `endpoint` option, i.e., `https://s3.amazonaws.com/BUCKET/KEY`.
  By default, the S3 client will use virtual hosted bucket addressing when possible (`https://BUCKET.s3.amazonaws.com/KEY`).

- `batch_size` (`integer`) (optional) (default: `10000`)

  Number of records to write before starting a new object.

- `batch_size_bytes` (`integer`) (optional) (default: `52428800` (= 50 MiB))

  Number of bytes (as Arrow buffer size) to write before starting a new object.

- `batch_timeout` (`duration`) (optional) (default: `30s` (30 seconds))

  Maximum interval between batch writes.

### format_spec

#### CSV

- `delimiter` (`string`) (optional) (default: `,`)

  Delimiter to use in the CSV file.

- `skip_header` (`boolean`) (optional) (default: `false`)

  If set to `true`, the CSV file will not contain a header row as the first row.

#### JSON

Reserved for future use.

#### Parquet

- `version` (`string`) (optional) (default: `v2Latest`)

  Parquet format version to use. Supported values are `v1.0`, `v2.4`, `v2.6` and `v2Latest`.
  `v2Latest` is an alias for the latest version available in the Parquet library which is currently `v2.6`.

  Useful when the reader consuming the Parquet files does not support the latest version.

- `root_repetition` (`string`) (optional) (default: `repeated`)

  [Repetition option to use for the root node](https://github.com/apache/arrow/issues/20243). Supported values are `undefined`, `required`, `optional` and `repeated`.

  Some Parquet readers require a specific root repetition option to be able to read the file. For example, importing Parquet files into [Snowflake](https://www.snowflake.com/en/) requires the root repetition to be `undefined`.

- `max_row_group_length` (`integer`) (optional) (default: `134217728` (= 128 * 1024 * 1024))

  The maximum number of rows in a single row group. Use a lower number to reduce memory usage when reading the Parquet files, and a higher number to increase the efficiency of reading the Parquet files.

### server_side_encryption_configuration

- `sse_kms_key_id` (`string`) (required)

  KMS Key ID appended to S3 API calls header. Used in conjunction with `server_side_encryption`.

- `server_side_encryption` (`string`) (required)

  The server-side encryption algorithm used when storing the object in S3. Supported values are `AES256`, `aws:kms` and `aws:kms:dsse`.


### credentials

- `local_profile` (`string`) (default: will use current credentials)

  [Local profile](https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-files.html) to use to authenticate this account with.
  Please note this should be set to the name of the profile.

  For example, with the following credentials file:

  ```toml copy
  [default]
  aws_access_key_id=xxxx
  aws_secret_access_key=xxxx

  [user1]
  aws_access_key_id=xxxx
  aws_secret_access_key=xxxx
  ```

  `local_profile` should be set to either `default` or `user1`.

- `role_arn` (`string`)

  If specified will use this to assume role.

- `role_session_name` (`string`)

  If specified will use this session name when assume role to `role_arn`.

- `external_id` (`string`)

  If specified will use this when assuming role to `role_arn`.

#### CSV

- `delimiter` (`string`) (optional) (default: `,`)

  Delimiter to use in the CSV file.

- `skip_header` (`boolean`) (optional) (default: `false`)

  If set to `true`, the CSV file will not contain a header row as the first row.


## Authentication

:authentication
