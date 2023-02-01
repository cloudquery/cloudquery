create or replace view view_azure_security_policy_parameters as
SELECT
    id,
    azure_policy_assignments.subscription_id,
    azure_policy_assignments."name",
    parameters.*,
    azure_policy_assignments.properties -> parameters.param ->> 'value' AS value
FROM
    azure_policy_assignments,
    jsonb_object_keys(azure_policy_assignments.properties) AS parameters ("param")
WHERE azure_policy_assignments."name" = 'SecurityCenterBuiltIn';