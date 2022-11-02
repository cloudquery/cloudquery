# CloudQuery SQLite Destination Plugin

This destination plugin let's you sync data from a CloudQuery source to a sqlite. This can be very useful for local data exploration as no db/service is required.

## PostgreSQL Spec

This is the top level spec used by the SQLite destination Plugin.

- `connection_string` (string) (required)

  path to a file. such as `./mydb.sql`