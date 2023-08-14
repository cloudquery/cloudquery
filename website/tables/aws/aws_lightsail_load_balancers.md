# Table: aws_lightsail_load_balancers

This table shows data for Lightsail Load Balancers.

https://docs.aws.amazon.com/lightsail/2016-11-28/api-reference/API_LoadBalancer.html

The primary key for this table is **arn**.

## Relations

The following tables depend on aws_lightsail_load_balancers:
  - [aws_lightsail_load_balancer_tls_certificates](aws_lightsail_load_balancer_tls_certificates)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|arn (PK)|`utf8`|
|tags|`json`|
|configuration_options|`json`|
|created_at|`timestamp[us, tz=UTC]`|
|dns_name|`utf8`|
|health_check_path|`utf8`|
|https_redirection_enabled|`bool`|
|instance_health_summary|`json`|
|instance_port|`int64`|
|ip_address_type|`utf8`|
|location|`json`|
|name|`utf8`|
|protocol|`utf8`|
|public_ports|`list<item: int64, nullable>`|
|resource_type|`utf8`|
|state|`utf8`|
|support_code|`utf8`|
|tls_certificate_summaries|`json`|
|tls_policy_name|`utf8`|

## Example Queries

These SQL queries are sampled from CloudQuery policies and are compatible with PostgreSQL.

### Unused Lightsail load balancers

```sql
SELECT
  'Unused Lightsail load balancers' AS title,
  account_id,
  arn AS resource_id,
  'fail' AS status
FROM
  aws_lightsail_load_balancers
WHERE
  COALESCE(jsonb_array_length(instance_health_summary), 0) = 0;
```


