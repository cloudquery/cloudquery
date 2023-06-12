insert into azure_policy_results
SELECT
    :'execution_time'                                                        AS execution_time,
    :'framework'                                                             AS framework,
    :'check_id'                                                              AS check_id,
    'Ensure default network access rule for Storage Accounts is set to deny' AS title,
    subscription_id                                                          AS subscription_id,
    id                                                                       AS resource_id,
    CASE
        WHEN properties->'networkAcls'->>'defaultAction' = 'Allow'
        THEN 'fail'
        ELSE 'pass'
    END                                                                      AS status
FROM azure_storage_accounts
