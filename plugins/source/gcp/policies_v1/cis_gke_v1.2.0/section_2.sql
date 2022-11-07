\set framework 'cis_gke_v1.2.0'
\echo "Creating CIS GKE V1.2.0 Section 2 Views"
\ir ../views/project_policy_members.sql
\echo "Executing CIS GKE V1.2.0 Section 2"

\echo "2 Control Plane Configuration"
\echo "2.1 Authentication and Authorization"
\set check_id '2.1.1'
\echo "Executing check 2.1.1"
\echo "Client certificate authentication should not be used for users (Manual)"
\ir ../queries/manual.sql

\echo "2.2 Logging"

\set check_id '2.2.1'
\echo "Executing check 2.2.1"
\echo "Ensure that a minimal audit policy is created (Manual)"
\ir ../queries/manual.sql

\set check_id '2.2.2'
\echo "Executing check 2.2.2"
\echo "Ensure that the audit policy covers key security concerns (Manual)"
\ir ../queries/manual.sql

