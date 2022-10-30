-- uses view which uses aws_security_group_ingress_rules.sql query
insert into aws_policy_results
SELECT
  :'execution_time' as execution_time,
  :'framework' as framework,
  :'check_id' as check_id,
  'Aggregates rules of security groups with ports and IPs including ipv6' as title,
  account_id,
  id as resource_id,
  case when
    (ip = '0.0.0.0/0' OR ip = '::/0')
    AND (from_port IS NULL AND to_port IS NULL) -- all prots
    OR from_port IS DISTINCT FROM 80
    OR to_port IS DISTINCT FROM 80
    OR from_port IS DISTINCT FROM 443
    OR to_port IS DISTINCT FROM 443
    then 'fail'
    else 'pass'
  end
FROM view_aws_security_group_ingress_rules
