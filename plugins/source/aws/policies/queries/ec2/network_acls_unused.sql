insert into aws_policy_results
select :'execution_time'                    as execution_time,
       :'framework'                         as framework,
       :'check_id'                          as check_id,
       'Unused network access control list' as title,
       account_id,
       arn                                  as resource_id,
       'fail'                               as status
from aws_ec2_network_acls
where coalesce(jsonb_array_length(associations), 0) = 0;
