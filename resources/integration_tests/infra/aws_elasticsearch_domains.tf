resource "aws_elasticsearch_domain" "aws_elasticsearch_domains_domain" {
  domain_name           = substr("elastic-domain-${var.test_suffix}", 0, 28)
  elasticsearch_version = "7.10"

  cluster_config {
    instance_type = "t3.small.elasticsearch"
  }

  snapshot_options {
    automated_snapshot_start_hour = 23
  }

  ebs_options {
    ebs_enabled = true
    volume_size = 10
  }

  tags = {
    Domain = "TestDomain"
  }
}