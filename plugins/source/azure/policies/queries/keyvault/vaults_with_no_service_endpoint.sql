WITH secured_vaults AS (SELECT v.cq_id, subnet_id
                        FROM azure_keyvault_vaults v,
                             UNNEST(v.network_acls_virtual_network_rules) AS subnet_id
                                 LEFT JOIN azure_network_virtual_network_subnets sb
                                           ON subnet_id = sb.id
                        WHERE v.network_acls_default_action = 'Deny'
                          AND sb.provisioning_state = 'Succeeded')
insert into azure_policy_results
SELECT
  :'execution_time',
  :'framework',
  :'check_id',
  'Key Vault should use a virtual network service endpoint',
  subscription_id,
  id,
  case
    when sv.cq_id IS NULL then 'fail' else 'pass'
  end
FROM azure_keyvault_vaults v
  LEFT JOIN secured_vaults sv ON v.cq_id = sv.cq_id