-- SELECT *
-- FROM gcp_log_metric_filters
-- WHERE
--     enabled = TRUE
--     AND "filter" ~ '\s*(\s*protoPayload.serviceName\s*=\s*"cloudresourcemanager.googleapis.com"\s*)\s*AND\s*(\s*ProjectOwnership\s*OR\s*projectOwnerInvitee\s*)\s*OR\s*(\s*protoPayload.serviceData.policyDelta.bindingDeltas.action\s*=\s*"REMOVE"\s*AND\s*protoPayload.serviceData.policyDelta.bindingDeltas.role\s*=\s*"roles/owner"\s*)\s*OR\s*(\s*protoPayload.serviceData.policyDelta.bindingDeltas.action\s*=\s*"ADD"\s*AND\s*protoPayload.serviceData.policyDelta.bindingDeltas.role\s*=\s*"roles/owner"\s*)\s*'; -- noqa


INSERT INTO gcp_policy_results (resource_id, execution_time, framework, check_id, title, project_id, status)
SELECT "filter"                                                                                          AS resource_id,
       :'execution_time'::timestamp                                                                      AS execution_time,
       :'framework'                                                                                      AS framework,
       :'check_id'                                                                                       AS check_id,
       'Ensure log metric filter and alerts exist for project ownership assignments/changes (Automated)' AS title,
       project_id                                                                                        AS project_id,
       CASE
           WHEN
                       disabled = FALSE
                   AND "filter" ~
                       '\s*(\s*protoPayload.serviceName\s*=\s*"cloudresourcemanager.googleapis.com"\s*)\s*AND\s*(\s*ProjectOwnership\s*OR\s*projectOwnerInvitee\s*)\s*OR\s*(\s*protoPayload.serviceData.policyDelta.bindingDeltas.action\s*=\s*"REMOVE"\s*AND\s*protoPayload.serviceData.policyDelta.bindingDeltas.role\s*=\s*"roles/owner"\s*)\s*OR\s*(\s*protoPayload.serviceData.policyDelta.bindingDeltas.action\s*=\s*"ADD"\s*AND\s*protoPayload.serviceData.policyDelta.bindingDeltas.role\s*=\s*"roles/owner"\s*)\s*'
               THEN 'fail'
           ELSE 'pass'
           END                                                                                           AS status
FROM gcp_logging_metrics;