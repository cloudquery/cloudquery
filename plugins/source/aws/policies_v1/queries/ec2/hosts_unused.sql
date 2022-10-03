insert into aws_policy_results
select :'execution_time'       as execution_time,
       :'framework'            as framework,
       :'check_id'             as check_id,
       'Unused dedicated host' as title,
       account_id,
       arn                     as resource_id,
       'fail'                  as status
from aws_ec2_hosts
where coalesce(jsonb_array_length(instances), 0) = 0;
