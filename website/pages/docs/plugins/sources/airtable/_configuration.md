import { Callout } from 'nextra-theme-docs';

```yaml copy
kind: source
# Common source-plugin configuration
spec:
  name: airtable
  registry: docker
  path: ghcr.io/cloudquery/cq-source-airtable:VERSION_SOURCE_AIRTABLE
  tables: ["*"]
  destinations: ["DESTINATION_NAME"]
  # airtable-specific configuration
  spec:
    access_token: "${AIRTABLE_ACCESS_TOKEN}" # required
    # endpoint_url: "https://api.airtable.com" # Optional, defaults to `https://api.airtable.com`
    # concurrency: 10000 # Optional, defaults to `10000`
```

<Callout type="info">
This example uses [environment variable expansion](/docs/advanced-topics/environment-variable-substitution) to read the token from an `AIRTABLE_ACCESS_TOKEN` environment variable. You can also hardcode the value in the configuration file, but this is not advised for production settings.
</Callout>
