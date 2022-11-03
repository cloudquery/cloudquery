# Table: aws_apprunner_vpc_ingress_connection

https://docs.aws.amazon.com/apprunner/latest/api/API_VpcIngressConnection.html

The primary key for this table is **arn**.


## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|region|String|
|arn (PK)|String|
|account_id|String|
|created_at|Timestamp|
|deleted_at|Timestamp|
|domain_name|String|
|ingress_vpc_configuration|JSON|
|service_arn|String|
|status|String|
|vpc_ingress_connection_name|String|