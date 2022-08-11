-- SELECT project_id,
--        gci."id",
--        gci.self_link AS link
-- FROM gcp_compute_instances gci
--          LEFT JOIN gcp_compute_instance_network_interfaces gcini ON
--     gci.id = gcini.instance_id
--          LEFT JOIN gcp_compute_instance_network_interface_access_configs gciniac ON
--     gcini.cq_id = gciniac.instance_network_interface_cq_id
-- WHERE gci."name" NOT LIKE 'gke-%'
--   AND (gciniac.nat_ip IS NOT NULL
--     OR gciniac.nat_ip != '')
-- GROUP BY 1, 2, 3
-- HAVING count(gciniac.*) > 0; -- noqa


INSERT INTO gcp_policy_results (resource_id, execution_time, framework, check_id, title, project_id, status)
SELECT DISTINCT gci.id                                                                     AS resource_id,
                :'execution_time'::timestamp                                               AS execution_time,
                :'framework'                                                               AS framework,
                :'check_id'                                                                AS check_id,
                'Ensure that Compute instances do not have public IP addresses (Automated' AS title,
                gci.project_id                                                             AS project_id,
                CASE
                    WHEN
                                gci."name" NOT LIKE 'gke-%'
                            AND (gciniac.nat_ip IS NOT NULL OR gciniac.nat_ip != '')
                        THEN 'fail'
                    ELSE 'pass'
                    END                                                                    AS status
FROM gcp_compute_instances gci
         LEFT JOIN gcp_compute_instance_network_interfaces gcini ON
    gci.id = gcini.instance_id
         LEFT JOIN gcp_compute_instance_network_interface_access_configs gciniac ON
    gcini.cq_id = gciniac.instance_network_interface_cq_id;