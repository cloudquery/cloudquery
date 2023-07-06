# Table: aws_ec2_eips

This table shows data for Amazon Elastic Compute Cloud (EC2) Eips.

https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_Address.html

The primary key for this table is **_cq_id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|tags|`json`|
|allocation_id|`utf8`|
|association_id|`utf8`|
|carrier_ip|`utf8`|
|customer_owned_ip|`utf8`|
|customer_owned_ipv4_pool|`utf8`|
|domain|`utf8`|
|instance_id|`utf8`|
|network_border_group|`utf8`|
|network_interface_id|`utf8`|
|network_interface_owner_id|`utf8`|
|private_ip_address|`utf8`|
|public_ip|`utf8`|
|public_ipv4_pool|`utf8`|

## Example Queries

These SQL queries are sampled from CloudQuery policies and are compatible with PostgreSQL.

### Unused EC2 EIP

```sql
SELECT
  'Unused EC2 EIP' AS title,
  account_id,
  allocation_id AS resource_id,
  'fail' AS status
FROM
  aws_ec2_eips
WHERE
  association_id IS NULL;
```

### Unused EC2 EIPs should be removed

```sql
SELECT
  'Unused EC2 EIPs should be removed' AS title,
  account_id,
  public_ip AS resource_id,
  CASE WHEN instance_id IS NULL THEN 'fail' ELSE 'pass' END AS status
FROM
  aws_ec2_eips;
```


