
# Table: azure_web_app_host_name_ssl_states
HostNameSslState SSL-enabled hostname
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|app_cq_id|uuid|Unique CloudQuery ID of azure_web_apps table (FK)|
|name|text|Hostname|
|ssl_state|text|SSL type Possible values include: 'SslStateDisabled', 'SslStateSniEnabled', 'SslStateIPBasedEnabled'|
|virtual_ip|text|Virtual IP address assigned to the hostname if IP based SSL is enabled|
|thumbprint|text|SSL certificate thumbprint|
|to_update|boolean|Set to true to update existing hostname|
|host_type|text|Indicates whether the hostname is a standard or repository hostname Possible values include: 'HostTypeStandard', 'HostTypeRepository'|
