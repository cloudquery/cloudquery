\set framework 'cis_v1.6.0'
\echo "Creating CIS V1.6.0 Section 3 Views"
-- \ir ../views/project_policy_members.sql
\echo "Executing CIS V1.6.0 Section 3"
\echo "3 Control Plane Configuration"

\echo "3.1 Authentication and Authorization"
\set check_id '3.1.1'
\echo "Executing check 3.1.1"
\echo "Client certificate authentication should not be used for users (Manual)"
\ir ../queries/manual.sql


\echo "3.2 Logging"
\set check_id '3.2.1'
\echo "Executing check 3.2.1"
\echo "Ensure that a minimal audit policy is created (Manual)"
\ir ../queries/manual.sql

\set check_id '3.2.2'
\echo "Executing check 3.2.2"
\echo "Ensure that the audit policy covers key security concerns (Manual)"
\ir ../queries/manual.sql