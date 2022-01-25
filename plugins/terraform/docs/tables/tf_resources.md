
# Table: tf_resources
Terraform resources
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|tf_data_cq_id|uuid|Unique CloudQuery ID of tf_data table (FK)|
|running_id|uuid|Unique fetch operation id|
|module|text|Resource module if exists|
|mode|text|Resource mode, for example: data, managed, etc|
|type|text|Resource type|
|name|text|Resource name|
|provider_path|text|Resource provider full path, for example: provider["registry.terraform.io/hashicorp/aws"]|
|provider|text|Resource provider name, for example: aws, gcp, etc|
