-- SELECT *
-- FROM gcp_log_metric_filters
-- WHERE
--     enabled = TRUE
--     AND "filter" ~ '\s*resource.type\s*=\s*"iam_role"\s*AND\s*protoPayload.methodName\s*=\s*"google.iam.admin.v1.CreateRole"\s*OR\s*protoPayload.methodName\s*=\s*"google.iam.admin.v1.DeleteRole"\s*OR\s*protoPayload.methodName\s*=\s*"google.iam.admin.v1.UpdateRole"\s*'; -- noqa


INSERT INTO gcp_policy_results (resource_id, execution_time, framework, check_id, title, project_id, status)
SELECT "filter"                                                                                 AS resource_id,
       :'execution_time'::timestamp                                                             AS execution_time,
       :'framework'                                                                             AS framework,
       :'check_id'                                                                              AS check_id,
       'Ensure that the log metric filter and alerts exist for Custom Role changes (Automated)' AS title,
       project_id                                                                               AS project_id,
       CASE
           WHEN
                       disabled = FALSE
                   AND "filter" ~
                       '\s*resource.type\s*=\s*"iam_role"\s*AND\s*protoPayload.methodName\s*=\s*"google.iam.admin.v1.CreateRole"\s*OR\s*protoPayload.methodName\s*=\s*"google.iam.admin.v1.DeleteRole"\s*OR\s*protoPayload.methodName\s*=\s*"google.iam.admin.v1.UpdateRole"\s*'
               THEN 'fail'
           ELSE 'pass'
           END                                                                                  AS status
FROM gcp_logging_metrics;