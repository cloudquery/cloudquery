# Table: aws_ec2_eips

https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_Address.html

The primary key for this table is **_cq_id**.



## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id (PK)|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|region|String|
|allocation_id|String|
|association_id|String|
|carrier_ip|String|
|customer_owned_ip|String|
|customer_owned_ipv4_pool|String|
|domain|String|
|instance_id|String|
|network_border_group|String|
|network_interface_id|String|
|network_interface_owner_id|String|
|private_ip_address|String|
|public_ip|String|
|public_ipv4_pool|String|
|tags|JSON|