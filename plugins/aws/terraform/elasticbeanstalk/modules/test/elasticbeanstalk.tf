// These are the policies that are attached to the elasticbeanstalk role when created in the console
resource "aws_iam_role_policy_attachment" "elastic-beanstalk-role-policy-attachment" {
  for_each = toset([
    "arn:aws:iam::aws:policy/AWSElasticBeanstalkWebTier", 
    "arn:aws:iam::aws:policy/AWSElasticBeanstalkMulticontainerDocker",
    "arn:aws:iam::aws:policy/AWSElasticBeanstalkWorkerTier",
  ])

  role       = "${aws_iam_role.elastic-beanstalk-role.name}"
  policy_arn = each.value
}

resource "aws_iam_role" "elastic-beanstalk-role" {
  name = "${var.prefix}-elastic-beanstalk-role"
  path = "/"

  assume_role_policy = jsonencode({
    "Version": "2008-10-17",
    "Statement": [
        {
            "Effect": "Allow",
            "Principal": {
                "Service": "ec2.amazonaws.com"
            },
            "Action": "sts:AssumeRole"
        }
    ]
  })
}

// This is usually created by the AWS console on first application
// We add it here so we ensure it exists before we create the environment
resource "aws_iam_instance_profile" "elastic-beanstalk-profile" {
  name = "${var.prefix}-elasticbeanstalk-ec2-role"
  role = aws_iam_role.elastic-beanstalk-role.name
}

resource "aws_s3_bucket_acl" "elastic-beanstalk-bucket-acl" {
  bucket = aws_s3_bucket.elastic-beanstalk-bucket.id
  acl    = "private"
}

resource "aws_s3_bucket_public_access_block" "elastic-beanstalk-bucket-public-access-block" {
  bucket = aws_s3_bucket.elastic-beanstalk-bucket.id

  block_public_acls       = true
  block_public_policy     = true
  ignore_public_acls      = true
  restrict_public_buckets = true
}

resource "aws_s3_bucket" "elastic-beanstalk-bucket" {
  // bucket names must me lowercase
  bucket = "${lower(var.prefix)}-elastic-beanstalk-bucket"
}

resource "aws_s3_object" "elastic-beanstalk-object" {
  bucket = aws_s3_bucket.elastic-beanstalk-bucket.id
  key    = "beanstalk/go-v1.zip"
  content = "go-v1.zip"
}

resource "aws_elastic_beanstalk_application_version" "elastic-beanstalk-version" {
  name        = "${var.prefix}-elastic-beanstalk-app-version"
  application = aws_elastic_beanstalk_application.example.name
  description = "application version created by terraform"
  bucket      = aws_s3_bucket.elastic-beanstalk-bucket.id
  key         = aws_s3_object.elastic-beanstalk-object.id
}

resource "aws_elastic_beanstalk_application" "example" {
  name        = "${var.prefix}-elastic-beanstalk-app"
  description = "Example Elastic Beanstalk Application"
}

resource "aws_elastic_beanstalk_environment" "env" {
  // bucket names must me lowercase
  name        = "${var.prefix}-elastic-beanstalk-env"
  application        = aws_elastic_beanstalk_application.example.name
  description = "Example Elastic Beanstalk Environment"
  solution_stack_name = "64bit Amazon Linux 2 v3.5.0 running Go 1"

  // See https://stackoverflow.com/questions/50806263/elastic-beanstalk-instance-profile-not-automatically-created-when-using-terraform
  // Required since the profile is not created automatically
  setting {
      namespace = "aws:autoscaling:launchconfiguration"
      name = "IamInstanceProfile"
      value = aws_iam_instance_profile.elastic-beanstalk-profile.name
  }
}
