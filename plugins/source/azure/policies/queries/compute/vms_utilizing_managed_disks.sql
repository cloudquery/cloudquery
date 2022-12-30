INSERT
INTO azure_policy_results (execution_time, framework, check_id, title, subscription_id, resource_id, status)
SELECT :'execution_time'                                              AS execution_time,
       :'framework'                                                   AS framework,
       :'check_id'                                                    AS check_id,
       'Ensure Virtual Machines are utilizing Managed Disks (Manual)' AS title,
       subscription_id                                                AS subscription_id,
       id                                                             AS resource_id,
       CASE
           WHEN properties -> 'storageProfile' -> 'osDisk' -> 'managedDisk' -> 'id' IS NULL
               THEN 'fail'
           ELSE 'pass'
           END                                                        AS status
FROM azure_compute_virtual_machines;
