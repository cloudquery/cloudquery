resource "aws_imagebuilder_image" "aws_ec2_images_image" {
  distribution_configuration_arn = aws_imagebuilder_distribution_configuration.aws_ec2_images_distribution_configuration.arn
  image_recipe_arn = aws_imagebuilder_image_recipe.aws_ec2_images_image_recipe.arn
  infrastructure_configuration_arn = aws_imagebuilder_infrastructure_configuration.aws_ec2_images_infrastructure_configuration.arn

  image_tests_configuration {
    image_tests_enabled = false
  }

  tags = {
    stage = "test"
  }
}

resource "aws_imagebuilder_distribution_configuration" "aws_ec2_images_distribution_configuration" {
  name = "ec2-images-dc-${var.test_prefix}${var.test_suffix}"

  distribution {

    ami_distribution_configuration {

      ami_tags = {
        CostCenter = "IT"
      }

      name = "aws-ec2-images-image-${var.test_prefix}${var.test_suffix}-{{ imagebuilder:buildDate }}"

      launch_permission {
        user_ids = [
          "123456789012"]
      }
    }

    region = data.aws_region.current.name
  }
}

resource "aws_imagebuilder_image_recipe" "aws_ec2_images_image_recipe" {
  block_device_mapping {
    device_name = "/dev/xvdb"

    ebs {
      delete_on_termination = true
      volume_size = 20
      volume_type = "gp2"
    }
  }

  component {
    component_arn = aws_imagebuilder_component.aws_ec2_images_hello_world.arn
  }

  name = "ec2-images-recipe-${var.test_prefix}${var.test_suffix}"
  parent_image = "arn:${data.aws_partition.current.partition}:imagebuilder:${data.aws_region.current.name}:aws:image/amazon-linux-2-x86/x.x.x"
  version = "0.1.1"

}

data "aws_partition" "current" {}

resource "aws_imagebuilder_infrastructure_configuration" "aws_ec2_images_infrastructure_configuration" {
  description = "ec2-images-ic-${var.test_prefix}${var.test_suffix}"
  instance_profile_name = aws_iam_instance_profile.aws_ec2_images_instance_profile.name
  instance_types = [
    "t2.nano",
    "t3.micro"]
  name = "ec2-images-ic-${var.test_prefix}${var.test_suffix}"
  terminate_instance_on_failure = true

  logging {
    s3_logs {
      s3_bucket_name = aws_s3_bucket.aws_ec2_images_bucket.bucket
      s3_key_prefix = "logs"
    }
  }

  tags = {
    foo = "bar"
  }
}

resource "aws_iam_instance_profile" "aws_ec2_images_instance_profile" {
  name = "ec2-images-ip-${var.test_prefix}${var.test_suffix}"
  role = aws_iam_role.aws_ec2_images_role.name
}

resource "aws_iam_role" "aws_ec2_images_role" {
  name = "ec2-images-role-${var.test_prefix}${var.test_suffix}"
  managed_policy_arns = [
    data.aws_iam_policy.aws_ec2_images_instance_profile_for_imagebuilder.arn,
    data.aws_iam_policy.aws_ec2_images_ecr_containers.arn,
    data.aws_iam_policy.aws_ec2_images_ssm.arn]
  inline_policy {
    name = "s3_logging"
    policy = data.aws_iam_policy_document.aws_ec2_images_inline_policy.json
  }

  assume_role_policy = data.aws_iam_policy_document.aws_ec2_images_instance-assume-role-policy.json
}

resource "aws_s3_bucket" "aws_ec2_images_bucket" {
  bucket = "ec2-images-logs-${var.test_prefix}${var.test_suffix}"
  acl = "private"

  force_destroy = true

  tags = {
    Name = "My bucket ${var.test_prefix}${var.test_suffix}"
    Environment = "test"
  }
}

resource "aws_imagebuilder_component" "aws_ec2_images_hello_world" {
  data = yamlencode({
    phases = [
      {
        name = "build"
        steps = [
          {
            action = "ExecuteBash"
            inputs = {
              commands = [
                "echo 'hello world'"]
            }
            name = "example-${var.test_prefix}${var.test_suffix}"
            onFailure = "Continue"
          }]
      }]
    schemaVersion = 1.0
  })
  name = "hello_world_${var.test_prefix}${var.test_suffix}"
  platform = "Linux"
  version = "1.0.0"
}

data "aws_iam_policy_document" "aws_ec2_images_instance-assume-role-policy" {
  statement {
    actions = [
      "sts:AssumeRole"]

    principals {
      type = "Service"
      identifiers = [
        "ec2.amazonaws.com"]
    }
  }
}


data "aws_iam_policy_document" "aws_ec2_images_inline_policy" {
  statement {
    actions = [
      "s3:PutObject"]
    resources = [
      "${aws_s3_bucket.aws_ec2_images_bucket.arn}/*",
    ]
  }
}

data "aws_iam_policy" "aws_ec2_images_instance_profile_for_imagebuilder" {
  name = "EC2InstanceProfileForImageBuilder"
}

data "aws_iam_policy" "aws_ec2_images_ecr_containers" {
  name = "EC2InstanceProfileForImageBuilderECRContainerBuilds"
}

data "aws_iam_policy" "aws_ec2_images_ssm" {
  name = "AmazonSSMManagedInstanceCore"
}
