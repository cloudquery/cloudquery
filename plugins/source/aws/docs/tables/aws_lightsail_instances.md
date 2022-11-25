# Table: aws_lightsail_instances

https://docs.aws.amazon.com/lightsail/2016-11-28/api-reference/API_Instance.html

The primary key for this table is **arn**.

## Relations

The following tables depend on aws_lightsail_instances:
  - [aws_lightsail_instance_port_states](aws_lightsail_instance_port_states.md)

## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|region|String|
|access_details|JSON|
|arn (PK)|String|
|add_ons|JSON|
|blueprint_id|String|
|blueprint_name|String|
|bundle_id|String|
|created_at|Timestamp|
|hardware|JSON|
|ip_address_type|String|
|ipv6_addresses|StringArray|
|is_static_ip|Bool|
|location|JSON|
|metadata_options|JSON|
|name|String|
|networking|JSON|
|private_ip_address|String|
|public_ip_address|String|
|resource_type|String|
|ssh_key_name|String|
|state|JSON|
|support_code|String|
|tags|JSON|
|username|String|