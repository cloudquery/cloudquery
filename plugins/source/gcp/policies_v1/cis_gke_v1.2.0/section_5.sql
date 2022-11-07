\set framework 'cis_gke_v1.2.0'
\echo "Creating CIS GKE V1.2.0 Section 5 Views"
\ir ../views/project_policy_members.sql
\echo "Executing CIS GKE V1.2.0 Section 5"


\echo "5 Managed services"
\echo "5.1 Image Registry and Image Scanning"

\set check_id '5.1.1'
\echo "Executing check 5.1.1"
\echo "Ensure Image Vulnerability Scanning using GCR Container Analysis or a third party provider (Manual)"
\ir ../queries/manual.sql

\set check_id '5.1.2'
\echo "Executing check 5.1.2"
\echo "Minimize user access to GCR (Manual)"
\ir ../queries/manual.sql

\set check_id '5.1.3'
\echo "Executing check 5.1.3"
\echo "Minimize cluster access to read-only for GCR (Manual)"
\ir ../queries/manual.sql

\set check_id '5.1.4'
\echo "Executing check 5.1.4"
\echo "Minimize Container Registries to only those approved (Manual)"
\ir ../queries/manual.sql

\echo "5.2 Identity and Access Management (IAM)"
\set check_id '5.2.1'
\echo "Executing check 5.2.1"
\echo "Ensure GKE clusters are not running using the Compute Engine default service account (Manual)"
\ir ../queries/manual.sql

\set check_id '5.2.2'
\echo "Executing check 5.2.2" 
\echo "Prefer using dedicated GCP Service Accounts and Workload Identity (Manual)"
\ir ../queries/manual.sql

\echo "5.3 Cloud Key Management Service (Cloud KMS)"
\set check_id '5.3.1'
\echo "Executing check 5.3.1"
\echo "Ensure Kubernetes Secrets are encrypted using keys managed in Cloud KMS (Manual)"
\ir ../queries/manual.sql

\echo "5.4 Node Metadata"
\set check_id '5.4.1'
\echo "Executing check 5.4.1"
\echo "Ensure legacy Compute Engine instance metadata APIs are Disabled (Automated)"
\ir ../queries/manual.sql

\set check_id '5.4.2'
\echo "Executing check 5.4.2"
\echo "Ensure the GKE Metadata Server is Enabled (Automated)"
\ir ../queries/manual.sql

\echo "5.5 Node Configuration and Maintenance"
\set check_id '5.5.1'
\echo "Executing check 5.5.1"
\echo "Ensure Container-Optimized OS (COS) is used for GKE node images (Automated)"
\ir ../queries/manual.sql

\set check_id '5.5.2'
\echo "Executing check 5.5.2"
\echo "Ensure Node Auto-Repair is enabled for GKE nodes (Automated)"
\ir ../queries/manual.sql

\set check_id '5.5.3'
\echo "Executing check 5.5.3"
\echo "Ensure Node Auto-Upgrade is enabled for GKE nodes (Automated)"
\ir ../queries/manual.sql

\set check_id '5.5.4'
\echo "Executing check 5.5.4"
\echo "When creating New Clusters - Automate GKE version management using Release Channels (Manual)"
\ir ../queries/manual.sql

\set check_id '5.5.5'
\echo "Executing check 5.5.5"
\echo "Ensure Shielded GKE Nodes are Enabled (Manual)"
\ir ../queries/manual.sql

\set check_id '5.5.6'
\echo "Executing check 5.5.6"
\echo "Ensure Integrity Monitoring for Shielded GKE Nodes is Enabled (Automated)"
\ir ../queries/manual.sql

\set check_id '5.5.7'
\echo "Executing check 5.5.7"
\echo "Ensure Secure Boot for Shielded GKE Nodes is Enabled (Automated)"
\ir ../queries/manual.sql


\echo "5.6 Cluster Networking"
\set check_id '5.6.1'
\echo "Executing check 5.6.1"
\echo "Enable VPC Flow Logs and Intranode Visibility (Automated)"
\ir ../queries/manual.sql

