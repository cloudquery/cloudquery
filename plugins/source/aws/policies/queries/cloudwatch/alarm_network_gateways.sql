insert into aws_policy_results
WITH matching_alarms AS (
    SELECT account_id, metric_name AS metric_name
    FROM view_aws_metric_filters_with_active_subscriptions
    WHERE pattern NOT LIKE '%NOT%'
    AND pattern LIKE '%($.eventName = CreateCustomerGateway)%'
    AND pattern LIKE '%($.eventName = DeleteCustomerGateway)%'
    AND pattern LIKE '%($.eventName = AttachInternetGateway)%'
    AND pattern LIKE '%($.eventName = CreateInternetGateway)%'
    AND pattern LIKE '%($.eventName = DeleteInternetGateway)%'
    AND pattern LIKE '%($.eventName = DetachInternetGateway)%')
SELECT  :'execution_time' as execution_time,
        :'framework' as framework,
        :'check_id' as check_id,
        'Ensure a log metric filter and alarm exist for changes to network gateways (Scored)' as title,
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
