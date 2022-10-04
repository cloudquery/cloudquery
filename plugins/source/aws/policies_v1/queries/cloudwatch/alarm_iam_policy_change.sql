insert into aws_policy_results
WITH matching_alarms AS (
    SELECT account_id, metric_name AS metric_name
    FROM view_aws_metric_filters_with_active_subscriptions
    WHERE pattern NOT LIKE '%NOT%'
    AND pattern LIKE '%($.eventName = DeleteGroupPolicy)%'
    AND pattern LIKE '%($.eventName = DeleteUserPolicy)%'
    AND pattern LIKE '%($.eventName = PutGroupPolicy)%'
    AND pattern LIKE '%($.eventName = PutRolePolicy)%'
    AND pattern LIKE '%($.eventName = PutUserPolicy)%'
    AND pattern LIKE '%($.eventName = CreatePolicy)%'
    AND pattern LIKE '%($.eventName = DeletePolicy)%'
    AND pattern LIKE '%($.eventName=CreatePolicyVersion)%'
    AND pattern LIKE '%($.eventName=DeletePolicyVersion)%'
    AND pattern LIKE '%($.eventName=AttachRolePolicy)%'
    AND pattern LIKE '%($.eventName=DetachRolePolicy)%'
    AND pattern LIKE '%($.eventName=AttachUserPolicy)%'
    AND pattern LIKE '%($.eventName = DetachUserPolicy)%'
    AND pattern LIKE '%($.eventName = AttachGroupPolicy)%'
    AND pattern LIKE '%($.eventName = DetachGroupPolicy)%')
SELECT  :'execution_time' as execution_time,
        :'framework' as framework,
        :'check_id' as check_id,
        'Ensure a log metric filter and alarm exist for IAM policy changes (Score)' as title,
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

