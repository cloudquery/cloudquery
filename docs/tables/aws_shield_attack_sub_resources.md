
# Table: aws_shield_attack_sub_resources
The attack information for the specified SubResource
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|attack_cq_id|uuid|Unique CloudQuery ID of aws_shield_attacks table (FK)|
|attack_vectors|jsonb|The list of attack types and associated counters|
|counters|jsonb|The counters that describe the details of the attack|
|id|text|The unique identifier (ID) of the SubResource|
|type|text|The SubResource type|
