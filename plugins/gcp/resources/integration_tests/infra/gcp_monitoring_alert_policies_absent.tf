resource "google_monitoring_alert_policy" "gcp_monitoring_alerts_absent_alert" {
  display_name = "alert-absent-${var.test_prefix}-${var.test_suffix}"
  combiner     = "OR"

  conditions {
    display_name = "test condition1"
    condition_absent {
      filter   = "metric.type=\"logging.googleapis.com/user/${google_logging_metric.gcp_monitoring_alerts_alert_absent_metric.name}\" AND resource.type=\"metric\""
      duration = "120s"
      aggregations {
        alignment_period   = "120s"
        per_series_aligner = "ALIGN_RATE"
      }
    }
  }

  user_labels = {
    foo = "bar"
  }

  depends_on = [
    google_logging_metric.gcp_monitoring_alerts_alert_absent_metric,
  time_sleep.gcp_monitoring_alerts_wait_for_absent_metric]
}


resource "time_sleep" "gcp_monitoring_alerts_wait_for_absent_metric" {
  depends_on = [
  google_logging_metric.gcp_monitoring_alerts_alert_absent_metric]

  create_duration = "10s"
}

resource "google_logging_metric" "gcp_monitoring_alerts_alert_absent_metric" {
  name   = "alerts-absent-metric-${var.test_prefix}-${var.test_suffix}"
  filter = "resource.type=gcs_bucket AND protoPayload.methodName=\"storage.setIamPermissions\""

  metric_descriptor {
    metric_kind = "DELTA"
    value_type  = "INT64"
  }
}

