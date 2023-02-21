insert into azure_policy_results
SELECT
  :'execution_time' as execution_time,
  :'framework' as framework,
  :'check_id' as check_id,
  'Ensure that Advanced Threat Protection (ATP) on a SQL server is set to "Enabled" (Automated)' as title,
  s.subscription_id,
  s.id AS server_id,
  case
    when p.properties->>'state' is distinct from 'Enabled'
    then 'fail' else 'pass'
  end
FROM azure_sql_servers s
    LEFT JOIN azure_sql_server_databases d ON
        s._cq_id = d._cq_parent_id
    LEFT JOIN azure_sql_server_database_threat_protections p ON
        d._cq_id = p._cq_parent_id
