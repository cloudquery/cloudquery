INSERT INTO k8s_policy_results (resource_id, execution_time, framework, check_id, title, context, namespace,
                               resource_name, status)
select DISTINCT (k8s_core_namespaces.uid)                      AS resource_id,
                :'execution_time'::timestamp                   AS execution_time,
                :'framework'                                   AS framework,
                :'check_id'                                    AS check_id,
                'Namespace enforces resource quota cpu limits' AS title,
                k8s_core_namespaces.context                    AS context,
                k8s_core_namespaces.name                       AS namespace,
                k8s_core_namespaces.name                       AS resource_name,
                CASE
                    WHEN
                        hard -> 'limits.cpu' IS NULL
                        THEN 'fail'
                    ELSE 'pass'
                    END                                        AS status
FROM k8s_core_namespaces
         LEFT JOIN k8s_core_resource_quotas
                   ON k8s_core_resource_quotas.namespace = k8s_core_namespaces.name