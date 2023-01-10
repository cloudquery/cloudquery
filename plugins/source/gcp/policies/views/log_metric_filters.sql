CREATE OR REPLACE VIEW gcp_log_metric_filters AS
SELECT
    gmap.project_id AS project_id,
    gmap.enabled AS enabled,
    glm."filter" AS filter
FROM gcp_monitoring_alert_policies gmap, JSONB_ARRAY_ELEMENTS(gmap.conditions) AS gmapc
    JOIN gcp_logging_metrics glm
        ON gmapc.value->'Condition'->'ConditionThreshold'->>'filter'
            LIKE '%metric.type="' || (glm.metric_descriptor->>'type') || '"%';
-- TODO check
