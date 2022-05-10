
# Table: aws_shield_protections
An object that represents a resource that is under DDoS protection.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text|The AWS Account ID of the resource.|
|region|text|The AWS Region of the resource.|
|tags|jsonb|The AWS tags of the resource.|
|application_automatic_response_configuration_status|text|Indicates whether automatic application layer DDoS mitigation is enabled for the protection|
|health_check_ids|text[]|The unique identifier (ID) for the Route 53 health check that's associated with the protection|
|id|text|The unique identifier (ID) of the protection|
|name|text|The name of the protection|
|arn|text|The ARN (Amazon Resource Name) of the protection|
|resource_arn|text|The ARN (Amazon Resource Name) of the Amazon Web Services resource that is protected|
