import { Callout } from 'nextra-theme-docs'

<Callout type="warning">

The Azure policies are currently supported by the v1 version of the plugin. Please 👍 [this issue](https://github.com/cloudquery/cloudquery/issues/5830) if you'd like us to prioritize the v2 version.

</Callout>

# Policies and Compliance Frameworks

Policies are a set of standard SQL queries that can be run to check the security and compliance of your cloud resources against best practice frameworks.

This page documents the available CloudQuery SQL Policies for Azure. See the [readme on GitHub](https://github.com/cloudquery/cloudquery/tree/main/plugins/source/azure/policies) for installation instructions.
## Azure CIS v1.3.0

### Requirements
Azure CIS v1.3.0 requires the following tables to be synced before the policy is executed:

```yaml
tables:
  - azure_compute_disks
  - azure_compute_virtual_machines
  - azure_keyvault_keys
  - azure_keyvault_secrets
  - azure_keyvault_vaults
  - azure_mysql_servers
  - azure_postgresql_configurations
  - azure_postgresql_firewall_rules
  - azure_postgresql_servers
  - azure_security_auto_provisioning_settings
  - azure_security_contacts
  - azure_security_pricings
  - azure_security_settings
  - azure_sql_database_blob_auditing_policies
  - azure_sql_database_threat_detection_policies
  - azure_sql_databases
  - azure_sql_encryption_protectors
  - azure_sql_server_admins
  - azure_sql_server_blob_auditing_policies
  - azure_sql_server_vulnerability_assessments
  - azure_sql_servers
  - azure_sql_transparent_data_encryptions
  - azure_web_apps
  - azure_web_publishing_profiles
  - azure_web_site_auth_settings
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
  - Ensure that Windows Defender ATP (WDATP) integration with Security Center is selected (Automatic)
  - Ensure that Microsoft Cloud App Security (MCAS) integration with Security Center is selected (Automatic)
  - Ensure that "Automatic provisioning of monitoring agent" is set to "On" (Automated)
  - Ensure any of the ASC Default policy setting is not set to "Disabled" (Automated)
  - Ensure "Additional email addresses" is configured with a security contact email (Automated)
  - Ensure that "Notify about alerts with the following severity" is set to "High" (Automated)
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
  - Ensure Virtual Machines are utilizing Managed Disks (Manual)
  - Ensure that ''OS and Data'' disks are encrypted with CMK (Automated)
  - Ensure that ''Unattached disks'' are encrypted with CMK (Automated)
  - Ensure that VHD''s are encrypted (Manual)
  - Ensure that the expiration date is set on all keys (Automated)
  - Ensure that the expiration date is set on all Secrets (Automated)
  - Ensure the key vault is recoverable (Automated)
  - Ensure App Service Authentication is set on Azure App Service (Automated)
  - Ensure web app redirects all HTTP traffic to HTTPS in Azure App Service (Automated)
  - Ensure the web app has ''Client Certificates (Incoming client certificates)'' set to ''On'' (Automated)
  - Ensure that Register with Azure Active Directory is enabled on App Service (Automated)
  - Ensure FTP deployments are disabled (Automated)
## Azure HIPAA HITRUST v9.2

### Requirements
Azure HIPAA HITRUST v9.2 requires the following tables to be synced before the policy is executed:

```yaml
tables:
  - azure_authorization_role_assignments
  - azure_authorization_role_definitions
  - azure_batch_accounts
  - azure_compute_virtual_machine_extensions
  - azure_compute_virtual_machine_scale_sets
  - azure_compute_virtual_machines
  - azure_container_managed_clusters
  - azure_container_registries
  - azure_cosmosdb_accounts
  - azure_datalake_store_accounts
  - azure_eventhub_namespaces
  - azure_eventhub_network_rule_sets
  - azure_keyvault_managed_hsms
  - azure_keyvault_vaults
  - azure_logic_diagnostic_settings
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
  - azure_sql_backup_long_term_retention_policies
  - azure_sql_database_vulnerability_assessment_scans
  - azure_sql_databases
  - azure_sql_encryption_protectors
  - azure_sql_managed_instance_encryption_protectors
  - azure_sql_managed_instance_vulnerability_assessments
  - azure_sql_managed_instances
  - azure_sql_server_blob_auditing_policies
  - azure_sql_server_vulnerability_assessments
  - azure_sql_servers
  - azure_sql_transparent_data_encryptions
  - azure_sql_virtual_network_rules
  - azure_storage_accounts
  - azure_streamanalytics_streaming_jobs
  - azure_subscriptions
  - azure_subscriptions_locations
  - azure_web_apps
  - azure_web_vnet_connections
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
