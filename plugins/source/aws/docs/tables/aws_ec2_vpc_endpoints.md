# Table: aws_ec2_vpc_endpoints

https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_VpcEndpoint.html

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
|creation_timestamp|Timestamp|
|dns_entries|JSON|
|dns_options|JSON|
|groups|JSON|
|ip_address_type|String|
|last_error|JSON|
|network_interface_ids|StringArray|
|owner_id|String|
|policy_document|String|
|private_dns_enabled|Bool|
|requester_managed|Bool|
|route_table_ids|StringArray|
|service_name|String|
|state|String|
|subnet_ids|StringArray|
|tags|JSON|
|vpc_endpoint_id|String|
|vpc_endpoint_type|String|
|vpc_id|String|