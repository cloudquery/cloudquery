# Table: aws_route53resolver_resolver_query_log_config_associations

This table shows data for Amazon Route 53 Resolver Resolver Query Log Config Associations.

https://docs.aws.amazon.com/Route53/latest/APIReference/API_route53resolver_ResolverQueryLogConfigAssociation.html

The composite primary key for this table is (**account_id**, **region**, **id**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id (PK)|`utf8`|
|region (PK)|`utf8`|
|creation_time|`utf8`|
|error|`utf8`|
|error_message|`utf8`|
|id (PK)|`utf8`|
|resolver_query_log_config_id|`utf8`|
|resource_id|`utf8`|
|status|`utf8`|