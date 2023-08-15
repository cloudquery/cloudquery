# Azure Source Plugin Contribution Guide

## Adding a new resource

The best way to add a new resource is to use an existing one as reference.

Mostly likely you'll need to create a new client via the `arm*.New<Resource>Client` and use the `NewListPager` to get the data.
See [the `azure_mysql_servers` resource](./resources/services/mysql/servers.go) for an example.

## Testing a new resource

The mock tests use an HTTP server to mock the Azure API. See [the `azure_mysql_servers` resource tests](./resources/services/mysql/servers_mock_test.go) for an example.

## Mapping a resource to a resource provider

If the resource uses `SubscriptionMultiplexRegisteredNamespace`, please ensure you're using the correct resource provider for the resource.
The mapping can be found [here](https://learn.microsoft.com/en-us/azure/azure-resource-manager/management/azure-services-resource-providers).
See [the `azure_mysql_servers` resource](./resources/services/mysql/servers.go#L18) for an example.

You can list all resource providers via the Azure CLI:
```bash
az provider list --query "[].{Provider:namespace, Status:registrationState}" --out table
```

And also enable a resource provider as explained in the [Azure docs](https://learn.microsoft.com/en-us/azure/azure-resource-manager/management/resource-providers-and-types#register-resource-provider).

## General Tips

- Keep transformations to a minimum. As far as possible, we aim to deliver an accurate reflection of what the Azure API provides.
- We generally only unroll structs one level deep. Nested structs should be transformed into JSON columns.
- It's recommended to split each resource addition into a separate PR. This makes it easier to review and merge.
- Before submitting a pull request, run `make gen` to generate documentation for the table. Include these generated files in the pull request.
- If you get stuck or need help, feel free to reach out on [Discord](https://www.cloudquery.io/discord). We are a friendly community and would love to help!
