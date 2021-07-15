
# Table: aws_mq_broker_configurations

## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|broker_cq_id|uuid|Unique CloudQuery ID of aws_mq_brokers table (FK)|
|account_id|text|The AWS Account ID of the resource.|
|region|text|The AWS Region of the resource.|
|arn|text|The ARN of the configuration.|
|authentication_strategy|text|The authentication strategy associated with the configuration.|
|created|timestamp without time zone|The date and time of the configuration revision.|
|description|text|The description of the configuration.|
|engine_type|text|The type of broker engine.|
|engine_version|text|The version of the broker engine.|
|id|text|The unique ID that Amazon MQ generates for the configuration.|
|latest_revision_created|timestamp without time zone|The date and time of the configuration revision.|
|latest_revision_description|text|The description of the configuration revision.|
|latest_revision|integer|The revision number of the configuration.|
|name|text|The name of the configuration.|
|tags|jsonb|The list of all tags associated with this configuration.|
