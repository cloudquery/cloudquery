# Table: aws_apprunner_vpc_ingress_connections

This table shows data for AWS App Runner VPC Ingress Connections.

https://docs.aws.amazon.com/apprunner/latest/api/API_VpcIngressConnection.html

Notes:
- 'account_id' has been renamed to 'source_account_id' to avoid conflict with the 'account_id' column that indicates what account this was synced from.

The primary key for this table is **arn**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|arn (PK)|`utf8`|
|source_account_id|`utf8`|
|tags|`json`|
|created_at|`timestamp[us, tz=UTC]`|
|deleted_at|`timestamp[us, tz=UTC]`|
|domain_name|`utf8`|
|ingress_vpc_configuration|`json`|
|service_arn|`utf8`|
|status|`utf8`|
|vpc_ingress_connection_arn|`utf8`|
|vpc_ingress_connection_name|`utf8`|