# Table: aws_ec2_flow_logs

This table shows data for Amazon Elastic Compute Cloud (EC2) Flow Logs.

https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_FlowLog.html

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
|creation_time|`timestamp[us, tz=UTC]`|
|deliver_cross_account_role|`utf8`|
|deliver_logs_error_message|`utf8`|
|deliver_logs_permission_arn|`utf8`|
|deliver_logs_status|`utf8`|
|destination_options|`json`|
|flow_log_id|`utf8`|
|flow_log_status|`utf8`|
|log_destination|`utf8`|
|log_destination_type|`utf8`|
|log_format|`utf8`|
|log_group_name|`utf8`|
|max_aggregation_interval|`int64`|
|resource_id|`utf8`|
|traffic_type|`utf8`|

## Example Queries

These SQL queries are sampled from CloudQuery policies and are compatible with PostgreSQL.

### VPC flow logging should be enabled in all VPCs

```sql
SELECT
  'VPC flow logging should be enabled in all VPCs' AS title,
  aws_ec2_vpcs.account_id,
  aws_ec2_vpcs.arn,
  CASE
  WHEN aws_ec2_flow_logs.resource_id IS NULL THEN 'fail'
  ELSE 'pass'
  END
FROM
  aws_ec2_vpcs
  LEFT JOIN aws_ec2_flow_logs ON
      aws_ec2_vpcs.vpc_id = aws_ec2_flow_logs.resource_id;
```


