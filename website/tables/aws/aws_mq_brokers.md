# Table: aws_mq_brokers

This table shows data for Amazon MQ Brokers.

https://docs.aws.amazon.com/amazon-mq/latest/api-reference/brokers.html

The primary key for this table is **arn**.

## Relations

The following tables depend on aws_mq_brokers:
  - [aws_mq_broker_configurations](aws_mq_broker_configurations)
  - [aws_mq_broker_users](aws_mq_broker_users)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|arn (PK)|`utf8`|
|actions_required|`json`|
|authentication_strategy|`utf8`|
|auto_minor_version_upgrade|`bool`|
|broker_arn|`utf8`|
|broker_id|`utf8`|
|broker_instances|`json`|
|broker_name|`utf8`|
|broker_state|`utf8`|
|configurations|`json`|
|created|`timestamp[us, tz=UTC]`|
|data_replication_metadata|`json`|
|data_replication_mode|`utf8`|
|deployment_mode|`utf8`|
|encryption_options|`json`|
|engine_type|`utf8`|
|engine_version|`utf8`|
|host_instance_type|`utf8`|
|ldap_server_metadata|`json`|
|logs|`json`|
|maintenance_window_start_time|`json`|
|pending_authentication_strategy|`utf8`|
|pending_data_replication_metadata|`json`|
|pending_data_replication_mode|`utf8`|
|pending_engine_version|`utf8`|
|pending_host_instance_type|`utf8`|
|pending_ldap_server_metadata|`json`|
|pending_security_groups|`list<item: utf8, nullable>`|
|publicly_accessible|`bool`|
|security_groups|`list<item: utf8, nullable>`|
|storage_type|`utf8`|
|subnet_ids|`list<item: utf8, nullable>`|
|tags|`json`|
|users|`json`|
|result_metadata|`json`|