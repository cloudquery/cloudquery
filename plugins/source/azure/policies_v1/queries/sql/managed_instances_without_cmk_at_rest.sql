WITH protected_instances AS (SELECT s.id AS instance_id
                             FROM azure_sql_managed_instances s
                                      LEFT JOIN azure_sql_managed_instance_encryption_protectors ep
                                                ON s.id = ep.sql_managed_instance_id
                             WHERE ep.server_key_type = 'AzureKeyVault'
                               AND ep.uri IS NOT NULL)
insert into azure_policy_results
SELECT
  :'execution_time',
  :'framework',
  :'check_id',
  'SQL managed instances should use customer-managed keys to encrypt data at rest',
  i.subscription_id,
  i.id AS instance_id,
  case
    when  p.instance_id IS NULL
      then 'fail' else 'pass'
  end
FROM azure_sql_managed_instances i
         LEFT JOIN protected_instances p ON p.instance_id = i.id