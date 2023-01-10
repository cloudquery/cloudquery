# Policies and Compliance Frameworks

Policies are a set of standard SQL queries that can be run to check the security and compliance of your cloud resources against best practice frameworks.

This page documents the available CloudQuery SQL Policies for GCP. See the [readme on GitHub](https://github.com/cloudquery/cloudquery/tree/main/plugins/source/gcp/policies) for installation instructions.
## GCP CIS v1.2.0

### Requirements
GCP CIS v1.2.0 requires the following tables to be synced before the policy is executed:

```yaml
tables:
  - gcp_bigquery_dataset_accesses
  - gcp_bigquery_dataset_tables
  - gcp_bigquery_datasets
  - gcp_bigquery_tables
  - gcp_compute_disks
  - gcp_compute_firewalls
  - gcp_compute_instance_service_accounts
  - gcp_compute_instances
  - gcp_compute_networks
  - gcp_compute_projects
  - gcp_compute_subnetworks
  - gcp_dns_managed_zone_dnssec_config_default_key_specs
  - gcp_dns_managed_zones
  - gcp_dns_policies
  - gcp_dns_policy_networks
  - gcp_firewall_allowed_rules
  - gcp_iam_service_account_keys
  - gcp_iam_service_accounts
  - gcp_kms_crypto_keys
  - gcp_kms_keyring_crypto_keys
  - gcp_kms_keyrings
  - gcp_log_metric_filters
  - gcp_logging_metrics
  - gcp_logging_sinks
  - gcp_project_policy_members
  - gcp_public_buckets_accesses
  - gcp_resourcemanager_project_policies
  - gcp_sql_instances
  - gcp_storage_buckets
```

### Queries
GCP CIS v1.2.0 performs the following checks:
  - Ensure that there are only GCP-managed service account keys for each service account (Automated)
  - Ensure that Service Account has no Admin privileges (Automated)
  - Ensure that IAM users are not assigned the Service Account User or Service Account Token Creator roles at project level (Automated)
  - Ensure user-managed/external keys for service accounts are rotated every 90 days or less (Automated)
  - Ensure that Separation of duties is enforced while assigning service account related roles to users (Automated)
  - Ensure that Cloud KMS cryptokeys are not anonymously or publicly accessible (Automated)
  - Ensure KMS encryption keys are rotated within a period of 90 days (Automated)
  - Ensure that Separation of duties is enforced while assigning KMS related roles to users (Automated)
  - Ensure that Cloud Audit Logging is configured properly across all services and all users from a project (Automated)
  - Ensure that sinks are configured for all log entries (Automated)
  - Ensure that retention policies on log buckets are configured using Bucket Lock (Automated)
  - Ensure log metric filter and alerts exist for project ownership assignments/changes (Automated)
  - Ensure that the log metric filter and alerts exist for Audit Configuration changes (Automated)
  - Ensure that the log metric filter and alerts exist for Custom Role changes (Automated)
  - Ensure that the log metric filter and alerts exist for VPC Network Firewall rule changes (Automated)
  - Ensure that the log metric filter and alerts exist for VPC network route changes (Automated)
  - Ensure that the log metric filter and alerts exist for VPC network changes (Automated)
  - Ensure that the log metric filter and alerts exist for Cloud Storage IAM permission changes (Automated)
  - Ensure that the log metric filter and alerts exist for SQL instance configuration changes (Automated)
  - Ensure that Cloud DNS logging is enabled for all VPC networks (Automated)
  - Ensure that the default network does not exist in a project (Automated)
  - Ensure legacy networks do not exist for a project (Automated)
  - Ensure that DNSSEC is enabled for Cloud DNS (Automated)
  - Ensure that DNSSEC is enabled for Cloud DNS (Automated)
  - Ensure that RSASHA1 is not used for the zone-signing key in Cloud DNS DNSSEC (Manual)
  - Ensure that SSH access is restricted from the internet (Automated)
  - Ensure that RDP access is restricted from the Internet (Automated)
  - Ensure that VPC Flow Logs is enabled for every subnet in a VPC Network (Automated)
  - GCP CIS3.10 Ensure Firewall Rules for instances behind Identity Aware Proxy (IAP) only allow the traffic from Google Cloud Loadbalancer (GCLB) Health Check and Proxy Addresses (Manual)
  - Ensure that instances are not configured to use the default service account (Automated)
  - Ensure that instances are not configured to use the default service account with full access to all Cloud APIs (Automated)
  - Ensure "Block Project-wide SSH keys" is enabled for VM instances (Automated)
  - Ensure oslogin is enabled for a Project (Automated)
  - Ensure "Enable connecting to serial ports" is not enabled for VM Instance (Automated)
  - Ensure that IP forwarding is not enabled on Instances (Automated)
  - Ensure VM disks for critical VMs are encrypted with Customer-Supplied Encryption Keys (CSEK) (Automated)
  - Ensure Compute instances are launched with Shielded VM enabled (Automated)
  - Ensure that Compute instances do not have public IP addresses (Automated
  - Ensure that Compute instances have Confidential Computing enabled (Automated)
  - Ensure that Cloud Storage bucket is not anonymously or publicly accessible (Automated)
  - Ensure that Cloud Storage buckets have uniform bucket-level access enabled (Automated)
  - Ensure "skip_show_database" database flag for Cloud SQL Mysql instance is set to "on" (Automated)
  - Ensure that the "local_infile" database flag for a Cloud SQL Mysql instance is set to "off" (Automated)
  - Ensure that the "log_checkpoints" database flag for Cloud SQL PostgreSQL instance is set to "on" (Automated)
  - Ensure "log_error_verbosity" database flag for Cloud SQL PostgreSQL instance is set to "DEFAULT" or stricter (Manual)
  - Ensure that the "log_connections" database flag for Cloud SQL PostgreSQL instance is set to "on" (Automated)
  - Ensure that the log_disconnections" database flag for Cloud SQL PostgreSQL instance is set to "on" (Automated)
  - Ensure "log_duration" database flag for Cloud SQL PostgreSQL instance is set to "on" (Manual)
  - Ensure that the "log_lock_waits" database flag for Cloud SQL PostgreSQL instance is set to "on" (Automated)
  - Ensure "log_hostname" database flag for Cloud SQL PostgreSQL instance is set appropriately (Automated)
  - Ensure "log_parser_stats" database flag for Cloud SQL PostgreSQL instance is set to "off" (Automated)
  - Ensure "log_planner_stats" database flag for Cloud SQL PostgreSQL instance is set to "off" (Automated)
  - Ensure "log_executor_stats" database flag for Cloud SQL PostgreSQL instance is set to "off" (Automated)
  - Ensure "log_statement_stats" database flag for Cloud SQL PostgreSQL instance is set to "off" (Automated)
  - Ensure that the "log_min_messages" database flag for Cloud SQL PostgreSQL instance is set appropriately (Manual)
  - Ensure that the "log_temp_files" database flag for Cloud SQL PostgreSQL instance is set to "0" (on) (Automated)
  - Ensure that the "log_min_duration_statement" database flag for Cloud SQL PostgreSQL instance is set to "-1" (disabled) (Automated)
  - Ensure "external scripts enabled" database flag for Cloud SQL SQL Server instance is set to "off" (Automated)
  - Ensure that the "cross db ownership chaining" database flag for Cloud SQL SQL Server instance is set to "off" (Automated)
  - Ensure "user connections" database flag for Cloud SQL SQL Server instance is set as appropriate (Automated)
  - Ensure "user options" database flag for Cloud SQL SQL Server instance is not configured (Automated)
  - Ensure "remote access" database flag for Cloud SQL SQL Server instance is set to "off" (Automated)
  - Ensure "3625 (trace flag)" database flag for Cloud SQL SQL Server instance is set to "off" (Automated)
  - Ensure that the "contained database authentication" database flag for Cloud SQL on the SQL Server instance is set to "off" (Automated)
  - Ensure that the Cloud SQL database instance requires all incoming connections to use SSL (Automated)
  - Ensure that Cloud SQL database instances are not open to the world (Automated)
  - Ensure that Cloud SQL database instances do not have public IPs (Automated)
  - Ensure that Cloud SQL database instances are configured with automated backups (Automated)
  - Ensure that BigQuery datasets are not anonymously or publicly accessible (Automated)
  - Ensure that a Default Customer-managed encryption key (CMEK) is specified for all BigQuery Data Sets (Automated)
  - Ensure that all BigQuery Tables are encrypted with Customer-managed encryption key (CMEK) (Automated)
