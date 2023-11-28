insert into aws_policy_results
select
    :'execution_time' as execution_time,
    :'framework' as framework,
    :'check_id' as check_id,
    'Ensure a log metric filter and alarm exist for usage of "root" account (Score)' as title,
    account_id,
    cloud_watch_logs_log_group_arn as resource_id,
    case
        when pattern NOT LIKE '%NOT%'
            AND pattern LIKE '%($.eventName = AcceptHandshake)%'
            AND pattern LIKE '%($.eventName = AttachPolicy)%'
            AND pattern LIKE '%($.eventName = CreateAccount)%'
            AND pattern LIKE '%($.eventName = CreateOrganizationalUnit)%'
            AND pattern LIKE '%($.eventName = CreatePolicy)%'
            AND pattern LIKE '%($.eventName = DeclineHandshake)%'
            AND pattern LIKE '%($.eventName = DeleteOrganization)%'
            AND pattern LIKE '%($.eventName = DeleteOrganizationalUnit)%'
            AND pattern LIKE '%($.eventName = DeletePolicy)%'
            AND pattern LIKE '%($.eventName = DetachPolicy)%'
            AND pattern LIKE '%($.eventName = DisablePolicyType)%'
            AND pattern LIKE '%($.eventName = EnablePolicyType)%'
            AND pattern LIKE '%($.eventName = InviteAccountToOrganization)%'
            AND pattern LIKE '%($.eventName = LeaveOrganization)%'
            AND pattern LIKE '%($.eventName = MoveAccount)%'
            AND pattern LIKE '%($.eventName = RemoveAccountFromOrganization)%'
            AND pattern LIKE '%($.eventName = UpdatePolicy)%'
            AND pattern LIKE '%($.eventName = UpdateOrganizationalUnit)%' then 'pass'
        else 'fail'
        end as title
from view_aws_log_metric_filter_and_alarm