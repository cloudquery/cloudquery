insert into aws_policy_results
select :'execution_time'         as execution_time,
       :'framework'              as framework,
       :'check_id'               as check_id,
       'Unused ELB target group' as title,
       account_id,
       arn                       as resource_id,
       'fail'                    as status
from aws_elbv2_target_groups
where array_length(load_balancer_arns, 1) = 0