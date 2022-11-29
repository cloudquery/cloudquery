# Table: azure_web_functions

https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/services/web/mgmt/2020-12-01/web#FunctionEnvelope

The primary key for this table is **id**.

## Relations
This table depends on [azure_web_apps](azure_web_apps.md).


## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|subscription_id|String|
|web_app_id|String|
|function_app_id|String|
|script_root_path_href|String|
|script_href|String|
|config_href|String|
|test_data_href|String|
|secrets_file_href|String|
|href|String|
|files|JSON|
|test_data|String|
|invoke_url_template|String|
|language|String|
|is_disabled|Bool|
|id (PK)|String|
|name|String|
|kind|String|
|type|String|