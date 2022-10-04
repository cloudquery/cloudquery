# Table: aws_route53_hosted_zone_resource_record_sets



The primary key for this table is **_cq_id**.

## Relations
This table depends on [`aws_route53_hosted_zones`](aws_route53_hosted_zones.md).

## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|UUID|
|_cq_parent_id|UUID|
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|account_id|String|
|hosted_zone_arn|String|
|name|String|
|type|String|
|alias_target|JSON|
|cidr_routing_config|JSON|
|failover|String|
|geo_location|JSON|
|health_check_id|String|
|multi_value_answer|Bool|
|region|String|
|resource_records|JSON|
|set_identifier|String|
|ttl|Int|
|traffic_policy_instance_id|String|
|weight|Int|