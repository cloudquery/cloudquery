################################################################################
# Cloud Functions Module
################################################################################

data "archive_file" "function_zip_archive" {
  type        = "zip"
  source_file = "${path.module}/hello_http.go"
  output_path = "./tmp/function.zip"
}

module "gcs_buckets" {
  source     = "terraform-google-modules/cloud-storage/google"
  version    = "~> 3.1.0"
  project_id = var.project_id
  names      = ["cloudfunctions"]
  prefix     = "cq-provider-${var.prefix}"

  labels = var.labels
}

resource "google_storage_bucket_object" "cloudfunctions_function_zip" {
  name   = "hello_http"
  bucket = module.gcs_buckets.name
  source = data.archive_file.function_zip_archive.output_path
}

resource "google_cloudfunctions_function" "cloudfunctions_function" {
  name        = "${var.prefix}-cloudfunctions"
  description = "Cloudquery cloudfunction test"
  runtime     = "go113"

  available_memory_mb   = 128
  source_archive_bucket = module.gcs_buckets.name
  source_archive_object = google_storage_bucket_object.cloudfunctions_function_zip.name
  trigger_http          = true
  entry_point           = "HelloHTTP"

  ingress_settings = "ALLOW_INTERNAL_ONLY"

  labels = var.labels

  environment_variables = {
    "ENV" = "dev"
  }

  build_environment_variables = {
    "ENV" = "dev"
  }
}
