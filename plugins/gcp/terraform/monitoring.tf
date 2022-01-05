################################################################################
# Monitoring Module - Helpers
################################################################################

resource "google_pubsub_topic" "notification_channel_topic" {
  name = "${local.prefix}-topic"

  labels = local.labels

}

resource "google_monitoring_notification_channel" "notification_channel" {
  display_name = "${local.prefix}-notification-channel"
  type         = "pubsub"
  labels       = {
    topic = google_pubsub_topic.notification_channel_topic.id
  }
}

################################################################################
# Monitoring Module
################################################################################

resource "google_monitoring_alert_policy" "alert_policy" {
  display_name          = "${local.prefix}-alert-policy"
  combiner              = "OR"
  conditions {
    display_name = "${local.prefix}-alert-policy-condition"
    condition_threshold {
      filter     = "metric.type=\"compute.googleapis.com/instance/disk/write_bytes_count\" AND resource.type=\"gce_instance\""
      duration   = "60s"
      comparison = "COMPARISON_GT"
      aggregations {
        alignment_period   = "60s"
        per_series_aligner = "ALIGN_RATE"
        group_by_fields    = ["resource.labels.instance_id"]
      }
      denominator_aggregations {
        alignment_period = "60s"
        group_by_fields  = ["resource.labels.instance_id"]
      }
    }
  }
  notification_channels = [google_monitoring_notification_channel.notification_channel.id]
  user_labels           = local.labels
}

resource "google_monitoring_alert_policy" "alert_policy_absent" {
  display_name = "${local.prefix}-alert-policy-absent"
  combiner     = "OR"
  conditions {
    display_name = "${local.prefix}-alert-policy-absent-condition"
    condition_absent {
      filter   = "metric.type=\"compute.googleapis.com/instance/disk/write_bytes_count\" AND resource.type=\"gce_instance\""
      duration = "3600s"
      aggregations {
        alignment_period   = "3600s"
        per_series_aligner = "ALIGN_RATE"
        group_by_fields    = ["resource.labels.instance_id"]
      }
    }
  }

  user_labels = local.labels
}