# Table: aws_ec2_byoip_cidrs

https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_ByoipCidr.html

The composite primary key for this table is (**account_id**, **region**, **cidr**).



## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id (PK)|String|
|region (PK)|String|
|cidr (PK)|String|
|description|String|
|state|String|
|status_message|String|