resource "google_monitoring_alert_policy" "gcp_monitoring_alerts_alert" {
  display_name = "alert-policies-${var.test_prefix}-${var.test_suffix}"
  combiner     = "OR"
  conditions {
    condition_threshold {
      filter             = "metric.type=\"logging.googleapis.com/user/${google_logging_metric.gcp_monitoring_alerts_alert_metric.name}\" resource.type=\"gke_container\""
      denominator_filter = "metric.type=\"logging.googleapis.com/user/${google_logging_metric.gcp_monitoring_alerts_alert_metric1.name}\" resource.type=\"gke_container\""
      comparison         = "COMPARISON_LT"
      threshold_value    = "0.6"
      duration           = "300s"

      aggregations {
        alignment_period = "60s"
      }

      denominator_aggregations {
        alignment_period = "60s"
      }
    }

    display_name = "StatefulSet has enough ready replicas"
  }

  user_labels = {
    foo = "bar"
  }

  depends_on = [
    google_logging_metric.gcp_monitoring_alerts_alert_metric,
    google_logging_metric.gcp_monitoring_alerts_alert_metric1,
  time_sleep.gcp_monitoring_alerts_wait_for_metric]
}


resource "time_sleep" "gcp_monitoring_alerts_wait_for_metric" {
  depends_on = [
  google_logging_metric.gcp_monitoring_alerts_alert_metric]

  create_duration = "10s"
}

resource "google_logging_metric" "gcp_monitoring_alerts_alert_metric" {
  name   = "alerts-metric-${var.test_prefix}-${var.test_suffix}"
  filter = "resource.type=gcs_bucket AND protoPayload.methodName=\"storage.setIamPermissions\""

  metric_descriptor {
    metric_kind = "DELTA"
    value_type  = "INT64"
  }
}


resource "google_logging_metric" "gcp_monitoring_alerts_alert_metric1" {
  name   = "alerts-metric1-${var.test_prefix}-${var.test_suffix}"
  filter = "resource.type=gcs_bucket"

  metric_descriptor {
    metric_kind = "DELTA"
    value_type  = "INT64"
  }
}