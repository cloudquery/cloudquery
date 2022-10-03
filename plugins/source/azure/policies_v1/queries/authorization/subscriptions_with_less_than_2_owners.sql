WITH owners_in_sub AS (SELECT a.subscription_id, COUNT(*) AS owners, d.id as id
                       FROM azure_authorization_role_assignments a
                                JOIN azure_authorization_role_definitions d ON a.properties_role_definition_id = d.id
                       WHERE role_name = 'Owner'
                         AND role_type = 'BuiltInRole' -- todo check if it checks only role or permissions list as well
                       GROUP BY d.id, a.subscription_id)

INSERT INTO azure_policy_results (execution_time, framework, check_id, title, subscription_id, resource_id, status)
SELECT :'execution_time'                                                  AS execution_time,
       :'framework'                                                       AS framework,
       :'check_id'                                                        AS check_id,
       'There should be more than one owner assigned to your subscription' AS title,
       subscription_id                                                    AS subscription_id,
       id                                                                 AS resource_id,
       CASE
           WHEN owners > 1
               THEN 'fail'
           ELSE 'pass'
           END                                                            AS status
FROM owners_in_sub