INSERT INTO azure_policy_results (execution_time, framework, check_id, title, subscription_id, resource_id, status)
SELECT :'execution_time'                                 AS execution_time,
       :'framework'                                      AS framework,
       :'check_id'                                       AS check_id,
       'Ensure the key vault is recoverable (Automated)' AS title,
       subscription_id                                   AS subscription_id,
       id                                                AS resource_id,
       CASE
           WHEN enable_soft_delete IS NOT TRUE OR enable_purge_protection IS NOT TRUE
               THEN 'fail'
           ELSE 'pass'
           END                                           AS status
FROM azure_keyvault_vaults;
