insert into azure_policy_results
SELECT
  :'execution_time',
  :'framework',
  :'check_id',
  'Ensure SQL server"s TDE protector is encrypted with Customer-managed key (Automated)',
  s.subscription_id,
  s.id AS server_id,
  case
    when p.kind != 'azurekeyvault'
      OR p.server_key_type IS DISTINCT FROM 'AzureKeyVault'
      OR uri IS NULL
    then 'fail' else 'pass'
  end
FROM azure_sql_servers s
         LEFT JOIN azure_sql_encryption_protectors p ON
    s.id = p.sql_server_id
