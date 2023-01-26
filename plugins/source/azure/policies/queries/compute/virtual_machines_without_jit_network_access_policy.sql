WITH jit_vms AS (SELECT properties->'virtualMachines'->>'id' AS vm_id
                 FROM azure_security_jit_network_access_policies
                 WHERE properties->>'provisioningState' = 'Succeeded')

INSERT
INTO azure_policy_results (execution_time, framework, check_id, title, subscription_id, resource_id, status)
SELECT :'execution_time'                                                                                   AS execution_time,
       :'framework'                                                                                        AS framework,
       :'check_id'                                                                                         AS check_id,
       'Management ports of virtual machines should be protected with just-in-time network access control' AS title,
       subscription_id                                                                                     AS subscription_id,
       id                                                                                                  AS resource_id,
       CASE
           WHEN j.vm_id = NULL
               THEN 'fail'
           ELSE 'pass'
           END                                                                                             AS status
FROM azure_compute_virtual_machines vm
         LEFT JOIN jit_vms j ON vm.id = j.vm_id
