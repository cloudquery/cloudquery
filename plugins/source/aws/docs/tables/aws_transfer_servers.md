# Table: aws_transfer_servers

https://docs.aws.amazon.com/transfer/latest/userguide/API_DescribedServer.html

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
|certificate|String|
|domain|String|
|endpoint_details|JSON|
|endpoint_type|String|
|host_key_fingerprint|String|
|identity_provider_details|JSON|
|identity_provider_type|String|
|logging_role|String|
|post_authentication_login_banner|String|
|pre_authentication_login_banner|String|
|protocol_details|JSON|
|protocols|StringArray|
|security_policy_name|String|
|server_id|String|
|state|String|
|user_count|Int|
|workflow_details|JSON|