
# Table: aws_ec2_vpc_endpoint_dns_entries
Describes a DNS entry.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|vpc_endpoint_cq_id|uuid|Unique CloudQuery ID of aws_ec2_vpc_endpoints table (FK)|
|dns_name|text|The DNS name.|
|hosted_zone_id|text|The ID of the private hosted zone.|
