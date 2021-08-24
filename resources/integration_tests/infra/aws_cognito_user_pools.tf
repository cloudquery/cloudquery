resource "aws_cognito_user_pool" "cognito_user_pool" {
  name = "cognito_user_pool${var.test_prefix}-${var.test_suffix}"
}