
# Table: azure_cdn_profile_rule_sets
RuleSet friendly RuleSet name mapping to the any RuleSet or secret related information
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|profile_cq_id|uuid|Unique CloudQuery ID of azure_cdn_profiles table (FK)|
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
