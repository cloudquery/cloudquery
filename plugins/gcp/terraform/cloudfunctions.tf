################################################################################
# Cloud Functions Module
################################################################################

data "archive_file" "function_zip_archive" {
  type        = "zip"
  source_file = "${path.module}/fixtures/cloudfunction/hello_http.go"
  output_path = "./tmp/function.zip"
}

resource "google_storage_bucket_object" "cloudfunctions_function_zip" {
  name   = "${local.prefix}-cloudfunction-function-helloworld"
  bucket = module.gcs_buckets.name
  source = data.archive_file.function_zip_archive.output_path
}

resource "google_cloudfunctions_function" "cloudfunctions_function" {
  name        = "${local.prefix}-cloudfunction-function"
  description = "Cloudquery cloudfunction test"
  runtime     = "go113"

  available_memory_mb   = 128
  source_archive_bucket = module.gcs_buckets.name
  source_archive_object = google_storage_bucket_object.cloudfunctions_function_zip.name
  trigger_http          = true
  entry_point           = "HelloHTTP"

  ingress_settings = "ALLOW_INTERNAL_ONLY"

  labels = local.labels

  environment_variables = {
    "ENV" = "dev"
  }

  build_environment_variables = {
    "ENV" = "dev"
  }
}
