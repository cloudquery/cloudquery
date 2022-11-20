# CloudQuery Snowflake Destination Plugin

This destination plugin let's you sync data from a CloudQuery source to an Snowflake database. This can be useful for local data exploration as no db/service is required.

## Snowflake Spec

This is the top level spec used by the Snowflake destination Plugin.

- `connection_string` (string) (required)

  path to a file. such as `./mydb.sql`