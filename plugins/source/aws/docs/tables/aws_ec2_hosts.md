# Table: aws_ec2_hosts

https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_Host.html

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
|allocation_time|Timestamp|
|allows_multiple_instance_types|String|
|auto_placement|String|
|availability_zone|String|
|availability_zone_id|String|
|available_capacity|JSON|
|client_token|String|
|host_id|String|
|host_properties|JSON|
|host_recovery|String|
|host_reservation_id|String|
|instances|JSON|
|member_of_service_linked_resource_group|Bool|
|outpost_arn|String|
|owner_id|String|
|release_time|Timestamp|
|state|String|
|tags|JSON|