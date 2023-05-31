insert into azure_policy_results
SELECT
    :'execution_time'                                                          AS execution_time,
    :'framework'                                                               AS framework,
    :'check_id'                                                                AS check_id,
    'Ensure storage for critical data are encrypted with Customer Managed Key' AS title,
    subscription_id                                                            AS subscription_id,
    id                                                                         AS resource_id,
    CASE
        WHEN properties->'encryption'->>'keySource' = 'Microsoft.Keyvault'
         AND properties->'encryption'->'keyvaultproperties' IS DISTINCT FROM NULL
        THEN 'pass'
        ELSE 'fail'
    END                                                                        AS status
FROM azure_storage_accounts
