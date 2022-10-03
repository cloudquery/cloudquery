--check if definition matches scopes
WITH assignable_scopes AS (
    SELECT
        cq_id,
        UNNEST(assignable_scopes) AS assignable_scope
    FROM azure_authorization_role_definitions v
),
meets_scopes AS (
    SELECT cq_id
    FROM assignable_scopes a
    WHERE a.assignable_scope = '/'
        OR a.assignable_scope = 'subscription'
    GROUP BY cq_id
),
--check if definition matches actions
definition_actions AS (
    SELECT
        role_definition_cq_id AS cq_id,
        UNNEST(actions) AS ACTION
    FROM azure_authorization_role_definition_permissions
),
meets_actions AS (
    SELECT cq_id
    FROM definition_actions
    WHERE "action" = '*'
)
SELECT
    d.subscription_id,
    d.id AS definition_id,
    d."name" AS definition_name
FROM azure_authorization_role_definitions d
    JOIN meets_actions a ON
        d.cq_id = a.cq_id
    JOIN meets_scopes s ON
        a.cq_id = s.cq_id;
