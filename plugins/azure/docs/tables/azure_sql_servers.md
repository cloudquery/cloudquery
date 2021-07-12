
# Table: azure_sql_servers
Azure sql server
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|subscription_id|text|Azure subscription id|
|identity_principal_id|uuid|The Azure Active Directory principal id|
|identity_type|text|The identity type Set this to 'SystemAssigned' in order to automatically create and assign an Azure Active Directory principal for the resource Possible values include: 'None', 'SystemAssigned', 'UserAssigned'|
|identity_tenant_id|uuid|The Azure Active Directory tenant id|
|kind|text|Kind of sql server This is metadata used for the Azure portal experience|
|administrator_login|text|Administrator username for the server Once created it cannot be changed|
|administrator_login_password|text|The administrator login password (required for server creation)|
|version|text|The version of the server|
|state|text|The state of the server|
|fully_qualified_domain_name|text|The fully qualified domain name of the server|
|minimal_tls_version|text|Minimal TLS version Allowed values: '10', '11', '12'|
|public_network_access|text|Whether or not public endpoint access is allowed for this server  Value is optional but if passed in, must be 'Enabled' or 'Disabled' Possible values include: 'ServerPublicNetworkAccessEnabled', 'ServerPublicNetworkAccessDisabled'|
|location|text|Resource location|
|tags|jsonb|Resource tags|
|id|text|Resource ID|
|name|text|Resource name|
|type|text|Resource type|
