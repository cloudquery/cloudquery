insert into aws_policy_results
select
  :execution_time as execution_time,
  :'framework' as framework,
  :'check_id' as check_id,
  'Ensure a log metric filter and alarm exist for S3 bucket policy changes (Scored)' as title,
  account_id,
  cloud_watch_logs_log_group_arn as resource_id,
  case 
    when pattern = '{ ($.eventSource = s3.amazonaws.com) '
      || '&& (($.eventName = PutBucketAcl) '
      || '|| ($.eventName = PutBucketPolicy) '
      || '|| ($.eventName = PutBucketCors) '
      || '|| ($.eventName = PutBucketLifecycle) '
      || '|| ($.eventName = PutBucketReplication) '
      || '|| ($.eventName = DeleteBucketPolicy) '
      || '|| ($.eventName = DeleteBucketCors) '
      || '|| ($.eventName = DeleteBucketLifecycle) '
      || '|| ($.eventName = DeleteBucketReplication)) }'
    then 'fail'
  end as status
from view_aws_log_metric_filter_and_alarm