\set check_id '5.6.2'
\echo "Executing check 5.6.2"
\echo "Ensure use of VPC-native clusters (Automated)"
\ir ../queries/manual.sql

\set check_id '5.6.3'
\echo "Executing check 5.6.3"
\echo "Ensure Master Authorized Networks is Enabled (Manual)"
\ir ../queries/manual.sql

\set check_id '5.6.4'
\echo "Executing check 5.6.4"
\echo "Ensure clusters are created with Private Endpoint Enabled and Public Access Disabled (Manual)"
\ir ../queries/manual.sql

\set check_id '5.6.5'
\echo "Executing check 5.6.5"
\echo "Ensure clusters are created with Private Nodes (Manual)"
\ir ../queries/manual.sql

\set check_id '5.6.6'
\echo "Executing check 5.6.6"
\echo "Consider firewalling GKE worker nodes (Manual)"
\ir ../queries/manual.sql

\set check_id '5.6.7'
\echo "Executing check 5.6.7"
\echo "Ensure Network Policy is Enabled and set as appropriate (Manual)"
\ir ../queries/manual.sql

\set check_id '5.6.8'
\echo "Executing check 5.6.8"
\echo "Ensure use of Google-managed SSL Certificates (Manual)"
\ir ../queries/manual.sql

\echo "5.7 Logging"
\set check_id '5.7.1'
\echo "Executing check 5.7.1"
\echo "Ensure Stackdriver Kubernetes Logging and Monitoring is Enabled (Automated)"
\ir ../queries/manual.sql

\set check_id '5.7.2'
\echo "Executing check 5.7.2"
\echo "Enable Linux auditd logging (Manual)"
\ir ../queries/manual.sql


\echo "5.8 Authentication and Authorization"
\set check_id '5.8.1'
\echo "Executing check 5.8.1"
\echo "Ensure Basic Authentication using static passwords is Disabled (Automated)"
\ir ../queries/manual.sql

\set check_id '5.8.2'
\echo "Executing check 5.8.2"
\echo "Ensure authentication using Client Certificates is Disabled (Automated)"
\ir ../queries/manual.sql

\set check_id '5.8.3'
\echo "Executing check 5.8.3"
\echo "Manage Kubernetes RBAC users with Google Groups for GKE (Manual)"
\ir ../queries/manual.sql

\set check_id '5.8.4'
\echo "Executing check 5.8.4"
\echo "Ensure Legacy Authorization (ABAC) is Disabled (Automated)"
\ir ../queries/manual.sql

\echo "5.9 Storage"
\set check_id '5.9.1'
\echo "Executing check 5.9.1"
\echo "Enable Customer-Managed Encryption Keys (CMEK) for GKE Persistent Disks (PD) (Manual)"
\ir ../queries/manual.sql


\echo "5.10 Other Cluster Configurations"
\set check_id '5.10.1'
\echo "Executing check 5.10.1"
\echo "Ensure Kubernetes Web UI is Disabled (Automated)"
\ir ../queries/manual.sql

\set check_id '5.10.2'
\echo "Executing check 5.10.2"
\echo "Ensure that Alpha clusters are not used for production workloads (Automated)"
\ir ../queries/manual.sql

\set check_id '5.10.3'
\echo "Executing check 5.10.3"
\echo "Ensure Pod Security Policy is Enabled and set as appropriate (Manual)"
\ir ../queries/manual.sql

\set check_id '5.10.4'
\echo "Executing check 5.10.4"
\echo "Consider GKE Sandbox for running untrusted workloads (Manual)"
\ir ../queries/manual.sql

\set check_id '5.10.5'
\echo "Executing check 5.10.5"
\echo "Ensure use of Binary Authorization (Automated)"
\ir ../queries/manual.sql

\set check_id '5.10.6'
\echo "Executing check 5.10.6"
\echo "Enable Cloud Security Command Center (Cloud SCC) (Manual)"
\ir ../queries/manual.sql