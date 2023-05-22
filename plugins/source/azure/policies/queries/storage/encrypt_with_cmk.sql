insert into azure_policy_results
SELECT
    :'execution_time'                                                                                                  AS execution_time,
    :'framework'                                                                                                       AS framework,
    :'check_id'                                                                                                        AS check_id,
    'Ensure the storage account containing the container with activity logs is encrypted with BYOK (Use Your Own Key)' AS title,
    subscription_id                                                                                                    AS subscription_id,
    id                                                                                                                 AS resource_id,
    CASE
        WHEN properties->'encryption'->>'keySource' = 'Microsoft.Keyvault' 
         AND properties->'encryption'->'keyvaultproperties' IS DISTINCT FROM NULL
        THEN 'pass'
        ELSE 'fail'
    END                                                                                                                AS status
FROM azure_storage_accounts
