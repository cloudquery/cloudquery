WITH combined AS (
    SELECT * FROM gcp_compute_firewalls gcf, JSONB_ARRAY_ELEMENTS(gcf.allowed) AS a
)

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
                                    ] <@ gcf.source_ranges AND NOT (gcf.value->>'I_p_protocol' = 'tcp' AND ARRAY(SELECT JSONB_ARRAY_ELEMENTS_TEXT(gcf.value->'ports')) @> ARRAY ['80'])
                        THEN 'fail'
                    ELSE 'pass'
                    END                                                                                                                                                                                    AS status
FROM combined AS gcf