--5.1.1 Ensure that the cluster-admin role is only used where required 
INSERT INTO k8s_policy_results (resource_id, execution_time, framework, check_id, title, context, namespace,
                                resource_name, status)
select uid                              AS resource_id,
        :'execution_time'::timestamp     AS execution_time,
        :'framework'                     AS framework,
        :'check_id'                      AS check_id,
        'cluster-admin role is only used where required' AS title,
        context                          AS context,
        namespace                        AS namespace,
        name                             AS resource_name,
        CASE WHEN
            role_ref->>'name' = 'cluster-admin' then 'fail'
            else 'pass'
            END AS status
FROM 
	k8s_rbac_cluster_role_bindings