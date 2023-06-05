INSERT INTO azure_policy_results
SELECT
    :'execution_time'                                 AS execution_time,
    :'framework'                                      AS framework,
    :'check_id'                                       AS check_id,
    'Ensure FTP deployments are disabled (Automated)' AS title,
    aawac.subscription_id                             AS subscription_id,
    aawac.id                                          AS resource_id,
    CASE
        WHEN aawac.properties->>'ftpsState' = 'AllAllowed'
        THEN 'fail'
        ELSE 'pass'
    END                                               AS status
FROM azure_appservice_web_apps as aawa
    JOIN azure_appservice_web_app_configurations aawac
        ON aawa._cq_id = aawac._cq_parent_id
