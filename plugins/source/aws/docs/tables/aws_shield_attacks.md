
# Table: aws_shield_attacks
The details of a DDoS attack
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text|The AWS Account ID of the resource.|
|attack_counters|jsonb|List of counters that describe the attack for the specified time period|
|id|text|The unique identifier (ID) of the attack|
|end_time|timestamp without time zone|The time the attack ended, in Unix time in seconds|
|mitigations|text[]|List of mitigation actions taken for the attack|
|resource_arn|text|The ARN (Amazon Resource Name) of the resource that was attacked|
|start_time|timestamp without time zone|The time the attack started, in Unix time in seconds|
