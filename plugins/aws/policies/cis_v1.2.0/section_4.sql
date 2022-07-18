\echo  "Executing CIS V1.2.0 Section 4"
\echo "Creating view_aws_security_group_ingress_rules"
\i views/security_group_ingress_rules.sql
\set check_id "4.1"
\echo "Executing check 4.1"
\i queries/ec2/no_broad_public_ingress_on_port_22.sql
\set check_id "4.2"
\echo "Executing check 4.2"
\i queries/ec2/no_broad_public_ingress_on_port_3389.sql
\set check_id "4.3"
\echo "Executing check 4.3"
\i queries/ec2/default_sg_no_access.sql
