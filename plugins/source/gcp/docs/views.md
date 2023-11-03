# GCP Resources View

The GCP Resources view, called `gcp_resources`, is a view that shows all the resources that are supported by the GCP plugin. It collects all resources into a single view and allows them to be queried by project id, region, id, name or description.

The view is not created as part of `cloudquery sync`. Instead, it needs to be created with a SQL query after the sync completes. The SQL query to create the `gcp_resources` view in PostgreSQL can be found [on GitHub](https://github.com/cloudquery/cloudquery/tree/main/plugins/source/gcp/views). 