insert into aws_policy_results
select :'execution_time'                 as execution_time,
       :'framework'                      as framework,
       :'check_id'                       as check_id,
       'Unused Lightsail load balancers' as title,
       account_id,
       arn                               as resource_id,
       'fail'                            as status
from aws_lightsail_load_balancers
where coalesce(jsonb_array_length(instance_health_summary), 0) = 0;
