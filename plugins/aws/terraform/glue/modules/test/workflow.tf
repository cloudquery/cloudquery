resource "aws_glue_workflow" "workflow" {
  name = "${var.prefix}-workflow"
  description = "Test workflow"
  max_concurrent_runs = 1
  tags = {
    "key" = "value"
  }
}
