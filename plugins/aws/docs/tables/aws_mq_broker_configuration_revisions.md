
# Table: aws_mq_broker_configuration_revisions

## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|broker_configuration_cq_id|uuid|Unique CloudQuery ID of aws_mq_broker_configurations table (FK)|
|configuration_id|text|Required|
|created|timestamp without time zone|Required|
|data|jsonb|Required|
|description|text|The description of the configuration.|
