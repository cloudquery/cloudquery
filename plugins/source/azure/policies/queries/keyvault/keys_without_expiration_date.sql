INSERT INTO azure_policy_results (execution_time, framework, check_id, title, subscription_id, resource_id, status)
SELECT :'execution_time'                                                AS execution_time,
       :'framework'                                                     AS framework,
       :'check_id'                                                      AS check_id,
       'Ensure that the expiration date is set on all keys (Automated)' AS title,
       akv.subscription_id                                              AS subscription_id,
       akvk.id                                                          AS resource_id,
       CASE
           WHEN (akvk.properties -> 'attributes'->>'enabled')::boolean = TRUE
            AND (akvk.properties -> 'attributes'->>'exp') IS NULL
           THEN 'fail'
           ELSE 'pass'
       END                                                              AS status
FROM azure_keyvault_keyvault akv
    JOIN azure_keyvault_keyvault_keys akvk
      ON akv._cq_id = akvk._cq_parent_id
