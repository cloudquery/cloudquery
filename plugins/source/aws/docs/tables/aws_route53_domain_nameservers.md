
# Table: aws_route53_domain_nameservers
Nameserver includes the following elements.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|domain_cq_id|uuid|Unique CloudQuery ID of aws_route53_domains table (FK)|
|name|text|The fully qualified host name of the name server|
|glue_ips|text[]|Glue IP address of a name server entry|
