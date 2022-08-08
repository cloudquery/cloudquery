resource "aws_glue_trigger" "example" {
  name = "${var.prefix}-glue-trigger"
  type = "CONDITIONAL"

  actions {
    job_name = aws_glue_job.example.name
  }

  predicate {
    conditions {
      job_name = aws_glue_job.example.name
      state    = "SUCCEEDED"
    }
  }
}