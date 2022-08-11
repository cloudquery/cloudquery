variable "test_suffix" {
  type = string
  default = ""
}



resource "random_id" "test_id" {
  keepers = {
    test_ids = "${var.test_suffix}"
  }

  byte_length = 8
}
