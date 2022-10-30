# Table: aws_ec2_internet_gateways

https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_InternetGateway.html

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
|attachments|JSON|
|internet_gateway_id|String|
|owner_id|String|