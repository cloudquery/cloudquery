insert into aws_policy_results
select :'execution_time'                         as execution_time,
       :'framework'                              as framework,
       :'check_id'                               as check_id,
       'Direct Connect LAGs with no connections' as title,
       account_id,
       arn                                       as resource_id,
       'fail'                                    as status
from aws_directconnect_lags
where number_of_connections == 0
  or array_length(connection_ids, 1) == 0