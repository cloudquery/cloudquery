resource "aws_codepipeline" "codepipeline" {
  name     = "${var.prefix}-codepipeline-pipeline"
  role_arn = aws_iam_role.codepipeline_role.arn

  artifact_store {
    location = aws_s3_bucket.codepipeline_artifact_bucket.bucket
    type     = "S3"

    encryption_key {
      id   = aws_kms_key.codepipeline_key.id
      type = "KMS"
    }
  }

  stage {
    name = "Source"

    action {
      name             = "Source"
      category         = "Source"
      owner            = "AWS"
      provider         = "CodeStarSourceConnection"
      version          = "1"
      output_artifacts = ["source_output"]

      configuration = {
        ConnectionArn    = aws_codestarconnections_connection.codestar_connection.arn
        FullRepositoryId = "cloudquery/cloudquery"
        BranchName       = "main"
      }
    }
  }

  stage {
    name = "Build"

    action {
      name             = "Build"
      category         = "Build"
      owner            = "AWS"
      provider         = "CodeBuild"
      input_artifacts  = ["source_output"]
      version          = "1"

      configuration = {
        ProjectName = "${var.prefix}-codebuild-project"
      }
    }
  }

}

resource "aws_codestarconnections_connection" "codestar_connection" {
  name          = "${var.prefix}-codestar-connection"
  provider_type = "GitHub"
}

resource "aws_s3_bucket" "codepipeline_artifact_bucket" {
  bucket = "${var.prefix}-codepipeline-artifact-bucket"
}

resource "aws_s3_bucket_acl" "codepipeline_artifact_bucket_acl" {
  bucket = aws_s3_bucket.codepipeline_artifact_bucket.id
  acl    = "private"
}

resource "aws_iam_role" "codepipeline_role" {
  name = "${var.prefix}-codepipeline-role"

  assume_role_policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Principal": {
        "Service": "codepipeline.amazonaws.com"
      },
      "Action": "sts:AssumeRole"
    }
  ]
}
EOF
}

resource "aws_iam_role_policy" "codepipeline_policy" {
  name = "${var.prefix}-codepipeline-policy"
  role = aws_iam_role.codepipeline_role.id

  policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect":"Allow",
      "Action": [
        "s3:GetObject",
        "s3:GetObjectVersion",
        "s3:GetBucketVersioning",
        "s3:PutObjectAcl",
        "s3:PutObject"
      ],
      "Resource": [
        "${aws_s3_bucket.codepipeline_artifact_bucket.arn}",
        "${aws_s3_bucket.codepipeline_artifact_bucket.arn}/*"
      ]
    },
    {
      "Effect": "Allow",
      "Action": [
        "codestar-connections:UseConnection"
      ],
      "Resource": "${aws_codestarconnections_connection.codestar_connection.arn}"
    },
    {
      "Effect": "Allow",
      "Action": [
        "codebuild:BatchGetBuilds",
        "codebuild:StartBuild"
      ],
      "Resource": "*"
    }
  ]
}
EOF
}

resource "aws_kms_key" "codepipeline_key" {
  description             = "Codepipeline KMS key"
  deletion_window_in_days = 10
}

resource "aws_iam_role" "codebuild_role" {
  name = "${var.prefix}-codebuild"

  assume_role_policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Principal": {
        "Service": "codebuild.amazonaws.com"
      },
      "Action": "sts:AssumeRole"
    }
  ]
}
EOF
}

resource "aws_codebuild_project" "codebuild_project" {
  name          = "${var.prefix}-codebuild-project"
  build_timeout = "5"
  service_role  = aws_iam_role.codebuild_role.arn

  artifacts {
    type = "CODEPIPELINE"
  }

  environment {
    compute_type                = "BUILD_GENERAL1_SMALL"
    image                       = "aws/codebuild/standard:1.0"
    type                        = "LINUX_CONTAINER"
    image_pull_credentials_type = "CODEBUILD"
  }

  source {
    type = "CODEPIPELINE"
  }

}