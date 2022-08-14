
# Table: aws_route53_hosted_zone_traffic_policy_instances
A complex type that contains settings for the new traffic policy instance.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|hosted_zone_cq_id|uuid|Unique CloudQuery ID of aws_route53_hosted_zones table (FK)|
|id|text|The ID that Amazon Route 53 assigned to the new traffic policy instance.|
|message|text|If State is Failed, an explanation of the reason for the failure.|
|name|text|The DNS name, such as www.|
|state|text|The value of State is one of the following values: Applied Amazon Route 53 has finished creating resource record sets, and changes have propagated to all Route 53 edge locations.|
|ttl|bigint|The TTL that Amazon Route 53 assigned to all of the resource record sets that it created in the specified hosted zone.|
|traffic_policy_id|text|The ID of the traffic policy that Amazon Route 53 used to create resource record sets in the specified hosted zone.|
|traffic_policy_type|text|The DNS type that Amazon Route 53 assigned to all of the resource record sets that it created for this traffic policy instance.|
|traffic_policy_version|integer|The version of the traffic policy that Amazon Route 53 used to create resource record sets in the specified hosted zone.|
|arn|text|Amazon Resource Name (ARN) of the route53 hosted zone traffic policy instance.|
