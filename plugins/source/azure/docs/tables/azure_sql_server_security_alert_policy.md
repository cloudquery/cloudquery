
# Table: azure_sql_server_security_alert_policy
List the server's threat detection policies
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|server_cq_id|uuid|Unique ID of azure_sql_servers table (FK)|
|id|text|Resource ID|
|name|text|Resource name|
|type|text|The virtual network rule type|
|state|text|Specifies the state of the policy, whether it is enabled or disabled or a policy has not been applied yet on the specific database. Possible values include: 'SecurityAlertPolicyStateNew', 'SecurityAlertPolicyStateEnabled', 'SecurityAlertPolicyStateDisabled'|
|disabled_alerts|text[]|Specifies an array of alerts that are disabled. Allowed values are: Sql_Injection, Sql_Injection_Vulnerability, Access_Anomaly, Data_Exfiltration, Unsafe_Action|
|email_addresses|text[]|Specifies an array of e-mail addresses to which the alert is sent.|
|email_account_admins|boolean|Specifies that the alert is sent to the account administrators.|
|storage_endpoint|text|Specifies the blob storage endpoint (e.g. https://MyAccount.blob.core.windows.net). This blob storage will hold all Threat Detection audit logs.|
|storage_account_access_key|text|Specifies the identifier key of the Threat Detection audit storage account.|
|retention_days|integer|Specifies the number of days to keep in the Threat Detection audit logs.|
|creation_time|timestamp without time zone|Specifies the UTC creation time of the policy.|
