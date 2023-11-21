create or replace view view_aws_log_metric_filter_and_alarm as
with af as (
  select distinct a.arn, a.actions_enabled, a.alarm_actions, m->'MetricStat'->'Metric'->>'MetricName' as metric_name -- TODO check
  from aws_cloudwatch_alarms a, jsonb_array_elements(a.metrics) as m
),
tes as (
  select trail_arn from aws_cloudtrail_trail_event_selectors
  where exists(
    select * from jsonb_array_elements(event_selectors) as es
    where es ->>'ReadWriteType' = 'All' and (es->>'IncludeManagementEvents')::boolean = TRUE
  ) or exists(
    select * from jsonb_array_elements(advanced_event_selectors) as aes
    where not exists(select * from jsonb_array_elements(aes ->'FieldSelectors') as aes_fs where aes_fs ->>'Field' = 'readOnly')
  )
)
select
    t.account_id,
    t.region,
    t.cloud_watch_logs_log_group_arn,
    mf.filter_pattern as pattern
from aws_cloudtrail_trails t
inner join tes on t.arn = tes.trail_arn
inner join aws_cloudwatchlogs_metric_filters mf on mf.log_group_name = t.cloudwatch_logs_log_group_name
inner join af on mf.filter_name = af.metric_name
inner join aws_sns_subscriptions ss on ss.topic_arn = ANY(af.alarm_actions)
where t.is_multi_region_trail = TRUE
    and (t.status->>'IsLogging')::boolean = TRUE
    and ss.arn like 'aws:arn:%'
