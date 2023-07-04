# Table: gcp_dns_policies

This table shows data for GCP DNS Policies.

https://cloud.google.com/dns/docs/reference/v1/policies#resource

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|project_id|`utf8`|
|alternative_name_server_config|`json`|
|description|`utf8`|
|enable_inbound_forwarding|`bool`|
|enable_logging|`bool`|
|id (PK)|`int64`|
|kind|`utf8`|
|name|`utf8`|
|networks|`json`|

## Example Queries

These SQL queries are sampled from CloudQuery policies and are compatible with PostgreSQL.

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


