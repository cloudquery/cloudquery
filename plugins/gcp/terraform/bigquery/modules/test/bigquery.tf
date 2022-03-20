data "google_client_config" "current" {
}

data "google_client_openid_userinfo" "me" {
}


################################################################################
# Bigquery Helper
################################################################################
resource "google_bigquery_dataset_access" "access" {
  dataset_id    = module.bigquery.bigquery_dataset.dataset_id
  role          = "OWNER"
  user_by_email = data.google_client_openid_userinfo.me.email
}

################################################################################
# Bigquery Module
################################################################################

module "bigquery" {
  source  = "terraform-google-modules/bigquery/google"
  version = "~> 5.3"

  dataset_id                  = "${var.prefix}bigquerydataset"
  dataset_name                = "${var.prefix}-bigquery-dataset"
  description                 = "CQ bigquery dataset"
  project_id                  =  var.project_id
  location                    = "US"
  default_table_expiration_ms = null

  tables = [
    {
      table_id           = "cq-provider-table"
      schema             = file("${path.module}/sample_bq_schema.json"),
      time_partitioning  = {
        type                     = "DAY"
        field                    = null
        require_partition_filter = false
        expiration_ms            = null
      },
      range_partitioning = null
      expiration_time    = null
      clustering         = ["fullVisitorId", "visitId"]
      labels             = var.labels
    }
  ]

  dataset_labels = var.labels

}