insert into aws_policy_results
select
  :'execution_time' as execution_time,
  :'framework' as framework,
  :'check_id' as check_id,
  'Elasticsearch domains should have audit logging enabled' as title,
  account_id,
  arn as resource_id,
  case when
    log_publishing_options -> 'AUDIT_LOGS' -> 'Enabled' is distinct from 'true'
    or log_publishing_options -> 'AUDIT_LOGS' -> 'CloudWatchLogsLogGroupArn' is null
    then 'fail'
    else 'pass'
  end as status
from aws_elasticsearch_domains
