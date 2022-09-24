# Environment variable substitution

CloudQuery configuration `.yml` files support substitution of values
from environment variables. This allows you to keep sensitive data (like passwords & tokens) or variable data (that you want to change without touching CloudQuery configuration) outside the configuration file and load them from environment variables at run-time.

## Example

And use it inside `postgresql.yml`:

```yaml
kind: "destination"
spec:
  name: "postgresql"
  spec:
    connection_string: ${PG_CONNECTION_STRING}
```

`PG_CONNECTION_STRING` will be sourced from the environment and replaced with the value of `${PG_CONNECTION_STRING}` before processing.
