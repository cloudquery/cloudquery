insert into azure_policy_results
SELECT
  :'execution_time' as execution_time,
  :'framework' as framework,
  :'check_id' as check_id,
  'Ensure that VA setting "Also send email notifications to admins and subscription owners" is set for a SQL server (Automated)' as title,
  s.subscription_id,
  s.id AS server_id,
  case
    when (a.recurring_scans->>'emailSubscriptionAdmins')::boolean IS DISTINCT FROM TRUE
      then 'fail' else 'pass'
  end
FROM azure_sql_servers s
    LEFT JOIN azure_sql_server_vulnerability_assessments a ON
        s.id = a.sql_server_id
