INSERT INTO azure_policy_results (execution_time, framework, check_id, title, subscription_id, resource_id, status)
SELECT :'execution_time'                                                               AS execution_time,
       :'framework'                                                                    AS framework,
       :'check_id'                                                                     AS check_id,
       'MFA should be enabled on accounts with write permissions on your subscription' AS title,
       subscription_id                                                             AS subscription_id,
       id                                                                           AS resource_id,
       CASE
           WHEN properties->'status'->>'code' IS DISTINCT FROM 'NotApplicable'
               AND properties->'status'->>'code' IS DISTINCT FROM 'Healthy'
               THEN 'fail'
           ELSE 'pass'
           END                                                                         AS status
FROM azure_security_assessments
WHERE name = '57e98606-6b1e-6193-0e3d-fe621387c16b'