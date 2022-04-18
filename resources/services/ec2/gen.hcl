service          = "aws"
output_directory = "."
add_generate     = true

resource "aws" "ec2" "egress_only_internet_gateways" {
  path = "github.com/aws/aws-sdk-go-v2/service/ec2/types.EgressOnlyInternetGateway"
  ignoreError "IgnoreAccessDenied" {
    path = "github.com/cloudquery/cq-provider-aws/client.IgnoreAccessDeniedServiceDisabled"
  }
  deleteFilter "AccountRegionFilter" {
    path = "github.com/cloudquery/cq-provider-aws/client.DeleteAccountRegionFilter"
  }
  multiplex "AwsAccountRegion" {
    path   = "github.com/cloudquery/cq-provider-aws/client.ServiceAccountRegionMultiplexer"
    params = ["ec2"]
  }


  options {
    primary_keys = ["arn"]
  }

  userDefinedColumn "account_id" {
    description = "The AWS Account ID of the resource."
    type        = "string"
    resolver "resolveAWSAccount" {
      path = "github.com/cloudquery/cq-provider-aws/client.ResolveAWSAccount"
    }
  }
  userDefinedColumn "region" {
    type        = "string"
    description = "The AWS Region of the resource."
    resolver "resolveAWSRegion" {
      path = "github.com/cloudquery/cq-provider-aws/client.ResolveAWSRegion"
    }
  }

  column "egress_only_internet_gateway_id" {
    rename = "id"
  }

  column "attachments" {
    type              = "json"
    generate_resolver = true
    description       = "Information about the attachment of the egress-only internet gateway."
  }

  column "tags" {
    type              = "json"
    generate_resolver = false
    description       = "The tags assigned to the egress-only internet gateway."
  }

  userDefinedColumn "arn" {
    type        = "string"
    description = "The Amazon Resource Name (ARN) for the egress-only internet gateway."
    generate_resolver = false
  }

}

resource "aws" "ec2" "network_interfaces" {
  path = "github.com/aws/aws-sdk-go-v2/service/ec2/types.NetworkInterface"
  ignoreError "IgnoreAccessDenied" {
    path = "github.com/cloudquery/cq-provider-aws/client.IgnoreAccessDeniedServiceDisabled"
  }
  deleteFilter "AccountRegionFilter" {
    path = "github.com/cloudquery/cq-provider-aws/client.DeleteAccountRegionFilter"
  }
  multiplex "AwsAccountRegion" {
    path   = "github.com/cloudquery/cq-provider-aws/client.ServiceAccountRegionMultiplexer"
    params = ["ec2"]
  }

  options {
    primary_keys = ["arn"]
  }

  userDefinedColumn "account_id" {
    description = "The AWS Account ID of the resource."
    type        = "string"
    resolver "resolveAWSAccount" {
      path = "github.com/cloudquery/cq-provider-aws/client.ResolveAWSAccount"
    }
  }
  userDefinedColumn "region" {
    type        = "string"
    description = "The AWS Region of the resource."
    resolver "resolveAWSRegion" {
      path = "github.com/cloudquery/cq-provider-aws/client.ResolveAWSRegion"
    }
  }

  column "network_interface_id" {
    rename = "id"
  }

  column "tag_set" {
    skip = true
  }

  column "groups" {
    type              = "json"
    generate_resolver = true
    description       = "The tags assigned to the egress-only internet gateway."
  }

  column "groups" {
    type              = "json"
    generate_resolver = true
    description       = "Describes a security group."
  }

  column "ipv4_prefixes" {
    type              = "stringArray"
    generate_resolver = true
    description       = "Describes an IPv4 prefix."
  }

  column "ipv6_addresses" {
    type              = "stringArray"
    generate_resolver = true
    description       = "Describes an IPv6 address associated with a network interface."
  }

  column "ipv6_prefixes" {
    type              = "stringArray"
    generate_resolver = true
    description       = "Describes the IPv6 prefix."
  }

  userDefinedColumn "arn" {
    type        = "string"
    description = "The Amazon Resource Name (ARN) for the egress-only internet gateway."
    generate_resolver = false
  }

  userDefinedColumn "tags" {
    type        = "json"
    description = "Any tags assigned to the network interface."
    generate_resolver = false
  }

}

resource "aws" "ec2" "hosts" {
  path = "github.com/aws/aws-sdk-go-v2/service/ec2/types.Host"
  ignoreError "IgnoreAccessDenied" {
    path = "github.com/cloudquery/cq-provider-aws/client.IgnoreAccessDeniedServiceDisabled"
  }
  deleteFilter "AccountRegionFilter" {
    path = "github.com/cloudquery/cq-provider-aws/client.DeleteAccountRegionFilter"
  }
  multiplex "AwsAccountRegion" {
    path   = "github.com/cloudquery/cq-provider-aws/client.ServiceAccountRegionMultiplexer"
    params = ["ec2"]
  }


  options {
    primary_keys = ["arn"]
  }

  userDefinedColumn "account_id" {
    description = "The AWS Account ID of the resource."
    type        = "string"
    resolver "resolveAWSAccount" {
      path = "github.com/cloudquery/cq-provider-aws/client.ResolveAWSAccount"
    }
  }
  userDefinedColumn "region" {
    type        = "string"
    description = "The AWS Region of the resource."
    resolver "resolveAWSRegion" {
      path = "github.com/cloudquery/cq-provider-aws/client.ResolveAWSRegion"
    }
  }

  column "host_id" {
    rename = "id"
  }

  column "host_properties" {
    skip_prefix = true
  }

  column "available_capacity_available_v_cpus" {
    rename = "available_capacity_available_vcpus"
  }

  column "total_v_cpus" {
    rename = "total_vcpus"
  }

  column "host_reservation_id" {
    rename = "reservation_id"
  }

  column "tags" {
    type              = "json"
    generate_resolver = false
    description       = "Any tags assigned to the Dedicated Host."
  }

  column "available_capacity" {
    skip_prefix = true
  }

  column "available_v_cpus" {
    rename = "available_vcpus"
  }

  userDefinedColumn "arn" {
    type        = "string"
    description = "The Amazon Resource Name (ARN) for the dedicated host."
    generate_resolver = false
  }

}