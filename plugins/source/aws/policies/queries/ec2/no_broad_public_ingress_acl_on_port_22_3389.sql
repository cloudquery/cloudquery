-- uses view which uses no_broad_public_ingress_acl_on_port_22_3389.sql query
insert into aws_policy_results
select
  :'execution_time',
  :'framework',
  :'check_id',
  'Ensure no Network ACLs allow ingress from 0.0.0.0/0 to remote server administration ports (Automated)',
  account_id,
  arn,
  case
    when bool_or(
      rule_action = 'allow'
      and (cidr_block = '0.0.0.0/0' or ipv6_cidr_block = '::/0')
      and (
        (port_range_from is null and port_range_to is null) -- all ports
        or 22 between port_range_from and port_range_to
        or 3389 between port_range_from and port_range_to
      )
    )
    then 'fail'
    else 'pass'
  end
from view_aws_nacl_ingress_rules
group by account_id, arn
