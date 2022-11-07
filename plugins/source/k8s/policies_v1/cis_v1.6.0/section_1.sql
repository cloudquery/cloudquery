\set framework 'cis_v1.6.0'
\echo "Creating CIS V1.6.0 Section 1 Views"
\ir ../views/project_policy_members.sql
\echo "Executing CIS V1.6.0 Section 1"
\echo "Control Plane Components"
\echo "1.1 Master Node Configuration Files"

\set check_id  '1.1.1' 
\echo "Executing check 1.1.1"
\echo "Ensure that the API server pod specification file permissions are set to 644 or more restrictive (Automated)"
\ir ../queries/manual.sql

\set check_id  '1.1.2' 
\echo "Executing check 1.1.2"
\echo "Ensure that the API server pod specification file ownership is set to root:root (Automated)"
\ir ../queries/manual.sql

\set check_id  '1.1.3' 
\echo "Executing check 1.1.3"
\echo "Ensure that the controller manager pod specification file permissions are set to 644 or more restrictive (Automated)"
\ir ../queries/manual.sql

\set check_id  '1.1.4' 
\echo "Executing check 1.1.4"
\echo "Ensure that the controller manager pod specification file ownership is set to root:root (Automated)"
\ir ../queries/manual.sql

\set check_id  '1.1.5' 
\echo "Executing check 1.1.5"
\echo "Ensure that the scheduler pod specification file permissions are set to 644 or more restrictive (Automated)"
\ir ../queries/manual.sql

\set check_id  '1.1.6' 
\echo "Executing check 1.1.6"
\echo "Ensure that the scheduler pod specification file ownership is set to root:root (Automated)"
\ir ../queries/manual.sql

\set check_id  '1.1.7' 
\echo "Executing check 1.1.7"
\echo "Ensure that the etcd pod specification file permissions are set to 644 or more restrictive (Automated)"
\ir ../queries/manual.sql

\set check_id  '1.1.8' 
\echo "Executing check 1.1.8"
\echo "Ensure that the etcd pod specification file ownership is set to root:root (Automated)"
\ir ../queries/manual.sql

\set check_id  '1.1.9' 
\echo "Executing check 1.1.9"
\echo "Ensure that the Container Network Interface file permissions are set to 644 or more restrictive (Manual)"
\ir ../queries/manual.sql

\set check_id  '1.1.10'
\echo "Executing check 1.1.10"
\echo  "Ensure that the Container Network Interface file ownership is set to root:root (Manual)"
\ir ../queries/manual.sql
