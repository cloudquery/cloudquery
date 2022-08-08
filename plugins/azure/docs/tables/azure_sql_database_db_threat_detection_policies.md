
# Table: azure_sql_database_db_threat_detection_policies
DatabaseSecurityAlertPolicy contains information about a database Threat Detection policy
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|database_cq_id|uuid|Unique CloudQuery ID of azure_sql_databases table (FK)|
|location|text|The geo-location where the resource lives|
|kind|text|Resource kind|
|state|text|Specifies the state of the policy|
|disabled_alerts|text|Specifies the semicolon-separated list of alerts that are disabled, or empty string to disable no alerts|
|email_addresses|text|Specifies the semicolon-separated list of e-mail addresses to which the alert is sent|
|email_account_admins|text|Specifies that the alert is sent to the account administrators|
|storage_endpoint|text|Specifies the blob storage endpoint (eg|
|storage_account_access_key|text|Specifies the identifier key of the Threat Detection audit storage account|
|retention_days|integer|Specifies the number of days to keep in the Threat Detection audit logs|
|use_server_default|text|Specifies whether to use the default server policy|
|id|text|Resource ID|
|name|text|Resource name|
|type|text|Resource type|
