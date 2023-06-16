INSERT INTO gcp_policy_results (resource_id, execution_time, framework, check_id, title, project_id, status)
SELECT DISTINCT gcn.name                                                                    AS resource_id,
                :'execution_time'::timestamp                                                AS execution_time,
                :'framework'                                                                AS framework,
                :'check_id'                                                                 AS check_id,
                'Ensure that Cloud DNS logging is enabled for all VPC networks (Automated)' AS title,
                gcn.project_id                                                              AS project_id,
                CASE
                    WHEN
                        gdp.enable_logging = FALSE
                        THEN 'fail'
                    ELSE 'pass'
                    END                                                                     AS status
FROM gcp_dns_policies gdp, JSONB_ARRAY_ELEMENTS(gdp.networks) AS gdpn
    JOIN gcp_compute_networks gcn ON gcn.self_link = REPLACE(gdpn->>'networkUrl', 'compute.googleapis', 'www.googleapis')
