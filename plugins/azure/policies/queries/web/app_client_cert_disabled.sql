INSERT INTO azure_policy_results (execution_time, framework, check_id, title, subscription_id, resource_id, status)
SELECT :'execution_time'                                                                                          AS execution_time,
       :'framework'                                                                                               AS framework,
       :'check_id'                                                                                                AS check_id,
       'Ensure the web app has ''Client Certificates (Incoming client certificates)'' set to ''On'' (Automated)' AS title,
       subscription_id                                                                                        AS subscription_id,
       id                                                                                                     AS resource_id,
       CASE
           WHEN kind LIKE 'app%' AND client_cert_enabled IS NOT TRUE
               THEN 'fail'
           ELSE 'pass'
           END                                                                                                    AS status
FROM azure_web_apps
