
# Table: aws_mq_broker_users

## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|broker_cq_id|uuid|Unique CloudQuery ID of aws_mq_brokers table (FK)|
|account_id|text|The AWS Account ID of the resource.|
|region|text|The AWS Region of the resource.|
|console_access|boolean|Enables access to the the ActiveMQ Web Console for the ActiveMQ user.|
|groups|text[]|The list of groups (20 maximum) to which the ActiveMQ user belongs|
|pending|jsonb|The status of the changes pending for the ActiveMQ user.|
|username|text|The username of the ActiveMQ user.|
