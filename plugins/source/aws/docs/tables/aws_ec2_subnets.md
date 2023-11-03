# Table: aws_ec2_subnets

This table shows data for Amazon Elastic Compute Cloud (EC2) Subnets.

https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_Subnet.html
The 'request_account_id' and 'request_region' columns are added to show from where the request was made.

The composite primary key for this table is (**request_account_id**, **request_region**, **arn**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|request_account_id (PK)|`utf8`|
|request_region (PK)|`utf8`|
|arn (PK)|`utf8`|
|tags|`json`|
|assign_ipv6_address_on_creation|`bool`|
|availability_zone|`utf8`|
|availability_zone_id|`utf8`|
|available_ip_address_count|`int64`|
|cidr_block|`utf8`|
|customer_owned_ipv4_pool|`utf8`|
|default_for_az|`bool`|
|enable_dns64|`bool`|
|enable_lni_at_device_index|`int64`|
|ipv6_cidr_block_association_set|`json`|
|ipv6_native|`bool`|
|map_customer_owned_ip_on_launch|`bool`|
|map_public_ip_on_launch|`bool`|
|outpost_arn|`utf8`|
|owner_id|`utf8`|
|private_dns_name_options_on_launch|`json`|
|state|`utf8`|
|subnet_arn|`utf8`|
|subnet_id|`utf8`|
|vpc_id|`utf8`|

## Example Queries

These SQL queries are sampled from CloudQuery policies and are compatible with PostgreSQL.

### EC2 subnets should not automatically assign public IP addresses

```sql
SELECT
  'EC2 subnets should not automatically assign public IP addresses' AS title,
  owner_id AS account_id,
  arn AS resource_id,
  CASE
  WHEN map_public_ip_on_launch IS true THEN 'fail'
  ELSE 'pass'
  END
FROM
  aws_ec2_subnets;
```

### EMR clusters should not have public IP addresses

```sql
SELECT
  'EMR clusters should not have public IP addresses' AS title,
  aws_emr_clusters.account_id,
  aws_emr_clusters.arn AS resource_id,
  CASE
  WHEN aws_ec2_subnets.map_public_ip_on_launch
  AND aws_emr_clusters.status->>'State' IN ('RUNNING', 'WAITING')
  THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  aws_emr_clusters
  LEFT JOIN aws_ec2_subnets ON
      aws_emr_clusters.ec2_instance_attributes->>'Ec2SubnetId'
      = aws_ec2_subnets.subnet_id;
```


