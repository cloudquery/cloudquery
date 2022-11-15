INSERT INTO k8s_policy_results (resource_id, execution_time, framework, check_id, title, context, namespace,
                                resource_name, status)
select uid                              AS resource_id,
        :'execution_time'::timestamp     AS execution_time,
        :'framework'                     AS framework,
        :'check_id'                      AS check_id,
        'Stateful sets enforce cpu limits' AS title,
        context                          AS context,
        namespace                        AS namespace,
        name                             AS resource_name,
        CASE
            WHEN
                  -- Every container needs to have a CPU limit for the check to pass
                  (SELECT COUNT(*) FROM stateful_set_containers WHERE stateful_set_containers.uid = k8s_apps_stateful_sets.uid AND
                  stateful_set_containers.container->'resources'->'limits'->>'cpu' IS NULL) > 0
                THEN 'fail'
                ELSE 'pass'
            END                          AS status
FROM k8s_apps_stateful_sets;