# Policies and Compliance Frameworks

Policies are a set of standard SQL queries that can be run to check the security and compliance of your cloud resources against best practice frameworks.

This page documents the available CloudQuery SQL Policies for Azure. See the [readme on GitHub](https://github.com/cloudquery/cloudquery/tree/main/plugins/source/azure/policies) for installation instructions.
## Azure CIS v1.3.0

### Requirements
Azure CIS v1.3.0 requires the following tables to be synced before the policy is executed:

```yaml
tables:
  - azure_appservice_web_app_auth_settings
  - azure_appservice_web_app_configurations
  - azure_appservice_web_apps
  - azure_compute_disks
  - azure_compute_virtual_machines
  - azure_keyvault_keyvault
  - azure_keyvault_keyvault_keys
  - azure_keyvault_keyvault_secrets
  - azure_monitor_activity_log_alerts
  - azure_monitor_diagnostic_settings
  - azure_monitor_resources
  - azure_monitor_subscription_diagnostic_settings
  - azure_mysql_servers
  - azure_network_security_groups
  - azure_policy_assignments
  - azure_postgresql_server_configurations
  - azure_postgresql_server_firewall_rules
  - azure_postgresql_servers
  - azure_security_auto_provisioning_settings
  - azure_security_pricings
  - azure_sql_server_admins
  - azure_sql_server_advanced_threat_protection_settings
  - azure_sql_server_blob_auditing_policies
  - azure_sql_server_database_blob_auditing_policies
  - azure_sql_server_databases
  - azure_sql_server_encryption_protectors
  - azure_sql_server_vulnerability_assessments
  - azure_sql_servers
  - azure_sql_transparent_data_encryptions
  - azure_storage_accounts
  - azure_storage_blob_services
  - azure_storage_containers
```

### Queries

Azure CIS v1.3.0 performs the following checks:
  - Ensure that Azure Defender is set to On for Servers (Automatic)
  - Ensure that Azure Defender is set to On for App Service (Automatic)
  - Ensure that Azure Defender is set to On for Azure SQL database servers (Automatic)
  - Ensure that Azure Defender is set to On for SQL servers on machines (Automatic)
  - Ensure that Azure Defender is set to On for Storage (Automatic)
  - Ensure that Azure Defender is set to On for Kubernetes (Automatic)
  - Ensure that Azure Defender is set to On for Container Registries (Automatic)
  - Ensure that Azure Defender is set to On for Key Vault (Manual)
  - Ensure that "Automatic provisioning of monitoring agent" is set to "On" (Automated)
  - Ensure any of the ASC Default policy setting is not set to "Disabled" (Automated)
  - Secure transfer to storage accounts should be enabled
  - Ensure that ''Public access level'' is set to Private for blob containers
  - Ensure default network access rule for Storage Accounts is set to deny
  - Ensure soft delete is enabled for Azure Storage
  - Ensure storage for critical data are encrypted with Customer Managed Key
  - Ensure that "Auditing" is set to "On" (Automated)
  - Ensure that "Data encryption" is set to "On" on a SQL Database (Automated)
  - Ensure that "Auditing" Retention is "greater than 90 days" (Automated)
  - Ensure that Advanced Threat Protection (ATP) on a SQL server is set to "Enabled" (Automated)
  - Ensure that Vulnerability Assessment (VA) is enabled on a SQL server by setting a Storage Account (Automated)
  - Ensure that VA setting Periodic Recurring Scans is enabled on a SQL server (Automated)
  - Ensure that VA setting Send scan reports to is configured for a SQL server (Automated)
  - Ensure that VA setting "Also send email notifications to admins and subscription owners" is set for a SQL server (Automated)
  - Ensure "Enforce SSL connection" is set to "ENABLED" for PostgreSQL Database Server (Automated)
  - Ensure "Enforce SSL connection" is set to "ENABLED" for MySQL Database Server (Automated)
  - Ensure server parameter "log_checkpoints" is set to "ON" for PostgreSQL Database Server (Automated)
  - Ensure server parameter "log_connections" is set to "ON" for PostgreSQL Database Server (Automated)
  - Ensure server parameter "log_disconnections" is set to "ON" for PostgreSQL Database Server (Automated)
  - Ensure server parameter "connection_throttling" is set to "ON" for PostgreSQL Database Server (Automated)
  - Ensure server parameter "log_retention_days" is greater than 3 days for PostgreSQL Database Server (Automated)
  - Ensure "Allow access to Azure services" for PostgreSQL Database Server is disabled (Automated)
  - Ensure that Azure Active Directory Admin is configured (Automated)
  - Ensure SQL server"s TDE protector is encrypted with Customer-managed key (Automated)
  - Ensure that a ''Diagnostics Setting'' exists
  - Ensure Diagnostic Setting captures appropriate categories
  - Ensure the storage container storing the activity logs is not publicly accessible
  - Ensure the storage account containing the container with activity logs is encrypted with BYOK (Use Your Own Key)
  - Ensure that logging for Azure Key Vault is ''Enabled''
  - Ensure that Activity Log Alert exists for Create Policy Assignment
  - Ensure that Activity Log Alert exists for Delete Policy Assignment
  - Ensure that Activity Log Alert exists for Create or Update Network Security Group
  - Ensure that Activity Log Alert exists for Delete Network Security Group
  - Ensure that Activity Log Alert exists for Create or Update Network Security Group Rule
  - Ensure that Activity Log Alert exists for Delete Network Security Group Rule
  - Ensure that Activity Log Alert exists for Create or Update Security Solution
  - Ensure that Activity Log Alert exists for Delete Security Solution
  - Ensure that Activity Log Alert exists for Create or Update or Delete SQL Server Firewall Rule
  - Ensure that Diagnostic Logs are enabled for all services which support it.
  - Ensure that RDP access is restricted from the Internet
  - Ensure that SSH access is restricted from the Internet
  - Ensure that UDP Services are restricted from the Internet
  - Ensure Virtual Machines are utilizing Managed Disks (Manual)
  - Ensure that ''OS and Data'' disks are encrypted with CMK (Automated)
  - Ensure that ''Unattached disks'' are encrypted with CMK (Automated)
  - Ensure that VHD''s are encrypted (Manual)
  - Ensure that the expiration date is set on all keys (Automated)
  - Ensure that the expiration date is set on all Secrets (Automated)
  - Ensure the key vault is recoverable (Automated)
  - Ensure App Service Authentication is set on Azure App Service (Automated)
  - Ensure web app redirects all HTTP traffic to HTTPS in Azure App Service (Automated)
  - Ensure web app is using the latest version of TLS encryption (Automated)
  - Ensure the web app has ''Client Certificates (Incoming client certificates)'' set to ''On'' (Automated)
  - Ensure that Register with Azure Active Directory is enabled on App Service (Automated)
  - Ensure FTP deployments are disabled (Automated)

### Dependent Views

Azure CIS v1.3.0 depends on the following views:

  - view_azure_nsg_dest_port_ranges<sup>*</sup>
  - view_azure_security_policy_parameters<sup>*</sup>

  <sup>*</sup> These views are automatically created or updated by this policy.
## Azure HIPAA HITRUST v9.2

### Requirements
Azure HIPAA HITRUST v9.2 requires the following tables to be synced before the policy is executed:

```yaml
tables:
  - azure_appservice_web_app_vnet_connections
  - azure_appservice_web_apps
  - azure_authorization_role_assignments
  - azure_authorization_role_definitions
  - azure_batch_account
  - azure_compute_virtual_machine_extensions
  - azure_compute_virtual_machine_scale_sets
  - azure_compute_virtual_machines
  - azure_containerregistry_registries
  - azure_containerservice_managed_clusters
  - azure_cosmos_database_accounts
  - azure_datalakestore_accounts
  - azure_eventhub_namespace_network_rule_sets
  - azure_eventhub_namespaces
  - azure_keyvault_keyvault
  - azure_keyvault_keyvault_managed_hsms
  - azure_logic_workflows
  - azure_mariadb_servers
  - azure_monitor_activity_log_alerts
  - azure_monitor_diagnostic_settings
  - azure_monitor_log_profiles
  - azure_monitor_resources
  - azure_mysql_servers
  - azure_network_interfaces
  - azure_network_security_groups
  - azure_network_virtual_networks
  - azure_network_watchers
  - azure_postgresql_servers
  - azure_redis_caches
  - azure_resources_links
  - azure_search_services
  - azure_security_assessments
  - azure_security_auto_provisioning_settings
  - azure_security_jit_network_access_policies
  - azure_sql_managed_instance_encryption_protectors
  - azure_sql_managed_instance_vulnerability_assessments
  - azure_sql_managed_instances
  - azure_sql_server_blob_auditing_policies
  - azure_sql_server_database_long_term_retention_policies
  - azure_sql_server_database_vulnerability_assessment_scans
  - azure_sql_server_database_vulnerability_assessments
  - azure_sql_server_databases
  - azure_sql_server_encryption_protectors
  - azure_sql_server_virtual_network_rules
  - azure_sql_server_vulnerability_assessments
  - azure_sql_servers
  - azure_sql_transparent_data_encryptions
  - azure_storage_accounts
  - azure_streamanalytics_streaming_jobs
  - azure_subscription_subscription_locations
  - azure_subscription_subscriptions
```

### Queries

Azure HIPAA HITRUST v9.2 performs the following checks:
  - MFA should be enabled on accounts with owner permissions on your subscription
  - MFA should be enabled on accounts with write permissions on your subscription
  - MFA should be enabled on accounts with owner permissions on your subscription
  - Management ports of virtual machines should be protected with just-in-time network access control
  - [Preview]: Container Registry should use a virtual network service endpoint
  - App Service should use a virtual network service endpoint
  - Cosmos DB should use a virtual network service endpoint
  - Event Hub should use a virtual network service endpoint
  - Gateway subnets should not be configured with a network security group
  - Internet-facing virtual machines should be protected with network security groups
  - Key Vault should use a virtual network service endpoint
  - SQL Server should use a virtual network service endpoint
  - Storage Accounts should use a virtual network service endpoint
  - Subnets should be associated with a Network Security Group
  - Virtual machines should be connected to an approved virtual network
  - API App should only be accessible over HTTPS
  - Enforce SSL connection should be enabled for MySQL database servers
  - Enforce SSL connection should be enabled for PostgreSQL database servers
  - Function App should only be accessible over HTTPS
  - Latest TLS version should be used in your API App
  - Latest TLS version should be used in your Function App
  - Latest TLS version should be used in your Web App
  - Only secure connections to your Azure Cache for Redis should be enabled
  - Secure transfer to storage accounts should be enabled
  - Web Application should only be accessible over HTTPS
  - A maximum of 3 owners should be designated for your subscription
  - There should be more than one owner assigned to your subscription
  - Resource logs in Azure Data Lake Store should be enabled
  - Resource logs in Logic Apps should be enabled
  - Resource logs in Batch accounts should be enabled
  - Resource logs in Virtual Machine Scale Sets should be enabled
  - Resource logs in Azure Stream Analytics should be enabled
  - Resource logs in Event Hub should be enabled
  - Resource logs in Search services should be enabled
  - Resource logs in App Services should be enabled
  - Auditing on SQL server should be enabled
  - Resource logs in Azure Key Vault Managed HSM should be enabled
  - Resource logs in Key Vault should be enabled
  - Azure Monitor should collect activity logs from all regions
  - Virtual machines should have the Log Analytics extension installed
  - The Log Analytics extension should be installed on Virtual Machine Scale Sets
  - Audit Windows machines on which the Log Analytics agent is not connected as expected
  - Azure Monitor log profile should collect logs for categories ''write,'' ''delete,'' and ''action''
  - Auto provisioning of the Log Analytics agent should be enabled on your subscription
  - An activity log alert should exist for specific Administrative operations
  - Role-Based Access Control (RBAC) should be used on Kubernetes Services
  - External accounts with owner permissions should be removed from your subscription
  - Windows machines should meet requirements for ''User Rights Assignment''
  - Deploy default Microsoft IaaSAntimalware extension for Windows Server
  - Endpoint protection solution should be installed on virtual machine scale sets
  - Microsoft Antimalware for Azure should be configured to automatically update protection signatures
  - System updates should be installed on your machines
  - Long-term geo-redundant backup should be enabled for Azure SQL Databases
  - Geo-redundant backup should be enabled for Azure Database for MySQL
  - Geo-redundant backup should be enabled for Azure Database for PostgreSQL
  - Geo-redundant backup should be enabled for Azure Database for MariaDB
  - All network ports should be restricted on network security groups associated to your virtual machine
  - Storage accounts should restrict network access
  - [Preview]: Network traffic data collection agent should be installed on Windows virtual machines
  - Virtual machines should be migrated to new Azure Resource Manager resources
  - [Preview]: Network traffic data collection agent should be installed on Linux virtual machines
  - Network Watcher should be enable
  - Ensure that "Data encryption" is set to "On" on a SQL Database (Automated)
  - Require encryption on Data Lake Store accounts
  - SQL managed instances should use customer-managed keys to encrypt data at rest
  - Ensure SQL server"s TDE protector is encrypted with Customer-managed key (Automated)
  - Ensure the web app has ''Client Certificates (Incoming client certificates)'' set to ''On'' (Automated)
  - CORS should not allow every resource to access your Web Applications
  - CORS should not allow every resource to access your Function Apps
  - CORS should not allow every resource to access your API App
  - Remote debugging should be turned off for Web Applications
  - Remote debugging should be turned off for Function Apps
  - Remote debugging should be turned off for API Apps
  - A vulnerability assessment solution should be enabled on your virtual machines
  - SQL databases should have vulnerability findings resolved
  - Vulnerability assessment should be enabled on SQL Managed Instance
  - Vulnerability assessment should be enabled on your SQL servers
  - Audit virtual machines without disaster recovery configured.
  - Azure Key Vault Managed HSM should have purge protection enabled
  - Ensure the key vault is recoverable (Automated)

### Dependent Views

Azure HIPAA HITRUST v9.2 depends on the following views:

  - view_azure_nsg_rules<sup>*</sup>

  <sup>*</sup> This view is automatically created or updated by this policy.
