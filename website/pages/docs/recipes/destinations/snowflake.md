# Snowflake Destination Plugin Recipes

Full spec options for Snowflake destination available [here](https://github.com/cloudquery/cloudquery/tree/main/plugins/destination/snowflake).

Note: Make sure you use environment variable expansion in production instead of committing the credentials to the configuration file directly.

## Basic

```yaml copy
kind: destination
spec:
  name: snowflake
  path: cloudquery/snowflake
  version: "VERSION_DESTINATION_SNOWFLAKE"
  spec:
    connection_string: ${SNOWFLAKE_CONNECTION_STRING}
```

