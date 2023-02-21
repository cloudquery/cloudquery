insert into azure_policy_results
SELECT
  :'execution_time' as execution_time,
  :'framework' as framework,
  :'check_id' as check_id,
  'Ensure that Vulnerability Assessment (VA) is enabled on a SQL server by setting a Storage Account (Automated)' as title,
  s.subscription_id,
  s.id AS server_id,
  case
    when a.properties->>'storageContainerPath' IS NULL OR a.properties->>'storageContainerPath' = ''
      then 'fail' else 'pass'
  end
FROM azure_sql_servers s
    LEFT JOIN azure_sql_server_vulnerability_assessments a ON
        s._cq_id = a._cq_parent_id
