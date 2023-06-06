INSERT INTO azure_policy_results
--check if definition matches scopes
WITH custom_roles AS (
    SELECT *
    FROM azure_authorization_role_definitions
    WHERE properties->>'type' = 'CustomRole'
),
assignable_scopes AS (
    SELECT
        _cq_id,
        scope AS assignable_scope
    FROM custom_roles,
         jsonb_array_elements_text(properties->'assignableScopes') scope
),
meets_scopes AS (
    SELECT
        _cq_id,
        bool_or(assignable_scope = '/' OR assignable_scope ~ '^\/subscriptions\/[0-9a-f]{8}(-[0-9a-f]{4}){3}-[0-9a-f]{12}$') AS has_wide_scope
    FROM assignable_scopes
    GROUP BY _cq_id
),
--check if definition matches actions
definition_actions AS (
    SELECT
        _cq_id,
        actions AS action
    FROM custom_roles,
         jsonb_array_elements(properties->'permissions') p,
         jsonb_array_elements_text(p->'actions') actions
),
meets_actions AS (
    SELECT
        _cq_id,
        bool_or("action" = '*') AS has_all_action
    FROM definition_actions
    GROUP BY _cq_id
)
SELECT
    :'execution_time'                                              AS execution_time,
    :'framework'                                                   AS framework,
    :'check_id'                                                    AS check_id,
    'Ensure That No Custom Subscription Administrator Roles Exist' AS title,
    subscription_id                                                AS subscription_id,
    id                                                             AS resource_id,
    CASE
        WHEN has_wide_scope AND has_all_action
        THEN 'fail'
        ELSE 'pass'
    END                                                            AS status
FROM custom_roles 
    JOIN meets_scopes USING (_cq_id) JOIN meets_actions USING (_cq_id)
