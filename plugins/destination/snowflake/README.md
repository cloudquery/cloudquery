# CloudQuery Snowflake Destination Plugin

The snowflake plugin helps you sync data to your Snowflake data warehouse. Note: This plugin is only for running the initial migrations and create the tables and schema. For the actual sync you need to use the file plugin and upload the data in csv/json to a storage bucket (S3/GCP/Azure Blob Storage) and use periodic `copy into` and/or snowpipe.

## Snowflake Spec

This is the top level spec used by the Snowflake destination Plugin.

- `connection_string` (string) (required)

  path to a file. such as `./mydb.sql`

Example DSNs:

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

`account` - Name assigned to your Snowflake account. If you are not on us-west-2 or AWS deployement, append the region and platform to the end, e.g., <account>.<region> or <account>.<region>.<platform>.

## Underlying library

We use the official [github.com/snowflakedb/gosnowflake](https://github.com/snowflakedb/gosnowflake) library to communicate with the database
