# CloudQuery vs Google Cloud Asset Inventory

Google Cloud Asset Inventory is the native GCP asset inventory.

**Key Differences:**

- **Resource Types**: CloudQuery [supports](https://github.com/cloudquery/cloudquery/tree/main/plugins/source/gcp/docs/tables) more than 80 types of resources while GCP currently supports about 140 types of resources (tables). Being an open-source project you can easily add the missing resources without being blocked by a vendor.
- **Database Agnostic and Raw Access to data**: CloudQuery supports multiple databases such as PostgreSQL, BigQuery and others. This makes it play nicely with the whole SQL eco-system and gives you the ability to re-use other tools like Grafana/BI. AWS Config is using a proprietary subset of SQL and database and thus doesn't give you the ability to re-use other tools easily.
- **Search**: CloudQuery doesn't impose any limits on what you can search or filter on due to vanilla PostgreSQL. GCP allows to query/search only on set of predefined attributes.
- **Cloud Agnostic**: CloudQuery gives you the ability to assess, audit and monitor [multi-cloud and SaaS infrastructure](https://hub.cloudquery.io).
- **Policy Language**: CloudQuery enables to codify and version controls security and compliance rules using HCL configuration with SQL as the query engine. GCP doesn't have such policy language where you can aggregate multiple queries and rules.
