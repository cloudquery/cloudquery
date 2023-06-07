insert into azure_policy_results
SELECT
    :'execution_time'                                                                   AS execution_time,
    :'framework'                                                                        AS framework,
    :'check_id'                                                                         AS check_id,
    'Ensure the storage container storing the activity logs is not publicly accessible' AS title,
    subscription_id                                                                     AS subscription_id,
    id                                                                                  AS resource_id,
    CASE
        WHEN properties->>'publicAccess' = 'None'
        THEN 'pass'
        ELSE 'fail'
    END                                                                                 AS status
FROM azure_storage_containers
WHERE
    name = 'insights-activity-logs'
