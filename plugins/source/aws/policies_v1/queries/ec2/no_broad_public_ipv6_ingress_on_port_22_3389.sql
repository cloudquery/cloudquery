-- uses view which uses aws_security_group_ingress_rules.sql query
insert into aws_policy_results
select
  :'execution_time',
  :'framework',
  :'check_id',
  'Ensure no security groups allow ingress from ::/0 to remote server administration ports (Automated)',
  account_id,
  arn,
  case when
      (ip = '::/0')
      and (
          (from_port is null and to_port is null) -- all ports
          or 22 between from_port and to_port
          or 3389 between from_port and to_port)
      then 'fail'
      else 'pass'
  end
from view_aws_security_group_ingress_rules
