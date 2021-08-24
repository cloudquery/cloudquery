resource "aws_s3_bucket" "aws_cloudfront_distributions_bucket" {
  bucket = "b-${var.test_prefix}-${var.test_suffix}"
  acl = "private"
}

resource "aws_cloudfront_origin_access_identity" "aws_cloudfront_distributions_access_identity" {
  comment = "access identity ${var.test_prefix}-${var.test_suffix}"
}

data "aws_iam_policy_document" "aws_cloudfront_distributions_s3_policy" {
  statement {
    actions = [
      "s3:*"
    ]
    resources = [
      "${aws_s3_bucket.aws_cloudfront_distributions_bucket.arn}/content/*"]

    principals {
      type = "AWS"
      identifiers = [
        aws_cloudfront_origin_access_identity.aws_cloudfront_distributions_access_identity.iam_arn]
    }
  }
}

resource "aws_s3_bucket_policy" "aws_cloudfront_distributions_bucket_policy" {
  bucket = aws_s3_bucket.aws_cloudfront_distributions_bucket.id
  policy = data.aws_iam_policy_document.aws_cloudfront_distributions_s3_policy.json
}

resource "aws_cloudfront_distribution" "aws_cloudfront_distributions_distribution" {
  origin {
    domain_name = aws_s3_bucket.aws_cloudfront_distributions_bucket.bucket_regional_domain_name
    origin_id = "s3origin${var.test_prefix}-${var.test_suffix}"

    s3_origin_config {
      origin_access_identity = aws_cloudfront_origin_access_identity.aws_cloudfront_distributions_access_identity.cloudfront_access_identity_path
    }
  }

  enabled = true
  is_ipv6_enabled = true
  comment = "Some comment"
  default_root_object = "index.html"

  //  logging_config {
  //    include_cookies = false
  //    bucket = "mylogs.s3.amazonaws.com"
  //    prefix = "myprefix"
  //  }

  //  aliases = [
  //    "mysite.example.com",
  //    "yoursite.example.com"]

  default_cache_behavior {
    allowed_methods = [
      "DELETE",
      "GET",
      "HEAD",
      "OPTIONS",
      "PATCH",
      "POST",
      "PUT"]
    cached_methods = [
      "GET",
      "HEAD"]
    target_origin_id = "s3origin${var.test_prefix}-${var.test_suffix}"

    forwarded_values {
      query_string = false

      cookies {
        forward = "none"
      }
    }

    viewer_protocol_policy = "allow-all"
    min_ttl = 0
    default_ttl = 3600
    max_ttl = 86400
  }

  custom_error_response {
    error_code = 404
    response_code = 404
    response_page_path = "/custom_404.html"
  }

  # Cache behavior with precedence 0
  ordered_cache_behavior {
    path_pattern = "/content/immutable/*"
    allowed_methods = [
      "GET",
      "HEAD",
      "OPTIONS"]
    cached_methods = [
      "GET",
      "HEAD",
      "OPTIONS"]
    target_origin_id = "s3origin${var.test_prefix}-${var.test_suffix}"

    forwarded_values {
      query_string = false
      headers = [
        "Origin"]

      cookies {
        forward = "none"
      }
    }

    min_ttl = 0
    default_ttl = 86400
    max_ttl = 31536000
    compress = true
    viewer_protocol_policy = "redirect-to-https"
  }

  # Cache behavior with precedence 1
  ordered_cache_behavior {
    path_pattern = "/content/*"
    allowed_methods = [
      "GET",
      "HEAD",
      "OPTIONS"]
    cached_methods = [
      "GET",
      "HEAD"]
    target_origin_id = "s3origin${var.test_prefix}-${var.test_suffix}"

    forwarded_values {
      query_string = false

      cookies {
        forward = "none"
      }
    }

    min_ttl = 0
    default_ttl = 3600
    max_ttl = 86400
    compress = true
    viewer_protocol_policy = "redirect-to-https"
  }

  price_class = "PriceClass_200"

  restrictions {
    geo_restriction {
      restriction_type = "whitelist"
      locations = [
        "US",
        "CA",
        "GB",
        "DE"]
    }
  }

  tags = {
    Environment = "production"
  }

  viewer_certificate {
    cloudfront_default_certificate = true
  }
}