resource "aws_glue_ml_transform" "aws_glue_ml_transform" {
  name     = "${var.prefix}-glue-ml-transform"
  role_arn = aws_iam_role.aws_iam_role.arn

  input_record_tables {
    database_name = aws_glue_catalog_table.aws_glue_catalog_table.database_name
    table_name    = aws_glue_catalog_table.aws_glue_catalog_table.name
  }

  parameters {
    transform_type = "FIND_MATCHES"

    find_matches_parameters {
      primary_key_column_name = "my_column_1"
    }
  }

}