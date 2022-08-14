resource "aws_ssm_document" "ssm_document" {
  name          = "${var.prefix}-ssm-cq-provider"
  document_type = "Command"

  content = <<DOC
  {
    "schemaVersion": "1.2",
    "description": "Check ip configuration of a Linux instance.",
    "parameters": {

    },
    "runtimeConfig": {
      "aws:runShellScript": {
        "properties": [
          {
            "id": "0.aws:runShellScript",
            "runCommand": ["ifconfig"]
          }
        ]
      }
    }
  }
DOC

  tags = {
    Name = "docs"
  }
}

resource "aws_ssm_association" "test_associations" {
  name = "AmazonCloudWatch-ManageAgent"
  compliance_severity = "UNSPECIFIED"
  max_concurrency = 1
  max_errors = 10
  association_name = "${var.prefix}-ssm-assoc"
  targets {
    key    = "tag:Environment"
    values = ["Development"]
  }
}
