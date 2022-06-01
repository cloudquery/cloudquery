# Changelog

All notable changes to this provider will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).


### [0.12.2](https://github.com/cloudquery/cq-provider-aws/compare/v0.12.1...v0.12.2) (2022-06-01)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/cq-provider-sdk to v0.10.6 ([#977](https://github.com/cloudquery/cq-provider-aws/issues/977)) ([2baacf2](https://github.com/cloudquery/cq-provider-aws/commit/2baacf2422d1a68a59f0e1dbde5712383a6f06c4))
* **errors:** Use `WithNoOverwrite` instead of `WithOptionalSeverity` ([#975](https://github.com/cloudquery/cq-provider-aws/issues/975)) ([e6d2086](https://github.com/cloudquery/cq-provider-aws/commit/e6d2086c937f810e18ba3c8aa9606f10cea22b78))
* Fixed `NotFound` errors exceptions for some resources ([#965](https://github.com/cloudquery/cq-provider-aws/issues/965)) ([d2cab56](https://github.com/cloudquery/cq-provider-aws/commit/d2cab566bbd7f4c314917eeae1681e6ffbf1b488))
* Update endpoints ([#972](https://github.com/cloudquery/cq-provider-aws/issues/972)) ([2af266a](https://github.com/cloudquery/cq-provider-aws/commit/2af266ad068940f6b95e356111ca1683f5a009f5))

### [0.12.1](https://github.com/cloudquery/cq-provider-aws/compare/v0.12.0...v0.12.1) (2022-05-31)


### Features

* Add tags for organizations.Account ([#942](https://github.com/cloudquery/cq-provider-aws/issues/942)) ([b1a350d](https://github.com/cloudquery/cq-provider-aws/commit/b1a350debbf25ac8d7c5ffb539632d31038674ad)), closes [#940](https://github.com/cloudquery/cq-provider-aws/issues/940)
* Add waf/wafv2 logging config ([#814](https://github.com/cloudquery/cq-provider-aws/issues/814)) ([ed6c836](https://github.com/cloudquery/cq-provider-aws/commit/ed6c8363bcea6668a1ae6d4fa97e1051b26b0527))


### Bug Fixes

* Classify MetadataException ([#953](https://github.com/cloudquery/cq-provider-aws/issues/953)) ([2b74e7b](https://github.com/cloudquery/cq-provider-aws/commit/2b74e7bba60e04a0cfead40f7188e7f7d9c1e9cc))
* **deps:** Update module github.com/cloudquery/cq-provider-sdk to v0.10.2 ([#951](https://github.com/cloudquery/cq-provider-aws/issues/951)) ([b5b4c97](https://github.com/cloudquery/cq-provider-aws/commit/b5b4c97e07d0c75cd646d6946c594d2718da028d))
* **deps:** Update module github.com/cloudquery/cq-provider-sdk to v0.10.3 ([#954](https://github.com/cloudquery/cq-provider-aws/issues/954)) ([21a5818](https://github.com/cloudquery/cq-provider-aws/commit/21a5818e250ad7dfab78a17aba338e87c72e275f))
* **deps:** Update module github.com/cloudquery/cq-provider-sdk to v0.10.4 ([#961](https://github.com/cloudquery/cq-provider-aws/issues/961)) ([648f6c1](https://github.com/cloudquery/cq-provider-aws/commit/648f6c1ea58d8c89c09e816f6d0476ed2864f4c7))
* **deps:** Update module github.com/cloudquery/cq-provider-sdk to v0.10.5 ([#967](https://github.com/cloudquery/cq-provider-aws/issues/967)) ([40a1d06](https://github.com/cloudquery/cq-provider-aws/commit/40a1d06fb19f375587c99f50d061af8ab3e41079))
* Remove relation tables PK ([#921](https://github.com/cloudquery/cq-provider-aws/issues/921)) ([036ce47](https://github.com/cloudquery/cq-provider-aws/commit/036ce474f801bbc25d7067f4359e88b778cbc503))
* Update endpoints ([#948](https://github.com/cloudquery/cq-provider-aws/issues/948)) ([3b5b193](https://github.com/cloudquery/cq-provider-aws/commit/3b5b193310b066ebe5fced2ea1bff6fd91fe6fca))
* Update endpoints ([#952](https://github.com/cloudquery/cq-provider-aws/issues/952)) ([c59523c](https://github.com/cloudquery/cq-provider-aws/commit/c59523cb41c6a5afcf999f9992999653409db141))
* Update endpoints ([#956](https://github.com/cloudquery/cq-provider-aws/issues/956)) ([5702860](https://github.com/cloudquery/cq-provider-aws/commit/5702860a8bd01163b7483205e0351dc1b76687df))
* Update endpoints ([#958](https://github.com/cloudquery/cq-provider-aws/issues/958)) ([df14874](https://github.com/cloudquery/cq-provider-aws/commit/df14874c0bfbfdf3d14cdc17eaf0c0e44a18bd71))
* Update timestamps fields ([#891](https://github.com/cloudquery/cq-provider-aws/issues/891)) ([48b9e6f](https://github.com/cloudquery/cq-provider-aws/commit/48b9e6f35c6794c57efb10d033bba9c084b1b451))

## [0.12.0](https://github.com/cloudquery/cq-provider-aws/compare/v0.12.0-rc2...v0.12.0) (2022-05-24)


### ⚠ BREAKING CHANGES

* Remove migrations (#933)

### Features

* Remove migrations ([#933](https://github.com/cloudquery/cq-provider-aws/issues/933)) ([37620e3](https://github.com/cloudquery/cq-provider-aws/commit/37620e330ba187c4da2ff02382423b92be91e318))


## [0.12.0-rc2](https://github.com/cloudquery/cq-provider-aws/compare/v0.12.0-rc1...v0.12.0-rc2) (2022-05-24)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/cq-provider-sdk to v0.10.1 ([#938](https://github.com/cloudquery/cq-provider-aws/issues/938)) ([7db6d1a](https://github.com/cloudquery/cq-provider-aws/commit/7db6d1a854f89f7e69149491eb44375e2aee7cdf))


### Miscellaneous Chores

* Release 0.12.0-rc2 ([#945](https://github.com/cloudquery/cq-provider-aws/issues/945)) ([4987b4a](https://github.com/cloudquery/cq-provider-aws/commit/4987b4a2a854ccdc1bb97d1d960a2783bb7ec260))

## [0.12.0-rc1](https://github.com/cloudquery/cq-provider-aws/compare/v0.11.8...v0.12.0-rc1) (2022-05-24)


### ⚠ BREAKING CHANGES

* Remove migrations (#933)

### Features

* Remove migrations ([#933](https://github.com/cloudquery/cq-provider-aws/issues/933)) ([37620e3](https://github.com/cloudquery/cq-provider-aws/commit/37620e330ba187c4da2ff02382423b92be91e318))


### Miscellaneous Chores

* Release 0.12.0-rc1 ([#943](https://github.com/cloudquery/cq-provider-aws/issues/943)) ([6d8048d](https://github.com/cloudquery/cq-provider-aws/commit/6d8048d37e7b3334ccc424b1a51d7bc0c93d16d6))

### [0.11.8](https://github.com/cloudquery/cq-provider-aws/compare/v0.11.7...v0.11.8) (2022-05-23)


### Features

* Parallelize Sagemaker Training Jobs ([c925608](https://github.com/cloudquery/cq-provider-aws/commit/c925608d56453e55d78ccdc4f8c4f65a222265cc))


### Bug Fixes

* **deps:** Update module github.com/cloudquery/cq-provider-sdk to v0.9.5 ([#935](https://github.com/cloudquery/cq-provider-aws/issues/935)) ([c7474f3](https://github.com/cloudquery/cq-provider-aws/commit/c7474f3cd2b0f348a19b9601192d02d2199baaf4))
* Ignore backup global settings in disabled region ([#923](https://github.com/cloudquery/cq-provider-aws/issues/923)) ([1100f6a](https://github.com/cloudquery/cq-provider-aws/commit/1100f6aa3988883d5d28769aa6868769ae8a5e37))

### [0.11.7](https://github.com/cloudquery/cq-provider-aws/compare/v0.11.6...v0.11.7) (2022-05-17)


### Bug Fixes

* Fix NotFound error in `aws_shield_*` resources ([#916](https://github.com/cloudquery/cq-provider-aws/issues/916)) ([fc9cdcc](https://github.com/cloudquery/cq-provider-aws/commit/fc9cdcc5ad804ed63ee27d027838882cfff82e57))
* Hardcoded region for Get-Caller-Identity call ([0f2091e](https://github.com/cloudquery/cq-provider-aws/commit/0f2091e4f75016cf25321a1c35ad5f36cf0b343c))

### [0.11.6](https://github.com/cloudquery/cq-provider-aws/compare/v0.11.5...v0.11.6) (2022-05-17)


### Features

* Add partition info to regions ([#898](https://github.com/cloudquery/cq-provider-aws/issues/898)) ([76d4587](https://github.com/cloudquery/cq-provider-aws/commit/76d4587f4d2b74cb6a73385021dd60cc9d23e678))
* Added Athena resources: data_catalogs, work_groups, named_queries, prepared_statements ([#804](https://github.com/cloudquery/cq-provider-aws/issues/804)) ([bf77311](https://github.com/cloudquery/cq-provider-aws/commit/bf7731126c9566040266db0ff6e606acce7eb87e))


### Bug Fixes

* Add Partition To client ([#899](https://github.com/cloudquery/cq-provider-aws/issues/899)) ([3a77950](https://github.com/cloudquery/cq-provider-aws/commit/3a77950cf7121a56adb9d6bc5c7b395ba34f085c))
* Classify auth failure ([#904](https://github.com/cloudquery/cq-provider-aws/issues/904)) ([c134c10](https://github.com/cloudquery/cq-provider-aws/commit/c134c10c8c147ca659e8e3e2b7db04ae4830c9f7))
* **deps:** Update module github.com/cloudquery/cq-provider-sdk to v0.9.4 ([#912](https://github.com/cloudquery/cq-provider-aws/issues/912)) ([c3cd011](https://github.com/cloudquery/cq-provider-aws/commit/c3cd01134ed9ccc2d4d26a058a6ea5ce131b5383))
* Fail on access errors during initial setup ([#906](https://github.com/cloudquery/cq-provider-aws/issues/906)) ([c512535](https://github.com/cloudquery/cq-provider-aws/commit/c512535fd67960150dd0e4800eabf9b9df0dd83b))
* Fix NotFound error in `aws_sqs_queues` ([#910](https://github.com/cloudquery/cq-provider-aws/issues/910)) ([3c893df](https://github.com/cloudquery/cq-provider-aws/commit/3c893df4b1cf1d65b33104e63f993a2694067698))
* Merge migrations ([#913](https://github.com/cloudquery/cq-provider-aws/issues/913)) ([5eb71ed](https://github.com/cloudquery/cq-provider-aws/commit/5eb71eda052cd26dfe2899376e2c4b738fce2007))
* Non standard partition fixes ([#894](https://github.com/cloudquery/cq-provider-aws/issues/894)) ([2172e49](https://github.com/cloudquery/cq-provider-aws/commit/2172e499c1e409b0af39ae120b69a5e0caee3e1a))
* Remove "increase max_retries" from the throttling error message ([#868](https://github.com/cloudquery/cq-provider-aws/issues/868)) ([2c850c0](https://github.com/cloudquery/cq-provider-aws/commit/2c850c0f073c05969efa09cd7667b86dbc96a899))

### [0.11.5](https://github.com/cloudquery/cq-provider-aws/compare/v0.11.4...v0.11.5) (2022-05-11)


### Features

* Add S3 Bucket fetch speed improvement ([#840](https://github.com/cloudquery/cq-provider-aws/issues/840)) ([0d57a54](https://github.com/cloudquery/cq-provider-aws/commit/0d57a5484f4a3d2d484cc03668ae508e86778ba6))
* Shield protections added ([#728](https://github.com/cloudquery/cq-provider-aws/issues/728)) ([fc8a308](https://github.com/cloudquery/cq-provider-aws/commit/fc8a308e8b3450809c0ac04befeaf13f1d25b35d))
* Support Xray groups and sampling rules ([#841](https://github.com/cloudquery/cq-provider-aws/issues/841)) ([e9c57b8](https://github.com/cloudquery/cq-provider-aws/commit/e9c57b88f58fd928712c4757954cdb8f3c453e31))


### Bug Fixes

* Add ON DELETE CASCADE to redshift snapshots. ([#880](https://github.com/cloudquery/cq-provider-aws/issues/880)) ([0009d0a](https://github.com/cloudquery/cq-provider-aws/commit/0009d0a55ac3d12bb4294ccfa3261c90d51e4bf0))
* Adjust some PKs ([#849](https://github.com/cloudquery/cq-provider-aws/issues/849)) ([45807e8](https://github.com/cloudquery/cq-provider-aws/commit/45807e8b54c0b861d102c855bbbca9e8a33450cf))
* Do not fail if rds parameter group is gone ([#887](https://github.com/cloudquery/cq-provider-aws/issues/887)) ([47c5032](https://github.com/cloudquery/cq-provider-aws/commit/47c50321548dd8f4f9296657149fc6be746f2905))
* Handle a case where autoscaling group is being deleted. ([#872](https://github.com/cloudquery/cq-provider-aws/issues/872)) ([28c19d4](https://github.com/cloudquery/cq-provider-aws/commit/28c19d4f4c32ed8e34a66ec431fabb46a721f0df))
* Ignore apigateway model-template fetch if model not exists ([#876](https://github.com/cloudquery/cq-provider-aws/issues/876)) ([0afb429](https://github.com/cloudquery/cq-provider-aws/commit/0afb429948c476c941cbaf816b06170fae6c5204))
* Lambda function alias duplicate PK ([#881](https://github.com/cloudquery/cq-provider-aws/issues/881)) ([2ad1fef](https://github.com/cloudquery/cq-provider-aws/commit/2ad1fefd80fb15f9210cc8a3fb8c5636db3d43da))
* Update migrations for 0.11.5 ([#886](https://github.com/cloudquery/cq-provider-aws/issues/886)) ([6758918](https://github.com/cloudquery/cq-provider-aws/commit/6758918bcf779f2e0fa735b069e90383ab951829))

### [0.11.4](https://github.com/cloudquery/cq-provider-aws/compare/v0.11.3...v0.11.4) (2022-05-10)


### Features

* Support AWS partitions ([#842](https://github.com/cloudquery/cq-provider-aws/issues/842)) ([6976653](https://github.com/cloudquery/cq-provider-aws/commit/6976653037def1afa334162ec6e8375e1c31237e))


### Bug Fixes

* Classify UnrecognizedClientException as Access error ([#862](https://github.com/cloudquery/cq-provider-aws/issues/862)) ([6dbfbc1](https://github.com/cloudquery/cq-provider-aws/commit/6dbfbc19543801513342142ad0091088e69cd912))
* **deps:** Update module github.com/cloudquery/cq-provider-sdk to v0.9.3 ([#856](https://github.com/cloudquery/cq-provider-aws/issues/856)) ([6a8ecee](https://github.com/cloudquery/cq-provider-aws/commit/6a8ecee4bba7a5f917c5421bae63549d034bf77c))
* Reduce retry params ([#855](https://github.com/cloudquery/cq-provider-aws/issues/855)) ([23e2fe1](https://github.com/cloudquery/cq-provider-aws/commit/23e2fe178fb3f59b05c5ab74db2c464bf3541872))
* Update endpoints ([#854](https://github.com/cloudquery/cq-provider-aws/issues/854)) ([05be3c2](https://github.com/cloudquery/cq-provider-aws/commit/05be3c2f283e86bdc802cbac2db2e13cea36f5e9))
* Update endpoints ([#870](https://github.com/cloudquery/cq-provider-aws/issues/870)) ([06a217b](https://github.com/cloudquery/cq-provider-aws/commit/06a217bebe32338c50147098a28b086f9d393176))

## [v0.9.1] - 2022-01-13
###### SDK Version: 0.6.1
### :spider: Fixed
* Config Status Recorder [#406](https://github.com/cloudquery/cq-provider-aws/pull/406)


## [v0.9.0] - 2022-01-10
###### SDK Version: 0.6.1
### :spider: Fixed
* chore: update partition_service_region.json [#398](https://github.com/cloudquery/cq-provider-aws/pull/398)
### 💥 Breaking Changes
* SG simplification [#363](https://github.com/cloudquery/cq-provider-aws/pull/363)
* upgrade initial migration to v0.8.5 due to protocol ugprade v3 [#390](https://github.com/cloudquery/cq-provider-aws/pull/390)
### 🚀 Added
* Support Wildcard for Region [#391](https://github.com/cloudquery/cq-provider-aws/pull/391)
* added asciicheck linter [#397](https://github.com/cloudquery/cq-provider-aws/pull/397)
* turn issue template into github forms [#388](https://github.com/cloudquery/cq-provider-aws/pull/388)

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
