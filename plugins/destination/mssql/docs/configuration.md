# Microsoft SQL Server destination plugin configuration reference

:badge

## Example Configuration

:configuration

:::callout{type="info"}
Make sure you use [environment variable expansion](/docs/advanced-topics/environment-variable-substitution) in production instead of committing the credentials to the configuration file directly.
:::

The Microsoft SQL Server destination utilizes batching, and supports [`batch_size`](/docs/reference/destination-spec#batch_size) and [`batch_size_bytes`](/docs/reference/destination-spec#batch_size_bytes).

## Microsoft SQL Server spec

This is the (nested) spec used by the Microsoft SQL Server destination plugin.

- `connection_string` (`string`) (required)

  Connection string to connect to the database.
  See [SDK documentation](https://github.com/microsoft/go-mssqldb#connection-parameters-and-dsn) for details.

  Example connection strings:
	- `"sqlserver://username:password@hostname/instance"` basic connection using a named instance
	- `"sqlserver://username:password@localhost?database=master&connection+timeout=30"` select "master" database and set connection timeout (default instance)


- `auth_mode` (`string`) (optional) (default: `ms`)

  If you need to authenticate via Azure Active Directory ensure you specify `azure` value.
  See [SDK documentation](https://github.com/microsoft/go-mssqldb#azure-active-directory-authentication) for more information.
  Supported values:

    - `ms` _connect to Microsoft SQL Server instance_
    - `azure` _connect to Azure SQL Server instance_

- `schema` (`string`) (optional) (default: `dbo`)

  Schema name to be used. 
  By default, Microsoft SQL Server destination plugin will use the [default](https://learn.microsoft.com/en-us/sql/relational-databases/security/authentication-access/ownership-and-user-schema-separation?view=sql-server-ver16#the-dbo-schema) schema named `dbo`.

- `batch_size` (`integer`) (optional) (default: `1000`)

  Maximum amount of items that may be grouped together to be written in a single write.

- `batch_size_bytes` (`integer`) (optional) (default: `5242880` (= 5 MiB))

  Maximum size of items that may be grouped together to be written in a single write.

- `batch_timeout` (`duration`) (optional) (default: `20s` (= 20 seconds))

  Maximum interval between batch writes.

### Verbose logging for debug

The Microsoft SQL Server destination can be run in debug mode.
To achieve this pass the `log` option to `connection_string`.
See [SDK documentation](https://github.com/microsoft/go-mssqldb#connection-parameters-and-dsn) for more details.

Note: This will use [SDK](https://github.com/microsoft/go-mssqldb) built-in logging
and might output data and sensitive information to logs.
Make sure not to use it in production environment.

```yaml copy
kind: destination
spec:
  name:     "mssql"
  path:     "cloudquery/mssql"
  registry:   "cloudquery"
  version:  "VERSION_DESTINATION_MSSQL"

  spec:
    connection_string: "${MSSQL_CONNECTION_STRING};log=255"
```
