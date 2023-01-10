WITH pod_volumes AS (SELECT uid, value AS volumes
                     FROM k8s_core_pods
                     CROSS JOIN jsonb_array_elements(spec_volumes) AS value)

INSERT INTO k8s_policy_results (resource_id, execution_time, framework, check_id, title, context, namespace,
                                resource_name, status)
select uid                              AS resource_id,
            :'execution_time'::timestamp      AS execution_time,
            :'framework'                      AS framework,
            :'check_id'                       AS check_id,
            'Pod volume don''t have a hostPath'            AS title,
            context                           AS context,
            namespace                         AS namespace,
            name                              AS resource_name,
            CASE WHEN
               (SELECT COUNT(*) FROM pod_volumes WHERE pod_volumes.uid = k8s_core_pods.uid AND
                 pod_volumes.volumes->>'hostPath' IS NOT NULL) > 0
               THEN 'fail'
               ELSE 'pass'
               END                          AS status
FROM k8s_core_pods;
