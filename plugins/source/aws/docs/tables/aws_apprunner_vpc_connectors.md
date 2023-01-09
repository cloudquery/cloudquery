# Table: aws_apprunner_vpc_connectors

https://docs.aws.amazon.com/apprunner/latest/api/API_VpcConnector.html

The primary key for this table is **arn**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|region|String|
|arn (PK)|String|
|tags|JSON|
|created_at|Timestamp|
|deleted_at|Timestamp|
|security_groups|StringArray|
|status|String|
|subnets|StringArray|
|vpc_connector_arn|String|
|vpc_connector_name|String|
|vpc_connector_revision|Int|