-- SELECT *
-- FROM gcp_log_metric_filters
-- WHERE
--     enabled = TRUE
--     AND "filter" ~ '\s*resource.type\s*=\s*gce_network\s*AND\s*protoPayload.methodName\s*=\s*"beta.compute.networks.insert"\s*OR\s*protoPayload.methodName\s*=\s*"beta.compute.networks.patch"\s*OR\s*protoPayload.methodName\s*=\s*"v1.compute.networks.delete"\s*OR\s*protoPayload.methodName\s*=\s*"v1.compute.networks.removePeering"\s*OR\s*protoPayload.methodName\s*=\s*"v1.compute.networks.addPeering"\s*'; -- noqa


INSERT INTO gcp_policy_results (resource_id, execution_time, framework, check_id, title, project_id, status)
SELECT "filter"                                                                                 AS resource_id,
       :'execution_time'::timestamp                                                             AS execution_time,
       :'framework'                                                                             AS framework,
       :'check_id'                                                                              AS check_id,
       'Ensure that the log metric filter and alerts exist for VPC network changes (Automated)' AS title,
       project_id                                                                               AS project_id,
       CASE
           WHEN
                       disabled = FALSE
                   AND "filter" ~
                       '\s*resource.type\s*=\s*gce_network\s*AND\s*protoPayload.methodName\s*=\s*"beta.compute.networks.insert"\s*OR\s*protoPayload.methodName\s*=\s*"beta.compute.networks.patch"\s*OR\s*protoPayload.methodName\s*=\s*"v1.compute.networks.delete"\s*OR\s*protoPayload.methodName\s*=\s*"v1.compute.networks.removePeering"\s*OR\s*protoPayload.methodName\s*=\s*"v1.compute.networks.addPeering"\s*'
               THEN 'fail'
           ELSE 'pass'
           END                                                                                  AS status
FROM gcp_logging_metrics;