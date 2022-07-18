CREATE OR REPLACE VIEW gcp_firewall_allowed_rules AS
SELECT
    gcf.project_id,
    gcf."name",
    gcf.network,
    gcf.self_link AS link,
    gcf.direction,
    gcf.source_ranges,
    gcfa.ip_protocol,
    gcfa.ports,
    pr.range_start,
    pr.range_end,
    pr.single_port
FROM gcp_compute_firewalls gcf
    LEFT JOIN gcp_compute_firewall_allowed gcfa ON
        gcf.cq_id = gcfa.firewall_cq_id
    LEFT JOIN (
        SELECT cq_id, range_start, range_end, single_port
        FROM
            (
                SELECT
                    cq_id,
                    Split_part(p, '-', 1) :: INTEGER AS range_start,
                    split_part(p, '-', 2) :: INTEGER AS range_end,
                    NULL AS single_port
                FROM ( SELECT cq_id, UNNEST(ports) AS p
                    FROM gcp_compute_firewall_allowed) AS f
                WHERE p ~ '^[0-9]+(-[0-9]+)$'
                UNION
                SELECT cq_id, NULL AS range_start, NULL AS range_end, p AS single_port
                FROM ( SELECT cq_id, UNNEST(ports) AS p
                    FROM gcp_compute_firewall_allowed) AS f
                WHERE p ~ '^[0-9]*$') AS s
    ) AS pr
    ON gcfa.cq_id = pr.cq_id;
