import { Callout } from 'nextra-theme-docs';

```yaml copy
kind: source
# Common source-plugin configuration
spec:
  name: bitbucket
  registry: docker
  path: ghcr.io/cloudquery/cq-source-bitbucket:VERSION_SOURCE_BITBUCKET
  tables: ["*"]
  destinations: ["DESTINATION_NAME"]
  # bitbucket-specific configuration
  spec:
    username: "${BITBUCKET_USERNAME}" # required
    password: "${BITBUCKET_PASSWORD}" # required
```

<Callout type="info">
This example uses [environment variable expansion](/docs/advanced-topics/environment-variable-substitution) to read the `username` and `password` options from `BITBUCKET_USERNAME` and `BITBUCKET_PASSWORD` environment variables respectively. You can also hardcode the value in the configuration file, but this is not advised for production settings.
</Callout>
