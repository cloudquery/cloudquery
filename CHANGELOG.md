# Changelog

All notable changes to this provider will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

<!-- 
## Unreleased

### 🚀 Added

### :gear: Changed

### :spider: Fixed

### 💥 Breaking Changes

-->

## [v0.8.5] - 2022-01-03
### :spider: Fixed
* Fixed PK error with elasticsearch domains [#384](https://github.com/cloudquery/cq-provider-aws/pull/384).
### :gear: Changed
* Updated to SDK version v0.6.1

## [v0.8.4] - 2021-12-23
###### SDK Version: 0.5.7

### 💥 Breaking Changes
* Renamed columns of `aws_sagemaker_model_containers` image_config_repository_auth_config_repository_credentials_provider_arn -> image_config_repository_auth_config_repo_creds_provider_arn [#356](https://github.com/cloudquery/cq-provider-aws/pull/356).

### 🚀 Added
* Added how to use AWS provider with MFA enabled roles [#351](https://github.com/cloudquery/cq-provider-aws/pull/351) resolves [#35](https://github.com/cloudquery/cq-provider-aws/issues/35).
* Added to github test to run fetch on PR to main [#359](https://github.com/cloudquery/cq-provider-aws/pull/359).
* Passed version to provider struct so it will be passed in protocol [#370](https://github.com/cloudquery/cq-provider-aws/pull/370).

### :gear: Changed
* Check unsupported regions for service and remove them from multiplexer to reduce unnecessary calls [#373](https://github.com/cloudquery/cq-provider-aws/pull/373).
* Notify failures on warning messages in sanity tests [#346](https://github.com/cloudquery/cq-provider-aws/pull/352).
* Upgraded to SDK Version [v0.5.7](https://github.com/cloudquery/cq-provider-sdk/blob/main/CHANGELOG.md).
* Improved docs for contributors [#353](https://github.com/cloudquery/cq-provider-aws/pull/353).
* ignore AccessDeniedException completely on `OrganizationsAccounts` [#369](https://github.com/cloudquery/cq-provider-aws/pull/369).

### :spider: Fixed
* Ensure maximum table name length [#356](https://github.com/cloudquery/cq-provider-aws/pull/356).
* Fixed PK error with secret manager resource [#361](https://github.com/cloudquery/cq-provider-aws/pull/361) closed [#354](https://github.com/cloudquery/cq-provider-aws/issues/354).
* Fixed region filtering [#367](https://github.com/cloudquery/cq-provider-aws/pull/367).
* Fixed bucket missing tag incorrect warning message [#372](https://github.com/cloudquery/cq-provider-aws/pull/372).


## [v0.8.3] - 2021-12-15
###### SDK Version: 0.5.5

### 🚀 Added
* Added Contribution [guide](https://github.com/cloudquery/cq-provider-aws/blob/main/.github/CONTRIBUTING.md) [#335](https://github.com/cloudquery/cq-provider-aws/pull/335).
* extended logging of aws authorization error [#347](https://github.com/cloudquery/cq-provider-aws/pull/347) fixes [#245](https://github.com/cloudquery/cq-provider-aws/issues/245).

### :gear: Changed
* renames column of `aws_ec2_subnets` from `subnet_arn` to `arn` [#346](https://github.com/cloudquery/cq-provider-aws/pull/346).
* Upgraded to SDK Version [v0.5.5](https://github.com/cloudquery/cq-provider-sdk/blob/main/CHANGELOG.md).

### :spider: Fixed
* Fixed Call to ListAccounts in a non-org user  [#337](https://github.com/cloudquery/cq-provider-aws/pull/337) [#349](https://github.com/cloudquery/cq-provider-aws/pull/349).


## [v0.8.2] - 2021-12-09
###### SDK Version: 0.5.4

### 🚀 Added
* Added `aws_rds_db_parameter_groups`, `aws_rds_cluster_parameter_groups` and `aws_rds_db_security_groups` resources [#333](https://github.com/cloudquery/cq-provider-aws/pull/333).

### :gear: Changed
* Upgraded to SDK Version [v0.5.4](https://github.com/cloudquery/cq-provider-sdk/blob/main/CHANGELOG.md).

### :spider: Fixed
* Fixed PK violation in `aws_lambda_runtimes` [#334](https://github.com/cloudquery/cq-provider-aws/issues/343).
* Removed unnecessary Mutliplex and DeleteFilter in `aws_elbv2_listeners` [#342](https://github.com/cloudquery/cq-provider-aws/pull/342).
* Fixed [#242](https://github.com/cloudquery/cq-provider-aws/issues/242)  using disabled detection with hardcoded "us-east-1" [#341](https://github.com/cloudquery/cq-provider-aws/pull/341).

## [v0.8.1] - 2021-12-08
###### SDK Version: 0.5.3

### :rocket: Added
* Added `aws_lambda_runtimes` [#338](https://github.com/cloudquery/cq-provider-aws/pull/338).
* Added DAX and DynamoDB tables [#324](https://github.com/cloudquery/cq-provider-aws/pull/324).
* Renamed `aws_directconnect_gateways` table columns [#300](https://github.com/cloudquery/cq-provider-aws/pull/300).
    - `aws_dynamodb_tables`.
    - `aws_dax_clusters`.
    - `aws_applicationautoscaling_policies`.

### :gear: Changed
* Renamed in aws cloudfront `aws_cloudfront_distribution_default_cache_behavior_lambda_functions` -> `aws_cloudfront_distribution_default_cache_behavior_functions` [#336](https://github.com/cloudquery/cq-provider-aws/pull/336).

### :spider: Fixed
* Fixed call to ListAccounts in non-org user [#337](https://github.com/cloudquery/cq-provider-aws/pull/337).

## [v0.8.0] - 2021-12-06
###### SDK Version: 0.5.3

### 💥 Breaking Changes
* Renamed columns of `aws_cloudfront_distributions` behaviour -> behavior [#207](https://github.com/cloudquery/cq-provider-aws/pull/207).
* Table `aws_emr_clusters` is dropped and recreated in this version.

### :rocket: Added
* Added `aws_ecs_task_definitions` resource [#317](https://github.com/cloudquery/cq-provider-aws/pull/317)
* Added `aws_ssm_documents`, `aws_ssm_instances` resources [#307](https://github.com/cloudquery/cq-provider-aws/pull/307)
* Added `aws_ec2_instances` columns: `state_transition_reason_time`, `boot_mode`, `metadata_options_http_protocol_ipv6`, `ipv4_prefixes`, `ipv6_prefixes` [#325](https://github.com/cloudquery/cq-provider-aws/pull/325)
* Added logging configuration to WAF & WAFv2 [#315](https://github.com/cloudquery/cq-provider-aws/pull/315)
* Enhanced `aws_config_recorders` with status information [#301](https://github.com/cloudquery/cq-provider-aws/pull/301)
* Added `aws_acm_certificates` [#313](https://github.com/cloudquery/cq-provider-aws/pull/313)
* Added `logging_status` to `aws_redshift_clusters` & `aws_redshift_cluster_parameters` relation table [#319](https://github.com/cloudquery/cq-provider-aws/pull/319)
* Added support for AWS SecretManager secrets [#312](https://github.com/cloudquery/cq-provider-aws/pull/321)
* Added support for Elasticbeanstalk Applications `aws_elasticbeanstalk_applications` [#316](https://github.com/cloudquery/cq-provider-aws/pull/316)
* Added support for RDS event subscriptions `aws_rds_event_subscriptions` [#322](https://github.com/cloudquery/cq-provider-aws/pull/322)
* Added full info for EMR clusters [#318](https://github.com/cloudquery/cq-provider-aws/pull/318)
* Added lacking columns to `aws_cloudfront_distributions` [#207](https://github.com/cloudquery/cq-provider-aws/pull/207)
* Added makefile for easy execution of cq-provider-aws [#330](https://github.com/cloudquery/cq-provider-aws/pull/330)

### :gear: Changed
* Upgraded to SDK Version [v0.5.3](https://github.com/cloudquery/cq-provider-sdk/blob/main/CHANGELOG.md).
* Updated client default `max_retries` and `backoff` [#332](https://github.com/cloudquery/cq-provider-aws/pull/330)


## [v0.7.0] - 2021-11-29
###### SDK Version: 0.5.2

### 💥 Breaking Changes
* Renamed `aws_directconnect_gateways` table columns [#300](https://github.com/cloudquery/cq-provider-aws/pull/300).
    - "aws_directconnect_gateways" rename from `direct_connect_gateway_state` to `state`.
    - "aws_directconnect_gateways" rename from `direct_connect_gateway_name` to `name`.
    - "aws_directconnect_gateways" drop column `directconnect_gateway_id`.
    - "aws_directconnect_gateway_associations" rename from `directconnect_gateway_cq_id` to `gateway_cq_id`.
    - "aws_directconnect_gateway_associations" rename from `directconnect_gateway_id` to `gateway_id`.

### :rocket: Added
* Added `KmsMasterKeyId` column to `sns` resource [#309](https://github.com/cloudquery/cq-provider-aws/pull/309).
* Added support for ECS cluster dependencies [#267](https://github.com/cloudquery/cq-provider-aws/pull/267) fixes [#260](https://github.com/cloudquery/cq-provider-aws/issues/260).
  -  Added `aws_ecs_cluster_services` relation of `aws_ecs_clusters`.
  -  Added `aws_ecs_cluster_container_instances` relation of `aws_ecs_clusters`.
* Added support for AWS sagemaker resources [#291](https://github.com/cloudquery/cq-provider-aws/pull/291).
* Added support for SSM instance and compliance Items `aws_ssm_instances` and `aws_ssm_instance_compliance_items` [#299](https://github.com/cloudquery/cq-provider-aws/pull/299).
* Added DMS Replication instance resources `aws_dms_instances` [#280](https://github.com/cloudquery/cq-provider-aws/pull/280)
* Added RDS Cluster, DB snapshots, attributes [#287](https://github.com/cloudquery/cq-provider-aws/pull/287).
* Added support for `aws_regions` table allowing to view all enabled regions for an account [#293](https://github.com/cloudquery/cq-provider-aws/pull/293).
* Added support for `aws_guardduty_detectors` [#286](https://github.com/cloudquery/cq-provider-aws/pull/286) resource.
* Added `aws_ec2_ebs_snapshots` [#283](https://github.com/cloudquery/cq-provider-aws/pull/283) and `aws_ec2_eips` [#284](https://github.com/cloudquery/cq-provider-aws/pull/284) resources.
* Added ARN column for security group resource [#278](https://github.com/cloudquery/cq-provider-aws/issues/277).
* Added `aws_codebuild_projects` resource [#270](https://github.com/cloudquery/cq-provider-aws/issues/270).
* Added  `aws_autoscaling_groups` resource [#268](https://github.com/cloudquery/cq-provider-aws/issues/268).
* Added AWS EMR block public access [#269](https://github.com/cloudquery/cq-provider-aws/pull/269) Closes [#249](https://github.com/cloudquery/cq-provider-aws/issues/249).
* Improved AWS Assume Role documentation [#264](https://github.com/cloudquery/cq-provider-aws/pull/264).
* Added Support S3 Account settings [#285](https://github.com/cloudquery/cq-provider-aws/pull/285) Fixes [#282](https://github.com/cloudquery/cq-provider-aws/issues/282).
* Stored data showing account password policy doesn't exist [#281](https://github.com/cloudquery/cq-provider-aws/issues/281).
  
  
### :gear: Changed
* Upgraded to SDK Version [v0.5.2](https://github.com/cloudquery/cq-provider-sdk/blob/main/CHANGELOG.md).
* Improved docs generation to remove stale docs [#294](https://github.com/cloudquery/cq-provider-aws/pull/294).


### :spider: Fixed
* Fixed `aws_ec2_ebs_volumes` pagination during fetch [#279](https://github.com/cloudquery/cq-provider-aws/issues/279).
* Fixed ignore error logic [#274](https://github.com/cloudquery/cq-provider-aws/pull/274) fixes [#265](https://github.com/cloudquery/cq-provider-aws/issues/265).


## [v0.6.4] - 2021-11-16
###### SDK Version: 0.5.1

### :spider: Fixed
* Fixed trails not fetched if they have no tags defined [#258](https://github.com/cloudquery/cq-provider-aws/issues/258).

### :rocket: Added
* Added Elbv2 Listeners resource [#256](https://github.com/cloudquery/cq-provider-aws/issues/256).


## [v0.6.3] - 2021-11-11
###### SDK Version: 0.5.1

### :spider: Fixed
* Fixed [#164](https://github.com/cloudquery/cq-provider-aws/issues/164) apigateway_api_keys does not return key value

### :rocket: Added
* Added Route53 Domains resource

## [v0.6.2] - 2021-11-03
###### SDK Version: 0.5.1

### :spider: Fixed
* Fixed [#241](https://github.com/cloudquery/cq-provider-aws/issues/241) Failed to fetch ApiGatewayV2: GetDomainNames
* Fixed [#236](https://github.com/cloudquery/cq-provider-aws/issues/236) error in `aws_cloud_trails` get tags request because ARNs in request were from different regions

## [v0.6.1] - 2021-10-29
###### SDK Version: 0.5.0
* added skip GetFunctionCodeSigningConfig for container functions [#230](https://github.com/cloudquery/cq-provider-aws/pull/230)

## [v0.6.0] - 2021-10-26
###### SDK Version: 0.5.0

### :rocket: Added
* Added ignore `AWSOrganizationsNotInUseException` error to ignore error filter. [#213](https://github.com/cloudquery/cq-provider-aws/pull/213)
* Added ignore error when regions are disabled for a specific service [#210](https://github.com/cloudquery/cq-provider-aws/issues/210)
* Increased testing coverage with assume role fetch workflow [#218](https://github.com/cloudquery/cq-provider-aws/pull/218)
* Added `vpc_id` to `emr_clusters` resource [#221](https://github.com/cloudquery/cq-provider-aws/issues/221)

### :spider: Fixed
* Fixed [#157](https://github.com/cloudquery/cq-provider-aws/issues/157) tags on multiple resources kms, gateways_v2_vpc_links, elbv2 groups/balancers, ecs clusters, directconnect gateways, cloudtrail trails, elasticsearch_domains, elasticbeanstalk_environments  [#191](https://github.com/cloudquery/cq-provider-aws/pull/191)
* Fixed duplicate of `id` field for `aws_apigatewayv2_vpc_links` - removed `vpc_link_id` field
* Fixed duplicate of `region` field for `aws_cloudtrail_trails` - removed `home_region` field
* Fixed naming according to convention `aws_elasticbeanstalk_environments`:`environment_name` -> `name`, `aws_kms_keys`: `key_id` -> `id`
* Fixed [Web ACL (WAF) attachment](https://github.com/cloudquery/cq-provider-aws/issues/209)
* Fixed violation in `aws_apigateway_domain_name_base_path_mappings_pk` [#222](https://github.com/cloudquery/cq-provider-aws/issues/222)

### :gear: Changed
* Upgraded to SDK Version [v0.5.0](https://github.com/cloudquery/cq-provider-sdk/blob/main/CHANGELOG.md#v050---2021-10-21)

## [v0.5.16] - 2021-10-07
###### SDK Version: v0.4.9

### :rocket: Added
* Added for SQS queues resource [#202](https://github.com/cloudquery/cq-provider-aws/issues/202).

### :gear: Changed
* Upgraded to SDK Version [v0.4.9](https://github.com/cloudquery/cq-provider-sdk/blob/main/CHANGELOG.md)

### :spider: Fixed
* Fixed migration tests [#203](https://github.com/cloudquery/cq-provider-aws/pull/203)

## [v0.5.15] - 2021-10-03
###### SDK Version: v0.4.7

### :gear: Changed
* Upgraded to SDK Version [v0.4.7](https://github.com/cloudquery/cq-provider-sdk/blob/main/CHANGELOG.md#v047---2021-09-23)

## [v0.5.14] - 2021-09-13
###### SDK Version: v0.4.4

### :spider: Fixed
* Fixed kms key fetching when provider tried to fetch aws managed keys rotation properties [#168](https://github.com/cloudquery/cq-provider-aws/pull/168)

### :gear: Changed
* Upgraded to SDK Version [v0.4.4](https://github.com/cloudquery/cq-provider-sdk/blob/main/CHANGELOG.md#v044---2021-09-13)

## [v0.5.13] - 2021-09-09
###### SDK Version: v0.4.3

### :spider: Fixed
remove region multiplex on web acl resource[#188](https://github.com/cloudquery/cq-provider-aws/pull/188)



## [v0.5.12] - 2021-09-09
###### SDK Version: v0.4.3

### :spider: Fixed
* Fixed bad migrations [#187](https://github.com/cloudquery/cq-provider-aws/pull/187)


## [v0.5.11] - 2021-09-09
###### SDK Version: v0.4.3

### :spider: Fixed
* Fixed kms key fetching when provider tried to fetch aws managed keys rotation properties [#168](https://github.com/cloudquery/cq-provider-aws/pull/168)

### :rocket: Added
* Add Elbv2 Attributes [#177](https://github.com/cloudquery/cq-provider-aws/pull/177)
* Added integration test for ec2.images resource [#184](https://github.com/cloudquery/cq-provider-aws/pull/184)

### :gear: Changed
* Upgraded to SDK Version [v0.4.3](https://github.com/cloudquery/cq-provider-sdk/blob/main/CHANGELOG.md)

## [v0.5.10] - 2021-08-19
###### SDK Version: 0.3.4

### :spider: Fixed
* Fixed bad multiplexes in iam_roles_policies, iam_group_policies, iam_user_policies [#168](https://github.com/cloudquery/cq-provider-aws/pull/168)

### :rocket: Added
* Added multiple e2e integration tests for resources for increased stability [#158](https://github.com/cloudquery/cq-provider-aws/pull/158)


## [v0.5.9] - 2021-08-19
###### SDK Version: 0.3.3-rc2

### :spider: Fixed
* Fixed bad migration SQL [#160](https://github.com/cloudquery/cq-provider-aws/pull/160)


## [v0.5.8] - 2021-08-19
###### SDK Version: 0.3.3-rc2

### :spider: Fixed
* Fixed bad multiplexes (cloudfront.distributions, cloudtrail.trails, cognito.user_pools) and bad PK s3_bucket_core_rules [#158](https://github.com/cloudquery/cq-provider-aws/pull/158) Thanks [@jbertman](https://github.com/jbertman) for reporting
* All providers must be wrapped in credentials cache, should fix [Assume Role issues](https://github.com/aws/aws-sdk-go-v2/issues/914) [#153](https://github.com/cloudquery/cq-provider-aws/pull/153)

### :rocket: Added
* Added support for provider e2e testing to improve stability, upcoming release should include more tests.


## [v0.5.7] - 2021-08-12
###### SDK Version: 0.3.2

### :spider: Fixed
* Updated organizational accounts call [#146](https://github.com/cloudquery/cq-provider-aws/pull/146)


## [v0.5.6] - 2021-08-12
###### SDK Version: 0.3.2

### :spider: Fixed
* Removed problematic s3 bucket primary keys [#144](https://github.com/cloudquery/cq-provider-aws/pull/144)


## [v0.5.5] - 2021-08-11
###### SDK Version: 0.3.2
### :gear: Changed
* Upgraded to SDK Version [0.3.2](https://github.com/cloudquery/cq-provider-sdk/blob/main/CHANGELOG.md#v032---2020-08-11)

### :spider: Fixed
* Fixed relation tables inserts - Delete filter issue [#141](https://github.com/cloudquery/cq-provider-aws/pull/141)

### :rocket: Added
* Added some resources missing Arn's [#142](https://github.com/cloudquery/cq-provider-aws/issues/142)

## [v0.5.4] - 2021-08-07
###### SDK Version: 0.3.1

### :spider: Fixed
* Fixed cognito identity pools [#138](https://github.com/cloudquery/cq-provider-aws/pull/138)
* Fixed Not fetching all Lambda functions [#135](https://github.com/cloudquery/cq-provider-aws/issues/135) [#136](https://github.com/cloudquery/cq-provider-aws/pull/136)
* Ignore access denied in s3 relations[#138](https://github.com/cloudquery/cq-provider-aws/pull/138)


## [v0.5.3] - 2021-08-04
###### SDK Version: 0.3.1

### :spider: Fixed
* Fixed [#130](https://github.com/cloudquery/cq-provider-aws/issues/130) IAM User Tags not persisting [#125](https://github.com/cloudquery/cq-provider-aws/pull/132)
* Fixed s3 bucket resource resolving errors [#131](https://github.com/cloudquery/cq-provider-aws/pull/131)

## [v0.5.2] - 2021-08-01
###### SDK Version: 0.3.1

### :rocket: Added
* Added support for arm64 [#128](https://github.com/cloudquery/cq-provider-aws/pull/128)

## [v0.5.1] - 2021-07-30
###### SDK Version: 0.3.1

### :rocket: Added
* Added new resource directconnect lags by [@James-Quigley](https://github.com/James-Quigley) [#122](https://github.com/cloudquery/cq-provider-aws/pull/122)

### :gear: Changed
* Updated SDK to version [0.3.1](https://github.com/cloudquery/cq-provider-sdk/blob/main/CHANGELOG.md#v031---2020-07-30)

### :spider: Fixed
* Fixed iam policies encoded policies parsing [#125](https://github.com/cloudquery/cq-provider-aws/pull/125)
* Fixed cognito_user_pools input [#126](https://github.com/cloudquery/cq-provider-aws/pull/126)


## [v0.5.0] - 2021-07-28
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

## [0.4.11] - 2021-07-15

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

