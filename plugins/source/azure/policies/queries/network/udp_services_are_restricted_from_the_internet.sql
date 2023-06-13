insert into azure_policy_results
WITH conditions AS (
    SELECT
        subscription_id,
        id,
        access = 'Allow' AS acceptAccess,
        protocol = '*' OR upper(protocol) = 'UDP' AS matchProtocol,
        direction = 'Inbound' AS isInbound,
        sourceAddressPrefix IN ('*', '0.0.0.0', '<nw>/0', '/0', 'internet', 'any') AS matchPrefix,
        (
            SELECT bool_or(port_number BETWEEN start_dest_port AND end_dest_port)
            FROM unnest(Array[53, 123, 161, 389, 1900]) AS port_number
        ) AS inRange
    FROM view_azure_nsg_dest_port_ranges
),
statuses_by_port_range AS (
    SELECT
        subscription_id,
        id,
        acceptAccess AND matchProtocol AND isInbound AND matchPrefix AND inRange AS failed
    FROM conditions
),
statuses AS (
    SELECT subscription_id, id, bool_or(failed) AS failed
    FROM statuses_by_port_range
    GROUP BY subscription_id, id
)
SELECT
    :'execution_time'                                           AS execution_time,
    :'framework'                                                AS framework,
    :'check_id'                                                 AS check_id,
    'Ensure that UDP Services are restricted from the Internet' AS title,
    subscription_id                                             AS subscription_id,
    id                                                          AS resource_id,
    CASE
        WHEN failed
        THEN 'fail'
        ELSE 'pass'
    END                                                         AS status
FROM statuses;
