resource "aws_iam_openid_connect_provider" "default_openid" {
  url = "https://openidprovider${var.test_suffix}.com"

  client_id_list = [
    "client_id_list12312312312312123123123123"]

  thumbprint_list = [
    "client_id_list12312312312312123123123123"]

  tags = {
    "test" = "integration"
    "hello" = "world"
  }
}

