-- SELECT gcn.id, gcn.project_id, gcn.self_link AS network, gcs.self_link AS subnetwork, gcs.enable_flow_logs
-- FROM gcp_compute_networks gcn
--     JOIN gcp_compute_subnetworks gcs ON
--         gcn.self_link = gcs.network
-- WHERE gcs.enable_flow_logs = FALSE;


INSERT INTO gcp_policy_results (resource_id, execution_time, framework, check_id, title, project_id, status)
SELECT DISTINCT gcn.name                                                                             AS resource_id,
                :'execution_time'::timestamp                                                         AS execution_time,
                :'framework'                                                                         AS framework,
                :'check_id'                                                                          AS check_id,
                'Ensure that VPC Flow Logs is enabled for every subnet in a VPC Network (Automated)' AS title,
                gcn.project_id                                                                       AS project_id,
                CASE
                    WHEN
                        gcs.enable_flow_logs = FALSE
                        THEN 'fail'
                    ELSE 'pass'
                    END                                                                              AS status
FROM gcp_compute_networks gcn
         JOIN gcp_compute_subnetworks gcs ON
    gcn.self_link = gcs.network;