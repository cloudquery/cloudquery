-- describes metric alarms and pattern of alarm filter that has active subscriptions
CREATE OR REPLACE VIEW view_aws_metric_filters_with_active_subscriptions AS
WITH filter_metric_name AS (SELECT mf.account_id,
                                   filter_pattern,
                                   jsonb_array_elements(metric_transformations) ->> 'MetricName' AS metric_name
                            FROM aws_cloudwatchlogs_metric_filters mf),
     alarms_actions_metric AS (SELECT arn, UNNEST(alarm_actions) AS topic_arn, fmn.metric_name, fmn.filter_pattern
                               FROM aws_cloudwatch_alarms aca
                                        INNER JOIN filter_metric_name fmn ON
                                           fmn.metric_name = aca.metric_name
                                       AND aca.account_id = fmn.account_id)
SELECT account_id, metric_name, filter_pattern
FROM aws_sns_topics st
         LEFT JOIN alarms_actions_metric aam ON
    aam.topic_arn = st.arn
WHERE st.subscriptions_confirmed > 0;