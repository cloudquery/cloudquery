
# Table: tf_resource_instances
Terraform resource instances
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|resource_id|uuid|Parent resource id|
|instance_id|text|Instance id|
|schema_version|bigint|Terraform schema version|
|attributes|jsonb|Instance attributes|
|dependencies|text[]|Instance dependencies array|
|create_before_destroy|boolean|Should resource should be created before destroying|
