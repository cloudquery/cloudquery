insert into azure_policy_results
SELECT 
  :'execution_time' as execution_time,
  :'framework' as framework,
  :'check_id' as check_id,
  'Ensure that VA setting Periodic Recurring Scans is enabled on a SQL server (Automated)' as title,
  s.subscription_id,
  s.id,
  case
    when (a.properties->'recurringScans'->>'isEnabled')::boolean IS DISTINCT FROM TRUE
      then 'fail' else 'pass'
  end
FROM azure_sql_servers s
    LEFT JOIN azure_sql_server_vulnerability_assessments a ON
        s._cq_id = a._cq_parent_id
