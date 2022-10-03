insert into aws_policy_results
select :'execution_time'        as execution_time,
       :'framework'             as framework,
       :'check_id'              as check_id,
       'Unused ACM certificate' as title,
       account_id,
       arn                      as resource_id,
       'fail'                   as status
from aws_acm_certificates
where array_length(in_use_by, 1) = 0