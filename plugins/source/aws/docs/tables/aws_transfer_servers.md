# Table: aws_transfer_servers

This table shows data for Transfer Servers.

https://docs.aws.amazon.com/transfer/latest/userguide/API_DescribedServer.html

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
|certificate|`utf8`|
|domain|`utf8`|
|endpoint_details|`json`|
|endpoint_type|`utf8`|
|host_key_fingerprint|`utf8`|
|identity_provider_details|`json`|
|identity_provider_type|`utf8`|
|logging_role|`utf8`|
|post_authentication_login_banner|`utf8`|
|pre_authentication_login_banner|`utf8`|
|protocol_details|`json`|
|protocols|`list<item: utf8, nullable>`|
|security_policy_name|`utf8`|
|server_id|`utf8`|
|state|`utf8`|
|structured_log_destinations|`list<item: utf8, nullable>`|
|user_count|`int64`|
|workflow_details|`json`|