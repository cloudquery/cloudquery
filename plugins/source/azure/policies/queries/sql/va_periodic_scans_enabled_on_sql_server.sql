insert into azure_policy_results
SELECT 
  :'execution_time',
  :'framework',
  :'check_id',
  'Ensure that VA setting Periodic Recurring Scans is enabled on a SQL server (Automated)',
  s.subscription_id,
  s.id,
  case
    when a.recurring_scans_is_enabled IS NULL OR a.recurring_scans_is_enabled != TRUE
      then 'fail' else 'pass'
  end
FROM azure_sql_servers s
    LEFT JOIN azure_sql_server_vulnerability_assessments a ON
        s.cq_id = a.server_cq_id
