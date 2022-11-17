# Table: azure_appservice_functions

https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice/v2#FunctionEnvelope

The primary key for this table is **id**.

## Relations
This table depends on [azure_appservice_sites](azure_appservice_sites.md).


## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|subscription_id|String|
|kind|String|
|config_href|String|
|files|JSON|
|function_app_id|String|
|href|String|
|invoke_url_template|String|
|is_disabled|Bool|
|language|String|
|script_href|String|
|script_root_path_href|String|
|secrets_file_href|String|
|test_data|String|
|test_data_href|String|
|id (PK)|String|
|name|String|
|type|String|
|site_id|String|