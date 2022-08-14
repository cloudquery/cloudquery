resource "aws_accessanalyzer_analyzer" "example" {
  analyzer_name = "${var.prefix}-accessanalyzer"
  tags = var.tags
}
