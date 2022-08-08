
# Table: aws_sns_subscriptions
A wrapper type for the attributes of an Amazon SNS subscription.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text|The AWS Account ID of the resource.|
|region|text|The AWS Region of the resource.|
|endpoint|text|The subscription's endpoint (format depends on the protocol).|
|owner|text|The subscription's owner.|
|protocol|text|The subscription's protocol.|
|arn|text|The subscription's ARN.|
|topic_arn|text|The ARN of the subscription's topic.|
