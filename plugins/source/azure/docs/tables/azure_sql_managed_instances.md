
# Table: azure_sql_managed_instances
ManagedInstance an Azure SQL managed instance
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|subscription_id|text|Azure subscription id|
|identity_principal_id|uuid|The Azure Active Directory principal id|
|identity_type|text|The identity type|
|identity_tenant_id|uuid|The Azure Active Directory tenant id|
|sku_name|text|The name of the SKU, typically, a letter + Number code, eg|
|sku_tier|text|The tier or edition of the particular SKU, eg|
|sku_size|text|Size of the particular SKU|
|sku_family|text|If the service has different generations of hardware, for the same SKU, then that can be captured here|
|sku_capacity|integer|Capacity of the particular SKU|
|provisioning_state|text|Possible values include: 'ProvisioningState1Creating', 'ProvisioningState1Deleting', 'ProvisioningState1Updating', 'ProvisioningState1Unknown', 'ProvisioningState1Succeeded', 'ProvisioningState1Failed'|
|managed_instance_create_mode|text|Specifies the mode of database creation|
|fully_qualified_domain_name|text|The fully qualified domain name of the managed instance|
|administrator_login|text|Administrator username for the managed instance|
|subnet_id|text|Subnet resource ID for the managed instance|
|state|text|The state of the managed instance|
|license_type|text|The license type|
|v_cores|integer|The number of vCores|
|storage_size_in_gb|integer|Storage size in GB|
|collation|text|Collation of the managed instance|
|dns_zone|text|The Dns Zone that the managed instance is in|
|dns_zone_partner|text|The resource id of another managed instance whose DNS zone this managed instance will share after creation|
|public_data_endpoint_enabled|boolean|Whether or not the public data endpoint is enabled|
|source_managed_instance_id|text|The resource identifier of the source managed instance associated with create operation of this instance|
|restore_point_in_time|timestamp without time zone||
|proxy_override|text|Connection type used for connecting to the instance|
|timezone_id|text|Id of the timezone|
|instance_pool_id|text|The Id of the instance pool this managed server belongs to|
|maintenance_configuration_id|text|Specifies maintenance configuration id to apply to this managed instance|
|minimal_tls_version|text|Minimal TLS version|
|storage_account_type|text|The storage account type used to store backups for this instance|
|zone_redundant|boolean|Whether or not the multi-az is enabled|
|location|text|Resource location|
|tags|jsonb|Resource tags|
|id|text|Resource ID|
|name|text|Resource name|
|type|text|Resource type|
