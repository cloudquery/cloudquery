insert into aws_policy_results
select
    :'execution_time' as execution_time,
    :'framework' as framework,
    :'check_id' as check_id,
    'Ensure CloudTrail log file validation is enabled' as title,
    account_id,
    arn as resource_id,
    case
      when bool_or(log_file_validation_enabled) then 'pass'
      else 'fail'
    end as status
from aws_cloudtrail_trails
group by account_id, arn
