INSERT INTO azure_policy_results (execution_time, framework, check_id, title, subscription_id, resource_id, status)
SELECT :'execution_time'                                                                     AS execution_time,
       :'framework'                                                                          AS framework,
       :'check_id'                                                                           AS check_id,
       'Azure Key Vault Managed HSM should have purge protection enabled' AS title,
       subscription_id                                                                       AS subscription_id,
       id                                                                                    AS resource_id,
       CASE
           WHEN properties_enable_purge_protection IS NOT TRUE
               OR properties_enable_soft_delete IS NOT TRUE THEN 'fail'
           ELSE 'pass'
      END                                                                               AS status
FROM azure_keyvault_managed_hsms;
