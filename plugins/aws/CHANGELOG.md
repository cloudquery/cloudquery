# Changelog

All notable changes to this provider will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [v0.5.1] - 2020-07-30
###### SDK Version: 0.3.1

### :rocket: Added
* Added new resource directconnect lags by [@James-Quigley](https://github.com/James-Quigley) [#122](https://github.com/cloudquery/cq-provider-aws/pull/122)

### :gear: Changed
* Updated SDK to version [0.3.1](https://github.com/cloudquery/cq-provider-sdk/blob/main/CHANGELOG.md#v031---2020-07-30)

### :spider: Fixed
* Fixed iam policies encoded policies parsing [#125](https://github.com/cloudquery/cq-provider-aws/pull/125)
* Fixed cognito_user_pools input [#126](https://github.com/cloudquery/cq-provider-aws/pull/126)


## [v0.5.0] - 2020-07-28
###### SDK Version: 0.3.0

### :rocket: Added

* Added a changelog :)

### :gear: Changed
* Upgraded to SDK Version [0.3.0](https://github.com/cloudquery/cq-provider-sdk/blob/main/CHANGELOG.md)
* **Breaking Change**: default CloudQuery "id" from `id` to `cq_id` [#41](https://github.com/cloudquery/cq-provider-sdk/pull/41)
* `aws_ec2_instance_licenses` table changed to `json` field `licenses`
* `aws_ecs_cluster_default_capacity_provider_strategies ` table changed to `json` field `default_capacity_provider_strategy`
* `aws_lambda_function_event_source_mapping_access_configurations  ` table changed to `json` field `source_access_configurations`
* `aws_rds_instance_status_infos  ` table changed to `json` field `status_infos`
* `aws_wafv2_rule_group_available_labels  ` table changed to `json` field `available_labels`
* `aws_wafv2_rule_group_consumed_labels  ` table changed to `json` field `consumed_labels`
* `aws_rds_cluster_db_cluster_option_group_memberships` table change to `json` field `db_cluster_option_group_memberships` [#118](https://github.com/cloudquery/cq-provider-aws/pull/118)

### :spider: Fixed
* Fixed AWS Debug flag will now write into log instead of stdout. [#119](https://github.com/cloudquery/cq-provider-aws/pull/119)

## [0.4.11] - 2020-07-15

Base version at which changelog was introduced.

###### SDK Version: 0.2.8

### Supported Resources
- accessanalyzer.analyzers
- apigateway.api_keys
- apigateway.client_certificates
- apigateway.domain_names
- apigateway.rest_apis
- apigateway.usage_plans
- apigateway.vpc_links
- apigatewayv2.apis
- apigatewayv2.domain_names
- apigatewayv2.vpc_links
- autoscaling.launch_configurations
- cloudfront.cache_policies
- cloudfront.distributions
- cloudtrail.trails
- cloudwatch.alarms
- cloudwatchlogs.filters
- cognito.identity_pools
- cognito.user_pools
- config.configuration_recorders
- config.conformance_packs
- directconnect.connections
- directconnect.gateways
- directconnect.virtual_gateways
- directconnect.virtual_interfaces
- ec2.byoip_cidrs
- ec2.customer_gateways
- ec2.ebs_volumes
- ec2.flow_logs
- ec2.images
- ec2.instances
- ec2.internet_gateways
- ec2.nat_gateways
- ec2.network_acls
- ec2.regional_config
- ec2.route_tables
- ec2.security_groups
- ec2.subnets
- ec2.transit_gateways
- ec2.vpc_endpoints
- ec2.vpc_peering_connections
- ec2.vpcs
- ec2.vpn_gateways
- ecr.repositories
- ecs.clusters
- efs.filesystems
- eks.clusters
- elasticbeanstalk.environments
- elasticsearch.domains
- elbv1.load_balancers
- elbv2.load_balancers
- elbv2.target_groups
- emr.clusters
- fsx.backups
- iam.accounts
- iam.groups
- iam.openid_connect_identity_providers
- iam.password_policies
- iam.policies
- iam.roles
- iam.saml_identity_providers
- iam.server_certificates
- iam.users
- iam.virtual_mfa_devices
- kms.keys
- lambda.functions
- lambda.layers
- mq.brokers
- organizations.accounts
- rds.certificates
- rds.clusters
- rds.db_subnet_groups
- rds.instances
- redshift.clusters
- redshift.subnet_groups
- route53.health_checks
- route53.hosted_zones
- route53.reusable_delegation_sets
- route53.traffic_policies
- s3.buckets
- sns.subscriptions
- sns.topics
- waf.rule_groups
- waf.rules
- waf.subscribed_rule_groups
- waf.web_acls
- wafv2.managed_rule_groups
- wafv2.rule_groups
- wafv2.web_acls

