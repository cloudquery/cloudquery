WITH daemonset_containers AS (SELECT uid, value AS container 
                               FROM k8s_apps_daemon_sets
                               CROSS JOIN jsonb_array_elements(spec_template->'spec'->'containers') AS value)

INSERT INTO k8s_policy_results (resource_id, execution_time, framework, check_id, title, context, namespace,
                                resource_name, status)
select uid                              AS resource_id,
        :'execution_time'::timestamp     AS execution_time,
        :'framework'                     AS framework,
        :'check_id'                      AS check_id,
        'Daemonset uses default namespace' AS title,
        context                          AS context,
        namespace                        AS namespace,
        name                             AS resource_name,
        CASE
            WHEN
                  
                  (SELECT * FROM k8s_apps_daemon_sets WHERE daemonset_containers.namespace = 'default')
                THEN 'fail'
                ELSE 'pass'
            END                          AS status
FROM k8s_apps_daemon_sets
