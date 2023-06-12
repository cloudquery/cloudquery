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
        scope
    FROM azure_monitor_activity_log_alerts, jsonb_array_elements_text(properties->'scopes') AS scope
),
conditions AS (
    SELECT
        fields.subscription_id AS subscription_id,
        fields.id AS id,
        scopes.scope AS scope,
        location = 'global'
            AND enabled
            AND equals = 'Microsoft.Network/networkSecurityGroups/delete'
            AND scopes.scope ~ '^\/subscriptions\/[0-9a-f]{8}(-[0-9a-f]{4}){3}-[0-9a-f]{12}$'
        AS condition
    FROM fields JOIN scopes ON fields.id = scopes.id
    WHERE field = 'operationName'
)
SELECT
    :'execution_time'                                                         AS execution_time,
    :'framework'                                                              AS framework,
    :'check_id'                                                               AS check_id,
    'Ensure that Activity Log Alert exists for Delete Network Security Group' AS title,
    subscription_id                                                           AS subscription_id,
    scope                                                                     AS resrouce_id,
    bool_or(condition)                                                        AS status
FROM conditions
GROUP BY subscription_id, scope