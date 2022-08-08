INSERT
INTO k8s_policy_results (resource_id, execution_time, framework, check_id, title, context, namespace,
                        resource_name, status)
select distinct uid                                          AS resource_id,
                :'execution_time'::timestamp                 AS execution_time,
                :'framework'                                 AS framework,
                :'check_id'                                  AS check_id,
                'Network policy default don''t allow egress' AS title,
                context                                      AS context,
                namespace                                    AS namespace,
                name                                         AS resource_name,
                CASE
                    WHEN
                            k8s_networking_network_policy_egress_ports.network_policy_egress_cq_id IS NULL
                            OR k8s_networking_network_policy_egress_to.network_policy_egress_cq_id IS NULL
                        THEN 'fail'
                    ELSE 'pass'
                    END                                      AS status
FROM k8s_networking_network_policies
         LEFT JOIN k8s_networking_network_policy_egress
                   ON k8s_networking_network_policy_egress.network_policy_cq_id =
                      k8s_networking_network_policies.cq_id
         LEFT JOIN k8s_networking_network_policy_egress_ports
                   ON k8s_networking_network_policy_egress_ports.network_policy_egress_cq_id =
                      k8s_networking_network_policy_egress.network_policy_uid::UUID
         LEFT JOIN k8s_networking_network_policy_egress_to
                   ON k8s_networking_network_policy_egress_to.network_policy_egress_cq_id =
                      k8s_networking_network_policy_egress.network_policy_uid::UUID