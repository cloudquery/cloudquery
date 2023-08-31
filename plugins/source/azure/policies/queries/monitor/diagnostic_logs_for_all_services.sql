INSERT INTO azure_policy_results
SELECT
    :'execution_time'                                                            AS execution_time,
    :'framework'                                                                 AS framework,
    :'check_id'                                                                  AS check_id,
    'Ensure that Diagnostic Logs are enabled for all services which support it.' AS title,
    amr.subscription_id                                                          AS subscription_id,
    amr.id                                                                       AS resource_id,
    CASE
        WHEN amds.id IS DISTINCT FROM NULL
        THEN 'pass'
        ELSE 'fail'
    END                                                                          AS status
FROM azure_monitor_resources AS amr
    LEFT JOIN azure_monitor_diagnostic_settings AS amds
        ON amr._cq_id = amds._cq_parent_id
