resource "aws_db_event_subscription" "rds_event_subscription" {
  name      = "rds-event-sub-${var.test_prefix}-${var.test_suffix}"
  sns_topic = aws_sns_topic.sns_test2.arn

  source_type = "db-instance"
  source_ids  = [aws_db_instance.rds_db_instance.id]

  event_categories = [
    "failure",
  ]
}
