INSERT INTO k8s_policy_results (resource_id, execution_time, framework, check_id, title, context, namespace,
                                resource_name, status)
select uid                              AS resource_id,
        :'execution_time'::timestamp     AS execution_time,
        :'framework'                     AS framework,
        :'check_id'                      AS check_id,
        'Replicaset has seccomp enabled' AS title,
        context                          AS context,
        namespace                        AS namespace,
        name                             AS resource_name,
        CASE
            WHEN                 
                  container->'resources'->'securityContext'->'seccompProfile'->>'type' != 'RuntimeDefault'
                THEN 'fail'
                ELSE 'pass'
            END                          AS status
FROM replica_set_containers;
