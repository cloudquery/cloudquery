-- SELECT gcf.project_id, gcf.id, gcf.name, gcf.self_link AS link, count(*) AS broken_rules
-- FROM gcp_compute_firewalls gcf
--     JOIN gcp_compute_firewall_allowed gcfa ON
--         gcf.cq_id = gcfa.firewall_cq_id
-- WHERE
--     NOT ARRAY [
--         '35.191.0.0/16', '130.211.0.0/22'
--     ] <@ gcf.source_ranges AND NOT (ip_protocol = 'tcp' AND ports @> ARRAY ['80'])
-- GROUP BY 1,2,3,4
-- HAVING count(*) > 0;


INSERT INTO gcp_policy_results (resource_id, execution_time, framework, check_id, title, project_id, status)
SELECT DISTINCT gcf.id                                                                                                                                                                                     AS resource_id,
                :'execution_time'::timestamp                                                                                                                                                               AS execution_time,
                :'framework'                                                                                                                                                                               AS framework,
                :'check_id'                                                                                                                                                                                AS check_id,
                'GCP CIS3.10 Ensure Firewall Rules for instances behind Identity Aware Proxy (IAP) only allow the traffic from Google Cloud Loadbalancer (GCLB) Health Check and Proxy Addresses (Manual)' AS title,
                gcf.project_id                                                                                                                                                                             AS project_id,
                CASE
                    WHEN
                            NOT ARRAY [
                                    '35.191.0.0/16', '130.211.0.0/22'
                                    ] <@ gcf.source_ranges AND NOT (gcfa.ip_protocol = 'tcp' AND gcfa.ports @> ARRAY ['80'])
                        THEN 'fail'
                    ELSE 'pass'
                    END                                                                                                                                                                                    AS status
FROM gcp_compute_firewalls gcf
         JOIN gcp_compute_firewall_allowed gcfa ON
    gcf.cq_id = gcfa.firewall_cq_id;