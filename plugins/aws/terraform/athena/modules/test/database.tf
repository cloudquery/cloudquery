resource "aws_s3_bucket" "athenabucket" {
  bucket        = "${var.prefix}athenabkt${var.prefix}"
  force_destroy = true
}

resource "aws_athena_database" "aws_athena_database" {
  name          = "${var.prefix}athenadatabase"
  bucket        = aws_s3_bucket.athenabucket.bucket
  force_destroy = true
}

resource "aws_glue_catalog_table" "aws_athena_database_table" {
  name          = "${var.prefix}aws_athena_databasetable"
  database_name = aws_athena_database.aws_athena_database.name
}