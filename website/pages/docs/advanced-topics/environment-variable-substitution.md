---
title: Environment and File Variable Substitution
description: CloudQuery configuration .yml files support substitution of values from environment variables. This allows you to keep sensitive data (like passwords & tokens) or variable data (that you want to change without touching CloudQuery configuration) outside the configuration file and load them from environment variables at run-time.
---

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

## Environment variables with multi-line JSON

Multi-line JSON, such as those required by the service account key for the GCP integration, can be imported by using pipe '|' operator. The substitution should be in the next line and it should be indented by a single tab before. You don't need to escape any characters while passing the variable.

Inside `gcp.yml`:

```yaml copy
kind: "source"
spec:
  name: "gcp"
  spec:
    service_account_key_json: |
      ${GCP_SERVICE_ACCOUNT_KEY_JSON}
```

## JSON files in older versions

If the file or environment variable being substituted in contains JSON, it should be imported as-is. If you're using CloudQuery version 3.5.0 or prior, it should be imported inside single quotes and content should be escaped with newlines removed.

```yaml copy
kind: "destination"
spec:
  name: "bigquery"
  spec:
    service_account_key_json: '${file:./path/to/secret/file.json}' # single quotes only for CLI versions 3.5.0 or prior
```
