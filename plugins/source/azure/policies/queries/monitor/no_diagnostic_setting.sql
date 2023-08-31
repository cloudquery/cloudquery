insert into azure_policy_results
SELECT 
    :'execution_time'                              AS execution_time,
    :'framework'                                   AS framework,
    :'check_id'                                    AS check_id,
    'Ensure that a ''Diagnostics Setting'' exists' AS title,
    amr.subscription_id                            AS subscription_id,
    amr.id                                         AS resource_id,
    CASE
        WHEN amds.properties IS NULL
        THEN 'fail'
        ELSE 'pass'
    END                                            AS status
FROM
    azure_monitor_resources AS amr
    LEFT JOIN azure_monitor_diagnostic_settings AS amds
        ON amr._cq_id = amds._cq_parent_id
