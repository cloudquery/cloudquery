insert into aws_policy_results
WITH matching_alarms AS (
    SELECT account_id, metric_name AS metric_name
    FROM view_aws_metric_filters_with_active_subscriptions
    WHERE pattern NOT LIKE '%NOT%'
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
      AND pattern LIKE '%($.eventName = UpdateOrganizationalUnit)%')
SELECT  :'execution_time' as execution_time,
        :'framework' as framework,
        :'check_id' as check_id,
        'Ensure a log metric filter and alarm exists for AWS Organizations changes (Automated)' as title,
        a.account_id,
        case when
                     count(ma.metric_name) = 0
                 then 'fail'
             else 'pass'
            end as status
FROM aws_accounts a
         LEFT JOIN matching_alarms ma ON
        a.account_id = ma.account_id
GROUP BY a.account_id, ma.metric_name;