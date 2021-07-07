
# Table: azure_sql_servers
Azure sql server
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|subscription_id|text|Azure subscription id|
|kind|text|Kind of sql server  This is metadata used for the Azure portal experience|
|fully_qualified_domain_name|text|The fully qualified domain name of the server|
|version|text|The version of the server Possible values include: 'TwoFullStopZero', 'OneTwoFullStopZero'|
|administrator_login|text|Administrator username for the server Can only be specified when the server is being created (and is required for creation)|
|administrator_login_password|text|The administrator login password (required for server creation)|
|external_administrator_sid|uuid|The ID of the Active Azure Directory object with admin permissions on this server Legacy parameter, always null To check for Active Directory admin, query /servers/{serverName}/administrators|
|external_administrator_login|text|The display name of the Azure Active Directory object with admin permissions on this server Legacy parameter, always null To check for Active Directory admin, query /servers/{serverName}/administrators|
|state|text|The state of the server Possible values include: 'ServerStateReady', 'ServerStateDisabled'|
|location|text|Resource location|
|tags|jsonb|Resource tags|
|id|text|Resource ID|
|name|text|Resource name|
|type|text|Resource type|
