# CloudQuery Snowflake Destination Plugin

The snowflake plugin helps you sync data to your Snowflake data warehouse.

There are two ways to sync data to Snowflake:

1. Direct (easy but not recommended for production or large data sets): This is the default mode of operation where CQ plugin will stream the results directly to the Snowflake database. There is no additional setup needed apart from authentication to Snowflake.

2. Loading via CSV/JSON from a remote storage: This is the standard way of loading data into Snowflake, it is recommended for production and large data sets. This mode requires a remote storage (e.g. S3, GCS, Azure Blob Storage) and a Snowflake stage to be created. The CQ plugin will stream the results to the remote storage. You can then load those files via a cronjob or via SnowPipe. This method is still in the works and will be updated soon with a guide.

## Snowflake Spec

This is the top level spec used by the Snowflake destination plugin.

- `connection_string` (string) (required)

  Snowflake `connection_string`

  Example `connection_string`:

  ```
  // user[:password]@account/database/schema[?param1=value1&paramN=valueN]
  // or
  // user[:password]@account/database[?param1=value1&paramN=valueN]
  // or
  // user[:password]@host:port/database/schema?account=user_account[?param1=value1&paramN=valueN]
  // or
  // host:port/database/schema?account=user_account[?param1=value1&paramN=valueN]
  ```

  From Snowflake documentation:

  `account` - Name assigned to your Snowflake account. If you are not on us-west-2 or AWS deployment, append the region and platform to the end, e.g., `<account>.<region> or <account>.<region>.<platform>`.


## Underlying library

We use the official [github.com/snowflakedb/gosnowflake](https://github.com/snowflakedb/gosnowflake) package for database connection.
