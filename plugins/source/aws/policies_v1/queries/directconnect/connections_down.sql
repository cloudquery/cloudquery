insert into aws_policy_results
select :'execution_time'                            as execution_time,
       :'framework'                                 as framework,
       :'check_id'                                  as check_id,
       'Direct Connect connections in "down" state' as title,
       account_id,
       arn                                          as resource_id,
       'fail'                                       as status
from aws_directconnect_connections
where connection_state = 'down'