\set check_id 'ELB.2'
\echo "Executing check ELB.2"
\i queries/elb/elbv1_cert_provided_by_acm.sql

\set check_id 'ELB.3'
\echo "Executing check ELB.3"
\i queries/elb/elbv1_https_or_tls.sql

\set check_id 'ELB.4'
\echo "Executing check ELB.4"
\i queries/elb/alb_drop_http_headers.sql

\set check_id 'ELB.5'
\echo "Executing check ELB.5"
\i queries/elb/alb_logging_enabled.sql

\set check_id 'ELB.6'
\echo "Executing check ELB.6"
\i queries/elb/alb_deletion_protection_enabled.sql

\set check_id 'ELB.7'
\echo "Executing check ELB.7"
\i queries/elb/elbv1_conn_draining_enabled.sql

\set check_id 'ELB.8'
\echo "Executing check ELB.8"
\i queries/elb/elbv1_https_predefined_policy.sql
