
# Table: aws_route53_traffic_policy_versions
A complex type that contains settings for a traffic policy.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|traffic_policy_cq_id|uuid|Unique CloudQuery ID of aws_route53_traffic_policies table (FK)|
|document|jsonb|The definition of a traffic policy in JSON format.|
|id|text|The ID that Amazon Route 53 assigned to a traffic policy when you created it.|
|name|text|The name that you specified when you created the traffic policy.|
|type|text|The DNS type of the resource record sets that Amazon Route 53 creates when you use a traffic policy to create a traffic policy instance.|
|version|integer|The version number that Amazon Route 53 assigns to a traffic policy.|
|comment|text|The comment that you specify in the CreateTrafficPolicy request, if any.|
