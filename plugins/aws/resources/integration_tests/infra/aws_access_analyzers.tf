resource "aws_accessanalyzer_analyzer" "analyzer" {
  analyzer_name = "analyzer-${var.test_prefix}${var.test_suffix}"
}
