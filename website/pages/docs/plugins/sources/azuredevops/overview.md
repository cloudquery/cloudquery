# Azure DevOps Source Plugin

import { getLatestVersion } from "../../../../../utils/versions";
import { Badge } from "../../../../../components/Badge";

<Badge text={"Latest: " + getLatestVersion("source", `azuredevops`)}/>

The CloudQuery Azure DevOps plugin reads information from your Azure DevOps account and loads it into any supported CloudQuery destination (e.g. PostgreSQL, Snowflake, BigQuery).

## Configuration

This example syncs from Azure DevOps to a Postgres destination. The (top level) source spec section is described in the [Source Spec Reference](/docs/reference/source-spec).

```yaml
kind: source
# Common source-plugin configuration
spec:
  name: azuredevops
  path: cloudquery/azuredevops
  version: "VERSION_SOURCE_AZUREDEVOPS"
  tables: ["*"]
  destinations: ["postgresql"]

  # Azure DevOps specific configuration
  spec:
    personal_access_token: "${AZURE_DEVOPS_PERSONAL_ACCESS_TOKEN}"
    organization_url: "${AZURE_DEVOPS_ORGANIZATION_URL}"
```

For more information on downloading, installing and running the CloudQuery CLI, see the [Quickstart guide](/docs/quickstart).

## Azure DevOps Spec

This is the (nested) spec used by the Azure DevOps source plugin.

- `personal_access_token` (string, required):
   
  An API token to access Azure DevOps resources. This can be obtained by [creating an API token](https://learn.microsoft.com/en-us/azure/devops/organizations/accounts/use-personal-access-tokens-to-authenticate?view=azure-devops&tabs=Windows#create-a-pat). It's recommended to allow only read access to the resources you need to sync.

- `organization_url` ([]string, required):

  The Azure DevOps organization URL. Should be in the format `https://dev.azure.com/{organization}`.