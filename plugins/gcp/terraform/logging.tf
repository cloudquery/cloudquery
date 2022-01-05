################################################################################
# Logging Module
################################################################################

resource "google_logging_metric" "gcp_logging_metric" {
  name   = "${local.prefix}logging-metric"
  filter = "protoPayload.methodName=\"cloudsql.instances.update\""

  metric_descriptor {
    metric_kind = "DELTA"
    value_type  = "DISTRIBUTION"
    unit        = "1"
    labels {
      key         = "mass"
      value_type  = "STRING"
      description = "amount of matter"
    }
    labels {
      key         = "sku"
      value_type  = "INT64"
      description = "Identifying number for item"
    }
    display_name = "My metric"
  }
  value_extractor = "EXTRACT(jsonPayload.request)"
  label_extractors = {
    "mass" = "EXTRACT(jsonPayload.request)"
    "sku"  = "EXTRACT(jsonPayload.id)"
  }
  bucket_options {
    linear_buckets {
      num_finite_buckets = 3
      width              = 1
      offset             = 1
    }
  }
}


resource "google_project_iam_binding" "log-writer" {
  role = "roles/storage.objectCreator"

  members = [
    google_logging_project_sink.gcp_logging_sink.writer_identity,
  ]
}

resource "google_logging_project_sink" "gcp_logging_sink" {
  name        = "${local.prefix}logging-sink"
  destination = "storage.googleapis.com/${module.gcs_buckets.name}"

  unique_writer_identity = true

  exclusions {
    name        = "${local.prefix}-logging-exclusion"
    description = "Exclude logs from namespace-1 in k8s"
    filter      = "resource.type = k8s_container resource.labels.namespace_name=\"namespace-1\""
  }
}