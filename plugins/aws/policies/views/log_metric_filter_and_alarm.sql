create or replace view view_aws_log_metric_filter_and_alarm as
select
    aws_cloudtrail_trails.account_id,
    aws_cloudtrail_trails.region,
    cloud_watch_logs_log_group_arn,
    pattern
from aws_cloudtrail_trails
inner join aws_cloudtrail_trail_event_selectors
    on
        aws_cloudtrail_trails.cq_id
        = aws_cloudtrail_trail_event_selectors.trail_cq_id
inner join aws_cloudwatchlogs_filters
    on
        aws_cloudtrail_trails.cloudwatch_logs_log_group_name
        = aws_cloudwatchlogs_filters.log_group_name
inner join aws_cloudwatch_alarm_metrics
    on
        aws_cloudwatchlogs_filters.name
        = aws_cloudwatch_alarm_metrics.metric_stat_metric_name
inner join
    aws_cloudwatch_alarms on
        aws_cloudwatch_alarm_metrics.alarm_cq_id
        = aws_cloudwatch_alarms.cq_id
inner join
    aws_sns_subscriptions on
        aws_sns_subscriptions.topic_arn
        = ANY(aws_cloudwatch_alarms.actions)
where is_multi_region_trail = TRUE
    and is_logging = TRUE
    and include_management_events = TRUE
    and read_write_type = 'All'
    and aws_sns_subscriptions.arn like 'aws:arn:%'
