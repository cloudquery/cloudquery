
# Table: aws_applicationautoscaling_policies
Information about a scaling policy to use with Application Auto Scaling
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text|The AWS Account ID of the resource.|
|region|text|The AWS Region of the resource.|
|namespace|text|The AWS Service Namespace of the resource.|
|creation_time|timestamp without time zone|The Unix timestamp for when the scaling policy was created. |
|arn|text|The Amazon Resource Name (ARN) of the scaling policy. |
|name|text|The name of the scaling policy. |
|type|text|The scaling policy type. |
|resource_id|text|The identifier of the resource associated with the scaling policy|
|scalable_dimension|text|The scalable dimension|
|service_namespace|text|The namespace of the Amazon Web Services service that provides the resource, or a custom-resource. |
|alarms|jsonb|The CloudWatch alarms associated with the scaling policy.|
|step_scaling_policy_configuration|jsonb|A step scaling policy.|
|target_tracking_scaling_policy_configuration|jsonb|A target tracking scaling policy.|
