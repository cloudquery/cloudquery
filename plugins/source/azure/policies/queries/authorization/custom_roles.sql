INSERT INTO azure_policy_results (execution_time, framework, check_id, title, subscription_id, resource_id, status)
SELECT :'execution_time'                                                                   AS execution_time,
       :'framework'                                                                        AS framework,
       :'check_id'                                                                         AS check_id,
       'External accounts with owner permissions should be removed from your subscription' AS title,
       mc.subscription_id                                                                  AS subscription_id,
       mc.id                                                                               AS resource_id,
       CASE
           WHEN (properties ->> 'enableRBAC')::boolean IS distinct from TRUE
               THEN 'fail'
           ELSE 'pass'
           END                                                                             AS status
FROM azure_containerservice_managed_clusters AS mc
         INNER JOIN azure_subscription_subscriptions AS sub ON sub.id = mc.subscription_id
