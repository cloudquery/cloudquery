INSERT INTO azure_policy_results (execution_time, framework, check_id, title, subscription_id, resource_id, status)
SELECT :'execution_time'                                                               AS execution_time,
       :'framework'                                                                    AS framework,
       :'check_id'                                                                     AS check_id,
       'MFA should be enabled on accounts with owner permissions on your subscription' AS title,
       subscription_id                                                             AS subscription_id,
       id                                                                           AS resource_id,
       CASE
           WHEN status->>'code' IS DISTINCT FROM 'NotApplicable'
               AND status->>'code' IS DISTINCT FROM 'Healthy'
               THEN 'fail'
           ELSE 'pass'
           END                                                                         AS status
FROM azure_security_assessments
WHERE name = '94290b00-4d0c-d7b4-7cea-064a9554e681'