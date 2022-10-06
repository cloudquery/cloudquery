WITH pod_containers AS (SELECT uid, value AS container 
                        FROM k8s_core_pods
                        CROSS JOIN jsonb_array_elements(spec_containers) AS value)

INSERT INTO k8s_policy_results (resource_id, execution_time, framework, check_id, title, context, namespace,
                                resource_name, status)
select uid                              AS resource_id,
        :'execution_time'::timestamp      AS execution_time,
        :'framework'                      AS framework,
        :'check_id'                       AS check_id,
        'Pod container privilege escalation disabled' AS title,
        context                           AS context,
        namespace                         AS namespace,
        name                              AS resource_name,
        CASE WHEN
            (SELECT COUNT(*) FROM pod_containers WHERE pod_containers.uid = k8s_core_pods.uid AND
              pod_containers.container->'securityContext'->>'allowPrivilegeEscalation' = 'true') > 0
            THEN 'fail'
            ELSE 'pass'
            END                          AS status
FROM k8s_core_pods;