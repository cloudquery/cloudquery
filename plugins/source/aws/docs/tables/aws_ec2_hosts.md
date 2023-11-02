# Table: aws_ec2_hosts

This table shows data for Amazon Elastic Compute Cloud (EC2) Hosts.

https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_Host.html

The primary key for this table is **arn**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|arn (PK)|`utf8`|
|tags|`json`|
|allocation_time|`timestamp[us, tz=UTC]`|
|allows_multiple_instance_types|`utf8`|
|asset_id|`utf8`|
|auto_placement|`utf8`|
|availability_zone|`utf8`|
|availability_zone_id|`utf8`|
|available_capacity|`json`|
|client_token|`utf8`|
|host_id|`utf8`|
|host_maintenance|`utf8`|
|host_properties|`json`|
|host_recovery|`utf8`|
|host_reservation_id|`utf8`|
|instances|`json`|
|member_of_service_linked_resource_group|`bool`|
|outpost_arn|`utf8`|
|owner_id|`utf8`|
|release_time|`timestamp[us, tz=UTC]`|
|state|`utf8`|

## Example Queries

These SQL queries are sampled from CloudQuery policies and are compatible with PostgreSQL.

### Unused dedicated host

```sql
SELECT
  'Unused dedicated host' AS title,
  account_id,
  arn AS resource_id,
  'fail' AS status
FROM
  aws_ec2_hosts
WHERE
  COALESCE(jsonb_array_length(instances), 0) = 0;
```


