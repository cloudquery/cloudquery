WITH deny_count AS (SELECT k8s_core_namespaces.uid,
                           k8s_core_namespaces.name AS namespace,
                           k8s_core_namespaces.context,
                           COUNT(*)
                           FILTER (WHERE policy_types @> ARRAY ['Egress'] AND pod_selector_match_labels::TEXT = '{}')
                    FROM k8s_core_namespaces
                             LEFT JOIN k8s_networking_network_policies
                                       ON k8s_networking_network_policies.namespace = k8s_core_namespaces.name
                             LEFT JOIN k8s_networking_network_policy_egress
                                       ON k8s_networking_network_policy_egress.network_policy_cq_id =
                                          k8s_networking_network_policies.cq_id
                    WHERE k8s_networking_network_policy_egress.cq_id IS NULL
                    GROUP BY k8s_core_namespaces.name,
                             k8s_core_namespaces.uid,
                             k8s_core_namespaces.context)


INSERT
INTO k8s_policy_results (resource_id, execution_time, framework, check_id, title, context, namespace,
                        resource_name, status)
select uid                                  AS resource_id,
       :'execution_time'::timestamp         AS execution_time,
       :'framework'                         AS framework,
       :'check_id'                          AS check_id,
       'Network policy default deny egress' AS title,
       context                              AS context,
       namespace                            AS namespace,
       namespace                            AS resource_name,
       CASE
           WHEN
               count = 0
               THEN 'fail'
           ELSE 'pass'
           END                              AS status
FROM deny_count