insert into azure_policy_results
SELECT
    :'execution_time'                                                                          AS execution_time,
    :'framework'                                                                               AS framework,
    :'check_id'                                                                                AS check_id,
    'Ensure that Network Security Group Flow Log retention period is ''greater than 90 days''' AS title,
    subscription_id                                                                            AS subscription_id,
    id                                                                                         AS resource_id,
    CASE
        WHEN (properties->'retentionPolicy'->>'days')::INT >= 90
        THEN 'pass'
        ELSE 'fail'
    END                                                                                        AS status
FROM azure_network_watcher_flow_logs
