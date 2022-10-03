INSERT INTO azure_policy_results (execution_time, framework, check_id, title, subscription_id, resource_id, status)
SELECT :'execution_time'                                                                     AS execution_time,
       :'framework'                                                                          AS framework,
       :'check_id'                                                                           AS check_id,
       'Deprecated accounts with owner permissions should be removed from your subscription' AS title,
       subscription_id                                                                       AS subscription_id,
       id                                                                                    AS resource_id,
       CASE
           WHEN a.code IS NULL
               THEN 'fail'
           ELSE 'pass'
           END                                                                               AS status
FROM azure_subscription_subscriptions s
         LEFT OUTER JOIN azure_security_assessments a
                         ON
                                     s.id = '/subscriptions/' || a.subscription_id
                                 AND a.name = 'e52064aa-6853-e252-a11e-dffc675689c2'
                                 AND (a.code IS NOT DISTINCT FROM 'NotApplicable'
                                 OR a.code IS NOT DISTINCT FROM 'Healthy')