create or replace view view_azure_nsg_rules as
WITH tmp AS (
    SELECT subscription_id, id, name, rls FROM azure_network_security_groups, jsonb_array_elements(properties -> 'securityRules') AS rls
),
ansg AS (
    SELECT tmp.*, dprs FROM tmp, jsonb_array_elements_text(rls->'destinationPortRanges') AS dprs
)
(
SELECT
    ansg.subscription_id AS subscription_id,
    ansg.id AS nsg_id,
    ansg."name" AS nsg_name,
    ansg.rls->>'id' AS rule_id,
    ansg.rls->>'access' AS access,
    ansg.rls->>'direction' AS direction,
    ansg.rls->>'sourceAddressPrefix' AS source_address_prefix,
    ansg.rls->>'protocol' AS protocol,
    split_part(ansg.dprs, '-', 1) :: INTEGER AS range_start,
    split_part(ansg.dprs, '-', 2) :: INTEGER AS range_end,
    NULL AS single_port
FROM ansg
WHERE ansg.dprs ~ '^[0-9]+(-[0-9]+)$'
)
UNION
(
SELECT
    ansg.subscription_id AS subscription_id,
    ansg.id AS nsg_id,
    ansg."name" AS nsg_name,
    ansg.rls->>'id' AS rule_id,
    ansg.rls->>'access' AS access,
    ansg.rls->>'direction' AS direction,
    ansg.rls->>'sourceAddressPrefix' AS source_address_prefix,
    ansg.rls->>'protocol' AS protocol,
    NULL AS range_start, NULL AS range_end,
    ansg.dprs AS single_port
FROM ansg
WHERE ansg.dprs ~ '^[0-9]*$'
    )
