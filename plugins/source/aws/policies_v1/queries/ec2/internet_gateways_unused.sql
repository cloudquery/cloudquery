insert into aws_policy_results
select :'execution_time'         as execution_time,
       :'framework'              as framework,
       :'check_id'               as check_id,
       'Unused internet gateway' as title,
       account_id,
       arn                       as resource_id,
       'fail'                    as status
from aws_ec2_internet_gateways
where coalesce(jsonb_array_length(attachments), 0) = 0;
