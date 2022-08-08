
# Table: aws_route53_traffic_policies
A complex type that contains information about the latest version of one traffic policy that is associated with the current AWS account.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text|The AWS Account ID of the resource.|
|id|text|The ID that Amazon Route 53 assigned to the traffic policy when you created it.|
|latest_version|integer|The version number of the latest version of the traffic policy.|
|name|text|The name that you specified for the traffic policy when you created it.|
|traffic_policy_count|integer|The number of traffic policies that are associated with the current AWS account.|
|type|text|The DNS type of the resource record sets that Amazon Route 53 creates when you use a traffic policy to create a traffic policy instance.|
|arn|text|The Amazon Resource Name (ARN) for the resource.|
