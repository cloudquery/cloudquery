\set framework 'cis_eks_v1.1.0'
\echo "Creating CIS EKS V1.1.0 Section 2 Views"
\ir ../views/project_policy_members.sql
\echo "Executing CIS EKS V1.1.0 Section 2"

\echo "2 Control Plane Configuration"
\echo "2.1 Logging"
\set check_id '2.1.1'
\echo "Executing check 2.1.1"
\echo "Enable audit Logs (Manual)"
\ir ../queries/manual.sql