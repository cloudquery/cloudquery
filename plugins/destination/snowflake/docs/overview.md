---
name: Snowflake
title: Snowflake Destination Plugin
description: CloudQuery Snowflake destination plugin documentation
---
# Snowflake Destination Plugin

:badge

The snowflake plugin helps you sync data to your Snowflake data warehouse.

There are two ways to sync data to Snowflake:

1. Direct (easy but not recommended for production or large data sets): This is the default mode of operation where CQ plugin will stream the results directly to the Snowflake database. There is no additional setup needed apart from authentication to Snowflake.

2. Loading via CSV/JSON from a remote storage: This is the standard way of loading data into Snowflake, it is recommended for production and large data sets. This mode requires a remote storage (e.g. S3, GCS, Azure Blob Storage) and a Snowflake stage to be created. The CQ plugin will stream the results to the remote storage. You can then load those files via a cronjob or via SnowPipe. This method is still in the works and will be updated soon with a guide.

## Example Config

:configuration

The Snowflake destination utilizes batching, and supports [`batch_size`](https://www.cloudquery.io/docs/cli/integrations/destinations#batch_size) and [`batch_size_bytes`](https://www.cloudquery.io/docs/cli/integrations/destinations#batch_size_bytes).


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

  Setting this to a negative number means no limit.

- `batch_size` (`integer`) (optional) (default: `1000`)

  Number of records to batch together before sending to the database.

- `batch_size_bytes` (`integer`) (optional) (default: `4194304` (= 4 MiB))

  Number of bytes (as Arrow buffer size) to batch together before sending to the database.

- `leave_stage_files` (boolean) (optional) (default: false)
     
  If set to true, intermediary files used to load data to the Snowflake stage are left in the temp directory. This can be useful for debugging purposes.

- `skip_write_setup` (boolean) (optional) (default: false)

  By default, every sync runs `create or replace file format cq_plugin_json_format` and `create or replace stage cq_plugin_stage` before writing. Replacing the stage removes every file currently in it, so concurrent syncs writing into the same Snowflake schema delete each other's staged files, which fails the subsequent `copy into`/`merge` with a missing file error.

  Set this to `true` to skip those statements when running multiple syncs against the same schema. The `cq_plugin_json_format` file format and the `cq_plugin_stage` stage must already exist in the target schema, otherwise the sync fails when the first file is uploaded. Create them once with:

  ```sql
  create file format if not exists cq_plugin_json_format type = 'JSON';
  create stage if not exists cq_plugin_stage file_format = cq_plugin_json_format;
  ```

## Underlying library

We use the official [github.com/snowflakedb/gosnowflake](https://github.com/snowflakedb/gosnowflake) package for database connection.
