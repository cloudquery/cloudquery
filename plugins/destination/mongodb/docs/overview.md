---
name: MongoDB
stage: GA
title: MongoDB Destination Plugin
description: CloudQuery MongoDB destination plugin documentation
---
# MongoDB Destination Plugin

:badge

This destination plugin lets you sync data from any CloudQuery source to a MongoDB database.

Supported database versions:

- MongoDB >= 4.4 (The same minimum version supported by the official [Go driver v2](https://github.com/mongodb/mongo-go-driver))

## Configuration

### Example

:configuration

:::callout{type="info"}
Make sure to use [environment variable substitution](/docs/advanced-topics/environment-variable-substitution) in production instead of committing the credentials to the configuration file directly.
:::

The MongoDB destination utilizes batching, and supports [`batch_size`](/docs/reference/destination-spec#batch_size) and [`batch_size_bytes`](/docs/reference/destination-spec#batch_size_bytes). 

### MongoDB Spec

This is the (nested) spec used by the MongoDB destination Plugin.

- `connection_string` (`string`) (required)

  MongoDB URI as described in the official MongoDB [documentation](https://www.mongodb.com/docs/manual/reference/connection-string/).
  
  Example connection strings:
  
  - `"mongodb://username:password@hostname:port/database"` basic connection
  - `"mongodb+srv://username:password@cluster.example.com/database"` connecting to a MongoDB Atlas cluster
  - `"mongodb://localhost:27017/myDatabase?authSource=admin"` specify authentication source

- `database` (`string`) (required)

  Database to sync the data to.

- `batch_size` (`integer`) (optional) (default: `1000`)

  Maximum amount of items that may be grouped together to be written in a single write.

- `batch_size_bytes` (`integer`) (optional) (default: `4194304` (= 4 MiB))

  Maximum size of items that may be grouped together to be written in a single write.

- `aws_credentials` ([aws_credentials](#aws_credentials)) (optional)

  Optional parameters to enable usage of AWS IAM credentials

- `write_retry` ([write_retry](#write_retry)) (optional)

  Tunes the exponential-backoff retry applied around each write batch to absorb transient network errors (e.g. `write tcp ...: broken pipe`) that are not covered by the driver's single built-in retry. Omit to use the defaults.




### write_retry

- `max_attempts` (`integer`) (optional) (default: `5`)

  Maximum number of write attempts per batch, including the initial attempt. Set to `1` to disable retries.

- `initial_backoff` (`duration`) (optional) (default: `"500ms"`)

  Initial backoff between retry attempts. Grows exponentially up to `max_backoff`.

- `max_backoff` (`duration`) (optional) (default: `"10s"`)

  Maximum backoff between retry attempts.

- `max_elapsed` (`duration`) (optional) (default: `"30s"`)

  Maximum total time to spend retrying a single write batch before giving up.



### aws_credentials

- `default` (`bool`) (optional)

  If set to `true` then AWS SDK will use the default credentials based on the AWS Credential chain

- `local_profile` (`string`)

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