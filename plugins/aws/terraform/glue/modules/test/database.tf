resource "aws_glue_catalog_database" "aws_glue_catalog_database" {
  name = "${var.prefix}-glue-database"

  create_table_default_permission {
    permissions = ["SELECT"]

    principal {
      data_lake_principal_identifier = "IAM_ALLOWED_PRINCIPALS"
    }
  }

}

resource "aws_glue_catalog_table" "aws_glue_catalog_table" {
  name          = "${var.prefix}-glue-table"
  database_name = aws_glue_catalog_database.aws_glue_catalog_database.name

  table_type = "EXTERNAL_TABLE"

  parameters = {
    EXTERNAL              = "TRUE"
    "parquet.compression" = "SNAPPY"
  }

  storage_descriptor {
    location      = "s3://${aws_s3_bucket.aws_s3_bucket.bucket}/event-streams/my-stream"
    input_format  = "org.apache.hadoop.hive.ql.io.parquet.MapredParquetInputFormat"
    output_format = "org.apache.hadoop.hive.ql.io.parquet.MapredParquetOutputFormat"

    ser_de_info {
      name                  = "${var.prefix}-glue-stream"
      serialization_library = "org.apache.hadoop.hive.ql.io.parquet.serde.ParquetHiveSerDe"

      parameters = {
        "serialization.format" = 1
      }
    }

    columns {
      name = "my_string"
      type = "string"
    }

    columns {
      name = "my_double"
      type = "double"
    }

    columns {
      name    = "my_date"
      type    = "date"
      comment = ""
    }

    columns {
      name    = "my_bigint"
      type    = "bigint"
      comment = ""
    }

    columns {
      name    = "my_struct"
      type    = "struct<my_nested_string:string>"
      comment = ""
    }
  }
}

resource "aws_glue_partition" "aws_glue_partition" {
  database_name    = aws_glue_catalog_database.aws_glue_catalog_database.name
  table_name       = aws_glue_catalog_table.aws_glue_catalog_table.name
  partition_values = ["some-value"]
}

resource "aws_glue_partition_index" "aws_glue_partition_index" {
  database_name = aws_glue_catalog_database.aws_glue_catalog_database.name
  table_name    = aws_glue_catalog_table.aws_glue_catalog_table.name

  partition_index {
    index_name = "${var.prefix}-glue-partitition"
    keys       = ["my_column_1", "my_column_2"]
  }
}