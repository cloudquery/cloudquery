# Snowflake Destination Plugin Recipes

Full spec options for the Snowflake destination plugin are available [here](/docs/plugins/destinations/snowflake/overview#snowflake-spec).

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
