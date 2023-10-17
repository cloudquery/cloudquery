```yaml copy
kind: source
# Common source-plugin configuration
spec:
  name: azuredevops
  path: cloudquery/azuredevops
  version: "VERSION_SOURCE_AZUREDEVOPS"
  tables: ["*"]
  destinations: ["DESTINATION_NAME"]

  # Azure DevOps specific configuration
  spec:
    personal_access_token: "${AZURE_DEVOPS_PERSONAL_ACCESS_TOKEN}"
    organization_url: "${AZURE_DEVOPS_ORGANIZATION_URL}"
    # Optional parameters
    # concurrency: 10000
```

This example uses the `AZURE_DEVOPS_PERSONAL_ACCESS_TOKEN` and `AZURE_DEVOPS_ORGANIZATION_URL` environment variables. You can also hardcode the values in the configuration file, but this is not advised for production settings.
