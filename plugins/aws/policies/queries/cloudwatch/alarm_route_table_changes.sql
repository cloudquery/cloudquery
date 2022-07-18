insert into aws_policy_results
select
  :execution_time as execution_time,
  :'framework' as framework,
  :'check_id' as check_id,
  'Ensure a log metric filter and alarm exist for route table changes (Scored)' as title,
  account_id,
  cloud_watch_logs_log_group_arn,
    case when
        pattern = '{ ($.eventName = CreateRoute) '
        || '|| ($.eventName = CreateRouteTable) '
        || '|| ($.eventName = ReplaceRoute) '
        || '|| ($.eventName = ReplaceRouteTableAssociation) '
        || '|| ($.eventName = DeleteRouteTable) '
        || '|| ($.eventName = DeleteRoute) '
        || '|| ($.eventName = DisassociateRouteTable) }' then 'pass'
        else 'fail'
    end as status
from view_aws_log_metric_filter_and_alarm
