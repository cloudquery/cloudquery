WITH replica_set_containers AS (SELECT uid, value AS container 
                               FROM k8s_apps_replica_sets
                               CROSS JOIN jsonb_array_elements(spec_template->'spec'->'containers') AS value)

INSERT INTO k8s_policy_results (resource_id, execution_time, framework, check_id, title, context, namespace,
                                resource_name, status)
select uid                              AS resource_id,
        :'execution_time'::timestamp                          AS execution_time,
        :'framework'                                          AS framework,
        :'check_id'                                           AS check_id,
        'ReplicaSet containers must run as non-root' AS title,
        context                                               AS context,
        namespace                                             AS namespace,
        name                                                  AS resource_name,
        CASE
            WHEN
                (SELECT COUNT(*) FROM replica_set_containers WHERE replica_set_containers.uid = k8s_apps_replica_sets.uid AND
                  replica_set_containers.container->'securityContext' ->> 'runAsNonRoot' IS DISTINCT FROM 'true') > 0
                THEN 'fail'
            ELSE 'pass'
            END                                               AS status
FROM k8s_apps_replica_sets;