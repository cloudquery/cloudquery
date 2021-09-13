resource "google_storage_bucket" "bucket_func" {
  name = "bucket-func-${var.test_prefix}${var.test_suffix}"
}

resource "google_storage_bucket_object" "bucket_object_function" {
  name     = "helloworld-${var.test_prefix}${var.test_suffix}"
  bucket   = google_storage_bucket.bucket_func.name
  filename = data.archive_file.function_zip_inline.output_path
}

resource "google_cloudfunctions_function" "helloworld_function" {
  name        = "helloworld-${var.test_prefix}${var.test_suffix}"
  description = "My function ${var.test_prefix}${var.test_suffix}"
  runtime     = "go113"

  available_memory_mb   = 128
  source_archive_bucket = google_storage_bucket.bucket_func.name
  source_archive_object = google_storage_bucket_object.bucket_object_function.name
  trigger_http          = true
  entry_point           = "HelloHTTP"
}



data "archive_file" "function_zip_inline" {
  type        = "zip"
  output_path = "./tmp/function.zip"
  source {
    content  = <<EOF
// Package helloworld provides a set of Cloud Functions samples.
package helloworld

import (
        "encoding/json"
        "fmt"
        "html"
        "net/http"
)

// HelloHTTP is an HTTP Cloud Function with a request parameter.
func HelloHTTP(w http.ResponseWriter, r *http.Request) {
        var d struct {
                Name string `json:"name"`
        }
        if err := json.NewDecoder(r.Body).Decode(&d); err != nil {
                fmt.Fprint(w, "Hello, World!")
                return
        }
        if d.Name == "" {
                fmt.Fprint(w, "Hello, World!")
                return
        }
        fmt.Fprintf(w, "Hello, %s!", html.EscapeString(d.Name))
}

EOF
    filename = "hello_http.go"
  }
}