WITH azure_nsg_rules as (SELECT ansg.id     AS nsg_id,
                                ansg."name" AS nsg_name,
                                ansgsr.id   AS rule_id,
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
                                  LEFT JOIN (SELECT cq_id, range_start, range_end, single_port
                                             FROM (SELECT cq_id,
                                                          Split_part(destination_port_range, '-', 1) :: INTEGER AS range_start,
                                                          split_part(destination_port_range, '-', 2) :: INTEGER AS range_end,
                                                          NULL                                                  AS single_port
                                                   FROM azure_network_security_group_security_rules
                                                   WHERE destination_port_range ~ '^[0-9]+(-[0-9]+)$'
                                                   UNION
                                                   SELECT cq_id,
                                                          NULL                   AS range_start,
                                                          NULL                   AS range_end,
                                                          destination_port_range AS single_port
                                                   FROM azure_network_security_group_security_rules
                                                   WHERE destination_port_range ~ '^[0-9]*$') AS s) AS pr
                                            ON ansgsr.cq_id = pr.cq_id;

)


INSERT INTO azure_policy_results (execution_time, framework, check_id, title, subscription_id, resource_id, status)
SELECT :'execution_time'                                            AS execution_time,
       :'framework'                                                 AS framework,
       :'check_id'                                                  AS check_id,
       'Management ports should be closed on your virtual machines' AS title,
       subscription_id                                          AS subscription_id,
       id                                                       AS resource_id,
       CASE
           WHEN source_address_prefix in ('0.0.0.0', '0.0.0.0/0', 'any', 'internet', '<nw>/0', '/0', '*')
               AND ((single_port = '3389' OR 3389 BETWEEN range_start AND range_end) or
                    (single_port = '22' OR 22 BETWEEN range_start AND range_end)
                   or single_port = '*')
               AND protocol = 'Tcp'
               AND "access" = 'Allow'
               AND direction = 'Inbound'
               THEN 'fail'
           ELSE 'pass'
           END                                                      AS status
FROM azure_nsg_rules