# Table: gcp_compute_ssl_policies

This table shows data for GCP Compute SSL Policies.

The primary key for this table is **self_link**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|project_id|`utf8`|
|creation_timestamp|`utf8`|
|custom_features|`list<item: utf8, nullable>`|
|description|`utf8`|
|enabled_features|`list<item: utf8, nullable>`|
|fingerprint|`utf8`|
|id|`int64`|
|kind|`utf8`|
|min_tls_version|`utf8`|
|name|`utf8`|
|profile|`utf8`|
|region|`utf8`|
|self_link (PK)|`utf8`|
|warnings|`json`|

## Example Queries

These SQL queries are sampled from CloudQuery policies and are compatible with PostgreSQL.

### Ensure no HTTPS or SSL proxy load balancers permit SSL policies with weak cipher suites (Manual)

```sql
SELECT
  gctsp.id AS resource_id,
  'Ensure no HTTPS or SSL proxy load balancers permit SSL policies with weak cipher suites (Manual)'
    AS title,
  gctsp.project_id AS project_id,
  CASE
  WHEN gctsp.ssl_policy
  NOT LIKE 'https://www.googleapis.com/compute/v1/projects/%/global/sslPolicies/%'
  THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  gcp_compute_target_https_proxies AS gctsp
UNION ALL
  SELECT
    DISTINCT
    gctsp.id AS resource_id,
    'Ensure no HTTPS or SSL proxy load balancers permit SSL policies with weak cipher suites (Manual)'
      AS title,
    gctsp.project_id AS project_id,
    CASE
    WHEN gctsp.ssl_policy
    LIKE 'https://www.googleapis.com/compute/v1/projects/%/global/sslPolicies/%'
    AND (p.min_tls_version != 'TLS_1_2' OR p.min_tls_version != 'TLS_1_3')
    AND (
        (p.profile = 'MODERN' OR p.profile = 'RESTRICTED')
        OR (
            p.profile = 'CUSTOM'
            AND ARRAY[
                'TLS_RSA_WITH_AES_128_GCM_SHA256',
                'TLS_RSA_WITH_AES_256_GCM_SHA384',
                'TLS_RSA_WITH_AES_128_CBC_SHA',
                'TLS_RSA_WITH_AES_256_CBC_SHA',
                'TLS_RSA_WITH_3DES_EDE_CBC_SHA'
              ]
              @> p.enabled_features
          )
      )
    THEN 'fail'
    ELSE 'pass'
    END
      AS status
  FROM
    gcp_compute_target_https_proxies AS gctsp
    JOIN gcp_compute_ssl_policies AS p ON gctsp.ssl_policy = p.self_link;
```


