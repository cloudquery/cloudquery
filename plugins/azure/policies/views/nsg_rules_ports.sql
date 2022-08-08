create or replace view view_azure_nsg_rules as
SELECT
    ansg.subscription_id AS subscription_id,
    ansg.id AS nsg_id,
    ansg."name" AS nsg_name,
    ansgsr.id AS rule_id,
    ansgsr."access",
    ansgsr.direction,
    ansgsr.source_address_prefix,
    ansgsr.protocol,
    pr.range_start,
    pr.range_end,
    pr.single_port
FROM azure_network_security_groups ansg
    LEFT JOIN azure_network_security_group_security_rules ansgsr ON
        ansg.cq_id = ansgsr.security_group_cq_id
    LEFT JOIN (
        SELECT cq_id, range_start, range_end, single_port
        FROM (
            SELECT
                cq_id,
                Split_part(destination_port_range, '-', 1) :: INTEGER AS range_start,
                split_part(destination_port_range, '-', 2) :: INTEGER AS range_end,
                NULL AS single_port
            FROM azure_network_security_group_security_rules
            WHERE destination_port_range ~ '^[0-9]+(-[0-9]+)$'
            UNION
            SELECT cq_id, NULL AS range_start, NULL AS range_end, destination_port_range AS single_port
            FROM azure_network_security_group_security_rules
            WHERE destination_port_range ~ '^[0-9]*$'
        ) AS s
    ) AS pr ON ansgsr.cq_id = pr.cq_id;