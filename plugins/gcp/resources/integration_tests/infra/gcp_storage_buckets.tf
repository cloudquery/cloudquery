resource "google_storage_bucket" "gpc_storage_buckets_bucket" {
  name          = "gcp-storage-buckets-${var.test_prefix}-${var.test_suffix}"
  location      = "US"
  force_destroy = true

  lifecycle_rule {
    condition {
      age = 3
    }
    action {
      type = "Delete"
    }
  }


  website {
    main_page_suffix = "index.html"
    not_found_page   = "404.html"
  }
  cors {
    origin          = ["http://image-store.com"]
    method          = ["GET", "HEAD", "PUT", "POST", "DELETE"]
    response_header = ["*"]
    max_age_seconds = 3600
  }
}

resource "google_storage_bucket_acl" "gpc_storage_buckets_bucket_acl" {
  bucket = google_storage_bucket.gpc_storage_buckets_bucket.name

  default_acl = "projectPrivate"
  role_entity = [
    "READER:allAuthenticatedUsers",
  ]
}

resource "google_storage_default_object_acl" "image-store-default-acl" {
  bucket = google_storage_bucket.gpc_storage_buckets_bucket.name

  role_entity = [
    "READER:allAuthenticatedUsers",
  ]
}
