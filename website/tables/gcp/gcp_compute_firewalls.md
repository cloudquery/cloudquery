# Table: gcp_compute_firewalls

This table shows data for GCP Compute Firewalls.

The primary key for this table is **self_link**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|project_id|`utf8`|
|allowed|`json`|
|creation_timestamp|`utf8`|
|denied|`json`|
|description|`utf8`|
|destination_ranges|`list<item: utf8, nullable>`|
|direction|`utf8`|
|disabled|`bool`|
|id|`int64`|
|kind|`utf8`|
|log_config|`json`|
|name|`utf8`|
|network|`utf8`|
|priority|`int64`|
|self_link (PK)|`utf8`|
|source_ranges|`list<item: utf8, nullable>`|
|source_service_accounts|`list<item: utf8, nullable>`|
|source_tags|`list<item: utf8, nullable>`|
|target_service_accounts|`list<item: utf8, nullable>`|
|target_tags|`list<item: utf8, nullable>`|

## Example Queries

These SQL queries are sampled from CloudQuery policies and are compatible with PostgreSQL.

### GCP CIS3.10 Ensure Firewall Rules for instances behind Identity Aware Proxy (IAP) only allow the traffic from Google Cloud Loadbalancer (GCLB) Health Check and Proxy Addresses (Manual)

```sql
WITH
  combined
    AS (
      SELECT
        *
      FROM
        gcp_compute_firewalls AS gcf, jsonb_array_elements(gcf.allowed) AS a
    )
SELECT
  DISTINCT
  gcf.name AS resource_id,
  'GCP CIS3.10 Ensure Firewall Rules for instances behind Identity Aware Proxy (IAP) only allow the traffic from Google Cloud Loadbalancer (GCLB) Health Check and Proxy Addresses (Manual)'
    AS title,
  gcf.project_id AS project_id,
  CASE
  WHEN NOT (ARRAY['35.191.0.0/16', '130.211.0.0/22'] <@ gcf.source_ranges)
  AND NOT
      (
        gcf.value->>'I_p_protocol' = 'tcp'
        AND ARRAY (SELECT jsonb_array_elements_text(gcf.value->'ports'))
          @> ARRAY['80']
      )
  THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  combined AS gcf;
```


