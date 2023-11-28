insert into aws_policy_results
select :'execution_time'                  as execution_time,
       :'framework'                       as framework,
       :'check_id'                        as check_id,
       'Disabled Lightsail distributions' as title,
       account_id,
       arn                                as resource_id,
       'fail'                             as status
from aws_lightsail_distributions
where is_enabled = false