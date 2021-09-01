resource "aws_sns_topic" "sns_test" {
  name                        = "sns-tests-topic-${var.test_suffix}.fifo"
  fifo_topic                  = true
  content_based_deduplication = true
  display_name                = "${var.test_prefix}-${var.test_suffix}"
  delivery_policy             = <<EOF
      {
        "http": {
          "defaultHealthyRetryPolicy": {
            "minDelayTarget": 20,
            "maxDelayTarget": 20,
            "numRetries": 3,
            "numMaxDelayRetries": 0,
            "numNoDelayRetries": 0,
            "numMinDelayRetries": 0,
            "backoffFunction": "linear"
          },
          "disableSubscriptionOverrides": false,
          "defaultThrottlePolicy": {
            "maxReceivesPerSecond": 1
          }
        }
      }
  EOF
}

resource "aws_sqs_queue" "sns_test_queue" {
  name                        = "sns-tests-queue${var.test_suffix}.fifo"
  fifo_queue                  = true
  content_based_deduplication = true

}

resource "aws_sns_topic_subscription" "sns_test_subscription" {
  topic_arn = aws_sns_topic.sns_test.arn
  protocol  = "sqs"
  endpoint  = aws_sqs_queue.sns_test_queue.arn
}