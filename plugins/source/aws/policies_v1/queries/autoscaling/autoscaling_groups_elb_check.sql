insert into aws_policy_results
select
    :'execution_time' as execution_time,
    :'framework' as framework,
    :'check_id' as check_id,
    'Auto Scaling groups associated with a load balancer should use health checks' as title,
    account_id,
    arn as resource_id,
    case
        when ARRAY_LENGTH(load_balancer_names, 1) > 0 and health_check_type is distinct from 'ELB' then 'fail'
        else 'pass'
    end as status
from aws_autoscaling_groups
