ALTER TABLE IF EXISTS aws_ec2_security_group_ip_permissions ADD COLUMN "permission_type" text;


ALTER TABLE IF EXISTS public.aws_ec2_security_group_ip_permission_ip_ranges DROP CONSTRAINT aws_ec2_security_group_ip_permission_ip_ranges_pk;


ALTER TABLE IF EXISTS aws_ec2_security_group_ip_permission_ip_ranges RENAME COLUMN "cidr_ip" TO "cidr";


ALTER TABLE IF EXISTS aws_ec2_security_group_ip_permission_ip_ranges ADD COLUMN "cidr_type" text;


DROP TABLE IF EXISTS public.aws_ec2_security_group_ip_permission_ipv6_ranges;
DROP TABLE IF EXISTS public.aws_ec2_security_group_ip_permissions_egress_ip_ranges;
DROP TABLE IF EXISTS public.aws_ec2_security_group_ip_permissions_egress_ipv6_ranges;
DROP TABLE IF EXISTS public.aws_ec2_security_group_ip_permissions_egress_prefix_list_ids;
DROP TABLE IF EXISTS public.aws_ec2_security_group_ip_permissions_egress_user_group_pairs;
DROP TABLE IF EXISTS public.aws_ec2_security_group_ip_permissions_egresses;
