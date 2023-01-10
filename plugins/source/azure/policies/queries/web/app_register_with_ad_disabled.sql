INSERT INTO azure_policy_results (execution_time, framework, check_id, title, subscription_id, resource_id, status)
SELECT :'execution_time'                                                                        AS execution_time,
       :'framework'                                                                             AS framework,
       :'check_id'                                                                              AS check_id,
       'Ensure that Register with Azure Active Directory is enabled on App Service (Automated)' AS title,
       subscription_id                                                                      AS subscription_id,
       id                                                                                   AS resource_id,
       CASE
           WHEN identity->>'principalId' IS NULL OR identity->>'principalId' = ''
               THEN 'fail'
           ELSE 'pass'
           END                                                                                  AS status
FROM azure_web_apps
