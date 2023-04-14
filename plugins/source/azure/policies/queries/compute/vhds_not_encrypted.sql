INSERT INTO azure_policy_results (execution_time, framework, check_id, title, subscription_id, resource_id, status)
SELECT :'execution_time'                           AS execution_time,
       :'framework'                                AS framework,
       :'check_id'                                 AS check_id,
       'Ensure that VHD''s are encrypted (Manual)' AS title,
       subscription_id                             AS subscription_id,
       id                                          AS resource_id,
       CASE
           WHEN (properties -> 'encryptionSettingsCollection' ->> 'enabled')::boolean IS DISTINCT FROM TRUE
           THEN 'fail'
           ELSE 'pass'
       END
FROM azure_compute_disks;
