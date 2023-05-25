insert into azure_policy_results
SELECT
    :'execution_time'                                 AS execution_time,
    :'framework'                                      AS framework,
    :'check_id'                                       AS check_id,
    'Ensure soft delete is enabled for Azure Storage' AS title
    subscription_id                                   AS subscription_id,
    id                                                AS resource_id,
    CASE
        WHEN (properties->'deleteRetentionPolicy'->>'enabled')::boolean
        THEN 'pass'
        ELSE 'fail'
    END                                               AS status
FROM azure_storage_blob_services
