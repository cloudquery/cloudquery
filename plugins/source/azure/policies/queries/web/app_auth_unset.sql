INSERT INTO azure_policy_results (execution_time, framework, check_id, title, subscription_id, resource_id, status)
SELECT :'execution_time'                                                           AS execution_time,
       :'framework'                                                                AS framework,
       :'check_id'                                                                 AS check_id,
       'Ensure App Service Authentication is set on Azure App Service (Automated)' AS title,
       awa.subscription_id                                                         AS subscription_id,
       awa.id                                                                      AS resource_id,
       CASE
           WHEN awaas.enabled IS NULL OR awaas.enabled != TRUE
               THEN 'fail'
           ELSE 'pass'
           END                                                                     AS status
FROM azure_web_apps awa
         LEFT JOIN azure_web_site_auth_settings awaas ON
    awa.id = awaas.web_app_id
