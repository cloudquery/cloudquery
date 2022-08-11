resource "aws_cloudformation_stack" "network" {
  name = "${var.prefix}-networking-stack"
  count =  1
  parameters = {
    VPCCidr = "10.0.0.0/16"
  }

  on_failure = "DO_NOTHING"

  timeout_in_minutes = 30

  tags = var.tags
  template_body = <<STACK
{
  "Parameters" : {
    "VPCCidr" : {
      "Type" : "String",
      "Default" : "10.0.0.0/16",
      "Description" : "Enter the CIDR block for the VPC. Default is 10.0.0.0/16."
    }
  },
  "Resources" : {
    "myVpc": {
      "Type" : "AWS::EC2::VPC",
      "Properties" : {
        "CidrBlock" : { "Ref" : "VPCCidr" },
        "Tags" : [
          {"Key": "Name", "Value": "Primary_CF_VPC"}
        ]
      }
    }
  }
}
STACK
}