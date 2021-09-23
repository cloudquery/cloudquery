# A bucket to store logs in
resource "google_storage_bucket" "log-bucket" {
  name          = "log_bucket_sink_test"
  force_destroy = true
  retention_policy {
    retention_period = 123
    is_locked        = true
  }
}

# Our sink; this logs all activity related to our "my-logged-instance" instance
resource "google_logging_project_sink" "gcp_logging_sinks_sink" {
  name        = "logging-sink-${var.test_prefix}-${var.test_suffix}"
  description = "a description"
  destination = "storage.googleapis.com/${google_storage_bucket.log-bucket.name}"

  unique_writer_identity = true

  exclusions {
    name        = "ex-${var.test_prefix}-${var.test_suffix}"
    description = "Exclude logs from namespace-1 in k8s"
    filter      = "resource.type = k8s_container resource.labels.namespace_name=\"namespace-1\""
  }
}

resource "google_project_iam_binding" "log-writer" {
  role = "roles/storage.objectCreator"

  members = [
    google_logging_project_sink.gcp_logging_sinks_sink.writer_identity,
  ]
}