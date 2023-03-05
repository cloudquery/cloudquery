# Table: aws_db_proxies

https://docs.aws.amazon.com/AmazonRDS/latest/APIReference/API_DBProxy.html

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
|auth|JSON|
|created_date|Timestamp|
|db_proxy_arn|String|
|db_proxy_name|String|
|debug_logging|Bool|
|endpoint|String|
|engine_family|String|
|idle_client_timeout|Int|
|require_tls|Bool|
|role_arn|String|
|status|String|
|updated_date|Timestamp|
|vpc_id|String|
|vpc_security_group_ids|StringArray|
|vpc_subnet_ids|StringArray|