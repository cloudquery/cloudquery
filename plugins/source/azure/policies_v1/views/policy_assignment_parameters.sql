create or replace view view_azure_security_policy_parameters as
SELECT
    id,
    azure_resources_policy_assignments.subscription_id,
    azure_resources_policy_assignments."name",
    parameters.*,
    azure_resources_policy_assignments.parameters -> parameters.param ->> 'value' AS value
FROM
    azure_resources_policy_assignments,
    jsonb_object_keys(azure_resources_policy_assignments.parameters) AS parameters ("param")
WHERE azure_resources_policy_assignments."name" = 'SecurityCenterBuiltIn';