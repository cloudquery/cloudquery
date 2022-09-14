
# Table: gcp_storage_bucket_acls
Access controls on the bucket.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|bucket_cq_id|uuid|Unique ID of gcp_storage_buckets table (FK)|
|bucket_id|text||
|bucket|text|The name of the bucket|
|domain|text|The domain associated with the entity, if any|
|email|text|The email address associated with the entity, if any|
|entity|text|The entity holding the permission, in one of the following forms: - user-userId - user-email - group-groupId - group-email - domain-domain - project-team-projectId - allUsers - allAuthenticatedUsers Examples: - The user liz@examplecom would be user-liz@examplecom - The group example@googlegroupscom would be group-example@googlegroupscom - To refer to all members of the Google Apps for Business domain examplecom, the entity would be domain-examplecom|
|entity_id|text|The ID for the entity, if any|
|etag|text|HTTP 11 Entity tag for the access-control entry|
|id|text|The ID of the access-control entry|
|kind|text|The kind of item this is For bucket access control entries, this is always storage#bucketAccessControl|
|project_team_project_number|text|The project number|
|project_team|text|The team|
|role|text|The access permission for the entity|
|self_link|text|The link to this access-control entry|
