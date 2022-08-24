
# Table: azure_cdn_profile_rule_set_rules
Rule friendly Rules name mapping to the any Rules or secret related information
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|profile_rule_set_cq_id|uuid|Unique CloudQuery ID of azure_cdn_profile_rule_sets table (FK)|
|order|bigint|The order in which the rules are applied for the endpoint|
|conditions|jsonb|A list of conditions that must be matched for the actions to be executed|
|actions|jsonb|A list of actions that are executed when all the conditions of a rule are satisfied|
|match_processing_behavior|text|If this rule is a match should the rules engine continue running the remaining rules or stop|
|provisioning_state|text|Provisioning status|
|deployment_status|text|Possible values include: 'DeploymentStatusNotStarted', 'DeploymentStatusInProgress', 'DeploymentStatusSucceeded', 'DeploymentStatusFailed'|
|id|text|Resource ID|
|name|text|Resource name|
|type|text|Resource type|
|created_by|text|An identifier for the identity that created the resource|
|created_by_type|text|The type of identity that created the resource|
|created_at_time|timestamp without time zone||
|last_modified_by|text|An identifier for the identity that last modified the resource|
|last_modified_by_type|text|The type of identity that last modified the resource|
|last_modified_at_time|timestamp without time zone||
