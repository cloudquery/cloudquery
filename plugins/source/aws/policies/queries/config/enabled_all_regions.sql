insert into aws_policy_results
select
    :'execution_time' as execution_time,
    :'framework' as framework,
    :'check_id' as check_id,
    'AWS Config should be enabled' as title,
    aws_regions.account_id,
    aws_regions.region as resource_id,
    case when
      aws_config_configuration_recorders.arn IS NULL
      OR (recording_group->>'IncludeGlobalResourceTypes')::boolean IS NOT TRUE
      OR (recording_group->>'AllSupported')::boolean IS NOT TRUE
      OR status_recording IS NOT TRUE
      OR status_last_status IS DISTINCT FROM 'SUCCESS'
    then 'fail'
    else 'pass'
    end as status
FROM
    aws_regions
    left join aws_config_configuration_recorders
      on aws_regions.region = aws_config_configuration_recorders.region
      and aws_regions.account_id = aws_config_configuration_recorders.account_id
