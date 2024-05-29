# MySQL destination plugin configuration reference

:badge

## Example Configuration

:configuration

:::callout{type="info"}
Make sure you use environment variable expansion in production instead of committing the credentials to the configuration file directly.
:::

The MySQL destination utilizes batching, and supports [`batch_size`](/docs/reference/destination-spec#batch_size) and [`batch_size_bytes`](/docs/reference/destination-spec#batch_size_bytes).

## MySQL spec

This is the (nested) spec used by the MySQL destination plugin.

- `connection_string` (`string`) (required)

  Connection string to connect to the database. See the [Go driver documentation](https://github.com/go-sql-driver/mysql#dsn-data-source-name) for details.
  
  - `"user:password@tcp(127.0.0.1:3306)/dbname"` connect with TCP
  - `"user:password@127.0.0.1:3306/dbname?charset=utf8mb4\u0026parseTime=True\u0026loc=Local"` connect and set charset, time parsing, and location
  - `"user:password@localhost:3306/dbname?timeout=30s\u0026readTimeout=1s\u0026writeTimeout=1s"` connect and set various timeouts
  - `"user:password@/dbname?loc=UTC\u0026allowNativePasswords=true\u0026tls=preferred"` connect and set location and native password allowance, and prefer TLS

- `batch_size` (`integer`) (optional) (default: `1000`)

  Maximum number of items that may be grouped together to be written in a single write.

- `batch_size_bytes` (`integer`) (optional) (default: `4194304` (= 4 MiB))

  Maximum size of items that may be grouped together to be written in a single write.
