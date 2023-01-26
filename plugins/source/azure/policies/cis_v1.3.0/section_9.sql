\echo "Executing CIS V1.3.0 Section 9"
\set check_id '9.1'
\echo "Executing check 9.1"
\ir ../queries/web/app_auth_unset.sql
\echo "Executing check 9.2"
\ir ../queries/web/app_allow_http.sql
\set check_id '9.3'
\echo "Executing check 9.3"
\ir ../queries/web/app_using_old_tls.sql
\set check_id '9.4'
\echo "Executing check 9.4"
\ir ../queries/web/app_client_cert_disabled.sql
\set check_id '9.5'
\echo "Executing check 9.5"
\ir ../queries/web/app_register_with_ad_disabled.sql
\set check_id '9.6'
\echo "Executing check 9.6"
\echo "Check must be done manually"
\set check_id '9.7'
\echo "Executing check 9.7"
\echo "Check must be done manually"
\set check_id '9.8'
\echo "Executing check 9.8"
\echo "Check must be done manually"
\set check_id '9.9'
\echo "Executing check 9.9"
\echo "Check must be done manually"
-- todo add a publishing profiles currently they are returned as XML document
-- \set check_id '9.10'
-- \echo "Executing check 9.10"
-- \ir ../queries/web/app_ftp_deployment_enabled.sql
\set check_id '9.11'
\echo "Executing check 9.11"
\echo "Check must be done manually"
