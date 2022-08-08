resource "aws_workspaces_directory" "directory" {
  directory_id = aws_directory_service_directory.service_directory.id
  subnet_ids   = slice(module.vpc.private_subnets, 0, 2)

  self_service_permissions {
    change_compute_type  = true
    increase_volume_size = true
    rebuild_workspace    = true
    restart_workspace    = true
    switch_running_mode  = true
  }

  workspace_access_properties {
    device_type_android    = "ALLOW"
    device_type_chromeos   = "ALLOW"
    device_type_ios        = "ALLOW"
    device_type_linux      = "DENY"
    device_type_osx        = "ALLOW"
    device_type_web        = "DENY"
    device_type_windows    = "ALLOW"
    device_type_zeroclient = "DENY"
  }

  workspace_creation_properties {
    enable_internet_access              = true
    enable_maintenance_mode             = true
    user_enabled_as_local_administrator = true
  }

  depends_on = [
    aws_iam_role_policy_attachment.workspaces_default_service_access,
    aws_iam_role_policy_attachment.workspaces_default_self_service_access
  ]
}

resource "aws_directory_service_directory" "service_directory" {
  name     = "${var.prefix}.cloudquery.local"
  password = random_password.this.result
  edition  = "Standard"
  type     = "MicrosoftAD"

  vpc_settings {
    vpc_id     = module.vpc.vpc_id
    subnet_ids = slice(module.vpc.private_subnets, 0, 2)
  }
}

##resource "aws_security_group" "workspace_directory_sg" {
##  name   = "${var.prefix}_workspace_directory_sg"
##  vpc_id = module.vpc.vpc_id
##
##  ingress {
##    from_port        = 0
##    to_port          = 0
##    protocol         = "-1"
##    cidr_blocks      = ["0.0.0.0/0"]
##    ipv6_cidr_blocks = ["::/0"]
##  }
##
##  egress {
##    from_port        = 0
##    to_port          = 0
##    protocol         = "-1"
##    cidr_blocks      = ["0.0.0.0/0"]
##    ipv6_cidr_blocks = ["::/0"]
##  }
##}
#
resource "random_password" "this" {
  length  = 40
  special = false
}

data "aws_iam_policy_document" "workspaces" {
  statement {
    actions = ["sts:AssumeRole"]

    principals {
      type        = "Service"
      identifiers = ["workspaces.amazonaws.com"]
    }
  }
}

resource "aws_iam_role" "workspaces_default_role" {
  name               = "workspaces_DefaultRole"
  assume_role_policy = data.aws_iam_policy_document.workspaces.json
}

resource "aws_iam_role_policy_attachment" "workspaces_default_service_access" {
  role       = aws_iam_role.workspaces_default_role.name
  policy_arn = "arn:aws:iam::aws:policy/AmazonWorkSpacesServiceAccess"
}

resource "aws_iam_role_policy_attachment" "workspaces_default_self_service_access" {
  role       = aws_iam_role.workspaces_default_role.name
  policy_arn = "arn:aws:iam::aws:policy/AmazonWorkSpacesSelfServiceAccess"
}

resource "aws_vpc_dhcp_options" "dns_resolver" {
  domain_name_servers = aws_directory_service_directory.service_directory.dns_ip_addresses
  domain_name         = "${var.prefix}.cloudquery.local"
  tags                = {
    Name = "${var.prefix}-dev-dhcp-option"
  }
}

resource "aws_vpc_dhcp_options_association" "dns_resolver" {
  vpc_id          = module.vpc.vpc_id
  dhcp_options_id = aws_vpc_dhcp_options.dns_resolver.id
}

module "vpc" {
  source  = "terraform-aws-modules/vpc/aws"
  version = "~> 3.0"

  name = "${var.prefix}-workspace-vpc"
  cidr = "10.101.0.0/18"

  azs             = ["us-east-1a", "us-east-1b", "us-east-1c"]
  private_subnets = ["10.101.3.0/24", "10.101.4.0/24", "10.101.5.0/24"]
  public_subnets  = ["10.101.6.0/24", "10.101.7.0/24", "10.101.8.0/24"]


  enable_nat_gateway     = true
  single_nat_gateway     = true
  one_nat_gateway_per_az = false
  enable_dns_hostnames   = true
  enable_dns_support     = true
}