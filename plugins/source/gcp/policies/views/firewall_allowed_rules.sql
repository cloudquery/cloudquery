CREATE OR REPLACE VIEW gcp_firewall_allowed_rules AS
WITH combined AS (
    SELECT * FROM gcp_compute_firewalls gcf, JSON_ARRAY_ELEMENTS(gcf.allowed) AS a
)
SELECT
    gcf.project_id,
    gcf."name",
    gcf.network,
    gcf.self_link AS link,
    gcf.direction,
    gcf.source_ranges,
    gcf.value->>'I_p_protocol' as ip_protocol,
    ARRAY(SELECT JSON_ARRAY_ELEMENTS_TEXT(gcf.value->'ports')) as ports,
    pr.range_start,
    pr.range_end,
    pr.single_port
FROM combined AS gcf
    LEFT JOIN (
        SELECT _cq_id, range_start, range_end, single_port
        FROM
            (
                SELECT
                    _cq_id,
                    split_part(p, '-', 1) :: INTEGER AS range_start,
                    split_part(p, '-', 2) :: INTEGER AS range_end,
                    NULL AS single_port
                FROM ( SELECT _cq_id, JSON_ARRAY_ELEMENTS_TEXT(value->'ports') AS p
                    FROM combined) AS f
                WHERE p ~ '^[0-9]+(-[0-9]+)$'
                UNION
                SELECT _cq_id, NULL AS range_start, NULL AS range_end, p AS single_port
                FROM ( SELECT _cq_id, JSON_ARRAY_ELEMENTS_TEXT(value->'ports') AS p
                    FROM combined) AS f
                WHERE p ~ '^[0-9]*$') AS s
    ) AS pr
    ON gcf._cq_id = pr._cq_id;
