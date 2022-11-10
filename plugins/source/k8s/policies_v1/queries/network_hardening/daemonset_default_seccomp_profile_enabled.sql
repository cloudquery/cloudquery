INSERT INTO k8s_policy_results (resource_id, execution_time, framework, check_id, title, context, namespace,
                                resource_name, status)
select uid                              AS resource_id,
        :'execution_time'::timestamp     AS execution_time,
        :'framework'                     AS framework,
        :'check_id'                      AS check_id,
        'Daemonset has seccomp enabled' AS title,
        context                          AS context,
        namespace                        AS namespace,
        name                             AS resource_name,
        CASE
            WHEN
                  
                  (SELECT * FROM k8s_apps_daemon_sets WHERE daemonset_containers.uid = k8s_apps_daemon_sets.uid AND
                  daemonset_containers.container->'resources'->'securityContext'->'seccompProfile'->>'type' != 'RuntimeDefault') > 0
                THEN 'fail'
                ELSE 'pass'
            END                          AS status
FROM k8s_apps_daemon_sets
