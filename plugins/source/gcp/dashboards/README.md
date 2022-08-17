# GCP Dashboards

This directory contains pre-built dashboards. Currently those are available only for Grafana, but you can create them in any other BI platform:

Checkout those tutorials:
* [Building Open Source Cloud Asset Inventory with CloudQuery and Grafana](https://www.cloudquery.io/blog/open-source-cloud-asset-inventory-with-cloudquery-and-grafana)
* [Building Open Source Cloud Asset Inventory with CloudQuery and Apache Superset](https://www.cloudquery.io/blog/cloud-asset-inventory-cloudquery-apache-superset)
* [Building Open Source Cloud Asset Inventory with CloudQuery and AWS QuickSight](https://www.cloudquery.io/blog/cloud-asset-inventory-cloudquery-aws-quicksight)
* [Building Open Source Cloud Asset Inventory with MetaBase](https://www.cloudquery.io/blog/cloud-asset-inventory-cloudquery-metabase)

## What's inside?

### GCP Asset Inventory

<img alt="GCP Asset Inventory" src="../dashboards/grafana/asset_inventory.png" width=50% height=50%>

#### Installation

1. Execute [this query](https://github.com/cloudquery/plugins/source/gcp/blob/main/views/resource.sql) to add the `gcp_resources` view.
2. Add the CloudQuery postgres database as a data source to Grafana (`Configuration -> Data Sources -> Add Data Source`)
3. Import [../dashboards/grafana/asset_inventory.json](../dashboards/grafana/asset_inventory.json) into Grafana (`Import -> Upload JSON File`).

### GCP Compliance and CSPM (Cloud Security Posture Management) Dashboard

<img alt="GCP Compliance and CSPM Dashboard" src="../dashboards/grafana/compliance.png" width=50% height=50%>

#### Installation

1. Execute one more of the GCP [policies](../policies/).
2. Add the CloudQuery postgres database as a data source to Grafana (`Configuration -> Data Sources -> Add Data Source`)
3. Import [../dashboards/grafana/compliance.json](../dashboards/grafana/compliance.json) into Grafana (`Import -> Upload JSON File`).
