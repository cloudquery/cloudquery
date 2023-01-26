INSERT INTO azure_policy_results (execution_time, framework, check_id, title, subscription_id, resource_id, status)
SELECT :'execution_time'                                                                   AS execution_time,
       :'framework'                                                                        AS framework,
       :'check_id'                                                                         AS check_id,
       'External accounts with owner permissions should be removed from your subscription' AS title,
       subscription_id                                                                     AS subscription_id,
       id                                                                                  AS resource_id,
       CASE
           WHEN a.properties->>'code' IS NULL
               THEN 'fail'
           ELSE 'pass'
           END                                                                             AS status
FROM azure_subscription_subscriptions s
         LEFT OUTER JOIN azure_security_assessments a
                         ON
                                     s.id = '/subscriptions/' || a.subscription_id
                                 AND a.name = 'c3b6ae71-f1f0-31b4-e6c1-d5951285d03d'
                                 AND (a.properties->>'code' IS NOT DISTINCT FROM 'NotApplicable'
                                 OR a.properties->>'code' IS NOT DISTINCT FROM 'Healthy')