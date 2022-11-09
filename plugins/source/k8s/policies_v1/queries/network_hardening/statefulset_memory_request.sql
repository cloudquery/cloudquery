With stateful_set_containers AS (SELECT uid, value AS container 
                               FROM k8s_apps_stateful_sets
                               CROSS JOIN jsonb_array_elements(spec_template->'spec'->'containers') AS value)

Insert Into k8s_policy_results (resource_id, execution_time, framework, check_id, title, context, namespace,
                                resource_name, status)
select uid                              AS resource_id,
        :'execution_time'::timestamp     AS execution_time,
        :'framework'                     AS framework,
        :'check_id'                      AS check_id,
        'Statefulsets enforce memory requests' AS title,
        context                          AS context,
        namespace                        AS namespace,
        name                             AS resource_name,
        CASE
            WHEN
                  -- Every container needs to have a memory request for the check to pass
                  (SELECT COUNT(*) FROM stateful_set_containers WHERE stateful_set_containers.uid = k8s_apps_stateful_sets.uid AND
                  stateful_set_containers.container->'resources'->'requests'->>'memory' IS NULL) > 0
                THEN 'fail'
                ELSE 'pass'
            END                          AS status
FROM k8s_apps_stateful_sets