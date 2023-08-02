insert into aws_policy_results
select
    :'execution_time' as execution_time,
    :'framework' as framework,
    :'check_id' as check_id,
    'CloudTrail trails should be integrated with CloudWatch Logs' as title,
    account_id,
    arn as resource_id,
    case
        when bool_or(
            cloud_watch_logs_log_group_arn is not null
            and (status->>'LatestCloudWatchLogsDeliveryTime')::timestamp - now() < '1 days'::INTERVAL
        )
        then 'pass'
        else 'fail'
    end as status
from aws_cloudtrail_trails
group by account_id, arn
