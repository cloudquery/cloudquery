insert into aws_policy_results
SELECT
  :'execution_time' as execution_time,
  :'framework' as framework,
  :'check_id' as check_id,
  'Elasticsearch domain error logging to CloudWatch Logs should be enabled' as title,
  account_id,
  arn as resource_id,
  case when
    log_publishing_options -> 'ES_APPLICATION_LOGS' -> 'Enabled' IS DISTINCT FROM 'true'
    OR log_publishing_options -> 'ES_APPLICATION_LOGS' -> 'CloudWatchLogsLogGroupArn' IS NULL
    then 'fail'
    else 'pass'
  end as status
FROM aws_elasticsearch_domains
