insert into azure_policy_results
SELECT
    :'execution_time'                                                           AS execution_time,
    :'framework'                                                                AS framework,
    :'check_id'                                                                 AS check_id,
    'Ensure that ''Public access level'' is set to Private for blob containers' AS title,
    azsc.subscription_id                                                        AS subscription_id,
    azsc.id                                                                     AS resrouce_id,
    CASE
        WHEN azsc.properties->>'publicAccess' = 'None'
         AND NOT (asa.properties->>'allowBlobPublicAccess')::boolean
        THEN 'pass'
        ELSE 'fail'
    END                                                                         AS status
FROM azure_storage_containers azsc
    JOIN azure_storage_accounts asa on azsc._cq_parent_id = asa._cq_id
