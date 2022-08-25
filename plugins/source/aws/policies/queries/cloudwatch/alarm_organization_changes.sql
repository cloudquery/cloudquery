insert into aws_policy_results
select
  :'execution_time' as execution_time,
  :'framework' as framework,
  :'check_id' as check_id,
  'Ensure a log metric filter and alarm exists for AWS Organizations changes (Automated)' as title,
  account_id,
  cloud_watch_logs_log_group_arn as resource_id,
  case when
      pattern = '{ ($.eventName = AcceptHandshake) '
      || '|| ($.eventName = AttachPolicy) '
      || '|| ($.eventName = CreateAccount) '
      || '|| ($.eventName = CreateOrganizationalUnit) '
      || '|| ($.eventName = CreatePolicy) '
      || '|| ($.eventName = DeclineHandshake) '
      || '|| ($.eventName = DeleteOrganization) '
      || '|| ($.eventName = DeleteOrganizationalUnit) '
      || '|| ($.eventName = DeletePolicy) '
      || '|| ($.eventName = DetachPolicy) '
      || '|| ($.eventName = DisablePolicyType) '
      || '|| ($.eventName = EnablePolicyType) '
      || '|| ($.eventName = InviteAccountToOrganization) '
      || '|| ($.eventName = LeaveOrganization) '
      || '|| ($.eventName = MoveAccount) '
      || '|| ($.eventName = RemoveAccountFromOrganization) '
      || '|| ($.eventName = UpdatePolicy) '
      || '|| ($.eventName = UpdateOrganizationalUnit) }'
      then 'pass'
      else 'fail'
  end as status
from view_aws_log_metric_filter_and_alarm
