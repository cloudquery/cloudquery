INSERT INTO k8s_policy_results (resource_id, execution_time, framework, check_id, title, context, namespace,
                               resource_name, status)
select DISTINCT (k8s_core_namespaces.uid)                       AS resource_id,
                :'execution_time'::timestamp                    AS execution_time,
                :'framework'                                    AS framework,
                :'check_id'                                     AS check_id,
                'Namespace enforces resource quota cpu request' AS title,
                k8s_core_namespaces.context                     AS context,
                k8s_core_namespaces.name                        AS namespace,
                k8s_core_namespaces.name                        AS resource_name,
                CASE
                    WHEN
                        (SELECT COUNT(*) FROM k8s_core_resource_quotas
                            WHERE namespace = k8s_core_namespaces.name
                            AND context = k8s_core_namespaces.context
                            AND spec_hard->>'requests.cpu' IS NOT NULL) = 0
                        THEN 'fail'
                    ELSE 'pass'
                    END                                         AS status
FROM k8s_core_namespaces;