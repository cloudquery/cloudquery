# Table: gcp_dns_managed_zones

This table shows data for GCP DNS Managed Zones.

https://cloud.google.com/dns/docs/reference/v1/managedZones#resource

The primary key for this table is **id**.

## Relations

The following tables depend on gcp_dns_managed_zones:
  - [gcp_dns_resource_record_sets](gcp_dns_resource_record_sets)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|project_id|`utf8`|
|cloud_logging_config|`json`|
|creation_time|`utf8`|
|description|`utf8`|
|dns_name|`utf8`|
|dnssec_config|`json`|
|forwarding_config|`json`|
|id (PK)|`int64`|
|kind|`utf8`|
|labels|`json`|
|name|`utf8`|
|name_server_set|`utf8`|
|name_servers|`list<item: utf8, nullable>`|
|peering_config|`json`|
|private_visibility_config|`json`|
|reverse_lookup_config|`json`|
|service_directory_config|`json`|
|visibility|`utf8`|

## Example Queries

These SQL queries are sampled from CloudQuery policies and are compatible with PostgreSQL.

### Ensure legacy networks do not exist for a project (Automated)

```sql
SELECT
  id AS resource_id,
  'Ensure legacy networks do not exist for a project (Automated)' AS title,
  project_id AS project_id,
  CASE
  WHEN dnssec_config->>'state' != 'on' THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  gcp_dns_managed_zones;
```

### Ensure that DNSSEC is enabled for Cloud DNS (Automated)

```sql
SELECT
  DISTINCT
  gdmz.id AS resource_id,
  'Ensure that DNSSEC is enabled for Cloud DNS (Automated)' AS title,
  gdmz.project_id AS project_id,
  CASE
  WHEN gdmzdcdks->>'keyType' = 'keySigning'
  AND gdmzdcdks->>'algorithm' = 'rsasha1'
  THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  gcp_dns_managed_zones AS gdmz,
  jsonb_array_elements(gdmz.dnssec_config->'defaultKeySpecs') AS gdmzdcdks;
```

### Ensure that RSASHA1 is not used for the zone-signing key in Cloud DNS DNSSEC (Manual)

```sql
SELECT
  DISTINCT
  gdmz.id AS resource_id,
  'Ensure that RSASHA1 is not used for the zone-signing key in Cloud DNS DNSSEC (Manual)'
    AS title,
  gdmz.project_id AS project_id,
  CASE
  WHEN gdmzdcdks->>'keyType' = 'zoneSigning'
  AND gdmzdcdks->>'algorithm' = 'rsasha1'
  THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  gcp_dns_managed_zones AS gdmz,
  jsonb_array_elements(gdmz.dnssec_config->'defaultKeySpecs') AS gdmzdcdks;
```

### Ensure that DNSSEC is enabled for Cloud DNS (Automated)

```sql
SELECT
  id AS resource_id,
  'Ensure that DNSSEC is enabled for Cloud DNS (Automated)' AS title,
  project_id AS project_id,
  CASE
  WHEN visibility != 'private'
  AND ((dnssec_config IS NULL) OR dnssec_config->>'state' = 'off')
  THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  gcp_dns_managed_zones;
```


