insert into aws_policy_results
select
  :'execution_time' as execution_time,
  :'framework' as framework,
  :'check_id' as check_id,
  'Ensure a log metric filter and alarm exist for changes to Network Access Control Lists (NACL) (Scored)' as title,
  account_id,
  cloud_watch_logs_log_group_arn as resource_id,
  case
      when pattern = '{ ($.eventName = CreateNetworkAcl) '
          || '|| ($.eventName = CreateNetworkAclEntry) '
          || '|| ($.eventName = DeleteNetworkAcl) '
          || '|| ($.eventName = DeleteNetworkAclEntry) '
          || '|| ($.eventName = ReplaceNetworkAclEntry) '
          || '|| ($.eventName = ReplaceNetworkAclAssociation) }' then 'pass'
      else 'fail'
  end as status
from view_aws_log_metric_filter_and_alarm
