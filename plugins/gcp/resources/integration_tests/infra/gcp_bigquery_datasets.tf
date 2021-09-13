resource "google_bigquery_dataset" "gcp_bigquery_datasets_ds" {
  dataset_id                  = "bigquerydataset${var.test_suffix}"
  friendly_name               = "bigquery_dataset_${var.test_prefix}${var.test_suffix}"
  description                 = "This is a test description"
  location                    = "EU"
  default_table_expiration_ms = 3600000

  labels = {
    env = "default"
  }


  access {
    role          = "OWNER"
    user_by_email = google_service_account.service_account.email
  }

  access {
    role   = "READER"
    domain = "hashicorp.com"
  }


}

resource "google_bigquery_table" "gcp_bigquery_datasets_tb1" {
  dataset_id          = google_bigquery_dataset.gcp_bigquery_datasets_ds.dataset_id
  table_id            = "test"
  deletion_protection = false

  time_partitioning {
    type = "DAY"
  }

  labels = {
    env = "default"
  }

  schema = <<EOF
[
  {
    "name": "permalink",
    "type": "STRING",
    "mode": "NULLABLE",
    "description": "The Permalink"
  },
  {
    "name": "state",
    "type": "STRING",
    "mode": "NULLABLE",
    "description": "State where the head office is located"
  }
]
EOF

  depends_on = [
  google_bigquery_dataset.gcp_bigquery_datasets_ds]

}