---
name: Snowflake
title: Snowflake Destination Plugin
description: CloudQuery Snowflake destination plugin documentation
---
# Snowflake Destination Plugin

:badge

The snowflake plugin helps you sync data to your Snowflake data warehouse. Authenticating to Snowflake is the only setup required; the plugin creates and migrates tables itself, and loads data through a Snowflake stage that it manages.

## How Data is Loaded

The plugin does not issue row-by-row `INSERT` statements. All data is bulk-loaded via a stage. For each batch of records, the plugin:

1. Writes the batch as newline-delimited JSON to a file in the local temp directory.
2. Uploads that file to an internal Snowflake named stage, `cq_plugin_stage`, using `PUT ... auto_compress=true`.
3. Loads the staged file into the target table:
   - Tables **without** a primary key are loaded with `COPY INTO`, which appends the rows.
   - Tables **with** a primary key are loaded with `MERGE INTO`, so matching rows are updated in place instead of duplicated.

The stage and its file format (`cq_plugin_json_format`) are created with `CREATE OR REPLACE` the first time a sync writes data. Any files still present in `cq_plugin_stage` are therefore discarded when the next sync begins — the stage is scratch space, not durable storage, and should not be read by anything other than the plugin.

## Example Config

:configuration

The Snowflake destination utilizes batching, and supports [`batch_size`](https://www.cloudquery.io/docs/cli/integrations/destinations#batch_size) and [`batch_size_bytes`](https://www.cloudquery.io/docs/cli/integrations/destinations#batch_size_bytes). Each batch becomes one staged file, so larger batches produce fewer and larger files.

## Continuous Loading with Snowpipe

The Snowflake destination does not create pipes and does not write to external stages. Because `cq_plugin_stage` is recreated on every sync (see [How Data is Loaded](#how-data-is-loaded)), you should not point a Snowpipe at it — files can be removed before the pipe has ingested them.

To load CloudQuery data with Snowpipe, write to object storage and load from an external stage instead of using this destination:

1. Sync to object storage with the [S3](https://www.cloudquery.io/hub/plugins/destination/cloudquery/s3/latest/docs), [GCS](https://www.cloudquery.io/hub/plugins/destination/cloudquery/gcs/latest/docs) or [Azure Blob Storage](https://www.cloudquery.io/hub/plugins/destination/cloudquery/azblob/latest/docs) destination. Give each table its own prefix so a pipe can target it:

   ```yaml copy
   kind: destination
   spec:
     name: "s3"
     path: "cloudquery/s3"
     registry: "cloudquery"
     version: "VERSION_DESTINATION_S3"
     write_mode: "append"
     spec:
       bucket: "bucket_name"
       region: "us-east-1"
       path: "cloudquery/{{TABLE}}/{{UUID}}.{{FORMAT}}"
       format: "parquet"
       # Snowflake recommends files of roughly 100-250 MB compressed
       batch_size_bytes: 209715200 # 200 MiB
   ```

2. Create the destination tables in Snowflake yourself. Snowpipe loads into existing tables only; it will not create or migrate schemas the way this destination does.

3. Create a storage integration and an external stage over that prefix:

   ```sql
   create file format cq_parquet_format type = parquet;

   create storage integration cq_s3_integration
     type = external_stage
     storage_provider = 's3'
     storage_aws_role_arn = 'arn:aws:iam::123456789012:role/snowflake-cloudquery'
     enabled = true
     storage_allowed_locations = ('s3://bucket_name/cloudquery/');

   create stage cq_external_stage
     url = 's3://bucket_name/cloudquery/'
     storage_integration = cq_s3_integration
     file_format = cq_parquet_format;
   ```

4. Create one pipe per table and wire up event notifications. Snowflake's [Automating Snowpipe for Amazon S3](https://docs.snowflake.com/en/user-guide/data-load-snowpipe-auto-s3) guide covers getting the notification channel (`show pipes`) and configuring the bucket to publish to it:

   ```sql
   create pipe cq_aws_ec2_instances_pipe
     auto_ingest = true
   as
     copy into aws_ec2_instances
     from @cq_external_stage/aws_ec2_instances/
     match_by_column_name = case_insensitive;
   ```

Things to be aware of with this approach:

- **Loads are append-only.** Snowpipe has no equivalent of the `MERGE INTO` this destination performs for tables with primary keys, and the object storage destinations write with `append` `write_mode`. Every sync adds a new set of rows, so deduplicate downstream — for example with a view that keeps the latest `_cq_sync_time` per `_cq_id`, or a scheduled task.
- **Latency is not guaranteed.** Snowflake [does not commit to a Snowpipe load latency](https://docs.snowflake.com/en/user-guide/data-load-snowpipe-intro) and recommends measuring it against your own workload.
- **File size matters.** Snowflake recommends files of roughly [100-250 MB compressed](https://docs.snowflake.com/en/user-guide/data-load-considerations-prepare); tune the storage destination's `batch_size_bytes` accordingly rather than leaving it at its default.
- **`auto_ingest` on internal stages is restricted.** It is only available for Snowflake accounts hosted on AWS, which is another reason to use an external stage. Alternatives are the Snowpipe REST API or a scheduled `alter pipe ... refresh`.

## Authentication

Authentication of the connection to Snowflake can be specified using:

* A username and password in the DSN:

  ```yaml
  kind: destination
  spec:
    name: snowflake
    send_sync_summary: true
    ...
    spec:
      connection_string: "user:pass@account/db/schema?warehouse=wh"
  ```

* A private key inline:

  ```yaml
  kind: destination
  spec:
    name: snowflake
    send_sync_summary: true
    ...
    spec:
      connection_string: "user@account/database/schema?warehouse=wh"
      private_key: |
        -----BEGIN PRIVATE KEY-----
        MIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQC2ajPRIbPtbxZ1
        3DONLA02eZJuCzsgIkBWov/Me5TL6cKN0gnY+mbA8OnNCH+9HSzgiU9P8XhTUrIN
        85diD+rj6uK+E0sSyxGk6HG17TyR5oBq8nz2hbZlbaNi/HO9qYoHQgAgMq908YBz
        ...
        DUmOIrBYEMf2nDTlTu/QVcKb
        -----END PRIVATE KEY-----
  ```

* A private key included from a file:

  ```yaml
  kind: destination
  spec:
    name: snowflake
    send_sync_summary: true
    ...
    spec:
      connection_string: "user@account/database/schema?warehouse=wh"
      private_key: "${file:./private.key}"
  ```

  where ./private.key is PEM-encoded private key file with contents of the form:

  ```txt
  -----BEGIN PRIVATE KEY-----
  MIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQC2ajPRIbPtbxZ1
  3DONLA02eZJuCzsgIkBWov/Me5TL6cKN0gnY+mbA8OnNCH+9HSzgiU9P8XhTUrIN
  85diD+rj6uK+E0sSyxGk6HG17TyR5oBq8nz2hbZlbaNi/HO9qYoHQgAgMq908YBz
  ...
  DUmOIrBYEMf2nDTlTu/QVcKb
  -----END PRIVATE KEY-----
  ```

* OAuth authentication when running in Snowpark container service

  ```yaml
  kind: destination
  spec:
    name: snowflake
    send_sync_summary: true
    ...
    spec:
      connection_string: "user:pass@account/db/schema?warehouse=wh&authenticator=oauth&token=token"
  ```

### Private Key Authentication Setup

The Snowflake guide for [Key Pair
Authentication](https://docs.snowflake.com/en/user-guide/key-pair-auth)
demonstrates how to create an RSA private key with the ability to authenticate
as a Snowflake user.

Note that encrypted private keys are not supported by the Snowflake Go SQL
driver, and hence not supported by the CloudQuery Snowflake plugin. You can
decrypt a private key in file enc.key and store it in a file dec.key using the
following command:

```bash
openssl pkcs8 -topk8 -nocrypt -in enc.key -out dec.key
```

## Snowflake Spec

This is the top level spec used by the Snowflake destination plugin.

- `connection_string` (`string`) (required)

  Snowflake `connection_string`.

  Example:

  ```yaml copy
  # user[:password]@account/database/schema?warehouse=user_warehouse[&param1=value1&paramN=valueN]
  # or
  # user[:password]@account/database?warehouse=user_warehouse[&param1=value1&paramN=valueN]
  # or
  # user[:password]@host:port/database/schema?account=user_account&warehouse=user_warehouse[&param1=value1&paramN=valueN]
  # or
  # host:port/database/schema?account=user_account&warehouse=user_warehouse[&param1=value1&paramN=valueN]
  ```

  From Snowflake documentation:

  `account` - Name assigned to your Snowflake account. If you are not on us-west-2 or AWS deployment, append the region and platform to the end, e.g., `<account>.<region> or <account>.<region>.<platform>`.

- `private_key` (`string`) (optional)

  A PEM-encoded private key for connecting to snowflake. Equivalent to adding
  `authenticator=snowflake_jwt&privateKey=...` to the `connection_string` but
  parses, validates, and correctly encodes the key for use with snowflake.

- `migrate_concurrency` (`integer`) (optional) (default: `1`)

  By default, tables are migrated one at a time.
  This option allows you to migrate multiple tables concurrently.
  This can be useful if you have a lot of tables to migrate and want to speed up the process.

  Must be `1` or greater.

- `batch_size` (`integer`) (optional) (default: `5000`)

  Number of records to batch together before sending to the database. Each batch is written to a single staged file.

- `batch_size_bytes` (`integer`) (optional) (default: `20971520` (= 20 MiB))

  Number of bytes (as Arrow buffer size) to batch together before sending to the database.

- `leave_stage_files` (`boolean`) (optional) (default: `false`)

  If set to `true`, the intermediary files used to load data into the Snowflake stage are left in the local temp directory instead of being deleted after upload. This can be useful for debugging purposes.

## Underlying library

We use the official [github.com/snowflakedb/gosnowflake](https://github.com/snowflakedb/gosnowflake) package for database connection.
