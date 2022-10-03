-- WITH project_policy_audit_configs AS (
--     SELECT
--         project_id,
--         jsonb_array_elements(p.policy -> 'auditConfigs') AS audit_config
--     FROM gcp_resource_manager_projects p
-- ),
-- log_types AS (
--     SELECT
--         project_id,
--         audit_config ->> 'service' AS "service",
--         jsonb_array_elements(audit_config -> 'auditLogConfigs') ->> 'logType' AS logs,
--         jsonb_array_elements(audit_config -> 'auditLogConfigs') ->> 'exemptedMembers' AS exempted
--     FROM project_policy_audit_configs
-- )
-- SELECT project_id, service, count(*)
-- FROM log_types
-- WHERE exempted IS NULL
--     AND logs IN ('DATA_READ', 'DATA_WRITE')
--     AND service = 'allServices'
-- GROUP BY project_id, service
-- --count(*) > 2 means DATA_READ and DATA_WRITE are there
-- HAVING count(*) = 2;

--
INSERT INTO gcp_policy_results (resource_id, execution_time, framework, check_id, title, project_id, status)
WITH project_policy_audit_configs AS (SELECT project_id,
                                             jsonb_array_elements(p.policy -> 'auditConfigs') AS audit_config
                                      FROM gcp_resource_manager_projects p),
     log_types AS (SELECT project_id,
                          audit_config ->> 'service'                                                    AS "service",
                          jsonb_array_elements(audit_config -> 'auditLogConfigs') ->> 'logType'         AS logs,
                          jsonb_array_elements(audit_config -> 'auditLogConfigs') ->> 'exemptedMembers' AS exempted
                   FROM project_policy_audit_configs)
SELECT service                                                                                                               AS resource_id,
       :'execution_time'::timestamp                                                                                          AS execution_time,
       :'framework'                                                                                                          AS framework,
       :'check_id'                                                                                                           AS check_id,
       'Ensure that Cloud Audit Logging is configured properly across all services and all users from a project (Automated)' AS title,
       "project_id"                                                                                                          AS project_id,
       CASE
           WHEN
                   exempted IS NULL
                   AND logs IN ('DATA_READ', 'DATA_WRITE')
                   AND service = 'allServices'
               THEN 'fail'
           ELSE 'pass'
           END                                                                                                               AS status
FROM log_types;