WITH subs AS (
    SELECT subscription_id, jsonb_array_elements(subnets) AS subnet, provisioning_state
    FROM azure_network_virtual_networks
), secured_servers AS (SELECT s._cq_id
                         FROM azure_sql_servers s
                                  LEFT JOIN azure_sql_virtual_network_rules r
                                            ON s.id = r.sql_server_id
                                  LEFT JOIN subs
                                            ON r.virtual_network_subnet_id = subs.subnet->>'id'
                         WHERE r.virtual_network_subnet_id IS NOT NULL
                           AND subs.provisioning_state = 'Succeeded')
insert into azure_policy_results
SELECT
  :'execution_time' as execution_time,
  :'framework' as framework,
  :'check_id' as check_id,
  'SQL Server should use a virtual network service endpoint' as title,
  subscription_id,
  id,
  case
    when ss._cq_id IS NULL
      then 'fail' else 'pass'
  end
FROM azure_sql_servers s
     LEFT JOIN secured_servers ss ON s._cq_id = ss._cq_id