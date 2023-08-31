# Table: aws_lightsail_instances

This table shows data for Lightsail Instances.

https://docs.aws.amazon.com/lightsail/2016-11-28/api-reference/API_Instance.html

The primary key for this table is **arn**.

## Relations

The following tables depend on aws_lightsail_instances:
  - [aws_lightsail_instance_port_states](aws_lightsail_instance_port_states)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|access_details|`json`|
|arn (PK)|`utf8`|
|tags|`json`|
|add_ons|`json`|
|blueprint_id|`utf8`|
|blueprint_name|`utf8`|
|bundle_id|`utf8`|
|created_at|`timestamp[us, tz=UTC]`|
|hardware|`json`|
|ip_address_type|`utf8`|
|ipv6_addresses|`list<item: utf8, nullable>`|
|is_static_ip|`bool`|
|location|`json`|
|metadata_options|`json`|
|name|`utf8`|
|networking|`json`|
|private_ip_address|`utf8`|
|public_ip_address|`utf8`|
|resource_type|`utf8`|
|ssh_key_name|`utf8`|
|state|`json`|
|support_code|`utf8`|
|username|`utf8`|

## Example Queries

These SQL queries are sampled from CloudQuery policies and are compatible with PostgreSQL.

### Lightsail instances should use IMDSv2

```sql
SELECT
  'Lightsail instances should use IMDSv2' AS title,
  account_id,
  arn AS resource_id,
  CASE
  WHEN metadata_options->>'HttpTokens' IS DISTINCT FROM 'required' THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  aws_lightsail_instances;
```


