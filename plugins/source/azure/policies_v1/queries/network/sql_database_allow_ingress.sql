insert into azure_policy_results
SELECT
  :'execution_time'
  :'framework',
  :'check_id',
  '',
  ass.subscription_id,
  ass.id AS server_id,
  case
    when assfr.start_ip_address = '0.0.0.0'
      OR ( assfr.start_ip_address = '255.255.255.255'
      AND assfr.end_ip_address = '0.0.0.0' )
    then 'fail' else 'pass'
  end
FROM azure_sql_servers ass
    LEFT JOIN
        azure_sql_server_firewall_rules assfr ON
            ass.cq_id = assfr.server_cq_id
