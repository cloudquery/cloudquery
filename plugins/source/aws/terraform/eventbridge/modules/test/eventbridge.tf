module "eventbridge" {
  source         = "terraform-aws-modules/eventbridge/aws"
  bus_name       = "${var.prefix}-bus"
  create_targets = false
  tags           = var.tags

  rules = {
    logs = {
      description   = "Capture log data"
      event_pattern = jsonencode({ "source" : ["my.app.logs"] })
    }
  }
}