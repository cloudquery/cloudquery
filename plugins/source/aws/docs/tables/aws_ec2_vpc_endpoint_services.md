# Table: aws_ec2_vpc_endpoint_services

https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_ServiceDetail.html

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
|acceptance_required|Bool|
|availability_zones|StringArray|
|base_endpoint_dns_names|StringArray|
|manages_vpc_endpoints|Bool|
|owner|String|
|payer_responsibility|String|
|private_dns_name|String|
|private_dns_name_verification_state|String|
|private_dns_names|JSON|
|service_id|String|
|service_name|String|
|service_type|JSON|
|supported_ip_address_types|StringArray|
|tags|JSON|
|vpc_endpoint_policy_supported|Bool|