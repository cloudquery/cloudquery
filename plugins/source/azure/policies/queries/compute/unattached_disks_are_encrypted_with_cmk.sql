INSERT
INTO azure_policy_results (execution_time, framework, check_id, title, subscription_id, resource_id, status)
SELECT :'execution_time'                                                     AS execution_time,
       :'framework'                                                          AS framework,
       :'check_id'                                                           AS check_id,
       'Ensure that ''Unattached disks'' are encrypted with CMK (Automated)' AS title,
       subscription_id                                                       AS subscription_id,
       id                                                                    AS resource_id,
       CASE
           WHEN properties -> 'encryption'->>'type' NOT LIKE '%CustomerKey%'
               THEN 'fail'
           ELSE 'pass'
           END                                                               AS status
FROM azure_compute_disks
WHERE properties ->> 'diskState' = 'Unattached'