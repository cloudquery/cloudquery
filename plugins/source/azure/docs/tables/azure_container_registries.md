
# Table: azure_container_registries
Azure compute disk
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|subscription_id|text|Azure subscription id|
|sku_name|text|The SKU name of the container registry|
|sku_tier|text|The SKU tier based on the SKU name|
|login_server|text|The URL that can be used to log into the container registry|
|creation_date|timestamp without time zone|The creation date of the container registry|
|provisioning_state|text|The provisioning state of the container registry at the time the operation was called|
|status|text|The short label for the status|
|status_message|text|The detailed message for the status, including alerts and error messages|
|status_timestamp|timestamp without time zone|The timestamp when the status was changed to the current value|
|admin_user_enabled|boolean|The value that indicates whether the admin user is enabled|
|storage_account_id|text|The resource ID of the storage account|
|network_rule_set_default_action|text|The default action of allow or deny when no other rules match|
|quarantine_policy_status|text|The value that indicates whether the policy is enabled or not|
|trust_policy_type|text|The type of trust policy|
|trust_policy_status|text|The value that indicates whether the policy is enabled or not|
|retention_policy_days|integer|The number of days to retain an untagged manifest after which it gets purged|
|retention_policy_last_updated_time|timestamp without time zone|The timestamp when the policy was last updated|
|retention_policy_status|text|The value that indicates whether the policy is enabled or not|
|id|text|The resource ID|
|name|text|The name of the resource|
|type|text|The type of the resource|
|location|text|The location of the resource|
|tags|jsonb|The tags of the resource|
