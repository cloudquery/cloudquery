WITH secured_servers AS (SELECT s.cq_id
                         FROM azure_sql_servers s
                                  LEFT JOIN azure_sql_server_virtual_network_rules r
                                            ON s.cq_id = r.server_cq_id
                                  LEFT JOIN azure_network_virtual_network_subnets sb
                                            ON r.subnet_id = sb.id
                         WHERE r.subnet_id IS NOT NULL
                           AND sb.provisioning_state = 'Succeeded')
insert into azure_policy_results
SELECT
  :'execution_time',
  :'framework',
  :'check_id',
  'SQL Server should use a virtual network service endpoint',
  subscription_id,
  id,
  case
    when ss.cq_id IS NULL
      then 'fail' else 'pass'
  end
FROM azure_sql_servers s
         LEFT JOIN secured_servers ss ON s.cq_id = ss.cq_id