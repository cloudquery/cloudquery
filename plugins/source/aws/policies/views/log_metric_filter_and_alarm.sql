create or replace view view_aws_log_metric_filter_and_alarm as
with af as (
  select distinct a.arn, a.actions_enabled, a.alarm_actions, m->'MetricStat'->'Metric'->>'MetricName' as metric_name -- TODO check
  from aws_cloudwatch_alarms a, jsonb_array_elements(a.metrics) as m
)
select
    t.account_id,
    t.region,
    t.cloud_watch_logs_log_group_arn,
    mf.filter_pattern as pattern
from aws_cloudtrail_trails t
inner join aws_cloudtrail_trail_event_selectors tes on t.arn = tes.trail_arn
inner join aws_cloudwatchlogs_metric_filters mf on mf.log_group_name = t.cloudwatch_logs_log_group_name
inner join af on mf.filter_name = af.metric_name
inner join aws_sns_subscriptions ss on ss.topic_arn = ANY(af.alarm_actions)
where t.is_multi_region_trail = TRUE
    and (t.status->>'IsLogging')::boolean = TRUE
    and tes.include_management_events = TRUE
    and tes.read_write_type = 'All'
    and ss.arn like 'aws:arn:%'
