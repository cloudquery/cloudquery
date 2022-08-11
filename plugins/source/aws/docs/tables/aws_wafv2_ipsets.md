
# Table: aws_wafv2_ipsets
Contains one or more IP addresses or blocks of IP addresses specified in Classless Inter-Domain Routing (CIDR) notation
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text|The AWS Account ID of the resource.|
|region|text|The AWS Region of the resource.|
|scope|text|Specifies whether this is for an Amazon CloudFront distribution or for a regional application.|
|arn|text|The Amazon Resource Name (ARN) of the entity.|
|addresses|cidr[]|Contains an array of strings that specify one or more IP addresses or blocks of IP addresses in Classless Inter-Domain Routing (CIDR) notation|
|ip_address_version|text|Specify IPV4 or IPV6.|
|id|text|A unique identifier for the set|
|name|text|The name of the IP set|
|description|text|A description of the IP set that helps with identification.|
|tags|jsonb|Resource tags|
