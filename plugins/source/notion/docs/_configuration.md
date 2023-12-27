```yaml
kind: source
spec:
  name: "notion"
  path: "cloudquery/notion"
  registry: "cloudquery"
  version: "VERSION_SOURCE_NOTION"
  destinations:
    - "postgresql"
  spec:
    bearer_token: "${NOTION_SECRET_KEY}"
```

:::callout{type="info"}
This example uses [environment variable expansion](/docs/advanced-topics/environment-variable-substitution) to read the token from an `NOTION_SECRET_KEY` environment variable. You can also hardcode the value in the configuration file, but this is not advised for production settings.
:::
