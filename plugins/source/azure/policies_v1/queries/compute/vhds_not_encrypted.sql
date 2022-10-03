WITH vm_disks AS (SELECT subscription_id,
                         id,
                         name,
                         jsonb_array_elements(instance_view -> 'disks') AS disk
                  FROM azure_compute_virtual_machines),
     disk_encryptions AS (SELECT subscription_id,
                                 id,
                                 name,
                                 disk -> 'name'                                             AS disk_name,
                                 (disk -> 'encryptionSettings' -> 0 ->> 'enabled')::BOOLEAN AS encryption_enabled
                          FROM vm_disks)
INSERT
INTO azure_policy_results (execution_time, framework, check_id, title, subscription_id, resource_id, status)
SELECT :'execution_time'                           AS execution_time,
       :'framework'                                AS framework,
       :'check_id'                                 AS check_id,
       'Ensure that VHD''s are encrypted (Manual)' AS title,
       subscription_id                             AS subscription_id,
       id                                          AS resource_id,
       case
           when encryption_enabled IS NULL OR encryption_enabled != TRUE
               then 'fail'
           else 'pass'
           end                                     AS status
FROM disk_encryptions
