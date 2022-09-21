
# Table: azure_cdn_profile_endpoint_delivery_policy_rules
DeliveryRule a rule that specifies a set of actions and conditions
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|profile_endpoint_cq_id|uuid|Unique CloudQuery ID of azure_cdn_profile_endpoints table (FK)|
|name|text|Name of the rule|
|order|bigint|The order in which the rules are applied for the endpoint|
|conditions|jsonb|A list of conditions that must be matched for the actions to be executed|
|actions|jsonb|A list of actions that are executed when all the conditions of a rule are satisfied|
