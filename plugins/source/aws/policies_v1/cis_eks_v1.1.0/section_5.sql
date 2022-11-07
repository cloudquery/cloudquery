\set framework 'cis_eks_v1.1.0'
\echo "Creating CIS EKS V1.1.0 Section 5 Views"
\ir ../views/project_policy_members.sql
\echo "Executing CIS EKS V1.1.0 Section 5"


\echo "5 Managed services"
\echo "5.1 Image Registry and Image Scanning"

\set check_id '5.1.1'
\echo "Executing check 5.1.1"
\echo "Ensure Image Vulnerability Scanning using Amazon ECR image scanning or a third party provider (Manual)"
\ir ../queries/manual.sql

\set check_id '5.1.2'
\echo "Executing check 5.1.2"
\echo "Minimize user access to Amazon ECR (Manual)"
\ir ../queries/manual.sql

\set check_id '5.1.3'
\echo "Executing check 5.1.3"
\echo "Minimize cluster access to read-only for Amazon ECR (Manual)"
\ir ../queries/manual.sql

\set check_id '5.1.4'
\echo "Executing check 5.1.4"
\echo "Minimize Container Registries to only those approved (Manual)"
\ir ../queries/manual.sql


\echo "5.2 Identity and Access Management (IAM)"
\set check_id '5.2.1'
\echo "Executing check 5.2.1"
\echo "Prefer using managed identities for workloads (Manual)"
\ir ../queries/manual.sql


\echo "5.3 AWS EKS Key Management Service"
\set check_id '5.3.1'
\echo "Executing check 5.3.1"
\echo "Ensure Kubernetes Secrets are encrypted using Customer Master Keys (CMKs) managed in AWS KMS (Manual)"
\ir ../queries/manual.sql


\echo "5.4 Cluster Networking"
\set check_id '5.4.1'
\echo "Executing check 5.4.1"
\echo "Restrict Access to the Control Plane Endpoint (Manual)"
\ir ../queries/manual.sql

\set check_id '5.4.2'
\echo "Executing check 5.4.2"
\echo "Ensure clusters are created with Private Endpoint Enabled and Public Access Disabled (Manual)"
\ir ../queries/manual.sql

\set check_id '5.4.3'
\echo "Executing check 5.4.3"
\echo "Ensure clusters are created with Private Nodes (Manual)"
\ir ../queries/manual.sql

\set check_id '5.4.4'
\echo "Executing check 5.4.4"
\echo "Ensure Network Policy is Enabled and set as appropriate (Manual)"
\ir ../queries/manual.sql

\set check_id '5.4.5'
\echo "Executing check 5.4.5"
\echo "Encrypt traffic to HTTPS load balancers with TLS certificates (Manual)"
\ir ../queries/manual.sql

\echo "5.5 Authentication and Authorization"
\set check_id '5.5.1'
\echo "Executing check 5.5.1"
\echo "Manage Kubernetes RBAC users with AWS IAM Authenticator for Kubernetes (Manual)"
\ir ../queries/manual.sql


\echo "5.6 Other Cluster Configurations"
\set check_id '5.6.1'
\echo "Executing check 5.6.1"
\echo "Consider Fargate for running untrusted workloads (Manual)"
\ir ../queries/manual.sql