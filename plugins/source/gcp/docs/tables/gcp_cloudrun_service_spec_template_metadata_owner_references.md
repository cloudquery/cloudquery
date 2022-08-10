
# Table: gcp_cloudrun_service_spec_template_metadata_owner_references
OwnerReference contains enough information to let you identify an owning object
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|service_cq_id|uuid|Unique CloudQuery ID of gcp_cloudrun_services table (FK)|
|api_version|text|API version of the referent|
|block_owner_deletion|boolean|If true, AND if the owner has the "foregroundDeletion" finalizer, then the owner cannot be deleted from the key-value store until this reference is removed|
|controller|boolean|If true, this reference points to the managing controller|
|kind|text|Kind of the referent|
|name|text|Name of the referent|
|uid|text|UID of the referent|
