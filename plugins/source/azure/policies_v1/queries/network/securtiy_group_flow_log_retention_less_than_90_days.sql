SELECT
  :'execution_time'
  :'framework',
  :'check_id',
  '',
  subscription_id,
  id,
  case
    when ansgfl.retention_policy_enabled != TRUE
      OR ansgfl.retention_policy_enabled IS NULL
      OR ansgfl.retention_policy_days < 90
      OR ansgfl.retention_policy_days IS NULL
    then 'fail' else 'pass'
  end
FROM azure_network_security_groups ansg
    LEFT JOIN azure_network_security_group_flow_logs ansgfl ON
        ansg.cq_id = ansgfl.security_group_cq_id
