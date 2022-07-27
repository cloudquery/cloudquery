resource "aws_glue_classifier" "aws_glue_classifier1" {
  name = "${var.prefix}-glue-classifier"

  csv_classifier {
    allow_single_column    = false
    contains_header        = "PRESENT"
    delimiter              = ","
    disable_value_trimming = false
    header                 = ["example1", "example2"]
    quote_symbol           = "'"
  }
}
