\echo  "Executing CIS V1.5.0 Section 5"
\echo "Creating view_aws_security_group_ingress_rules"
\ir ../views/security_group_ingress_rules.sql
\echo "Creating view_aws_nacl_allow_ingress_rules"
\ir ../views/networks_acls_ingress_rules.sql
\set check_id '5.1'
\echo "Executing check 5.1"
\ir ../queries/ec2/no_broad_public_ingress_acl_on_port_22_3389.sql
\set check_id '5.2'
\echo "Executing check 5.2"
\ir ../queries/ec2/no_broad_public_ipv4_ingress_on_port_22_3389.sql
\set check_id '5.3'
\ir ../queries/ec2/no_broad_public_ipv6_ingress_on_port_22_3389.sql
\set check_id '5.4'
\echo "Executing check 5.4"
\ir ../queries/ec2/default_sg_no_access.sql
\set check_id '5.5'
    -- manual
