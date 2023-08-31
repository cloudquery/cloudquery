insert into azure_policy_results
WITH diagnosis_logs AS (
    SELECT
        amds.subscription_id,
        amds.id || '/' || (coalesce(logs->>'category', logs->>'categoryGroup'))::text AS id,
        logs->>'category' IS DISTINCT FROM NULL AS hasCategory,
        (logs->'retentionPolicy'->>'days')::int >= 180 AS satisfyRetentionDays
    FROM azure_monitor_resources as amr
        LEFT JOIN azure_monitor_diagnostic_settings as amds ON amr._cq_id = amds._cq_parent_id,
        jsonb_array_elements(amds.properties->'logs') AS logs
    WHERE amr.type = 'Microsoft.KeyVault/vaults'
)
SELECT
    :'execution_time'                                        AS execution_time,
    :'framework'                                             AS framework,
    :'check_id'                                              AS check_id,
    'Ensure that logging for Azure Key Vault is ''Enabled''' AS title,
    subscription_id                                          AS subscription_id,
    id                                                       AS resource_id,
    CASE
        WHEN hasCategory AND satisfyRetentionDays
        THEN 'pass'
        ELSE 'fail'
    END                                                      AS status
FROM diagnosis_logs
