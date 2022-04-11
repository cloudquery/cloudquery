resource "aws_codepipeline_webhook" "codepipeline_webhook" {
  name            = "${var.prefix}-codepipeline-webhook"
  authentication  = "UNAUTHENTICATED"
  target_action   = "Source"
  target_pipeline = aws_codepipeline.codepipeline.name

  filter {
    json_path    = "$.ref"
    match_equals = "refs/heads/{Branch}"
  }

  tags = merge(
    {Name = "${var.prefix}-codepipeline-webhook"},
    var.tags
  )
}
