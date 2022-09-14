
# Table: gcp_bigquery_dataset_accesses

## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|dataset_cq_id|uuid||
|dataset_id|text||
|target_types|text[]|Which resources in the dataset this entry applies to.|
|domain|text|A domain to grant access to Any users signed in with the domain specified will be granted the specified access Example: "examplecom" Maps to IAM policy member "domain:DOMAIN"|
|group_by_email|text|An email address of a Google Group to grant access to Maps to IAM policy member "group:GROUP"|
|iam_member|text|Some other type of member that appears in the IAM Policy but isn't a user, group, domain, or special group|
|role|text|An IAM role ID that should be granted to the user, group, or domain specified in this access entry The following legacy mappings will be applied: OWNER  roles/bigquerydataOwner WRITER roles/bigquerydataEditor READER  roles/bigquerydataViewer This field will accept any of the above formats, but will return only the legacy format For example, if you set this field to "roles/bigquerydataOwner", it will be returned back as "OWNER"|
|routine_dataset_id|text|The ID of the dataset containing this routine|
|routine_project_id|text|The ID of the project containing this routine|
|routine_id|text|The ID of the routine The ID must contain only letters (a-z, A-Z), numbers (0-9), or underscores (_) The maximum length is 256 characters|
|special_group|text|A special group to grant access to Possible values include: projectOwners: Owners of the enclosing project projectReaders: Readers of the enclosing project projectWriters: Writers of the enclosing project allAuthenticatedUsers: All authenticated BigQuery users Maps to similarly-named IAM members|
|user_by_email|text|An email address of a user to grant access to For example: fred@examplecom Maps to IAM policy member "user:EMAIL" or "serviceAccount:EMAIL"|
|view_dataset_id|text|The ID of the dataset containing this table|
|view_project_id|text|The ID of the project containing this table|
|view_table_id|text|The ID of the table The ID must contain only letters (a-z, A-Z), numbers (0-9), or underscores (_) The maximum length is 1,024 characters|
