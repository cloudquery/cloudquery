INSERT
INTO azure_policy_results (execution_time, framework, check_id, title, subscription_id, resource_id, status)
SELECT :'execution_time'                                                      AS execution_time,
       :'framework'                                                           AS framework,
       :'check_id'                                                            AS check_id,
       'Ensure that ''OS and Data'' disks are encrypted with CMK (Automated)' AS title,
       v.subscription_id                                                      AS subscription_id,
       v.id                                                                   AS resource_id,
       CASE
           WHEN d.properties -> 'encryption' ->> 'type' NOT LIKE '%CustomerKey%'
               THEN 'fail'
           ELSE 'pass'
           END                                                                AS status
FROM azure_compute_virtual_machines v
         JOIN azure_compute_disks d ON
    LOWER(v.id) = LOWER(d.properties ->> 'managedBy')