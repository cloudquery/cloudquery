
# Table: azure_front_door_rules_engine_rule_match_conditions
A list of match conditions that must meet in order for the actions of the rule to run. Having no match conditions means the actions will always run.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|front_door_rules_engine_rule_cq_id|uuid|Unique CloudQuery ID of azure_front_door_rules_engine_rules table (FK)|
|match_variable|text|Match variable|
|selector|text|Name of selector in request header or request body to be matched|
|operator|text|Describes operator to apply to the match condition|
|negate_condition|boolean|Describes if this is negate condition or not|
|match_value|text[]|Match values to match against|
|transforms|text[]|List of transforms|
