
# Table: aws_mq_brokers

## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text|The AWS Account ID of the resource.|
|region|text|The AWS Region of the resource.|
|authentication_strategy|text|The authentication strategy used to secure the broker.|
|auto_minor_version_upgrade|boolean|Enables automatic upgrades to new minor versions for brokers, as Apache releases the versions.|
|broker_arn|text|The Amazon Resource Name (ARN) of the broker.|
|broker_id|text|The unique ID that Amazon MQ generates for the broker.|
|broker_instances|jsonb|A list of information about allocated brokers.|
|broker_name|text|The name of the broker|
|broker_state|text|The status of the broker.|
|created|timestamp without time zone|The time when the broker was created.|
|deployment_mode|text|The deployment mode of the broker.|
|encryption_options_use_aws_owned_key|boolean|Enables the use of an AWS owned CMK using AWS Key Management Service (KMS).|
|encryption_options_kms_key_id|text|The symmetric customer master key (CMK) to use for the AWS Key Management Service (KMS).|
|engine_type|text|The type of broker engine.|
|engine_version|text|The version of the broker engine|
|host_instance_type|text|The broker's instance type.|
|ldap_server_metadata|jsonb|The metadata of the LDAP server used to authenticate and authorize connections to the broker.|
|logs|jsonb|The list of information about logs currently enabled and pending to be deployed for the specified broker.|
|maintenance_window_start_time|jsonb|The parameters that determine the WeeklyStartTime.|
|pending_authentication_strategy|text|The authentication strategy that will be applied when the broker is rebooted.|
|pending_engine_version|text|The version of the broker engine to upgrade to|
|pending_host_instance_type|text|The host instance type of the broker to upgrade to|
|pending_ldap_server_metadata|jsonb|The metadata of the LDAP server that will be used to authenticate and authorize connections to the broker once it is rebooted.|
|pending_security_groups|text[]|The list of pending security groups to authorize connections to brokers.|
|publicly_accessible|boolean|Enables connections from applications outside of the VPC that hosts the broker's subnets.|
|security_groups|text[]|The list of security groups (1 minimum, 5 maximum) that authorizes connections to brokers.|
|storage_type|text|The broker's storage type.|
|subnet_ids|text[]|The list of groups that define which subnets and IP ranges the broker can use from different Availability Zones|
|tags|jsonb|The list of all tags associated with this broker.|
