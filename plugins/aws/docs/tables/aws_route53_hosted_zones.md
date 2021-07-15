
# Table: aws_route53_hosted_zones
A complex type that contains general information about the hosted zone.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text|The AWS Account ID of the resource.|
|tags|jsonb|The tags associated with the hosted zone.|
|arn|text|Amazon Resource Name (ARN) of the route53 hosted zone.|
|delegation_set_id|text|A complex type that lists the Amazon Route 53 name servers for the specified hosted zone.|
|caller_reference|text|The value that you specified for CallerReference when you created the hosted zone.|
|id|text|The ID that Amazon Route 53 assigned to the hosted zone when you created it.|
|name|text|The name of the domain.|
|config_comment|text|Any comments that you want to include about the hosted zone.|
|config_private_zone|boolean|A value that indicates whether this is a private hosted zone.|
|linked_service_description|text|If the health check or hosted zone was created by another service, an optional description that can be provided by the other service.|
|linked_service_principal|text|If the health check or hosted zone was created by another service, the service that created the resource.|
|resource_record_set_count|bigint|The number of resource record sets in the hosted zone.|
