CREATE OR REPLACE VIEW view_azure_nsg_dest_port_ranges AS
WITH security_rules AS (
    SELECT
        subscription_id,
        security_rules->>'id' AS id,
        security_rules->'properties'->>'access' AS access,
        security_rules->'properties'->>'protocol' AS protocol,
        security_rules->'properties'->>'direction' AS direction,
        security_rules->'properties'->>'sourceAddressPrefix' AS sourceAddressPrefix,
        security_rules->'properties'->>'destinationPortRange' AS destinationPortRange,
        security_rules->'properties'->'destinationPortRanges' AS destinationPortRanges
    FROM
        azure_network_security_groups ansg,
        jsonb_array_elements(ansg.properties->'securityRules') AS security_rules
),
dest_port AS (
    SELECT
        id,
        destinationPortRange
    FROM
        security_rules
    WHERE
        destinationPortRange <> '' OR destinationPortRange IS NOT NULL
),
unified_port_ranges AS (
    SELECT
        id,
        split_part(port_range, '-', 1) :: INT AS start_dest_port,
        split_part(port_range, '-', 2) :: INT AS end_dest_port
    FROM
        security_rules AS sr, jsonb_array_elements_text(sr.destinationPortRanges) AS port_range
    WHERE
        port_range  ~ '^[0-9]+-[0-9]+$'
    UNION
    SELECT
        id,
        port_range :: INT AS start_dest_port,
        port_range :: INT AS end_dest_port
    FROM
        security_rules AS sr, jsonb_array_elements_text(sr.destinationPortRanges) AS port_range
    WHERE
        port_range ~ '^[0-9]+$'
    UNION
    SELECT
        id,
        1 AS start_port,
        65535 AS end_port
    FROM
        dest_port
    WHERE
        destinationPortRange = '*'
    UNION
    SELECT
        id,
        destinationPortRange :: INT AS start_dest_port,
        destinationPortRange :: INT AS end_dest_port
    FROM
        dest_port
    WHERE
        destinationPortRange ~ '^[0-9]+$'
    UNION
    SELECT
        id,
        split_part(destinationPortRange, '-', 1) :: INT AS start_dest_port,
        split_part(destinationPortRange, '-', 2) :: INT AS end_dest_port
    FROM
        dest_port
    WHERE
        destinationPortRange ~ '^[0-9]+-[0-9]+$'
)
SELECT
    sr.id AS id,
    sr.subscription_id AS subscription_id,
    access,
    protocol,
    direction,
    sourceAddressPrefix,
    start_dest_port,
    end_dest_port
FROM
    security_rules sr
        JOIN unified_port_ranges
            ON sr.id = unified_port_ranges.id
