INSERT INTO azure_policy_results (execution_time, framework, check_id, title, subscription_id, resource_id, status)
SELECT :'execution_time'                                                                   AS execution_time,
       :'framework'                                                                        AS framework,
       :'check_id'                                                                         AS check_id,
       'Role-Based Access Control (RBAC) should be used on Kubernetes Services' AS title,
       subscription_id                                                                 AS subscription_id,
       id                                                                               AS resource_id,
       CASE
           WHEN (properties ->> 'enableRBAC')::boolean IS distinct from TRUE
               THEN 'fail'
           ELSE 'pass'
           END                                                                             AS status
FROM azure_containerservice_managed_clusters