resource "aws_glue_user_defined_function" "aws_glue_user_defined_function" {
  name          = "${var.prefix}-glue-user-function"
  catalog_id    = aws_glue_catalog_database.aws_glue_catalog_database.catalog_id
  database_name = aws_glue_catalog_database.aws_glue_catalog_database.name
  class_name    = "class"
  owner_name    = "owner"
  owner_type    = "GROUP"

  resource_uris {
    resource_type = "ARCHIVE"
    uri           = "uri"
  }
}