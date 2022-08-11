-- TODO FIXME
/* Find Security groups that give access to ipv4 addresses */
select t.arn
from (
    select
        /* create arn for sg */
        arn,
        /* Calculate total number of IPs a SG rule gives access to */
        (
            SPLIT_PART(
                HOST(BROADCAST(cidr::CIDR)), '.', 1
            )::BIGINT * 16777216
            + SPLIT_PART(HOST(BROADCAST(cidr::CIDR)), '.', 2)::BIGINT * 65536
            + SPLIT_PART(HOST(BROADCAST(cidr::CIDR)), '.', 3)::BIGINT * 256
            + SPLIT_PART(HOST(BROADCAST(cidr::CIDR)), '.', 4)::BIGINT
        ) - (
            SPLIT_PART(HOST(cidr::CIDR), '.', 1)::BIGINT * 16777216
            + SPLIT_PART(HOST(cidr::CIDR), '.', 2)::BIGINT * 65536
            + SPLIT_PART(HOST(cidr::CIDR), '.', 3)::BIGINT * 256
            + SPLIT_PART(HOST(cidr::CIDR), '.', 4)::BIGINT
        ) as totalips
    from aws_ec2_security_groups
    inner join
        aws_ec2_security_group_ip_permissions on
            aws_ec2_security_groups.cq_id
            = aws_ec2_security_group_ip_permissions.security_group_cq_id
    inner join
        aws_ec2_security_group_ip_permission_ip_ranges on
            aws_ec2_security_group_ip_permissions.cq_id
            = aws_ec2_security_group_ip_permission_ip_ranges.security_group_ip_permission_cq_id
            and aws_ec2_security_group_ip_permission_ip_ranges.cidr_type='ipv4'
    where (
            (
                from_port is null
                and to_port is null
            )
            or 22 between from_port
            and to_port
        )
) as t
group by t.arn
having SUM(t.totalips) = 4294967295
/* this value is the total number of ips in ipv4 space ie 0.0.0.0/0 */
union
/* Find Security groups that give access to ipv6 addresses */
select t.arn
from (
        select
            /* create arn for sg */
            arn,
            /* Calculate total number of IPs a SG rule gives access to */
            ROUND(
                2 ^ (
                    128 - MASKLEN(
                        aws_ec2_security_group_ip_permission_ip_ranges.cidr::CIDR
                    )
                )::NUMERIC
            ) as totalips
        from aws_ec2_security_groups
        inner join
            aws_ec2_security_group_ip_permissions on
                aws_ec2_security_groups.cq_id
                = aws_ec2_security_group_ip_permissions.security_group_cq_id
        inner join
            aws_ec2_security_group_ip_permission_ip_ranges on
                aws_ec2_security_group_ip_permissions.cq_id
                = aws_ec2_security_group_ip_permission_ip_ranges.security_group_ip_permission_cq_id
            and aws_ec2_security_group_ip_permission_ip_ranges.cidr_type='ipv6'
        where (
                (
                    from_port is null
                    and to_port is null
                )
                or 22 between from_port
                and to_port
            )
    ) as t
group by t.arn
having SUM(t.totalips) = 340282366920938463463374607431768211456;

/* this value is the total number of ips in ipv6 space ie ::/0 */
