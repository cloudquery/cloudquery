# CloudQuery Azure Source Plugin

The CloudQuery Azure source plugin extracts information from many of the supported services by Microsoft Azure.

## Authentication

The plugin needs to be authenticated with your Azure account in order to fetch information about your cloud setup.

You can either authenticate with `az login` (when running locally), or by using a "service principal" and exporting environment variables (appropriate for automated deployments).

You can find out more about authentication with Azure at Azure's [documentation](https://github.com/Azure/azure-sdk-for-go) for the GoLang sdk.

### Authentication with `az login`

First, install the [Azure CLI](https://docs.microsoft.com/en-us/cli/azure/install-azure-cli) (`az`). Then, login with the Azure CLI:

```bash
az login
```

You are now authenticated!

### Authentication with Environment Variables

You will need to create a service principal for the plugin to use:

#### Creating a service principal

First, install the Azure CLI (`az`).

Then, login with the Azure CLI:

```bash
az login
```

Then, create the service principal the plugin will use to access your cloud deployment. WARNING: The output of
`az ad sp create-for-rbac` contains credentials that you must protect - Make sure to handle with appropriate care.
This example uses bash - The commands for CMD and PowerShell are similar.

```bash
export SUBSCRIPTION_ID=<YOUR_SUBSCRIPTION_ID>
az account set --subscription $SUBSCRIPTION_ID
az provider register --namespace 'Microsoft.Security'

# Create a service-principal for the plugin
az ad sp create-for-rbac --name cloudquery-sp --scopes /subscriptions/$SUBSCRIPTION_ID --role Reader
```

(you can, of course, choose any name you'd like for your service-principal, `cloudquery-sp` is just an example.
If the service principal doesn't exist it will create a new one, otherwise it will update an existing one)

The output of `az ad sp create-for-rbac` should look like this:

```json
{
  "appId": "YOUR AZURE_CLIENT_ID",
  "displayName": "cloudquery-sp",
  "password": "YOUR AZURE_CLIENT_SECRET",
  "tenant": "YOUR AZURE_TENANT_ID"
}
```

#### Exporting environment variables

Next, you need to export the environment variables that plugin will use to sync your cloud configuration.
Copy them from the output of `az ad sp create-for-rbac` (or, take the opportunity to show off your jq-foo).
The example shows how to export environment variables for Linux - exporting for CMD and PowerShell is similar.

- `AZURE_TENANT_ID` is `tenant` in the JSON.
- `AZURE_CLIENT_ID` is `appId` in the JSON.
- `AZURE_CLIENT_SECRET` is `password` in the JSON.

```bash
export AZURE_TENANT_ID=<YOUR AZURE_TENANT_ID>
export AZURE_CLIENT_ID=<YOUR AZURE_CLIENT_ID>
export AZURE_CLIENT_SECRET=<YOUR AZURE_CLIENT_SECRET>
export AZURE_SUBSCRIPTION_ID=$SUBSCRIPTION_ID
```

## Query Examples

### Find all MySQL servers

```sql
SELECT * FROM azure_mysql_servers;
```

### Find storage accounts that are allowing non-HTTPS traffic

```sql
SELECT * from azure_storage_accounts where enable_https_traffic_only = false;
```

### Find all expired key vaults

```sql
SELECT * from azure_keyvault_vault_keys where attributes_expires >= extract(epoch from now()) * 1000;
```
