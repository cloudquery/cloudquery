insert into azure_policy_results
SELECT
  :'execution_time' as execution_time,
  :'framework' as framework,
  :'check_id' as check_id,
  'Ensure SQL server"s TDE protector is encrypted with Customer-managed key (Automated)' as title,
  s.subscription_id,
  s.id AS server_id,
  case
    when p.kind != 'azurekeyvault'
      OR p.properties->>'serverKeyType' IS DISTINCT FROM 'AzureKeyVault'
      OR p.properties->>'uri' IS NULL
    then 'fail' else 'pass'
  end
FROM azure_sql_servers s
         LEFT JOIN azure_sql_server_encryption_protectors p ON
    s._cq_id = p._cq_parent_id
