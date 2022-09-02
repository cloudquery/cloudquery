-- describes metric alarms and pattern of alarm filter that has active subscriptions
create or replace view view_aws_metric_filters_with_active_subscriptions as
WITH filter_metric_name AS (SELECT account_id, pattern, acfmt.metric_name
        FROM aws_cloudwatchlogs_filters acf
                 INNER JOIN aws_cloudwatchlogs_filter_metric_transformations acfmt ON
            acfmt.filter_cq_id = acf.cq_id),
     alarms_actions_metric AS (SELECT arn, UNNEST(actions) AS topic_arn, fmn.metric_name, fmn.pattern
       FROM aws_cloudwatch_alarms aca
                INNER JOIN filter_metric_name fmn ON
                   fmn.metric_name = aca.metric_name
               AND aca.account_id = fmn.account_id)
SELECT account_id, metric_name, pattern
FROM aws_sns_topics st
         LEFT JOIN alarms_actions_metric aam ON
    aam.topic_arn = st.arn
WHERE st.subscriptions_confirmed > 0;