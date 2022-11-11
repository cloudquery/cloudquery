insert into azure_policy_results
SELECT
  :'execution_time',
  :'framework',
  :'check_id',
  'Ensure "Allow access to Azure services" for PostgreSQL Database Server is disabled (Automated)',
  aps.subscription_id,
  aps.id AS server_id,
  case
    when apsfr."name" = 'AllowAllAzureIps'
      OR (apsfr.start_ip_address = '0.0.0.0'
      AND apsfr.end_ip_address = '0.0.0.0')
    then 'fail' else 'pass'
  end
FROM azure_postgresql_servers aps
    LEFT JOIN azure_postgresql_firewall_rules apsfr ON
        aps.id = apsfr.postgresql_server_id
