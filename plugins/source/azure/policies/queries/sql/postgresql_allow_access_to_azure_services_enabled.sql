insert into azure_policy_results
SELECT
  :'execution_time' as execution_time,
  :'framework' as framework,
  :'check_id' as check_id,
  'Ensure "Allow access to Azure services" for PostgreSQL Database Server is disabled (Automated)' as title,
  aps.subscription_id,
  aps.id AS server_id,
  case
    when apsfr."name" = 'AllowAllAzureIps'
      OR (apsfr.properties->>'startIPAddress' = '0.0.0.0'
      AND apsfr.properties->>'endIPAddress' = '0.0.0.0')
    then 'fail' else 'pass'
  end
FROM azure_postgresql_servers aps
    LEFT JOIN azure_postgresql_server_firewall_rules apsfr ON
        aps._cq_id = apsfr._cq_parent_id
