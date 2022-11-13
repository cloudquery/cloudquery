WITH default_request_cpu_limits AS (
   SELECT context, namespace, value->'default_request'->>'cpu' AS default_request_cpu_limit
   FROM k8s_core_limit_ranges CROSS JOIN jsonb_array_elements(k8s_core_limit_ranges.spec_limits))

INSERT
INTO k8s_policy_results (resource_id, execution_time, framework, check_id, title, context, namespace,
                        resource_name, status)
select uid                                         AS resource_id,
       :'execution_time'::timestamp                AS execution_time,
       :'framework'                                AS framework,
       :'check_id'                                 AS check_id,
       'Namespaces CPU request resource quota' AS title,
       context                                     AS context,
       name                                        AS namespace,
       name                                        AS resource_name,
       CASE
           WHEN
               (SELECT COUNT(default_request_cpu_limit) FROM default_request_cpu_limits 
                  WHERE namespace = k8s_core_namespaces.name
                  AND context = k8s_core_namespaces.context) = 0
               THEN 'fail'
           ELSE 'pass'
           END                                     AS status
FROM k8s_core_namespaces;