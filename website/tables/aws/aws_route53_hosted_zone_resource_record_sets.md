# Table: aws_route53_hosted_zone_resource_record_sets

This table shows data for Amazon Route 53 Hosted Zone Resource Record Sets.

https://docs.aws.amazon.com/Route53/latest/APIReference/API_ResourceRecordSet.html

The primary key for this table is **_cq_id**.

## Relations

This table depends on [aws_route53_hosted_zones](aws_route53_hosted_zones).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|hosted_zone_arn|`utf8`|
|name|`utf8`|
|type|`utf8`|
|alias_target|`json`|
|cidr_routing_config|`json`|
|failover|`utf8`|
|geo_location|`json`|
|health_check_id|`utf8`|
|multi_value_answer|`bool`|
|region|`utf8`|
|resource_records|`json`|
|set_identifier|`utf8`|
|ttl|`int64`|
|traffic_policy_instance_id|`utf8`|
|weight|`int64`|