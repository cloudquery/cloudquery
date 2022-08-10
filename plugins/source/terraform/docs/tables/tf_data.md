
# Table: tf_data
Terraform meta data
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|backend_type|text|Terraform backend type|
|backend_name|text|Terraform backend name|
|version|bigint|Terraform backend version|
|terraform_version|text|Terraform version|
|serial|bigint|Incremental number which describe the state version|
|lineage|text|The "lineage" is a unique ID assigned to a state when it is created|
