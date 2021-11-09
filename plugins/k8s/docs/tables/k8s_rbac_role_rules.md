
# Table: k8s_rbac_role_rules
PolicyRule holds information that describes a policy rule, but does not contain information about who the rule applies to or which namespace the rule applies to.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|role_cq_id|uuid|Unique CloudQuery ID of k8s_rbac_roles table (FK)|
|verbs|text[]|Verbs is a list of Verbs that apply to ALL the ResourceKinds and AttributeRestrictions contained in this rule|
|api_groups|text[]|APIGroups is the name of the APIGroup that contains the resources|
|resources|text[]|Resources is a list of resources this rule applies to|
|resource_names|text[]|ResourceNames is an optional white list of names that the rule applies to|
|non_resource_urls|text[]|NonResourceURLs is a set of partial urls that a user should have access to|
