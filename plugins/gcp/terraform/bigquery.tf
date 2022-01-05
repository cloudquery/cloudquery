################################################################################
# Bigquery Helper
################################################################################
resource "google_bigquery_dataset_access" "access" {
  dataset_id    = module.bigquery.bigquery_dataset.dataset_id
  role          = "OWNER"
  user_by_email = module.service_accounts.email
}

################################################################################
# Bigquery Module
################################################################################

module "bigquery" {
  source  = "terraform-google-modules/bigquery/google"
  version = "~> 5.3"

  dataset_id                  = "${replace(local.prefix, "-", "")}bigquerydataset"
  dataset_name                = "${local.prefix}-bigquery-dataset"
  description                 = "CQ bigquery dataset"
  project_id                  = local.project
  location                    = "US"
  default_table_expiration_ms = null

  tables = [
    {
      table_id           = "${local.prefix}-bigquery-table"
      schema             = file("fixtures/bigquery/sample_bq_schema.json"),
      time_partitioning  = {
        type                     = "DAY"
        field                    = null
        require_partition_filter = false
        expiration_ms            = null
      },
      range_partitioning = null
      expiration_time    = null
      clustering         = ["fullVisitorId", "visitId"]
      labels             = local.labels
    }
  ]

  dataset_labels = local.labels

}