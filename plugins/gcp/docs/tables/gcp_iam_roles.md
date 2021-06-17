
# Table: gcp_iam_roles
A role in the Identity and Access Management API
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|project_id|text|GCP Project Id of the resource|
|deleted|boolean|The current deleted state of the role This field is read only It will be ignored in calls to CreateRole and UpdateRole|
|description|text|A human-readable description for the role|
|etag|text|Used to perform a consistent read-modify-write|
|included_permissions|text[]|The names of the permissions this role grants when bound in an IAM policy|
|name|text|The name of the role When Role is used in CreateRole, the role name must not be set When Role is used in output and other input such as UpdateRole, the role name is the complete path, eg, roles/loggingviewer for predefined roles and organizations/{ORGANIZATION_ID}/roles/loggingviewer for custom roles|
|stage|text|The current launch stage of the role If the `ALPHA` launch stage has been selected for a role, the `stage` field will not be included in the returned definition for the role  Possible values:   "ALPHA" - The user has indicated this role is currently in an Alpha phase If this launch stage is selected, the `stage` field will not be included when requesting the definition for a given role   "BETA" - The user has indicated this role is currently in a Beta phase   "GA" - The user has indicated this role is generally available   "DEPRECATED" - The user has indicated this role is being deprecated   "DISABLED" - This role is disabled and will not contribute permissions to any members it is granted to in policies   "EAP" - The user has indicated this role is currently in an EAP phase|
|title|text|A human-readable title for the role Typically this is limited to 100 UTF-8 bytes|
