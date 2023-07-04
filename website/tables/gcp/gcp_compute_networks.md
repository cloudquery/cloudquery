# Table: gcp_compute_networks

This table shows data for GCP Compute Networks.

The primary key for this table is **self_link**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|project_id|`utf8`|
|ipv4_range|`utf8`|
|auto_create_subnetworks|`bool`|
|creation_timestamp|`utf8`|
|description|`utf8`|
|enable_ula_internal_ipv6|`bool`|
|firewall_policy|`utf8`|
|gateway_ipv4|`utf8`|
|id|`int64`|
|internal_ipv6_range|`utf8`|
|kind|`utf8`|
|mtu|`int64`|
|name|`utf8`|
|network_firewall_policy_enforcement_order|`utf8`|
|peerings|`json`|
|routing_config|`json`|
|self_link (PK)|`utf8`|
|self_link_with_id|`utf8`|
|subnetworks|`list<item: utf8, nullable>`|

## Example Queries

These SQL queries are sampled from CloudQuery policies and are compatible with PostgreSQL.

### Ensure that the default network does not exist in a project (Automated)

```sql
SELECT
  name AS resource_id,
  'Ensure that the default network does not exist in a project (Automated)'
    AS title,
  project_id AS project_id,
  CASE WHEN name = 'default' THEN 'fail' ELSE 'pass' END AS status
FROM
  gcp_compute_networks;
```

### Ensure that VPC Flow Logs is enabled for every subnet in a VPC Network (Automated)

```sql
SELECT
  DISTINCT
  gcn.name AS resource_id,
  'Ensure that VPC Flow Logs is enabled for every subnet in a VPC Network (Automated)'
    AS title,
  gcn.project_id AS project_id,
  CASE WHEN gcs.enable_flow_logs = false THEN 'fail' ELSE 'pass' END AS status
FROM
  gcp_compute_networks AS gcn
  JOIN gcp_compute_subnetworks AS gcs ON gcn.self_link = gcs.network;
```

### Ensure that Cloud DNS logging is enabled for all VPC networks (Automated)

```sql
SELECT
  DISTINCT
  gcn.name AS resource_id,
  'Ensure that Cloud DNS logging is enabled for all VPC networks (Automated)'
    AS title,
  gcn.project_id AS project_id,
  CASE WHEN gdp.enable_logging = false THEN 'fail' ELSE 'pass' END AS status
FROM
  gcp_dns_policies AS gdp,
  jsonb_array_elements(gdp.networks) AS gdpn
  JOIN gcp_compute_networks AS gcn ON
      gcn.self_link
      = replace(gdpn->>'networkUrl', 'compute.googleapis', 'www.googleapis');
```


