# Environment variable substitution

CloudQuery configuration `.yml` files supports substitution of values
from environment variables. This allows to extract security sensitive data (like passwords etc) or variable data (that you want to change without touching CloudQuery configuration) from configuration file and store in the environment variable

## Example

And use it inside `gcp.yml`:

```yaml
kind: destination
spec:
  name: postgresql
  tables: [""]
  spec:
    connection_string: ${PG_CONNECTION_STRING}
```

`PG_CONNECTION_STRING` will be sourced from the environment and replaces with the `${PG_CONNECTION_STRING}` before processing.
