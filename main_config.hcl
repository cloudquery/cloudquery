// Configuration AutoGenerated by CloudQuery CLI
cloudquery {
  plugin_directory = "./cq/providers"
  policy_directory = "./cq/policies"

  provider "aws" {
    version = "v0.11.0"
  }

#  provider "gcp" {
#    version = "latest"
#  }


  #  provider "azure" {
#    version = "latest"
#
#  }

  connection {
    dsn = "host=localhost user=postgres password=pass database=aws port=5432 sslmode=disable"
  }
}

#provider "gcp" {
#  configuration {}
#  resources = [
#    "compute.subnetworks"
#  ]
#
#}

#
#provider "azure" {
#  configuration {}
#
#  resources = []
#
#}
#
#provider "aws" {
#  configuration {
#    aws_debug = false
#  }
#  // list of resources to fetch
#  resources = []
#  // enables partial fetching, allowing for any failures to not stop full resource pull
#  enable_partial_fetch = true
#  // Limit provider to fetch only 5 resources at a given time
#  max_parallel_resource_fetch_limit = 5
#}
#
#
provider "aws" {
  alias = "aws-2"
  configuration {
    aws_debug = false
  }
  // list of resources to fetch
  resources = [
  "ec2.instances"
  ]
  // enables partial fetching, allowing for any failures to not stop full resource pull
#  enable_partial_fetch = true
}

#policy "aws-cis" {
#  source = "github.com/cloudquery-policies/aws//cis_v1.2.0"
#}
#
#policy "aws" {
#  source = "./my-policies/security1"
#}

modules {
  // drift configuration block
  drift "s3" {
    terraform {
      backend = "s3"
      bucket = "cq-provider-aws-tf"
      keys = [ "*" ]
    }
    provider "aws" {
      account_ids     = ["1234567891011"]
    }
  }
}