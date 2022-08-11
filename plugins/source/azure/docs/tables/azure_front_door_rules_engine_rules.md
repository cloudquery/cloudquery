
# Table: azure_front_door_rules_engine_rules
A list of rules that define a particular Rules Engine Configuration.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|front_door_rules_engine_cq_id|uuid|Unique CloudQuery ID of azure_front_door_rules_engines table (FK)|
|name|text|A name to refer to the rule|
|priority|integer|A priority assigned to the rule|
|request_header_actions|jsonb|A list of header actions to apply from the request from AFD to the origin.|
|response_header_actions|jsonb|A list of header actions to apply from the response from AFD to the client.|
|route_configuration_override|jsonb|Override the route configuration|
|match_processing_behavior|text|If the rule is a match should the rules engine continue running the remaining rules or stop|
