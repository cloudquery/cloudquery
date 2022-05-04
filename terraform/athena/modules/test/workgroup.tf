resource "aws_athena_workgroup" "aws_athena_workgroup" {
  name = "${var.prefix}_athena_workgroup"
  description = "test description"

  configuration {
    enforce_workgroup_configuration      = true
    publish_cloudwatch_metrics_enabled = true
  }
}

resource "aws_athena_named_query" "aws_athena_named_query" {
  name      = "${var.prefix}_athena_named_query"
  description = "test description"

  workgroup = aws_athena_workgroup.aws_athena_workgroup.id
  database  = aws_athena_database.aws_athena_database.name
  query     = "SELECT * FROM ${aws_athena_database.aws_athena_database.name} limit 10;"
}