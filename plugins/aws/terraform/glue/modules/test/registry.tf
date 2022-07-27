resource "aws_glue_registry" "aws_glue_registry" {
  registry_name = "${var.prefix}-glue-registry"
}

resource "aws_glue_schema" "aws_glue_schema" {
  schema_name       = "example"
  registry_arn      = aws_glue_registry.aws_glue_registry.arn
  data_format       = "AVRO"
  compatibility     = "NONE"
  schema_definition = "{\"type\": \"record\", \"name\": \"r1\", \"fields\": [ {\"name\": \"f1\", \"type\": \"int\"}, {\"name\": \"f2\", \"type\": \"string\"} ]}"
}