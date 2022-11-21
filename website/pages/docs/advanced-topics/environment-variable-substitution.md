# Environment and file variable substitution

CloudQuery configuration `.yml` files support substitution of values
from environment variables. This allows you to keep sensitive data (like passwords & tokens) or variable data (that you want to change without touching CloudQuery configuration) outside the configuration file and load them from environment variables at run-time.

## Environment variable substitution example

Inside `postgresql.yml`:

```yaml copy
kind: "destination"
spec:
  name: "postgresql"
  spec:
    connection_string: ${PG_CONNECTION_STRING}
```

`PG_CONNECTION_STRING` will be sourced from the environment and replaced with the value of `${PG_CONNECTION_STRING}` before processing.

## File variable substitution example

Inside `postgresql.yml`:

```yaml copy
kind: "destination"
spec:
  name: "postgresql"
  spec:
    connection_string: ${file:./path/to/secret/file}
```

Local path `./path/to/secret/file` will be read and replaced with the contents of the file before processing.