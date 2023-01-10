# Table: aws_apprunner_vpc_ingress_connections

https://docs.aws.amazon.com/apprunner/latest/api/API_VpcIngressConnection.html

Notes:
- 'account_id' has been renamed to 'source_account_id' to avoid conflict with the 'account_id' column that indicates what account this was synced from.

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
|source_account_id|String|
|tags|JSON|
|created_at|Timestamp|
|deleted_at|Timestamp|
|domain_name|String|
|ingress_vpc_configuration|JSON|
|service_arn|String|
|status|String|
|vpc_ingress_connection_arn|String|
|vpc_ingress_connection_name|String|