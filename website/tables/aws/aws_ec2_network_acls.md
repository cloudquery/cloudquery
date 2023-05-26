# Table: aws_ec2_network_acls

This table shows data for Amazon Elastic Compute Cloud (EC2) Network ACLs.

https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_NetworkAcl.html

The primary key for this table is **arn**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|utf8|
|_cq_sync_time|timestamp[us, tz=UTC]|
|_cq_id|uuid|
|_cq_parent_id|uuid|
|account_id|utf8|
|region|utf8|
|arn (PK)|utf8|
|tags|json|
|associations|json|
|entries|json|
|is_default|bool|
|network_acl_id|utf8|
|owner_id|utf8|
|vpc_id|utf8|