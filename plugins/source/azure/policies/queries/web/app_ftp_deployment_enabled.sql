INSERT INTO azure_policy_results (execution_time, framework, check_id, title, subscription_id, resource_id, status)
SELECT :'execution_time'                                 AS execution_time,
       :'framework'                                      AS framework,
       :'check_id'                                       AS check_id,
       'Ensure FTP deployments are disabled (Automated)' AS title,
       subscription_id                                   AS subscription_id,
       id                                                AS resource_id,
       CASE
           WHEN p.user_name NOT LIKE concat('%', a."name", '%')
               THEN 'fail'
           ELSE 'pass'
           END                                           AS status
FROM azure_web_apps a
         LEFT JOIN azure_web_app_publishing_profiles p ON
    a.cq_id = p.app_cq_id
