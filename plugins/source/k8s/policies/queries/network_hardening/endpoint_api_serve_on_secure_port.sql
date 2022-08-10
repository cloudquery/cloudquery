INSERT INTO k8s_policy_results (resource_id, execution_time, framework, check_id, title, context, namespace,
                               resource_name, status)
select uid                                   AS resource_id,
       :'execution_time'::timestamp          AS execution_time,
       :'framework'                          AS framework,
       :'check_id'                           AS check_id,
       'Endpoint API served on secure port"' AS title,
       context                               AS context,
       namespace                             AS namespace,
       k8s_core_endpoints.name               AS resource_name,
       CASE
           WHEN
                       k8s_core_endpoints.name = 'kubernetes'
                   AND NOT (
                           k8s_core_endpoint_subset_ports.name = 'https'
                       AND (
                                       k8s_core_endpoint_subset_ports.port = '443'
                                   OR k8s_core_endpoint_subset_ports.port = '6443'
                               )
                   )
               THEN 'fail'
           ELSE 'pass'
           END                               AS status
FROM k8s_core_endpoint_subset_ports
         JOIN k8s_core_endpoint_subsets
              ON k8s_core_endpoint_subsets.cq_id = k8s_core_endpoint_subset_ports.endpoint_subset_cq_id
         JOIN k8s_core_endpoints ON k8s_core_endpoints.cq_id = k8s_core_endpoint_subsets.endpoint_cq_id