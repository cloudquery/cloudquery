# Environment variable substitution

`cloudquery.yml` supports substitution of values from environment variables. This allows to extract security sensitive data (like passwords etc) or variable data (that you want to change without touching CloudQuery configuration) from configuration file and store in the environment variable. To use the feature, set an environment variable before running CLI, adding a `CQ_VAR` prefix to your desired name:

```bash
export CQ_VAR_AWS_VERSION=latest
export CQ_VAR_ARN=some_value
export CQ_VAR_DSN="postgres://postgres:pass@localhost:5432/postgres?sslmode=disable"
```

And use it inside `cloudquery.yml`:

```yaml
cloudquery:
  providers:
    - name: aws
      version: "${AWS_VERSION}"

  connection:
    dsn: "${DSN}"

providers:
  - name: aws
    accounts:
      - name: "<YOUR ID>"
        role_arn: "${ARN}"
```

Note that only environment variables starting with `CQ_VAR` are available for use in `cloudquery.yml`, and their prefix is removed.
