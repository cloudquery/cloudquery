# Azure Resources View

The Azure Resources view, called `azure_resources`, is a view that shows all the resources that are supported by the Azure plugin. It collects all resources into a single view and allows them to be queried by subscription id, id, name, kind or location.

The view is not created as part of `cloudquery sync`. Instead, it needs to be created with a SQL query after the sync completes. The SQL query to create the `azure_resources` view in PostgreSQL can be found [on GitHub](https://github.com/cloudquery/cloudquery/tree/main/plugins/source/azure/views). 