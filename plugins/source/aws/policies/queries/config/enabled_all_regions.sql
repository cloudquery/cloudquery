insert into aws_policy_results
select
    :'execution_time' as execution_time,
    :'framework' as framework,
    :'check_id' as check_id,
    'AWS Config should be enabled' as title,
    account_id,
    arn as resource_id,
    case when
      (recording_group->>'IncludeGlobalResourceTypes')::boolean IS NOT TRUE
      OR (recording_group->>'AllSupported')::boolean IS NOT TRUE
      OR status_recording IS NOT TRUE
      OR status_last_status IS DISTINCT FROM 'SUCCESS'
    then 'fail'
    else 'pass'
    end as status
FROM
    aws_config_configuration_recorders
