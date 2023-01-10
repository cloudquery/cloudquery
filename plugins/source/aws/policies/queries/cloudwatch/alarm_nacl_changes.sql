insert into aws_policy_results
select
  :'execution_time' as execution_time,
  :'framework' as framework,
  :'check_id' as check_id,
  'Ensure a log metric filter and alarm exist for changes to Network Access Control Lists (NACL) (Scored)' as title,
  account_id,
  cloud_watch_logs_log_group_arn as resource_id,
  case
      when pattern NOT LIKE '%NOT%'
          AND pattern LIKE '%($.eventName = CreateNetworkAcl)%'
          AND pattern LIKE '%($.eventName = CreateNetworkAclEntry)%'
          AND pattern LIKE '%($.eventName = DeleteNetworkAcl)%'
          AND pattern LIKE '%($.eventName = DeleteNetworkAclEntry)%'
          AND pattern LIKE '%($.eventName = ReplaceNetworkAclAssociation)%'
          AND pattern LIKE '%($.eventName = ReplaceNetworkAclEntry)%' then 'pass'
      else 'fail'
  end as status
from view_aws_log_metric_filter_and_alarm
