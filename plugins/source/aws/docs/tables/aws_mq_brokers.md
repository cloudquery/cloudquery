# Table: aws_mq_brokers

https://docs.aws.amazon.com/amazon-mq/latest/api-reference/brokers.html

The primary key for this table is **arn**.

## Relations

The following tables depend on aws_mq_brokers:
  - [aws_mq_broker_configurations](aws_mq_broker_configurations.md)
  - [aws_mq_broker_users](aws_mq_broker_users.md)

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
|actions_required|JSON|
|authentication_strategy|String|
|auto_minor_version_upgrade|Bool|
|broker_id|String|
|broker_instances|JSON|
|broker_name|String|
|broker_state|String|
|configurations|JSON|
|created|Timestamp|
|deployment_mode|String|
|encryption_options|JSON|
|engine_type|String|
|engine_version|String|
|host_instance_type|String|
|ldap_server_metadata|JSON|
|logs|JSON|
|maintenance_window_start_time|JSON|
|pending_authentication_strategy|String|
|pending_engine_version|String|
|pending_host_instance_type|String|
|pending_ldap_server_metadata|JSON|
|pending_security_groups|StringArray|
|publicly_accessible|Bool|
|security_groups|StringArray|
|storage_type|String|
|subnet_ids|StringArray|
|tags|JSON|
|users|JSON|
|result_metadata|JSON|