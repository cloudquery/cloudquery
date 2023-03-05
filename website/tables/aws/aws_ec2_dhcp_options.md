# Table: aws_ec2_dhcp_options

https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DhcpOptions.html

The composite primary key for this table is (**account_id**, **region**, **dhcp_options_id**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id (PK)|String|
|region (PK)|String|
|tags|JSON|
|dhcp_configurations|JSON|
|dhcp_options_id (PK)|String|
|owner_id|String|