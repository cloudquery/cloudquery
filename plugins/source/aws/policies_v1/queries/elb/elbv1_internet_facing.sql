insert into aws_policy_results
select
    :'execution_time' as execution_time,
    :'framework' as framework,
    :'check_id' as check_id,
    'Find all Classic ELBs that are Internet Facing' AS title,
    account_id,
    arn as resource_id,
    case when scheme = 'internet-facing' then 'fail' else 'pass' end as status
from
    aws_elbv1_load_balancers
