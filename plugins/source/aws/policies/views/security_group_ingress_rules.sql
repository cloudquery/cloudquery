create or replace view view_aws_security_group_ingress_rules as
with sg_rules_ports as (
    select
        aws_ec2_security_groups.account_id,
        aws_ec2_security_groups.region,
        aws_ec2_security_groups.group_name,
        aws_ec2_security_groups.arn,
        aws_ec2_security_groups.id,
        aws_ec2_security_group_ip_permissions.from_port,
        aws_ec2_security_group_ip_permissions.to_port,
        aws_ec2_security_group_ip_permissions.ip_protocol,
        aws_ec2_security_group_ip_permissions.cq_id as permission_id
    from aws_ec2_security_groups
    left join
        aws_ec2_security_group_ip_permissions on
            aws_ec2_security_groups.cq_id = aws_ec2_security_group_ip_permissions.security_group_cq_id
    where aws_ec2_security_group_ip_permissions.permission_type = 'ingress'
)

select
    sg_rules_ports.*,
    aws_ec2_security_group_ip_permission_ip_ranges.cidr as ip
from sg_rules_ports
left join
    aws_ec2_security_group_ip_permission_ip_ranges on
        sg_rules_ports.permission_id = aws_ec2_security_group_ip_permission_ip_ranges.security_group_ip_permission_cq_id
