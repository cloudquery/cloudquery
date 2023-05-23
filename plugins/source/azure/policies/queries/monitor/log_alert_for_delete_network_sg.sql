INSERT INTO azure_policy_results
WITH fields AS (
    SELECT
        subscription_id,
        id,
        location,
        (properties->'enabled')::boolean AS enabled,
        conditions->>'field' AS field,
        conditions->>'equals' AS equals
    FROM azure_monitor_activity_log_alerts, jsonb_array_elements(properties->'condition'->'allOf') AS conditions
),
scopes AS (
    SELECT
        subscription_id,
        id,
        scope::text
    FROM azure_monitor_activity_log_alerts, jsonb_array_elements(properties->'scopes') AS scope
)
SELECT
    :'execution_time'                                                         AS execution_time,
    :'framework'                                                              AS framework,
    :'check_id'                                                               AS check_id,
    'Ensure that Activity Log Alert exists for Delete Network Security Group' AS title,
    fields.subscription_id                                                    AS subscription_id,
    fields.id                                                                 AS resource_id,
    CASE
        WHEN location = 'global'
         AND enabled
         AND equals = 'Microsoft.Network/networkSecurityGroups/delete'
         AND scopes.scope ~ '^"\/subscriptions\/[0-9a-f]{8}(-[0-9a-f]{4}){3}-[0-9a-f]{12}"$'
        THEN 'pass'
        ELSE 'fail'
    END                                                                       AS status
FROM fields JOIN scopes ON fields.id = scopes.id
WHERE field = 'operationName'
