WITH protected_instances AS (SELECT s.id AS instance_id
                             FROM azure_sql_managed_instances s
                                      LEFT JOIN azure_sql_managed_instance_encryption_protectors ep
                                                ON s._cq_id = ep._cq_parent_id
                             WHERE ep.properties ->> 'serverKeyType' = 'AzureKeyVault'
                               AND ep.properties ->> 'uri' IS NOT NULL)
insert into azure_policy_results
SELECT
  :'execution_time' as execution_time,
  :'framework' as framework,
  :'check_id' as check_id,
  'SQL managed instances should use customer-managed keys to encrypt data at rest' as title,
  i.subscription_id,
  i.id AS instance_id,
  case
    when  p.instance_id IS NULL
      then 'fail' else 'pass'
  end
FROM azure_sql_managed_instances i
         LEFT JOIN protected_instances p ON p.instance_id = i.id