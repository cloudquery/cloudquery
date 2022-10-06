-- Join every row in the replica_set table with its json array of containers.
WITH replica_set_containers AS (SELECT uid, value AS container 
                               FROM k8s_apps_replica_sets
                               CROSS JOIN jsonb_array_elements(spec_template->'spec'->'containers') AS value)

INSERT INTO k8s_policy_results (resource_id, execution_time, framework, check_id, title, context, namespace,
                                resource_name, status)
select uid                              AS resource_id,
        :'execution_time'::timestamp     AS execution_time,
        :'framework'                     AS framework,
        :'check_id'                      AS check_id,
        'Replicaset enforces cpu limits' AS title,
        context                          AS context,
        namespace                        AS namespace,
        name                             AS resource_name,
        CASE
            WHEN
                  -- Every container needs to have a CPU limit for the check to pass
                  (SELECT COUNT(*) FROM replica_set_containers WHERE replica_set_containers.uid = k8s_apps_replica_sets.uid AND
                  replica_set_containers.container->'resources'->'limits'->>'cpu' IS NULL) > 0
                THEN 'fail'
                ELSE 'pass'
            END                          AS status
FROM k8s_apps_replica_sets