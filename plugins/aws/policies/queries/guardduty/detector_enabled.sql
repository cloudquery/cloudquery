insert into aws_policy_results
with enabled_detector_regions as (
    select region
    from aws_guardduty_detectors
    where status = 'ENABLED'
)

select
    :execution_time as execution_time,
    :'framework' as framework,
    :'check_id' as check_id,
    'GuardDuty should be enabled' AS title,
    r.account_id,
    r.region AS resource_id,
    case when
        enabled = TRUE and e.region is null
    then 'fail' else 'pass' end AS status
from aws_regions r
left join enabled_detector_regions e on e.region = r.region
union
-- Add any detector that is enabled but all data sources are disabled
select
    :execution_time as execution_time,
    :'framework' as framework,
    :'check_id' as check_id,
    'GuardDuty should be enabled (detectors)' AS title,
    account_id,
    region AS resource_id,
    case when
        data_sources_s3_logs_status != 'ENABLED' and data_sources_cloud_trail_status != 'ENABLED'
        and data_sources_dns_logs_status != 'ENABLED' and data_sources_flow_logs_status != 'ENABLED'
    then 'fail' else 'pass' end AS status
from aws_guardduty_detectors
where
    status = 'ENABLED'
