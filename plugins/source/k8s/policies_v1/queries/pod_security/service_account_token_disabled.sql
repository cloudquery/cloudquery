INSERT INTO k8s_policy_results (resource_id, execution_time, framework, check_id, title, context, namespace,
                               resource_name, status)
select DISTINCT uid                                    AS resource_id,
                :'execution_time'::timestamp           AS execution_time,
                :'framework'                           AS framework,
                :'check_id'                            AS check_id,
                'Pod service account tokens disabled"' AS title,
                context                                AS context,
                namespace                              AS namespace,
                name                                   AS resource_name,
                CASE
                    WHEN
                        automount_service_account_token
                        THEN 'fail'
                    ELSE 'pass'
                    END                                AS status
FROM k8s_core_service_accounts
