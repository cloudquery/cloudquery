# MySQL destination plugin configuration reference

:badge

## Example Configuration

<Configuration/>

:::callout{type="info"}
Make sure you use environment variable expansion in production instead of committing the credentials to the configuration file directly.
:::

The MySQL destination utilizes batching, and supports [`batch_size`](/docs/reference/destination-spec#batch_size) and [`batch_size_bytes`](/docs/reference/destination-spec#batch_size_bytes).

## MySQL spec

This is the (nested) spec used by the MySQL destination plugin.

- `connection_string` (`string`) (required)

  Connection string to connect to the database. See the [Go driver documentation](https://github.com/go-sql-driver/mysql#dsn-data-source-name) for more details.

- `batch_size` (`integer`) (optional) (default: `1000`)

  This parameter controls the maximum amount of items may be grouped together to be written as a single write.

- `batch_size_bytes` (`integer`) (optional) (default: `4194304` (= 4 MiB))

  This parameter controls the maximum size of items that may be grouped together to be written as a single write.
