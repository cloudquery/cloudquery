# Changelog

All notable changes to this provider will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).


## [22.13.2](https://github.com/cloudquery/cloudquery/compare/plugins-source-aws-v22.13.1...plugins-source-aws-v22.13.2) (2023-09-29)


### Bug Fixes

* Fix for `example_queries` skip on error message ([#14122](https://github.com/cloudquery/cloudquery/issues/14122)) ([95b3641](https://github.com/cloudquery/cloudquery/commit/95b3641a1483a6dc9054023b1b8f0512c2a810cc))
* Fix query for finding unused target groups for AWS ([#13616](https://github.com/cloudquery/cloudquery/issues/13616)) ([58b07cd](https://github.com/cloudquery/cloudquery/commit/58b07cd5b94889965bfbb8ffc8c9bf66a6579593))

## [22.13.1](https://github.com/cloudquery/cloudquery/compare/plugins-source-aws-v22.13.0...plugins-source-aws-v22.13.1) (2023-09-27)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.11.1 ([#14063](https://github.com/cloudquery/cloudquery/issues/14063)) ([5a0ff7b](https://github.com/cloudquery/cloudquery/commit/5a0ff7b67890478c371385b379e0a8ef0c2f4865))
* Don't stop on error for Regional WebACL Resources ([#14045](https://github.com/cloudquery/cloudquery/issues/14045)) ([33ab5b0](https://github.com/cloudquery/cloudquery/commit/33ab5b0f86a2e7bef9d1f71d01d4481d7af0445d))

## [22.13.0](https://github.com/cloudquery/cloudquery/compare/plugins-source-aws-v22.12.0...plugins-source-aws-v22.13.0) (2023-09-26)


### This Release has the Following Changes to Tables
- Table `aws_ec2_vpc_endpoint_connections` was added

### Features

* Add support for AWS EC2 VPC Endpoint Connections ([#14044](https://github.com/cloudquery/cloudquery/issues/14044)) ([0fa640e](https://github.com/cloudquery/cloudquery/commit/0fa640e1ca458358bc497064c33d4b74ec67656d)), closes [#14030](https://github.com/cloudquery/cloudquery/issues/14030)
* **services:** Support newly added regions ([#14032](https://github.com/cloudquery/cloudquery/issues/14032)) ([c11bfce](https://github.com/cloudquery/cloudquery/commit/c11bfcefd472526f106c245ac4077f7a4bec0ca0))


### Bug Fixes

* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.10.2 ([#13988](https://github.com/cloudquery/cloudquery/issues/13988)) ([aebaddf](https://github.com/cloudquery/cloudquery/commit/aebaddfc5ca0d7574b8cd72e9e074ec612472dbe))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.11.0 ([#14031](https://github.com/cloudquery/cloudquery/issues/14031)) ([ac7cdc4](https://github.com/cloudquery/cloudquery/commit/ac7cdc4f7d71599dad89b3170bb7bda676984228))

## [22.12.0](https://github.com/cloudquery/cloudquery/compare/plugins-source-aws-v22.11.0...plugins-source-aws-v22.12.0) (2023-09-21)


### Features

* **services:** Support newly added regions ([#13938](https://github.com/cloudquery/cloudquery/issues/13938)) ([c4e810b](https://github.com/cloudquery/cloudquery/commit/c4e810bea14e854771bd779cb04f3e4cc2d0ee40))


### Bug Fixes

* **deps:** Update github.com/cloudquery/arrow/go/v14 digest to ffb7089 ([#13879](https://github.com/cloudquery/cloudquery/issues/13879)) ([f95ced5](https://github.com/cloudquery/cloudquery/commit/f95ced5daa2b123bd71ddff75bd76b3b008790c1))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.10.0 ([#13978](https://github.com/cloudquery/cloudquery/issues/13978)) ([2efdf55](https://github.com/cloudquery/cloudquery/commit/2efdf55aed94a14c35c51632ff61ed454caaf5a5))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.8.0 ([#13950](https://github.com/cloudquery/cloudquery/issues/13950)) ([15b0b69](https://github.com/cloudquery/cloudquery/commit/15b0b6925932613ed2915a3255b3466f21a5c7bf))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.9.0 ([#13960](https://github.com/cloudquery/cloudquery/issues/13960)) ([f074076](https://github.com/cloudquery/cloudquery/commit/f074076a21dc0b8cadfdc3adb9731473d24d28b1))
* Flipped condition in query for SecretsManager.4 ([#13864](https://github.com/cloudquery/cloudquery/issues/13864)) ([76c21fb](https://github.com/cloudquery/cloudquery/commit/76c21fbec6dbb2ca03327a6a93024b8ecd538084))
* Flipped condition on secrets manager policy ([#13862](https://github.com/cloudquery/cloudquery/issues/13862)) ([c604b70](https://github.com/cloudquery/cloudquery/commit/c604b707c5ac63b9814915ccffb94b84bdbf2e32))

## [22.11.0](https://github.com/cloudquery/cloudquery/compare/plugins-source-aws-v22.10.0...plugins-source-aws-v22.11.0) (2023-09-12)


### Features

* **services:** Support newly added regions ([#13790](https://github.com/cloudquery/cloudquery/issues/13790)) ([9aa3ff3](https://github.com/cloudquery/cloudquery/commit/9aa3ff3bab520f68e240974c915b725cf17d2c8b))


### Bug Fixes

* **deps:** Update github.com/cloudquery/arrow/go/v14 digest to 483f6b2 ([#13780](https://github.com/cloudquery/cloudquery/issues/13780)) ([8d31b44](https://github.com/cloudquery/cloudquery/commit/8d31b44f787f42d47f186cdcc4a5739a3a370a5f))

## [22.10.0](https://github.com/cloudquery/cloudquery/compare/plugins-source-aws-v22.9.0...plugins-source-aws-v22.10.0) (2023-09-07)


### Features

* Add scheduler option to AWS ([#13757](https://github.com/cloudquery/cloudquery/issues/13757)) ([521918f](https://github.com/cloudquery/cloudquery/commit/521918f94b1783be75bb7f4e9024b95ee7be7c3a))


### Bug Fixes

* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.7.1 ([#13713](https://github.com/cloudquery/cloudquery/issues/13713)) ([73004dc](https://github.com/cloudquery/cloudquery/commit/73004dcabd05bf474d8b5960b8c747a894b98560))
* Issue [#13433](https://github.com/cloudquery/cloudquery/issues/13433) - match on lowercase "redirect" in aws policy  ([#13751](https://github.com/cloudquery/cloudquery/issues/13751)) ([539e029](https://github.com/cloudquery/cloudquery/commit/539e029412e17549cb4ad552ef89b1be90d58387))

## [22.9.0](https://github.com/cloudquery/cloudquery/compare/plugins-source-aws-v22.8.2...plugins-source-aws-v22.9.0) (2023-09-05)


### Features

* Add `aws_s3_bucket_notification_configurations` table ([#13608](https://github.com/cloudquery/cloudquery/issues/13608)) ([d3b08ee](https://github.com/cloudquery/cloudquery/commit/d3b08eef77264cd0d87a1b4635a502d43c20232e))
* Add `aws_s3_bucket_object_lock_configurations` table ([#13609](https://github.com/cloudquery/cloudquery/issues/13609)) ([1debcd3](https://github.com/cloudquery/cloudquery/commit/1debcd30b508aad74b18f7cee1d6c7c8062aa78e)), closes [#13606](https://github.com/cloudquery/cloudquery/issues/13606)
* **services:** Support newly added regions ([#13617](https://github.com/cloudquery/cloudquery/issues/13617)) ([bb67e06](https://github.com/cloudquery/cloudquery/commit/bb67e06d6f7c89a436e10dd476721a3dc2951d0e))


### Bug Fixes

* **deps:** Update `github.com/cloudquery/plugin-sdk/v4` to `v4.7.0` ([#13623](https://github.com/cloudquery/cloudquery/issues/13623)) ([871a792](https://github.com/cloudquery/cloudquery/commit/871a792ebf26ba36ca0b2452c591979460242494))
* **deps:** Update github.com/99designs/go-keychain digest to 9cf53c8 ([#13561](https://github.com/cloudquery/cloudquery/issues/13561)) ([a170256](https://github.com/cloudquery/cloudquery/commit/a17025657e92b017fe3c8bd37abfaa2354e6e818))
* **deps:** Update github.com/apache/arrow/go/v14 digest to a526ba6 ([#13562](https://github.com/cloudquery/cloudquery/issues/13562)) ([248672b](https://github.com/cloudquery/cloudquery/commit/248672beb020828cde1cb608d5c1ed6d656c777b))
* **deps:** Update github.com/cloudquery/arrow/go/v14 digest to cd3d411 ([#13598](https://github.com/cloudquery/cloudquery/issues/13598)) ([f22bfa6](https://github.com/cloudquery/cloudquery/commit/f22bfa6b2d4fd0caeacf0726ccd307db38f8860c))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.6.0 ([#13492](https://github.com/cloudquery/cloudquery/issues/13492)) ([c305876](https://github.com/cloudquery/cloudquery/commit/c305876e3d92944aa6c1a26547f786fdc5b50e23))
* Don't reference public schema in policies ([#13619](https://github.com/cloudquery/cloudquery/issues/13619)) ([163aa94](https://github.com/cloudquery/cloudquery/commit/163aa942db80e7401769bf7f8fe67cda35b2ecba))
* Use Pagination for `aws_ec2_images` ([#13560](https://github.com/cloudquery/cloudquery/issues/13560)) ([a3779e8](https://github.com/cloudquery/cloudquery/commit/a3779e8476515e3bc6a583aa3d32774485cdffc0))

## [22.8.2](https://github.com/cloudquery/cloudquery/compare/plugins-source-aws-v22.8.1...plugins-source-aws-v22.8.2) (2023-08-30)


### Bug Fixes

* Flush state backend to fix incremental fetching of `aws_cloudtrail_events` ([#13430](https://github.com/cloudquery/cloudquery/issues/13430)) ([5ed403a](https://github.com/cloudquery/cloudquery/commit/5ed403a434bd6b5cb3721eef96b860c9103e2ffe))

## [22.8.1](https://github.com/cloudquery/cloudquery/compare/plugins-source-aws-v22.8.0...plugins-source-aws-v22.8.1) (2023-08-29)


### Bug Fixes

* Store AWS config per account ([#13391](https://github.com/cloudquery/cloudquery/issues/13391)) ([a4245d7](https://github.com/cloudquery/cloudquery/commit/a4245d7deea76ddd5add531f4fd6d769627647a6)), closes [#13389](https://github.com/cloudquery/cloudquery/issues/13389)

## [22.8.0](https://github.com/cloudquery/cloudquery/compare/plugins-source-aws-v22.7.0...plugins-source-aws-v22.8.0) (2023-08-29)


### Features

* Implement CIS AWS v1.5.0 Section 1.16 and 1.20 ([#13290](https://github.com/cloudquery/cloudquery/issues/13290)) ([7eb3e06](https://github.com/cloudquery/cloudquery/commit/7eb3e06f93da9678bc78ab83c5acd398cb706bdd))
* **services:** Support newly added regions ([#13273](https://github.com/cloudquery/cloudquery/issues/13273)) ([c6727a9](https://github.com/cloudquery/cloudquery/commit/c6727a9e34b640ebad626b6ef70707ca488f562d))
* **services:** Support newly added regions ([#13358](https://github.com/cloudquery/cloudquery/issues/13358)) ([4ed4a9a](https://github.com/cloudquery/cloudquery/commit/4ed4a9a1da0fb856556cc5fb42e0b418a95097eb))


### Bug Fixes

* **deps:** Update `github.com/cloudquery/arrow/go/v13` to `github.com/apache/arrow/go/v14` ([#13341](https://github.com/cloudquery/cloudquery/issues/13341)) ([feb8f87](https://github.com/cloudquery/cloudquery/commit/feb8f87d8d761eb9c49ce84329ad0397f730a918))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.5.2 ([#13262](https://github.com/cloudquery/cloudquery/issues/13262)) ([5c55aa3](https://github.com/cloudquery/cloudquery/commit/5c55aa35282786375e8ce9493b2a4878e0fb27bc))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.5.5 ([#13285](https://github.com/cloudquery/cloudquery/issues/13285)) ([e076abd](https://github.com/cloudquery/cloudquery/commit/e076abd9d67813a29ced0c1b7b1664fd728b9ba8))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.5.6 ([#13345](https://github.com/cloudquery/cloudquery/issues/13345)) ([a995a05](https://github.com/cloudquery/cloudquery/commit/a995a0598a209e0fe3ba09f4ced2a052dc14b67a))
* Race condition in Initializing services ([#13300](https://github.com/cloudquery/cloudquery/issues/13300)) ([73a093d](https://github.com/cloudquery/cloudquery/commit/73a093dc0de33362eeb91077ddb14c9a8bacf06a))

## [22.7.0](https://github.com/cloudquery/cloudquery/compare/plugins-source-aws-v22.6.0...plugins-source-aws-v22.7.0) (2023-08-21)


### Features

* Instantiate services at sync time rather than during init phase of sync ([#13059](https://github.com/cloudquery/cloudquery/issues/13059)) ([99e6889](https://github.com/cloudquery/cloudquery/commit/99e6889c49f66e66d03fbf76064d779a77281f70))


### Bug Fixes

* AWS foundational security controls for S3.8 query reference ([#13065](https://github.com/cloudquery/cloudquery/issues/13065)) ([7cde3d2](https://github.com/cloudquery/cloudquery/commit/7cde3d2d063c65c714812a9d98dfaea0f60b84d0)), closes [#13064](https://github.com/cloudquery/cloudquery/issues/13064)
* **deps:** Update AWS modules ([#13246](https://github.com/cloudquery/cloudquery/issues/13246)) ([49ee475](https://github.com/cloudquery/cloudquery/commit/49ee4752fff1e6fc06e0b50e0450bdadd4373d16))
* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to 5b83d4f ([#13203](https://github.com/cloudquery/cloudquery/issues/13203)) ([b0a4b8c](https://github.com/cloudquery/cloudquery/commit/b0a4b8ccf7c429bf5a6ed88866865212015b68e4))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.5.1 ([#13195](https://github.com/cloudquery/cloudquery/issues/13195)) ([a184c37](https://github.com/cloudquery/cloudquery/commit/a184c3786ad49df8564344773e9b96f617ef87a1))
* Panic while fetching `aws_codecommit_repositories` ([#13223](https://github.com/cloudquery/cloudquery/issues/13223)) ([4d9bcc1](https://github.com/cloudquery/cloudquery/commit/4d9bcc188723f318a794ad343dc2c2ac5b3da018))
* Remove unused query ([#13152](https://github.com/cloudquery/cloudquery/issues/13152)) ([f050699](https://github.com/cloudquery/cloudquery/commit/f050699ba1bc229368a566764890a06ca90d53f0))

## [22.6.0](https://github.com/cloudquery/cloudquery/compare/plugins-source-aws-v22.5.1...plugins-source-aws-v22.6.0) (2023-08-15)


### This Release has the Following Changes to Tables
- Table `aws_elbv2_load_balancers`: column added with name `enforce_security_group_inbound_rules_on_private_link_traffic` and type `utf8`
- Table `aws_fsx_data_repository_tasks`: column added with name `release_configuration` and type `json`

### Features

* **services:** Support newly added regions ([#13062](https://github.com/cloudquery/cloudquery/issues/13062)) ([cbee2c2](https://github.com/cloudquery/cloudquery/commit/cbee2c23ad3f81c042284b222b16373469cca02e))


### Bug Fixes

* **deps:** Update AWS modules ([#13012](https://github.com/cloudquery/cloudquery/issues/13012)) ([d163f2b](https://github.com/cloudquery/cloudquery/commit/d163f2b46cd7c9a0961b8bced02ae8adc9b43bb1))
* **deps:** Update AWS modules ([#13013](https://github.com/cloudquery/cloudquery/issues/13013)) ([0a12c5b](https://github.com/cloudquery/cloudquery/commit/0a12c5bc1c44ba06c0320ae3c794ba6065ec20ea))
* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to e9683e1 ([#13015](https://github.com/cloudquery/cloudquery/issues/13015)) ([6557696](https://github.com/cloudquery/cloudquery/commit/65576966d3bd14297499a5b85d3b4fc2c7918df3))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.5.0 ([#13068](https://github.com/cloudquery/cloudquery/issues/13068)) ([7bb0e4b](https://github.com/cloudquery/cloudquery/commit/7bb0e4ba654971726e16a6a501393e3831170307))

## [22.5.1](https://github.com/cloudquery/cloudquery/compare/plugins-source-aws-v22.5.0...plugins-source-aws-v22.5.1) (2023-08-09)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.4.0 ([#12850](https://github.com/cloudquery/cloudquery/issues/12850)) ([0861200](https://github.com/cloudquery/cloudquery/commit/086120054b45213947e95be954ba6164b9cf6587))

## [22.5.0](https://github.com/cloudquery/cloudquery/compare/plugins-source-aws-v22.4.0...plugins-source-aws-v22.5.0) (2023-08-08)


### This Release has the Following Changes to Tables
- Table `aws_rds_cluster_snapshots`: column added with name `db_cluster_resource_id` and type `utf8`
- Table `aws_rds_clusters`: column added with name `local_write_forwarding_status` and type `utf8`
- Table `aws_rds_engine_versions`: column added with name `supports_local_write_forwarding` and type `bool`
- Table `aws_resiliencehub_alarm_recommendations`: column added with name `app_component_names` and type `list<item: utf8, nullable>`
- Table `aws_resiliencehub_app_assessments`: column added with name `drift_status` and type `utf8`
- Table `aws_resiliencehub_app_assessments`: column added with name `version_name` and type `utf8`
- Table `aws_resiliencehub_app_versions`: column added with name `creation_time` and type `timestamp[us, tz=UTC]`
- Table `aws_resiliencehub_app_versions`: column added with name `identifier` and type `int64`
- Table `aws_resiliencehub_app_versions`: column added with name `version_name` and type `utf8`
- Table `aws_resiliencehub_apps`: column added with name `drift_status` and type `utf8`
- Table `aws_resiliencehub_apps`: column added with name `event_subscriptions` and type `json`
- Table `aws_resiliencehub_apps`: column added with name `last_drift_evaluation_time` and type `timestamp[us, tz=UTC]`
- Table `aws_resiliencehub_apps`: column added with name `permission_model` and type `json`
- Table `aws_route53_hosted_zones`: column added with name `delegation_set` and type `json`
- Table `aws_scheduler_schedules`: column added with name `action_after_completion` and type `utf8`

### Features

* Replace DelegationSetId with full DelegationSet in aws_route53_hosted_zones ([#12737](https://github.com/cloudquery/cloudquery/issues/12737)) ([a78012d](https://github.com/cloudquery/cloudquery/commit/a78012db43ee529bb18f1dbdea9a0845ac169549))
* **services:** Support newly added regions ([#12797](https://github.com/cloudquery/cloudquery/issues/12797)) ([a64d5a9](https://github.com/cloudquery/cloudquery/commit/a64d5a9cd6fa907f18d985e037f3d14f02ff5980))


### Bug Fixes

* **deps:** Update AWS modules ([#12775](https://github.com/cloudquery/cloudquery/issues/12775)) ([945d570](https://github.com/cloudquery/cloudquery/commit/945d570c8d73e5f91c4854de279ba23d93e0081f))
* **deps:** Update AWS modules ([#12776](https://github.com/cloudquery/cloudquery/issues/12776)) ([91148ac](https://github.com/cloudquery/cloudquery/commit/91148ac0746c2aab49e2946f17a4c3da5cbde76c))
* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to f53878d ([#12778](https://github.com/cloudquery/cloudquery/issues/12778)) ([6f5d58e](https://github.com/cloudquery/cloudquery/commit/6f5d58e3b84d3c76b1d1a3d6c5a488f77995a057))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.2.4 ([#12718](https://github.com/cloudquery/cloudquery/issues/12718)) ([f059a15](https://github.com/cloudquery/cloudquery/commit/f059a159a2ee406ab2b0a33792c244cd217025a6))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.2.5 ([#12731](https://github.com/cloudquery/cloudquery/issues/12731)) ([d267239](https://github.com/cloudquery/cloudquery/commit/d267239aa3aca5f94bd36a8db1ec0d9f7dc0865f))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.2.6 ([#12799](https://github.com/cloudquery/cloudquery/issues/12799)) ([fb0e0d7](https://github.com/cloudquery/cloudquery/commit/fb0e0d75ab010f421c834e58d93676de76fcb423))

## [22.4.0](https://github.com/cloudquery/cloudquery/compare/plugins-source-aws-v22.3.0...plugins-source-aws-v22.4.0) (2023-08-01)


### This Release has the Following Changes to Tables
- Table `aws_appmesh_meshes` was added
- Table `aws_appmesh_virtual_gateways` was added
- Table `aws_appmesh_virtual_nodes` was added
- Table `aws_appmesh_virtual_routers` was added
- Table `aws_appmesh_virtual_services` was added
- Table `aws_ecr_repository_lifecycle_policies` was added

### Features

* **resources:** Add Support for AWS App Mesh resources ([#12582](https://github.com/cloudquery/cloudquery/issues/12582)) ([e0ca2be](https://github.com/cloudquery/cloudquery/commit/e0ca2be01f613fc40660c9cf438a65af17028f24))
* **resources:** Add support for ECR Lifecycle policy ([#12644](https://github.com/cloudquery/cloudquery/issues/12644)) ([344b6e6](https://github.com/cloudquery/cloudquery/commit/344b6e61d715b68c8c9bfeb75acef93e73a79790)), closes [#12594](https://github.com/cloudquery/cloudquery/issues/12594)

## [22.3.0](https://github.com/cloudquery/cloudquery/compare/plugins-source-aws-v22.2.0...plugins-source-aws-v22.3.0) (2023-08-01)


### This Release has the Following Changes to Tables
- Table `aws_appflow_flows` was added
- Table `aws_auditmanager_assessments` was added
- Table `aws_backup_report_plans` was added
- Table `aws_cloudformation_stacks`: column added with name `retain_except_on_create` and type `bool`
- Table `aws_ec2_ebs_snapshots`: column added with name `sse_type` and type `utf8`
- Table `aws_ec2_ebs_volumes`: column added with name `sse_type` and type `utf8`
- Table `aws_emr_notebook_executions` was added
- Table `aws_emr_release_labels` was added
- Table `aws_emr_steps` was added
- Table `aws_emr_supported_instance_types` was added
- Table `aws_rds_instances`: column added with name `percent_progress` and type `utf8`

### Features

* **resources:** Add support for Amazon Appflow Flows ([#12575](https://github.com/cloudquery/cloudquery/issues/12575)) ([43ed08e](https://github.com/cloudquery/cloudquery/commit/43ed08ee453b0043d43a3ff295b37055606fe6a5))
* **resources:** Add Support for AWS Audit Manager Assessments ([#12573](https://github.com/cloudquery/cloudquery/issues/12573)) ([ab5a939](https://github.com/cloudquery/cloudquery/commit/ab5a9392dc3fa8a49b6c2adb679895870ddd9af6))
* **resources:** Add support for AWS Backup Report Plan ([#12578](https://github.com/cloudquery/cloudquery/issues/12578)) ([5fa1af1](https://github.com/cloudquery/cloudquery/commit/5fa1af103a6e0e0b62d05f4d5d6f902c23a9e03e))
* **resources:** Adding additional EMR cluster resources ([#12562](https://github.com/cloudquery/cloudquery/issues/12562)) ([4a25c5c](https://github.com/cloudquery/cloudquery/commit/4a25c5cc03465d55fae305165a5d9311f1fd9c67))
* **services:** Support newly added regions ([#12671](https://github.com/cloudquery/cloudquery/issues/12671)) ([5af2d31](https://github.com/cloudquery/cloudquery/commit/5af2d3123635e71287c90deb54df06a83f9b3432))


### Bug Fixes

* **deps:** Update AWS modules ([#12591](https://github.com/cloudquery/cloudquery/issues/12591)) ([20eb1bf](https://github.com/cloudquery/cloudquery/commit/20eb1bf2328438a3478d667a78cc5fe03f96fa4f))
* **deps:** Update AWS modules ([#12592](https://github.com/cloudquery/cloudquery/issues/12592)) ([80ad5c5](https://github.com/cloudquery/cloudquery/commit/80ad5c564d5928def2a14ffa6ed6e7e034b39a36))
* **deps:** Update github.com/apache/arrow/go/v13 digest to 112f949 ([#12659](https://github.com/cloudquery/cloudquery/issues/12659)) ([48d73a9](https://github.com/cloudquery/cloudquery/commit/48d73a93e678994f43171c363f5a75c29547b0b9))
* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to 3452eb0 ([#12595](https://github.com/cloudquery/cloudquery/issues/12595)) ([c1c0949](https://github.com/cloudquery/cloudquery/commit/c1c09490b17f2e64435e05d745890cdb8b22310d))
* **deps:** Update github.com/cockroachdb/cockroachdb-parser digest to 302c9ad ([#12664](https://github.com/cloudquery/cloudquery/issues/12664)) ([924509c](https://github.com/cloudquery/cloudquery/commit/924509c409fcf008c93f67fc6a0c5dcf4b2bddc5))
* **deps:** Update github.com/gocarina/gocsv digest to 99d496c ([#12667](https://github.com/cloudquery/cloudquery/issues/12667)) ([428f719](https://github.com/cloudquery/cloudquery/commit/428f71968fd0ebe8a20e99c771f647267f614894))
* **deps:** Update github.com/petermattis/goid digest to 80aa455 ([#12669](https://github.com/cloudquery/cloudquery/issues/12669)) ([a140396](https://github.com/cloudquery/cloudquery/commit/a140396153d62d3e68646d58a7749426aa2cc9fe))
* Detecting conditions for CIS AWS v1.5.0 Section 1 ([#12670](https://github.com/cloudquery/cloudquery/issues/12670)) ([f7bd160](https://github.com/cloudquery/cloudquery/commit/f7bd160e69fc3a222968df0c5f77ee4ed460fb51))
* **resources:** Handle Pagination for AWS Code Commit Repositories ([#12653](https://github.com/cloudquery/cloudquery/issues/12653)) ([6f37e56](https://github.com/cloudquery/cloudquery/commit/6f37e560535357f1c0a34a5acb2d631c1785e742))

## [22.2.0](https://github.com/cloudquery/cloudquery/compare/plugins-source-aws-v22.1.0...plugins-source-aws-v22.2.0) (2023-07-27)


### This Release has the Following Changes to Tables
- Table `aws_cloudformation_stack_instance_resource_drifts` was added
- Table `aws_cloudformation_stack_instance_summaries` was added
- Table `aws_emr_studio_session_mappings` was added
- Table `aws_emr_studios` was added
- Table `aws_route53recoverycontrolconfig_clusters` was added
- Table `aws_route53recoverycontrolconfig_control_panels` was added
- Table `aws_route53recoverycontrolconfig_routing_controls` was added
- Table `aws_route53recoverycontrolconfig_safety_rules` was added
- Table `aws_route53recoveryreadiness_cells` was added
- Table `aws_route53recoveryreadiness_readiness_checks` was added
- Table `aws_route53recoveryreadiness_recovery_groups` was added
- Table `aws_route53recoveryreadiness_resource_sets` was added
- Table `aws_s3_multi_region_access_points` was added

### Features

* **resources-s3:** Add support for S3 Multi Region Access Points ([#12525](https://github.com/cloudquery/cloudquery/issues/12525)) ([00c6a3f](https://github.com/cloudquery/cloudquery/commit/00c6a3fa781c29d1e47216d79dc098d6ad1777c3))
* **resources:** Add additional emr resources for studio and studio session mapping ([#12529](https://github.com/cloudquery/cloudquery/issues/12529)) ([66a20a5](https://github.com/cloudquery/cloudquery/commit/66a20a5b2f97ebadee017b165fd413dfaa6d4086))
* **resources:** Add Support for Amazon Route 53 ARC Recovery Control Configuration ([#12460](https://github.com/cloudquery/cloudquery/issues/12460)) ([8c3109b](https://github.com/cloudquery/cloudquery/commit/8c3109b9deaad7cfa82fd0ff32f2ab4ad842cc88))
* **resources:** Add Support for Cloudformation Stack Instance Summaries and Stack Instance Resource Drifts ([#12495](https://github.com/cloudquery/cloudquery/issues/12495)) ([e1d9a74](https://github.com/cloudquery/cloudquery/commit/e1d9a746718cad26473e8a6a117ff55f381ecffe)), closes [#12461](https://github.com/cloudquery/cloudquery/issues/12461)
* **resources:** Add Support for Route53 ARC Recovery Readiness ([#12459](https://github.com/cloudquery/cloudquery/issues/12459)) ([7d6d4c0](https://github.com/cloudquery/cloudquery/commit/7d6d4c0d55d94f5ec8e34518154fd4c45b7ccbf0))

## [22.1.0](https://github.com/cloudquery/cloudquery/compare/plugins-source-aws-v22.0.1...plugins-source-aws-v22.1.0) (2023-07-25)


### Features

* **aws-services:** Support newly added regions ([#12463](https://github.com/cloudquery/cloudquery/issues/12463)) ([0c6f414](https://github.com/cloudquery/cloudquery/commit/0c6f4142d7f2d75a6804edc4b755340570dc32c2))
* **aws:** Improve error message ([#12456](https://github.com/cloudquery/cloudquery/issues/12456)) ([ce05f5c](https://github.com/cloudquery/cloudquery/commit/ce05f5cda2c018bb478ac1e9e6c0c120241a6936)), closes [#12169](https://github.com/cloudquery/cloudquery/issues/12169)


### Bug Fixes

* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to 10df4b9 ([#12443](https://github.com/cloudquery/cloudquery/issues/12443)) ([e385283](https://github.com/cloudquery/cloudquery/commit/e38528309f862f37bc7e278f9b69cf92d5aa5bd5))
* Resolve attributes for S3 buckets with no policy status ([#12457](https://github.com/cloudquery/cloudquery/issues/12457)) ([0593de9](https://github.com/cloudquery/cloudquery/commit/0593de92c2a2bcd3cdb1800ecca8f02a2015e087)), closes [#12393](https://github.com/cloudquery/cloudquery/issues/12393)
* **resources-web-acls:** Ignore `nil` responses from `GetWebACLForResource` ([#12454](https://github.com/cloudquery/cloudquery/issues/12454)) ([ec8714d](https://github.com/cloudquery/cloudquery/commit/ec8714d6c27153e7115306e8164884cddb7bd205))

## [22.0.1](https://github.com/cloudquery/cloudquery/compare/plugins-source-aws-v22.0.0...plugins-source-aws-v22.0.1) (2023-07-22)


### This Release has the Following Changes to Tables
- Table `aws_cloudformation_template_summaries`: column added with name `warnings` and type `json`
- Table `aws_ec2_instance_types`: column added with name `nitro_tpm_info` and type `json`
- Table `aws_ec2_instance_types`: column added with name `nitro_tpm_support` and type `utf8`
- Table `aws_rds_db_snapshots`: column added with name `db_system_id` and type `utf8`
- Table `aws_route53resolver_resolver_endpoints`: column added with name `outpost_arn` and type `utf8`
- Table `aws_route53resolver_resolver_endpoints`: column added with name `preferred_instance_type` and type `utf8`

### Bug Fixes

* **aws:** Make `aws_s3_buckets` properties nullable to clarify successful resolver results ([#12432](https://github.com/cloudquery/cloudquery/issues/12432)) ([d61502b](https://github.com/cloudquery/cloudquery/commit/d61502bab088c804c33584239635830489bd12b3))
* **deps:** Update AWS modules ([#12441](https://github.com/cloudquery/cloudquery/issues/12441)) ([9d7f2df](https://github.com/cloudquery/cloudquery/commit/9d7f2df624c2ea459a4b2e796267dd96ccaaf543))
* **resources-acm-certificates:** Filter by all key usages ([#12436](https://github.com/cloudquery/cloudquery/issues/12436)) ([dc6504e](https://github.com/cloudquery/cloudquery/commit/dc6504eb5199b17a1743dbde4dc292b40d763e7a))

## [22.0.0](https://github.com/cloudquery/cloudquery/compare/plugins-source-aws-v21.1.0...plugins-source-aws-v22.0.0) (2023-07-20)


### This Release has the Following Changes to Tables
- Table `aws_appconfig_deployment_strategies` was added
- Table `aws_dynamodb_table_continuous_backups`: primary key constraint added to column `table_arn` (:warning: breaking)
- Table `aws_dynamodb_table_continuous_backups`: primary key constraint removed from column `_cq_id` (:warning: breaking)
- Table `aws_lambda_function_versions`: primary key constraint added to column `function_arn` (:warning: breaking)
- Table `aws_lambda_function_versions`: primary key constraint added to column `version` (:warning: breaking)
- Table `aws_lambda_function_versions`: primary key constraint removed from column `_cq_id` (:warning: breaking)
- Table `aws_regions`: primary key constraint added to column `account_id` (:warning: breaking)
- Table `aws_regions`: primary key constraint added to column `region` (:warning: breaking)
- Table `aws_regions`: primary key constraint removed from column `_cq_id` (:warning: breaking)
- Table `aws_s3_bucket_encryption_rules`: primary key constraint added to column `bucket_arn` (:warning: breaking)
- Table `aws_s3_bucket_encryption_rules`: primary key constraint removed from column `_cq_id` (:warning: breaking)

### ⚠ BREAKING CHANGES

* **aws:** Define composite primary key for regions ([#12415](https://github.com/cloudquery/cloudquery/issues/12415))
* **aws:** Define primary key for s3_bucket_encryption_rules ([#12408](https://github.com/cloudquery/cloudquery/issues/12408))
* **aws:** Define primary key for dynamodb_table_continuous_backups ([#12409](https://github.com/cloudquery/cloudquery/issues/12409))
* **aws:** Define composite primary key for lambda_function_versions ([#12402](https://github.com/cloudquery/cloudquery/issues/12402))

### Features

* **aws:** Define composite primary key for lambda_function_versions ([#12402](https://github.com/cloudquery/cloudquery/issues/12402)) ([d1add18](https://github.com/cloudquery/cloudquery/commit/d1add18c5435c01579cbe86eb46164e656d4cd66))
* **aws:** Define composite primary key for regions ([#12415](https://github.com/cloudquery/cloudquery/issues/12415)) ([681ea97](https://github.com/cloudquery/cloudquery/commit/681ea9731fc13dcf0965e97513cfc209589b4765))
* **aws:** Define primary key for dynamodb_table_continuous_backups ([#12409](https://github.com/cloudquery/cloudquery/issues/12409)) ([7d4a657](https://github.com/cloudquery/cloudquery/commit/7d4a657afaed5e129e1650b06595abd8d333572d))
* **aws:** Define primary key for s3_bucket_encryption_rules ([#12408](https://github.com/cloudquery/cloudquery/issues/12408)) ([a026b98](https://github.com/cloudquery/cloudquery/commit/a026b989ad6cafd61f76374dcc5fa6bc43a57073))


### Bug Fixes

* **sync:** Pass `DeterministicCQID` option to scheduler ([#12424](https://github.com/cloudquery/cloudquery/issues/12424)) ([eaac2e6](https://github.com/cloudquery/cloudquery/commit/eaac2e684bd6e9744a38be8eef3a1f5e77d163f5))
* Use correct name for `DeploymentStrategies` table ([#12428](https://github.com/cloudquery/cloudquery/issues/12428)) ([dda4205](https://github.com/cloudquery/cloudquery/commit/dda42056ec2c3e0038ca72a95469cf75cf9d22a4))

## [21.1.0](https://github.com/cloudquery/cloudquery/compare/plugins-source-aws-v21.0.0...plugins-source-aws-v21.1.0) (2023-07-18)


### This Release has the Following Changes to Tables
- Table `aws_backup_jobs` was added

### Features

* **resources:** Add Backup Jobs ([#12389](https://github.com/cloudquery/cloudquery/issues/12389)) ([4e6825c](https://github.com/cloudquery/cloudquery/commit/4e6825ce431876ac21678bc170dc7c366c79fb00))


### Bug Fixes

* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.2.3 ([#12307](https://github.com/cloudquery/cloudquery/issues/12307)) ([8f14e4d](https://github.com/cloudquery/cloudquery/commit/8f14e4de7bf4d4c833f501135ea0610916a42f8b))

## [21.0.0](https://github.com/cloudquery/cloudquery/compare/plugins-source-aws-v20.1.0...plugins-source-aws-v21.0.0) (2023-07-18)


### This Release has the Following Changes to Tables
- Table `aws_servicecatalog_launch_paths` was added
- Table `aws_servicecatalog_portfolios`: column `created_time` removed from table (:warning: breaking)
- Table `aws_servicecatalog_portfolios`: column `description` removed from table (:warning: breaking)
- Table `aws_servicecatalog_portfolios`: column `display_name` removed from table (:warning: breaking)
- Table `aws_servicecatalog_portfolios`: column `id` removed from table (:warning: breaking)
- Table `aws_servicecatalog_portfolios`: column `provider_name` removed from table (:warning: breaking)
- Table `aws_servicecatalog_portfolios`: column added with name `budgets` and type `json`
- Table `aws_servicecatalog_portfolios`: column added with name `portfolio_detail` and type `json`
- Table `aws_servicecatalog_portfolios`: column added with name `region` and type `utf8`
- Table `aws_servicecatalog_portfolios`: column added with name `tag_options` and type `json`
- Table `aws_servicecatalog_products`: column `created_time` removed from table (:warning: breaking)
- Table `aws_servicecatalog_products`: column `product_arn` removed from table (:warning: breaking)
- Table `aws_servicecatalog_products`: column `product_view_summary` removed from table (:warning: breaking)
- Table `aws_servicecatalog_products`: column `source_connection` removed from table (:warning: breaking)
- Table `aws_servicecatalog_products`: column `status` removed from table (:warning: breaking)
- Table `aws_servicecatalog_products`: column added with name `budgets` and type `json`
- Table `aws_servicecatalog_products`: column added with name `product_view_detail` and type `json`
- Table `aws_servicecatalog_products`: column added with name `provisioning_artifact_summaries` and type `json`
- Table `aws_servicecatalog_products`: column added with name `region` and type `utf8`
- Table `aws_servicecatalog_products`: column added with name `tag_options` and type `json`
- Table `aws_servicecatalog_provisioned_products`: column added with name `region` and type `utf8`
- Table `aws_servicecatalog_provisioning_artifacts` was added
- Table `aws_servicecatalog_provisioning_parameters` was added

### ⚠ BREAKING CHANGES

* **aws:** Fix Service Catalog resources ([#12117](https://github.com/cloudquery/cloudquery/issues/12117))

### Features

* **aws-services:** Support newly added regions ([#12274](https://github.com/cloudquery/cloudquery/issues/12274)) ([1121683](https://github.com/cloudquery/cloudquery/commit/112168331886438191e4b46c61ad49f6fa748130))


### Bug Fixes

* **aws:** Fix Service Catalog resources ([#12117](https://github.com/cloudquery/cloudquery/issues/12117)) ([c942005](https://github.com/cloudquery/cloudquery/commit/c942005be14abeab4cc7eb489747a9a02151f019))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.2.0 ([#12256](https://github.com/cloudquery/cloudquery/issues/12256)) ([eaec331](https://github.com/cloudquery/cloudquery/commit/eaec33165345ad51fdb6ddbffbf8a1199ebd6384))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.2.1 ([#12272](https://github.com/cloudquery/cloudquery/issues/12272)) ([557ca69](https://github.com/cloudquery/cloudquery/commit/557ca69a7dee9dabb80e6afb6f41f205fd8a80d8))

## [20.1.0](https://github.com/cloudquery/cloudquery/compare/plugins-source-aws-v20.0.1...plugins-source-aws-v20.1.0) (2023-07-17)


### This Release has the Following Changes to Tables
- Table `aws_appconfig_applications` was added
- Table `aws_appconfig_configuration_profiles` was added
- Table `aws_appconfig_environments` was added
- Table `aws_appconfig_hosted_configuration_versions` was added
- Table `aws_appstream_app_blocks`: column added with name `app_block_errors` and type `json`
- Table `aws_appstream_app_blocks`: column added with name `packaging_type` and type `utf8`
- Table `aws_appstream_app_blocks`: column added with name `post_setup_script_details` and type `json`
- Table `aws_appstream_app_blocks`: column added with name `state` and type `utf8`
- Table `aws_ec2_instance_types`: column added with name `nitro_enclaves_support` and type `utf8`

### Features

* **aws:** Add support for `App Config` ([#12150](https://github.com/cloudquery/cloudquery/issues/12150)) ([48038aa](https://github.com/cloudquery/cloudquery/commit/48038aa882c7ec33e47cf4fabf299249ccb10683))


### Bug Fixes

* **deps:** Update AWS modules ([#12216](https://github.com/cloudquery/cloudquery/issues/12216)) ([d2e77a7](https://github.com/cloudquery/cloudquery/commit/d2e77a7db3da96eef586ff04243e541d68e937dd))
* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to 8e2219b ([#12220](https://github.com/cloudquery/cloudquery/issues/12220)) ([24e8fb5](https://github.com/cloudquery/cloudquery/commit/24e8fb588740896fe11a660e8b80231e741b753c))

## [20.0.1](https://github.com/cloudquery/cloudquery/compare/plugins-source-aws-v20.0.0...plugins-source-aws-v20.0.1) (2023-07-14)


### Bug Fixes

* **aws:** Make S3 column resolvers non-blocking ([#12165](https://github.com/cloudquery/cloudquery/issues/12165)) ([c384406](https://github.com/cloudquery/cloudquery/commit/c3844063a93ab4970e0400e870da6f03b702151f))
* **deps:** Update module github.com/cloudquery/plugin-pb-go to v1.7.0 ([#12166](https://github.com/cloudquery/cloudquery/issues/12166)) ([94390dd](https://github.com/cloudquery/cloudquery/commit/94390dde19d0c37fee9d035219d62f6ae7edb127))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.1.0 ([#12174](https://github.com/cloudquery/cloudquery/issues/12174)) ([80f0289](https://github.com/cloudquery/cloudquery/commit/80f02892a4cf876c4bf4dd4fd9367afb3770ad26))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.1.1 ([#12185](https://github.com/cloudquery/cloudquery/issues/12185)) ([cfaff16](https://github.com/cloudquery/cloudquery/commit/cfaff16d89800235b6e3015eeb6957d5783d1393))
* **deps:** Upgrade source plugins to SDK v4.0.0 release ([#12135](https://github.com/cloudquery/cloudquery/issues/12135)) ([c20a111](https://github.com/cloudquery/cloudquery/commit/c20a111d591101fb1bbc42292accc953af38e8a6))

## [20.0.0](https://github.com/cloudquery/cloudquery/compare/plugins-source-aws-v19.2.0...plugins-source-aws-v20.0.0) (2023-07-12)


### This Release has the Following Changes to Tables
- Table `aws_ec2_eips`: primary key constraint added to column `account_id` (:warning: breaking)
- Table `aws_ec2_eips`: primary key constraint added to column `allocation_id` (:warning: breaking)
- Table `aws_ec2_eips`: primary key constraint added to column `region` (:warning: breaking)
- Table `aws_ec2_eips`: primary key constraint removed from column `_cq_id` (:warning: breaking)
- Table `aws_networkmanager_global_networks` was added
- Table `aws_networkmanager_links` was added
- Table `aws_networkmanager_sites` was added
- Table `aws_networkmanager_transit_gateway_registrations` was added

### ⚠ BREAKING CHANGES

* **aws:** Define primary key for eips ([#11728](https://github.com/cloudquery/cloudquery/issues/11728))
* Upgrades the awspricing source plugin to use plugin-sdk v4. This version does not contain any user-facing breaking changes, but because it is now using CloudQuery gRPC protocol v3, it does require use of a destination plugin that also supports protocol v3. All recent destination plugin versions support this.

### Features

* Add table_options support for aws_securityhub_findings table ([#11955](https://github.com/cloudquery/cloudquery/issues/11955)) ([c9eff12](https://github.com/cloudquery/cloudquery/commit/c9eff1252fb2a6768373b954dabc291d52904fb6))
* **aws-policies:** Add in AWS security account contact query ([#11729](https://github.com/cloudquery/cloudquery/issues/11729)) ([c9d7294](https://github.com/cloudquery/cloudquery/commit/c9d7294a6294daf6906fe68dd623aa7b3df87b4b))
* **aws-policies:** Add sns logging of delivery status to AWS Policies ([#12074](https://github.com/cloudquery/cloudquery/issues/12074)) ([80f0b88](https://github.com/cloudquery/cloudquery/commit/80f0b88eb4f751175d396278440560867dfe2d6c))
* **aws-policies:** Update sqs encryption for aws foundational security policies ([#11777](https://github.com/cloudquery/cloudquery/issues/11777)) ([30d415c](https://github.com/cloudquery/cloudquery/commit/30d415ca186a136da420ad75e6bf3d5457c12ab2))
* **aws-policies:** Update ssm queries for aws policies ([#12067](https://github.com/cloudquery/cloudquery/issues/12067)) ([2b9180f](https://github.com/cloudquery/cloudquery/commit/2b9180f5024f9d949dcf49ba216aa9136ffb469a))
* **aws-services:** Support newly added regions ([#11922](https://github.com/cloudquery/cloudquery/issues/11922)) ([6680d7a](https://github.com/cloudquery/cloudquery/commit/6680d7a03353c8b0b807b682495af312109aa7c1))
* **aws-services:** Support newly added regions ([#12120](https://github.com/cloudquery/cloudquery/issues/12120)) ([15ea38c](https://github.com/cloudquery/cloudquery/commit/15ea38c54fd07368ecb3c570c4ae4a40dbf4cde8))
* **aws:** Add Support for `ecs:ListTasks` in `table_options` ([#11986](https://github.com/cloudquery/cloudquery/issues/11986)) ([3016c16](https://github.com/cloudquery/cloudquery/commit/3016c16d6b59ba4c5c1182ee9601b0bda8fd3591)), closes [#11981](https://github.com/cloudquery/cloudquery/issues/11981)
* **aws:** Define primary key for eips ([#11728](https://github.com/cloudquery/cloudquery/issues/11728)) ([fa48d4a](https://github.com/cloudquery/cloudquery/commit/fa48d4a65fcd8150e29f8c85720554ec1c58cfbb))
* **aws:** Support networkmanager resources ([#12123](https://github.com/cloudquery/cloudquery/issues/12123)) ([a642ce0](https://github.com/cloudquery/cloudquery/commit/a642ce04f09a73341e8916538c3f4411a7c390f8))
* Upgrades the awspricing source plugin to use plugin-sdk v4. This version does not contain any user-facing breaking changes, but because it is now using CloudQuery gRPC protocol v3, it does require use of a destination plugin that also supports protocol v3. All recent destination plugin versions support this. ([7d50d29](https://github.com/cloudquery/cloudquery/commit/7d50d29e6fcdf44579112ecfbcb92908a8bc1247))


### Bug Fixes

* **aws:** Skip fetching tags for `aws_kafka_cluster_operations` ([#11973](https://github.com/cloudquery/cloudquery/issues/11973)) ([2b62ba4](https://github.com/cloudquery/cloudquery/commit/2b62ba4c3f8ae801fff76f7ecf626a1e2abdfb84))
* **aws:** Validate table relations not just top level table ([#12121](https://github.com/cloudquery/cloudquery/issues/12121)) ([e13d931](https://github.com/cloudquery/cloudquery/commit/e13d9317f1c3032278fd4c3fa81004126d083772))
* **deps:** Update github.com/apache/arrow/go/v13 digest to 5a06b2e ([#11857](https://github.com/cloudquery/cloudquery/issues/11857)) ([43c2f5f](https://github.com/cloudquery/cloudquery/commit/43c2f5f3a893e5286f67c4943a9d1bc2736e2aeb))
* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to 0a52533 ([#12091](https://github.com/cloudquery/cloudquery/issues/12091)) ([927cefa](https://github.com/cloudquery/cloudquery/commit/927cefa943ec3969a2ec39b628bc1eba545a2108))
* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to a2a76eb ([#12104](https://github.com/cloudquery/cloudquery/issues/12104)) ([311f474](https://github.com/cloudquery/cloudquery/commit/311f4749af2491a606f29483190717a5fe238da6))
* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to df3b664 ([#11882](https://github.com/cloudquery/cloudquery/issues/11882)) ([9635b22](https://github.com/cloudquery/cloudquery/commit/9635b22b10a2cd9ca0f91819cffb7f4ba75dc2d9))
* **deps:** Update github.com/cockroachdb/cockroachdb-parser digest to c9c144e ([#11863](https://github.com/cloudquery/cloudquery/issues/11863)) ([1547efd](https://github.com/cloudquery/cloudquery/commit/1547efd045abbf8905f5a7bf5856571603bd64d9))
* **deps:** Update github.com/cockroachdb/logtags digest to 21c5414 ([#11864](https://github.com/cloudquery/cloudquery/issues/11864)) ([da48b1f](https://github.com/cloudquery/cloudquery/commit/da48b1fc86576ea5777505e5bb59ecaf0febf7ca))
* **deps:** Update github.com/gocarina/gocsv digest to 99d496c ([#11865](https://github.com/cloudquery/cloudquery/issues/11865)) ([c3de686](https://github.com/cloudquery/cloudquery/commit/c3de686895d4d15d4687ea7b505466195b09f546))
* **deps:** Update github.com/golang/geo digest to 6adc566 ([#11866](https://github.com/cloudquery/cloudquery/issues/11866)) ([edb7ed8](https://github.com/cloudquery/cloudquery/commit/edb7ed83896842e6f174079ba020f18d713f6f91))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/networkfirewall to v1.28.3 ([#12079](https://github.com/cloudquery/cloudquery/issues/12079)) ([a27fa21](https://github.com/cloudquery/cloudquery/commit/a27fa21c7252036ec4e3d79c5133b98e61a4acd0))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/securityhub to v1.33.2 ([#12081](https://github.com/cloudquery/cloudquery/issues/12081)) ([e77f93e](https://github.com/cloudquery/cloudquery/commit/e77f93ea85457943b1cd2021f5dfd4ac69d37785))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/servicediscovery to v1.21.7 ([#12082](https://github.com/cloudquery/cloudquery/issues/12082)) ([01f8b59](https://github.com/cloudquery/cloudquery/commit/01f8b599f992dc1aadf955f5b04e67ce5d09affb))
* **deps:** Update module github.com/cloudquery/plugin-pb-go to v1.5.0 ([#11850](https://github.com/cloudquery/cloudquery/issues/11850)) ([3255857](https://github.com/cloudquery/cloudquery/commit/3255857938bf16862d52491f5c2a8a0fa53faef0))
* **deps:** Update module github.com/cloudquery/plugin-pb-go to v1.6.0 ([#11916](https://github.com/cloudquery/cloudquery/issues/11916)) ([421e752](https://github.com/cloudquery/cloudquery/commit/421e7529360965175c8d156ff006d2b703ee9da2))
* **postgresql:** Rerun release please ([#12002](https://github.com/cloudquery/cloudquery/issues/12002)) ([9d12843](https://github.com/cloudquery/cloudquery/commit/9d12843462d1019d26bc239f8f928bf5f62940cf))

## [19.2.0](https://github.com/cloudquery/cloudquery/compare/plugins-source-aws-v19.1.0...plugins-source-aws-v19.2.0) (2023-06-30)


### This Release has the Following Changes to Tables
- Table `aws_codecommit_repositories` was added
- Table `aws_route53resolver_firewall_configs` was added
- Table `aws_route53resolver_firewall_domain_lists` was added
- Table `aws_route53resolver_firewall_rule_group_associations` was added
- Table `aws_route53resolver_firewall_rule_groups` was added
- Table `aws_route53resolver_resolver_endpoints` was added
- Table `aws_route53resolver_resolver_query_log_config_associations` was added
- Table `aws_route53resolver_resolver_query_log_configs` was added
- Table `aws_route53resolver_resolver_rule_associations` was added
- Table `aws_route53resolver_resolver_rules` was added

### Features

* **aws:** Add AWS CodeCommit Repositories ([#11827](https://github.com/cloudquery/cloudquery/issues/11827)) ([a198fd8](https://github.com/cloudquery/cloudquery/commit/a198fd8e5890d289a464c580497a72f139717e5c)), closes [#11819](https://github.com/cloudquery/cloudquery/issues/11819)
* **aws:** Add support for Route53resolver service ([#11818](https://github.com/cloudquery/cloudquery/issues/11818)) ([9bbd610](https://github.com/cloudquery/cloudquery/commit/9bbd610d3585feab47332e318988bd7dcee14fce))


### Bug Fixes

* **aws:** Add test to ensure documentation is unique for each table ([#11816](https://github.com/cloudquery/cloudquery/issues/11816)) ([107f98b](https://github.com/cloudquery/cloudquery/commit/107f98beaf5ee5bb9e757018492fe9d7e20d6324))
* **aws:** Codebuild builds ([#11805](https://github.com/cloudquery/cloudquery/issues/11805)) ([8c560f0](https://github.com/cloudquery/cloudquery/commit/8c560f0626410b0802b28ff037b4cca0c90beeaa))
* **aws:** Fix CodeBuild SourceCredentials Mock Test ([#11808](https://github.com/cloudquery/cloudquery/issues/11808)) ([cce4376](https://github.com/cloudquery/cloudquery/commit/cce437606797c1c49626adaf13c12e169d019185))
* **aws:** Fix Timestream Regional issue ([#11795](https://github.com/cloudquery/cloudquery/issues/11795)) ([c53399a](https://github.com/cloudquery/cloudquery/commit/c53399a119aee33d9e314e32030c4b8740966ad2))

## [19.1.0](https://github.com/cloudquery/cloudquery/compare/plugins-source-aws-v19.0.0...plugins-source-aws-v19.1.0) (2023-06-27)


### This Release has the Following Changes to Tables
- Table `aws_codeartifact_domains` was added
- Table `aws_codeartifact_repositories` was added
- Table `aws_codebuild_builds` was added
- Table `aws_codebuild_source_credentials` was added
- Table `aws_detective_graph_members` was added
- Table `aws_detective_graphs` was added
- Table `aws_ec2_hosts`: column added with name `asset_id` and type `utf8`
- Table `aws_ec2_vpn_connections` was added
- Table `aws_ecr_pull_through_cache_rules` was added
- Table `aws_eks_cluster_addons` was added
- Table `aws_eks_cluster_oidc_identity_provider_configs` was added
- Table `aws_inspector2_findings`: column added with name `code_vulnerability_details` and type `json`
- Table `aws_inspector2_findings`: column added with name `epss` and type `json`
- Table `aws_mq_broker_users`: column added with name `replication_user` and type `bool`
- Table `aws_mq_brokers`: column added with name `data_replication_metadata` and type `json`
- Table `aws_mq_brokers`: column added with name `data_replication_mode` and type `utf8`
- Table `aws_mq_brokers`: column added with name `pending_data_replication_metadata` and type `json`
- Table `aws_mq_brokers`: column added with name `pending_data_replication_mode` and type `utf8`
- Table `aws_networkfirewall_firewalls` was added
- Table `aws_networkfirewall_tls_inspection_configurations` was added
- Table `aws_redshift_clusters`: column added with name `custom_domain_certificate_arn` and type `utf8`
- Table `aws_redshift_clusters`: column added with name `custom_domain_certificate_expiry_date` and type `timestamp[us, tz=UTC]`
- Table `aws_redshift_clusters`: column added with name `custom_domain_name` and type `utf8`
- Table `aws_servicediscovery_instances` was added
- Table `aws_servicediscovery_namespaces` was added
- Table `aws_servicediscovery_services` was added
- Table `aws_signer_signing_profiles` was added
- Table `aws_stepfunctions_executions`: column added with name `state_machine_alias_arn` and type `utf8`
- Table `aws_stepfunctions_executions`: column added with name `state_machine_version_arn` and type `utf8`
- Table `aws_stepfunctions_map_run_executions`: column added with name `state_machine_alias_arn` and type `utf8`
- Table `aws_stepfunctions_map_run_executions`: column added with name `state_machine_version_arn` and type `utf8`
- Table `aws_stepfunctions_state_machines`: column added with name `description` and type `utf8`
- Table `aws_stepfunctions_state_machines`: column added with name `revision_id` and type `utf8`
- Table `aws_transfer_servers`: column added with name `structured_log_destinations` and type `list<item: utf8, nullable>`
- Table `aws_wellarchitected_lens_review_improvements` was added
- Table `aws_wellarchitected_lens_reviews` was added
- Table `aws_wellarchitected_lenses` was added
- Table `aws_wellarchitected_share_invitations` was added
- Table `aws_wellarchitected_workload_milestones` was added
- Table `aws_wellarchitected_workload_shares` was added
- Table `aws_wellarchitected_workloads` was added

### Features

* Add AWS Well-Architected resources ([#11697](https://github.com/cloudquery/cloudquery/issues/11697)) ([83174fe](https://github.com/cloudquery/cloudquery/commit/83174fef5f43b39aa99129c2bf7213aea7e280be)), closes [#11664](https://github.com/cloudquery/cloudquery/issues/11664)
* **aws-services:** Support newly added regions ([#11778](https://github.com/cloudquery/cloudquery/issues/11778)) ([afbf0ec](https://github.com/cloudquery/cloudquery/commit/afbf0ecc7c24f5917207845602a24af7942e9c67))
* **aws:** Add Support for `builds` and `source credentials` for AWS CodeBuild  ([#11705](https://github.com/cloudquery/cloudquery/issues/11705)) ([4eed4dc](https://github.com/cloudquery/cloudquery/commit/4eed4dc1816ddb5c4b5e0cda50489ef912bd0f94))
* **aws:** Add Support for AWS CodeArtifact Domains and Repositories ([#11698](https://github.com/cloudquery/cloudquery/issues/11698)) ([1838151](https://github.com/cloudquery/cloudquery/commit/18381512bb1d738521601b9f514da220aad2f366))
* **aws:** Add Support for AWS Detective `Graphs` and `GraphMembers` ([#11767](https://github.com/cloudquery/cloudquery/issues/11767)) ([b40d97d](https://github.com/cloudquery/cloudquery/commit/b40d97d384650d69ec60dce8b4c2e8564d065e24))
* **aws:** Add support for AWS Signer Profiles ([#11765](https://github.com/cloudquery/cloudquery/issues/11765)) ([e03b797](https://github.com/cloudquery/cloudquery/commit/e03b797bcdd3a477f849f2b166eb8fbced248f00))
* **aws:** Add Support for Cloud Map (servicediscovery) resources ([#11702](https://github.com/cloudquery/cloudquery/issues/11702)) ([0cefa8c](https://github.com/cloudquery/cloudquery/commit/0cefa8cb16999470fa4d3ad26c5617724346d265))
* **aws:** Add support for EC2 VPN Connections ([#11769](https://github.com/cloudquery/cloudquery/issues/11769)) ([0f0c340](https://github.com/cloudquery/cloudquery/commit/0f0c340cdb6fa772f3b3a92580430a0041e67546))
* **aws:** Add Support for ECR `PullThroughCacheRules` ([#11770](https://github.com/cloudquery/cloudquery/issues/11770)) ([9d72446](https://github.com/cloudquery/cloudquery/commit/9d72446e1346fca7f7eef22844248700716d8259))
* **aws:** Add support for EKS `AddOns` and `IdentityProviderConfigs` ([#11764](https://github.com/cloudquery/cloudquery/issues/11764)) ([24c348d](https://github.com/cloudquery/cloudquery/commit/24c348d10806c4c238db2c3508f793193c201fe8))
* **aws:** Add Support for Network firewall resources `Firewalls` and `TLSInspectionConfigurations` ([#11776](https://github.com/cloudquery/cloudquery/issues/11776)) ([6e7ae44](https://github.com/cloudquery/cloudquery/commit/6e7ae4479d48acbd65b3e65a4b1f9d819925093f))
* **aws:** Update all AWS dependencies ([#11783](https://github.com/cloudquery/cloudquery/issues/11783)) ([18d9fa7](https://github.com/cloudquery/cloudquery/commit/18d9fa705e34661e3099073cdacd0088a7ef82d0))


### Bug Fixes

* **aws:** AppRunner Tag errors ([#11786](https://github.com/cloudquery/cloudquery/issues/11786)) ([d2d333d](https://github.com/cloudquery/cloudquery/commit/d2d333d50dbff848779edffc16dde8b632c595fe))
* **aws:** Ensure that all certificates are synced ([#11761](https://github.com/cloudquery/cloudquery/issues/11761)) ([7e5b201](https://github.com/cloudquery/cloudquery/commit/7e5b2014ad44fb045cb82fad94519ba6ab0f63b2))
* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to 0656028 ([#11739](https://github.com/cloudquery/cloudquery/issues/11739)) ([7a6ad49](https://github.com/cloudquery/cloudquery/commit/7a6ad49e8402d51e914d6fdc444956c89db91ad3))
* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to 8366a22 ([#11717](https://github.com/cloudquery/cloudquery/issues/11717)) ([8eeff5b](https://github.com/cloudquery/cloudquery/commit/8eeff5b17486d72845f830b99983f950fee7f5a0))
* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to 95d3199 ([#11708](https://github.com/cloudquery/cloudquery/issues/11708)) ([03f214f](https://github.com/cloudquery/cloudquery/commit/03f214f3dfd719b74ce9eb698ba255a8cf7528c7))
* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to f060192 ([#11730](https://github.com/cloudquery/cloudquery/issues/11730)) ([c7019c2](https://github.com/cloudquery/cloudquery/commit/c7019c26c311f29b66c90fc5d461a0daf71d191c))
* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to f0dffc6 ([#11689](https://github.com/cloudquery/cloudquery/issues/11689)) ([18ac0e9](https://github.com/cloudquery/cloudquery/commit/18ac0e9dbef31d06701f1f13d263ad840ac60c5e))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/accessanalyzer to v1.19.14 ([#11733](https://github.com/cloudquery/cloudquery/issues/11733)) ([ab8242f](https://github.com/cloudquery/cloudquery/commit/ab8242f8f94894f94d1dddbacdcfd2e6fcc7c174))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/account to v1.10.8 ([#11734](https://github.com/cloudquery/cloudquery/issues/11734)) ([94e285b](https://github.com/cloudquery/cloudquery/commit/94e285b3ed5b3ba485d1d7b36109659506d85610))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/acm to v1.17.13 ([#11735](https://github.com/cloudquery/cloudquery/issues/11735)) ([c763315](https://github.com/cloudquery/cloudquery/commit/c76331514e13bc2771b475bc911c0e6620880c50))
* **deps:** Update module github.com/cloudquery/plugin-pb-go to v1.2.0 ([#11720](https://github.com/cloudquery/cloudquery/issues/11720)) ([7ef521d](https://github.com/cloudquery/cloudquery/commit/7ef521db1423c6f0de197b08c73adf22c896f999))
* **deps:** Update module github.com/cloudquery/plugin-pb-go to v1.2.1 ([#11722](https://github.com/cloudquery/cloudquery/issues/11722)) ([309be72](https://github.com/cloudquery/cloudquery/commit/309be7276d7de157013c281b6fb3934513898b3f))
* **deps:** Update module github.com/cloudquery/plugin-pb-go to v1.3.3 ([#11726](https://github.com/cloudquery/cloudquery/issues/11726)) ([f0ca611](https://github.com/cloudquery/cloudquery/commit/f0ca61195014bde707761a15efa27a92955b59db))
* **deps:** Update module github.com/cloudquery/plugin-pb-go to v1.3.4 ([#11753](https://github.com/cloudquery/cloudquery/issues/11753)) ([cd4fe1c](https://github.com/cloudquery/cloudquery/commit/cd4fe1c54f85f8511252bebd5671361618ddb0d3))

## [19.0.0](https://github.com/cloudquery/cloudquery/compare/plugins-source-aws-v18.4.0...plugins-source-aws-v19.0.0) (2023-06-20)


### This Release has the Following Changes to Tables
- Table `aws_cloudfront_functions` was added
- Table `aws_cloudfront_origin_access_identities` was added
- Table `aws_cloudfront_origin_request_policies` was added
- Table `aws_cloudfront_response_headers_policies` was added
- Table `aws_cloudtrail_channels` was added
- Table `aws_cloudtrail_imports` was added
- Table `aws_ec2_capacity_reservations` was added
- Table `aws_elbv2_load_balancer_web_acls` was added
- Table `aws_elbv2_load_balancers`: column `web_acl_arn` removed from table (:warning: breaking)
- Table `aws_organizations_account_parents` was added
- Table `aws_organizations_organizational_unit_parents` was added
- Table `aws_organizations_organizational_units`: column `account_id` removed from table (:warning: breaking)
- Table `aws_organizations_organizational_units`: column added with name `request_account_id (PK)` and type `utf8` (:warning: breaking)
- Table `aws_s3_buckets`: column added with name `policy_status` and type `json`
- Table `aws_ssoadmin_account_assignments` was removed (:warning: breaking)
- Table `aws_ssoadmin_instances`: primary key constraint added to column `instance_arn` (:warning: breaking)
- Table `aws_ssoadmin_instances`: primary key constraint removed from column `_cq_id` (:warning: breaking)
- Table `aws_ssoadmin_permission_set_account_assignments` was added
- Table `aws_ssoadmin_permission_set_customer_managed_policies` was added
- Table `aws_ssoadmin_permission_set_inline_policies` was added
- Table `aws_ssoadmin_permission_set_managed_policies` was added
- Table `aws_ssoadmin_permission_set_permissions_boundaries` was added
- Table `aws_ssoadmin_permission_sets`: column `inline_policy` removed from table (:warning: breaking)
- Table `aws_ssoadmin_permission_sets`: column added with name `instance_arn (PK)` and type `utf8` (:warning: breaking)
- Table `aws_ssoadmin_permission_sets`: primary key constraint added to column `permission_set_arn` (:warning: breaking)
- Table `aws_ssoadmin_permission_sets`: primary key constraint removed from column `_cq_id` (:warning: breaking)

### ⚠ BREAKING CHANGES

* **aws:** Move `web_acl_arn` to its own table ([#11421](https://github.com/cloudquery/cloudquery/issues/11421))
* **aws:** Add support For fully describing the organizational hierarchy ([#11633](https://github.com/cloudquery/cloudquery/issues/11633))
* **aws:** Support all policy types in Identity Center ([#10985](https://github.com/cloudquery/cloudquery/issues/10985))

### Features

* **aws-services:** Support newly added regions ([#11673](https://github.com/cloudquery/cloudquery/issues/11673)) ([8c0ab9d](https://github.com/cloudquery/cloudquery/commit/8c0ab9d7fc83a874e9df0ac089a14156f740c2a2))
* **aws:** Add PolicyStatus to WrappedBucket ([#11657](https://github.com/cloudquery/cloudquery/issues/11657)) ([55d966a](https://github.com/cloudquery/cloudquery/commit/55d966ab28152945ca76dc73f8c3761cca90e1d9))
* **aws:** Add support for Cloudfront Functions ([#11669](https://github.com/cloudquery/cloudquery/issues/11669)) ([102067a](https://github.com/cloudquery/cloudquery/commit/102067ac11d02ef9f37dcde8aed7f5357ace777f))
* **aws:** Add Support for Cloudtrail Channels ([#11670](https://github.com/cloudquery/cloudquery/issues/11670)) ([0dc13de](https://github.com/cloudquery/cloudquery/commit/0dc13de97deb8c62f53b880eb866c80244e0f1ab))
* **aws:** Add Support for Cloudtrail Imports ([#11671](https://github.com/cloudquery/cloudquery/issues/11671)) ([c908289](https://github.com/cloudquery/cloudquery/commit/c90828998dee1742240ea3a354760f5a7e15b2c9))
* **aws:** Add support for EC2 Capacity Reservations ([#11666](https://github.com/cloudquery/cloudquery/issues/11666)) ([70d6052](https://github.com/cloudquery/cloudquery/commit/70d6052c8305463b802ac25151e12c03136afd24))
* **aws:** Add support For fully describing the organizational hierarchy ([#11633](https://github.com/cloudquery/cloudquery/issues/11633)) ([f66995b](https://github.com/cloudquery/cloudquery/commit/f66995baa71a8bc0b4f2c917bffa29ff174a88b5))
* **aws:** Add Support for more Cloudfront Resources ([#11668](https://github.com/cloudquery/cloudquery/issues/11668)) ([52e6ad9](https://github.com/cloudquery/cloudquery/commit/52e6ad978389f4eb3ab84f554bb2e692385a935c))
* **aws:** Support all policy types in Identity Center ([#10985](https://github.com/cloudquery/cloudquery/issues/10985)) ([a8ab255](https://github.com/cloudquery/cloudquery/commit/a8ab255d369a2b88778da005b82fd7e59a085e00))


### Bug Fixes

* **aws:** Move `web_acl_arn` to its own table ([#11421](https://github.com/cloudquery/cloudquery/issues/11421)) ([cdda682](https://github.com/cloudquery/cloudquery/commit/cdda682263381d0e4a2262e5d32cb8d13ab25d05))
* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to 1e68c51 ([#11637](https://github.com/cloudquery/cloudquery/issues/11637)) ([46043bc](https://github.com/cloudquery/cloudquery/commit/46043bce410f86ba42390a6b190f9232fc2a1ded))
* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to 43638cb ([#11672](https://github.com/cloudquery/cloudquery/issues/11672)) ([3c60bbb](https://github.com/cloudquery/cloudquery/commit/3c60bbbb0233b17f934583766938780745145864))
* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to b0832be ([#11651](https://github.com/cloudquery/cloudquery/issues/11651)) ([71e8c29](https://github.com/cloudquery/cloudquery/commit/71e8c29624494a3e1cd104e46266a610ce57c83c))
* **deps:** Update module github.com/aws/aws-sdk-go-v2 to v1.18.1 ([#11652](https://github.com/cloudquery/cloudquery/issues/11652)) ([4230b52](https://github.com/cloudquery/cloudquery/commit/4230b52a19e91b84fc38348291c371c6c8a735af))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/config to v1.18.27 ([#11653](https://github.com/cloudquery/cloudquery/issues/11653)) ([4b45408](https://github.com/cloudquery/cloudquery/commit/4b454088055dcbd265e6cbb09420f7dae66865b5))
* **deps:** Update module github.com/cloudquery/plugin-pb-go to v1.1.0 ([#11665](https://github.com/cloudquery/cloudquery/issues/11665)) ([d8947c9](https://github.com/cloudquery/cloudquery/commit/d8947c9efa6ab8bf3952ad9d929e8ed81f2dea55))
* Use ServiceAccountRegion multiplexer for aws_availability_zones ([#11686](https://github.com/cloudquery/cloudquery/issues/11686)) ([7f4788f](https://github.com/cloudquery/cloudquery/commit/7f4788fd9eae914b918bc17c80bee350f16408e5))

## [18.4.0](https://github.com/cloudquery/cloudquery/compare/plugins-source-aws-v18.3.0...plugins-source-aws-v18.4.0) (2023-06-15)


### This Release has the Following Changes to Tables
- Table `aws_alpha_cloudwatch_metric_statistics` was added
- Table `aws_alpha_cloudwatch_metrics` was added
- Table `aws_alpha_costexplorer_cost_custom` was added
- Table `aws_applicationautoscaling_scalable_targets`: column added with name `scalable_target_arn` and type `utf8`
- Table `aws_appsync_graphql_apis`: column added with name `api_type` and type `utf8`
- Table `aws_appsync_graphql_apis`: column added with name `dns` and type `json`
- Table `aws_appsync_graphql_apis`: column added with name `merged_api_execution_role_arn` and type `utf8`
- Table `aws_appsync_graphql_apis`: column added with name `owner_contact` and type `utf8`
- Table `aws_appsync_graphql_apis`: column added with name `owner` and type `utf8`
- Table `aws_appsync_graphql_apis`: column added with name `visibility` and type `utf8`
- Table `aws_athena_work_group_query_executions`: column added with name `substatement_type` and type `utf8`
- Table `aws_backup_protected_resources` was added
- Table `aws_backup_vault_recovery_points`: column added with name `resource_name` and type `utf8`
- Table `aws_cloudformation_stack_sets`: column added with name `regions` and type `list<item: utf8, nullable>`
- Table `aws_cloudwatchlogs_log_groups`: column added with name `inherited_properties` and type `list<item: utf8, nullable>`
- Table `aws_computeoptimizer_ebs_volume_recommendations`: column added with name `tags` and type `json`
- Table `aws_computeoptimizer_ec2_instance_recommendations`: column added with name `external_metric_status` and type `json`
- Table `aws_computeoptimizer_ec2_instance_recommendations`: column added with name `instance_state` and type `utf8`
- Table `aws_computeoptimizer_ec2_instance_recommendations`: column added with name `tags` and type `json`
- Table `aws_computeoptimizer_ecs_service_recommendations`: column added with name `tags` and type `json`
- Table `aws_computeoptimizer_lambda_function_recommendations`: column added with name `tags` and type `json`
- Table `aws_dynamodb_tables`: column added with name `deletion_protection_enabled` and type `bool`
- Table `aws_ec2_hosts`: column added with name `host_maintenance` and type `utf8`
- Table `aws_ec2_instance_connect_endpoints` was added
- Table `aws_ec2_instances`: column added with name `current_instance_boot_mode` and type `utf8`
- Table `aws_elasticache_replication_groups`: column added with name `cluster_mode` and type `utf8`
- Table `aws_emr_cluster_instance_fleets`: column added with name `resize_specifications` and type `json`
- Table `aws_frauddetector_event_types`: column added with name `event_orchestration` and type `json`
- Table `aws_glue_database_tables`: column added with name `federated_table` and type `json`
- Table `aws_glue_databases`: column added with name `federated_database` and type `json`
- Table `aws_guardduty_detectors`: column added with name `features` and type `json`
- Table `aws_iot_jobs`: column added with name `destination_package_versions` and type `list<item: utf8, nullable>`
- Table `aws_iot_jobs`: column added with name `scheduled_job_rollouts` and type `json`
- Table `aws_kafka_cluster_operations`: column added with name `vpc_connection_info` and type `json`
- Table `aws_lambda_function_event_source_mappings`: column added with name `document_db_event_source_config` and type `json`
- Table `aws_lightsail_disks`: column added with name `auto_mount_status` and type `utf8`
- Table `aws_mwaa_environments`: column added with name `startup_script_s3_object_version` and type `utf8`
- Table `aws_mwaa_environments`: column added with name `startup_script_s3_path` and type `utf8`
- Table `aws_neptune_clusters`: column added with name `global_cluster_identifier` and type `utf8`
- Table `aws_neptune_clusters`: column added with name `pending_modified_values` and type `json`
- Table `aws_networkfirewall_firewall_policies`: column added with name `policy_variables` and type `json`
- Table `aws_ram_resource_share_permissions`: column added with name `feature_set` and type `utf8`
- Table `aws_ram_resource_share_permissions`: column added with name `permission_type` and type `utf8`
- Table `aws_ram_resource_share_permissions`: column added with name `tags` and type `json`
- Table `aws_rds_cluster_snapshots`: column added with name `storage_type` and type `utf8`
- Table `aws_rds_clusters`: column added with name `io_optimized_next_allowed_modification_time` and type `timestamp[us, tz=UTC]`
- Table `aws_rds_instances`: column added with name `read_replica_source_db_cluster_identifier` and type `utf8`
- Table `aws_resiliencehub_app_version_resource_mappings`: column added with name `eks_source_name` and type `utf8`
- Table `aws_resiliencehub_app_version_resources`: column added with name `additional_info` and type `json`
- Table `aws_resiliencehub_app_version_resources`: column added with name `excluded` and type `bool`
- Table `aws_resiliencehub_app_version_resources`: column added with name `parent_resource_name` and type `utf8`
- Table `aws_resiliencehub_app_version_resources`: column added with name `source_type` and type `utf8`
- Table `aws_sagemaker_models`: column added with name `deployment_recommendation` and type `json`
- Table `aws_securityhub_hubs`: column added with name `control_finding_generator` and type `utf8`
- Table `aws_timestream_tables`: column added with name `schema` and type `json`
- Table `aws_wafv2_web_acls`: column added with name `association_config` and type `json`

### Features

* **aws:** Add `aws_alpha_cloudwatch_metric*` tables to fetch Cloudwatch Metrics and statistics ([#11402](https://github.com/cloudquery/cloudquery/issues/11402)) ([07b76d2](https://github.com/cloudquery/cloudquery/commit/07b76d218664595692d05bc1d0b6aa1a9b288766))
* **aws:** Add Support for EC2 Instance Connect Endpoints ([#11531](https://github.com/cloudquery/cloudquery/issues/11531)) ([73ae77b](https://github.com/cloudquery/cloudquery/commit/73ae77b755b37de21687937e5b7f75917355a3d3))
* **aws:** Add Support for fully customizable Cost Explorer Data syncing ([#11185](https://github.com/cloudquery/cloudquery/issues/11185)) ([04fd769](https://github.com/cloudquery/cloudquery/commit/04fd769b5b194b0020c0cf36f39f41c62216bd88)), closes [#10805](https://github.com/cloudquery/cloudquery/issues/10805)
* **aws:** Add support for Protected Resources in AWS Backup ([#11522](https://github.com/cloudquery/cloudquery/issues/11522)) ([12a1bc0](https://github.com/cloudquery/cloudquery/commit/12a1bc09605aa019a6e4196ae94bcc11abfcbc5b))
* **aws:** Table Options Validation ([#11548](https://github.com/cloudquery/cloudquery/issues/11548)) ([ea858fc](https://github.com/cloudquery/cloudquery/commit/ea858fc26f13c564bcae3a409d36c2c6ebb8f009))


### Bug Fixes

* **aws:** Upgrade all AWS SDK to latest version ([#11626](https://github.com/cloudquery/cloudquery/issues/11626)) ([71f787f](https://github.com/cloudquery/cloudquery/commit/71f787f81c800cfd171a924d1158fe3100479c35))
* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to 4d76231 ([#11532](https://github.com/cloudquery/cloudquery/issues/11532)) ([6f04233](https://github.com/cloudquery/cloudquery/commit/6f042333acbd2506f7800ccb89a8c5cbfb7ad8d4))
* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to d864719 ([#11611](https://github.com/cloudquery/cloudquery/issues/11611)) ([557a290](https://github.com/cloudquery/cloudquery/commit/557a2903af272b8e2e4c9eebb36e39cd8a41a805))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v3 to v3.10.6 ([#11473](https://github.com/cloudquery/cloudquery/issues/11473)) ([7272133](https://github.com/cloudquery/cloudquery/commit/72721336632e127dd37de4541f2f503bf4f73fb6))

## [18.3.0](https://github.com/cloudquery/cloudquery/compare/plugins-source-aws-v18.2.0...plugins-source-aws-v18.3.0) (2023-06-13)


### Features

* **aws-services:** Support newly added regions ([#11451](https://github.com/cloudquery/cloudquery/issues/11451)) ([5adf992](https://github.com/cloudquery/cloudquery/commit/5adf992c115c37cd7a4f4cf57789622acf993893))


### Bug Fixes

* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to d8eacf8 ([#11449](https://github.com/cloudquery/cloudquery/issues/11449)) ([742dafd](https://github.com/cloudquery/cloudquery/commit/742dafd5bf5cdc8facb94fda5de1d84c88897cbd))

## [18.2.0](https://github.com/cloudquery/cloudquery/compare/plugins-source-aws-v18.1.0...plugins-source-aws-v18.2.0) (2023-06-12)


### This Release has the Following Changes to Tables
- Table `aws_ec2_image_last_launched_times` was added

### Features

* **aws-services:** Support newly added regions ([#11446](https://github.com/cloudquery/cloudquery/issues/11446)) ([ab16ec5](https://github.com/cloudquery/cloudquery/commit/ab16ec5009958c0654b049318aee9822ce504b2e))
* **aws:** Add support for EC2 Image Last Launched Time ([#11224](https://github.com/cloudquery/cloudquery/issues/11224)) ([eaee4df](https://github.com/cloudquery/cloudquery/commit/eaee4dfb2c72441a7702eb648741f87f0e63d225))


### Bug Fixes

* **aws:** Resolve All types of attached Resources in `aws_wafv2_web_acls` ([#11420](https://github.com/cloudquery/cloudquery/issues/11420)) ([799211f](https://github.com/cloudquery/cloudquery/commit/799211f17ac6dd6b01c39d7a002714c1f6af7e0e))
* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to 0f7bd3b ([#11412](https://github.com/cloudquery/cloudquery/issues/11412)) ([dd1e2e8](https://github.com/cloudquery/cloudquery/commit/dd1e2e892d95515fd7332339262abaefd2a256c5))
* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to 7f6aaff ([#11432](https://github.com/cloudquery/cloudquery/issues/11432)) ([55dfebc](https://github.com/cloudquery/cloudquery/commit/55dfebc064608fb47caaf3b8e68c8002de8a7dc3))
* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to 8f72077 ([#11395](https://github.com/cloudquery/cloudquery/issues/11395)) ([d91fc5c](https://github.com/cloudquery/cloudquery/commit/d91fc5ce24f64c29fff6988b19ec2c2775cc379b))
* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to 90670b8 ([#11279](https://github.com/cloudquery/cloudquery/issues/11279)) ([a6cdc91](https://github.com/cloudquery/cloudquery/commit/a6cdc912e4b38a3faf798c5147a986ffe2539643))
* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to b359e74 ([#11405](https://github.com/cloudquery/cloudquery/issues/11405)) ([5d92765](https://github.com/cloudquery/cloudquery/commit/5d927659bd4f7c445a0e312487f1655ffb9a60f6))
* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to e258cfb ([#11391](https://github.com/cloudquery/cloudquery/issues/11391)) ([eacbe9a](https://github.com/cloudquery/cloudquery/commit/eacbe9ad3ea16d88f27c4593fa2774574ac8fe4e))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/acmpca to v1.21.13 ([#11423](https://github.com/cloudquery/cloudquery/issues/11423)) ([86b3afe](https://github.com/cloudquery/cloudquery/commit/86b3afe5af3e997d0526a4be64465d1c2ed14b1b))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/directconnect to v1.18.15 ([#11424](https://github.com/cloudquery/cloudquery/issues/11424)) ([44c33ee](https://github.com/cloudquery/cloudquery/commit/44c33ee861a896553c45cc052a139b2566c98af3))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2 to v1.19.11 ([#11425](https://github.com/cloudquery/cloudquery/issues/11425)) ([7089538](https://github.com/cloudquery/cloudquery/commit/7089538299881c1673a2c2e8a63101917246bbb7))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/elastictranscoder to v1.14.10 ([#11426](https://github.com/cloudquery/cloudquery/issues/11426)) ([f87fe41](https://github.com/cloudquery/cloudquery/commit/f87fe414ffe46d7d59b01a128f987c12c3410017))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/firehose to v1.16.12 ([#11427](https://github.com/cloudquery/cloudquery/issues/11427)) ([f9a494b](https://github.com/cloudquery/cloudquery/commit/f9a494bb0b128695a4ee9f0695551fe7c4971fe1))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/fsx to v1.28.13 ([#11428](https://github.com/cloudquery/cloudquery/issues/11428)) ([043aa2b](https://github.com/cloudquery/cloudquery/commit/043aa2bbf628b3b6f8c002d3dd930858b5c09e22))

## [18.1.0](https://github.com/cloudquery/cloudquery/compare/plugins-source-aws-v18.0.0...plugins-source-aws-v18.1.0) (2023-06-06)


### Features

* **aws-services:** Support newly added regions ([#11100](https://github.com/cloudquery/cloudquery/issues/11100)) ([3b79f9f](https://github.com/cloudquery/cloudquery/commit/3b79f9fc2958c687c4c238887bdd6d4d0c5c3078))
* **aws-services:** Support newly added regions ([#11227](https://github.com/cloudquery/cloudquery/issues/11227)) ([c65ed23](https://github.com/cloudquery/cloudquery/commit/c65ed23ed7fc137a7c6082e8f303e15b41648627))


### Bug Fixes

* **deps:** Update github.com/apache/arrow/go/v13 digest to e07e22c ([#11151](https://github.com/cloudquery/cloudquery/issues/11151)) ([5083cf7](https://github.com/cloudquery/cloudquery/commit/5083cf720f0ae98e07448ba2ae1116048e2d3a90))
* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to 20b0de9 ([#11199](https://github.com/cloudquery/cloudquery/issues/11199)) ([dc3565d](https://github.com/cloudquery/cloudquery/commit/dc3565d3fd6a640d9d10b4fd3a7fe6009a9d02a5))
* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to 88d5dc2 ([#11226](https://github.com/cloudquery/cloudquery/issues/11226)) ([9f306bc](https://github.com/cloudquery/cloudquery/commit/9f306bcaf3833b4611f0df5c50277be43aa19cbb))
* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to a7aad4c ([#11184](https://github.com/cloudquery/cloudquery/issues/11184)) ([8a0822e](https://github.com/cloudquery/cloudquery/commit/8a0822e31fc0eef99de2cdd2bd6d7e4c8b4131bf))
* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to c67fb39 ([#11169](https://github.com/cloudquery/cloudquery/issues/11169)) ([dcb0f92](https://github.com/cloudquery/cloudquery/commit/dcb0f9296a770a5cc2eb6bffd6b1ee30fbccb5dc))
* **deps:** Update github.com/gocarina/gocsv digest to 9ddd7fd ([#10467](https://github.com/cloudquery/cloudquery/issues/10467)) ([43f9525](https://github.com/cloudquery/cloudquery/commit/43f9525b176c8e204428b51ac1c05fe25728c4fb))
* **deps:** Update golang.org/x/exp digest to 2e198f4 ([#11155](https://github.com/cloudquery/cloudquery/issues/11155)) ([c46c62b](https://github.com/cloudquery/cloudquery/commit/c46c62b68692f527485d7f4b84265abc5dc1142c))
* **deps:** Update google.golang.org/genproto digest to e85fd2c ([#11156](https://github.com/cloudquery/cloudquery/issues/11156)) ([dbe7e92](https://github.com/cloudquery/cloudquery/commit/dbe7e9293d693a6821570e0e0b80202a936b6d3c))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/dax to v1.12.10 ([#11071](https://github.com/cloudquery/cloudquery/issues/11071)) ([deb9c9c](https://github.com/cloudquery/cloudquery/commit/deb9c9c14e70378bb3fbfa4299771e785083d501))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/directconnect to v1.18.14 ([#11206](https://github.com/cloudquery/cloudquery/issues/11206)) ([dbcbbcc](https://github.com/cloudquery/cloudquery/commit/dbcbbcc6551ab45cec48637000c6a6bfe0812dda))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/dynamodbstreams to v1.14.11 ([#11207](https://github.com/cloudquery/cloudquery/issues/11207)) ([20be49f](https://github.com/cloudquery/cloudquery/commit/20be49f0175ee2225869e26db27322581594de0b))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/ecr to v1.18.11 ([#11208](https://github.com/cloudquery/cloudquery/issues/11208)) ([155e8fe](https://github.com/cloudquery/cloudquery/commit/155e8fe2decf84db11aa227027ef2aa22657826c))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/eks to v1.27.12 ([#11209](https://github.com/cloudquery/cloudquery/issues/11209)) ([a75d0ff](https://github.com/cloudquery/cloudquery/commit/a75d0ffda4c338510037f84252ec75ead2fcf58d))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/elasticbeanstalk to v1.15.10 ([#11210](https://github.com/cloudquery/cloudquery/issues/11210)) ([9508730](https://github.com/cloudquery/cloudquery/commit/9508730ff6438e4f9b75299d5f6d9959afecea85))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/elasticloadbalancing to v1.15.10 ([#11211](https://github.com/cloudquery/cloudquery/issues/11211)) ([ea9973d](https://github.com/cloudquery/cloudquery/commit/ea9973d981a734abf34c731ac94ffb35a61be02a))
* **deps:** Update module github.com/cloudquery/plugin-pb-go to v1.0.9 ([#11240](https://github.com/cloudquery/cloudquery/issues/11240)) ([f92cd4b](https://github.com/cloudquery/cloudquery/commit/f92cd4bfe3c3d0088964d52ab9cd01ca4cf622e1))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v3 to v3.10.3 ([#11150](https://github.com/cloudquery/cloudquery/issues/11150)) ([dc00994](https://github.com/cloudquery/cloudquery/commit/dc00994e32936af7e9893c93561d0f9df225a929))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v3 to v3.10.4 ([#11244](https://github.com/cloudquery/cloudquery/issues/11244)) ([8fceef6](https://github.com/cloudquery/cloudquery/commit/8fceef6f9041e173923555d8ff221cfe83b424c2))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v3 to v3.7.0 ([#11113](https://github.com/cloudquery/cloudquery/issues/11113)) ([487bf87](https://github.com/cloudquery/cloudquery/commit/487bf871afe360cb8d9d592dfea48837d6e7cf27))

## [18.0.0](https://github.com/cloudquery/cloudquery/compare/plugins-source-aws-v17.4.0...plugins-source-aws-v18.0.0) (2023-05-29)


### This Release has the Following Changes to Tables
- Table `aws_apigateway_rest_api_authorizers`: column `provider_ar_ns` removed from table (:warning: breaking)
- Table `aws_apigateway_rest_api_authorizers`: column added with name `provider_arns` and type `list<item: utf8, nullable>`
- Table `aws_autoscaling_groups`: column `target_group_ar_ns` removed from table (:warning: breaking)
- Table `aws_autoscaling_groups`: column added with name `target_group_arns` and type `list<item: utf8, nullable>`
- Table `aws_cloudformation_stacks`: column `notification_ar_ns` removed from table (:warning: breaking)
- Table `aws_cloudformation_stacks`: column added with name `notification_arns` and type `list<item: utf8, nullable>`
- Table `aws_cognito_identity_pools`: column `open_id_connect_provider_ar_ns` removed from table (:warning: breaking)
- Table `aws_cognito_identity_pools`: column `saml_provider_ar_ns` removed from table (:warning: breaking)
- Table `aws_cognito_identity_pools`: column added with name `open_id_connect_provider_arns` and type `list<item: utf8, nullable>`
- Table `aws_cognito_identity_pools`: column added with name `saml_provider_arns` and type `list<item: utf8, nullable>`
- Table `aws_ssoadmin_permission_sets`: column added with name `request_account_id` and type `utf8`
- Table `aws_ssoadmin_permission_sets`: column added with name `request_region` and type `utf8`

### ⚠ BREAKING CHANGES

* **aws:** Change names of columns which had `_ar_ns` instead of `_arns` ([#10802](https://github.com/cloudquery/cloudquery/issues/10802))
* This release introduces an internal change to our type system to use [Apache Arrow](https://arrow.apache.org/). This should not have any visible breaking changes, however due to the size of the change we are introducing it under a major version bump to communicate that it might have some bugs that we weren't able to catch during our internal tests. If you encounter an issue during the upgrade, please submit a [bug report](https://github.com/cloudquery/cloudquery/issues/new/choose). You will also need to update destinations depending on which one you use:
    - Azure Blob Storage >= v3.2.0
    - BigQuery >= v3.0.0
    - ClickHouse >= v3.1.1
    - DuckDB >= v1.1.6
    - Elasticsearch >= v2.0.0
    - File >= v3.2.0
    - Firehose >= v2.0.2
    - GCS >= v3.2.0
    - Gremlin >= v2.1.10
    - Kafka >= v3.0.1
    - Meilisearch >= v2.0.1
    - Microsoft SQL Server >= v4.2.0
    - MongoDB >= v2.0.1
    - MySQL >= v2.0.2
    - Neo4j >= v3.0.0
    - PostgreSQL >= v4.2.0
    - S3 >= v4.4.0
    - Snowflake >= v2.1.1
    - SQLite >= v2.2.0

### Features

* Update to use [Apache Arrow](https://arrow.apache.org/) type system ([#10797](https://github.com/cloudquery/cloudquery/issues/10797)) ([e355d14](https://github.com/cloudquery/cloudquery/commit/e355d14b8dac61226e1fecd53cf9a84fc79b4640))


### Bug Fixes

* **aws:** Change names of columns which had `_ar_ns` instead of `_arns` ([#10802](https://github.com/cloudquery/cloudquery/issues/10802)) ([e00ac44](https://github.com/cloudquery/cloudquery/commit/e00ac44bfb779c0c787409383de07698a4c37a8f))
* **aws:** Remove Hardcoded fix for AWS issue ([#10972](https://github.com/cloudquery/cloudquery/issues/10972)) ([ede53a7](https://github.com/cloudquery/cloudquery/commit/ede53a7eb262bdb6f713f538acc7ebdcddb80087))
* **aws:** Support syncing AWS SSO Account Assignments for non management accounts ([#10881](https://github.com/cloudquery/cloudquery/issues/10881)) ([a715e4f](https://github.com/cloudquery/cloudquery/commit/a715e4f24f973d8382f9d901723baf6ee6116d18))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs to v1.20.11 ([#11066](https://github.com/cloudquery/cloudquery/issues/11066)) ([be8e23b](https://github.com/cloudquery/cloudquery/commit/be8e23bc22de28a14b7f203ffc2823a58ed0cadf))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/codebuild to v1.20.13 ([#11067](https://github.com/cloudquery/cloudquery/issues/11067)) ([c3c831a](https://github.com/cloudquery/cloudquery/commit/c3c831a2983948af8bb5769ed550aa7ce287e121))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/cognitoidentity to v1.15.11 ([#11068](https://github.com/cloudquery/cloudquery/issues/11068)) ([6708fec](https://github.com/cloudquery/cloudquery/commit/6708fecfa7a7838d60e6287fdbcf44abfdbe564c))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider to v1.22.10 ([#11069](https://github.com/cloudquery/cloudquery/issues/11069)) ([28698dc](https://github.com/cloudquery/cloudquery/commit/28698dc02a9deff8df5d0fb225ce8517ae4bd1d8))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/costexplorer to v1.25.10 ([#11070](https://github.com/cloudquery/cloudquery/issues/11070)) ([8da3107](https://github.com/cloudquery/cloudquery/commit/8da31078b1e0e71bee8e4aef8fab9c141b69f0f5))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.45.0 ([#11041](https://github.com/cloudquery/cloudquery/issues/11041)) ([035e461](https://github.com/cloudquery/cloudquery/commit/035e461ae26ec14ea918a3e8b918802b1ff770e4))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v3 to v3.6.7 ([#11043](https://github.com/cloudquery/cloudquery/issues/11043)) ([3c6d885](https://github.com/cloudquery/cloudquery/commit/3c6d885c3d201b0b39cbc1406c6e54a57ec5ed5f))

## [17.4.0](https://github.com/cloudquery/cloudquery/compare/plugins-source-aws-v17.3.1...plugins-source-aws-v17.4.0) (2023-05-23)


### Features

* **aws-services:** Support newly added regions ([#10894](https://github.com/cloudquery/cloudquery/issues/10894)) ([a9b6633](https://github.com/cloudquery/cloudquery/commit/a9b6633fc0ed56bf9cba11699e3f9b62e3033067))
* **aws:** Only instantiate a single SDK per Account ([#10794](https://github.com/cloudquery/cloudquery/issues/10794)) ([1a3a4d3](https://github.com/cloudquery/cloudquery/commit/1a3a4d3eb51bcbae65a4e24d61a2f89c905a9a56))


### Bug Fixes

* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/apigateway to v1.16.11 ([#10875](https://github.com/cloudquery/cloudquery/issues/10875)) ([2a3f966](https://github.com/cloudquery/cloudquery/commit/2a3f96640221383d31da8cecf7e8b8fccb96deae))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/apigatewayv2 to v1.13.11 ([#10876](https://github.com/cloudquery/cloudquery/issues/10876)) ([a9087cd](https://github.com/cloudquery/cloudquery/commit/a9087cde6b2507ad69a222e113e8d2d11997535a))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/autoscalingplans to v1.13.10 ([#10878](https://github.com/cloudquery/cloudquery/issues/10878)) ([ad70bbf](https://github.com/cloudquery/cloudquery/commit/ad70bbffa642ace31ffc963ac8ea9749d21803d2))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/cloudhsmv2 to v1.14.10 ([#10879](https://github.com/cloudquery/cloudquery/issues/10879)) ([9d5887d](https://github.com/cloudquery/cloudquery/commit/9d5887df32cea8feb82a6345b8c1caeacb515cff))

## [17.3.1](https://github.com/cloudquery/cloudquery/compare/plugins-source-aws-v17.3.0...plugins-source-aws-v17.3.1) (2023-05-18)


### This Release has the Following Changes to Tables
- Table `aws_cloudformation_stack_templates`: column added with name `template_body_text` and type `String`

### Bug Fixes

* **aws-s3:** Always use models.WrappedBucket to avoid panic ([#10827](https://github.com/cloudquery/cloudquery/issues/10827)) ([be55852](https://github.com/cloudquery/cloudquery/commit/be558527b244670cf07349f3128768d06371b2b3))
* **aws:** Handle YAML cloudformation template bodies ([#10826](https://github.com/cloudquery/cloudquery/issues/10826)) ([498f394](https://github.com/cloudquery/cloudquery/commit/498f39479d9810a6b87a9c552ac915981b050fd9))

## [17.3.0](https://github.com/cloudquery/cloudquery/compare/plugins-source-aws-v17.2.0...plugins-source-aws-v17.3.0) (2023-05-16)


### This Release has the Following Changes to Tables
- Table `aws_acmpca_certificate_authorities` was added
- Table `aws_cloudformation_stack_templates` was added
- Table `aws_cloudformation_template_summaries` was added
- Table `aws_config_config_rule_compliance_details` was added
- Table `aws_efs_access_points` was added

### Features

* Add AWS Config Compliance Details table ([#10544](https://github.com/cloudquery/cloudquery/issues/10544)) ([9b43a2a](https://github.com/cloudquery/cloudquery/commit/9b43a2af489389a4cdd340e7880e899c149566aa))
* Add Cloudformation Template Summaries table ([#10571](https://github.com/cloudquery/cloudquery/issues/10571)) ([3515db9](https://github.com/cloudquery/cloudquery/commit/3515db985572016dff8388efd463b7e7260a54e4))
* **aws-services:** Support newly added regions ([#10806](https://github.com/cloudquery/cloudquery/issues/10806)) ([52b5e0f](https://github.com/cloudquery/cloudquery/commit/52b5e0f81e1b093fa82e2b409c448f4e29efdfaf))
* **aws:** Add private certificate authorities ([#10691](https://github.com/cloudquery/cloudquery/issues/10691)) ([9c97e85](https://github.com/cloudquery/cloudquery/commit/9c97e858ed4a0287b5cdc4ad1c7ca0b3c26255eb))
* **aws:** Add Support for Cloudformation Templates ([#10701](https://github.com/cloudquery/cloudquery/issues/10701)) ([7a23c2e](https://github.com/cloudquery/cloudquery/commit/7a23c2e59ebfb37349cbc0be8d17ae76e071d12f))
* **aws:** Add support for EFS Access Point ([#10803](https://github.com/cloudquery/cloudquery/issues/10803)) ([d994c85](https://github.com/cloudquery/cloudquery/commit/d994c851a3dc1ddaa3fbe571d7a25f87649a087b))
* **aws:** Support Table level inputs ([#10564](https://github.com/cloudquery/cloudquery/issues/10564)) ([161b11b](https://github.com/cloudquery/cloudquery/commit/161b11b1889a1283bfcfd7c5edfc108544f2ded3))
* **deps:** Upgrade to Apache Arrow v13 (latest `cqmain`) ([#10605](https://github.com/cloudquery/cloudquery/issues/10605)) ([a55da3d](https://github.com/cloudquery/cloudquery/commit/a55da3dbefafdc68a6bda2d5f1d334d12dd97b97))


### Bug Fixes

* **aws-policies:** Api Gateway stage logging for REST ([#10625](https://github.com/cloudquery/cloudquery/issues/10625)) ([f0d6f57](https://github.com/cloudquery/cloudquery/commit/f0d6f57b4b8876984edc327cefa68cf7f063c941))
* **aws-policies:** Api Gateway stage logging for websockets ([#10702](https://github.com/cloudquery/cloudquery/issues/10702)) ([e667400](https://github.com/cloudquery/cloudquery/commit/e6674006fe326859226cf6e01f754485bb40b4ad))
* **aws:** Change column type of `aws_cloudformation_stack_templates.template_body` ([#10752](https://github.com/cloudquery/cloudquery/issues/10752)) ([75b9785](https://github.com/cloudquery/cloudquery/commit/75b97858739d7a6e28f43d4dc9ecc2610b9ca062))
* **aws:** Fix the case where resrouce_id is null in ECS.2 of foundational policy ([#10692](https://github.com/cloudquery/cloudquery/issues/10692)) ([f5cf2d8](https://github.com/cloudquery/cloudquery/commit/f5cf2d866ab565e11be20f6c4e9f2bf0b8766eb8))
* **aws:** Handle Cloudfront Regions in different partitions ([#10690](https://github.com/cloudquery/cloudquery/issues/10690)) ([158aab1](https://github.com/cloudquery/cloudquery/commit/158aab10de740ef18564e7288893477be96dbbbe))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/config to v1.18.25 ([#10786](https://github.com/cloudquery/cloudquery/issues/10786)) ([caca1a4](https://github.com/cloudquery/cloudquery/commit/caca1a41e298c06afb6f474b8fd911c4544a2eec))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/acm to v1.17.11 ([#10789](https://github.com/cloudquery/cloudquery/issues/10789)) ([9122f84](https://github.com/cloudquery/cloudquery/commit/9122f843867d7099e3e872e1236b7595d5f02b04))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/amp to v1.16.11 ([#10790](https://github.com/cloudquery/cloudquery/issues/10790)) ([431905f](https://github.com/cloudquery/cloudquery/commit/431905f2b4eb34a664e68d34ae7689293f1586e4))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/amplify to v1.13.10 ([#10791](https://github.com/cloudquery/cloudquery/issues/10791)) ([81d175b](https://github.com/cloudquery/cloudquery/commit/81d175b96e5eb46a9009d85abaf3e28b9eac8c60))
* **deps:** Update module github.com/cloudquery/plugin-pb-go to v1.0.8 ([#10798](https://github.com/cloudquery/cloudquery/issues/10798)) ([27ff430](https://github.com/cloudquery/cloudquery/commit/27ff430527932d59a4d488a6767547eda8853940))

## [17.2.0](https://github.com/cloudquery/cloudquery/compare/plugins-source-aws-v17.1.0...plugins-source-aws-v17.2.0) (2023-05-09)


### This Release has the Following Changes to Tables
- Table `aws_config_configuration_aggregators` was added
- Table `aws_config_remediation_configurations` was added
- Table `aws_config_retention_configurations` was added
- Table `aws_networkfirewall_firewall_policies` was added
- Table `aws_networkfirewall_rule_groups` was added
- Table `aws_securityhub_enabled_standards` was added
- Table `aws_securityhub_hubs` was added
- Table `aws_wafregional_rule_groups`: column added with name `rule_ids` and type `StringArray`

### Features

* **aws-services:** Support newly added regions ([#10598](https://github.com/cloudquery/cloudquery/issues/10598)) ([e56bae2](https://github.com/cloudquery/cloudquery/commit/e56bae2a0d344b3ccccedd403932395da06c61fe))
* **aws:** Add `aws_securityhub_hubs` and `aws_securityhub_enabled_standards` tables ([#10553](https://github.com/cloudquery/cloudquery/issues/10553)) ([bc77f53](https://github.com/cloudquery/cloudquery/commit/bc77f5315331071e765d858829814086534d4750)), closes [#1592](https://github.com/cloudquery/cloudquery/issues/1592)
* **aws:** Add networkfirewall resources ([#10547](https://github.com/cloudquery/cloudquery/issues/10547)) ([74ada09](https://github.com/cloudquery/cloudquery/commit/74ada094e737ab2892b884bdad0332489968d3e4))
* **aws:** Add rule_ids to wafregional rule_group ([#10594](https://github.com/cloudquery/cloudquery/issues/10594)) ([4637baa](https://github.com/cloudquery/cloudquery/commit/4637baaed1d725a26bcbdb923c5bc9280446b8dd))
* **aws:** Refine AWS Org Error Message ([#10569](https://github.com/cloudquery/cloudquery/issues/10569)) ([b761ebc](https://github.com/cloudquery/cloudquery/commit/b761ebc2e91ee20492439e60b65322eb400436a3))
* More AWS Config resources ([#10509](https://github.com/cloudquery/cloudquery/issues/10509)) ([3db4ebb](https://github.com/cloudquery/cloudquery/commit/3db4ebb9394cda0279887a507aaee256261910f5))


### Bug Fixes

* **aws-policies:** Api Gateway xray tracing enabled query ([#10597](https://github.com/cloudquery/cloudquery/issues/10597)) ([341d849](https://github.com/cloudquery/cloudquery/commit/341d84931ac4fd698505e6894fe69111d8a7ebb2))
* **aws:** Use GetBucketLocation for S3 bucket region ([#10550](https://github.com/cloudquery/cloudquery/issues/10550)) ([7f0128f](https://github.com/cloudquery/cloudquery/commit/7f0128f2ba1af1cb88cfa4de93cfee148959c488)), closes [#10548](https://github.com/cloudquery/cloudquery/issues/10548)
* **deps:** Update module github.com/aws/aws-sdk-go-v2/config to v1.18.23 ([#10576](https://github.com/cloudquery/cloudquery/issues/10576)) ([eeb13d5](https://github.com/cloudquery/cloudquery/commit/eeb13d5b1b6b6fcb32764c8711bfbb79da35f9a8))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/internal/v4a to v1.0.25 ([#10585](https://github.com/cloudquery/cloudquery/issues/10585)) ([00d7449](https://github.com/cloudquery/cloudquery/commit/00d744988953f8a3f9c177970645be7b0f34fc69))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/accessanalyzer to v1.19.12 ([#10586](https://github.com/cloudquery/cloudquery/issues/10586)) ([4f929fe](https://github.com/cloudquery/cloudquery/commit/4f929fe811675ca9c520cf4cba11ed6debee7239))
* Remove unused line in mockassert library ([#10608](https://github.com/cloudquery/cloudquery/issues/10608)) ([3fc8708](https://github.com/cloudquery/cloudquery/commit/3fc8708ecc66b1642e6963aa1094392d5b2533f2))

## [17.1.0](https://github.com/cloudquery/cloudquery/compare/plugins-source-aws-v17.0.0...plugins-source-aws-v17.1.0) (2023-05-02)


### Features

* **aws-services:** Support newly added regions ([#10480](https://github.com/cloudquery/cloudquery/issues/10480)) ([f0b7aad](https://github.com/cloudquery/cloudquery/commit/f0b7aaddcafdbc5d6d6cc7525cfe8fb34bc1d10a))
* **aws:** Add Support for Cost Explorer data ([#10288](https://github.com/cloudquery/cloudquery/issues/10288)) ([c22f9b7](https://github.com/cloudquery/cloudquery/commit/c22f9b7677fa16030d08463bf159e06d385ea39c))
* **aws:** Make Cost And Forecast tables fixed time intervals ([#10479](https://github.com/cloudquery/cloudquery/issues/10479)) ([0560cb3](https://github.com/cloudquery/cloudquery/commit/0560cb3edc9c0abec45fc7416b1707c30741acc8))


### Bug Fixes

* **aws:** Cost Explorer Date Range ([#10458](https://github.com/cloudquery/cloudquery/issues/10458)) ([1bc7fdd](https://github.com/cloudquery/cloudquery/commit/1bc7fdd85a0d7d444ab838e3953a67c669d301e7))
* **aws:** Explicitly set region in all SDK calls ([#10453](https://github.com/cloudquery/cloudquery/issues/10453)) ([32a52ba](https://github.com/cloudquery/cloudquery/commit/32a52ba0d3b1342dd2346d47c13787e62e9e191a))
* **deps:** Update github.com/apache/arrow/go/v12 digest to 0ea1a10 ([#10461](https://github.com/cloudquery/cloudquery/issues/10461)) ([022709f](https://github.com/cloudquery/cloudquery/commit/022709f710cc6d95aee60260d6f58991698bbf42))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v2 to v2.5.0 ([#10390](https://github.com/cloudquery/cloudquery/issues/10390)) ([f706688](https://github.com/cloudquery/cloudquery/commit/f706688b2f5b8393d09d57020d31fb1d280f0dbd))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v2 to v2.5.1 ([#10448](https://github.com/cloudquery/cloudquery/issues/10448)) ([cc85b93](https://github.com/cloudquery/cloudquery/commit/cc85b939fe945939caf72f8c08095e1e744b9ee8))

## [17.0.0](https://github.com/cloudquery/cloudquery/compare/plugins-source-aws-v16.3.0...plugins-source-aws-v17.0.0) (2023-04-25)


### ⚠ BREAKING CHANGES

* **aws:** Move EC2 EBS Snapshot `attributes` column to standalone table ([#10247](https://github.com/cloudquery/cloudquery/issues/10247))
* **aws-resources:** Query issue on `aws_cloudwatchlogs_log_group_data_protection_policies` ([#10243](https://github.com/cloudquery/cloudquery/issues/10243))
* **aws:** Fix Primary Key for `aws_directconnect_gateways` ([#10233](https://github.com/cloudquery/cloudquery/issues/10233))

### Features

* **aws-services:** Support newly added regions ([#10289](https://github.com/cloudquery/cloudquery/issues/10289)) ([dd6c034](https://github.com/cloudquery/cloudquery/commit/dd6c034fa99002fc4cde912d240c1f386d9171f0))
* **aws:** Store Spec in Client ([#10286](https://github.com/cloudquery/cloudquery/issues/10286)) ([2561d5c](https://github.com/cloudquery/cloudquery/commit/2561d5c05159c19fe8f860a2f060d30b78c2f15b))


### Bug Fixes

* **aws-docs:** Policy docs should include tables used in views ([#10250](https://github.com/cloudquery/cloudquery/issues/10250)) ([08c4d91](https://github.com/cloudquery/cloudquery/commit/08c4d91b86e65a45c7f08d9c55bcbcfc92753d3a))
* **aws-resources:** Query issue on `aws_cloudwatchlogs_log_group_data_protection_policies` ([#10243](https://github.com/cloudquery/cloudquery/issues/10243)) ([4e9bb39](https://github.com/cloudquery/cloudquery/commit/4e9bb39b2540e1878d88f43c4ac69154ae74b353)), closes [#10216](https://github.com/cloudquery/cloudquery/issues/10216)
* **aws:** Fix Primary Key for `aws_directconnect_gateways` ([#10233](https://github.com/cloudquery/cloudquery/issues/10233)) ([fc9094f](https://github.com/cloudquery/cloudquery/commit/fc9094f57f2b2e67a040c68725f2c6271f3da2f9))
* **aws:** Move EC2 EBS Snapshot `attributes` column to standalone table ([#10247](https://github.com/cloudquery/cloudquery/issues/10247)) ([bdb421b](https://github.com/cloudquery/cloudquery/commit/bdb421bd7a23254c9cacf1be13ea19bba52a9d23))
* **aws:** Remove Resource Specific Parallelization For S3 to use only the SDK parallelization  ([#10255](https://github.com/cloudquery/cloudquery/issues/10255)) ([07c7edb](https://github.com/cloudquery/cloudquery/commit/07c7edb9eabcf7dfddf4db33e5e639746c9504d0))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/feature/s3/manager to v1.11.63 ([#10267](https://github.com/cloudquery/cloudquery/issues/10267)) ([7a8a4c7](https://github.com/cloudquery/cloudquery/commit/7a8a4c787bf2849b799014f51d32bec42942d16d))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/amp to v1.16.8 ([#10268](https://github.com/cloudquery/cloudquery/issues/10268)) ([d4ec528](https://github.com/cloudquery/cloudquery/commit/d4ec52819c876b2fd501e30aeaf24ff2c5f7a2f7))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/amplify to v1.13.8 ([#10269](https://github.com/cloudquery/cloudquery/issues/10269)) ([8d5d163](https://github.com/cloudquery/cloudquery/commit/8d5d163793725877545c525af03a52e06e7c6bd4))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/apigateway to v1.16.9 ([#10270](https://github.com/cloudquery/cloudquery/issues/10270)) ([dc8fd93](https://github.com/cloudquery/cloudquery/commit/dc8fd93d66ca1a7921f358dc7cd9606d2c2cae90))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/apigatewayv2 to v1.13.9 ([#10271](https://github.com/cloudquery/cloudquery/issues/10271)) ([e3e3f0a](https://github.com/cloudquery/cloudquery/commit/e3e3f0a759910a21b03b195651eb65039f7c2b0a))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/autoscalingplans to v1.13.8 ([#10272](https://github.com/cloudquery/cloudquery/issues/10272)) ([cd7bcee](https://github.com/cloudquery/cloudquery/commit/cd7bcee51920be5e19049c164f8fe6ed80b61413))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v2 to v2.3.8 ([#10213](https://github.com/cloudquery/cloudquery/issues/10213)) ([f358666](https://github.com/cloudquery/cloudquery/commit/f35866611cd206c37e6e9f9ad3329561e4cb32af))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v2 to v2.4.0 ([#10278](https://github.com/cloudquery/cloudquery/issues/10278)) ([a0a713e](https://github.com/cloudquery/cloudquery/commit/a0a713e8490b970b9d8bfaa1b50e01f43ff51c36))
* **policies-cloudtrail:** Add region criteria to inner join ([#10246](https://github.com/cloudquery/cloudquery/issues/10246)) ([c9fd369](https://github.com/cloudquery/cloudquery/commit/c9fd369827010e048f9715400d4907d0436de73a))

## [16.3.0](https://github.com/cloudquery/cloudquery/compare/plugins-source-aws-v16.2.0...plugins-source-aws-v16.3.0) (2023-04-20)


### Features

* **aws-services:** Support newly added regions ([#10152](https://github.com/cloudquery/cloudquery/issues/10152)) ([d395a94](https://github.com/cloudquery/cloudquery/commit/d395a94344b2406247f28c79564d18470eb754a7))
* **aws:** Add Support for Config Delivery Channels ([#10150](https://github.com/cloudquery/cloudquery/issues/10150)) ([361df3f](https://github.com/cloudquery/cloudquery/commit/361df3f03b005f2c2a8ccb3c4be86a5190788024))
* **aws:** Parallelize Initialization of Accounts after discovery ([#10177](https://github.com/cloudquery/cloudquery/issues/10177)) ([3838e80](https://github.com/cloudquery/cloudquery/commit/3838e8094a3c247cfee80772945dce54678f38f2))
* **aws:** Upgrade to `github.com/cloudquery/plugin-sdk/v2` ([#9938](https://github.com/cloudquery/cloudquery/issues/9938)) ([a3fb436](https://github.com/cloudquery/cloudquery/commit/a3fb4366d91be52418edbe526b640729d61467f7)), closes [#9937](https://github.com/cloudquery/cloudquery/issues/9937)


### Bug Fixes

* **aws:** Update EBS Snapshot Permissions Check Query ([#10149](https://github.com/cloudquery/cloudquery/issues/10149)) ([f65d9da](https://github.com/cloudquery/cloudquery/commit/f65d9da21cd57660c7ee8af76f720245400faf45)), closes [#10140](https://github.com/cloudquery/cloudquery/issues/10140)
* **deps:** Update module github.com/aws/aws-sdk-go-v2/config to v1.18.21 ([#10127](https://github.com/cloudquery/cloudquery/issues/10127)) ([3bcde69](https://github.com/cloudquery/cloudquery/commit/3bcde697c5f927fa4eab52ea4293f1f7724812d1))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/feature/s3/manager to v1.11.62 ([#10129](https://github.com/cloudquery/cloudquery/issues/10129)) ([13f8670](https://github.com/cloudquery/cloudquery/commit/13f867006cd17c92bc1b18022ab3a210266258d8))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/accessanalyzer to v1.19.10 ([#10131](https://github.com/cloudquery/cloudquery/issues/10131)) ([eefbad5](https://github.com/cloudquery/cloudquery/commit/eefbad516aff7599ab67dedb03d41b0dbb94b88c))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/acm to v1.17.9 ([#10132](https://github.com/cloudquery/cloudquery/issues/10132)) ([7f6d235](https://github.com/cloudquery/cloudquery/commit/7f6d235266ed3daaa8d6aed8feb8502cf2ac6773))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v2 to v2.2.0 ([#10135](https://github.com/cloudquery/cloudquery/issues/10135)) ([cf33b89](https://github.com/cloudquery/cloudquery/commit/cf33b892ead0bb231e3956aa70967de552a21624))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v2 to v2.2.2 ([#10143](https://github.com/cloudquery/cloudquery/issues/10143)) ([8f887e0](https://github.com/cloudquery/cloudquery/commit/8f887e05de2096e8efd1e55863a8cf3c7620ccc3))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v2 to v2.3.0 ([#10163](https://github.com/cloudquery/cloudquery/issues/10163)) ([9a7f214](https://github.com/cloudquery/cloudquery/commit/9a7f21460772200e7a588409ebc7eb19f97b195b))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v2 to v2.3.1 ([#10175](https://github.com/cloudquery/cloudquery/issues/10175)) ([5b53423](https://github.com/cloudquery/cloudquery/commit/5b53423e72672f6c2bfb8ae00cfce1641410443e))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v2 to v2.3.3 ([#10187](https://github.com/cloudquery/cloudquery/issues/10187)) ([b185248](https://github.com/cloudquery/cloudquery/commit/b1852480b6ec8b721d94c72d8435051352f26932))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v2 to v2.3.4 ([#10196](https://github.com/cloudquery/cloudquery/issues/10196)) ([c6d2f59](https://github.com/cloudquery/cloudquery/commit/c6d2f59c7d77177a351cb82ecdc381dec6aad30c))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v2 to v2.3.5 ([#10200](https://github.com/cloudquery/cloudquery/issues/10200)) ([5a33693](https://github.com/cloudquery/cloudquery/commit/5a33693fe29f7068b03d80be1859d6e479c42c0d))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v2 to v2.3.6 ([#10208](https://github.com/cloudquery/cloudquery/issues/10208)) ([91c80a7](https://github.com/cloudquery/cloudquery/commit/91c80a795b46480447cfaef67c4db721a31e3206))

## [16.2.0](https://github.com/cloudquery/cloudquery/compare/plugins-source-aws-v16.1.0...plugins-source-aws-v16.2.0) (2023-04-12)


### This Release has the Following Changes to Tables
- Table `aws_cloudwatchlogs_log_group_data_protection_policies` was added

### Features

* **aws:** Add Support for Cloudwatch Logs Data Protection Policy ([#9818](https://github.com/cloudquery/cloudquery/issues/9818)) ([9014726](https://github.com/cloudquery/cloudquery/commit/901472697dbd6370079659d0f08d8ef01865b987))


### Bug Fixes

* **aws:** Update module github.com/cloudquery/plugin-sdk to v1.45.0  ([#9856](https://github.com/cloudquery/cloudquery/issues/9856)) ([7aa0b5f](https://github.com/cloudquery/cloudquery/commit/7aa0b5fe0520a9905717e299f2702dd1f362ed5b))

## [16.1.0](https://github.com/cloudquery/cloudquery/compare/plugins-source-aws-v16.0.1...plugins-source-aws-v16.1.0) (2023-04-11)


### Features

* **aws-services:** Support newly added regions ([#9821](https://github.com/cloudquery/cloudquery/issues/9821)) ([4d7f388](https://github.com/cloudquery/cloudquery/commit/4d7f38890433ee97310794b11656c6386aab41c3))
* **aws:** Parallelize resolution of  `aws_iot_billing_groups` ([#9809](https://github.com/cloudquery/cloudquery/issues/9809)) ([e759661](https://github.com/cloudquery/cloudquery/commit/e7596615a12ab19d59bc21f79a6de91f4898d7be))
* **aws:** Parallelize resolution of `aws_iot_ca_certificates` ([#9808](https://github.com/cloudquery/cloudquery/issues/9808)) ([d491460](https://github.com/cloudquery/cloudquery/commit/d49146046a0efb743f37044edce2cd24da7376da))
* **aws:** Parallelize resolution of `aws_iot_jobs` ([#9810](https://github.com/cloudquery/cloudquery/issues/9810)) ([9b85678](https://github.com/cloudquery/cloudquery/commit/9b85678a17d0d7a8d1b4d44593bd4594e1cef658))
* **aws:** Parallelize resolution of `aws_iot_policies` ([#9806](https://github.com/cloudquery/cloudquery/issues/9806)) ([b9a9d06](https://github.com/cloudquery/cloudquery/commit/b9a9d06fa9ec4d303f27281ac733c226ec1ac9f7))
* **aws:** Parallelize resolution of `aws_iot_security_profiles` ([#9807](https://github.com/cloudquery/cloudquery/issues/9807)) ([45705f1](https://github.com/cloudquery/cloudquery/commit/45705f1910478c8f7e56070a6632299b897dc621))
* **aws:** Parallelize resolution of `aws_iot_streams` ([#9804](https://github.com/cloudquery/cloudquery/issues/9804)) ([9b622e0](https://github.com/cloudquery/cloudquery/commit/9b622e087c946849cd5aaa9a8d8d506e1020afa4))
* **aws:** Parallelize resolution of `aws_iot_thing_groups` ([#9805](https://github.com/cloudquery/cloudquery/issues/9805)) ([f5047b9](https://github.com/cloudquery/cloudquery/commit/f5047b928f948d37fac4052ba27b75ffe2ded254))
* **aws:** Parallelize resolution of `aws_iot_topic_rules` ([#9803](https://github.com/cloudquery/cloudquery/issues/9803)) ([d344815](https://github.com/cloudquery/cloudquery/commit/d34481566ff997527cbdc7720f76efa6795048c5))


### Bug Fixes

* **aws-policies:** Update Query to properly handle a string and array ([#9815](https://github.com/cloudquery/cloudquery/issues/9815)) ([012347f](https://github.com/cloudquery/cloudquery/commit/012347ff6c533e72f1f48f03fb973e46b31dabe9)), closes [#9763](https://github.com/cloudquery/cloudquery/issues/9763)
* **aws:** Inspector Classic fetch details use proper limits ([#9816](https://github.com/cloudquery/cloudquery/issues/9816)) ([225a796](https://github.com/cloudquery/cloudquery/commit/225a796a83832a003bb2b509f602130d29e8ae59))
* **aws:** Replace Manually paginated calls with paginator where available ([#9765](https://github.com/cloudquery/cloudquery/issues/9765)) ([b64d152](https://github.com/cloudquery/cloudquery/commit/b64d152719c48fdc80118e88951c3dd1c4f45135))
* **aws:** Replace more manual pagination with Paginator objects ([#9754](https://github.com/cloudquery/cloudquery/issues/9754)) ([1d27dca](https://github.com/cloudquery/cloudquery/commit/1d27dcac0ccece45e2a2adea4803b5d9f5a458af))
* **deps:** Update module github.com/aws/aws-sdk-go-v2 to v1.17.8 ([#9781](https://github.com/cloudquery/cloudquery/issues/9781)) ([69bb790](https://github.com/cloudquery/cloudquery/commit/69bb790afbeac9ff01a41e71c8f631fb60fe64d1))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/config to v1.18.20 ([#9782](https://github.com/cloudquery/cloudquery/issues/9782)) ([1febd5b](https://github.com/cloudquery/cloudquery/commit/1febd5bbd944459a2fcbe380eb90385ecccfb079))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/feature/s3/manager to v1.11.61 ([#9791](https://github.com/cloudquery/cloudquery/issues/9791)) ([f9dcef8](https://github.com/cloudquery/cloudquery/commit/f9dcef81bb81da123b6820ef2c4b204325e64203))
* Fix case on arn attribute ([#9757](https://github.com/cloudquery/cloudquery/issues/9757)) ([0719095](https://github.com/cloudquery/cloudquery/commit/0719095cdcaa5526f58286431cbe2e1dd2228c0b))
* Use aws.ToString(output.NextToken) when checking tokens ([#9750](https://github.com/cloudquery/cloudquery/issues/9750)) ([7670494](https://github.com/cloudquery/cloudquery/commit/767049489fffa359beac6a4ec00ad76ceeb16224))

## [16.0.1](https://github.com/cloudquery/cloudquery/compare/plugins-source-aws-v16.0.0...plugins-source-aws-v16.0.1) (2023-04-05)


### Bug Fixes

* **aws:** Properly Handle Pagination for Guardduty Findings ([#9746](https://github.com/cloudquery/cloudquery/issues/9746)) ([b9249e0](https://github.com/cloudquery/cloudquery/commit/b9249e0b960b3b3776d83e9bdb1a283f548817a4))

## [16.0.0](https://github.com/cloudquery/cloudquery/compare/plugins-source-aws-v15.7.0...plugins-source-aws-v16.0.0) (2023-04-05)


### This Release has the Following Changes to Tables
- Table `aws_ec2_subnets`: column `account_id` removed from table (:warning: breaking)
- Table `aws_ec2_subnets`: column `region` removed from table (:warning: breaking)
- Table `aws_ec2_subnets`: column added with name `request_account_id (PK)` and type `String` (:warning: breaking)
- Table `aws_ec2_subnets`: column added with name `request_region (PK)` and type `String` (:warning: breaking)
- Table `aws_iam_group_last_accessed_details`: column `arn` removed from table (:warning: breaking)
- Table `aws_iam_group_last_accessed_details`: column added with name `group_arn (PK)` and type `String` (:warning: breaking)
- Table `aws_iam_group_last_accessed_details`: primary key constraint added to column `account_id` (:warning: breaking)
- Table `aws_iam_group_policies`: column `group_id` removed from table (:warning: breaking)
- Table `aws_iam_group_policies`: primary key constraint added to column `account_id` (:warning: breaking)
- Table `aws_iam_group_policies`: primary key constraint added to column `group_arn` (:warning: breaking)
- Table `aws_iam_group_policies`: primary key constraint added to column `policy_name` (:warning: breaking)
- Table `aws_iam_group_policies`: primary key constraint removed from column `_cq_id` (:warning: breaking)
- Table `aws_iam_groups`: column `id` removed from table (:warning: breaking)
- Table `aws_iam_groups`: primary key constraint added to column `arn` (:warning: breaking)
- Table `aws_iam_role_last_accessed_details`: column `arn` removed from table (:warning: breaking)
- Table `aws_iam_role_last_accessed_details`: column added with name `role_arn (PK)` and type `String` (:warning: breaking)
- Table `aws_iam_role_last_accessed_details`: primary key constraint added to column `account_id` (:warning: breaking)
- Table `aws_iam_role_policies`: column order changed for `policy_name`
- Table `aws_iam_roles`: column `id` removed from table (:warning: breaking)
- Table `aws_iam_roles`: primary key constraint added to column `arn` (:warning: breaking)
- Table `aws_iam_signing_certificates`: primary key constraint added to column `account_id` (:warning: breaking)
- Table `aws_iam_ssh_public_keys`: primary key constraint added to column `account_id` (:warning: breaking)
- Table `aws_iam_ssh_public_keys`: primary key constraint added to column `user_arn` (:warning: breaking)
- Table `aws_iam_user_groups`: primary key constraint added to column `account_id` (:warning: breaking)
- Table `aws_iam_user_groups`: primary key constraint added to column `user_arn` (:warning: breaking)
- Table `aws_iam_user_groups`: primary key constraint removed from column `user_id` (:warning: breaking)
- Table `aws_iam_user_last_accessed_details`: column `arn` removed from table (:warning: breaking)
- Table `aws_iam_user_last_accessed_details`: column added with name `user_arn (PK)` and type `String` (:warning: breaking)
- Table `aws_iam_user_last_accessed_details`: primary key constraint added to column `account_id` (:warning: breaking)
- Table `aws_iam_user_policies`: primary key constraint added to column `account_id` (:warning: breaking)
- Table `aws_iam_user_policies`: primary key constraint added to column `policy_name` (:warning: breaking)
- Table `aws_iam_user_policies`: primary key constraint added to column `user_arn` (:warning: breaking)
- Table `aws_iam_user_policies`: primary key constraint removed from column `_cq_id` (:warning: breaking)
- Table `aws_iam_users`: column `id` removed from table (:warning: breaking)
- Table `aws_iam_users`: primary key constraint added to column `arn` (:warning: breaking)
- Table `aws_organizations_accounts`: column `account_id` removed from table (:warning: breaking)
- Table `aws_organizations_accounts`: column added with name `request_account_id (PK)` and type `String` (:warning: breaking)
- Table `aws_organizations_roots`: column `account_id` removed from table (:warning: breaking)
- Table `aws_organizations_roots`: column added with name `request_account_id (PK)` and type `String` (:warning: breaking)

### ⚠ BREAKING CHANGES

* **aws:** Add `request_account_id` to `aws_organizations_accounts` primary key ([#9733](https://github.com/cloudquery/cloudquery/issues/9733))
* **aws:** Add `request_account_id` to `aws_organizations_roots` primary key ([#9732](https://github.com/cloudquery/cloudquery/issues/9732))
* **aws:** Add Columns to `aws_ec2_subnets` Primary Key to handle when subnet is shared ([#9731](https://github.com/cloudquery/cloudquery/issues/9731))
* **aws:** Primary Key changed for `aws_iam_signing_certificates`  ([#9677](https://github.com/cloudquery/cloudquery/issues/9677))
* **aws:** Primary Key changed for `aws_iam_ssh_public_keys`  ([#9677](https://github.com/cloudquery/cloudquery/issues/9677))
* **aws:** Primary Key changed for `aws_iam_user_groups`  ([#9677](https://github.com/cloudquery/cloudquery/issues/9677))
* **aws:** Primary Key changed for `aws_iam_user_policies`  ([#9677](https://github.com/cloudquery/cloudquery/issues/9677))
* **aws:** In `aws_iam_user_policies`, primary Key was changed from `(_cq_id)` to `(account_id, user_arn, policy_name)` so in `overwrite` policies are updated and not always appended
* **aws:** Move Role and Group Attached Policies to Separate tables ([#9508](https://github.com/cloudquery/cloudquery/issues/9508))
* **aws:** Replace `arn` with `id` in `aws_eventbridge_event_bus_targets` primary key ([#9648](https://github.com/cloudquery/cloudquery/issues/9648))
* **aws:** Replace `arn` with `id` in `aws_eventbridge_event_bus_targets` primary key as `arn` refers to the ARN of the resource that is being invoked, not the ARN of the target resource.  The `Id` is what uniquely identifies the target within a specific rule.
* **aws:** Change Primary keys for `aws_dynamodb_global_tables` to include `region` ([#9651](https://github.com/cloudquery/cloudquery/issues/9651))
* **aws:** Change `aws_inspector2_findings` Primary Key to Include `request_account_id` and `request_region`  ([#9650](https://github.com/cloudquery/cloudquery/issues/9650))
* **aws:** Add `cluster_arn` to `aws_ecs_cluster_services` primary key as cluster ID is absent in the "old format" ARNs. See [ecs-account-settings.html#ecs-resource-ids](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs-account-settings.html#ecs-resource-ids) for more details.
* **aws:** Remove `region` from `aws_docdb_global_clusters` as global tables are visible from many regions. Additionally, `aws_docdb_global_clusters` now uses account-level multiplexing instead of region-level, so only 1 entry would be returned when previously it could produce several entries.
* **aws:** GuardDuty detector findings PK should include detector ([#9555](https://github.com/cloudquery/cloudquery/issues/9555))
* **aws:** Remove `region` from `aws_neptune_global_clusters` as global clusters are visible from many regions. Additionally, `aws_neptune_global_clusters` now uses account-level multiplexing instead of region-level, so only 1 entry would be returner when previously it could produce several entries.
* **aws:** Missing partial/parent PK for `aws_ssm_instance_patches` ([#8437](https://github.com/cloudquery/cloudquery/issues/8437))
* **aws:** Add `updated_at` to `aws_securityhub_findings` PK ([#8914](https://github.com/cloudquery/cloudquery/issues/8914))
* **aws:** Missing partial PK for aws_ram_resource_share_permissions ([#8433](https://github.com/cloudquery/cloudquery/issues/8433))
* **aws:** Replace the type of columns ending with `_at` to `Timestamp`. Initially it wasn't the case as the API markup of AWS doesn't include the proper annotations. See https://github.com/aws/aws-sdk/issues/492 for more details. The following tables are affected:
    * `aws_guardduty_detector_findings`: `created_at`, `updated_at`
    * `aws_guardduty_detector_members`: `invited_at`, `updated_at`
    * `aws_guardduty_detectors`: `created_at`, `updated_at`
    * `aws_securityhub_findings`: `created_at`, `first_observed_at`, `last_observed_at`, `updated_at`

### Features

* **aws-services:** Support newly added regions ([#9654](https://github.com/cloudquery/cloudquery/issues/9654)) ([69d5f35](https://github.com/cloudquery/cloudquery/commit/69d5f351864c03d83e667692484724ad061d246c))
* **aws:** ELBv2 Listener Rules ([#9484](https://github.com/cloudquery/cloudquery/issues/9484)) ([c235fca](https://github.com/cloudquery/cloudquery/commit/c235fca5de1380c28d795d348b9f845dabc20b63))
* **aws:** More Redshift resources ([#9486](https://github.com/cloudquery/cloudquery/issues/9486)) ([dacc048](https://github.com/cloudquery/cloudquery/commit/dacc04894139e3c99a201d2bed95f9808379b37c))
* **aws:** Move Role and Group Attached Policies to Separate tables ([#9508](https://github.com/cloudquery/cloudquery/issues/9508)) ([18c5351](https://github.com/cloudquery/cloudquery/commit/18c53511372af69aa5d58f5083329effa19fd344))


### Bug Fixes

* **aws:** Add `cluster_arn` to `aws_ecs_cluster_services` primary key ([#9559](https://github.com/cloudquery/cloudquery/issues/9559)) ([3f0c7ee](https://github.com/cloudquery/cloudquery/commit/3f0c7eeffde45cbd2fd355d08c20d649686e9128))
* **aws:** Add `request_account_id` to `aws_organizations_accounts` primary key ([#9733](https://github.com/cloudquery/cloudquery/issues/9733)) ([3c7d7a8](https://github.com/cloudquery/cloudquery/commit/3c7d7a898205b0e4a418af39aeaf33ede057f1d2))
* **aws:** Add `request_account_id` to `aws_organizations_roots` primary key ([#9732](https://github.com/cloudquery/cloudquery/issues/9732)) ([2fbcb6b](https://github.com/cloudquery/cloudquery/commit/2fbcb6bc9f597de927dc5a2b49d8c1e57b3d1fa0))
* **aws:** Add `updated_at` to `aws_securityhub_findings` PK ([#8914](https://github.com/cloudquery/cloudquery/issues/8914)) ([59b25d3](https://github.com/cloudquery/cloudquery/commit/59b25d39da116570a7a8c7fa1f91906fa6864cb3))
* **aws:** Add Columns to `aws_ec2_subnets` Primary Key to handle when subnet is shared ([#9731](https://github.com/cloudquery/cloudquery/issues/9731)) ([c18928f](https://github.com/cloudquery/cloudquery/commit/c18928f3aa70c819ef0399032b18b135ff59615c))
* **aws:** Change `aws_inspector2_findings` Primary Key to Include `request_account_id` and `request_region`  ([#9650](https://github.com/cloudquery/cloudquery/issues/9650)) ([e51c680](https://github.com/cloudquery/cloudquery/commit/e51c680153a22eabeed38a658bf49816e4bc6e5a))
* **aws:** Change Primary keys for `aws_dynamodb_global_tables` to include `region` ([#9651](https://github.com/cloudquery/cloudquery/issues/9651)) ([63baf42](https://github.com/cloudquery/cloudquery/commit/63baf42cf4405d5d5b673566d11adef788393461))
* **aws:** GuardDuty detector findings PK should include detector ([#9555](https://github.com/cloudquery/cloudquery/issues/9555)) ([71f0826](https://github.com/cloudquery/cloudquery/commit/71f0826a1d3ffca339860819feb459b344dcaea3))
* **aws:** Implement proper tag pagination ([#9646](https://github.com/cloudquery/cloudquery/issues/9646)) ([bcb82d6](https://github.com/cloudquery/cloudquery/commit/bcb82d6a6c3530726e9805dbc8f8e1dc964df341))
* **aws:** Missing partial PK for aws_ram_resource_share_permissions ([#8433](https://github.com/cloudquery/cloudquery/issues/8433)) ([bf48d05](https://github.com/cloudquery/cloudquery/commit/bf48d053e99ca5f87d60480026be4224580c8e5d))
* **aws:** Missing partial/parent PK for `aws_ssm_instance_patches` ([#8437](https://github.com/cloudquery/cloudquery/issues/8437)) ([49bdcd3](https://github.com/cloudquery/cloudquery/commit/49bdcd35ecc18de5b8be34655d522f821ad0ccf1))
* **aws:** Only send images once ([#9644](https://github.com/cloudquery/cloudquery/issues/9644)) ([d883e84](https://github.com/cloudquery/cloudquery/commit/d883e84f38405373bcf3e172b382c089d5dc37c3))
* **aws:** Primary Key changed for `aws_iam_groups`  ([#9677](https://github.com/cloudquery/cloudquery/issues/9677)) ([03b80a1](https://github.com/cloudquery/cloudquery/commit/03b80a10b84c932ecb3530058b5316ce94ff8649))
* **aws:** Primary Key changed for `aws_iam_signing_certificates`  ([#9677](https://github.com/cloudquery/cloudquery/issues/9677)) ([03b80a1](https://github.com/cloudquery/cloudquery/commit/03b80a10b84c932ecb3530058b5316ce94ff8649))
* **aws:** Primary Key changed for `aws_iam_ssh_public_keys`  ([#9677](https://github.com/cloudquery/cloudquery/issues/9677)) ([03b80a1](https://github.com/cloudquery/cloudquery/commit/03b80a10b84c932ecb3530058b5316ce94ff8649))
* **aws:** Primary Key changed for `aws_iam_user_groups`  ([#9677](https://github.com/cloudquery/cloudquery/issues/9677)) ([03b80a1](https://github.com/cloudquery/cloudquery/commit/03b80a10b84c932ecb3530058b5316ce94ff8649))
* **aws:** Primary Key changed for `aws_iam_user_policies`  ([#9677](https://github.com/cloudquery/cloudquery/issues/9677)) ([03b80a1](https://github.com/cloudquery/cloudquery/commit/03b80a10b84c932ecb3530058b5316ce94ff8649))
* **aws:** Primary Key changed for `aws_iam_users`  ([#9677](https://github.com/cloudquery/cloudquery/issues/9677)) ([03b80a1](https://github.com/cloudquery/cloudquery/commit/03b80a10b84c932ecb3530058b5316ce94ff8649))
* **aws:** Remove `group_id` from `aws_iam_group_policies`  ([#9677](https://github.com/cloudquery/cloudquery/issues/9677)) ([03b80a1](https://github.com/cloudquery/cloudquery/commit/03b80a10b84c932ecb3530058b5316ce94ff8649))
* **aws:** Remove `id` from `aws_iam_groups`  ([#9677](https://github.com/cloudquery/cloudquery/issues/9677)) ([03b80a1](https://github.com/cloudquery/cloudquery/commit/03b80a10b84c932ecb3530058b5316ce94ff8649))
* **aws:** Remove `id` from `aws_iam_roles`  ([#9677](https://github.com/cloudquery/cloudquery/issues/9677)) ([03b80a1](https://github.com/cloudquery/cloudquery/commit/03b80a10b84c932ecb3530058b5316ce94ff8649))
* **aws:** Remove `id` from `aws_iam_users`  ([#9677](https://github.com/cloudquery/cloudquery/issues/9677)) ([03b80a1](https://github.com/cloudquery/cloudquery/commit/03b80a10b84c932ecb3530058b5316ce94ff8649))
* **aws:** Remove `region` from `aws_docdb_global_clusters` ([#9558](https://github.com/cloudquery/cloudquery/issues/9558)) ([ea9750b](https://github.com/cloudquery/cloudquery/commit/ea9750be6d36a2c5b44328a98f867aedc977c871))
* **aws:** Remove `region` from `aws_neptune_global_clusters` ([#9556](https://github.com/cloudquery/cloudquery/issues/9556)) ([f09e767](https://github.com/cloudquery/cloudquery/commit/f09e7670b78f1813c911707d1f31593de1f55ab3))
* **aws:** Remove `tags` from `aws_eventbridge_event_bus_targets` ([#9648](https://github.com/cloudquery/cloudquery/issues/9648)) ([fa64254](https://github.com/cloudquery/cloudquery/commit/fa642542fa4417796ee14be7b295fa8ad55d2e80))
* **aws:** Rename `arn` field in `aws_iam_group_last_accessed_details` ([#9677](https://github.com/cloudquery/cloudquery/issues/9677)) ([03b80a1](https://github.com/cloudquery/cloudquery/commit/03b80a10b84c932ecb3530058b5316ce94ff8649))
* **aws:** Rename `arn` field in `aws_iam_role_last_accessed_details` ([#9677](https://github.com/cloudquery/cloudquery/issues/9677)) ([03b80a1](https://github.com/cloudquery/cloudquery/commit/03b80a10b84c932ecb3530058b5316ce94ff8649))
* **aws:** Rename `arn` field in `aws_iam_user_last_accessed_details` ([#9677](https://github.com/cloudquery/cloudquery/issues/9677)) ([03b80a1](https://github.com/cloudquery/cloudquery/commit/03b80a10b84c932ecb3530058b5316ce94ff8649))
* **aws:** Replace `arn` with `id` in `aws_eventbridge_event_bus_targets` primary key ([#9648](https://github.com/cloudquery/cloudquery/issues/9648)) ([fa64254](https://github.com/cloudquery/cloudquery/commit/fa642542fa4417796ee14be7b295fa8ad55d2e80))
* **aws:** Replace the type of columns ending with `_at` to `Timestamp` ([#8912](https://github.com/cloudquery/cloudquery/issues/8912)) ([d3b9f71](https://github.com/cloudquery/cloudquery/commit/d3b9f710a21df7b22c89646f5e2cfd636156a7fb))
* **aws:** Tag Pagination Properly Handled ([#9611](https://github.com/cloudquery/cloudquery/issues/9611)) ([1af31aa](https://github.com/cloudquery/cloudquery/commit/1af31aad2afbe9e0dfb92e08fd3998122cb5a34b))
* **aws:** Update Primary key for `aws_iam_group_policies`  ([#9677](https://github.com/cloudquery/cloudquery/issues/9677)) ([03b80a1](https://github.com/cloudquery/cloudquery/commit/03b80a10b84c932ecb3530058b5316ce94ff8649))
* **deps:** Update github.com/gocarina/gocsv digest to 9a18a84 ([#9563](https://github.com/cloudquery/cloudquery/issues/9563)) ([4c5432b](https://github.com/cloudquery/cloudquery/commit/4c5432b2f57b66ccfe39f9fff66a161feaedfb6b))
* **deps:** Update golang.org/x/exp digest to 10a5072 ([#9587](https://github.com/cloudquery/cloudquery/issues/9587)) ([31f913f](https://github.com/cloudquery/cloudquery/commit/31f913f8e3538a2ba41b089bb11eae78aaf42ab2))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/feature/s3/manager to v1.11.60 ([#9616](https://github.com/cloudquery/cloudquery/issues/9616)) ([d155d28](https://github.com/cloudquery/cloudquery/commit/d155d28f4956be7b2e32ed163f62b4e05432cf6f))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/ecrpublic to v1.15.6 ([#9617](https://github.com/cloudquery/cloudquery/issues/9617)) ([710b564](https://github.com/cloudquery/cloudquery/commit/710b564636a9904d3b1e658420043b7a411bb9e7))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/efs to v1.19.9 ([#9618](https://github.com/cloudquery/cloudquery/issues/9618)) ([399e56e](https://github.com/cloudquery/cloudquery/commit/399e56eb271e6515593335075048916561d85a3e))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/eks to v1.27.8 ([#9619](https://github.com/cloudquery/cloudquery/issues/9619)) ([9ad0589](https://github.com/cloudquery/cloudquery/commit/9ad05893da9dbe40bff26eb3727ac8b93d8bbc70))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/elasticache to v1.26.6 ([#9620](https://github.com/cloudquery/cloudquery/issues/9620)) ([f33de4f](https://github.com/cloudquery/cloudquery/commit/f33de4f7a9cb0fd138a2e6d41b72185e8224a4e8))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/elasticbeanstalk to v1.15.6 ([#9621](https://github.com/cloudquery/cloudquery/issues/9621)) ([3c144ba](https://github.com/cloudquery/cloudquery/commit/3c144ba6efac23b9727fc586a75c93fb344aa58e))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/elasticloadbalancing to v1.15.6 ([#9622](https://github.com/cloudquery/cloudquery/issues/9622)) ([da307f7](https://github.com/cloudquery/cloudquery/commit/da307f789d0c36ce084a0d9686e57e08b71ac7cc))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2 to v1.19.7 ([#9623](https://github.com/cloudquery/cloudquery/issues/9623)) ([1ca1a82](https://github.com/cloudquery/cloudquery/commit/1ca1a82bfd0bdb089fb9b6c0bfb0df12e76f94e8))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/elasticsearchservice to v1.18.7 ([#9624](https://github.com/cloudquery/cloudquery/issues/9624)) ([efa8c2c](https://github.com/cloudquery/cloudquery/commit/efa8c2cea2240f31ad1ecce493c6002a41896ec3))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/elastictranscoder to v1.14.6 ([#9625](https://github.com/cloudquery/cloudquery/issues/9625)) ([3322c94](https://github.com/cloudquery/cloudquery/commit/3322c94947ab6be38e2f51da0d7afbbe41c56243))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/firehose to v1.16.8 ([#9626](https://github.com/cloudquery/cloudquery/issues/9626)) ([279dd25](https://github.com/cloudquery/cloudquery/commit/279dd25a7beb80528c98564832916c825d34c6cf))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/fsx to v1.28.7 ([#9627](https://github.com/cloudquery/cloudquery/issues/9627)) ([3e0491a](https://github.com/cloudquery/cloudquery/commit/3e0491a269e05131d98191fa6cf173d904879968))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/glacier to v1.14.7 ([#9628](https://github.com/cloudquery/cloudquery/issues/9628)) ([9d10696](https://github.com/cloudquery/cloudquery/commit/9d1069602f5e5c95d636dd6a4988d4e5c6d021f3))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/iam to v1.19.8 ([#9629](https://github.com/cloudquery/cloudquery/issues/9629)) ([807fc7c](https://github.com/cloudquery/cloudquery/commit/807fc7c75e802072c1ed4444558fdde444adb88c))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/identitystore to v1.16.6 ([#9630](https://github.com/cloudquery/cloudquery/issues/9630)) ([e3dc54c](https://github.com/cloudquery/cloudquery/commit/e3dc54c63f2a55579f894f04df3b780baa908408))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/inspector to v1.13.6 ([#9631](https://github.com/cloudquery/cloudquery/issues/9631)) ([55706fc](https://github.com/cloudquery/cloudquery/commit/55706fce020cd4f26f019d9c35c0f932fbf54374))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/inspector2 to v1.11.7 ([#9632](https://github.com/cloudquery/cloudquery/issues/9632)) ([791f9a4](https://github.com/cloudquery/cloudquery/commit/791f9a47cc96a3fedab5ec8ed3ab5173af3b672a))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/internal/endpoint-discovery to v1.7.25 ([#9634](https://github.com/cloudquery/cloudquery/issues/9634)) ([16a220f](https://github.com/cloudquery/cloudquery/commit/16a220fc62ad738fcda8d811083e1261ccb73282))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/kafka to v1.19.7 ([#9635](https://github.com/cloudquery/cloudquery/issues/9635)) ([540952d](https://github.com/cloudquery/cloudquery/commit/540952d7610c2c3afc96743e277c99beff7f5a62))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/kinesis to v1.17.8 ([#9636](https://github.com/cloudquery/cloudquery/issues/9636)) ([f08edc4](https://github.com/cloudquery/cloudquery/commit/f08edc44d8e4f3eec23f4b14628c110cfa316354))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/kms to v1.20.8 ([#9637](https://github.com/cloudquery/cloudquery/issues/9637)) ([e2c1a79](https://github.com/cloudquery/cloudquery/commit/e2c1a797d88d54c3330bbf39e7c0eb53e26f21d1))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.44.1 ([#9520](https://github.com/cloudquery/cloudquery/issues/9520)) ([202c31b](https://github.com/cloudquery/cloudquery/commit/202c31b2788c3df35b5df7d07fdc750f92e7bb23))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.44.2 ([#9661](https://github.com/cloudquery/cloudquery/issues/9661)) ([a27dc84](https://github.com/cloudquery/cloudquery/commit/a27dc84a9b67b68b5b75b04dd3afe13e2c556082))
* **deps:** Update module github.com/mattn/go-isatty to v0.0.18 ([#9609](https://github.com/cloudquery/cloudquery/issues/9609)) ([5b2908e](https://github.com/cloudquery/cloudquery/commit/5b2908e8260c6e48f8c5fd6b8bd6c772f0c779d1))

## [15.7.0](https://github.com/cloudquery/cloudquery/compare/plugins-source-aws-v15.6.0...plugins-source-aws-v15.7.0) (2023-03-28)


### This Release has the Following Changes to Tables
- Table `aws_cloudformation_stack_set_operation_results` was added
- Table `aws_cloudformation_stack_set_operations` was added
- Table `aws_cloudformation_stack_sets` was added
- Table `aws_dynamodb_backups` was added
- Table `aws_dynamodb_exports` was added
- Table `aws_dynamodb_global_tables` was added
- Table `aws_dynamodbstreams_streams` was added
- Table `aws_eventbridge_event_bus_targets` was added
- Table `aws_eventbridge_replays`: column added with name `description` and type `String`
- Table `aws_eventbridge_replays`: column added with name `destination` and type `JSON`
- Table `aws_eventbridge_replays`: column added with name `replay_arn` and type `String`
- Table `aws_guardduty_detector_filters` was added
- Table `aws_guardduty_detector_findings` was added
- Table `aws_guardduty_detector_intel_sets` was added
- Table `aws_guardduty_detector_ip_sets` was added
- Table `aws_guardduty_detector_publishing_destinations` was added

### Features

* **aws-services:** Support newly added regions ([#9466](https://github.com/cloudquery/cloudquery/issues/9466)) ([be51234](https://github.com/cloudquery/cloudquery/commit/be5123482351c2f6e4bf362e5b8fa03c0d712c9c))
* **aws:** Add Support For DynamoDB Resources ([#9376](https://github.com/cloudquery/cloudquery/issues/9376)) ([bde468b](https://github.com/cloudquery/cloudquery/commit/bde468bcbc73ba32e0fcf075e256248560f6a693))
* **aws:** Add Support for DynamoDB Streams ([#9399](https://github.com/cloudquery/cloudquery/issues/9399)) ([18cd04c](https://github.com/cloudquery/cloudquery/commit/18cd04c086dc70315bb40fef9af864edefa138eb))
* **aws:** Add Support for Stack Sets ([#7924](https://github.com/cloudquery/cloudquery/issues/7924)) ([555a240](https://github.com/cloudquery/cloudquery/commit/555a2407ef0ba27bd3e08324b6122fed6839e471))
* **aws:** More EventBridge resources ([#9408](https://github.com/cloudquery/cloudquery/issues/9408)) ([86a4c7b](https://github.com/cloudquery/cloudquery/commit/86a4c7bdf303a97c10042765ecaabc5d7f8e4342))
* **aws:** More GuardDuty resources ([#9394](https://github.com/cloudquery/cloudquery/issues/9394)) ([5898a88](https://github.com/cloudquery/cloudquery/commit/5898a88a2757678ecd2e65520c11b52fdcf074d4)), closes [#7709](https://github.com/cloudquery/cloudquery/issues/7709)


### Bug Fixes

* **aws-policies:** Fix query for open critical ports for AWS security groups ([#9410](https://github.com/cloudquery/cloudquery/issues/9410)) ([a8a6dc7](https://github.com/cloudquery/cloudquery/commit/a8a6dc7d740a705f6805382a032a7508b9001ae2))
* **aws:** Only resolve ec2 image launch permissions when the image is owned by the AWS Account ([#9406](https://github.com/cloudquery/cloudquery/issues/9406)) ([16b33c4](https://github.com/cloudquery/cloudquery/commit/16b33c4f752e024e67a10fae083af263250fdb21))
* **deps:** Update module github.com/aws/aws-sdk-go-v2 to v1.17.7 ([#9425](https://github.com/cloudquery/cloudquery/issues/9425)) ([c8a4ab1](https://github.com/cloudquery/cloudquery/commit/c8a4ab1aaf52a1ae68f816b26b6bf7c47910501e))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/config to v1.18.19 ([#9426](https://github.com/cloudquery/cloudquery/issues/9426)) ([2017697](https://github.com/cloudquery/cloudquery/commit/2017697a59970f61c79e713054e8d3e4e482c453))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/credentials to v1.13.18 ([#9427](https://github.com/cloudquery/cloudquery/issues/9427)) ([b2ef029](https://github.com/cloudquery/cloudquery/commit/b2ef0292574d3fa03b7cba8d8a6d25031210079a))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/feature/s3/manager to v1.11.59 ([#9429](https://github.com/cloudquery/cloudquery/issues/9429)) ([71c69a1](https://github.com/cloudquery/cloudquery/commit/71c69a110732f30c61e490360dfe0320fe5e211f))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/accessanalyzer to v1.19.8 ([#9434](https://github.com/cloudquery/cloudquery/issues/9434)) ([41b06d4](https://github.com/cloudquery/cloudquery/commit/41b06d4c9eacd65a8a3bbbf2bb31c20862d7bf93))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/acm to v1.17.7 ([#9435](https://github.com/cloudquery/cloudquery/issues/9435)) ([ede1ccc](https://github.com/cloudquery/cloudquery/commit/ede1ccc0b9581f4c248ad5d2c3c53acd8dae0f0c))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/amp to v1.16.6 ([#9436](https://github.com/cloudquery/cloudquery/issues/9436)) ([7947747](https://github.com/cloudquery/cloudquery/commit/794774704fe830498a0e71f4ac57ae63563b2496))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/amplify to v1.13.6 ([#9437](https://github.com/cloudquery/cloudquery/issues/9437)) ([90ea1c5](https://github.com/cloudquery/cloudquery/commit/90ea1c5a3318245fa45c9d8ccc8e08a1b177ee9c))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/apigateway to v1.16.7 ([#9438](https://github.com/cloudquery/cloudquery/issues/9438)) ([2af4fd8](https://github.com/cloudquery/cloudquery/commit/2af4fd8bf49f7bfd455b775dcff1bd143c03c273))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/apigatewayv2 to v1.13.7 ([#9439](https://github.com/cloudquery/cloudquery/issues/9439)) ([41f2b46](https://github.com/cloudquery/cloudquery/commit/41f2b4600cc5a3de801bb4b0911f9b8b7929aaaa))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/autoscalingplans to v1.13.6 ([#9440](https://github.com/cloudquery/cloudquery/issues/9440)) ([7b52b17](https://github.com/cloudquery/cloudquery/commit/7b52b1746df2e26bf8c7bf628399aaa1fe748a8e))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/cloudhsmv2 to v1.14.6 ([#9441](https://github.com/cloudquery/cloudquery/issues/9441)) ([d5217fd](https://github.com/cloudquery/cloudquery/commit/d5217fd68d26588690ffa57081057591431c41bd))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/cloudwatch to v1.25.7 ([#9442](https://github.com/cloudquery/cloudquery/issues/9442)) ([5a4976e](https://github.com/cloudquery/cloudquery/commit/5a4976e6b038f8ed286eef7f5c79c00a7c297af4))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs to v1.20.7 ([#9443](https://github.com/cloudquery/cloudquery/issues/9443)) ([b507da7](https://github.com/cloudquery/cloudquery/commit/b507da7107e9eb16c9de4bfcd7eca17406733a90))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/codebuild to v1.20.7 ([#9444](https://github.com/cloudquery/cloudquery/issues/9444)) ([ff20bd5](https://github.com/cloudquery/cloudquery/commit/ff20bd5179f4b9c79e423f11e3f53446695d8c47))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/codepipeline to v1.14.6 ([#9451](https://github.com/cloudquery/cloudquery/issues/9451)) ([795ce95](https://github.com/cloudquery/cloudquery/commit/795ce9516ef207fc35c2a3b18ce63cedb9bbf0aa))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/cognitoidentity to v1.15.6 ([#9452](https://github.com/cloudquery/cloudquery/issues/9452)) ([5b0d211](https://github.com/cloudquery/cloudquery/commit/5b0d211646f3d3616490e4e930e0b8d4e4a21fc3))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider to v1.22.6 ([#9453](https://github.com/cloudquery/cloudquery/issues/9453)) ([85c0356](https://github.com/cloudquery/cloudquery/commit/85c0356b766df1465aa6fba05e72678a93036a0c))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/computeoptimizer to v1.21.5 ([#9454](https://github.com/cloudquery/cloudquery/issues/9454)) ([1a5a796](https://github.com/cloudquery/cloudquery/commit/1a5a796065f4032fd1c9d6ab5bc324036f4126a9))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/dax to v1.12.6 ([#9455](https://github.com/cloudquery/cloudquery/issues/9455)) ([1571135](https://github.com/cloudquery/cloudquery/commit/15711353a2a6537a2609fcf8077a410f19b0b556))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/directconnect to v1.18.8 ([#9456](https://github.com/cloudquery/cloudquery/issues/9456)) ([7abcd34](https://github.com/cloudquery/cloudquery/commit/7abcd3485e0740e5a0dddf464ca00a62c56f085e))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/docdb to v1.20.6 ([#9457](https://github.com/cloudquery/cloudquery/issues/9457)) ([09f7a0e](https://github.com/cloudquery/cloudquery/commit/09f7a0e07bad30830af896e2cfc6e30e5c9a3f74))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/ecr to v1.18.7 ([#9458](https://github.com/cloudquery/cloudquery/issues/9458)) ([b719a11](https://github.com/cloudquery/cloudquery/commit/b719a11ca1bf8c54c862719ad6274ba17f77273d))

## [15.6.0](https://github.com/cloudquery/cloudquery/compare/plugins-source-aws-v15.5.0...plugins-source-aws-v15.6.0) (2023-03-23)


### Features

* **aws:** Autoscaling Plans ([#9328](https://github.com/cloudquery/cloudquery/issues/9328)) ([b438d31](https://github.com/cloudquery/cloudquery/commit/b438d311a0ac994e3268b4dd739ee2b2d4451997))


### This Release has the Following Changes to Tables
- Table `aws_autoscaling_plan_resources` was added
- Table `aws_autoscaling_plans` was added

### This Release has the Following Changes to Tables
- Table `aws_autoscaling_plan_resources` was added
- Table `aws_autoscaling_plans` was added

### Bug Fixes

* **aws:** Paginate EBS Snapshots for more consistent throughput ([#9374](https://github.com/cloudquery/cloudquery/issues/9374)) ([156013a](https://github.com/cloudquery/cloudquery/commit/156013a358597840b2c312561b09d4a3688e1d45))
* **aws:** Skip backtrack fetching if backtrack is disabled for RDS cluster ([#9352](https://github.com/cloudquery/cloudquery/issues/9352)) ([3109a6c](https://github.com/cloudquery/cloudquery/commit/3109a6c628428c2740461f0c38e183d6baed3774)), closes [#9351](https://github.com/cloudquery/cloudquery/issues/9351)

## [15.5.0](https://github.com/cloudquery/cloudquery/compare/plugins-source-aws-v15.4.0...plugins-source-aws-v15.5.0) (2023-03-21)


### This Release has the Following Changes to Tables
- Table `aws_applicationautoscaling_scalable_targets` was added
- Table `aws_applicationautoscaling_scaling_activities` was added
- Table `aws_applicationautoscaling_scheduled_actions` was added
- Table `aws_ssm_document_versions` was added
- Table `aws_ssm_sessions` was added

### Features

* Add `member_id` to `aws_identitystore_group_memberships` ([#9297](https://github.com/cloudquery/cloudquery/issues/9297)) ([62ce2ae](https://github.com/cloudquery/cloudquery/commit/62ce2ae82eb4f27f282387b2ab5ce96e1807f58d))
* **aws:** Add Support for Sagemaker Apps ([#9021](https://github.com/cloudquery/cloudquery/issues/9021)) ([e73ea51](https://github.com/cloudquery/cloudquery/commit/e73ea515b817296fd1ab7987ee5990df39941ec6))
* **aws:** Add tables for Batch Jobs, Job Queues, Job Definitions and Compute Environments ([#9162](https://github.com/cloudquery/cloudquery/issues/9162)) ([0933304](https://github.com/cloudquery/cloudquery/commit/09333041d15b82c6aa47dc4dca6c78d7481f7955))
* **aws:** More ApplicationAutoScaling resources ([#9261](https://github.com/cloudquery/cloudquery/issues/9261)) ([6314a4f](https://github.com/cloudquery/cloudquery/commit/6314a4f8e06e0791c9c75d15bad9bf29ad34c026))
* **aws:** SSM Document Versions, SSM Sessions ([#9255](https://github.com/cloudquery/cloudquery/issues/9255)) ([1490cd9](https://github.com/cloudquery/cloudquery/commit/1490cd96f51fc4a41c80e08197ab581eac92c314)), closes [#7907](https://github.com/cloudquery/cloudquery/issues/7907)


### Bug Fixes

* **aws-policies:** Fix AWS Default Security Group query ([#9267](https://github.com/cloudquery/cloudquery/issues/9267)) ([b598824](https://github.com/cloudquery/cloudquery/commit/b598824ebe70de67b7b384d20235ab22600f561a))
* **aws:** AWS Password Policy Saved Query ([#9079](https://github.com/cloudquery/cloudquery/issues/9079)) ([00e5f86](https://github.com/cloudquery/cloudquery/commit/00e5f86bc2482f149c12ef1597cd718eae6341d8))
* **aws:** Fix DynamoDB PITR Query ([059fdf7](https://github.com/cloudquery/cloudquery/commit/059fdf760f28b977070989f70b61ea8fdad4acbf))
* **aws:** Resolve IoT certificates in parallel instead of serially ([#9264](https://github.com/cloudquery/cloudquery/issues/9264)) ([26557bb](https://github.com/cloudquery/cloudquery/commit/26557bb64590ef0567552a4caec3fab712a0d373))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/config to v1.18.18 ([#9227](https://github.com/cloudquery/cloudquery/issues/9227)) ([f630ecc](https://github.com/cloudquery/cloudquery/commit/f630ecc28c19e8388626c823954dca9f561e3920))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/feature/s3/manager to v1.11.58 ([#9229](https://github.com/cloudquery/cloudquery/issues/9229)) ([f8654b4](https://github.com/cloudquery/cloudquery/commit/f8654b4deaaa1a38c5f653a382c1eb6cff6cec74))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/accessanalyzer to v1.19.7 ([#9230](https://github.com/cloudquery/cloudquery/issues/9230)) ([0766214](https://github.com/cloudquery/cloudquery/commit/0766214e263644107f45da06ea420726baebcf6b))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/computeoptimizer to v1.21.4 ([#9231](https://github.com/cloudquery/cloudquery/issues/9231)) ([d760cbd](https://github.com/cloudquery/cloudquery/commit/d760cbd6f4b0b1480ef651af669ce3157d20fa0d))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/directconnect to v1.18.7 ([#9232](https://github.com/cloudquery/cloudquery/issues/9232)) ([43cd203](https://github.com/cloudquery/cloudquery/commit/43cd203a8d1142c042e91eb571ef4dd518ebf66d))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/docdb to v1.20.5 ([#9233](https://github.com/cloudquery/cloudquery/issues/9233)) ([28e73b3](https://github.com/cloudquery/cloudquery/commit/28e73b325bbaef9b61be2ad6007ddbc1f1fdf88a))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/ecr to v1.18.6 ([#9234](https://github.com/cloudquery/cloudquery/issues/9234)) ([809edfa](https://github.com/cloudquery/cloudquery/commit/809edfa6bf241889572b900c8ef53d58f7109328))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/ecrpublic to v1.15.5 ([#9235](https://github.com/cloudquery/cloudquery/issues/9235)) ([d681303](https://github.com/cloudquery/cloudquery/commit/d68130320ffcd1e28c9d70bef2b3f3e0af952f66))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/efs to v1.19.8 ([#9236](https://github.com/cloudquery/cloudquery/issues/9236)) ([f0451ff](https://github.com/cloudquery/cloudquery/commit/f0451ff4ced4cf6a70b8744bd3ca34db50c041fe))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/eks to v1.27.7 ([#9237](https://github.com/cloudquery/cloudquery/issues/9237)) ([f36e818](https://github.com/cloudquery/cloudquery/commit/f36e818c5eff8de623a757da57cad695616c66a3))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/elasticache to v1.26.5 ([#9238](https://github.com/cloudquery/cloudquery/issues/9238)) ([8386ed9](https://github.com/cloudquery/cloudquery/commit/8386ed92fe4136beb806ae7ad518fea4e3a9c233))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/elasticbeanstalk to v1.15.5 ([#9239](https://github.com/cloudquery/cloudquery/issues/9239)) ([17f20d4](https://github.com/cloudquery/cloudquery/commit/17f20d4ddc005c142056d76c5447a2df626949ac))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/elasticloadbalancing to v1.15.5 ([#9240](https://github.com/cloudquery/cloudquery/issues/9240)) ([5e35494](https://github.com/cloudquery/cloudquery/commit/5e3549406f8e513c43d1d6dde13d798e42bc0d5e))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2 to v1.19.6 ([#9241](https://github.com/cloudquery/cloudquery/issues/9241)) ([3c535cf](https://github.com/cloudquery/cloudquery/commit/3c535cf5c8d2d03f5a711c0a1657316da3193fb9))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/elasticsearchservice to v1.18.6 ([#9242](https://github.com/cloudquery/cloudquery/issues/9242)) ([9555a04](https://github.com/cloudquery/cloudquery/commit/9555a04f629703e61af6343ead0adc1154372678))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/elastictranscoder to v1.14.5 ([#9243](https://github.com/cloudquery/cloudquery/issues/9243)) ([bcde085](https://github.com/cloudquery/cloudquery/commit/bcde0854cd49ecef5c9a86d77aa70b815344520b))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/firehose to v1.16.7 ([#9244](https://github.com/cloudquery/cloudquery/issues/9244)) ([43a4a7e](https://github.com/cloudquery/cloudquery/commit/43a4a7e5c84cc317a64dca6bb333125b813ac748))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/fsx to v1.28.6 ([#9245](https://github.com/cloudquery/cloudquery/issues/9245)) ([1946362](https://github.com/cloudquery/cloudquery/commit/1946362c77f517cae64a77dff3721422552b88a8))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/glacier to v1.14.6 ([#9246](https://github.com/cloudquery/cloudquery/issues/9246)) ([248ae09](https://github.com/cloudquery/cloudquery/commit/248ae09a2093f2996ae849137fa495dca61b73dc))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/iam to v1.19.6 ([#9247](https://github.com/cloudquery/cloudquery/issues/9247)) ([d8df47a](https://github.com/cloudquery/cloudquery/commit/d8df47a4bfe31d41369be666d4efb7b867d9bce4))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/identitystore to v1.16.5 ([#9248](https://github.com/cloudquery/cloudquery/issues/9248)) ([f79a817](https://github.com/cloudquery/cloudquery/commit/f79a817fb2df626cc60b38439afb536cbcf7fc89))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/inspector to v1.13.5 ([#9249](https://github.com/cloudquery/cloudquery/issues/9249)) ([b9060b1](https://github.com/cloudquery/cloudquery/commit/b9060b18bda146b6c60a9d7439b0d1aac8c26aeb))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/inspector2 to v1.11.6 ([#9250](https://github.com/cloudquery/cloudquery/issues/9250)) ([a66de92](https://github.com/cloudquery/cloudquery/commit/a66de920b8d5cb74d294b3d1d8d2897043982f04))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.44.0 ([#9167](https://github.com/cloudquery/cloudquery/issues/9167)) ([49d6477](https://github.com/cloudquery/cloudquery/commit/49d647730a85ea6fae51e97194ba61c0625d1331))
* Update endpoints ([#9268](https://github.com/cloudquery/cloudquery/issues/9268)) ([1f0d3f9](https://github.com/cloudquery/cloudquery/commit/1f0d3f9a615a1a8df461ce5ff82547be6209c754))

## [15.4.0](https://github.com/cloudquery/cloudquery/compare/plugins-source-aws-v15.3.0...plugins-source-aws-v15.4.0) (2023-03-15)


### Features

* **aws:** Add AWS EC2 AccountAttributes ([#8937](https://github.com/cloudquery/cloudquery/issues/8937)) ([90b9b3b](https://github.com/cloudquery/cloudquery/commit/90b9b3babba95c0db392acfc59d72aa908ce2f53))
* **aws:** Add AWS Organizations Delegated Services ([#8938](https://github.com/cloudquery/cloudquery/issues/8938)) ([349cf0d](https://github.com/cloudquery/cloudquery/commit/349cf0d05fd82a9e016d766b37a4cc9a3722b4cf))
* **aws:** Add Compute Optimizer Resource ([#8930](https://github.com/cloudquery/cloudquery/issues/8930)) ([da8cc77](https://github.com/cloudquery/cloudquery/commit/da8cc771034688badb2d67bf3de9ad47f08207c4))
* **aws:** Add EC2 image launch permissions resource (`aws_ec2_image_launch_permissions`) ([#8689](https://github.com/cloudquery/cloudquery/issues/8689)) ([025786f](https://github.com/cloudquery/cloudquery/commit/025786fd802c5ea5a593db4788e63f2dc90114b6))
* **aws:** Add Support for VPC Service Permissions ([#8808](https://github.com/cloudquery/cloudquery/issues/8808)) ([05d3342](https://github.com/cloudquery/cloudquery/commit/05d334295f0f5afcb7b14067b966558273442bb5))


### Bug Fixes

* **aws-resources-functions:** Save function configuration instead of failing on `AccessDenied` or `AccessDeniedException` errors in `aws_lambda_functions` ([#8870](https://github.com/cloudquery/cloudquery/issues/8870)) ([ec0b9fd](https://github.com/cloudquery/cloudquery/commit/ec0b9fda20b830eb32e81f6001fd0ca7e9bfc7d0))
* **aws:** Fix description for `aws_organizations_delegated_administrators` ([#8935](https://github.com/cloudquery/cloudquery/issues/8935)) ([2dc6675](https://github.com/cloudquery/cloudquery/commit/2dc6675fbbec7f8a1e81df78ce955960086d1bae))
* **aws:** Fixed `aws_ecs_clusters` fetch to include cluster settings ([#9101](https://github.com/cloudquery/cloudquery/issues/9101)) ([6bc7933](https://github.com/cloudquery/cloudquery/commit/6bc79334a023f3b6d2694ebce8df46fd80a0bca9))
* **aws:** Log Error On Skipped Multiplexer ([#8799](https://github.com/cloudquery/cloudquery/issues/8799)) ([ea30c54](https://github.com/cloudquery/cloudquery/commit/ea30c5405903e49a007aac439cd142135800e708))
* **deps:** Update module github.com/aws/aws-sdk-go-v2 to v1.17.6 ([#8882](https://github.com/cloudquery/cloudquery/issues/8882)) ([5fa0031](https://github.com/cloudquery/cloudquery/commit/5fa0031ff61a92ff1fc086c1fd8b201a5417af36))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/config to v1.18.16 ([#8883](https://github.com/cloudquery/cloudquery/issues/8883)) ([82ffe4d](https://github.com/cloudquery/cloudquery/commit/82ffe4d5aada3b0d3a174fa7a7722ce1a3719993))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/feature/ec2/imds to v1.12.24 ([#8885](https://github.com/cloudquery/cloudquery/issues/8885)) ([674fec4](https://github.com/cloudquery/cloudquery/commit/674fec4c02af4d39613d064ef7d88be62e0a160a))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/feature/s3/manager to v1.11.56 ([#8886](https://github.com/cloudquery/cloudquery/issues/8886)) ([8a3db4b](https://github.com/cloudquery/cloudquery/commit/8a3db4b90501b32fbcc87e5800e2f34fa0b299b7))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/internal/ini to v1.3.31 ([#8889](https://github.com/cloudquery/cloudquery/issues/8889)) ([f8fdb07](https://github.com/cloudquery/cloudquery/commit/f8fdb074573c9fcf394f0f0969156beaaf0ef592))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/internal/v4a to v1.0.22 ([#8890](https://github.com/cloudquery/cloudquery/issues/8890)) ([3c5b412](https://github.com/cloudquery/cloudquery/commit/3c5b41286590308a47207460c93f132e28c8e0a3))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/accessanalyzer to v1.19.6 ([#8891](https://github.com/cloudquery/cloudquery/issues/8891)) ([b90fb07](https://github.com/cloudquery/cloudquery/commit/b90fb07e1e3a3d8e5778b4504b5a38da6bf76d8e))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/acm to v1.17.6 ([#8892](https://github.com/cloudquery/cloudquery/issues/8892)) ([65d5c27](https://github.com/cloudquery/cloudquery/commit/65d5c27dd37e94ad0e94f2223c38b431d231063b))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/amp to v1.16.5 ([#8893](https://github.com/cloudquery/cloudquery/issues/8893)) ([0adec38](https://github.com/cloudquery/cloudquery/commit/0adec388a238244d974e0c67786c5bffdcdef20f))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/amplify to v1.13.5 ([#8894](https://github.com/cloudquery/cloudquery/issues/8894)) ([1f171cc](https://github.com/cloudquery/cloudquery/commit/1f171ccbf5bdf93690f47f7e7554904a3d7ce1af))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/apigateway to v1.16.6 ([#8895](https://github.com/cloudquery/cloudquery/issues/8895)) ([dd84bc9](https://github.com/cloudquery/cloudquery/commit/dd84bc95ad13dc2b692e3cecad927c371d3f785a))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/apigatewayv2 to v1.13.6 ([#8896](https://github.com/cloudquery/cloudquery/issues/8896)) ([5699e4b](https://github.com/cloudquery/cloudquery/commit/5699e4b8df1c1f4f3288c5399016bd977d781a45))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/applicationautoscaling to v1.17.6 ([#8897](https://github.com/cloudquery/cloudquery/issues/8897)) ([1729bed](https://github.com/cloudquery/cloudquery/commit/1729bed54d0889cf1be79bc57c24b924cbe8deb4))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/cloudhsmv2 to v1.14.5 ([#8898](https://github.com/cloudquery/cloudquery/issues/8898)) ([f453645](https://github.com/cloudquery/cloudquery/commit/f4536456eae22700fb2b03fa2c934e2b4ab7ca7b))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/cloudwatch to v1.25.5 ([#8899](https://github.com/cloudquery/cloudquery/issues/8899)) ([34cbb8d](https://github.com/cloudquery/cloudquery/commit/34cbb8d99e0f5183bd885023dc8a138322b860ba))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs to v1.20.6 ([#8900](https://github.com/cloudquery/cloudquery/issues/8900)) ([a19edb4](https://github.com/cloudquery/cloudquery/commit/a19edb498a73ebd8ccc9672ad9fd68d934c82e34))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/codebuild to v1.20.6 ([#8901](https://github.com/cloudquery/cloudquery/issues/8901)) ([319ef28](https://github.com/cloudquery/cloudquery/commit/319ef28cf46460c5602671866d7d61677df52cb9))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/codepipeline to v1.14.5 ([#8907](https://github.com/cloudquery/cloudquery/issues/8907)) ([9751c0a](https://github.com/cloudquery/cloudquery/commit/9751c0a3657cb4c86c748c5a0b57443894286f9b))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/cognitoidentity to v1.15.5 ([#8908](https://github.com/cloudquery/cloudquery/issues/8908)) ([3bc3e3a](https://github.com/cloudquery/cloudquery/commit/3bc3e3ab402ab834c3f9b9f2438ed0e79224f313))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider to v1.22.5 ([#8909](https://github.com/cloudquery/cloudquery/issues/8909)) ([f5dd4cc](https://github.com/cloudquery/cloudquery/commit/f5dd4cc0ab3d4f34d17779f4b9ab02a4d32439cf))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/configservice to v1.29.6 ([#8910](https://github.com/cloudquery/cloudquery/issues/8910)) ([6475fbc](https://github.com/cloudquery/cloudquery/commit/6475fbc4a33a243cb91e80ec67b2ef60220ff730))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/dax to v1.12.5 ([#8911](https://github.com/cloudquery/cloudquery/issues/8911)) ([f7a6c1d](https://github.com/cloudquery/cloudquery/commit/f7a6c1dd2a1cefd3cb7a65ee1da02800d94ab4f7))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.43.0 ([#8949](https://github.com/cloudquery/cloudquery/issues/8949)) ([31dfc63](https://github.com/cloudquery/cloudquery/commit/31dfc634850b699ba7bb7876399270a7367d6c7e))
* Update endpoints ([#8942](https://github.com/cloudquery/cloudquery/issues/8942)) ([806e490](https://github.com/cloudquery/cloudquery/commit/806e49006963974718bafd13ed9eedeea64e1e7d))

## [15.3.0](https://github.com/cloudquery/cloudquery/compare/plugins-source-aws-v15.2.0...plugins-source-aws-v15.3.0) (2023-03-08)


### Features

* Add AWS plugin support for EC2 Launch Templates ([#8675](https://github.com/cloudquery/cloudquery/issues/8675)) ([9ea75a2](https://github.com/cloudquery/cloudquery/commit/9ea75a297d196b7eece6ec888213d86c13c56b89))
* Add RDS BackTracks, Events and Option Groups ([#8671](https://github.com/cloudquery/cloudquery/issues/8671)) ([b8edd0f](https://github.com/cloudquery/cloudquery/commit/b8edd0ff841da36388496f5deff13ef32992be1d))
* **aws-resources:** Add IAM Access Advisor tables: `aws_iam_group_last_accessed_details`, `aws_iam_policy_last_accessed_details`, `aws_iam_role_last_accessed_details` and `aws_iam_user_last_accessed_details`. These might be slow to sync on some accounts. You can skip them if needed via `skip_tables: ["aws_iam_*_last_accessed_details"]` ([bc53529](https://github.com/cloudquery/cloudquery/commit/bc535299a56a597f7c9d46e87900584b6cae44b0))
* **aws:** Add EC2 DHCP Options (aws_ec2_dhcp_options) ([#8678](https://github.com/cloudquery/cloudquery/issues/8678)) ([ec99007](https://github.com/cloudquery/cloudquery/commit/ec990071678ff8c9ef889084334656f7b79de8e4))
* **aws:** Add Support for CloudTrail Events With Incremental Table support ([#8333](https://github.com/cloudquery/cloudquery/issues/8333)) ([a90b95c](https://github.com/cloudquery/cloudquery/commit/a90b95c809f648730b453d9b8a672a0eb5e8eb68))
* **aws:** Add Support for S3 Bucket Static Website ([#8497](https://github.com/cloudquery/cloudquery/issues/8497)) ([255e9b5](https://github.com/cloudquery/cloudquery/commit/255e9b5be3673dc3d39672bcf97321a0472ecd81))
* **aws:** Add Support for syncing Delegated Administrators ([#8342](https://github.com/cloudquery/cloudquery/issues/8342)) ([7c1dc40](https://github.com/cloudquery/cloudquery/commit/7c1dc40bc07a2b8211ca2e884748e45ffc7a6f0b))
* **aws:** EC2 Spot resources ([#8679](https://github.com/cloudquery/cloudquery/issues/8679)) ([ef464b5](https://github.com/cloudquery/cloudquery/commit/ef464b521e48e72b94c00bfb124393788de8a0e3))
* **aws:** Support for Org Policies, Organizational Units and Roots ([#8134](https://github.com/cloudquery/cloudquery/issues/8134)) ([7cf277e](https://github.com/cloudquery/cloudquery/commit/7cf277e837f88bea6622fadc9d7f682a5030d6d6))
* **docs:** Render tables as a part of the Website and add a [tables search box](https://www.cloudquery.io/tables). The equivalent of the GitHub README.md file is now under each plugin's docs section, for example https://www.cloudquery.io/docs/plugins/sources/aws/tables. The Website HTML page is built from the GitHub markdown file located under each plugin's path in our Website code, for example https://github.com/cloudquery/cloudquery/blob/main/website/pages/docs/plugins/sources/aws/tables.md. For the list of all plugins table files as they are stored on GitHub see https://github.com/cloudquery/cloudquery/tree/main/website/tables ([342b0c5](https://github.com/cloudquery/cloudquery/commit/342b0c569fd28ee26ea3e09ec6d787f85c49f16c))


### Bug Fixes

* **aws-resources:** Use QueryUnescape instead of PathUnescape when decoding iam policy version list document ([#8522](https://github.com/cloudquery/cloudquery/issues/8522)) ([3f28658](https://github.com/cloudquery/cloudquery/commit/3f286582b7ecfeaa5680ae5048f7193ab2e6d5af))
* **aws:** EC2 Image Duplicate Records ([#8434](https://github.com/cloudquery/cloudquery/issues/8434)) ([e4e1599](https://github.com/cloudquery/cloudquery/commit/e4e15990dd9fab116cf9d0d28f538fa4f22ca57e))
* **deps:** Update github.com/gocarina/gocsv digest to 70c27cb ([#8559](https://github.com/cloudquery/cloudquery/issues/8559)) ([edae209](https://github.com/cloudquery/cloudquery/commit/edae209be4399ce08bc2458940a44671a0b062a2))
* **deps:** Update golang.org/x/exp digest to c95f2b4 ([#8560](https://github.com/cloudquery/cloudquery/issues/8560)) ([9c3bd5b](https://github.com/cloudquery/cloudquery/commit/9c3bd5b68f9741a360fde6c54bf3f5f3efe06d9e))
* **deps:** Update module github.com/aws/aws-sdk-go-v2 to v1.17.5 ([#8446](https://github.com/cloudquery/cloudquery/issues/8446)) ([e86922b](https://github.com/cloudquery/cloudquery/commit/e86922b62e01d609bcdbacc6afdc2e51febeb7f0))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/config to v1.18.15 ([#8447](https://github.com/cloudquery/cloudquery/issues/8447)) ([98cb352](https://github.com/cloudquery/cloudquery/commit/98cb352834ea715bcb9365b2c124dc98eb9474db))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/feature/ec2/imds to v1.12.23 ([#8449](https://github.com/cloudquery/cloudquery/issues/8449)) ([c59f43e](https://github.com/cloudquery/cloudquery/commit/c59f43e23944c0ffb4f9762bd3efe70a41e4731f))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/feature/s3/manager to v1.11.55 ([#8450](https://github.com/cloudquery/cloudquery/issues/8450)) ([416a435](https://github.com/cloudquery/cloudquery/commit/416a435304cbef7c228b6ee1bc90ec9d1197ae1c))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/internal/ini to v1.3.30 ([#8453](https://github.com/cloudquery/cloudquery/issues/8453)) ([912401b](https://github.com/cloudquery/cloudquery/commit/912401b0b64ff41ad864403ab0cc3f280a0a6355))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/internal/v4a to v1.0.21 ([#8454](https://github.com/cloudquery/cloudquery/issues/8454)) ([7820d00](https://github.com/cloudquery/cloudquery/commit/7820d00414bebb5890beb2ac26326ce0d5a44199))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/accessanalyzer to v1.19.5 ([#8455](https://github.com/cloudquery/cloudquery/issues/8455)) ([4e55283](https://github.com/cloudquery/cloudquery/commit/4e552832f39514ef133281d310d3181eb65644af))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/acm to v1.17.5 ([#8456](https://github.com/cloudquery/cloudquery/issues/8456)) ([670bd96](https://github.com/cloudquery/cloudquery/commit/670bd965d16708fdaa03dc103f171efd14a520a6))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/amp to v1.16.4 ([#8457](https://github.com/cloudquery/cloudquery/issues/8457)) ([b605e4e](https://github.com/cloudquery/cloudquery/commit/b605e4eac1b97ab45a471f57c35d28c42db2571e))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/amplify to v1.13.4 ([#8458](https://github.com/cloudquery/cloudquery/issues/8458)) ([b203f30](https://github.com/cloudquery/cloudquery/commit/b203f306a716ed07fa06aab7feabaf0e7560a74b))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/apigateway to v1.16.5 ([#8459](https://github.com/cloudquery/cloudquery/issues/8459)) ([729099c](https://github.com/cloudquery/cloudquery/commit/729099c331e18ebb5a55f32e46b4e97f65264e54))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/apigatewayv2 to v1.13.5 ([#8460](https://github.com/cloudquery/cloudquery/issues/8460)) ([01a9d06](https://github.com/cloudquery/cloudquery/commit/01a9d067beea937d7d3ced01b780bc0d713d9f61))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/applicationautoscaling to v1.17.5 ([#8461](https://github.com/cloudquery/cloudquery/issues/8461)) ([38603e6](https://github.com/cloudquery/cloudquery/commit/38603e60aa7a86542c8cdd1c768f4e9f6a9c0b10))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/athena to v1.22.4 ([#8462](https://github.com/cloudquery/cloudquery/issues/8462)) ([03cd7cc](https://github.com/cloudquery/cloudquery/commit/03cd7cc248cca98c7e8d46f89dafbf56f1fc540b))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/cloudhsmv2 to v1.14.4 ([#8463](https://github.com/cloudquery/cloudquery/issues/8463)) ([670b86c](https://github.com/cloudquery/cloudquery/commit/670b86cc0ee51d3bf73b2dbfcee91cf0418c640b))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/cloudwatch to v1.25.4 ([#8464](https://github.com/cloudquery/cloudquery/issues/8464)) ([8e70932](https://github.com/cloudquery/cloudquery/commit/8e70932e41b480361c9c45f92541784473642998))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs to v1.20.5 ([#8465](https://github.com/cloudquery/cloudquery/issues/8465)) ([677d41f](https://github.com/cloudquery/cloudquery/commit/677d41f49b4c9e000f28dcad67901b2a0b314b0c))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/codebuild to v1.20.5 ([#8466](https://github.com/cloudquery/cloudquery/issues/8466)) ([4fdfde6](https://github.com/cloudquery/cloudquery/commit/4fdfde62b8a82a05793dea32c2e1787b13234e46))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/codepipeline to v1.14.4 ([#8467](https://github.com/cloudquery/cloudquery/issues/8467)) ([564d1cd](https://github.com/cloudquery/cloudquery/commit/564d1cdd64b3200193757dc26ad1697932f1d0fd))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/cognitoidentity to v1.15.4 ([#8468](https://github.com/cloudquery/cloudquery/issues/8468)) ([c4b0d80](https://github.com/cloudquery/cloudquery/commit/c4b0d805d0e99902daac4512ed87fe347e09efc9))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider to v1.22.4 ([#8469](https://github.com/cloudquery/cloudquery/issues/8469)) ([a3e1580](https://github.com/cloudquery/cloudquery/commit/a3e1580052008962adc8ba0a1d5152f1717535dc))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/configservice to v1.29.5 ([#8470](https://github.com/cloudquery/cloudquery/issues/8470)) ([3c479c6](https://github.com/cloudquery/cloudquery/commit/3c479c6ceda77ea69314b58c37a7348e373d192b))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/databasemigrationservice to v1.23.5 ([#8697](https://github.com/cloudquery/cloudquery/issues/8697)) ([31e3a71](https://github.com/cloudquery/cloudquery/commit/31e3a717ab6004fe5e0ccaa84c20d54402f93948))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/dax to v1.12.4 ([#8698](https://github.com/cloudquery/cloudquery/issues/8698)) ([aefccae](https://github.com/cloudquery/cloudquery/commit/aefccae69ae44f369e164e8e787a3c4fcfd45126))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/directconnect to v1.18.5 ([#8699](https://github.com/cloudquery/cloudquery/issues/8699)) ([b10bcf3](https://github.com/cloudquery/cloudquery/commit/b10bcf3a9701dd969fc109ef35de809e20d4b812))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/docdb to v1.20.4 ([#8700](https://github.com/cloudquery/cloudquery/issues/8700)) ([c2ab59c](https://github.com/cloudquery/cloudquery/commit/c2ab59c18c2656ffed61648331dc2272d8e05b4c))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/dynamodb to v1.18.6 ([#8701](https://github.com/cloudquery/cloudquery/issues/8701)) ([ad00fa7](https://github.com/cloudquery/cloudquery/commit/ad00fa79fda0996acff4a94005dfdf7cb04d784b))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/ecr to v1.18.5 ([#8702](https://github.com/cloudquery/cloudquery/issues/8702)) ([a6bbc91](https://github.com/cloudquery/cloudquery/commit/a6bbc917d2febfc7a259a503b87ae6acd5375701))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/ecrpublic to v1.15.4 ([#8703](https://github.com/cloudquery/cloudquery/issues/8703)) ([7902150](https://github.com/cloudquery/cloudquery/commit/79021503e450014c382ca46d2c22b0d608289a8f))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/efs to v1.19.7 ([#8704](https://github.com/cloudquery/cloudquery/issues/8704)) ([9805d2b](https://github.com/cloudquery/cloudquery/commit/9805d2b9c7c50a5598fe1d05f476fb23a66f30ee))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/eks to v1.27.5 ([#8705](https://github.com/cloudquery/cloudquery/issues/8705)) ([843d015](https://github.com/cloudquery/cloudquery/commit/843d01559ae24b90050ffd031e072d4dfd5bcd6a))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/elasticache to v1.26.4 ([#8706](https://github.com/cloudquery/cloudquery/issues/8706)) ([09a47d5](https://github.com/cloudquery/cloudquery/commit/09a47d5df6179bf8f5f25080ad18174dd99c5b46))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/elasticbeanstalk to v1.15.4 ([#8707](https://github.com/cloudquery/cloudquery/issues/8707)) ([642a303](https://github.com/cloudquery/cloudquery/commit/642a303c45e71b81b1b307fffb72ed277f9f9ce1))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/elasticloadbalancing to v1.15.4 ([#8708](https://github.com/cloudquery/cloudquery/issues/8708)) ([b9628b7](https://github.com/cloudquery/cloudquery/commit/b9628b78bdf39260d6315915e0f3a1ac62386e4d))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2 to v1.19.5 ([#8709](https://github.com/cloudquery/cloudquery/issues/8709)) ([c15a071](https://github.com/cloudquery/cloudquery/commit/c15a0712bf51f4babd9570607df765d25279ff4b))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/elasticsearchservice to v1.18.5 ([#8710](https://github.com/cloudquery/cloudquery/issues/8710)) ([ac19d06](https://github.com/cloudquery/cloudquery/commit/ac19d06b7e5ca7fb40e3eaa044a66567cfe4858d))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/elastictranscoder to v1.14.4 ([#8711](https://github.com/cloudquery/cloudquery/issues/8711)) ([7d5b530](https://github.com/cloudquery/cloudquery/commit/7d5b530f1ad84c744ade92617cf971cb532a2422))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/firehose to v1.16.5 ([#8712](https://github.com/cloudquery/cloudquery/issues/8712)) ([50eef2e](https://github.com/cloudquery/cloudquery/commit/50eef2e6ec0b7dc0638ac7aa2a46523ce1600c4a))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/fsx to v1.28.5 ([#8713](https://github.com/cloudquery/cloudquery/issues/8713)) ([babee7c](https://github.com/cloudquery/cloudquery/commit/babee7c222f9f0a7acc2302d7ba2ddc429b51768))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/glacier to v1.14.5 ([#8714](https://github.com/cloudquery/cloudquery/issues/8714)) ([4cefef8](https://github.com/cloudquery/cloudquery/commit/4cefef83c8056443a829546b6fb6946d3ded106d))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/guardduty to v1.17.6 ([#8715](https://github.com/cloudquery/cloudquery/issues/8715)) ([c0858fe](https://github.com/cloudquery/cloudquery/commit/c0858fe8c6239a50a98f552ed8d353dafd72d990))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/iam to v1.19.4 ([#8716](https://github.com/cloudquery/cloudquery/issues/8716)) ([968a740](https://github.com/cloudquery/cloudquery/commit/968a740cdaa6172e4d18c6bea422c40c3dd2084d))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/identitystore to v1.16.4 ([#8717](https://github.com/cloudquery/cloudquery/issues/8717)) ([05088bc](https://github.com/cloudquery/cloudquery/commit/05088bc19b086e6d1ae86f031836a12ad6e4038e))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/inspector to v1.13.4 ([#8718](https://github.com/cloudquery/cloudquery/issues/8718)) ([7730a9e](https://github.com/cloudquery/cloudquery/commit/7730a9eba33a19372d6b8621d2be85884dd7300d))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/inspector2 to v1.11.5 ([#8719](https://github.com/cloudquery/cloudquery/issues/8719)) ([89a11f9](https://github.com/cloudquery/cloudquery/commit/89a11f9046911bb09f76b7195b7117edbbebff7b))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.39.0 ([#8344](https://github.com/cloudquery/cloudquery/issues/8344)) ([9c57544](https://github.com/cloudquery/cloudquery/commit/9c57544d06f9a774adcc659bcabd2518a905bdaa))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.39.1 ([#8371](https://github.com/cloudquery/cloudquery/issues/8371)) ([e3274c1](https://github.com/cloudquery/cloudquery/commit/e3274c109739bc107387627d340a713470c3a3c1))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.40.0 ([#8401](https://github.com/cloudquery/cloudquery/issues/8401)) ([4cf36d6](https://github.com/cloudquery/cloudquery/commit/4cf36d68684f37c0407332930766c1ba60807a93))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.41.0 ([#8682](https://github.com/cloudquery/cloudquery/issues/8682)) ([ea9d065](https://github.com/cloudquery/cloudquery/commit/ea9d065ae9f77c6dd990570974630ae6ac3f153e))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.42.0 ([#8725](https://github.com/cloudquery/cloudquery/issues/8725)) ([b83b277](https://github.com/cloudquery/cloudquery/commit/b83b277a2421d1caf46a26c3229041b27a3da148))
* **deps:** Update module github.com/stretchr/testify to v1.8.2 ([#8599](https://github.com/cloudquery/cloudquery/issues/8599)) ([2ec8086](https://github.com/cloudquery/cloudquery/commit/2ec808677328410cc96c97a693ef65022d314c32))
* **docs:** Fix documentation link for AWS S3 bucket websites ([#8502](https://github.com/cloudquery/cloudquery/issues/8502)) ([dcccd40](https://github.com/cloudquery/cloudquery/commit/dcccd40bbeb2df852af46a2e36cf0fcd6524d806))
* Fix typo in aws_organizations_delegated_administrators table name ([#8477](https://github.com/cloudquery/cloudquery/issues/8477)) ([bb081c2](https://github.com/cloudquery/cloudquery/commit/bb081c2ea905cff268f6a6c7c2481cca2a121cfc))
* Update endpoints ([#8499](https://github.com/cloudquery/cloudquery/issues/8499)) ([22fce4e](https://github.com/cloudquery/cloudquery/commit/22fce4e38675f2263fe595c12771440773ad5282))
* Update endpoints ([#8737](https://github.com/cloudquery/cloudquery/issues/8737)) ([3582b8f](https://github.com/cloudquery/cloudquery/commit/3582b8fd12e2e9376f99712fac66025634629608))

## [15.2.0](https://github.com/cloudquery/cloudquery/compare/plugins-source-aws-v15.1.0...plugins-source-aws-v15.2.0) (2023-02-21)


### Features

* **aws:** CloudFormation stack_resources to include stack id ([#8141](https://github.com/cloudquery/cloudquery/issues/8141)) ([2c6bfe2](https://github.com/cloudquery/cloudquery/commit/2c6bfe283ac9b0a3080508608746e4440235043f))
* **aws:** Support RDS Reserved Instances ([#8260](https://github.com/cloudquery/cloudquery/issues/8260)) ([6aef2d1](https://github.com/cloudquery/cloudquery/commit/6aef2d12445c18856ffcdb96f5502fed78e924a0))


### Bug Fixes

* **deps:** Update module github.com/aws/aws-sdk-go-v2/config to v1.18.13 ([#8231](https://github.com/cloudquery/cloudquery/issues/8231)) ([1eb436d](https://github.com/cloudquery/cloudquery/commit/1eb436d4db2f467419413c250c9fd1252d0a2fa5))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/feature/s3/manager to v1.11.53 ([#8233](https://github.com/cloudquery/cloudquery/issues/8233)) ([3bc3b86](https://github.com/cloudquery/cloudquery/commit/3bc3b8613a2e59fea4e0838d3b751e4da12b8379))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/internal/v4a to v1.0.20 ([#8234](https://github.com/cloudquery/cloudquery/issues/8234)) ([6516f73](https://github.com/cloudquery/cloudquery/commit/6516f735ac2edb576afbe168bf56f9d5b25eef71))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/accessanalyzer to v1.19.3 ([#8235](https://github.com/cloudquery/cloudquery/issues/8235)) ([a1d1072](https://github.com/cloudquery/cloudquery/commit/a1d1072899350deab5478f483748ecf149a34226))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/acm to v1.17.3 ([#8236](https://github.com/cloudquery/cloudquery/issues/8236)) ([a5fea7d](https://github.com/cloudquery/cloudquery/commit/a5fea7df92491edf5604db0da79ecfe82af5ee76))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/amp to v1.16.2 ([#8237](https://github.com/cloudquery/cloudquery/issues/8237)) ([b8040ea](https://github.com/cloudquery/cloudquery/commit/b8040eae26d0575a4f1cd8cd0ea8dbbc6023c1dc))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/amplify to v1.13.2 ([#8238](https://github.com/cloudquery/cloudquery/issues/8238)) ([4e719a8](https://github.com/cloudquery/cloudquery/commit/4e719a868026084cf1b08f15954f40a02c9b2620))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/apigateway to v1.16.3 ([#8239](https://github.com/cloudquery/cloudquery/issues/8239)) ([a3a4f79](https://github.com/cloudquery/cloudquery/commit/a3a4f79789bf28b5f1af0d7c47830319bd0273c9))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/apigatewayv2 to v1.13.3 ([#8240](https://github.com/cloudquery/cloudquery/issues/8240)) ([63170f6](https://github.com/cloudquery/cloudquery/commit/63170f6a3457b5347963b1a9b2015b0864008a33))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/applicationautoscaling to v1.17.3 ([#8241](https://github.com/cloudquery/cloudquery/issues/8241)) ([1ee38d6](https://github.com/cloudquery/cloudquery/commit/1ee38d608eae116989b51617308f70053f0b6f5f))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/athena to v1.22.2 ([#8242](https://github.com/cloudquery/cloudquery/issues/8242)) ([bd980fa](https://github.com/cloudquery/cloudquery/commit/bd980fa3a232a3b0c93ebe8ac201c9d8bb1f06e1))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/cloudhsmv2 to v1.14.2 ([#8243](https://github.com/cloudquery/cloudquery/issues/8243)) ([c849f73](https://github.com/cloudquery/cloudquery/commit/c849f73fd4177168be017b74d6ffd3cea18fb94f))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs to v1.20.3 ([#8244](https://github.com/cloudquery/cloudquery/issues/8244)) ([5abb3d4](https://github.com/cloudquery/cloudquery/commit/5abb3d4332433dc2b0710da1de280ca5b896c992))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/codebuild to v1.20.3 ([#8245](https://github.com/cloudquery/cloudquery/issues/8245)) ([16e80e3](https://github.com/cloudquery/cloudquery/commit/16e80e3900cd175d6ce6a36ba4b69d30381cac72))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/codepipeline to v1.14.2 ([#8246](https://github.com/cloudquery/cloudquery/issues/8246)) ([bad0617](https://github.com/cloudquery/cloudquery/commit/bad0617608b1674d04b68d9e3e2aa5e30306c0f6))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/cognitoidentity to v1.15.2 ([#8247](https://github.com/cloudquery/cloudquery/issues/8247)) ([bfbeee0](https://github.com/cloudquery/cloudquery/commit/bfbeee0289d576bd29f9841e1678545849208299))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider to v1.22.2 ([#8248](https://github.com/cloudquery/cloudquery/issues/8248)) ([65ece07](https://github.com/cloudquery/cloudquery/commit/65ece07ce9dd4ac33449747abe9440c2dab62d60))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/configservice to v1.29.3 ([#8249](https://github.com/cloudquery/cloudquery/issues/8249)) ([d91fcde](https://github.com/cloudquery/cloudquery/commit/d91fcde58b264aad60a4500d33f40f9f47e66d65))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/databasemigrationservice to v1.23.3 ([#8250](https://github.com/cloudquery/cloudquery/issues/8250)) ([762e076](https://github.com/cloudquery/cloudquery/commit/762e0760cfd83f590359fbcb8732f8659e15aaac))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/dax to v1.12.2 ([#8251](https://github.com/cloudquery/cloudquery/issues/8251)) ([7e1a123](https://github.com/cloudquery/cloudquery/commit/7e1a12336d703309ab3d3099206634e065a7b3df))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/directconnect to v1.18.3 ([#8252](https://github.com/cloudquery/cloudquery/issues/8252)) ([23292ba](https://github.com/cloudquery/cloudquery/commit/23292ba1d2f7c1afddbf4c72533242b1dac1b85f))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/dynamodb to v1.18.3 ([#8253](https://github.com/cloudquery/cloudquery/issues/8253)) ([2d767aa](https://github.com/cloudquery/cloudquery/commit/2d767aa5a4f6332a65fed0ab07a1fd07ad2c6a02))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/ecr to v1.18.3 ([#8254](https://github.com/cloudquery/cloudquery/issues/8254)) ([fd6fb50](https://github.com/cloudquery/cloudquery/commit/fd6fb50ad70b1b6ada599415e72388b7eee94077))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.38.2 ([#8156](https://github.com/cloudquery/cloudquery/issues/8156)) ([ac2d2d7](https://github.com/cloudquery/cloudquery/commit/ac2d2d70d5c4bc45fb8734bd4deb8a1e36074f6d))
* **deps:** Update module golang.org/x/net to v0.7.0 [SECURITY] ([#8176](https://github.com/cloudquery/cloudquery/issues/8176)) ([fc4cef8](https://github.com/cloudquery/cloudquery/commit/fc4cef86dce4ca76ca8397e897ab744e48975834))
* Update endpoints ([#8265](https://github.com/cloudquery/cloudquery/issues/8265)) ([a7cb153](https://github.com/cloudquery/cloudquery/commit/a7cb153fcc9400a59e70f06f29cb3c610bc1d19d))

## [15.1.0](https://github.com/cloudquery/cloudquery/compare/plugins-source-aws-v15.0.1...plugins-source-aws-v15.1.0) (2023-02-16)


### Features

* Add Support resources (includes AWS Trusted Advisor) ([#8140](https://github.com/cloudquery/cloudquery/issues/8140)) ([a49aae0](https://github.com/cloudquery/cloudquery/commit/a49aae04723d69efacf6a3a5344f25fcfffd4c25))
* **aws:** Use ServiceAccountRegion Multiplexer ([#8158](https://github.com/cloudquery/cloudquery/issues/8158)) ([a06e02f](https://github.com/cloudquery/cloudquery/commit/a06e02f9eef709f21d4d13027c692f64c10f1003))


### Bug Fixes

* **aws:** Fix error for Empty Backup Notification ([#8164](https://github.com/cloudquery/cloudquery/issues/8164)) ([d9be357](https://github.com/cloudquery/cloudquery/commit/d9be357e3f2cefd29d0050be25e13543c9967db1))
* **aws:** Fix error handling int `aws_alternate_contact` fetching ([#8152](https://github.com/cloudquery/cloudquery/issues/8152)) ([fcdf778](https://github.com/cloudquery/cloudquery/commit/fcdf7789fb70f7e972c581fba1f0e88970ec146e))
* **aws:** Fix S3 Access Points ([#8160](https://github.com/cloudquery/cloudquery/issues/8160)) ([25caebf](https://github.com/cloudquery/cloudquery/commit/25caebf5607c5a2366576447ce2d4ef30fa8fa93))
* **aws:** Ignore Default ECR Errors ([#8162](https://github.com/cloudquery/cloudquery/issues/8162)) ([e88e137](https://github.com/cloudquery/cloudquery/commit/e88e137c46585f1982e043a146f8b23fd41bd2fe))

## [15.0.1](https://github.com/cloudquery/cloudquery/compare/plugins-source-aws-v15.0.0...plugins-source-aws-v15.0.1) (2023-02-16)


### Bug Fixes

* **aws:** Use `credentialScope.region` if exists when generating regions data ([#8131](https://github.com/cloudquery/cloudquery/issues/8131)) ([915e829](https://github.com/cloudquery/cloudquery/commit/915e829e84bc8ff38dfe46c0288d5d90e93103e2))

## [15.0.0](https://github.com/cloudquery/cloudquery/compare/plugins-source-aws-v14.0.0...plugins-source-aws-v15.0.0) (2023-02-15)


### ⚠ BREAKING CHANGES

* **aws:** Step Function Executions And MapRuns ([#8130](https://github.com/cloudquery/cloudquery/issues/8130))
* **aws:** Add `arn` to `aws_ec2_managed_prefix_lists` PK ([#8119](https://github.com/cloudquery/cloudquery/issues/8119))

### Bug Fixes

* **aws:** Add `arn` to `aws_ec2_managed_prefix_lists` PK ([#8119](https://github.com/cloudquery/cloudquery/issues/8119)) ([148b06c](https://github.com/cloudquery/cloudquery/commit/148b06ce82be72bc9cf78af95743c2988d4a8263))
* **aws:** Step Function Executions And MapRuns ([#8130](https://github.com/cloudquery/cloudquery/issues/8130)) ([21e4ea1](https://github.com/cloudquery/cloudquery/commit/21e4ea1cbcfe0ca1112ea7f2a01c9a860e72d2e2))

## [14.0.0](https://github.com/cloudquery/cloudquery/compare/plugins-source-aws-v13.3.0...plugins-source-aws-v14.0.0) (2023-02-14)


### ⚠ BREAKING CHANGES

* **aws:** Use `log_group_arn` + `filter_name` for `aws_cloudwatchlogs_metric_filters` PK ([#8061](https://github.com/cloudquery/cloudquery/issues/8061))
* **aws:** Add `account_id` and `region` to `aws_ec2_transit_gateways` PK ([#8060](https://github.com/cloudquery/cloudquery/issues/8060))
* **aws:** Rename `user_arn` field to `arn` in `aws_quicksight_group_members`
* **aws:** Resource Access Manager (RAM) PK issues ([#8010](https://github.com/cloudquery/cloudquery/issues/8010))
* **aws:** Add `resource_share_arn` and `region` to `aws_ram_principals` PK ([#7985](https://github.com/cloudquery/cloudquery/issues/7985))
* **aws:** Add `account_id` and `region` to PK in `aws_directconnect_gateways` ([#7984](https://github.com/cloudquery/cloudquery/issues/7984))
* **aws:** Use `physical_resource_identifier` instead of `physical_resource_id` in `aws_resiliencehub_app_version_resources` and `aws_resiliencehub_app_version_resource_mappings` PKs
* **aws:** PK for `aws_s3_bucket_grants` ([#7822](https://github.com/cloudquery/cloudquery/issues/7822))
* **aws:** PK for `aws_docdb_certificates` ([#7820](https://github.com/cloudquery/cloudquery/issues/7820))
* **aws:** PK for `aws_ec2_images` ([#7821](https://github.com/cloudquery/cloudquery/issues/7821))
* **aws:** Fix PK  for `aws_rds_certificates` ([#7817](https://github.com/cloudquery/cloudquery/issues/7817))

### Features

* Athena resource view creation ([#7908](https://github.com/cloudquery/cloudquery/issues/7908)) ([a8769ea](https://github.com/cloudquery/cloudquery/commit/a8769eaed8c116940717ed84daf356b50924f49f))
* **aws-resources:** Add EMR Cluster Instances ([#7814](https://github.com/cloudquery/cloudquery/issues/7814)) ([f2883d0](https://github.com/cloudquery/cloudquery/commit/f2883d0f31b2c6f9d1c587802386c221b7c26c12))
* **aws-resources:** Add EMR Instance Fleets and Groups ([#7818](https://github.com/cloudquery/cloudquery/issues/7818)) ([df8cfea](https://github.com/cloudquery/cloudquery/commit/df8cfeab17d0b31ebb96019648a3be6469ef330f))
* **aws-resources:** Add EMR Security Configuration ([#7812](https://github.com/cloudquery/cloudquery/issues/7812)) ([3d70111](https://github.com/cloudquery/cloudquery/commit/3d7011122d5e1940cefa209315bef9cd8649bde9))
* **aws:** Add Method and Integration Resources for API Gateway ([#7923](https://github.com/cloudquery/cloudquery/issues/7923)) ([f3d669a](https://github.com/cloudquery/cloudquery/commit/f3d669a8eca6ae855bbbb34d81dbed1263c51b66))
* **aws:** Add New Elasticache Resources and Attributes ([#7925](https://github.com/cloudquery/cloudquery/issues/7925)) ([e18ca20](https://github.com/cloudquery/cloudquery/commit/e18ca20e3801b83deca965f95fe121f49978c091))
* **aws:** Add New Validation for AWS Tags ([#7651](https://github.com/cloudquery/cloudquery/issues/7651)) ([2440e89](https://github.com/cloudquery/cloudquery/commit/2440e89863a903d65beb875667b5d68bc94704f9))
* **aws:** Add Support for CloudWatch Filter Subscriptions ([#8073](https://github.com/cloudquery/cloudquery/issues/8073)) ([cbc2fb1](https://github.com/cloudquery/cloudquery/commit/cbc2fb17f77201ba99ec79570a65e6fc98ad6bcc))
* **aws:** Add Support for Directconnect Locations ([#7906](https://github.com/cloudquery/cloudquery/issues/7906)) ([0187098](https://github.com/cloudquery/cloudquery/commit/0187098f47932400ffd9d06988408a8263dd27dd))
* **aws:** Add Support for EC2 Managed Prefix List ([#7942](https://github.com/cloudquery/cloudquery/issues/7942)) ([743fdd2](https://github.com/cloudquery/cloudquery/commit/743fdd270baaff5ec581eea9a08aff0e6359b6e0))
* **aws:** Add Support for Route53 Operations ([#7944](https://github.com/cloudquery/cloudquery/issues/7944)) ([306f0bd](https://github.com/cloudquery/cloudquery/commit/306f0bd82c91d1d90836daf6fb3b42139cb047ab))
* **aws:** Step Function Resources ([#7911](https://github.com/cloudquery/cloudquery/issues/7911)) ([ffb3275](https://github.com/cloudquery/cloudquery/commit/ffb3275ee92d2fafef053c4bcfb69a885adfdc65))
* **aws:** Support Availability Zones ([#7914](https://github.com/cloudquery/cloudquery/issues/7914)) ([f082057](https://github.com/cloudquery/cloudquery/commit/f082057a9e7083e59d1bee4eddd43705b68b6598))


### Bug Fixes

* **aws:** Add `account_id` and `region` to `aws_appstream_images` PK ([#7972](https://github.com/cloudquery/cloudquery/issues/7972)) ([a7cfd33](https://github.com/cloudquery/cloudquery/commit/a7cfd33cc09f5c607552abd141b6d5a266d4a46e))
* **aws:** Add `account_id` and `region` to `aws_ec2_transit_gateways` PK ([#8060](https://github.com/cloudquery/cloudquery/issues/8060)) ([37fbf63](https://github.com/cloudquery/cloudquery/commit/37fbf6393aa722ad4321d10e439eb33a9db09805))
* **aws:** Add `account_id` and `region` to PK in `aws_directconnect_gateways` ([#7984](https://github.com/cloudquery/cloudquery/issues/7984)) ([588bc88](https://github.com/cloudquery/cloudquery/commit/588bc88c2cc1a0f7434c738884c9c29f04b7c974))
* **aws:** Add `account_id` and region to `aws_quicksight_*` resources ([4d0fdb4](https://github.com/cloudquery/cloudquery/commit/4d0fdb48d95439afe33eff1deef642e3602e0f7b))
* **aws:** Add `request_account_id` & `request_region` to `aws_securityhub_findings` PK ([#7971](https://github.com/cloudquery/cloudquery/issues/7971)) ([aa74371](https://github.com/cloudquery/cloudquery/commit/aa7437154b5ffd89e2ae3c2aa49bda3ded6f99e9))
* **aws:** Add `resource_share_arn` and `region` to `aws_ram_principals` PK ([#7985](https://github.com/cloudquery/cloudquery/issues/7985)) ([fd24a5b](https://github.com/cloudquery/cloudquery/commit/fd24a5b0afc3c260827931a326b43a6973677ff7))
* **aws:** Add parent table ARNs as PKs to child tables for `aws_resiliencehub_*` resources ([2f70712](https://github.com/cloudquery/cloudquery/commit/2f70712201a0981861e294d87969ba8fc4627c90))
* **aws:** Fix `aws_iam_accounts` unmarshaling ([#7899](https://github.com/cloudquery/cloudquery/issues/7899)) ([abf28ed](https://github.com/cloudquery/cloudquery/commit/abf28edbbeb05e119a6ce70dd170f811d6d32683))
* **aws:** Fix PK  for `aws_rds_certificates` ([#7817](https://github.com/cloudquery/cloudquery/issues/7817)) ([8082b2a](https://github.com/cloudquery/cloudquery/commit/8082b2a542e9f75637db11b270965d4e626b3075))
* **aws:** PK for `aws_docdb_certificates` ([#7820](https://github.com/cloudquery/cloudquery/issues/7820)) ([b2b06ea](https://github.com/cloudquery/cloudquery/commit/b2b06eac297f5893729bf3c080b1734548c46906))
* **aws:** PK for `aws_ec2_images` ([#7821](https://github.com/cloudquery/cloudquery/issues/7821)) ([d568cde](https://github.com/cloudquery/cloudquery/commit/d568cde6afcf20e55642f768d6721e05e7825a47))
* **aws:** PK for `aws_s3_bucket_grants` ([#7822](https://github.com/cloudquery/cloudquery/issues/7822)) ([4efbf32](https://github.com/cloudquery/cloudquery/commit/4efbf32f77146b61b62956ca9ef41c925d644d4f))
* **aws:** Remove `account_id` and `region` columns from `aws_resiliencehub_*` PKs ([2f70712](https://github.com/cloudquery/cloudquery/commit/2f70712201a0981861e294d87969ba8fc4627c90))
* **aws:** Rename `user_arn` field to `arn` in `aws_quicksight_group_members` ([4d0fdb4](https://github.com/cloudquery/cloudquery/commit/4d0fdb48d95439afe33eff1deef642e3602e0f7b))
* **aws:** Resource Access Manager (RAM) PK issues ([#8010](https://github.com/cloudquery/cloudquery/issues/8010)) ([901dcf4](https://github.com/cloudquery/cloudquery/commit/901dcf45ee93fa2ef62a9ec851ca4a59549018fc))
* **aws:** Use `log_group_arn` + `filter_name` for `aws_cloudwatchlogs_metric_filters` PK ([#8061](https://github.com/cloudquery/cloudquery/issues/8061)) ([d29f19a](https://github.com/cloudquery/cloudquery/commit/d29f19aaddaa2d0b0e4056b5cd0dd5bebd7848b6))
* **aws:** Use `physical_resource_identifier` instead of `physical_resource_id` in `aws_resiliencehub_app_version_resources` and `aws_resiliencehub_app_version_resource_mappings` PKs ([2f70712](https://github.com/cloudquery/cloudquery/commit/2f70712201a0981861e294d87969ba8fc4627c90))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/codepipeline to v1.14.1 ([#7945](https://github.com/cloudquery/cloudquery/issues/7945)) ([bd1eee8](https://github.com/cloudquery/cloudquery/commit/bd1eee86371a9aeab9f48c66348ea50148df34f1))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/cognitoidentity to v1.15.1 ([#7946](https://github.com/cloudquery/cloudquery/issues/7946)) ([7747eab](https://github.com/cloudquery/cloudquery/commit/7747eabe9c45bf904d7f104265ed033595d3c42c))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider to v1.22.1 ([#7947](https://github.com/cloudquery/cloudquery/issues/7947)) ([f3307f2](https://github.com/cloudquery/cloudquery/commit/f3307f28fb83fb7c4bc3b829e041e6a70ddb632c))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/configservice to v1.29.2 ([#7948](https://github.com/cloudquery/cloudquery/issues/7948)) ([db7b7af](https://github.com/cloudquery/cloudquery/commit/db7b7afea96d383acb9a598bd12d79d5050a4fb0))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/databasemigrationservice to v1.23.2 ([#7949](https://github.com/cloudquery/cloudquery/issues/7949)) ([ff3e2d1](https://github.com/cloudquery/cloudquery/commit/ff3e2d1d3c0f28514b004502d3126334269d18f3))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/dax to v1.12.1 ([#7950](https://github.com/cloudquery/cloudquery/issues/7950)) ([3127200](https://github.com/cloudquery/cloudquery/commit/31272000e3e62103a011629e09a6644c68dfe9fa))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/directconnect to v1.18.2 ([#7951](https://github.com/cloudquery/cloudquery/issues/7951)) ([c220d49](https://github.com/cloudquery/cloudquery/commit/c220d49a2171b772285774eb6ba8e827933e9bfc))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/docdb to v1.20.2 ([#7952](https://github.com/cloudquery/cloudquery/issues/7952)) ([f021f80](https://github.com/cloudquery/cloudquery/commit/f021f80f731d0b3cebd36095122374719660646a))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/dynamodb to v1.18.2 ([#7953](https://github.com/cloudquery/cloudquery/issues/7953)) ([b50cb36](https://github.com/cloudquery/cloudquery/commit/b50cb360654dfd178bd5f1c7a94caad5fbe8900a))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/ecr to v1.18.2 ([#7954](https://github.com/cloudquery/cloudquery/issues/7954)) ([610165a](https://github.com/cloudquery/cloudquery/commit/610165a10f9d60756a3b16a99b89d19dc968c242))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/ecrpublic to v1.15.1 ([#7955](https://github.com/cloudquery/cloudquery/issues/7955)) ([70a2b8b](https://github.com/cloudquery/cloudquery/commit/70a2b8be78105c5a1ff933abbe9adaaaef4dd157))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/ecs to v1.23.2 ([#7956](https://github.com/cloudquery/cloudquery/issues/7956)) ([66b4cee](https://github.com/cloudquery/cloudquery/commit/66b4cee6229698b83c82437c7e4dd0d97ea478d5))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/efs to v1.19.3 ([#7957](https://github.com/cloudquery/cloudquery/issues/7957)) ([202d467](https://github.com/cloudquery/cloudquery/commit/202d4674044ab3e7c881952afab241bbfefbfa00))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/eks to v1.27.2 ([#7958](https://github.com/cloudquery/cloudquery/issues/7958)) ([8f08e65](https://github.com/cloudquery/cloudquery/commit/8f08e651a3b945a81f0c14ebfe15f4afef50d55f))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/elasticache to v1.26.2 ([#7959](https://github.com/cloudquery/cloudquery/issues/7959)) ([4ca09b8](https://github.com/cloudquery/cloudquery/commit/4ca09b8a411df13652a2b1ab4738f900fc743569))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/elasticbeanstalk to v1.15.1 ([#7960](https://github.com/cloudquery/cloudquery/issues/7960)) ([e859fd0](https://github.com/cloudquery/cloudquery/commit/e859fd0d0853fc60c9e6dac57bcc0df9dd230206))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/elasticloadbalancing to v1.15.2 ([#7961](https://github.com/cloudquery/cloudquery/issues/7961)) ([d7b1163](https://github.com/cloudquery/cloudquery/commit/d7b11632370295be2d2155dda9065615bb219ccd))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2 to v1.19.3 ([#7962](https://github.com/cloudquery/cloudquery/issues/7962)) ([4afa209](https://github.com/cloudquery/cloudquery/commit/4afa209e539ae92170b5bbeeea3e8be195a93995))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/elasticsearchservice to v1.18.2 ([#7963](https://github.com/cloudquery/cloudquery/issues/7963)) ([69752f3](https://github.com/cloudquery/cloudquery/commit/69752f3ccc5feaa8cc7d230d423e44ca38fc1fb7))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/elastictranscoder to v1.14.1 ([#7964](https://github.com/cloudquery/cloudquery/issues/7964)) ([d8ca305](https://github.com/cloudquery/cloudquery/commit/d8ca305ef73b08e168612e948b78f1aa70738484))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/emr to v1.22.2 ([#7965](https://github.com/cloudquery/cloudquery/issues/7965)) ([1ee800c](https://github.com/cloudquery/cloudquery/commit/1ee800c1e00846f02ef7bcbf150bb3edd3b9a67c))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/firehose to v1.16.2 ([#7966](https://github.com/cloudquery/cloudquery/issues/7966)) ([48a4faf](https://github.com/cloudquery/cloudquery/commit/48a4faf6dd60ec5aebcd36536a6160de5c028e60))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/fsx to v1.28.2 ([#7967](https://github.com/cloudquery/cloudquery/issues/7967)) ([2ff2fce](https://github.com/cloudquery/cloudquery/commit/2ff2fce8ea67ac3c6996bf7aac257e4e0e3d6aa5))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/glacier to v1.14.2 ([#7968](https://github.com/cloudquery/cloudquery/issues/7968)) ([0a5fe22](https://github.com/cloudquery/cloudquery/commit/0a5fe22309fa33b14f0d1d5ae36869e899b136c7))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/guardduty to v1.17.2 ([#7969](https://github.com/cloudquery/cloudquery/issues/7969)) ([7fcb80a](https://github.com/cloudquery/cloudquery/commit/7fcb80a0f6e64eaee3618e648ff4201476b69c52))
* **deps:** Update module github.com/cloudquery/codegen to v0.2.1 ([#7875](https://github.com/cloudquery/cloudquery/issues/7875)) ([9d10ac1](https://github.com/cloudquery/cloudquery/commit/9d10ac103dc28a46afa351ed0687c0a22a49ecee))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.36.0 ([#7809](https://github.com/cloudquery/cloudquery/issues/7809)) ([c85a9cb](https://github.com/cloudquery/cloudquery/commit/c85a9cb697477520e94a1fd260c56b89da62fc87))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.36.1 ([#7930](https://github.com/cloudquery/cloudquery/issues/7930)) ([39dccc1](https://github.com/cloudquery/cloudquery/commit/39dccc1bf81f4eb02d181ba0c47b37038a4c5455))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.37.0 ([#7933](https://github.com/cloudquery/cloudquery/issues/7933)) ([dc9cffb](https://github.com/cloudquery/cloudquery/commit/dc9cffbf37bbc6fae73a20bf47e6bbf17e74d1f9))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.37.1 ([#8008](https://github.com/cloudquery/cloudquery/issues/8008)) ([c47aac0](https://github.com/cloudquery/cloudquery/commit/c47aac0b5e3190a04299713651b97e360043911f))
* Fix ECR image ARN to use repository name ([#7839](https://github.com/cloudquery/cloudquery/issues/7839)) ([e585d61](https://github.com/cloudquery/cloudquery/commit/e585d615174cb52ea16cf23cee21151b8b1f4a2b))
* Update endpoints ([#7795](https://github.com/cloudquery/cloudquery/issues/7795)) ([7f5260c](https://github.com/cloudquery/cloudquery/commit/7f5260c6fa74dee57697aa7a950099a645461c8a))
* Update endpoints ([#7797](https://github.com/cloudquery/cloudquery/issues/7797)) ([3e72d3e](https://github.com/cloudquery/cloudquery/commit/3e72d3e345ecd9101d82e3b2091150ca5c24f0a8))
* Update endpoints ([#7798](https://github.com/cloudquery/cloudquery/issues/7798)) ([0a21b16](https://github.com/cloudquery/cloudquery/commit/0a21b16176ced82323eee8b9e37b575e8082dc40))
* Update endpoints ([#7799](https://github.com/cloudquery/cloudquery/issues/7799)) ([ea62bbc](https://github.com/cloudquery/cloudquery/commit/ea62bbccf5acb4b46a45ea7bbdf97c83b4eadcd2))
* Update endpoints ([#7810](https://github.com/cloudquery/cloudquery/issues/7810)) ([c797331](https://github.com/cloudquery/cloudquery/commit/c7973312ab23f845aa9437dbf604fdc6bb315f4f))
* Update endpoints ([#7811](https://github.com/cloudquery/cloudquery/issues/7811)) ([5b1566c](https://github.com/cloudquery/cloudquery/commit/5b1566cd73094042c2fbf835a0a509957ca44e2d))
* Update endpoints ([#7816](https://github.com/cloudquery/cloudquery/issues/7816)) ([e1cca0c](https://github.com/cloudquery/cloudquery/commit/e1cca0c0ab9e08ecdbfdd114ed72df874d4d32dd))
* Update endpoints ([#7991](https://github.com/cloudquery/cloudquery/issues/7991)) ([3e81234](https://github.com/cloudquery/cloudquery/commit/3e81234bf633e6ede7741f71b95ad05cf8833dc9))

## [13.3.0](https://github.com/cloudquery/cloudquery/compare/plugins-source-aws-v13.2.0...plugins-source-aws-v13.3.0) (2023-02-07)


### Features

* **aws:** Add IAM Signing Certificate ([#7699](https://github.com/cloudquery/cloudquery/issues/7699)) ([35344f3](https://github.com/cloudquery/cloudquery/commit/35344f3ed590da7a5256c45de3e7d8c0d51a6229))
* **aws:** Add Support for S3 Accesspoints ([#7704](https://github.com/cloudquery/cloudquery/issues/7704)) ([0c792f9](https://github.com/cloudquery/cloudquery/commit/0c792f9caf74d9e39720783b3ffafb5eccd285db))

## [13.2.0](https://github.com/cloudquery/cloudquery/compare/plugins-source-aws-v13.1.0...plugins-source-aws-v13.2.0) (2023-02-07)


### Features

* Add Resilience Hub resources ([#7299](https://github.com/cloudquery/cloudquery/issues/7299)) ([1b54e7e](https://github.com/cloudquery/cloudquery/commit/1b54e7e9a910801617937cdbf74fa3be31a2b496))
* **aws:** Add support for Amplify Apps ([#7695](https://github.com/cloudquery/cloudquery/issues/7695)) ([8ee00c5](https://github.com/cloudquery/cloudquery/commit/8ee00c5d1077905d798bb70521a40cc6143e0e1d))
* **aws:** Add support for ECS Tasksets ([#7688](https://github.com/cloudquery/cloudquery/issues/7688)) ([b77e8b5](https://github.com/cloudquery/cloudquery/commit/b77e8b5f6e59b08d07c4467e32206a2174303c0d))
* **aws:** Add support for EKS Fargate Profiles ([#7693](https://github.com/cloudquery/cloudquery/issues/7693)) ([1f668db](https://github.com/cloudquery/cloudquery/commit/1f668db8e8e0ebe289131124c03cd2cee8a92165))
* **AWS:** Add support for EKS Node Groups ([#7692](https://github.com/cloudquery/cloudquery/issues/7692)) ([4a5650e](https://github.com/cloudquery/cloudquery/commit/4a5650e3136030438265dd8d11e3149a0652d81c))
* **aws:** Add Support for Secrets Manager Secret Versions ([#7701](https://github.com/cloudquery/cloudquery/issues/7701)) ([041c341](https://github.com/cloudquery/cloudquery/commit/041c3412b9afa20331c49dc6e12a65e3de6fe4cd))


### Bug Fixes

* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.34.0 ([#7719](https://github.com/cloudquery/cloudquery/issues/7719)) ([6a33085](https://github.com/cloudquery/cloudquery/commit/6a33085c75adcf2387f7bbb5aa4f7a84ce7e2957))
* Update endpoints ([#7702](https://github.com/cloudquery/cloudquery/issues/7702)) ([03d1eeb](https://github.com/cloudquery/cloudquery/commit/03d1eeb30cfb2c4657ce254751b4b8d4823b6bde))

## [13.1.0](https://github.com/cloudquery/cloudquery/compare/plugins-source-aws-v13.0.0...plugins-source-aws-v13.1.0) (2023-02-06)


### Features

* **aws-resources:** Add EC2 EBS Volume Statuses ([#7638](https://github.com/cloudquery/cloudquery/issues/7638)) ([a23c6a3](https://github.com/cloudquery/cloudquery/commit/a23c6a3725f827225c4a13de729dd629fdb5fd7d))


### Bug Fixes

* **aws:** Turn tags into maps ([#7678](https://github.com/cloudquery/cloudquery/issues/7678)) ([acaa654](https://github.com/cloudquery/cloudquery/commit/acaa6540e6f7187afc16e669b125cafa00bc93a8))
* **aws:** Update EC2 Tag structure ([#7621](https://github.com/cloudquery/cloudquery/issues/7621)) ([8d97115](https://github.com/cloudquery/cloudquery/commit/8d9711567ed8eca5e814a1e8840d25923f190bc6))
* **deps:** Update module github.com/aws/aws-sdk-go-v2 to v1.17.4 ([#7652](https://github.com/cloudquery/cloudquery/issues/7652)) ([2196050](https://github.com/cloudquery/cloudquery/commit/2196050848b7abdafa9174af97151d0dbdf629c4))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/config to v1.18.12 ([#7653](https://github.com/cloudquery/cloudquery/issues/7653)) ([59daf42](https://github.com/cloudquery/cloudquery/commit/59daf423f2992c89db3db542c000286800d4ca61))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/feature/ec2/imds to v1.12.22 ([#7655](https://github.com/cloudquery/cloudquery/issues/7655)) ([4e56621](https://github.com/cloudquery/cloudquery/commit/4e56621f73f515874c15eddb6da8b349d0889d6c))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/feature/s3/manager to v1.11.51 ([#7656](https://github.com/cloudquery/cloudquery/issues/7656)) ([43a0c59](https://github.com/cloudquery/cloudquery/commit/43a0c59ca701281fa558c7a73a7673e019ad3ad6))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/internal/ini to v1.3.29 ([#7659](https://github.com/cloudquery/cloudquery/issues/7659)) ([60f15d7](https://github.com/cloudquery/cloudquery/commit/60f15d7cadfb3323c9b072869e252cdc7dfb0aab))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/internal/v4a to v1.0.19 ([#7660](https://github.com/cloudquery/cloudquery/issues/7660)) ([9035012](https://github.com/cloudquery/cloudquery/commit/9035012d6ac2d41bdbdf0e2bf6f025f1bbac058b))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/accessanalyzer to v1.19.2 ([#7661](https://github.com/cloudquery/cloudquery/issues/7661)) ([dfbb566](https://github.com/cloudquery/cloudquery/commit/dfbb566ba2c3b633e0ec05c10bbc494a8ca8c405))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/account to v1.8.1 ([#7662](https://github.com/cloudquery/cloudquery/issues/7662)) ([ceebda7](https://github.com/cloudquery/cloudquery/commit/ceebda78dc0aece658c2b213654777366e432287))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/acm to v1.17.2 ([#7663](https://github.com/cloudquery/cloudquery/issues/7663)) ([0d29d75](https://github.com/cloudquery/cloudquery/commit/0d29d75efb06eaf0b4a63de2d1564b6ba34cc88d))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/amp to v1.16.1 ([#7664](https://github.com/cloudquery/cloudquery/issues/7664)) ([72e3613](https://github.com/cloudquery/cloudquery/commit/72e36138fe5f2ef4536926cee16de81f15beccfe))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/apigateway to v1.16.2 ([#7665](https://github.com/cloudquery/cloudquery/issues/7665)) ([a84fd09](https://github.com/cloudquery/cloudquery/commit/a84fd0940b0be6dd2521f5630eff5f8acaa1dc1b))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/apigatewayv2 to v1.13.2 ([#7666](https://github.com/cloudquery/cloudquery/issues/7666)) ([335fb4f](https://github.com/cloudquery/cloudquery/commit/335fb4f3103d78713cff840bf68be9c530ee544e))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/applicationautoscaling to v1.17.2 ([#7667](https://github.com/cloudquery/cloudquery/issues/7667)) ([be3e871](https://github.com/cloudquery/cloudquery/commit/be3e871aadb2c365e06f0b50289863c1c19cce51))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/apprunner to v1.16.1 ([#7668](https://github.com/cloudquery/cloudquery/issues/7668)) ([01aa892](https://github.com/cloudquery/cloudquery/commit/01aa89289524610ee75ff60537cd4d09e513bc34))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/athena to v1.22.1 ([#7669](https://github.com/cloudquery/cloudquery/issues/7669)) ([dbfdb6d](https://github.com/cloudquery/cloudquery/commit/dbfdb6dc4307ac8a588bc31264f3e39b1010b7af))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/autoscaling to v1.26.2 ([#7670](https://github.com/cloudquery/cloudquery/issues/7670)) ([8e1dc76](https://github.com/cloudquery/cloudquery/commit/8e1dc76eed2c69986c596b9526f32cc8da56ef0e))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/backup to v1.19.2 ([#7671](https://github.com/cloudquery/cloudquery/issues/7671)) ([36e1ea5](https://github.com/cloudquery/cloudquery/commit/36e1ea58cd40466e23448772a368e61de25861f9))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/cloudfront to v1.24.1 ([#7672](https://github.com/cloudquery/cloudquery/issues/7672)) ([9a246a2](https://github.com/cloudquery/cloudquery/commit/9a246a2fc07eac19f9b36606700b6bea65b8fc06))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/cloudhsmv2 to v1.14.1 ([#7673](https://github.com/cloudquery/cloudquery/issues/7673)) ([63c0e33](https://github.com/cloudquery/cloudquery/commit/63c0e3399d8f3fdcc37f11262fb28d2337d50ddb))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/cloudwatch to v1.25.2 ([#7674](https://github.com/cloudquery/cloudquery/issues/7674)) ([88d2b93](https://github.com/cloudquery/cloudquery/commit/88d2b937f092f8339d48ecd3411c11e9b5b9896c))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs to v1.20.2 ([#7675](https://github.com/cloudquery/cloudquery/issues/7675)) ([4b04056](https://github.com/cloudquery/cloudquery/commit/4b040567fa048381f26fa19a9c52910035904c0e))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/codebuild to v1.20.2 ([#7676](https://github.com/cloudquery/cloudquery/issues/7676)) ([c0d4e3c](https://github.com/cloudquery/cloudquery/commit/c0d4e3c701f214b7f82926aedb2083cf03685377))
* Update endpoints ([#7626](https://github.com/cloudquery/cloudquery/issues/7626)) ([91cf8ad](https://github.com/cloudquery/cloudquery/commit/91cf8ad6367c0834faf04520a842e7b13ea14cfe))
* Update endpoints ([#7641](https://github.com/cloudquery/cloudquery/issues/7641)) ([3ceebc6](https://github.com/cloudquery/cloudquery/commit/3ceebc69136fa8752a52ffcdd0f42e5cd0142592))
* Update endpoints ([#7648](https://github.com/cloudquery/cloudquery/issues/7648)) ([dd8c4fe](https://github.com/cloudquery/cloudquery/commit/dd8c4fe56a749c288a7f83e80d125b19087ff626))

## [13.0.0](https://github.com/cloudquery/cloudquery/compare/plugins-source-aws-v12.2.1...plugins-source-aws-v13.0.0) (2023-02-01)


### ⚠ BREAKING CHANGES

* **aws:** PK for `aws_eventbridge_event_bus_rules` ([#7394](https://github.com/cloudquery/cloudquery/issues/7394))
* **aws:** PKs for `aws_apigatewayv2_*` ([#7380](https://github.com/cloudquery/cloudquery/issues/7380))
* **aws:** PK for `aws_cloudtrail_trails` ([#7467](https://github.com/cloudquery/cloudquery/issues/7467))
* **aws:** PK for `aws_route53_hosted_zone_traffic_policy_instances` ([#7412](https://github.com/cloudquery/cloudquery/issues/7412))
* **aws:** PK for `aws_qldb_ledger_journal_kinesis_streams` ([#7409](https://github.com/cloudquery/cloudquery/issues/7409))
* **aws:** PK for `aws_mq_broker_configurations` ([#7407](https://github.com/cloudquery/cloudquery/issues/7407))
* **aws:** PKs for `aws_lightsail_` tables ([#7401](https://github.com/cloudquery/cloudquery/issues/7401))
* **aws:** PK for `aws_lambda_layer_versions` ([#7399](https://github.com/cloudquery/cloudquery/issues/7399))
* **aws:** PK for `aws_iam_user_groups` ([#7397](https://github.com/cloudquery/cloudquery/issues/7397))
* **aws:** PK for `aws_glue_registry_schemas` ([#7395](https://github.com/cloudquery/cloudquery/issues/7395))
* **aws:** PK for `aws_apigateway_vpc_links` ([#7379](https://github.com/cloudquery/cloudquery/issues/7379))
* **aws:** PK for `aws_apigateway_usage_plan_keys` ([#7376](https://github.com/cloudquery/cloudquery/issues/7376))
* **aws:** PK for `aws_apigateway_usage_plans` ([#7375](https://github.com/cloudquery/cloudquery/issues/7375))
* **aws:** PK for `aws_apigateway_rest_api_stages` ([#7373](https://github.com/cloudquery/cloudquery/issues/7373))
* **aws:** PK for `aws_apigateway_rest_api_resources` ([#7369](https://github.com/cloudquery/cloudquery/issues/7369))
* **aws:** PK for `aws_apigateway_rest_api_request_validators` ([#7368](https://github.com/cloudquery/cloudquery/issues/7368))
* **aws:** PK for `aws_apigateway_rest_api_models` ([#7366](https://github.com/cloudquery/cloudquery/issues/7366))
* **aws:** PK for `aws_apigateway_rest_api_gateway_responses` ([#7364](https://github.com/cloudquery/cloudquery/issues/7364))
* **aws:** PK for `aws_apigateway_rest_api_documentation_versions` ([#7360](https://github.com/cloudquery/cloudquery/issues/7360))
* **aws:** PK for `aws_apigateway_rest_api_deployments` ([#7356](https://github.com/cloudquery/cloudquery/issues/7356))
* **aws:** PK for `aws_apigateway_rest_api_documentation_parts` ([#7355](https://github.com/cloudquery/cloudquery/issues/7355))
* **aws:** PK for `aws_apigateway_rest_api_authorizers` ([#7350](https://github.com/cloudquery/cloudquery/issues/7350))
* **aws:** PK for `aws_apigateway_domain_name_base_path_mappings`
* **aws:** RDS Engine Version PKs ([#7202](https://github.com/cloudquery/cloudquery/issues/7202))

### Features

* **aws:** Add Support for Securityhub findings ([#7204](https://github.com/cloudquery/cloudquery/issues/7204)) ([d96496a](https://github.com/cloudquery/cloudquery/commit/d96496a8b78b23fbcba48bab408d43f4f2a4304c))


### Bug Fixes

* **aws:** PK for `aws_apigateway_domain_name_base_path_mappings` ([8a945be](https://github.com/cloudquery/cloudquery/commit/8a945be178b94a74fdeb215adabb3ff859f409b4))
* **aws:** PK for `aws_apigateway_rest_api_authorizers` ([#7350](https://github.com/cloudquery/cloudquery/issues/7350)) ([33e110f](https://github.com/cloudquery/cloudquery/commit/33e110fe88061d17f83bbcb48800e789318c98a7))
* **aws:** PK for `aws_apigateway_rest_api_deployments` ([#7356](https://github.com/cloudquery/cloudquery/issues/7356)) ([d5f6fd5](https://github.com/cloudquery/cloudquery/commit/d5f6fd518cb4768c205abfe16ef067eae43dce5b))
* **aws:** PK for `aws_apigateway_rest_api_documentation_parts` ([#7355](https://github.com/cloudquery/cloudquery/issues/7355)) ([2d81f86](https://github.com/cloudquery/cloudquery/commit/2d81f86aae5a891cca167495a90a4d17d4d229f5))
* **aws:** PK for `aws_apigateway_rest_api_documentation_versions` ([#7360](https://github.com/cloudquery/cloudquery/issues/7360)) ([d6a5aea](https://github.com/cloudquery/cloudquery/commit/d6a5aead9ef2dc3a517c81c6d294b5f67cf96584))
* **aws:** PK for `aws_apigateway_rest_api_gateway_responses` ([#7364](https://github.com/cloudquery/cloudquery/issues/7364)) ([6e779ac](https://github.com/cloudquery/cloudquery/commit/6e779ac8338f24c3eeff417a53cc63ad12f82416))
* **aws:** PK for `aws_apigateway_rest_api_models` ([#7366](https://github.com/cloudquery/cloudquery/issues/7366)) ([aad7dd6](https://github.com/cloudquery/cloudquery/commit/aad7dd66ae22c56319655a7575858b36a7086cd2))
* **aws:** PK for `aws_apigateway_rest_api_request_validators` ([#7368](https://github.com/cloudquery/cloudquery/issues/7368)) ([0394e31](https://github.com/cloudquery/cloudquery/commit/0394e31b3a461b316abe6ace8f4bcb688577e3aa))
* **aws:** PK for `aws_apigateway_rest_api_resources` ([#7369](https://github.com/cloudquery/cloudquery/issues/7369)) ([0c8fa30](https://github.com/cloudquery/cloudquery/commit/0c8fa3057c0d2699fcc002cb4ea84a5d43f63151))
* **aws:** PK for `aws_apigateway_rest_api_stages` ([#7373](https://github.com/cloudquery/cloudquery/issues/7373)) ([b49ee55](https://github.com/cloudquery/cloudquery/commit/b49ee55c425ad2801483ff1a01def52e8477091b))
* **aws:** PK for `aws_apigateway_usage_plan_keys` ([#7376](https://github.com/cloudquery/cloudquery/issues/7376)) ([4d7a76f](https://github.com/cloudquery/cloudquery/commit/4d7a76f36804df040bb0a49eed0979792e728b0d))
* **aws:** PK for `aws_apigateway_usage_plans` ([#7375](https://github.com/cloudquery/cloudquery/issues/7375)) ([bac329d](https://github.com/cloudquery/cloudquery/commit/bac329db0573516ecf0c43cc0b6c145654da2169))
* **aws:** PK for `aws_apigateway_vpc_links` ([#7379](https://github.com/cloudquery/cloudquery/issues/7379)) ([60bd130](https://github.com/cloudquery/cloudquery/commit/60bd13095a3894f22234fcfd926ec64804b76b37))
* **aws:** PK for `aws_cloudtrail_trails` ([#7467](https://github.com/cloudquery/cloudquery/issues/7467)) ([f580207](https://github.com/cloudquery/cloudquery/commit/f580207491ad0dd6e6d3e51109c9fec11751a590))
* **aws:** PK for `aws_eventbridge_event_bus_rules` ([#7394](https://github.com/cloudquery/cloudquery/issues/7394)) ([e1e7405](https://github.com/cloudquery/cloudquery/commit/e1e7405035dce99a7770a0a0cf434b6cdc39454f))
* **aws:** PK for `aws_glue_registry_schemas` ([#7395](https://github.com/cloudquery/cloudquery/issues/7395)) ([de43500](https://github.com/cloudquery/cloudquery/commit/de4350070bb724db9680aca132306d9fad16e7c1))
* **aws:** PK for `aws_iam_user_groups` ([#7397](https://github.com/cloudquery/cloudquery/issues/7397)) ([ced05f6](https://github.com/cloudquery/cloudquery/commit/ced05f6ea4f0f0da6d0a7bfbcad670b646565559))
* **aws:** PK for `aws_lambda_layer_versions` ([#7399](https://github.com/cloudquery/cloudquery/issues/7399)) ([744c124](https://github.com/cloudquery/cloudquery/commit/744c124b90c5ba82b4ba5b7a7080053a65a38bb6))
* **aws:** PK for `aws_mq_broker_configurations` ([#7407](https://github.com/cloudquery/cloudquery/issues/7407)) ([3b6f383](https://github.com/cloudquery/cloudquery/commit/3b6f383d55cac085be447d1b84829e4e065d620b))
* **aws:** PK for `aws_qldb_ledger_journal_kinesis_streams` ([#7409](https://github.com/cloudquery/cloudquery/issues/7409)) ([101cdb6](https://github.com/cloudquery/cloudquery/commit/101cdb65d2f8ae67990ea45fa01584a439035820))
* **aws:** PK for `aws_route53_hosted_zone_traffic_policy_instances` ([#7412](https://github.com/cloudquery/cloudquery/issues/7412)) ([4bf36e7](https://github.com/cloudquery/cloudquery/commit/4bf36e72dbc2579add6c4feb2c8a6ab12a2bafa5))
* **aws:** PKs for `aws_apigatewayv2_*` ([#7380](https://github.com/cloudquery/cloudquery/issues/7380)) ([48f5463](https://github.com/cloudquery/cloudquery/commit/48f5463f6e8452930b55a866218d404b63e1818b))
* **aws:** PKs for `aws_lightsail_` tables ([#7401](https://github.com/cloudquery/cloudquery/issues/7401)) ([0b218b0](https://github.com/cloudquery/cloudquery/commit/0b218b08cfa7592276eb05a882ee1d78a6faccec))
* **aws:** RDS Engine Version PKs ([#7202](https://github.com/cloudquery/cloudquery/issues/7202)) ([a49984e](https://github.com/cloudquery/cloudquery/commit/a49984e381e23e7d3e419e9dbf3bf3e652be3e0b))
* **deps:** Update github.com/gocarina/gocsv digest to 763e25b ([#7529](https://github.com/cloudquery/cloudquery/issues/7529)) ([9aaa696](https://github.com/cloudquery/cloudquery/commit/9aaa696877cb2657bd6a26579a5b33594b11b829))
* **deps:** Update golang.org/x/exp digest to f062dba ([#7531](https://github.com/cloudquery/cloudquery/issues/7531)) ([59d5575](https://github.com/cloudquery/cloudquery/commit/59d55758b0951553b8d246d1e78b4e3917ff1976))
* **deps:** Update google.golang.org/genproto digest to 1c01626 ([#7533](https://github.com/cloudquery/cloudquery/issues/7533)) ([c549c27](https://github.com/cloudquery/cloudquery/commit/c549c275077f1cdfb9df0b3f3c129cbf0b150552))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/feature/s3/manager to v1.11.49 ([#7309](https://github.com/cloudquery/cloudquery/issues/7309)) ([16da39d](https://github.com/cloudquery/cloudquery/commit/16da39d4bd8a6851329cbd25c2d80801b1872663))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/accessanalyzer to v1.19.1 ([#7310](https://github.com/cloudquery/cloudquery/issues/7310)) ([3bb3d78](https://github.com/cloudquery/cloudquery/commit/3bb3d78e0cb0bccd2ada91faa033cf330355fe14))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/acm to v1.17.1 ([#7311](https://github.com/cloudquery/cloudquery/issues/7311)) ([829e2c0](https://github.com/cloudquery/cloudquery/commit/829e2c00f5f0609f5d2144bb4ddef4d707cf6c5e))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/apigateway to v1.16.1 ([#7312](https://github.com/cloudquery/cloudquery/issues/7312)) ([1fd914d](https://github.com/cloudquery/cloudquery/commit/1fd914d6115006d5e40ed6529ccf4b57157b6360))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/apigatewayv2 to v1.13.1 ([#7313](https://github.com/cloudquery/cloudquery/issues/7313)) ([b454a62](https://github.com/cloudquery/cloudquery/commit/b454a621138d8380b6b0da0e687ad8c5e01c1050))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/applicationautoscaling to v1.17.1 ([#7314](https://github.com/cloudquery/cloudquery/issues/7314)) ([b6a0807](https://github.com/cloudquery/cloudquery/commit/b6a08073e398bdb0398c8fb698bc6f18b3ae41ab))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/autoscaling to v1.26.1 ([#7315](https://github.com/cloudquery/cloudquery/issues/7315)) ([956fddc](https://github.com/cloudquery/cloudquery/commit/956fddc71193c9bbfb1f7d5601231d7c68cc23ab))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/cloudformation to v1.25.1 ([#7316](https://github.com/cloudquery/cloudquery/issues/7316)) ([de1cb2d](https://github.com/cloudquery/cloudquery/commit/de1cb2d2697224b6dfdbcf84d3ced6fa603d0c27))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/cloudtrail to v1.22.1 ([#7317](https://github.com/cloudquery/cloudquery/issues/7317)) ([3a454c4](https://github.com/cloudquery/cloudquery/commit/3a454c4b3e54bd233c73da66891443bd409c8f8a))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/cloudwatch to v1.25.1 ([#7318](https://github.com/cloudquery/cloudquery/issues/7318)) ([4103fb9](https://github.com/cloudquery/cloudquery/commit/4103fb960e1139d38e6b11ea0ba2666ee773251c))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs to v1.20.1 ([#7319](https://github.com/cloudquery/cloudquery/issues/7319)) ([12cdbdb](https://github.com/cloudquery/cloudquery/commit/12cdbdb25e9200a9c5b5e415550031fab7786d4e))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/configservice to v1.29.1 ([#7320](https://github.com/cloudquery/cloudquery/issues/7320)) ([3535303](https://github.com/cloudquery/cloudquery/commit/3535303b8e2918bf9a739b6471ced7254ee9f3f6))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/databasemigrationservice to v1.23.1 ([#7321](https://github.com/cloudquery/cloudquery/issues/7321)) ([4c3f0ae](https://github.com/cloudquery/cloudquery/commit/4c3f0ae36592586f81cc0f2bfc4b06bc593e0827))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/directconnect to v1.18.1 ([#7322](https://github.com/cloudquery/cloudquery/issues/7322)) ([0ea2664](https://github.com/cloudquery/cloudquery/commit/0ea2664f14eba40d0a111e36e833079fb88d843e))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/docdb to v1.20.1 ([#7323](https://github.com/cloudquery/cloudquery/issues/7323)) ([03ca605](https://github.com/cloudquery/cloudquery/commit/03ca6051fd1e3ff2cc07897a0cd3411c798cdf79))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/dynamodb to v1.18.1 ([#7324](https://github.com/cloudquery/cloudquery/issues/7324)) ([6097895](https://github.com/cloudquery/cloudquery/commit/60978956c6a414b6e6e03157b2f9693536d4a65d))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/ecr to v1.18.1 ([#7325](https://github.com/cloudquery/cloudquery/issues/7325)) ([66cd4f7](https://github.com/cloudquery/cloudquery/commit/66cd4f7441296bb8d82ed18b1e48fb93242714f0))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/ecs to v1.23.1 ([#7326](https://github.com/cloudquery/cloudquery/issues/7326)) ([7794001](https://github.com/cloudquery/cloudquery/commit/7794001930ac968fa7c67bb62a16824102359762))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/elasticache to v1.26.1 ([#7327](https://github.com/cloudquery/cloudquery/issues/7327)) ([4c203a5](https://github.com/cloudquery/cloudquery/commit/4c203a5237b0047001cc4456512cf31f3008726b))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/elasticloadbalancing to v1.15.1 ([#7328](https://github.com/cloudquery/cloudquery/issues/7328)) ([7393b13](https://github.com/cloudquery/cloudquery/commit/7393b132576e087010dab20048bd517646f2e8d7))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2 to v1.19.1 ([#7329](https://github.com/cloudquery/cloudquery/issues/7329)) ([d288551](https://github.com/cloudquery/cloudquery/commit/d288551fdcc079356a4d9c051db41746a5d552b6))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/elasticsearchservice to v1.18.1 ([#7330](https://github.com/cloudquery/cloudquery/issues/7330)) ([5e8e42c](https://github.com/cloudquery/cloudquery/commit/5e8e42cb4e2ac9eb9f5665fbab53e0dfcf30a3ad))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/emr to v1.22.1 ([#7331](https://github.com/cloudquery/cloudquery/issues/7331)) ([d914db7](https://github.com/cloudquery/cloudquery/commit/d914db792548a118b4b2d1c9b4cf4a12c29ecc2e))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.32.0 ([#7334](https://github.com/cloudquery/cloudquery/issues/7334)) ([b684122](https://github.com/cloudquery/cloudquery/commit/b68412222219f9ca160c0753290709d52de7fcd6))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.33.0 ([#7595](https://github.com/cloudquery/cloudquery/issues/7595)) ([c5adc75](https://github.com/cloudquery/cloudquery/commit/c5adc750d4b0242563997c04c582f8da27913095))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.33.1 ([#7614](https://github.com/cloudquery/cloudquery/issues/7614)) ([2fe665c](https://github.com/cloudquery/cloudquery/commit/2fe665cdd80d88c5699bb203bd7accd604dfba99))
* Update endpoints ([#7521](https://github.com/cloudquery/cloudquery/issues/7521)) ([004d433](https://github.com/cloudquery/cloudquery/commit/004d433b82b03e4b69eb33233bfc693f1ef2ad36))

## [12.2.1](https://github.com/cloudquery/cloudquery/compare/plugins-source-aws-v12.2.0...plugins-source-aws-v12.2.1) (2023-01-27)


### Bug Fixes

* Cloudfront Policy PK ([#7294](https://github.com/cloudquery/cloudquery/issues/7294)) ([7949fff](https://github.com/cloudquery/cloudquery/commit/7949fff3be36663e36e8b6020c0c1639b44d9aa4))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.31.0 ([#7228](https://github.com/cloudquery/cloudquery/issues/7228)) ([36e8549](https://github.com/cloudquery/cloudquery/commit/36e8549f722658d909865723630fad1b2821db62))
* Update endpoints ([#7232](https://github.com/cloudquery/cloudquery/issues/7232)) ([8b7a8d0](https://github.com/cloudquery/cloudquery/commit/8b7a8d0cad158276e8eddc7bf02c408c15d6fc46))

## [12.2.0](https://github.com/cloudquery/cloudquery/compare/plugins-source-aws-v12.1.0...plugins-source-aws-v12.2.0) (2023-01-26)


### Features

* **aws:** Add support for Lambda Runtime Management ([#7152](https://github.com/cloudquery/cloudquery/issues/7152)) ([888e6e2](https://github.com/cloudquery/cloudquery/commit/888e6e201996f512ee277bab81fec2e60b51331c))


### Bug Fixes

* **aws:** Built in Resolvers ([#7203](https://github.com/cloudquery/cloudquery/issues/7203)) ([6b32744](https://github.com/cloudquery/cloudquery/commit/6b32744654b22122b9b1225683122509c945cf9b))
* **aws:** Cloudtrail pks ([#7199](https://github.com/cloudquery/cloudquery/issues/7199)) ([06d8ff6](https://github.com/cloudquery/cloudquery/commit/06d8ff60dda19c6c7c325bd5e7cf10b3121bd524))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.29.0 ([#7121](https://github.com/cloudquery/cloudquery/issues/7121)) ([b7441c9](https://github.com/cloudquery/cloudquery/commit/b7441c93c274ae3a6009474a2b28f44a172dd6dc))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.30.0 ([#7222](https://github.com/cloudquery/cloudquery/issues/7222)) ([73ca21c](https://github.com/cloudquery/cloudquery/commit/73ca21c4259545f7e949c9d780d8184db475d2ac))
* Update endpoints ([#7142](https://github.com/cloudquery/cloudquery/issues/7142)) ([7116865](https://github.com/cloudquery/cloudquery/commit/7116865aa5eebfeec864acd33c983cd72dbe355b))
* Update endpoints ([#7198](https://github.com/cloudquery/cloudquery/issues/7198)) ([506392f](https://github.com/cloudquery/cloudquery/commit/506392f8a125ab7ffcfdd577d2f8333d6394973e))

## [12.1.0](https://github.com/cloudquery/cloudquery/compare/plugins-source-aws-v12.0.0...plugins-source-aws-v12.1.0) (2023-01-24)


### Features

* **aws:** Add aws/iam/instance_profiles table ([#6985](https://github.com/cloudquery/cloudquery/issues/6985)) ([a7e3cd6](https://github.com/cloudquery/cloudquery/commit/a7e3cd6991da006bcb6bd2dda9bc7880f1e4d842))


### Bug Fixes

* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/backup to v1.19.1 ([#6990](https://github.com/cloudquery/cloudquery/issues/6990)) ([0bdda17](https://github.com/cloudquery/cloudquery/commit/0bdda172bb5bbf3d3b7147cc6e662dc471edd6c3))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/cloudwatch to v1.25.0 ([#6995](https://github.com/cloudquery/cloudquery/issues/6995)) ([8b1ffde](https://github.com/cloudquery/cloudquery/commit/8b1ffdea7fef3b39f8f1e67cf0207850a813f500))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs to v1.20.0 ([#6996](https://github.com/cloudquery/cloudquery/issues/6996)) ([40ff6b2](https://github.com/cloudquery/cloudquery/commit/40ff6b263b4ad0e96b00495731ed8f4247f87b52))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/codebuild to v1.20.1 ([#6991](https://github.com/cloudquery/cloudquery/issues/6991)) ([060797f](https://github.com/cloudquery/cloudquery/commit/060797f9e778199de3bf6639a24238eeef4f5681))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/ec2 to v1.80.0 ([#6997](https://github.com/cloudquery/cloudquery/issues/6997)) ([b669149](https://github.com/cloudquery/cloudquery/commit/b6691495c341f092f022007dc9ea7f677c450526))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/efs to v1.19.2 ([#6992](https://github.com/cloudquery/cloudquery/issues/6992)) ([0da6e11](https://github.com/cloudquery/cloudquery/commit/0da6e11df93b813944302a9f5a6768d3b3db2a93))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/glue to v1.40.0 ([#6998](https://github.com/cloudquery/cloudquery/issues/6998)) ([96b4b5d](https://github.com/cloudquery/cloudquery/commit/96b4b5dfb1711200133326a05e1361759b8ff12c))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/quicksight to v1.30.0 ([#6999](https://github.com/cloudquery/cloudquery/issues/6999)) ([655f478](https://github.com/cloudquery/cloudquery/commit/655f478c5cbb8681f76fb1d115374741b2762394))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/sagemaker to v1.63.0 ([#7000](https://github.com/cloudquery/cloudquery/issues/7000)) ([2206b56](https://github.com/cloudquery/cloudquery/commit/2206b567e1f894364d1e7c443bfe1a46779cbdcb))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/servicecatalogappregistry to v1.16.1 ([#6993](https://github.com/cloudquery/cloudquery/issues/6993)) ([b73997a](https://github.com/cloudquery/cloudquery/commit/b73997adcdcdb6a8663db5cdac3be65cb842f7c1))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/wafv2 to v1.24.2 ([#6994](https://github.com/cloudquery/cloudquery/issues/6994)) ([f48384c](https://github.com/cloudquery/cloudquery/commit/f48384c31796d4f612315643a4b013e77c136027))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.28.0 ([#7009](https://github.com/cloudquery/cloudquery/issues/7009)) ([12ac005](https://github.com/cloudquery/cloudquery/commit/12ac005428a355d06a5939fbe06a82d49533e662))
* Update endpoints ([#6953](https://github.com/cloudquery/cloudquery/issues/6953)) ([d894388](https://github.com/cloudquery/cloudquery/commit/d8943882602869b3f93e4b709bc2d1654543612f))
* Update endpoints ([#6980](https://github.com/cloudquery/cloudquery/issues/6980)) ([24b0219](https://github.com/cloudquery/cloudquery/commit/24b0219ca14628da3f3188efb463915064bf9a0d))
* Update endpoints ([#6983](https://github.com/cloudquery/cloudquery/issues/6983)) ([59ae5e2](https://github.com/cloudquery/cloudquery/commit/59ae5e29b7c4cd8a8ae5f81ad62bf823396ce39e))
* Update endpoints ([#7054](https://github.com/cloudquery/cloudquery/issues/7054)) ([d039572](https://github.com/cloudquery/cloudquery/commit/d039572fe6bdc4f96d46c94190f09a6434744fb6))

## [12.0.0](https://github.com/cloudquery/cloudquery/compare/plugins-source-aws-v11.0.1...plugins-source-aws-v12.0.0) (2023-01-18)


### ⚠ BREAKING CHANGES

* **aws:** Add region to `aws_backup_global_settings` primary key ([#6844](https://github.com/cloudquery/cloudquery/issues/6844))
* **aws:** Remove redundant fields from `aws_lambda_runtimes` ([#6849](https://github.com/cloudquery/cloudquery/issues/6849))
* **aws:** Add `region` to `aws_docdb_engine_versions` primary key ([#6846](https://github.com/cloudquery/cloudquery/issues/6846))

### Features

* **aws:** Add Elasticbeanstalk Application Tags ([#6915](https://github.com/cloudquery/cloudquery/issues/6915)) ([fd66b78](https://github.com/cloudquery/cloudquery/commit/fd66b78738d0831b0979866b2034da539505d883))
* **aws:** Add missing descriptions ([#6847](https://github.com/cloudquery/cloudquery/issues/6847)) ([d62b50b](https://github.com/cloudquery/cloudquery/commit/d62b50b57e5100ffa2adcf2d79b958501139d579))
* **aws:** Add RDS DB Proxies ([#6831](https://github.com/cloudquery/cloudquery/issues/6831)) ([8233160](https://github.com/cloudquery/cloudquery/commit/823316000a922df54ad4e713979267f0e8e4d9ab))
* **aws:** Remove more codegen ([#6853](https://github.com/cloudquery/cloudquery/issues/6853)) ([795e40d](https://github.com/cloudquery/cloudquery/commit/795e40d05ef7cb8f5dbe7b5fa94cb2c02e0f1c16))


### Bug Fixes

* **aws:** Add `db_engine_version_description` to `aws_rds_engine_versions` primary key ([#6851](https://github.com/cloudquery/cloudquery/issues/6851)) ([779be98](https://github.com/cloudquery/cloudquery/commit/779be9866beea65b95961119229f922bde4e9c97))
* **aws:** Add `name` and `vendor_name` to `aws_wafv2_managed_rule_groups` primary key ([#6843](https://github.com/cloudquery/cloudquery/issues/6843)) ([d555ed5](https://github.com/cloudquery/cloudquery/commit/d555ed5ad36ca2dd3b04437e35bc65420d1b7aa0))
* **aws:** Add `region` to `aws_docdb_engine_versions` primary key ([#6846](https://github.com/cloudquery/cloudquery/issues/6846)) ([0f624ab](https://github.com/cloudquery/cloudquery/commit/0f624abaeb0ff4ddcf33b5ad92a7f577c2ab0c19))
* **aws:** Add `region` to `aws_glue_datacatalog_encryption_settings` primary key ([#6845](https://github.com/cloudquery/cloudquery/issues/6845)) ([9b8b5b1](https://github.com/cloudquery/cloudquery/commit/9b8b5b1de07e7bb824b74e72d8c3b97794ebc4fd))
* **aws:** Add `region` to `aws_ram_resource_types` primary key ([#6850](https://github.com/cloudquery/cloudquery/issues/6850)) ([99eddd6](https://github.com/cloudquery/cloudquery/commit/99eddd64636a685fe33fba395a81cea9312e929b))
* **aws:** Add region to `aws_backup_global_settings` primary key ([#6844](https://github.com/cloudquery/cloudquery/issues/6844)) ([397e02a](https://github.com/cloudquery/cloudquery/commit/397e02ad7d3dd1f0786fe17870e7b780f1c5dd8a))
* **aws:** Remove redundant fields from `aws_lambda_runtimes` ([#6849](https://github.com/cloudquery/cloudquery/issues/6849)) ([cc5af56](https://github.com/cloudquery/cloudquery/commit/cc5af567cfab772ff5d7d7582a7c9604cb49fd37))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs to v1.19.0 ([#6813](https://github.com/cloudquery/cloudquery/issues/6813)) ([c623e0a](https://github.com/cloudquery/cloudquery/commit/c623e0a7272963b44878a7cc60b8fe8721e2bdfb))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/ec2 to v1.78.0 ([#6814](https://github.com/cloudquery/cloudquery/issues/6814)) ([902bee0](https://github.com/cloudquery/cloudquery/commit/902bee073c1affc5588f39a4380d5a2fbb9f0440))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/ecrpublic to v1.15.0 ([#6815](https://github.com/cloudquery/cloudquery/issues/6815)) ([df72894](https://github.com/cloudquery/cloudquery/commit/df7289443715b19b4ff2efb7e203009e0fddbcc9))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/lambda to v1.28.0 ([#6816](https://github.com/cloudquery/cloudquery/issues/6816)) ([56fe84f](https://github.com/cloudquery/cloudquery/commit/56fe84f72e30dbe51ded739577c161321a633bb0))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/rds to v1.40.0 ([#6817](https://github.com/cloudquery/cloudquery/issues/6817)) ([6227bef](https://github.com/cloudquery/cloudquery/commit/6227befc5b237c6a16dacb5f9461616dee3f8676))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/resourcegroups to v1.14.0 ([#6818](https://github.com/cloudquery/cloudquery/issues/6818)) ([c343395](https://github.com/cloudquery/cloudquery/commit/c343395247cd65705f4d5fda129c75a3942f6841))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/secretsmanager to v1.18.1 ([#6812](https://github.com/cloudquery/cloudquery/issues/6812)) ([99df2d9](https://github.com/cloudquery/cloudquery/commit/99df2d92cb8622132820e9c7d7ce68ba602c908a))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/shield to v1.18.0 ([#6819](https://github.com/cloudquery/cloudquery/issues/6819)) ([0f18621](https://github.com/cloudquery/cloudquery/commit/0f18621d869af0663d26455a058f2ed71b394b62))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/sns to v1.19.0 ([#6820](https://github.com/cloudquery/cloudquery/issues/6820)) ([95bbafd](https://github.com/cloudquery/cloudquery/commit/95bbafd40859c1b8484dc1a54d9da2f486ebe8b0))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/sqs to v1.20.0 ([#6821](https://github.com/cloudquery/cloudquery/issues/6821)) ([02974bc](https://github.com/cloudquery/cloudquery/commit/02974bc91797b54b28b0cd29ed12894feb6b0161))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/ssm to v1.35.0 ([#6822](https://github.com/cloudquery/cloudquery/issues/6822)) ([82250e0](https://github.com/cloudquery/cloudquery/commit/82250e0ad236e1b4b3395b1cf718ca55f8ee5d85))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/ssoadmin to v1.16.0 ([#6823](https://github.com/cloudquery/cloudquery/issues/6823)) ([75412e5](https://github.com/cloudquery/cloudquery/commit/75412e5a8ab598731dadd527a1782630e3403fd6))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/timestreamwrite to v1.15.0 ([#6824](https://github.com/cloudquery/cloudquery/issues/6824)) ([8912b57](https://github.com/cloudquery/cloudquery/commit/8912b579dbf08c66b4395b7a133ba973d4899bb7))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/transfer to v1.28.0 ([#6825](https://github.com/cloudquery/cloudquery/issues/6825)) ([18f6ad7](https://github.com/cloudquery/cloudquery/commit/18f6ad7a0e15b5c8592882b7ccb00bd4d087a84d))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/waf to v1.12.0 ([#6826](https://github.com/cloudquery/cloudquery/issues/6826)) ([0c8c0d3](https://github.com/cloudquery/cloudquery/commit/0c8c0d3ad12a625b2ebed08a1d90a8855883112a))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/wafregional to v1.13.1 ([#6827](https://github.com/cloudquery/cloudquery/issues/6827)) ([d3a03fc](https://github.com/cloudquery/cloudquery/commit/d3a03fc1902d2f942d34f22c8febd4b8ee53ba6b))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/wafv2 to v1.24.1 ([#6828](https://github.com/cloudquery/cloudquery/issues/6828)) ([d24c78e](https://github.com/cloudquery/cloudquery/commit/d24c78e75e100cd5364cc01a3e3f9257870f6a7a))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/workspaces to v1.28.0 ([#6829](https://github.com/cloudquery/cloudquery/issues/6829)) ([8ce0af6](https://github.com/cloudquery/cloudquery/commit/8ce0af610a2e385ada5f42223d3620ab0d620707))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/xray to v1.16.0 ([#6830](https://github.com/cloudquery/cloudquery/issues/6830)) ([0d88a34](https://github.com/cloudquery/cloudquery/commit/0d88a342167ed5ba8829b4eaace1b19fcf3a921a))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.25.1 ([#6805](https://github.com/cloudquery/cloudquery/issues/6805)) ([9da0ce2](https://github.com/cloudquery/cloudquery/commit/9da0ce283f50410eb9274375ec1d22131a80d937))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.26.0 ([#6839](https://github.com/cloudquery/cloudquery/issues/6839)) ([6ccda8d](https://github.com/cloudquery/cloudquery/commit/6ccda8d0bc6e7ce75f4a64a18911e349ccaac277))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.27.0 ([#6856](https://github.com/cloudquery/cloudquery/issues/6856)) ([545799b](https://github.com/cloudquery/cloudquery/commit/545799bb0481087e187b5f27c88f5dde9c99f2f0))
* Update endpoints ([#6798](https://github.com/cloudquery/cloudquery/issues/6798)) ([041040d](https://github.com/cloudquery/cloudquery/commit/041040d2b7674b7b4ffd524bb654b69d08964bf7))
* Update endpoints ([#6912](https://github.com/cloudquery/cloudquery/issues/6912)) ([bb77355](https://github.com/cloudquery/cloudquery/commit/bb77355f7d5ab1e9d661d6b99e694c873601f923))

## [11.0.1](https://github.com/cloudquery/cloudquery/compare/plugins-source-aws-v11.0.0...plugins-source-aws-v11.0.1) (2023-01-12)


### Bug Fixes

* **aws:** Correctly fill in `grantee_id` column to `aws_s3_bucket_grants` ([#6772](https://github.com/cloudquery/cloudquery/issues/6772)) ([2cf0451](https://github.com/cloudquery/cloudquery/commit/2cf0451476faaaa014fa723be0e9732cb51d21da))
* Update endpoints ([#6774](https://github.com/cloudquery/cloudquery/issues/6774)) ([0523a1a](https://github.com/cloudquery/cloudquery/commit/0523a1a0e2b7ccd01a5222f58bbc484cf6968ca4))

## [11.0.0](https://github.com/cloudquery/cloudquery/compare/plugins-source-aws-v10.1.0...plugins-source-aws-v11.0.0) (2023-01-12)


### ⚠ BREAKING CHANGES

* introduce `aws_s3_bucket_grants` explicit primary key `(bucket_arn, grantee_id)`. If you've previously synced this table you'll need to drop it for the PK change (migration) to succeed.

### Features

* **aws:** Add Org resource policies ([#6743](https://github.com/cloudquery/cloudquery/issues/6743)) ([80dcf8e](https://github.com/cloudquery/cloudquery/commit/80dcf8e9c5c13adcd91f58046d712bf87d6c1d3f))
* **aws:** Add support for xray resource policy ([#4833](https://github.com/cloudquery/cloudquery/issues/4833)) ([b68dc35](https://github.com/cloudquery/cloudquery/commit/b68dc350487194de6d5600663a147cca21123c82))


### Bug Fixes

* **aws:** Add PKs To IAM Resources ([#6741](https://github.com/cloudquery/cloudquery/issues/6741)) ([f7cdb07](https://github.com/cloudquery/cloudquery/commit/f7cdb07d19447cfc3651892b7d8d870ce3f29c15))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.24.2 ([#6695](https://github.com/cloudquery/cloudquery/issues/6695)) ([694ab9f](https://github.com/cloudquery/cloudquery/commit/694ab9f3e20473146e3620d7b03bb17eb259d697))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.25.0 ([#6745](https://github.com/cloudquery/cloudquery/issues/6745)) ([9c41854](https://github.com/cloudquery/cloudquery/commit/9c418547c3bbff97449765e337182230fb5e40d5))
* introduce `aws_s3_bucket_grants` explicit primary key `(bucket_arn, grantee_id)`. If you've previously synced this table you'll need to drop it for the PK change (migration) to succeed. ([cf35801](https://github.com/cloudquery/cloudquery/commit/cf3580131cd9957b4a569e4ff44acc529b532826))

## [10.1.0](https://github.com/cloudquery/cloudquery/compare/plugins-source-aws-v10.0.0...plugins-source-aws-v10.1.0) (2023-01-11)


### Features

* **aws:** Add support for Savingsplans ([#6660](https://github.com/cloudquery/cloudquery/issues/6660)) ([6566ac5](https://github.com/cloudquery/cloudquery/commit/6566ac566816ab3de6d2a3433db88501b38b3564))


### Bug Fixes

* Update endpoints ([#6652](https://github.com/cloudquery/cloudquery/issues/6652)) ([6a53fa9](https://github.com/cloudquery/cloudquery/commit/6a53fa92f4e6b9097f5759384d07a513db1bdbbb))

## [10.0.0](https://github.com/cloudquery/cloudquery/compare/plugins-source-aws-v9.2.0...plugins-source-aws-v10.0.0) (2023-01-10)


### ⚠ BREAKING CHANGES

* **deps:** `aws_amp_rule_groups_namespaces` column `data` type changed from `IntArray` to `ByteArray`
* **deps:** `aws_iam_virtual_mfa_devices` column `base32_string_seed` type changed from `IntArray` to `ByteArray`
* **deps:** `aws_iam_virtual_mfa_devices` column `qr_code_png` type changed from `IntArray` to `ByteArray`

### Bug Fixes

* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.24.1 ([31405ae](https://github.com/cloudquery/cloudquery/commit/31405aec106a1ea9eef7a05fc46cb5f6dfabebce))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.24.1 ([31405ae](https://github.com/cloudquery/cloudquery/commit/31405aec106a1ea9eef7a05fc46cb5f6dfabebce))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.24.1 ([31405ae](https://github.com/cloudquery/cloudquery/commit/31405aec106a1ea9eef7a05fc46cb5f6dfabebce))

## [9.2.0](https://github.com/cloudquery/cloudquery/compare/plugins-source-aws-v9.1.1...plugins-source-aws-v9.2.0) (2023-01-10)


### Features

* **aws:** KMS Key policies ([#6562](https://github.com/cloudquery/cloudquery/issues/6562)) ([d8571aa](https://github.com/cloudquery/cloudquery/commit/d8571aa0172de38ef0ac063738b0fbab7c32fbf7)), closes [#6559](https://github.com/cloudquery/cloudquery/issues/6559)
* Move AWS to use built-in transformations (no codegen) ([#6337](https://github.com/cloudquery/cloudquery/issues/6337)) ([926d278](https://github.com/cloudquery/cloudquery/commit/926d2788a2108097288519709d951ce5dc47a1c3))


### Bug Fixes

* **aws:** Enable Partial Resolution of Lambda ([#6391](https://github.com/cloudquery/cloudquery/issues/6391)) ([e96a410](https://github.com/cloudquery/cloudquery/commit/e96a410394568145d80cddf73c16b59fe83f408d))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/config to v1.18.8 ([#6393](https://github.com/cloudquery/cloudquery/issues/6393)) ([ffba44f](https://github.com/cloudquery/cloudquery/commit/ffba44f1318eb401d2b7ce2fa91c155d8925d90d))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/feature/s3/manager to v1.11.47 ([#6395](https://github.com/cloudquery/cloudquery/issues/6395)) ([71ec9b9](https://github.com/cloudquery/cloudquery/commit/71ec9b99328ae4b7b0739a0c22258a805b586948))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/accessanalyzer to v1.19.0 ([#6396](https://github.com/cloudquery/cloudquery/issues/6396)) ([f19a34c](https://github.com/cloudquery/cloudquery/commit/f19a34cc0fa7897379d3d7e80432b26a0ad73f9e))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/account to v1.8.0 ([#6397](https://github.com/cloudquery/cloudquery/issues/6397)) ([98a7681](https://github.com/cloudquery/cloudquery/commit/98a76815f1d1b98dd2278dfc0d57b775345ea46a))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/acm to v1.17.0 ([#6398](https://github.com/cloudquery/cloudquery/issues/6398)) ([f59f6bd](https://github.com/cloudquery/cloudquery/commit/f59f6bda99e6fce0209d106db7b6389d9ec21700))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/amp to v1.16.0 ([#6399](https://github.com/cloudquery/cloudquery/issues/6399)) ([128ae76](https://github.com/cloudquery/cloudquery/commit/128ae76a9cd921c9b63e9a02c9d5261077e814b8))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/apigateway to v1.16.0 ([#6400](https://github.com/cloudquery/cloudquery/issues/6400)) ([4570f15](https://github.com/cloudquery/cloudquery/commit/4570f15b08f46c6a2b3e8b258bb17abb2a8b688f))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/apigatewayv2 to v1.13.0 ([#6401](https://github.com/cloudquery/cloudquery/issues/6401)) ([ca3b156](https://github.com/cloudquery/cloudquery/commit/ca3b1565fe5800637c03dfe892b56fea36418451))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/applicationautoscaling to v1.16.0 ([#6360](https://github.com/cloudquery/cloudquery/issues/6360)) ([afa651e](https://github.com/cloudquery/cloudquery/commit/afa651e3357170e3f4a19149c0ecadde0c4a0cca))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/applicationautoscaling to v1.17.0 ([#6402](https://github.com/cloudquery/cloudquery/issues/6402)) ([5c239ed](https://github.com/cloudquery/cloudquery/commit/5c239ed5b1294aa1356e45fcc74382eaad970a0c))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/apprunner to v1.16.0 ([#6403](https://github.com/cloudquery/cloudquery/issues/6403)) ([a312c4b](https://github.com/cloudquery/cloudquery/commit/a312c4b03bca32f9c158e08d4c19984859a0a546))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/appstream to v1.19.0 ([#6404](https://github.com/cloudquery/cloudquery/issues/6404)) ([1cd08e1](https://github.com/cloudquery/cloudquery/commit/1cd08e1591674c78ce3f688bd8b601ba4994cd96))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/appsync to v1.18.0 ([#6405](https://github.com/cloudquery/cloudquery/issues/6405)) ([76d4df5](https://github.com/cloudquery/cloudquery/commit/76d4df52517fcac585b9759045b3a1f650e6af7c))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/athena to v1.22.0 ([#6406](https://github.com/cloudquery/cloudquery/issues/6406)) ([815a0f9](https://github.com/cloudquery/cloudquery/commit/815a0f9a5cff400f4e8cc23ba53584b115817bac))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/autoscaling to v1.26.0 ([#6407](https://github.com/cloudquery/cloudquery/issues/6407)) ([26e995c](https://github.com/cloudquery/cloudquery/commit/26e995cbdbe73a7e42926b62e4f85e9689cf67a8))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/backup to v1.19.0 ([#6408](https://github.com/cloudquery/cloudquery/issues/6408)) ([01accc8](https://github.com/cloudquery/cloudquery/commit/01accc866a6f3020883ac50edf7bf979a73149df))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/cloudformation to v1.25.0 ([#6409](https://github.com/cloudquery/cloudquery/issues/6409)) ([89ffac6](https://github.com/cloudquery/cloudquery/commit/89ffac690b18edb055cc892ab341907d7c8b27d4))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/cloudfront to v1.24.0 ([#6410](https://github.com/cloudquery/cloudquery/issues/6410)) ([74a96fb](https://github.com/cloudquery/cloudquery/commit/74a96fb2e34202369d357b51c8d44aa130e2ad95))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/cloudhsmv2 to v1.14.0 ([#6411](https://github.com/cloudquery/cloudquery/issues/6411)) ([d12e25c](https://github.com/cloudquery/cloudquery/commit/d12e25cadf8e37d9a33269241c4a5ed14c04be83))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/cloudtrail to v1.22.0 ([#6412](https://github.com/cloudquery/cloudquery/issues/6412)) ([c1765b8](https://github.com/cloudquery/cloudquery/commit/c1765b859c644487b035ca52077f9764ccf80a2a))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/cloudwatch to v1.24.0 ([#6413](https://github.com/cloudquery/cloudquery/issues/6413)) ([c6324f3](https://github.com/cloudquery/cloudquery/commit/c6324f382f040b959819f5a6bcd8a93fd0b4530c))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs to v1.17.4 ([#6359](https://github.com/cloudquery/cloudquery/issues/6359)) ([8ca9085](https://github.com/cloudquery/cloudquery/commit/8ca9085225f996bd3fc77f4bff4c76790d68ad3b))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs to v1.18.0 ([#6414](https://github.com/cloudquery/cloudquery/issues/6414)) ([35b62ac](https://github.com/cloudquery/cloudquery/commit/35b62acbc10ee60318260e06a3ef74445ec14772))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/codebuild to v1.20.0 ([#6444](https://github.com/cloudquery/cloudquery/issues/6444)) ([b91ca03](https://github.com/cloudquery/cloudquery/commit/b91ca037c14a6e228ef0e95d54757757f505b747))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/codepipeline to v1.14.0 ([#6445](https://github.com/cloudquery/cloudquery/issues/6445)) ([bb7d3f2](https://github.com/cloudquery/cloudquery/commit/bb7d3f29aef0ff6b8307664f2546cf9f497b3e6a))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/cognitoidentity to v1.15.0 ([#6446](https://github.com/cloudquery/cloudquery/issues/6446)) ([c93c305](https://github.com/cloudquery/cloudquery/commit/c93c3059e570dbace6a6cf18f0f48b8dcfcff449))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider to v1.22.0 ([#6447](https://github.com/cloudquery/cloudquery/issues/6447)) ([4ffbfd5](https://github.com/cloudquery/cloudquery/commit/4ffbfd52267bb6386527d8f23df671b2752cbbde))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/configservice to v1.29.0 ([#6448](https://github.com/cloudquery/cloudquery/issues/6448)) ([2433cbe](https://github.com/cloudquery/cloudquery/commit/2433cbe134d1f5041eb7cfb7e08dbfa783f47f6a))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/databasemigrationservice to v1.23.0 ([#6449](https://github.com/cloudquery/cloudquery/issues/6449)) ([9c2bb7c](https://github.com/cloudquery/cloudquery/commit/9c2bb7c22a0cb5c8fd38708c311386bc6e36f19f))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/dax to v1.12.0 ([#6450](https://github.com/cloudquery/cloudquery/issues/6450)) ([da537c6](https://github.com/cloudquery/cloudquery/commit/da537c6d28a10e4334ca42a2180bae763159146b))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/directconnect to v1.18.0 ([#6451](https://github.com/cloudquery/cloudquery/issues/6451)) ([89e34fe](https://github.com/cloudquery/cloudquery/commit/89e34fe8db2ffb40b6b1a6373e13db28346b3f6e))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/docdb to v1.20.0 ([#6452](https://github.com/cloudquery/cloudquery/issues/6452)) ([472536a](https://github.com/cloudquery/cloudquery/commit/472536a5a39c2d9176adbe3db2ca2f4b192ff2df))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/dynamodb to v1.18.0 ([#6453](https://github.com/cloudquery/cloudquery/issues/6453)) ([561db49](https://github.com/cloudquery/cloudquery/commit/561db492f96428cc50e6e8dbaa5afb1ee5f8eadf))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/ecr to v1.18.0 ([#6454](https://github.com/cloudquery/cloudquery/issues/6454)) ([ecdbe98](https://github.com/cloudquery/cloudquery/commit/ecdbe985cc0834d3bc60708b48ef9769dd9a2a52))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/ecrpublic to v1.14.0 ([#6455](https://github.com/cloudquery/cloudquery/issues/6455)) ([7984da9](https://github.com/cloudquery/cloudquery/commit/7984da97223ad938d141c0be2aac09557cc89bcd))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/ecs to v1.23.0 ([#6456](https://github.com/cloudquery/cloudquery/issues/6456)) ([94b8b4c](https://github.com/cloudquery/cloudquery/commit/94b8b4c4dfc0c1c8bab2e05d5db59d1793193753))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/efs to v1.19.0 ([#6457](https://github.com/cloudquery/cloudquery/issues/6457)) ([a56841a](https://github.com/cloudquery/cloudquery/commit/a56841ae4b07aa42f20ccd556d54f2e7824f97aa))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/eks to v1.27.0 ([#6458](https://github.com/cloudquery/cloudquery/issues/6458)) ([4a74a54](https://github.com/cloudquery/cloudquery/commit/4a74a54e138f5271727281fbaeb5afaa5ff38303))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/elasticache to v1.26.0 ([#6459](https://github.com/cloudquery/cloudquery/issues/6459)) ([5c70e45](https://github.com/cloudquery/cloudquery/commit/5c70e455c721bab217ca7ef5bc4b9e58de25fb60))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/elasticbeanstalk to v1.15.0 ([#6460](https://github.com/cloudquery/cloudquery/issues/6460)) ([3333cfb](https://github.com/cloudquery/cloudquery/commit/3333cfb377deebedd278632540987f1c9988b281))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/elasticloadbalancing to v1.15.0 ([#6461](https://github.com/cloudquery/cloudquery/issues/6461)) ([d72d213](https://github.com/cloudquery/cloudquery/commit/d72d213e3daabf1ac4520866136bcb5990fb7531))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2 to v1.19.0 ([#6462](https://github.com/cloudquery/cloudquery/issues/6462)) ([57ba804](https://github.com/cloudquery/cloudquery/commit/57ba8046dab6ab21110326a36ba3b6cf2e9a026e))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/elasticsearchservice to v1.18.0 ([#6463](https://github.com/cloudquery/cloudquery/issues/6463)) ([e282e84](https://github.com/cloudquery/cloudquery/commit/e282e8476eea9345fd7a41a72ee4eaaa2f26351a))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/elastictranscoder to v1.14.0 ([#6464](https://github.com/cloudquery/cloudquery/issues/6464)) ([163a581](https://github.com/cloudquery/cloudquery/commit/163a58103ff3a5454c83e82fca0685ecd9ee2192))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/emr to v1.22.0 ([#6465](https://github.com/cloudquery/cloudquery/issues/6465)) ([f149f08](https://github.com/cloudquery/cloudquery/commit/f149f084163145ec03023122161e39899f2ab763))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/eventbridge to v1.17.0 ([#6473](https://github.com/cloudquery/cloudquery/issues/6473)) ([188848f](https://github.com/cloudquery/cloudquery/commit/188848f482ef1f0657fd78203dfa279709bff53b))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/firehose to v1.16.0 ([#6474](https://github.com/cloudquery/cloudquery/issues/6474)) ([de60adb](https://github.com/cloudquery/cloudquery/commit/de60adb2c5ffa8eea22e202089d25f01bff373b6))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/frauddetector to v1.21.0 ([#6475](https://github.com/cloudquery/cloudquery/issues/6475)) ([19b154a](https://github.com/cloudquery/cloudquery/commit/19b154ac1da2e3792a8a25a4f6c11f4fb9dabc86))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/fsx to v1.28.0 ([#6476](https://github.com/cloudquery/cloudquery/issues/6476)) ([7b39554](https://github.com/cloudquery/cloudquery/commit/7b39554e05b53e7ae54679c19f36e71d45fcb3c8))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/glacier to v1.14.0 ([#6477](https://github.com/cloudquery/cloudquery/issues/6477)) ([8d0887f](https://github.com/cloudquery/cloudquery/commit/8d0887fb4a1d48b0e7b5b8e8b1732f7aa57a7142))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/glue to v1.39.0 ([#6478](https://github.com/cloudquery/cloudquery/issues/6478)) ([4001267](https://github.com/cloudquery/cloudquery/commit/4001267f92b58fbd6069a170a6173b80f8d53417))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/guardduty to v1.17.0 ([#6479](https://github.com/cloudquery/cloudquery/issues/6479)) ([651dfe4](https://github.com/cloudquery/cloudquery/commit/651dfe437778b2e3df663a33c4787fba5b0ab61d))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/iam to v1.19.0 ([#6480](https://github.com/cloudquery/cloudquery/issues/6480)) ([30cf79d](https://github.com/cloudquery/cloudquery/commit/30cf79d234d86dc3ab36350ebcbdca9696174d03))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/identitystore to v1.16.0 ([#6481](https://github.com/cloudquery/cloudquery/issues/6481)) ([c7dd4a6](https://github.com/cloudquery/cloudquery/commit/c7dd4a64734ad60be65a9c31d823d311f53ddd7e))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/inspector to v1.13.0 ([#6482](https://github.com/cloudquery/cloudquery/issues/6482)) ([e628a51](https://github.com/cloudquery/cloudquery/commit/e628a51e8912222ec5a6a44b1dc117768cce30d5))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/inspector2 to v1.11.0 ([#6483](https://github.com/cloudquery/cloudquery/issues/6483)) ([fe8fd70](https://github.com/cloudquery/cloudquery/commit/fe8fd70ca0f1782f734f1e45730ca09384b86664))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/iot to v1.33.0 ([#6484](https://github.com/cloudquery/cloudquery/issues/6484)) ([d66ca25](https://github.com/cloudquery/cloudquery/commit/d66ca2585ac6a61a963e80d4b3e5fc0b59713260))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/kafka to v1.19.0 ([#6485](https://github.com/cloudquery/cloudquery/issues/6485)) ([acb2683](https://github.com/cloudquery/cloudquery/commit/acb26838a6d09ce9581224964de386e56caae3a9))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/kinesis to v1.17.0 ([#6486](https://github.com/cloudquery/cloudquery/issues/6486)) ([31b551f](https://github.com/cloudquery/cloudquery/commit/31b551f89e3e48e8e669e2ae960205d06d5fba4b))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/kms to v1.20.0 ([#6487](https://github.com/cloudquery/cloudquery/issues/6487)) ([5a8b81d](https://github.com/cloudquery/cloudquery/commit/5a8b81da157189a9ec92522b06992b20b969bc6c))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/lambda to v1.27.0 ([#6488](https://github.com/cloudquery/cloudquery/issues/6488)) ([de83eda](https://github.com/cloudquery/cloudquery/commit/de83edafcafd92df94ad9dc23967ffc0f278d3c8))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/lightsail to v1.25.0 ([#6489](https://github.com/cloudquery/cloudquery/issues/6489)) ([99ce478](https://github.com/cloudquery/cloudquery/commit/99ce47890a9d303537eb833e60a95fa1f63495cc))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/mq to v1.14.0 ([#6490](https://github.com/cloudquery/cloudquery/issues/6490)) ([d5b465c](https://github.com/cloudquery/cloudquery/commit/d5b465cf656c706fdf588657f6e31fd2dbe40021))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/mwaa to v1.14.0 ([#6491](https://github.com/cloudquery/cloudquery/issues/6491)) ([c565b0a](https://github.com/cloudquery/cloudquery/commit/c565b0aa691019c1a11058b062fa2d2a30bceb7c))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/neptune to v1.19.0 ([#6492](https://github.com/cloudquery/cloudquery/issues/6492)) ([d373abd](https://github.com/cloudquery/cloudquery/commit/d373abd5031631885e6c59a904f11459a337b156))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/organizations to v1.18.0 ([#6495](https://github.com/cloudquery/cloudquery/issues/6495)) ([bbb2457](https://github.com/cloudquery/cloudquery/commit/bbb2457db330818f24496719f4bceb32384bc9cb))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/qldb to v1.15.0 ([#6496](https://github.com/cloudquery/cloudquery/issues/6496)) ([5f01dcf](https://github.com/cloudquery/cloudquery/commit/5f01dcfaf5b4c07f05905e9b25adea9482270379))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/quicksight to v1.29.0 ([#6497](https://github.com/cloudquery/cloudquery/issues/6497)) ([6c6788c](https://github.com/cloudquery/cloudquery/commit/6c6788c677dc7632f01dd029c155c8c7ce18c622))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/ram to v1.17.0 ([#6498](https://github.com/cloudquery/cloudquery/issues/6498)) ([306bcc5](https://github.com/cloudquery/cloudquery/commit/306bcc5d81b2b86014ab5b2cf08d59f9951c2426))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/rds to v1.39.0 ([#6499](https://github.com/cloudquery/cloudquery/issues/6499)) ([ed4cd10](https://github.com/cloudquery/cloudquery/commit/ed4cd10e81c576f1c03c6c77400db17a673db843))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/redshift to v1.27.0 ([#6500](https://github.com/cloudquery/cloudquery/issues/6500)) ([4c82e8b](https://github.com/cloudquery/cloudquery/commit/4c82e8be4121cd1e9c6025ee5e1c55021ca35ca8))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/resourcegroups to v1.13.0 ([#6501](https://github.com/cloudquery/cloudquery/issues/6501)) ([9efd121](https://github.com/cloudquery/cloudquery/commit/9efd1217a33b38944f91e64f0e3cfc50c3ce2d82))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/route53 to v1.26.0 ([#6502](https://github.com/cloudquery/cloudquery/issues/6502)) ([135dc6b](https://github.com/cloudquery/cloudquery/commit/135dc6b79f07e6af16aa44f18c19b8b585e2152f))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/route53domains to v1.14.0 ([#6503](https://github.com/cloudquery/cloudquery/issues/6503)) ([ba2aa9e](https://github.com/cloudquery/cloudquery/commit/ba2aa9e975664b4249c093c329a7c5f9d2884bae))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/s3control to v1.29.0 ([#6505](https://github.com/cloudquery/cloudquery/issues/6505)) ([47fa203](https://github.com/cloudquery/cloudquery/commit/47fa203c1be135d6e06ca6a39c580a502d24e29f))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/sagemaker to v1.62.0 ([#6506](https://github.com/cloudquery/cloudquery/issues/6506)) ([57207cf](https://github.com/cloudquery/cloudquery/commit/57207cfd05228f8e257764cabee6b6f14b01d1c1))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/scheduler to v1.1.0 ([#6507](https://github.com/cloudquery/cloudquery/issues/6507)) ([a56a3e0](https://github.com/cloudquery/cloudquery/commit/a56a3e0b885b913d0cee5bc17e01350e967cf866))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/secretsmanager to v1.18.0 ([#6508](https://github.com/cloudquery/cloudquery/issues/6508)) ([83f6116](https://github.com/cloudquery/cloudquery/commit/83f6116fa8e4e355502a07850def33304059d08a))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/servicecatalog to v1.16.0 ([#6509](https://github.com/cloudquery/cloudquery/issues/6509)) ([1e69f8d](https://github.com/cloudquery/cloudquery/commit/1e69f8dd2ea31117bbe0bd3ff2056eab849ccd67))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/servicecatalogappregistry to v1.16.0 ([#6510](https://github.com/cloudquery/cloudquery/issues/6510)) ([6a7f928](https://github.com/cloudquery/cloudquery/commit/6a7f928051691ca17e219af069d2422ed373cc33))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/servicequotas to v1.14.0 ([#6511](https://github.com/cloudquery/cloudquery/issues/6511)) ([d12e3b8](https://github.com/cloudquery/cloudquery/commit/d12e3b8a86b6d838935b0c0404e1041445ab5249))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/ses to v1.15.0 ([#6512](https://github.com/cloudquery/cloudquery/issues/6512)) ([4f8863c](https://github.com/cloudquery/cloudquery/commit/4f8863cd922df654ef4860897ee22f8437192539))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/sesv2 to v1.16.0 ([#6513](https://github.com/cloudquery/cloudquery/issues/6513)) ([a4987ae](https://github.com/cloudquery/cloudquery/commit/a4987ae191b9f5e660d29ca4a99f1469dd9e1fe7))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/sfn to v1.17.0 ([#6514](https://github.com/cloudquery/cloudquery/issues/6514)) ([c8e7d9c](https://github.com/cloudquery/cloudquery/commit/c8e7d9cc72fd6ab71dc8916ef6ddc8902bf5b92b))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/ssm to v1.34.0 ([#6361](https://github.com/cloudquery/cloudquery/issues/6361)) ([2413dba](https://github.com/cloudquery/cloudquery/commit/2413dba0ba477e5ba958c4c97e7d330275ea4323))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.18.0 ([#6339](https://github.com/cloudquery/cloudquery/issues/6339)) ([158365a](https://github.com/cloudquery/cloudquery/commit/158365a78dfa4389074f716a0f581f18fedc1080))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.19.0 ([#6363](https://github.com/cloudquery/cloudquery/issues/6363)) ([ae6967c](https://github.com/cloudquery/cloudquery/commit/ae6967c22002c554a083f444eb611ac3e6d2698f))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.20.0 ([#6376](https://github.com/cloudquery/cloudquery/issues/6376)) ([d6187ec](https://github.com/cloudquery/cloudquery/commit/d6187ec584f13be4fe9362dd393385b19d386113))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.21.0 ([#6382](https://github.com/cloudquery/cloudquery/issues/6382)) ([5baea40](https://github.com/cloudquery/cloudquery/commit/5baea40d2aec4e807db839c928be2e037d572bef))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.22.0 ([#6516](https://github.com/cloudquery/cloudquery/issues/6516)) ([b7e4e73](https://github.com/cloudquery/cloudquery/commit/b7e4e737a5f4d8f254960426ea8ba555d8f9b944))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.23.0 ([#6522](https://github.com/cloudquery/cloudquery/issues/6522)) ([ce24f1d](https://github.com/cloudquery/cloudquery/commit/ce24f1d64394cbb5ab07dcaa4af66c53f77f700f))
* Update endpoints ([#6327](https://github.com/cloudquery/cloudquery/issues/6327)) ([11d6973](https://github.com/cloudquery/cloudquery/commit/11d6973618a135c00f6232b3462a6f82cb501380))
* Update endpoints ([#6440](https://github.com/cloudquery/cloudquery/issues/6440)) ([499f421](https://github.com/cloudquery/cloudquery/commit/499f421fa96b2d3c63f0679b32e00f98cc0da4b3))
* Update endpoints ([#6558](https://github.com/cloudquery/cloudquery/issues/6558)) ([45c4753](https://github.com/cloudquery/cloudquery/commit/45c4753d3d81a799a46a6fdf8782c043b5666c77))

## [9.1.1](https://github.com/cloudquery/cloudquery/compare/plugins-source-aws-v9.1.0...plugins-source-aws-v9.1.1) (2023-01-03)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.17.0 ([#6256](https://github.com/cloudquery/cloudquery/issues/6256)) ([b19f6cd](https://github.com/cloudquery/cloudquery/commit/b19f6cd8e2c39994aeb19d78e78e927d6c3cf580))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.17.2 ([#6260](https://github.com/cloudquery/cloudquery/issues/6260)) ([805972a](https://github.com/cloudquery/cloudquery/commit/805972aa67ce54e3358501c6b7ee5d85e5f65cac))

## [9.1.0](https://github.com/cloudquery/cloudquery/compare/plugins-source-aws-v9.0.1...plugins-source-aws-v9.1.0) (2023-01-02)


### Features

* **aws:** Route53 Add transfer lock Status ([#6022](https://github.com/cloudquery/cloudquery/issues/6022)) ([c3f0e37](https://github.com/cloudquery/cloudquery/commit/c3f0e37bb5ee63e2b829be19095310c036d81271))


### Bug Fixes

* **aws:** Remove account validation ([#6226](https://github.com/cloudquery/cloudquery/issues/6226)) ([98dc0b2](https://github.com/cloudquery/cloudquery/commit/98dc0b2470e21c7397ca874c17409913df968abf))
* **deps:** Update github.com/gocarina/gocsv digest to 1fea7ae ([#6168](https://github.com/cloudquery/cloudquery/issues/6168)) ([6fc737b](https://github.com/cloudquery/cloudquery/commit/6fc737b1ba919e87c2bbabc9fa5aad014196f70d))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/apigateway to v1.15.28 ([#6141](https://github.com/cloudquery/cloudquery/issues/6141)) ([5a7b876](https://github.com/cloudquery/cloudquery/commit/5a7b87661425ddd183da4294fadeed2e99b2ca14))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/cloudfront to v1.23.0 ([#6159](https://github.com/cloudquery/cloudquery/issues/6159)) ([6e42815](https://github.com/cloudquery/cloudquery/commit/6e428155856e1dafc7726ac375476f88c8463d9a))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/elasticache to v1.25.0 ([#6111](https://github.com/cloudquery/cloudquery/issues/6111)) ([b1f5d4b](https://github.com/cloudquery/cloudquery/commit/b1f5d4bd194c5265c3ff463b09f9bf2f70aa61b5))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/emr to v1.21.0 ([#6142](https://github.com/cloudquery/cloudquery/issues/6142)) ([d68af16](https://github.com/cloudquery/cloudquery/commit/d68af16ad392fa6561871edd1052de88603bca9a))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/rds to v1.38.0 ([#6112](https://github.com/cloudquery/cloudquery/issues/6112)) ([549e290](https://github.com/cloudquery/cloudquery/commit/549e2904070da2c86c745ab7db573fbaaf95570b))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/secretsmanager to v1.17.0 ([#6143](https://github.com/cloudquery/cloudquery/issues/6143)) ([2d99f4f](https://github.com/cloudquery/cloudquery/commit/2d99f4f953df0378dd0255cd5d3856f797d4b8d8))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/transfer to v1.27.0 ([#6049](https://github.com/cloudquery/cloudquery/issues/6049)) ([87cf1f6](https://github.com/cloudquery/cloudquery/commit/87cf1f65f0e79f2c8d8e6a5c30ab6264074afa8a))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.14.0 ([#6025](https://github.com/cloudquery/cloudquery/issues/6025)) ([35b2cfc](https://github.com/cloudquery/cloudquery/commit/35b2cfc7fc7bcdaceb7ee674e3a17f0f5673b366))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.15.0 ([#6071](https://github.com/cloudquery/cloudquery/issues/6071)) ([684b525](https://github.com/cloudquery/cloudquery/commit/684b525aaa285fcae70dd87af56679c1205adebe))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.15.1 ([#6079](https://github.com/cloudquery/cloudquery/issues/6079)) ([650659c](https://github.com/cloudquery/cloudquery/commit/650659c3c6766df571868e2ec3a2007cb76696eb))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.16.0 ([#6098](https://github.com/cloudquery/cloudquery/issues/6098)) ([7bacdf3](https://github.com/cloudquery/cloudquery/commit/7bacdf3364716eab08fa1a84ae4047b42edeee7e))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.16.1 ([#6214](https://github.com/cloudquery/cloudquery/issues/6214)) ([53b2415](https://github.com/cloudquery/cloudquery/commit/53b241508d7511d4b5fa74cc4262d180c1e6df66))
* Update endpoints ([#6136](https://github.com/cloudquery/cloudquery/issues/6136)) ([808ffab](https://github.com/cloudquery/cloudquery/commit/808ffab30a149b80d0ae00c893fff6f281a2d9d8))
* Update endpoints ([#6158](https://github.com/cloudquery/cloudquery/issues/6158)) ([dd773b8](https://github.com/cloudquery/cloudquery/commit/dd773b8756195d8ceb8d70cc4e34a4b3f28e151d))

## [9.0.1](https://github.com/cloudquery/cloudquery/compare/plugins-source-aws-v9.0.0...plugins-source-aws-v9.0.1) (2022-12-27)


### Bug Fixes

* **aws_ses_active_receipt_rule_sets:** Don't call DescribeActiveReceiptRuleSet on unsupported regions ([#5997](https://github.com/cloudquery/cloudquery/issues/5997)) ([6a63147](https://github.com/cloudquery/cloudquery/commit/6a631478014ee6cf2d21f3426fe903715e67e3c0))
* **aws_ses_active_receipt_rule_sets:** Don't sync empty return values from DescribeActiveReceiptRuleSet ([#5992](https://github.com/cloudquery/cloudquery/issues/5992)) ([5837069](https://github.com/cloudquery/cloudquery/commit/58370697f69f10b3ada7c50b5dd6b5ce28cf0f91))

## [9.0.0](https://github.com/cloudquery/cloudquery/compare/plugins-source-aws-v8.1.0...plugins-source-aws-v9.0.0) (2022-12-27)


### ⚠ BREAKING CHANGES

* **aws:** Organizational Unit (OU) IDs are now validated to match either `ou-` or `r-` formats

### Features

* **aws:** Add AWS Managed Service for Prometheus ([#5875](https://github.com/cloudquery/cloudquery/issues/5875)) ([8cad8e4](https://github.com/cloudquery/cloudquery/commit/8cad8e4101bd7fc7703628189304904f0a804371))
* **aws:** Add more Elastic Search resources ([#5867](https://github.com/cloudquery/cloudquery/issues/5867)) ([0762a59](https://github.com/cloudquery/cloudquery/commit/0762a59bf2a2f9a1842461c72df25f06aef66726))
* **aws:** Fetch tags for aws_elasticache_clusters ([#5911](https://github.com/cloudquery/cloudquery/issues/5911)) ([65b5093](https://github.com/cloudquery/cloudquery/commit/65b5093e65b9201c2678c59113148e92a03a2f89)), closes [#5899](https://github.com/cloudquery/cloudquery/issues/5899)
* **aws:** Paginate Ec2 Describe Images ([#5878](https://github.com/cloudquery/cloudquery/issues/5878)) ([9a37b52](https://github.com/cloudquery/cloudquery/commit/9a37b52c1ff692d4003fe05af36c3e67bc69e1de))
* **aws:** Support Custom endpoints ([#5942](https://github.com/cloudquery/cloudquery/issues/5942)) ([746ba65](https://github.com/cloudquery/cloudquery/commit/746ba6501727e5bfe93e45720ccee410d56ab6c3))
* **aws:** Support recursive listing of AWS orgs and skipping of OUs and accounts ([7908fd4](https://github.com/cloudquery/cloudquery/commit/7908fd4b0e766f9e549d274d45e21ffd192976b0))


### Bug Fixes

* **deps:** Update module github.com/aws/aws-sdk-go-v2/config to v1.18.7 ([#5879](https://github.com/cloudquery/cloudquery/issues/5879)) ([c2c082e](https://github.com/cloudquery/cloudquery/commit/c2c082e08adc1ec652dcb218f165e491ec5bd878))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/credentials to v1.13.7 ([#5811](https://github.com/cloudquery/cloudquery/issues/5811)) ([135427a](https://github.com/cloudquery/cloudquery/commit/135427a71316201e83a57397737741eabdec9ac8))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/feature/s3/manager to v1.11.46 ([#5880](https://github.com/cloudquery/cloudquery/issues/5880)) ([a4e8c81](https://github.com/cloudquery/cloudquery/commit/a4e8c81b7fca19cc80b3dc367425c91a164f2ecd))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/cloudhsmv2 to v1.13.24 ([#5881](https://github.com/cloudquery/cloudquery/issues/5881)) ([741b48b](https://github.com/cloudquery/cloudquery/commit/741b48b9ccdf843d91774ee81b4142b5971f4192))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/ec2 to v1.77.0 ([#5883](https://github.com/cloudquery/cloudquery/issues/5883)) ([e3dcb98](https://github.com/cloudquery/cloudquery/commit/e3dcb983cf593780a398b5d7e247cb42c42f6a99))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/fsx to v1.27.0 ([#5950](https://github.com/cloudquery/cloudquery/issues/5950)) ([f9e542d](https://github.com/cloudquery/cloudquery/commit/f9e542d295b2ced02011398a550837bbd6c23eae))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/identitystore to v1.15.10 ([#5904](https://github.com/cloudquery/cloudquery/issues/5904)) ([3a49662](https://github.com/cloudquery/cloudquery/commit/3a496629346e850797476d1d9700a4f1d5a4cef4))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/inspector2 to v1.10.0 ([#5951](https://github.com/cloudquery/cloudquery/issues/5951)) ([9586d36](https://github.com/cloudquery/cloudquery/commit/9586d3663edf0e9885e63386531c8dd65508e27b))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/rds to v1.37.0 ([#5932](https://github.com/cloudquery/cloudquery/issues/5932)) ([bc04b3e](https://github.com/cloudquery/cloudquery/commit/bc04b3e967c1fb082cf130691aeb8de20f50ecbd))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/route53domains to v1.13.0 ([#5884](https://github.com/cloudquery/cloudquery/issues/5884)) ([06acf5a](https://github.com/cloudquery/cloudquery/commit/06acf5acbc5f59984c0b8ffcc48aba726db161b4))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/sagemaker to v1.59.0 ([#5819](https://github.com/cloudquery/cloudquery/issues/5819)) ([7ee4849](https://github.com/cloudquery/cloudquery/commit/7ee4849b948d880e20e6d2fd5263a386a1d2c1ad))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/sagemaker to v1.60.0 ([#5885](https://github.com/cloudquery/cloudquery/issues/5885)) ([ba72fd8](https://github.com/cloudquery/cloudquery/commit/ba72fd84edb59d01eb12f5d06a3fab7d0e3648da))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/sagemaker to v1.61.0 ([#5906](https://github.com/cloudquery/cloudquery/issues/5906)) ([5f12d55](https://github.com/cloudquery/cloudquery/commit/5f12d55de0832803ba8e4a486debcda59b10d984))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/secretsmanager to v1.16.11 ([#5931](https://github.com/cloudquery/cloudquery/issues/5931)) ([f0bdd7f](https://github.com/cloudquery/cloudquery/commit/f0bdd7f96e82dce648b6f385d45639152d46c675))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/ssm to v1.33.4 ([#5905](https://github.com/cloudquery/cloudquery/issues/5905)) ([8977bb6](https://github.com/cloudquery/cloudquery/commit/8977bb635a7c59a215cb608d75b0a0130b966a24))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/ssoadmin to v1.15.16 ([#5882](https://github.com/cloudquery/cloudquery/issues/5882)) ([be860f3](https://github.com/cloudquery/cloudquery/commit/be860f34a9861b6f4494a67a118008f9b112962f))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/transfer to v1.26.0 ([#5907](https://github.com/cloudquery/cloudquery/issues/5907)) ([fa63d48](https://github.com/cloudquery/cloudquery/commit/fa63d48b59af5200e4389c9121dbf5c9f45241b8))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.13.1 ([#5897](https://github.com/cloudquery/cloudquery/issues/5897)) ([ad15915](https://github.com/cloudquery/cloudquery/commit/ad15915f2951a75729859f6f1377ed789f8ba115))
* Update endpoints ([#5876](https://github.com/cloudquery/cloudquery/issues/5876)) ([3a62218](https://github.com/cloudquery/cloudquery/commit/3a6221862c57b1f127f3545ed61010707030244e))
* Update endpoints ([#5903](https://github.com/cloudquery/cloudquery/issues/5903)) ([6979d07](https://github.com/cloudquery/cloudquery/commit/6979d07daea741f94d79a329156db7ac8476fc76))
* Update endpoints ([#5929](https://github.com/cloudquery/cloudquery/issues/5929)) ([2b8e6a1](https://github.com/cloudquery/cloudquery/commit/2b8e6a13e9268e5661121883b7e85de13e507148))
* Update endpoints ([#5948](https://github.com/cloudquery/cloudquery/issues/5948)) ([8829544](https://github.com/cloudquery/cloudquery/commit/8829544be65e7a33adbfe3c9d4902961d29164a9))

## [8.1.0](https://github.com/cloudquery/cloudquery/compare/plugins-source-aws-v8.0.0...plugins-source-aws-v8.1.0) (2022-12-20)


### Features

* **aws:** Add aws_organizations ([#5677](https://github.com/cloudquery/cloudquery/issues/5677)) ([06c57fd](https://github.com/cloudquery/cloudquery/commit/06c57fd8d36f4c1c84748e036bdd31c2bfd03392)), closes [#5621](https://github.com/cloudquery/cloudquery/issues/5621)

## [8.0.0](https://github.com/cloudquery/cloudquery/compare/plugins-source-aws-v7.4.0...plugins-source-aws-v8.0.0) (2022-12-20)


### ⚠ BREAKING CHANGES

* **aws:** Rename `id` column to `domain_id` for `aws_elasticsearch_domains`
* **aws:** The primary key for the `aws_ses_identies` table is now `arn` instead of `account_id `, `name ` and `region`.

### Features

* **aws:** Add `arn` to `aws_ses_identies` ([6338855](https://github.com/cloudquery/cloudquery/commit/633885520d9fc8d8c9c9e4fb4d9b017c9ecfc149))
* **aws:** Add Elastic Transcoder resources ([#5804](https://github.com/cloudquery/cloudquery/issues/5804)) ([bb8d28e](https://github.com/cloudquery/cloudquery/commit/bb8d28e12bd620ea89c360ccc078bd0e480afb79))
* **aws:** AWS SES resources: `aws_ses_active_receipt_rule_sets` & `aws_ses_custom_verification_email_templates` ([#5792](https://github.com/cloudquery/cloudquery/issues/5792)) ([6ce287d](https://github.com/cloudquery/cloudquery/commit/6ce287d94f0c26fbfe7bd71347eb4407fdbddd7f))
* **aws:** Make `arn` primary key of `aws_elasticsearch_domains`. Rename `id` column to `domain_id` ([c4d4a22](https://github.com/cloudquery/cloudquery/commit/c4d4a2292cb1d4d38ea43935f6a7620422bcca6b))


### Bug Fixes

* **aws-codegen:** Check error returned by formatAndWriteFile ([#5793](https://github.com/cloudquery/cloudquery/issues/5793)) ([ed0d9d7](https://github.com/cloudquery/cloudquery/commit/ed0d9d78ffe5598bcd0bbabfb5862f2cf4731ddd))
* **deps:** Update module github.com/aws/aws-sdk-go-v2 to v1.17.3 ([#5685](https://github.com/cloudquery/cloudquery/issues/5685)) ([e540ee5](https://github.com/cloudquery/cloudquery/commit/e540ee53c58b580f64d1317dbaffddf1b3950e20))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/config to v1.18.5 ([#5686](https://github.com/cloudquery/cloudquery/issues/5686)) ([72d50ea](https://github.com/cloudquery/cloudquery/commit/72d50ea25d1cd5018f163ead27909aee3872e0ed))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/config to v1.18.6 ([#5810](https://github.com/cloudquery/cloudquery/issues/5810)) ([6c5ef8d](https://github.com/cloudquery/cloudquery/commit/6c5ef8d8346248ed4efa8b714f662e3d97243103))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/credentials to v1.13.5 ([#5687](https://github.com/cloudquery/cloudquery/issues/5687)) ([8b8cd71](https://github.com/cloudquery/cloudquery/commit/8b8cd71ced04471fcf1c451d13facb1817881e6f))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/feature/s3/manager to v1.11.44 ([#5688](https://github.com/cloudquery/cloudquery/issues/5688)) ([5e37ab3](https://github.com/cloudquery/cloudquery/commit/5e37ab37f8db8a4d9bde577ba0fe03b08d9d7ebd))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/feature/s3/manager to v1.11.45 ([#5812](https://github.com/cloudquery/cloudquery/issues/5812)) ([a1fe91b](https://github.com/cloudquery/cloudquery/commit/a1fe91bc46ad9eeaea55a00e6ac5c8ff5cdfea8d))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/accessanalyzer to v1.18.2 ([#5689](https://github.com/cloudquery/cloudquery/issues/5689)) ([801834a](https://github.com/cloudquery/cloudquery/commit/801834a0acc9b8da84dfd2fd87edfad6445abf61))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/account to v1.7.13 ([#5690](https://github.com/cloudquery/cloudquery/issues/5690)) ([4e8c448](https://github.com/cloudquery/cloudquery/commit/4e8c448b196e6de8461ebd1d22e4601ed2f751a4))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/acm to v1.16.5 ([#5691](https://github.com/cloudquery/cloudquery/issues/5691)) ([da5760d](https://github.com/cloudquery/cloudquery/commit/da5760d5f578222605e47a787a8e2f3991d134a3))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/apigateway to v1.15.27 ([#5692](https://github.com/cloudquery/cloudquery/issues/5692)) ([2d838df](https://github.com/cloudquery/cloudquery/commit/2d838dfdca2a37b70fe36e59176a70a86ee7c1b3))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/apigatewayv2 to v1.12.25 ([#5693](https://github.com/cloudquery/cloudquery/issues/5693)) ([748dd66](https://github.com/cloudquery/cloudquery/commit/748dd6677e6fae25e21ece18c15ce2ca20e7bea8))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/applicationautoscaling to v1.15.25 ([#5694](https://github.com/cloudquery/cloudquery/issues/5694)) ([9d74fbb](https://github.com/cloudquery/cloudquery/commit/9d74fbbee0f0d9beb2be25866e41dad518419fd3))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/apprunner to v1.15.2 ([#5695](https://github.com/cloudquery/cloudquery/issues/5695)) ([6d7be4b](https://github.com/cloudquery/cloudquery/commit/6d7be4b57695dff944840b0f888a5c8913ab1590))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/appstream to v1.18.3 ([#5696](https://github.com/cloudquery/cloudquery/issues/5696)) ([6cffcd8](https://github.com/cloudquery/cloudquery/commit/6cffcd868c3d422a67242591fe41ac1384ddbfdf))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/appsync to v1.17.1 ([#5697](https://github.com/cloudquery/cloudquery/issues/5697)) ([476b82b](https://github.com/cloudquery/cloudquery/commit/476b82b7502e1fe55bd1a883435f51764c7d4237))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/athena to v1.20.3 ([#5698](https://github.com/cloudquery/cloudquery/issues/5698)) ([2e9aa33](https://github.com/cloudquery/cloudquery/commit/2e9aa33984a499558d481c7de818a9569187aff8))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/athena to v1.21.0 ([#5813](https://github.com/cloudquery/cloudquery/issues/5813)) ([3164e31](https://github.com/cloudquery/cloudquery/commit/3164e31c34df0e38ac245718b9b83eb0ec60b06d))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/autoscaling to v1.25.1 ([#5699](https://github.com/cloudquery/cloudquery/issues/5699)) ([f4ae75f](https://github.com/cloudquery/cloudquery/commit/f4ae75f9020614e50df63b24a4338f96190c1d12))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/backup to v1.18.2 ([#5700](https://github.com/cloudquery/cloudquery/issues/5700)) ([f592456](https://github.com/cloudquery/cloudquery/commit/f592456c7d53e3348321bffa5b415da29364e3a1))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/cloudformation to v1.24.3 ([#5701](https://github.com/cloudquery/cloudquery/issues/5701)) ([3c8f3de](https://github.com/cloudquery/cloudquery/commit/3c8f3de2ca8d74716fb5f7415bd06f9b52696831))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/cloudfront to v1.22.1 ([#5702](https://github.com/cloudquery/cloudquery/issues/5702)) ([c12fe3f](https://github.com/cloudquery/cloudquery/commit/c12fe3f151ee21eb9ad67561b51ea674729478d6))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/cloudfront to v1.22.2 ([#5710](https://github.com/cloudquery/cloudquery/issues/5710)) ([5b79177](https://github.com/cloudquery/cloudquery/commit/5b79177666d6016f22a06e2b657042ca83521132))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/cloudhsmv2 to v1.13.23 ([#5703](https://github.com/cloudquery/cloudquery/issues/5703)) ([cbef456](https://github.com/cloudquery/cloudquery/commit/cbef456a65cd0dd1f70c126ce12ea8343b50479b))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/cloudtrail to v1.21.0 ([#5631](https://github.com/cloudquery/cloudquery/issues/5631)) ([ea6ce38](https://github.com/cloudquery/cloudquery/commit/ea6ce381d9f2da05966be1990855d8339f1f92b1))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/cloudtrail to v1.21.1 ([#5704](https://github.com/cloudquery/cloudquery/issues/5704)) ([609b3c7](https://github.com/cloudquery/cloudquery/commit/609b3c7cde08508efb68e3a24eb083a2aeafdffb))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/cloudwatch to v1.23.0 ([#5667](https://github.com/cloudquery/cloudquery/issues/5667)) ([a2ecabc](https://github.com/cloudquery/cloudquery/commit/a2ecabc5b42f1e3b3835ca95f54699ec71fac4e6))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/cloudwatch to v1.23.1 ([#5705](https://github.com/cloudquery/cloudquery/issues/5705)) ([2674a72](https://github.com/cloudquery/cloudquery/commit/2674a72aa6a4b7b20f230fb7bde46124c5ac111b))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs to v1.17.3 ([#5711](https://github.com/cloudquery/cloudquery/issues/5711)) ([aa9eb9c](https://github.com/cloudquery/cloudquery/commit/aa9eb9cdc3fa93009a0651424598b20ee52a4cd7))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/codebuild to v1.19.21 ([#5712](https://github.com/cloudquery/cloudquery/issues/5712)) ([3d5ef99](https://github.com/cloudquery/cloudquery/commit/3d5ef99f814cdc70e553f24e7f38e4856f7af8d3))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/codepipeline to v1.13.21 ([#5713](https://github.com/cloudquery/cloudquery/issues/5713)) ([cf58ee4](https://github.com/cloudquery/cloudquery/commit/cf58ee400db38f9aba3dd2a87e380cad50524acb))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/cognitoidentity to v1.14.6 ([#5714](https://github.com/cloudquery/cloudquery/issues/5714)) ([5f5df9f](https://github.com/cloudquery/cloudquery/commit/5f5df9ff80721f7469a0c5c34c719ea1769f94b0))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider to v1.21.4 ([#5715](https://github.com/cloudquery/cloudquery/issues/5715)) ([2d8535d](https://github.com/cloudquery/cloudquery/commit/2d8535d7dc9dc95b1741775f5a141270ba946500))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/configservice to v1.28.2 ([#5716](https://github.com/cloudquery/cloudquery/issues/5716)) ([6d6411a](https://github.com/cloudquery/cloudquery/commit/6d6411abae4f413c5f44ba185c6d63e77fc52aaa))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/databasemigrationservice to v1.22.3 ([#5717](https://github.com/cloudquery/cloudquery/issues/5717)) ([3e1046b](https://github.com/cloudquery/cloudquery/commit/3e1046ba6c3503fca330858761c39073d72543ca))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/dax to v1.11.21 ([#5718](https://github.com/cloudquery/cloudquery/issues/5718)) ([e1db9ce](https://github.com/cloudquery/cloudquery/commit/e1db9ce91255583a575ef92531f6e428af8b6609))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/directconnect to v1.17.25 ([#5719](https://github.com/cloudquery/cloudquery/issues/5719)) ([b6ffd5a](https://github.com/cloudquery/cloudquery/commit/b6ffd5a50722b02ce71bc058554cf70602cd81ea))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/docdb to v1.19.18 ([#5720](https://github.com/cloudquery/cloudquery/issues/5720)) ([d24c2ea](https://github.com/cloudquery/cloudquery/commit/d24c2eaf9fecf344b53902fabf0f265dea9e067d))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/dynamodb to v1.17.9 ([#5721](https://github.com/cloudquery/cloudquery/issues/5721)) ([ee536b7](https://github.com/cloudquery/cloudquery/commit/ee536b7b69c22a0d33f64b2f35a856c428038297))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/ec2 to v1.76.1 ([#5722](https://github.com/cloudquery/cloudquery/issues/5722)) ([a189df3](https://github.com/cloudquery/cloudquery/commit/a189df38047f673e3a2bf96ba58e7c41baccc0a6))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/ecr to v1.17.25 ([#5723](https://github.com/cloudquery/cloudquery/issues/5723)) ([1d62a32](https://github.com/cloudquery/cloudquery/commit/1d62a329144ee52e273ae5b1de82a003065302e5))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/ecrpublic to v1.13.22 ([#5724](https://github.com/cloudquery/cloudquery/issues/5724)) ([a848301](https://github.com/cloudquery/cloudquery/commit/a848301709484602e370050b6cb9b764d2b040bf))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/ecs to v1.22.0 ([#5814](https://github.com/cloudquery/cloudquery/issues/5814)) ([c7ac305](https://github.com/cloudquery/cloudquery/commit/c7ac305ca98d348bba41d75ad3e918d71324d272))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/efs to v1.18.3 ([#5725](https://github.com/cloudquery/cloudquery/issues/5725)) ([12c183f](https://github.com/cloudquery/cloudquery/commit/12c183f03911399f91832823601e534037264397))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/eks to v1.26.0 ([#5815](https://github.com/cloudquery/cloudquery/issues/5815)) ([3fd4cd1](https://github.com/cloudquery/cloudquery/commit/3fd4cd1def0892cc709fb4ac9668676f30628f95))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/elasticache to v1.24.3 ([#5726](https://github.com/cloudquery/cloudquery/issues/5726)) ([5dce040](https://github.com/cloudquery/cloudquery/commit/5dce040a0bfd1bb160ab81175ee8f44d06f541e8))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/elasticbeanstalk to v1.14.22 ([#5727](https://github.com/cloudquery/cloudquery/issues/5727)) ([fdb22e6](https://github.com/cloudquery/cloudquery/commit/fdb22e6dadc52cb70cace2a76571f7ab3736a3d7))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/elasticloadbalancing to v1.14.25 ([#5728](https://github.com/cloudquery/cloudquery/issues/5728)) ([c59755c](https://github.com/cloudquery/cloudquery/commit/c59755cf6811e36b922f752400bf516664f01264))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2 to v1.18.28 ([#5729](https://github.com/cloudquery/cloudquery/issues/5729)) ([3ff161d](https://github.com/cloudquery/cloudquery/commit/3ff161d85252e36c0009ac5b634ad72a21ccf7e0))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/elasticsearchservice to v1.17.4 ([#5730](https://github.com/cloudquery/cloudquery/issues/5730)) ([dc3d54f](https://github.com/cloudquery/cloudquery/commit/dc3d54f042230994d6cf1907bb5dabd0632c0518))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/emr to v1.20.18 ([#5731](https://github.com/cloudquery/cloudquery/issues/5731)) ([2fd7043](https://github.com/cloudquery/cloudquery/commit/2fd70430d9c537fb6805f523d8d2d7e58edc0eb0))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/eventbridge to v1.16.22 ([#5732](https://github.com/cloudquery/cloudquery/issues/5732)) ([ada7cd2](https://github.com/cloudquery/cloudquery/commit/ada7cd29c5a4843d66a1dd20ba74ca7bef2862eb))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/firehose to v1.15.2 ([#5734](https://github.com/cloudquery/cloudquery/issues/5734)) ([fd324c9](https://github.com/cloudquery/cloudquery/commit/fd324c93764f0ca8f2a595508c9721a368addf7d))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/frauddetector to v1.20.14 ([#5735](https://github.com/cloudquery/cloudquery/issues/5735)) ([702266a](https://github.com/cloudquery/cloudquery/commit/702266a8e08184dca35b56243c8faaeff896ee13))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/fsx to v1.26.2 ([#5736](https://github.com/cloudquery/cloudquery/issues/5736)) ([f3769ea](https://github.com/cloudquery/cloudquery/commit/f3769ea1811c4fab64310313f8f4faaf9bfe946f))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/glacier to v1.13.21 ([#5737](https://github.com/cloudquery/cloudquery/issues/5737)) ([d8e9ed1](https://github.com/cloudquery/cloudquery/commit/d8e9ed17c74ef52a51ae20b98469d300334a7ab8))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/glue to v1.38.1 ([#5816](https://github.com/cloudquery/cloudquery/issues/5816)) ([4327315](https://github.com/cloudquery/cloudquery/commit/43273155675ae96bf9412ac944dafe370a4a40dc))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/guardduty to v1.16.4 ([#5630](https://github.com/cloudquery/cloudquery/issues/5630)) ([cd3754f](https://github.com/cloudquery/cloudquery/commit/cd3754f83e6d47abbebc6aea6d3d02952f640d30))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/guardduty to v1.16.6 ([#5738](https://github.com/cloudquery/cloudquery/issues/5738)) ([1f7e8d7](https://github.com/cloudquery/cloudquery/commit/1f7e8d74950b00113f05cd8a9879d8c63cfd5e52))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/iam to v1.18.25 ([#5739](https://github.com/cloudquery/cloudquery/issues/5739)) ([7360cc9](https://github.com/cloudquery/cloudquery/commit/7360cc9f3afe8ccfc0ae473d6214314e6372bd11))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/identitystore to v1.15.9 ([#5740](https://github.com/cloudquery/cloudquery/issues/5740)) ([f7ef2ca](https://github.com/cloudquery/cloudquery/commit/f7ef2cad487662ec4fdf7f3f6882d079e1771f01))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/inspector to v1.12.21 ([#5741](https://github.com/cloudquery/cloudquery/issues/5741)) ([7be6e29](https://github.com/cloudquery/cloudquery/commit/7be6e297ad48ed2c2af8c576a36265d8e65d94c8))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/inspector2 to v1.9.2 ([#5742](https://github.com/cloudquery/cloudquery/issues/5742)) ([0d560c5](https://github.com/cloudquery/cloudquery/commit/0d560c5816a5333afdf9abe660a6ae6f2a54671c))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/iot to v1.32.2 ([#5743](https://github.com/cloudquery/cloudquery/issues/5743)) ([373a46b](https://github.com/cloudquery/cloudquery/commit/373a46b6b39bc6357b9fe241c618075933f1a9f0))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/kafka to v1.18.2 ([#5744](https://github.com/cloudquery/cloudquery/issues/5744)) ([a06048e](https://github.com/cloudquery/cloudquery/commit/a06048e490635331ee70c0ea8dbc50340b3b9628))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/kinesis to v1.16.0 ([#5817](https://github.com/cloudquery/cloudquery/issues/5817)) ([063bd82](https://github.com/cloudquery/cloudquery/commit/063bd82ca3ba21edb0090bebd2bd0e6ce656d1df))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/kms to v1.19.3 ([#5666](https://github.com/cloudquery/cloudquery/issues/5666)) ([c033c81](https://github.com/cloudquery/cloudquery/commit/c033c814ad263f0fb3a77f956915c29c3ddb2d73))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/kms to v1.19.4 ([#5745](https://github.com/cloudquery/cloudquery/issues/5745)) ([0b20016](https://github.com/cloudquery/cloudquery/commit/0b20016beb51ea8595bce77d24a628729a736f4e))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/lambda to v1.26.2 ([#5746](https://github.com/cloudquery/cloudquery/issues/5746)) ([a502d72](https://github.com/cloudquery/cloudquery/commit/a502d72ca2eb9572ad0d907840b1324f6552afe0))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/lightsail to v1.24.2 ([#5747](https://github.com/cloudquery/cloudquery/issues/5747)) ([94dbd2b](https://github.com/cloudquery/cloudquery/commit/94dbd2ba7bb2f7f964809f90286c779ab16a58c6))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/mq to v1.13.18 ([#5748](https://github.com/cloudquery/cloudquery/issues/5748)) ([d544794](https://github.com/cloudquery/cloudquery/commit/d54479403dab2519fa9c8f44fbea6d238fd327b4))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/mwaa to v1.13.14 ([#5749](https://github.com/cloudquery/cloudquery/issues/5749)) ([674ef9b](https://github.com/cloudquery/cloudquery/commit/674ef9bbb7bb76fd8a20e6da2d0b29472e09ddd2))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/neptune to v1.18.5 ([#5750](https://github.com/cloudquery/cloudquery/issues/5750)) ([b625e9e](https://github.com/cloudquery/cloudquery/commit/b625e9ecb3ca74ea73450a8f279e828304a64f5b))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/organizations to v1.17.2 ([#5751](https://github.com/cloudquery/cloudquery/issues/5751)) ([d51a8b2](https://github.com/cloudquery/cloudquery/commit/d51a8b2f256bd6da3bb4c7cfbfc8d5b9454131bf))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/qldb to v1.14.22 ([#5752](https://github.com/cloudquery/cloudquery/issues/5752)) ([37e3d83](https://github.com/cloudquery/cloudquery/commit/37e3d830b52a00cef3cf5ae43261a9603a155d3f))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/quicksight to v1.28.3 ([#5753](https://github.com/cloudquery/cloudquery/issues/5753)) ([68cac4a](https://github.com/cloudquery/cloudquery/commit/68cac4a0b41ff1d013c1d6e74a77f64b8e5ccdff))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/ram to v1.16.26 ([#5754](https://github.com/cloudquery/cloudquery/issues/5754)) ([532b7d6](https://github.com/cloudquery/cloudquery/commit/532b7d6b0ec75d818648ce07a3c24fd15e7cba9f))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/rds to v1.35.0 ([#5632](https://github.com/cloudquery/cloudquery/issues/5632)) ([98b796a](https://github.com/cloudquery/cloudquery/commit/98b796a9029ac5155d9cd93e13e775dd02b26eab))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/rds to v1.35.1 ([#5755](https://github.com/cloudquery/cloudquery/issues/5755)) ([f4502d4](https://github.com/cloudquery/cloudquery/commit/f4502d45117559e6e5ea5271d7698e5eba47d659))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/rds to v1.36.0 ([#5818](https://github.com/cloudquery/cloudquery/issues/5818)) ([b3ac1a7](https://github.com/cloudquery/cloudquery/commit/b3ac1a7dcfff5899fd837169ddb5d401af8e1c3f))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/redshift to v1.26.18 ([#5763](https://github.com/cloudquery/cloudquery/issues/5763)) ([f49fdee](https://github.com/cloudquery/cloudquery/commit/f49fdee20bd5054476e9ed07d5761e124cc82b29))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/resourcegroups to v1.12.23 ([#5764](https://github.com/cloudquery/cloudquery/issues/5764)) ([6d273d5](https://github.com/cloudquery/cloudquery/commit/6d273d508f8e63524bff7d5841156dd9ce151fb2))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/route53 to v1.25.2 ([#5765](https://github.com/cloudquery/cloudquery/issues/5765)) ([8c703d3](https://github.com/cloudquery/cloudquery/commit/8c703d31584b038ca303b4e1a4d3267b940d93ff))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/route53domains to v1.12.21 ([#5766](https://github.com/cloudquery/cloudquery/issues/5766)) ([bc28931](https://github.com/cloudquery/cloudquery/commit/bc28931417c882c78c1195038479100fc23c8c0e))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/s3control to v1.28.2 ([#5768](https://github.com/cloudquery/cloudquery/issues/5768)) ([e3e15ab](https://github.com/cloudquery/cloudquery/commit/e3e15ab252998b5b56776b5b32c32ef05e3d8b09))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/scheduler to v1.0.3 ([#5769](https://github.com/cloudquery/cloudquery/issues/5769)) ([9dab834](https://github.com/cloudquery/cloudquery/commit/9dab8349c34c31898d468606d68a68748a148c7e))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/secretsmanager to v1.16.10 ([#5770](https://github.com/cloudquery/cloudquery/issues/5770)) ([03f1784](https://github.com/cloudquery/cloudquery/commit/03f17841a2ccd5f21d8afd82f8874bfcfcd971cf))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/servicecatalog to v1.15.2 ([#5771](https://github.com/cloudquery/cloudquery/issues/5771)) ([bfc8e7c](https://github.com/cloudquery/cloudquery/commit/bfc8e7c607950dccb89d4966f70d0afe47fc2402))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/servicecatalogappregistry to v1.15.2 ([#5772](https://github.com/cloudquery/cloudquery/issues/5772)) ([dd57905](https://github.com/cloudquery/cloudquery/commit/dd5790583b7e62a300ce4aabc90923c0143c2d09))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/servicequotas to v1.13.23 ([#5773](https://github.com/cloudquery/cloudquery/issues/5773)) ([18c0bf8](https://github.com/cloudquery/cloudquery/commit/18c0bf857ea9dba4b0d2fdfb109347b40f5cc58e))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/sesv2 to v1.15.3 ([#5774](https://github.com/cloudquery/cloudquery/issues/5774)) ([0e513d3](https://github.com/cloudquery/cloudquery/commit/0e513d3f0f7fb49b880c0011461891e38eec91f5))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/sfn to v1.16.2 ([#5775](https://github.com/cloudquery/cloudquery/issues/5775)) ([5a6b8ac](https://github.com/cloudquery/cloudquery/commit/5a6b8ac728edba4a33155da574d45fd1fcbe1711))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/shield to v1.17.13 ([#5776](https://github.com/cloudquery/cloudquery/issues/5776)) ([771ecd6](https://github.com/cloudquery/cloudquery/commit/771ecd6d7eec63c8d0f6ac0ec515376672f288af))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/sns to v1.18.8 ([#5777](https://github.com/cloudquery/cloudquery/issues/5777)) ([44952a1](https://github.com/cloudquery/cloudquery/commit/44952a1a4d87bee9623c0346693b51f0c0cff2bc))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/sqs to v1.19.17 ([#5778](https://github.com/cloudquery/cloudquery/issues/5778)) ([c2f7d0e](https://github.com/cloudquery/cloudquery/commit/c2f7d0e32e726a7152f6eeb222fa333c941bd5d8))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/ssm to v1.33.3 ([#5779](https://github.com/cloudquery/cloudquery/issues/5779)) ([f182bc2](https://github.com/cloudquery/cloudquery/commit/f182bc28e1719d29ae053b3b69881ba781265ba1))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/ssoadmin to v1.15.15 ([#5780](https://github.com/cloudquery/cloudquery/issues/5780)) ([d802734](https://github.com/cloudquery/cloudquery/commit/d802734ff54c41de9fc291e86439bd77e41737bc))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/timestreamwrite to v1.14.8 ([#5781](https://github.com/cloudquery/cloudquery/issues/5781)) ([0e4284d](https://github.com/cloudquery/cloudquery/commit/0e4284d890f8efd1ebe982a5fb2b3b790f3142ab))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/transfer to v1.25.2 ([#5782](https://github.com/cloudquery/cloudquery/issues/5782)) ([ac48bb4](https://github.com/cloudquery/cloudquery/commit/ac48bb44f91dd9933ab264eab5f49110728db6c3))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/waf to v1.11.21 ([#5783](https://github.com/cloudquery/cloudquery/issues/5783)) ([4917f32](https://github.com/cloudquery/cloudquery/commit/4917f32c156ce5b26ea6b0f6aa19d93a14f49f5f))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/wafregional to v1.12.22 ([#5784](https://github.com/cloudquery/cloudquery/issues/5784)) ([df15d53](https://github.com/cloudquery/cloudquery/commit/df15d53c52540f4e5ca1511c5355aa3a3f95ac96))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/wafv2 to v1.23.4 ([#5785](https://github.com/cloudquery/cloudquery/issues/5785)) ([0bb05c0](https://github.com/cloudquery/cloudquery/commit/0bb05c064da505de0798745d181742e47cd1d73c))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/workspaces to v1.27.2 ([#5786](https://github.com/cloudquery/cloudquery/issues/5786)) ([9d8af2f](https://github.com/cloudquery/cloudquery/commit/9d8af2f04f60c72d39d191d2732f33e337896e40))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/xray to v1.15.3 ([#5787](https://github.com/cloudquery/cloudquery/issues/5787)) ([a8076ef](https://github.com/cloudquery/cloudquery/commit/a8076efc7df7b188e9e2f2832c88c09d9e57654d))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.12.2 ([#5583](https://github.com/cloudquery/cloudquery/issues/5583)) ([d721c4e](https://github.com/cloudquery/cloudquery/commit/d721c4e06b8a97b5373215aca0e4ed64942ac489))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.12.3 ([#5639](https://github.com/cloudquery/cloudquery/issues/5639)) ([6452d0e](https://github.com/cloudquery/cloudquery/commit/6452d0ed5a44abad9d7530af6e79cde6504d0c4c))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.12.4 ([#5649](https://github.com/cloudquery/cloudquery/issues/5649)) ([b4aa889](https://github.com/cloudquery/cloudquery/commit/b4aa889e396db3b0887d1684e4bc07da6050af43))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.12.5 ([#5661](https://github.com/cloudquery/cloudquery/issues/5661)) ([b354b8a](https://github.com/cloudquery/cloudquery/commit/b354b8a3683fa2bc918c1002afac487427d65a5f))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.12.6 ([#5790](https://github.com/cloudquery/cloudquery/issues/5790)) ([8e2663c](https://github.com/cloudquery/cloudquery/commit/8e2663c17c3347afd5e53f665462adc3e709c96c))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.12.7 ([#5797](https://github.com/cloudquery/cloudquery/issues/5797)) ([15da529](https://github.com/cloudquery/cloudquery/commit/15da5294786fa2656228ca5bbc48ef1fc44e486b))
* Update endpoints ([#5628](https://github.com/cloudquery/cloudquery/issues/5628)) ([b117530](https://github.com/cloudquery/cloudquery/commit/b1175301e95ad81b800cd0190acdcea56de534eb))
* Update endpoints ([#5664](https://github.com/cloudquery/cloudquery/issues/5664)) ([1cc7d60](https://github.com/cloudquery/cloudquery/commit/1cc7d60a4f8ad91ebf32ab8ba77ca7a6e2f49e57))
* Update endpoints ([#5682](https://github.com/cloudquery/cloudquery/issues/5682)) ([c7c79a2](https://github.com/cloudquery/cloudquery/commit/c7c79a28509fb5436af19efe8a2055cd46c0e594))
* Update endpoints ([#5709](https://github.com/cloudquery/cloudquery/issues/5709)) ([3d65725](https://github.com/cloudquery/cloudquery/commit/3d65725922ae5b950075052d995c59bf7b568ec3))
* Update endpoints ([#5808](https://github.com/cloudquery/cloudquery/issues/5808)) ([d93f88c](https://github.com/cloudquery/cloudquery/commit/d93f88cb6efdb5e5f668f87da6a01153130c7485))

## [7.4.0](https://github.com/cloudquery/cloudquery/compare/plugins-source-aws-v7.3.2...plugins-source-aws-v7.4.0) (2022-12-13)


### Features

* Add more cases where DateTime types in CSV responses of AWS IAM… ([#5526](https://github.com/cloudquery/cloudquery/issues/5526)) ([a43bf39](https://github.com/cloudquery/cloudquery/commit/a43bf39f0ad12b93786518a7f304aa254febe5ce))
* **aws:** Add IAM SSH Public Keys ([#5538](https://github.com/cloudquery/cloudquery/issues/5538)) ([5bd2b4d](https://github.com/cloudquery/cloudquery/commit/5bd2b4d09f975bf06f5f35b0b601534dc95e4eb8))


### Bug Fixes

* **aws-functions:** Return error in case of access denied ([#5537](https://github.com/cloudquery/cloudquery/issues/5537)) ([1fe79d3](https://github.com/cloudquery/cloudquery/commit/1fe79d38c1d877031e4ef95e289616c8b24d449e))
* **aws:** Deterministic Multiplexer ([#5513](https://github.com/cloudquery/cloudquery/issues/5513)) ([eaf6e99](https://github.com/cloudquery/cloudquery/commit/eaf6e99b53deffe681d5d03d9753023f95c1d485))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/appstream to v1.18.2 ([#5566](https://github.com/cloudquery/cloudquery/issues/5566)) ([e7b7c24](https://github.com/cloudquery/cloudquery/commit/e7b7c24cb7bde369ca1bbe568d94a2a024b7d80b))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/athena to v1.20.2 ([#5523](https://github.com/cloudquery/cloudquery/issues/5523)) ([9385a69](https://github.com/cloudquery/cloudquery/commit/9385a69ad93b6221ced9c21c299d5fe7153306ad))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/autoscaling to v1.25.0 ([#5524](https://github.com/cloudquery/cloudquery/issues/5524)) ([5009a35](https://github.com/cloudquery/cloudquery/commit/5009a3584ec51c2d5cc87fd0efc3f8517078b41e))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs to v1.17.2 ([#5535](https://github.com/cloudquery/cloudquery/issues/5535)) ([2ce06fc](https://github.com/cloudquery/cloudquery/commit/2ce06fc3f47e8842ba5a1c7b44a87c49564a24ca))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/ec2 to v1.76.0 ([#5569](https://github.com/cloudquery/cloudquery/issues/5569)) ([979a780](https://github.com/cloudquery/cloudquery/commit/979a780d030e9a36d4ecd605b492a5a0c4e90cbd))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/mq to v1.13.17 ([#5567](https://github.com/cloudquery/cloudquery/issues/5567)) ([6075e71](https://github.com/cloudquery/cloudquery/commit/6075e7149d3779a488afe794f61f26fae91ad436))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/rds to v1.34.0 ([#5570](https://github.com/cloudquery/cloudquery/issues/5570)) ([8fd69ba](https://github.com/cloudquery/cloudquery/commit/8fd69ba50d7cef86b3f87ff40537d408b7ede96f))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/wafv2 to v1.23.3 ([#5568](https://github.com/cloudquery/cloudquery/issues/5568)) ([1782f00](https://github.com/cloudquery/cloudquery/commit/1782f00b045758864cc47408d624b9aeac93a714))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.12.0 ([#5539](https://github.com/cloudquery/cloudquery/issues/5539)) ([fb71293](https://github.com/cloudquery/cloudquery/commit/fb71293d5cfe1b2ef32ba83d604ac3c48e662bce))
* Update endpoints ([#5520](https://github.com/cloudquery/cloudquery/issues/5520)) ([88a88a7](https://github.com/cloudquery/cloudquery/commit/88a88a71dc69c88f4706c1ad31e69d0ecf9be332))
* Update endpoints ([#5565](https://github.com/cloudquery/cloudquery/issues/5565)) ([6b8c15d](https://github.com/cloudquery/cloudquery/commit/6b8c15dfd351c31f1fa79344faad5937c3b3bf69))

## [7.3.2](https://github.com/cloudquery/cloudquery/compare/plugins-source-aws-v7.3.1...plugins-source-aws-v7.3.2) (2022-12-08)


### Bug Fixes

* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/cloudfront to v1.22.0 ([#5495](https://github.com/cloudquery/cloudquery/issues/5495)) ([637bde2](https://github.com/cloudquery/cloudquery/commit/637bde2ab06926e5dfd7023b1c9de95e25bff4b5))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider to v1.21.3 ([#5493](https://github.com/cloudquery/cloudquery/issues/5493)) ([56bc65a](https://github.com/cloudquery/cloudquery/commit/56bc65aac3967d4c5557832cc6165d5877ca0501))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/eks to v1.25.0 ([#5496](https://github.com/cloudquery/cloudquery/issues/5496)) ([fb7bb80](https://github.com/cloudquery/cloudquery/commit/fb7bb8090888e558532341de07cefe65a801615c))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/kms to v1.19.2 ([#5494](https://github.com/cloudquery/cloudquery/issues/5494)) ([f84804b](https://github.com/cloudquery/cloudquery/commit/f84804b3a266ef1e7634cf02c3734dee33647bda))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.11.2 ([#5497](https://github.com/cloudquery/cloudquery/issues/5497)) ([c1876cf](https://github.com/cloudquery/cloudquery/commit/c1876cf793b43d825a25fb3c9ba4996e4b09964f))
* Update endpoints ([#5489](https://github.com/cloudquery/cloudquery/issues/5489)) ([dc8255c](https://github.com/cloudquery/cloudquery/commit/dc8255c50d65d5ce59512142c21b6f8f344988a2))

## [7.3.1](https://github.com/cloudquery/cloudquery/compare/plugins-source-aws-v7.3.0...plugins-source-aws-v7.3.1) (2022-12-07)


### Bug Fixes

* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/cognitoidentity to v1.14.5 ([#5454](https://github.com/cloudquery/cloudquery/issues/5454)) ([b671f08](https://github.com/cloudquery/cloudquery/commit/b671f08ec6c3f86255fecbad0db707c28ff2044e))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/rds to v1.33.0 ([#5455](https://github.com/cloudquery/cloudquery/issues/5455)) ([00331c2](https://github.com/cloudquery/cloudquery/commit/00331c22be15fc27d0c908444381b75923f90571))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.11.1 ([#5458](https://github.com/cloudquery/cloudquery/issues/5458)) ([58b7432](https://github.com/cloudquery/cloudquery/commit/58b74321cd253c9a843c8c103f324abb93952195))
* Update endpoints ([#5453](https://github.com/cloudquery/cloudquery/issues/5453)) ([6f610fd](https://github.com/cloudquery/cloudquery/commit/6f610fd91f6d85a8c64aca2a442701644178c8ad))

## [7.3.0](https://github.com/cloudquery/cloudquery/compare/plugins-source-aws-v7.2.3...plugins-source-aws-v7.3.0) (2022-12-06)


### Features

* **aws:** Add inline_policy field to permission_sets table ([#5280](https://github.com/cloudquery/cloudquery/issues/5280)) ([bce8bb7](https://github.com/cloudquery/cloudquery/commit/bce8bb798a08d21095457cab4171b37bbc72de3c))
* **website:** Add plugins tables ([#5259](https://github.com/cloudquery/cloudquery/issues/5259)) ([c336f4e](https://github.com/cloudquery/cloudquery/commit/c336f4e25e192ffdd4c211d4a35b67b71d01d1f8))


### Bug Fixes

* AWS Endpoint Generation ([#5268](https://github.com/cloudquery/cloudquery/issues/5268)) ([9cf87c4](https://github.com/cloudquery/cloudquery/commit/9cf87c4ea0241ae56496b74f7724321d08de6c63))
* **aws:** Unused EIP policy - check associations not instance_ids ([#5378](https://github.com/cloudquery/cloudquery/issues/5378)) ([79907db](https://github.com/cloudquery/cloudquery/commit/79907dba2d9773770c9bae06405ff316f02e3d0c))
* **deps:** Update github.com/gocarina/gocsv digest to c8ef781 ([#5199](https://github.com/cloudquery/cloudquery/issues/5199)) ([9a4249c](https://github.com/cloudquery/cloudquery/commit/9a4249ce316f2aa0934d1a6f9132d0590919589e))
* **deps:** Update module github.com/aws/aws-sdk-go-v2 to v1.17.2 ([#5297](https://github.com/cloudquery/cloudquery/issues/5297)) ([001999e](https://github.com/cloudquery/cloudquery/commit/001999ee890f7ec11d3d422c3a130346f5fb6491))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/config to v1.18.4 ([#5298](https://github.com/cloudquery/cloudquery/issues/5298)) ([550558b](https://github.com/cloudquery/cloudquery/commit/550558b69fba25cfc5cb7767c6039aa1715b3564))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/feature/s3/manager to v1.11.43 ([#5300](https://github.com/cloudquery/cloudquery/issues/5300)) ([934d3a2](https://github.com/cloudquery/cloudquery/commit/934d3a2ef8e3d68cc60000ad35e2dd748b530563))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/accessanalyzer to v1.18.0 ([#5214](https://github.com/cloudquery/cloudquery/issues/5214)) ([a2a1636](https://github.com/cloudquery/cloudquery/commit/a2a1636c865aee75260d3e3c169ee3a04999b80d))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/accessanalyzer to v1.18.1 ([#5301](https://github.com/cloudquery/cloudquery/issues/5301)) ([f2ff665](https://github.com/cloudquery/cloudquery/commit/f2ff6651e8162bae9d7647b317f54f5afb34c91e))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/account to v1.7.12 ([#5302](https://github.com/cloudquery/cloudquery/issues/5302)) ([3c4b649](https://github.com/cloudquery/cloudquery/commit/3c4b64932cb868f8af45847c58da0eacb0c2cd43))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/acm to v1.16.4 ([#5303](https://github.com/cloudquery/cloudquery/issues/5303)) ([33373a8](https://github.com/cloudquery/cloudquery/commit/33373a8c2485411254af8c3baf88bc31a193ca5d))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/apigateway to v1.15.26 ([#5304](https://github.com/cloudquery/cloudquery/issues/5304)) ([99928c3](https://github.com/cloudquery/cloudquery/commit/99928c362a38c1b0a3ad1a2efd059363d35fce56))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/apigatewayv2 to v1.12.24 ([#5305](https://github.com/cloudquery/cloudquery/issues/5305)) ([5ad584c](https://github.com/cloudquery/cloudquery/commit/5ad584c58d080982cd4a67ce8f0f1e7390f76e6f))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/applicationautoscaling to v1.15.24 ([#5306](https://github.com/cloudquery/cloudquery/issues/5306)) ([d30723b](https://github.com/cloudquery/cloudquery/commit/d30723bd894f4df0e3f8088250942dbdbe8ce4e5))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/apprunner to v1.15.1 ([#5307](https://github.com/cloudquery/cloudquery/issues/5307)) ([28aadc7](https://github.com/cloudquery/cloudquery/commit/28aadc719bc04711a65f8cbd7da7513f3a787a92))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/appstream to v1.18.1 ([#5308](https://github.com/cloudquery/cloudquery/issues/5308)) ([8ec0779](https://github.com/cloudquery/cloudquery/commit/8ec077958c1c299f557cf7981913c773288e9ad4))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/appsync to v1.17.0 ([#5411](https://github.com/cloudquery/cloudquery/issues/5411)) ([5453ffa](https://github.com/cloudquery/cloudquery/commit/5453ffa6f84e683d4e5f252a032a28ef22145ddd))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/athena to v1.20.0 ([#5215](https://github.com/cloudquery/cloudquery/issues/5215)) ([fcecd6e](https://github.com/cloudquery/cloudquery/commit/fcecd6eaa9c033f624d19827d2a670d787b8ffa4))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/athena to v1.20.1 ([#5309](https://github.com/cloudquery/cloudquery/issues/5309)) ([69fdaef](https://github.com/cloudquery/cloudquery/commit/69fdaef5029add8bc68b8f7baa5aaed82ff5b204))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/autoscaling to v1.24.4 ([#5310](https://github.com/cloudquery/cloudquery/issues/5310)) ([62ef3e9](https://github.com/cloudquery/cloudquery/commit/62ef3e91ca6b3ab2bc6e20cd20161e8143cc2963))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/backup to v1.18.1 ([#5311](https://github.com/cloudquery/cloudquery/issues/5311)) ([4bbbfb5](https://github.com/cloudquery/cloudquery/commit/4bbbfb5b498df3304c6fc18c479fecccbe607d04))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/cloudformation to v1.24.2 ([#5312](https://github.com/cloudquery/cloudquery/issues/5312)) ([661bb44](https://github.com/cloudquery/cloudquery/commit/661bb44e264a78cdebf59b22c1a0681ed4f94efd))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/cloudfront to v1.21.1 ([#5313](https://github.com/cloudquery/cloudquery/issues/5313)) ([8399d58](https://github.com/cloudquery/cloudquery/commit/8399d58e4685759f67d84558c38e6efd88422fc8))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/cloudhsmv2 to v1.13.22 ([#5314](https://github.com/cloudquery/cloudquery/issues/5314)) ([bcdeb31](https://github.com/cloudquery/cloudquery/commit/bcdeb313dba6de34302d1304168ac602004ab202))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/cloudtrail to v1.20.4 ([#5315](https://github.com/cloudquery/cloudquery/issues/5315)) ([645c497](https://github.com/cloudquery/cloudquery/commit/645c497aa3c2318ea5f06860a0fd8a636060bf95))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/cloudwatch to v1.22.1 ([#5316](https://github.com/cloudquery/cloudquery/issues/5316)) ([268d434](https://github.com/cloudquery/cloudquery/commit/268d434d07a863e6f8be82ff2e4bf99218bdf402))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs to v1.17.1 ([#5317](https://github.com/cloudquery/cloudquery/issues/5317)) ([73cd9a8](https://github.com/cloudquery/cloudquery/commit/73cd9a8ea230f808e1bbfe9cc93a560a3dd57a3d))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/codebuild to v1.19.20 ([#5318](https://github.com/cloudquery/cloudquery/issues/5318)) ([fa38434](https://github.com/cloudquery/cloudquery/commit/fa38434dfe838ead8a0de5394ed9935546be12ff))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/codepipeline to v1.13.20 ([#5319](https://github.com/cloudquery/cloudquery/issues/5319)) ([bb672ae](https://github.com/cloudquery/cloudquery/commit/bb672ae912c5f795f47af477eb7dd7021467c4d2))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/cognitoidentity to v1.14.4 ([#5320](https://github.com/cloudquery/cloudquery/issues/5320)) ([d4b5f0d](https://github.com/cloudquery/cloudquery/commit/d4b5f0d7c9c95747456b672ad4071f78bb138f45))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider to v1.21.2 ([#5321](https://github.com/cloudquery/cloudquery/issues/5321)) ([f0f2eaf](https://github.com/cloudquery/cloudquery/commit/f0f2eaf8dbf8a809ae684d5071eeaaed06b21b65))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/configservice to v1.28.1 ([#5322](https://github.com/cloudquery/cloudquery/issues/5322)) ([aa4a29f](https://github.com/cloudquery/cloudquery/commit/aa4a29f8f00d6a9c78b1be252e4fd050bfcc16cc))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/databasemigrationservice to v1.22.2 ([#5323](https://github.com/cloudquery/cloudquery/issues/5323)) ([cdc9021](https://github.com/cloudquery/cloudquery/commit/cdc902129f8ea9e28192a8f8af9a8b78db2999aa))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/dax to v1.11.20 ([#5324](https://github.com/cloudquery/cloudquery/issues/5324)) ([6cc4eb0](https://github.com/cloudquery/cloudquery/commit/6cc4eb0483e44f1c2387c9e157f342607b934910))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/directconnect to v1.17.24 ([#5325](https://github.com/cloudquery/cloudquery/issues/5325)) ([6444aa9](https://github.com/cloudquery/cloudquery/commit/6444aa99b1a2728a9ff6cde2380171a220eb45b8))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/docdb to v1.19.17 ([#5326](https://github.com/cloudquery/cloudquery/issues/5326)) ([3d3bf10](https://github.com/cloudquery/cloudquery/commit/3d3bf10f5a4c1112b906fefc5cf9f0dd8f28bb19))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/dynamodb to v1.17.8 ([#5327](https://github.com/cloudquery/cloudquery/issues/5327)) ([7f6cd8a](https://github.com/cloudquery/cloudquery/commit/7f6cd8abcc65237e23c41122c95591477b109345))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/ec2 to v1.74.1 ([#5328](https://github.com/cloudquery/cloudquery/issues/5328)) ([190ffff](https://github.com/cloudquery/cloudquery/commit/190ffff0f667a896146606f3dfdb4ddb3b22ee9a))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/ec2 to v1.75.0 ([#5412](https://github.com/cloudquery/cloudquery/issues/5412)) ([a652b73](https://github.com/cloudquery/cloudquery/commit/a652b732cabe5a4ecc510c79742a027c620be0ff))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/ecr to v1.17.24 ([#5329](https://github.com/cloudquery/cloudquery/issues/5329)) ([0c208f2](https://github.com/cloudquery/cloudquery/commit/0c208f2e465f56ad6895b0c3757b137a43a5def6))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/ecrpublic to v1.13.20 ([#5203](https://github.com/cloudquery/cloudquery/issues/5203)) ([b3ed262](https://github.com/cloudquery/cloudquery/commit/b3ed262ce84ef524499c1a777465bdbf49ccbda2))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/ecrpublic to v1.13.21 ([#5330](https://github.com/cloudquery/cloudquery/issues/5330)) ([15a24bc](https://github.com/cloudquery/cloudquery/commit/15a24bc659fa2c18d1e69f05f8749a65d0f0129c))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/ecs to v1.20.1 ([#5331](https://github.com/cloudquery/cloudquery/issues/5331)) ([79da6fc](https://github.com/cloudquery/cloudquery/commit/79da6fce1238ce61f51d7667c2e4d8ef644084b8))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/efs to v1.18.1 ([#5332](https://github.com/cloudquery/cloudquery/issues/5332)) ([ed58abb](https://github.com/cloudquery/cloudquery/commit/ed58abb1a29562d9a409e8d24730382ea6c98e97))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/eks to v1.24.1 ([#5333](https://github.com/cloudquery/cloudquery/issues/5333)) ([e085850](https://github.com/cloudquery/cloudquery/commit/e0858501ccb3d72ce1ea624acead6d766ef372ee))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/elasticache to v1.24.2 ([#5334](https://github.com/cloudquery/cloudquery/issues/5334)) ([eab93ed](https://github.com/cloudquery/cloudquery/commit/eab93ed4673b6408ecb3ef07a040210807ecb7f1))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/elasticbeanstalk to v1.14.21 ([#5335](https://github.com/cloudquery/cloudquery/issues/5335)) ([60d3eb6](https://github.com/cloudquery/cloudquery/commit/60d3eb6d52d29c40dc65061c2a0f1b2035b26793))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/elasticloadbalancing to v1.14.24 ([#5336](https://github.com/cloudquery/cloudquery/issues/5336)) ([b28fa9d](https://github.com/cloudquery/cloudquery/commit/b28fa9daefa0a484f51011c3cc8dcf8a39fe52b3))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2 to v1.18.27 ([#5337](https://github.com/cloudquery/cloudquery/issues/5337)) ([9c774bb](https://github.com/cloudquery/cloudquery/commit/9c774bb4019dd8a1a318fa67884fc0637ed89f38))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/elasticsearchservice to v1.17.3 ([#5338](https://github.com/cloudquery/cloudquery/issues/5338)) ([11aca3b](https://github.com/cloudquery/cloudquery/commit/11aca3b879e013fc55cdf073b1b5a500ed41e33e))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/emr to v1.20.17 ([#5339](https://github.com/cloudquery/cloudquery/issues/5339)) ([127add0](https://github.com/cloudquery/cloudquery/commit/127add01b5554790b6e9af6dbb18b5298ffbe54c))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/eventbridge to v1.16.21 ([#5340](https://github.com/cloudquery/cloudquery/issues/5340)) ([d68e79e](https://github.com/cloudquery/cloudquery/commit/d68e79e1503d8b0cc4d9d9d29e3b1b21019e9e13))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/firehose to v1.15.1 ([#5341](https://github.com/cloudquery/cloudquery/issues/5341)) ([c6f13d0](https://github.com/cloudquery/cloudquery/commit/c6f13d06960531a55b3611bab183f886eb19a636))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/frauddetector to v1.20.13 ([#5342](https://github.com/cloudquery/cloudquery/issues/5342)) ([63f5e5f](https://github.com/cloudquery/cloudquery/commit/63f5e5fbfc178833aa2e1a063657d4bda122c32e))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/fsx to v1.26.1 ([#5343](https://github.com/cloudquery/cloudquery/issues/5343)) ([f4e012a](https://github.com/cloudquery/cloudquery/commit/f4e012a336425441c383b7204422a3b5d0042d40))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/glacier to v1.13.20 ([#5344](https://github.com/cloudquery/cloudquery/issues/5344)) ([8122416](https://github.com/cloudquery/cloudquery/commit/812241697c644bdb1ae202bbadcb3baae456f788))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/glue to v1.37.0 ([#5216](https://github.com/cloudquery/cloudquery/issues/5216)) ([a2f9b30](https://github.com/cloudquery/cloudquery/commit/a2f9b302087bed6f58231b3142c7d2ec83cd78c3))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/glue to v1.37.1 ([#5351](https://github.com/cloudquery/cloudquery/issues/5351)) ([a2916c6](https://github.com/cloudquery/cloudquery/commit/a2916c642609866181e8edf6b2bbccc4fe98369e))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/guardduty to v1.16.3 ([#5352](https://github.com/cloudquery/cloudquery/issues/5352)) ([0020cfb](https://github.com/cloudquery/cloudquery/commit/0020cfb8c68967d509d3ebd7152b10357ddb9ff5))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/iam to v1.18.24 ([#5353](https://github.com/cloudquery/cloudquery/issues/5353)) ([566d7c6](https://github.com/cloudquery/cloudquery/commit/566d7c62c65ee8c8ce585ddeb3aae3f66bc5608b))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/identitystore to v1.15.8 ([#5354](https://github.com/cloudquery/cloudquery/issues/5354)) ([055b345](https://github.com/cloudquery/cloudquery/commit/055b3457dc0035b42755032db73d63345f47854d))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/inspector to v1.12.20 ([#5355](https://github.com/cloudquery/cloudquery/issues/5355)) ([8861d4b](https://github.com/cloudquery/cloudquery/commit/8861d4b1486687a14340c8d11a0614c91906bd9d))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/inspector2 to v1.9.1 ([#5356](https://github.com/cloudquery/cloudquery/issues/5356)) ([c0dda14](https://github.com/cloudquery/cloudquery/commit/c0dda144add6d0b1f6842dd83f95ad6600104fb8))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/iot to v1.32.1 ([#5357](https://github.com/cloudquery/cloudquery/issues/5357)) ([7c5d74e](https://github.com/cloudquery/cloudquery/commit/7c5d74e40843f3023413f63fae53442f927a2422))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/kafka to v1.18.1 ([#5358](https://github.com/cloudquery/cloudquery/issues/5358)) ([3aaa284](https://github.com/cloudquery/cloudquery/commit/3aaa284f3ac712d9b2a17bbf1a6075dea32dab49))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/kinesis to v1.15.25 ([#5359](https://github.com/cloudquery/cloudquery/issues/5359)) ([f6ddc88](https://github.com/cloudquery/cloudquery/commit/f6ddc882d70bc7ba349bd9cf8767d79c5aaca45b))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/kms to v1.19.1 ([#5360](https://github.com/cloudquery/cloudquery/issues/5360)) ([7868d85](https://github.com/cloudquery/cloudquery/commit/7868d85d9955063a4e01d5cafefeae8778a34d18))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/lambda to v1.26.1 ([#5361](https://github.com/cloudquery/cloudquery/issues/5361)) ([2d7b3f7](https://github.com/cloudquery/cloudquery/commit/2d7b3f7f3e9df1c67536d3cbc2c053eb44ca4565))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/lightsail to v1.24.1 ([#5362](https://github.com/cloudquery/cloudquery/issues/5362)) ([af28155](https://github.com/cloudquery/cloudquery/commit/af28155a10df7b95bd1e83f9fd0d14c2de394056))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/mq to v1.13.16 ([#5363](https://github.com/cloudquery/cloudquery/issues/5363)) ([6369466](https://github.com/cloudquery/cloudquery/commit/636946605fdb874c7eed321721d2c241af0fda86))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/mwaa to v1.13.13 ([#5364](https://github.com/cloudquery/cloudquery/issues/5364)) ([ffd29a4](https://github.com/cloudquery/cloudquery/commit/ffd29a480122ba40dd2164e880a3b5ebc6f893b9))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/neptune to v1.18.4 ([#5365](https://github.com/cloudquery/cloudquery/issues/5365)) ([898a0f5](https://github.com/cloudquery/cloudquery/commit/898a0f54868f66ab8c99de5c750a990405604f66))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/organizations to v1.17.1 ([#5366](https://github.com/cloudquery/cloudquery/issues/5366)) ([99b1514](https://github.com/cloudquery/cloudquery/commit/99b151479ad1c064d328328509a305fb9f4c373d))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/qldb to v1.14.21 ([#5367](https://github.com/cloudquery/cloudquery/issues/5367)) ([84bc23c](https://github.com/cloudquery/cloudquery/commit/84bc23c34b3c59703da9b724ae38046b6990638b))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/quicksight to v1.28.1 ([#5368](https://github.com/cloudquery/cloudquery/issues/5368)) ([4a16c77](https://github.com/cloudquery/cloudquery/commit/4a16c77328197c5d6233413ee95c01a261f15659))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/ram to v1.16.25 ([#5369](https://github.com/cloudquery/cloudquery/issues/5369)) ([eb02411](https://github.com/cloudquery/cloudquery/commit/eb024111fd373939ebef41c596c651261cf17453))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/rds to v1.31.1 ([#5370](https://github.com/cloudquery/cloudquery/issues/5370)) ([beb6bc2](https://github.com/cloudquery/cloudquery/commit/beb6bc29fef98f4c0cc18897751c1804b5dc41f1))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/rds to v1.32.0 ([#5413](https://github.com/cloudquery/cloudquery/issues/5413)) ([159ab4d](https://github.com/cloudquery/cloudquery/commit/159ab4d21e90f8836cc82bdbc5da80287649a379))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/redshift to v1.26.17 ([#5371](https://github.com/cloudquery/cloudquery/issues/5371)) ([af9f487](https://github.com/cloudquery/cloudquery/commit/af9f4874d878e5c580e8d405253b7c743e504cb6))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/resourcegroups to v1.12.22 ([#5372](https://github.com/cloudquery/cloudquery/issues/5372)) ([81c224d](https://github.com/cloudquery/cloudquery/commit/81c224d524c9ebb18d6d9edf3d0e78bff344ef52))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/route53 to v1.25.1 ([#5373](https://github.com/cloudquery/cloudquery/issues/5373)) ([fb15bcd](https://github.com/cloudquery/cloudquery/commit/fb15bcde6d9f11e3993dc351f8ea39620e18b141))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/route53domains to v1.12.20 ([#5374](https://github.com/cloudquery/cloudquery/issues/5374)) ([bb1d6fb](https://github.com/cloudquery/cloudquery/commit/bb1d6fb9b48f128a51abd8a0008589161a6704d9))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/s3control to v1.28.0 ([#5217](https://github.com/cloudquery/cloudquery/issues/5217)) ([d6412aa](https://github.com/cloudquery/cloudquery/commit/d6412aa9427b1a43f06e961c1e61143dd14d6615))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/s3control to v1.28.1 ([#5390](https://github.com/cloudquery/cloudquery/issues/5390)) ([e72a2d7](https://github.com/cloudquery/cloudquery/commit/e72a2d7b4275887b4f710255d263469409e92b45))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/sagemaker to v1.56.0 ([#5218](https://github.com/cloudquery/cloudquery/issues/5218)) ([1c9755d](https://github.com/cloudquery/cloudquery/commit/1c9755d87f3037bbb5eef1ab0de66b06da601ea8))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/sagemaker to v1.56.1 ([#5391](https://github.com/cloudquery/cloudquery/issues/5391)) ([d82c2dc](https://github.com/cloudquery/cloudquery/commit/d82c2dc609b7075b492dda79f5f785be9c408fb9))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/scheduler to v1.0.2 ([#5392](https://github.com/cloudquery/cloudquery/issues/5392)) ([50ac42f](https://github.com/cloudquery/cloudquery/commit/50ac42f145094d42d6762f84a3638559205fa3ea))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/secretsmanager to v1.16.9 ([#5393](https://github.com/cloudquery/cloudquery/issues/5393)) ([6e305bf](https://github.com/cloudquery/cloudquery/commit/6e305bff10ff00d2ad813643efed05a0cdfcc797))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/servicecatalog to v1.15.1 ([#5394](https://github.com/cloudquery/cloudquery/issues/5394)) ([730ea19](https://github.com/cloudquery/cloudquery/commit/730ea197aeb668194cd14b3a1f01a05ff41653ad))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/servicecatalogappregistry to v1.15.1 ([#5395](https://github.com/cloudquery/cloudquery/issues/5395)) ([c0b3a84](https://github.com/cloudquery/cloudquery/commit/c0b3a8410549d07e41c60247f8f637dbbebb7c94))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/servicequotas to v1.13.21 ([#5396](https://github.com/cloudquery/cloudquery/issues/5396)) ([e2ae614](https://github.com/cloudquery/cloudquery/commit/e2ae614b8a7e6675bea11554692f802e57fb51fe))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/sesv2 to v1.15.2 ([#5397](https://github.com/cloudquery/cloudquery/issues/5397)) ([26af6cc](https://github.com/cloudquery/cloudquery/commit/26af6ccd813d7b1dc554b0b63aa48947cf69e613))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/sfn to v1.16.0 ([#5284](https://github.com/cloudquery/cloudquery/issues/5284)) ([7cc803b](https://github.com/cloudquery/cloudquery/commit/7cc803b84382cc78814ecc8a274202c3194bbee5))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/sfn to v1.16.1 ([#5398](https://github.com/cloudquery/cloudquery/issues/5398)) ([17d0a70](https://github.com/cloudquery/cloudquery/commit/17d0a703b163ba64ae312d43e55b9e9382e39a2c))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/shield to v1.17.12 ([#5399](https://github.com/cloudquery/cloudquery/issues/5399)) ([816ceb8](https://github.com/cloudquery/cloudquery/commit/816ceb85a04012b638450c4cff41f38207c15daa))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/sns to v1.18.7 ([#5400](https://github.com/cloudquery/cloudquery/issues/5400)) ([d11461d](https://github.com/cloudquery/cloudquery/commit/d11461d9dbf4f4fba743edcda4666bff5fd4aa5d))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/sqs to v1.19.16 ([#5401](https://github.com/cloudquery/cloudquery/issues/5401)) ([771e3d9](https://github.com/cloudquery/cloudquery/commit/771e3d95fb25db29c0626008ebe6434e3e57e735))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/ssm to v1.33.2 ([#5402](https://github.com/cloudquery/cloudquery/issues/5402)) ([ce6b440](https://github.com/cloudquery/cloudquery/commit/ce6b440e4e7e7f411ecce861466b1523fa53e12b))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/ssoadmin to v1.15.14 ([#5403](https://github.com/cloudquery/cloudquery/issues/5403)) ([ed7475a](https://github.com/cloudquery/cloudquery/commit/ed7475a9c4db50abe6cd196f8138b4626aded5f6))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/timestreamwrite to v1.14.7 ([#5404](https://github.com/cloudquery/cloudquery/issues/5404)) ([0133015](https://github.com/cloudquery/cloudquery/commit/01330150149c15ea998d26374e89310ed8e6928f))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/transfer to v1.25.1 ([#5405](https://github.com/cloudquery/cloudquery/issues/5405)) ([4131a6e](https://github.com/cloudquery/cloudquery/commit/4131a6e713b6dd2106c1f118048600d2426cfe79))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/waf to v1.11.20 ([#5406](https://github.com/cloudquery/cloudquery/issues/5406)) ([33f2a2c](https://github.com/cloudquery/cloudquery/commit/33f2a2cfa74ee7b9c1bb7f55360cf6eea6f15cf2))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/wafregional to v1.12.21 ([#5407](https://github.com/cloudquery/cloudquery/issues/5407)) ([e725946](https://github.com/cloudquery/cloudquery/commit/e725946549b2775eeee1f5e29dbb6a534cc15d20))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/wafv2 to v1.23.2 ([#5408](https://github.com/cloudquery/cloudquery/issues/5408)) ([e11e087](https://github.com/cloudquery/cloudquery/commit/e11e08708d2ee8413916d920029c2a037c172e07))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/workspaces to v1.27.1 ([#5409](https://github.com/cloudquery/cloudquery/issues/5409)) ([3a2c24b](https://github.com/cloudquery/cloudquery/commit/3a2c24bf34723fcee11e864d17b65f49c150fdee))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/xray to v1.15.2 ([#5410](https://github.com/cloudquery/cloudquery/issues/5410)) ([6d0be25](https://github.com/cloudquery/cloudquery/commit/6d0be25f5772a1263db27680ced6c9d11a02f293))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.11.0 ([#5416](https://github.com/cloudquery/cloudquery/issues/5416)) ([2e7ca35](https://github.com/cloudquery/cloudquery/commit/2e7ca35922fdb14fd717f582aaaa9693dae2ef4c))
* Update endpoints ([#5283](https://github.com/cloudquery/cloudquery/issues/5283)) ([5bc7bb8](https://github.com/cloudquery/cloudquery/commit/5bc7bb89bc70c34b334e3962dbf2e4d47f6af07c))
* Update endpoints ([#5384](https://github.com/cloudquery/cloudquery/issues/5384)) ([b3850cf](https://github.com/cloudquery/cloudquery/commit/b3850cf29be3b272978a75fea1e207b45fe8fb3a))

## [7.2.3](https://github.com/cloudquery/cloudquery/compare/plugins-source-aws-v7.2.2...plugins-source-aws-v7.2.3) (2022-11-30)


### Bug Fixes

* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/backup to v1.18.0 ([#5128](https://github.com/cloudquery/cloudquery/issues/5128)) ([12f64ca](https://github.com/cloudquery/cloudquery/commit/12f64caa0389209299e651e25075452e393359eb))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/cloudwatch to v1.22.0 ([#5129](https://github.com/cloudquery/cloudquery/issues/5129)) ([74d3996](https://github.com/cloudquery/cloudquery/commit/74d3996b7e4f8afee94986fdef96e183e7724230))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs to v1.17.0 ([#5130](https://github.com/cloudquery/cloudquery/issues/5130)) ([8e10157](https://github.com/cloudquery/cloudquery/commit/8e1015782a3cb194fa5c7782ce525dd87fc28324))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/configservice to v1.28.0 ([#5159](https://github.com/cloudquery/cloudquery/issues/5159)) ([dc526ae](https://github.com/cloudquery/cloudquery/commit/dc526ae3cc0a200b055403a2c20038de93185484))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/ec2 to v1.74.0 ([#5160](https://github.com/cloudquery/cloudquery/issues/5160)) ([a64d4b9](https://github.com/cloudquery/cloudquery/commit/a64d4b9ad4ed43773877b87aeb3a14d986823969))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/ecs to v1.20.0 ([#5131](https://github.com/cloudquery/cloudquery/issues/5131)) ([eb7d628](https://github.com/cloudquery/cloudquery/commit/eb7d628a49d4adc7e4321f2c2f74ef066fb7b830))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/efs to v1.18.0 ([#5132](https://github.com/cloudquery/cloudquery/issues/5132)) ([95a0a76](https://github.com/cloudquery/cloudquery/commit/95a0a76cb172a17ca6055eb3d8217905a071c144))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/eks to v1.24.0 ([#5161](https://github.com/cloudquery/cloudquery/issues/5161)) ([0a4e65e](https://github.com/cloudquery/cloudquery/commit/0a4e65ee0a2541ea11f5e0cdf8cf0f2b96445558))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/firehose to v1.14.23 ([#5127](https://github.com/cloudquery/cloudquery/issues/5127)) ([25b2ff3](https://github.com/cloudquery/cloudquery/commit/25b2ff3d89d0161e29682737645f15e021529def))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/firehose to v1.15.0 ([#5162](https://github.com/cloudquery/cloudquery/issues/5162)) ([a948774](https://github.com/cloudquery/cloudquery/commit/a9487747d6518377b158bc5acad5e7ff37d7012b))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/fsx to v1.26.0 ([#5163](https://github.com/cloudquery/cloudquery/issues/5163)) ([e6b1e3b](https://github.com/cloudquery/cloudquery/commit/e6b1e3b10b96c6a5ceb28389efc91da4d9aca590))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/glue to v1.36.0 ([#5164](https://github.com/cloudquery/cloudquery/issues/5164)) ([f80c8a7](https://github.com/cloudquery/cloudquery/commit/f80c8a764f09344d074bf093aa57d88f035c9308))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/inspector2 to v1.9.0 ([#5165](https://github.com/cloudquery/cloudquery/issues/5165)) ([7cb89b3](https://github.com/cloudquery/cloudquery/commit/7cb89b3cedf1bff6ccabcf85505c9ca8e716691b))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/iot to v1.32.0 ([#5133](https://github.com/cloudquery/cloudquery/issues/5133)) ([e8034ac](https://github.com/cloudquery/cloudquery/commit/e8034acd800a189fa15a1ed2c35c405bfb09bf5d))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/kms to v1.19.0 ([#5166](https://github.com/cloudquery/cloudquery/issues/5166)) ([53c1aab](https://github.com/cloudquery/cloudquery/commit/53c1aab13489ce47174c77d1422dacc6d4050456))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/lambda to v1.26.0 ([#5167](https://github.com/cloudquery/cloudquery/issues/5167)) ([10fb2f8](https://github.com/cloudquery/cloudquery/commit/10fb2f841a973878b37ce849e45da0a81e58ea4b))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/organizations to v1.17.0 ([#5135](https://github.com/cloudquery/cloudquery/issues/5135)) ([0fe3552](https://github.com/cloudquery/cloudquery/commit/0fe35529e7c55ddbecbfc477b4262cfe61170953))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/quicksight to v1.28.0 ([#5168](https://github.com/cloudquery/cloudquery/issues/5168)) ([cfb7caf](https://github.com/cloudquery/cloudquery/commit/cfb7caf53b0d06b2949917d56f1c2ef8094a8cae))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/rds to v1.31.0 ([#5136](https://github.com/cloudquery/cloudquery/issues/5136)) ([540fe16](https://github.com/cloudquery/cloudquery/commit/540fe16ce55a75133d7e8438c0cbce475d90cc3d))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/s3control to v1.27.0 ([#5169](https://github.com/cloudquery/cloudquery/issues/5169)) ([7f6556e](https://github.com/cloudquery/cloudquery/commit/7f6556e882ef0b1418c0ad0afa525749589f7972))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.10.0 ([#5153](https://github.com/cloudquery/cloudquery/issues/5153)) ([ea1f77e](https://github.com/cloudquery/cloudquery/commit/ea1f77e910f430287600e74cedd7d3f4ae79eb18))
* Update endpoints ([#5137](https://github.com/cloudquery/cloudquery/issues/5137)) ([a428171](https://github.com/cloudquery/cloudquery/commit/a4281710bbacd21019a4a3736b26a70f0341dd20))
* Update endpoints ([#5158](https://github.com/cloudquery/cloudquery/issues/5158)) ([96aedd1](https://github.com/cloudquery/cloudquery/commit/96aedd196d6192b4793877490e457d7c66af354a))

## [7.2.2](https://github.com/cloudquery/cloudquery/compare/plugins-source-aws-v7.2.1...plugins-source-aws-v7.2.2) (2022-11-28)


### Bug Fixes

* Add refresh window and jitter for cached credentials ([#5097](https://github.com/cloudquery/cloudquery/issues/5097)) ([919aa8c](https://github.com/cloudquery/cloudquery/commit/919aa8c0ad711cd6686d81cc6f5ff60cb1d66c36))
* **deps:** Update plugin-sdk for aws to v1.8.1 ([#5032](https://github.com/cloudquery/cloudquery/issues/5032)) ([75a1e28](https://github.com/cloudquery/cloudquery/commit/75a1e28e6da252c062b7c32bd9d1c4419b4f08ec))
* **deps:** Update plugin-sdk for aws to v1.8.2 ([#5074](https://github.com/cloudquery/cloudquery/issues/5074)) ([7112a9a](https://github.com/cloudquery/cloudquery/commit/7112a9ae9282f1e5523d5d87ed78324edd9e6875))
* **deps:** Update plugin-sdk for aws to v1.9.0 ([#5092](https://github.com/cloudquery/cloudquery/issues/5092)) ([7906991](https://github.com/cloudquery/cloudquery/commit/790699183f33dfe31374b5f4a1113c99a109b463))
* Update content ([#5005](https://github.com/cloudquery/cloudquery/issues/5005)) ([ae188d0](https://github.com/cloudquery/cloudquery/commit/ae188d0f58f2684091ee2f0915a51af9b14795fb))
* Update endpoints ([#5106](https://github.com/cloudquery/cloudquery/issues/5106)) ([33b1f08](https://github.com/cloudquery/cloudquery/commit/33b1f08ad7a2a8bd81e5906bd022a837be28711a))

## [7.2.1](https://github.com/cloudquery/cloudquery/compare/plugins-source-aws-v7.2.0...plugins-source-aws-v7.2.1) (2022-11-24)


### Bug Fixes

* Allow sync to complete when one or more accounts fail to authenticate ([#5030](https://github.com/cloudquery/cloudquery/issues/5030)) ([1e81ee6](https://github.com/cloudquery/cloudquery/commit/1e81ee6f218d9257615c844bc16166b31f5199cd))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/directconnect to v1.17.23 ([#5006](https://github.com/cloudquery/cloudquery/issues/5006)) ([74368aa](https://github.com/cloudquery/cloudquery/commit/74368aa39997621d102f1c2827ac850988afc85a))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/docdb to v1.19.16 ([#5007](https://github.com/cloudquery/cloudquery/issues/5007)) ([88664c0](https://github.com/cloudquery/cloudquery/commit/88664c036e1a07881ef7f91f398b907b2c71d68d))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/dynamodb to v1.17.7 ([#5008](https://github.com/cloudquery/cloudquery/issues/5008)) ([f286234](https://github.com/cloudquery/cloudquery/commit/f2862340522a566b03b8e76bc8e5d9a63031218b))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/ecs to v1.19.2 ([#5009](https://github.com/cloudquery/cloudquery/issues/5009)) ([ff3f633](https://github.com/cloudquery/cloudquery/commit/ff3f63301bc03962f8d08f3813d002505901f7b3))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/elasticache to v1.24.1 ([#5010](https://github.com/cloudquery/cloudquery/issues/5010)) ([3ca2f97](https://github.com/cloudquery/cloudquery/commit/3ca2f97098c0fc754d9f53551d9d0f7e63cfcb31))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/eventbridge to v1.16.20 ([#5011](https://github.com/cloudquery/cloudquery/issues/5011)) ([1d8dd9f](https://github.com/cloudquery/cloudquery/commit/1d8dd9ff82b811615fc984a079b3299674af9b71))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/kinesis to v1.15.24 ([#5012](https://github.com/cloudquery/cloudquery/issues/5012)) ([c6ffc49](https://github.com/cloudquery/cloudquery/commit/c6ffc4930d070b6ce8cb3008268a7fa835977c7b))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/kms to v1.18.18 ([#5013](https://github.com/cloudquery/cloudquery/issues/5013)) ([694a32c](https://github.com/cloudquery/cloudquery/commit/694a32c0a2eb4035dbb6ef3d03aacbce43cd8cd4))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/lambda to v1.25.1 ([#5014](https://github.com/cloudquery/cloudquery/issues/5014)) ([44a4579](https://github.com/cloudquery/cloudquery/commit/44a4579213cbc1bf9d795659e6484119edffb8c9))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/neptune to v1.18.3 ([#5015](https://github.com/cloudquery/cloudquery/issues/5015)) ([3feadc9](https://github.com/cloudquery/cloudquery/commit/3feadc9deb941a55f904cb1f8d2bf37fd0d85d25))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/rds to v1.30.1 ([#5016](https://github.com/cloudquery/cloudquery/issues/5016)) ([5e886fa](https://github.com/cloudquery/cloudquery/commit/5e886fa57db856050f3d229fbca8653300150959))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/redshift to v1.26.16 ([#5017](https://github.com/cloudquery/cloudquery/issues/5017)) ([730c996](https://github.com/cloudquery/cloudquery/commit/730c9969c685ec9c6d2b4edc8051854162675395))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/s3control to v1.26.1 ([#5020](https://github.com/cloudquery/cloudquery/issues/5020)) ([483f18f](https://github.com/cloudquery/cloudquery/commit/483f18f0fb5f42166c935d7e32a0391d9975c43f))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/secretsmanager to v1.16.8 ([#5021](https://github.com/cloudquery/cloudquery/issues/5021)) ([32a0ab0](https://github.com/cloudquery/cloudquery/commit/32a0ab05deb417229768625a6a5a260b20817439))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/sfn to v1.15.1 ([#5022](https://github.com/cloudquery/cloudquery/issues/5022)) ([4cd64ed](https://github.com/cloudquery/cloudquery/commit/4cd64ed5c05d615076f512ba3fe065a79847f46f))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/sns to v1.18.6 ([#5023](https://github.com/cloudquery/cloudquery/issues/5023)) ([36e40fc](https://github.com/cloudquery/cloudquery/commit/36e40fc270ada15b6c02d7c4b2e8664e3aaf7f73))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/sqs to v1.19.15 ([#5024](https://github.com/cloudquery/cloudquery/issues/5024)) ([114982c](https://github.com/cloudquery/cloudquery/commit/114982c03f502413bfd18945533977da09f5b481))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/ssm to v1.33.1 ([#5025](https://github.com/cloudquery/cloudquery/issues/5025)) ([fcc47e1](https://github.com/cloudquery/cloudquery/commit/fcc47e1c7e94359697519b49e85564831babba2d))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/xray to v1.15.1 ([#5026](https://github.com/cloudquery/cloudquery/issues/5026)) ([c034c20](https://github.com/cloudquery/cloudquery/commit/c034c20b1871ba7b4308cb1f33ef4a93a117e63e))

## [7.2.0](https://github.com/cloudquery/cloudquery/compare/plugins-source-aws-v7.1.4...plugins-source-aws-v7.2.0) (2022-11-23)


### Features

* Add support for AWS Step Functions ([#4832](https://github.com/cloudquery/cloudquery/issues/4832)) ([08892a7](https://github.com/cloudquery/cloudquery/commit/08892a73277890a15cd27dcac1c3136bbf8f0921))
* **aws:** Add ssoadmin permission_sets, account_assignments ([#4817](https://github.com/cloudquery/cloudquery/issues/4817)) ([4ae00ca](https://github.com/cloudquery/cloudquery/commit/4ae00ca0a6444f0844ea1a94b09adb1671e281dd))
* **aws:** Add Support for Cloudwatch Logs Resource Policy ([#4883](https://github.com/cloudquery/cloudquery/issues/4883)) ([84cb081](https://github.com/cloudquery/cloudquery/commit/84cb081b1854fffae4704fc812e9c3dfd8aa15a5))
* **aws:** Validate Service Multiplexer ([#4882](https://github.com/cloudquery/cloudquery/issues/4882)) ([9619722](https://github.com/cloudquery/cloudquery/commit/9619722491e8e7d74835e06a8119bbe27edbe370))


### Bug Fixes

* **deps:** Update module github.com/aws/aws-sdk-go-v2/config to v1.18.3 ([#4943](https://github.com/cloudquery/cloudquery/issues/4943)) ([e4aaf3f](https://github.com/cloudquery/cloudquery/commit/e4aaf3f0976a0836301b5de70a8e933c7abb5365))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/feature/s3/manager to v1.11.42 ([#4945](https://github.com/cloudquery/cloudquery/issues/4945)) ([e76c57d](https://github.com/cloudquery/cloudquery/commit/e76c57d49f44a1a370d1e2397b667be4df5c15c3))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/accessanalyzer to v1.17.3 ([#4946](https://github.com/cloudquery/cloudquery/issues/4946)) ([fbf128e](https://github.com/cloudquery/cloudquery/commit/fbf128e012ee276a0ebb6eb227502e473d96c3c0))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/acm to v1.16.3 ([#4947](https://github.com/cloudquery/cloudquery/issues/4947)) ([fde1bdf](https://github.com/cloudquery/cloudquery/commit/fde1bdfdc061d48e4c518c4a06ee3628fcec4053))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/apigateway to v1.15.25 ([#4948](https://github.com/cloudquery/cloudquery/issues/4948)) ([7ceee4a](https://github.com/cloudquery/cloudquery/commit/7ceee4a4aeb4754dd3a1c66547da01024040767e))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/apigatewayv2 to v1.12.23 ([#4949](https://github.com/cloudquery/cloudquery/issues/4949)) ([0edf753](https://github.com/cloudquery/cloudquery/commit/0edf753d75ea613a29827c8af71980dd07690b17))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/applicationautoscaling to v1.15.23 ([#4952](https://github.com/cloudquery/cloudquery/issues/4952)) ([c9fb1fd](https://github.com/cloudquery/cloudquery/commit/c9fb1fd7950ef9afc42d32ecdb785c17d6879f56))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/autoscaling to v1.24.3 ([#4953](https://github.com/cloudquery/cloudquery/issues/4953)) ([b3478bf](https://github.com/cloudquery/cloudquery/commit/b3478bfb79bc3e2b75f54a437da9e6e8e2228d28))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/cloudformation to v1.24.1 ([#4954](https://github.com/cloudquery/cloudquery/issues/4954)) ([55800f6](https://github.com/cloudquery/cloudquery/commit/55800f6c04e4d0de4f0acc3af8eb72a1b0881de1))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/cloudtrail to v1.20.3 ([#4955](https://github.com/cloudquery/cloudquery/issues/4955)) ([3c27da6](https://github.com/cloudquery/cloudquery/commit/3c27da6a8f155e399d3f1afb48af94b551724e29))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/cloudwatch to v1.21.11 ([#4956](https://github.com/cloudquery/cloudquery/issues/4956)) ([b9213b3](https://github.com/cloudquery/cloudquery/commit/b9213b3edf5d8a9f487b149f325fb42bfb1f7bc9))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs to v1.16.4 ([#4957](https://github.com/cloudquery/cloudquery/issues/4957)) ([c1a3791](https://github.com/cloudquery/cloudquery/commit/c1a379186ad89dca86c9af8ccc538c2eba7f33c3))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/configservice to v1.27.5 ([#4958](https://github.com/cloudquery/cloudquery/issues/4958)) ([2ab12d8](https://github.com/cloudquery/cloudquery/commit/2ab12d8a7740b4dd4462644b3769971665f1f872))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/databasemigrationservice to v1.22.1 ([#4959](https://github.com/cloudquery/cloudquery/issues/4959)) ([85d4d39](https://github.com/cloudquery/cloudquery/commit/85d4d394d199547f5d764f431f829edd88fc76a1))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/ec2 to v1.72.1 ([#4950](https://github.com/cloudquery/cloudquery/issues/4950)) ([b3dea9a](https://github.com/cloudquery/cloudquery/commit/b3dea9a835866ae867867cedd3e3bd1794250ab9))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/ecr to v1.17.23 ([#4951](https://github.com/cloudquery/cloudquery/issues/4951)) ([731ff3d](https://github.com/cloudquery/cloudquery/commit/731ff3d558b0e06a31ce5655a43e7c8d81d7c929))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/elasticloadbalancing to v1.14.23 ([#4960](https://github.com/cloudquery/cloudquery/issues/4960)) ([93babd5](https://github.com/cloudquery/cloudquery/commit/93babd5c79058756c2b58c3565adb4812d23640e))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2 to v1.18.26 ([#4961](https://github.com/cloudquery/cloudquery/issues/4961)) ([465c9b7](https://github.com/cloudquery/cloudquery/commit/465c9b70f2740320961f4119d35d2ca1ed1cc306))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/elasticsearchservice to v1.17.2 ([#4962](https://github.com/cloudquery/cloudquery/issues/4962)) ([aa231c6](https://github.com/cloudquery/cloudquery/commit/aa231c6b5fafad05813f9e39d5b83fe81163f339))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/emr to v1.20.16 ([#4963](https://github.com/cloudquery/cloudquery/issues/4963)) ([5964e7c](https://github.com/cloudquery/cloudquery/commit/5964e7cf7d729e225b4a4fd56fe71460a5c32aeb))
* **deps:** Update plugin-sdk for aws to v1.8.0 ([#4966](https://github.com/cloudquery/cloudquery/issues/4966)) ([16817f4](https://github.com/cloudquery/cloudquery/commit/16817f4b470d234f240832464fd4bc0ed6d30ccc))
* Update endpoints ([#4939](https://github.com/cloudquery/cloudquery/issues/4939)) ([c09f7b1](https://github.com/cloudquery/cloudquery/commit/c09f7b1c7970e481595cf1d66c5a37fedf9d2f83))

## [7.1.4](https://github.com/cloudquery/cloudquery/compare/plugins-source-aws-v7.1.3...plugins-source-aws-v7.1.4) (2022-11-22)


### Bug Fixes

* **aws:** Typo in lambda query ([#4901](https://github.com/cloudquery/cloudquery/issues/4901)) ([c05f49c](https://github.com/cloudquery/cloudquery/commit/c05f49ca1d08157983a065f1a0fe30bf158cc930))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/route53 to v1.25.0 ([#4899](https://github.com/cloudquery/cloudquery/issues/4899)) ([6cb81ca](https://github.com/cloudquery/cloudquery/commit/6cb81ca0e5bd5d53f03955d8634ebd4aa86c64e8))
* **deps:** Update plugin-sdk for aws to v1.7.0 ([#4903](https://github.com/cloudquery/cloudquery/issues/4903)) ([4ba5acd](https://github.com/cloudquery/cloudquery/commit/4ba5acdcc4118c15986b40e5a0fdfcd368a971c3))

## [7.1.3](https://github.com/cloudquery/cloudquery/compare/plugins-source-aws-v7.1.2...plugins-source-aws-v7.1.3) (2022-11-21)


### Bug Fixes

* Added password enabled column to `aws_iam_credential_reports` ([#4840](https://github.com/cloudquery/cloudquery/issues/4840)) ([a172cb9](https://github.com/cloudquery/cloudquery/commit/a172cb9648d46ab3efcbc6e6b2aae039e3f85e96))
* **aws:** Fix logger in aws_fsx_file_caches ([#4830](https://github.com/cloudquery/cloudquery/issues/4830)) ([ea6642c](https://github.com/cloudquery/cloudquery/commit/ea6642c24638287009b40259d04016a2bc71e1e1))
* **aws:** Fixed typo in JSON key ([#4868](https://github.com/cloudquery/cloudquery/issues/4868)) ([5ba36cd](https://github.com/cloudquery/cloudquery/commit/5ba36cd1af62824fef28291cb08432bc44912433)), closes [#4837](https://github.com/cloudquery/cloudquery/issues/4837)
* Region Ordering Failures ([#4870](https://github.com/cloudquery/cloudquery/issues/4870)) ([38ee9cd](https://github.com/cloudquery/cloudquery/commit/38ee9cdde554c9e873191120dc25568349c8ef0b))

## [7.1.2](https://github.com/cloudquery/cloudquery/compare/plugins-source-aws-v7.1.1...plugins-source-aws-v7.1.2) (2022-11-21)


### Bug Fixes

* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/cloudfront to v1.21.0 ([#4820](https://github.com/cloudquery/cloudquery/issues/4820)) ([9f9e4fe](https://github.com/cloudquery/cloudquery/commit/9f9e4fe58de01777a3eed3f2e929218712a20080))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/dynamodb to v1.17.6 ([#4819](https://github.com/cloudquery/cloudquery/issues/4819)) ([3a49510](https://github.com/cloudquery/cloudquery/commit/3a49510776ebfea80f7157d0b22e0e07bd7406c8))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/ec2 to v1.72.0 ([#4821](https://github.com/cloudquery/cloudquery/issues/4821)) ([28eedaf](https://github.com/cloudquery/cloudquery/commit/28eedafb01677105ae55f8ef3d572584b575e89d))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/glue to v1.35.0 ([#4822](https://github.com/cloudquery/cloudquery/issues/4822)) ([f1a6685](https://github.com/cloudquery/cloudquery/commit/f1a6685e5899c12455a6de777789047c4a557697))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/lambda to v1.25.0 ([#4823](https://github.com/cloudquery/cloudquery/issues/4823)) ([5b904fe](https://github.com/cloudquery/cloudquery/commit/5b904fec9cd424c6b0c28e1f1ae36675b564e241))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/quicksight to v1.27.0 ([#4824](https://github.com/cloudquery/cloudquery/issues/4824)) ([346e14c](https://github.com/cloudquery/cloudquery/commit/346e14c5881e4fe04421b0a7e6c78055844bc170))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/sagemaker to v1.55.0 ([#4825](https://github.com/cloudquery/cloudquery/issues/4825)) ([c167808](https://github.com/cloudquery/cloudquery/commit/c1678083625f6c8064789eeda683d0499a99620d))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/servicecatalog to v1.15.0 ([#4829](https://github.com/cloudquery/cloudquery/issues/4829)) ([5cd1607](https://github.com/cloudquery/cloudquery/commit/5cd1607a13d237334f2a8405a4a3524f8e84a9f1))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/transfer to v1.25.0 ([#4826](https://github.com/cloudquery/cloudquery/issues/4826)) ([bb95b83](https://github.com/cloudquery/cloudquery/commit/bb95b8308ef3d292102b05bb9af015f6c0ec4c78))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/workspaces to v1.27.0 ([#4827](https://github.com/cloudquery/cloudquery/issues/4827)) ([e330138](https://github.com/cloudquery/cloudquery/commit/e330138e6f7ae9e24d84434c94c840d07cc274ce))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/xray to v1.15.0 ([#4828](https://github.com/cloudquery/cloudquery/issues/4828)) ([370d4e8](https://github.com/cloudquery/cloudquery/commit/370d4e859256d3cb0bf42f7d38b677ec694b20ef))
* **deps:** Update plugin-sdk for aws to v1.6.0 ([#4841](https://github.com/cloudquery/cloudquery/issues/4841)) ([09267c6](https://github.com/cloudquery/cloudquery/commit/09267c6462872363ced5fb1f3e50beea5cf4b536))
* Update endpoints ([#4816](https://github.com/cloudquery/cloudquery/issues/4816)) ([9845388](https://github.com/cloudquery/cloudquery/commit/9845388627d8b0f1636a991601a7088d6c6e2b57))

## [7.1.1](https://github.com/cloudquery/cloudquery/compare/plugins-source-aws-v7.1.0...plugins-source-aws-v7.1.1) (2022-11-18)


### Bug Fixes

* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/rds to v1.30.0 ([#4782](https://github.com/cloudquery/cloudquery/issues/4782)) ([3be128c](https://github.com/cloudquery/cloudquery/commit/3be128c5cea7ce6bee5da8c9565d5b85d3c3a908))
* Fixed cloudfront policy queries with null json fields ([#4773](https://github.com/cloudquery/cloudquery/issues/4773)) ([43095c4](https://github.com/cloudquery/cloudquery/commit/43095c45a0dd343f262018e864a542879d02021b))

## [7.1.0](https://github.com/cloudquery/cloudquery/compare/plugins-source-aws-v7.0.1...plugins-source-aws-v7.1.0) (2022-11-18)


### Features

* **aws:** Add Support For Config Rules and Compliance ([#4730](https://github.com/cloudquery/cloudquery/issues/4730)) ([100f4ba](https://github.com/cloudquery/cloudquery/commit/100f4ba6775eb90c6a228790b9392c8a8130f0f8))
* **aws:** Support AWS Account Contacts ([#4734](https://github.com/cloudquery/cloudquery/issues/4734)) ([bdf3867](https://github.com/cloudquery/cloudquery/commit/bdf3867ada5379d2706b64450870ad32a0833952))


### Bug Fixes

* **aws:** Cloudfront control fails for specific configuration ([#4735](https://github.com/cloudquery/cloudquery/issues/4735)) ([7507bae](https://github.com/cloudquery/cloudquery/commit/7507bae1e7572055f32565d96b79c78ad66ed8e3))
* **aws:** Remove ARN helper usage ([#4714](https://github.com/cloudquery/cloudquery/issues/4714)) ([dde430f](https://github.com/cloudquery/cloudquery/commit/dde430f33ed9adcfb7b27d0c2468250310f2fe7e)), closes [#4689](https://github.com/cloudquery/cloudquery/issues/4689)
* **deps:** Update module github.com/aws/aws-sdk-go-v2/config to v1.18.1 ([#4736](https://github.com/cloudquery/cloudquery/issues/4736)) ([db70d2a](https://github.com/cloudquery/cloudquery/commit/db70d2a602fc4edfc74ed61fd7d28ada6da6a3af))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/config to v1.18.2 ([#4775](https://github.com/cloudquery/cloudquery/issues/4775)) ([136fb42](https://github.com/cloudquery/cloudquery/commit/136fb4213da150f8f9e4a68019fbe7fe94397370))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/credentials to v1.13.2 ([#4737](https://github.com/cloudquery/cloudquery/issues/4737)) ([2b2433d](https://github.com/cloudquery/cloudquery/commit/2b2433dd3e091147d0a8beed86159fa7b670e783))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/feature/s3/manager to v1.11.40 ([#4738](https://github.com/cloudquery/cloudquery/issues/4738)) ([69b01cc](https://github.com/cloudquery/cloudquery/commit/69b01cc65d9a62a4b734c956315a95aa1daafa86))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/feature/s3/manager to v1.11.41 ([#4776](https://github.com/cloudquery/cloudquery/issues/4776)) ([d0b2bae](https://github.com/cloudquery/cloudquery/commit/d0b2baeb33e047accc5ef0ac3b68f1605866fe4f))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/accessanalyzer to v1.17.2 ([#4739](https://github.com/cloudquery/cloudquery/issues/4739)) ([3c85f78](https://github.com/cloudquery/cloudquery/commit/3c85f78c4b9322d93173d5029014c904f50fd5b8))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/acm to v1.16.2 ([#4740](https://github.com/cloudquery/cloudquery/issues/4740)) ([e2a734d](https://github.com/cloudquery/cloudquery/commit/e2a734dda89db5137a7738d864c484f5f0d5cc12))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/apigateway to v1.15.24 ([#4741](https://github.com/cloudquery/cloudquery/issues/4741)) ([6bec735](https://github.com/cloudquery/cloudquery/commit/6bec7352510f791026f17cb9640b54a51c22f6eb))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/apigatewayv2 to v1.12.22 ([#4742](https://github.com/cloudquery/cloudquery/issues/4742)) ([56bebb8](https://github.com/cloudquery/cloudquery/commit/56bebb853e41e87abc841b3e04cf087fbecf72ca))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/applicationautoscaling to v1.15.22 ([#4756](https://github.com/cloudquery/cloudquery/issues/4756)) ([ba77807](https://github.com/cloudquery/cloudquery/commit/ba778079bc4e9f660c9e978d7ef8c78a57bb0d78))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/appsync to v1.16.0 ([#4788](https://github.com/cloudquery/cloudquery/issues/4788)) ([21d3b1d](https://github.com/cloudquery/cloudquery/commit/21d3b1d56f9b08a31804a0effa8114678481b473))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/autoscaling to v1.24.2 ([#4743](https://github.com/cloudquery/cloudquery/issues/4743)) ([c366ea3](https://github.com/cloudquery/cloudquery/commit/c366ea390c500182479e3d6e593a36455240d876))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/cloudformation to v1.24.0 ([#4789](https://github.com/cloudquery/cloudquery/issues/4789)) ([3616ed2](https://github.com/cloudquery/cloudquery/commit/3616ed2fe372ac3fcce682c04a9c5672b0fd707a))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/cloudtrail to v1.20.2 ([#4744](https://github.com/cloudquery/cloudquery/issues/4744)) ([c0c32d3](https://github.com/cloudquery/cloudquery/commit/c0c32d31424fdeda83d52ba0ad936650704bba17))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/cloudwatch to v1.21.10 ([#4745](https://github.com/cloudquery/cloudquery/issues/4745)) ([f1c7a99](https://github.com/cloudquery/cloudquery/commit/f1c7a99669bf22e1cf1377e82aef854b2117256c))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs to v1.16.3 ([#4746](https://github.com/cloudquery/cloudquery/issues/4746)) ([1d020a0](https://github.com/cloudquery/cloudquery/commit/1d020a054fff421217b8a99b07169f826c19e810))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/configservice to v1.27.4 ([#4747](https://github.com/cloudquery/cloudquery/issues/4747)) ([ec6aed0](https://github.com/cloudquery/cloudquery/commit/ec6aed0b01ef2af72d0ff5dac0d04957a09e9a49))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/databasemigrationservice to v1.21.16 ([#4748](https://github.com/cloudquery/cloudquery/issues/4748)) ([a04a343](https://github.com/cloudquery/cloudquery/commit/a04a3436c243c7c5a2f5a88b6f287a0024c3d480))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/databasemigrationservice to v1.22.0 ([#4790](https://github.com/cloudquery/cloudquery/issues/4790)) ([cffe44a](https://github.com/cloudquery/cloudquery/commit/cffe44a2fdbcac604347d211d52ab56a2c2e0f31))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/directconnect to v1.17.22 ([#4749](https://github.com/cloudquery/cloudquery/issues/4749)) ([d98c79e](https://github.com/cloudquery/cloudquery/commit/d98c79e8e5b85bfebfb3a5a7bb466c9e3e409b15))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/docdb to v1.19.15 ([#4757](https://github.com/cloudquery/cloudquery/issues/4757)) ([d058d34](https://github.com/cloudquery/cloudquery/commit/d058d347e48e7438a92f4ef69d8053d593758846))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/dynamodb to v1.17.5 ([#4750](https://github.com/cloudquery/cloudquery/issues/4750)) ([9b0e200](https://github.com/cloudquery/cloudquery/commit/9b0e200b654f7bb839db845d31edee94447ec980))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/ec2 to v1.70.1 ([#4751](https://github.com/cloudquery/cloudquery/issues/4751)) ([6b75a34](https://github.com/cloudquery/cloudquery/commit/6b75a3444cc0ed9e8e032f1cd509c9e5e594eedb))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/ec2 to v1.71.0 ([#4791](https://github.com/cloudquery/cloudquery/issues/4791)) ([58368df](https://github.com/cloudquery/cloudquery/commit/58368df813a911976848f67346699bd162ba7346))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/ecr to v1.17.22 ([#4752](https://github.com/cloudquery/cloudquery/issues/4752)) ([b769ff7](https://github.com/cloudquery/cloudquery/commit/b769ff78e3eda16ee3f0cf3bf5e8ba0a140ff675))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/ecs to v1.19.1 ([#4753](https://github.com/cloudquery/cloudquery/issues/4753)) ([95b3bc4](https://github.com/cloudquery/cloudquery/commit/95b3bc42b5c47e938af4b3843b460884fe62c42d))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/eks to v1.23.0 ([#4792](https://github.com/cloudquery/cloudquery/issues/4792)) ([a7e546b](https://github.com/cloudquery/cloudquery/commit/a7e546b3079a9dd34d20ac303686274dbdfa3565))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/elasticache to v1.24.0 ([#4793](https://github.com/cloudquery/cloudquery/issues/4793)) ([c392bfb](https://github.com/cloudquery/cloudquery/commit/c392bfbdfc5b69d4399ad339d03b0471c790d81e))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/elasticloadbalancing to v1.14.22 ([#4758](https://github.com/cloudquery/cloudquery/issues/4758)) ([1a385c3](https://github.com/cloudquery/cloudquery/commit/1a385c3f02bdb2549cb16ad99786541a71296aa8))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2 to v1.18.24 ([#4759](https://github.com/cloudquery/cloudquery/issues/4759)) ([236ac56](https://github.com/cloudquery/cloudquery/commit/236ac56bb1bac5d4b8bad50e2acc1b0dfdda8f68))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2 to v1.18.25 ([#4777](https://github.com/cloudquery/cloudquery/issues/4777)) ([62dda2b](https://github.com/cloudquery/cloudquery/commit/62dda2b5d6ae77fefba0f8e6a2df5e5081e50d24))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/elasticsearchservice to v1.17.1 ([#4760](https://github.com/cloudquery/cloudquery/issues/4760)) ([02b41f6](https://github.com/cloudquery/cloudquery/commit/02b41f6d242217ad544f03e8081d2e8da1f38bcc))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/emr to v1.20.15 ([#4761](https://github.com/cloudquery/cloudquery/issues/4761)) ([0d4e1d5](https://github.com/cloudquery/cloudquery/commit/0d4e1d5a7c6879e020d1cd7c6617e2d9b6e1b7f5))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/eventbridge to v1.16.19 ([#4762](https://github.com/cloudquery/cloudquery/issues/4762)) ([79fedc4](https://github.com/cloudquery/cloudquery/commit/79fedc4ed5b38a0f4ee1263655d523312df46e6d))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/kinesis to v1.15.23 ([#4763](https://github.com/cloudquery/cloudquery/issues/4763)) ([ed16606](https://github.com/cloudquery/cloudquery/commit/ed166065f526b73809baac7d7d3c4ce2302e72ec))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/kms to v1.18.17 ([#4764](https://github.com/cloudquery/cloudquery/issues/4764)) ([a43ec2a](https://github.com/cloudquery/cloudquery/commit/a43ec2ace88a57870246a4eec389d15ed49415c0))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/lambda to v1.24.11 ([#4765](https://github.com/cloudquery/cloudquery/issues/4765)) ([65fa361](https://github.com/cloudquery/cloudquery/commit/65fa36111d1d16660ca5d8eb8392d86ce24fcea1))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/neptune to v1.18.2 ([#4766](https://github.com/cloudquery/cloudquery/issues/4766)) ([01d762b](https://github.com/cloudquery/cloudquery/commit/01d762bcc5bc090560f95533b5e070d0215e89e9))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/redshift to v1.26.15 ([#4778](https://github.com/cloudquery/cloudquery/issues/4778)) ([2c7a42e](https://github.com/cloudquery/cloudquery/commit/2c7a42e57035d58eb6855d8ba1a334aecdfdf878))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/s3control to v1.26.0 ([#4783](https://github.com/cloudquery/cloudquery/issues/4783)) ([430164e](https://github.com/cloudquery/cloudquery/commit/430164e75d77f80c069a041946ccf02b82248960))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/secretsmanager to v1.16.7 ([#4780](https://github.com/cloudquery/cloudquery/issues/4780)) ([15cda18](https://github.com/cloudquery/cloudquery/commit/15cda1819fd9f927af98d81539db2fb124198b8b))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/servicecatalogappregistry to v1.15.0 ([#4784](https://github.com/cloudquery/cloudquery/issues/4784)) ([114ff80](https://github.com/cloudquery/cloudquery/commit/114ff80c0a40cce8d02d93a453969a769f590ccc))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/sns to v1.18.5 ([#4781](https://github.com/cloudquery/cloudquery/issues/4781)) ([72f220d](https://github.com/cloudquery/cloudquery/commit/72f220ddd9da2cfc2737649b63dbf15244a06341))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/sqs to v1.19.14 ([#4786](https://github.com/cloudquery/cloudquery/issues/4786)) ([8e76293](https://github.com/cloudquery/cloudquery/commit/8e76293b7ac2b9f0abb8c9aabff249c57b5b2f9d))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/ssm to v1.33.0 ([#4785](https://github.com/cloudquery/cloudquery/issues/4785)) ([fcd03ff](https://github.com/cloudquery/cloudquery/commit/fcd03ff104389c327dc8571d6d6824951b13f589))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/timestreamwrite to v1.14.6 ([#4787](https://github.com/cloudquery/cloudquery/issues/4787)) ([4cc2f6b](https://github.com/cloudquery/cloudquery/commit/4cc2f6b12eb34f5b291b4001b6abac695409d985))
* Fix Name tag reference in grafana dashboard ([#4801](https://github.com/cloudquery/cloudquery/issues/4801)) ([bd38047](https://github.com/cloudquery/cloudquery/commit/bd38047e797e46191eae5852c83688b1db007575))
* Update endpoints ([#4732](https://github.com/cloudquery/cloudquery/issues/4732)) ([f5bb4d0](https://github.com/cloudquery/cloudquery/commit/f5bb4d0906b95ad78c4d2d6b9008e3fc471f8ad7))
* Update endpoints ([#4772](https://github.com/cloudquery/cloudquery/issues/4772)) ([33dc095](https://github.com/cloudquery/cloudquery/commit/33dc095563dbac2f66548de6be5513735973fd1e))

## [7.0.1](https://github.com/cloudquery/cloudquery/compare/plugins-source-aws-v7.0.0...plugins-source-aws-v7.0.1) (2022-11-16)


### Bug Fixes

* Improve AWS credentials error message and local profile docs ([#4708](https://github.com/cloudquery/cloudquery/issues/4708)) ([2a159b8](https://github.com/cloudquery/cloudquery/commit/2a159b83528a7f1127c7292add2ea299a56bc0c7))

## [7.0.0](https://github.com/cloudquery/cloudquery/compare/plugins-source-aws-v6.2.0...plugins-source-aws-v7.0.0) (2022-11-16)


### ⚠ BREAKING CHANGES

* **aws:** Removed `aws_resource_share_associated_principals` and `aws_resource_share_associated_resources` relations of `aws_ram_resource_shares` and created a new top level table to hold this data `aws_resource_share_associations`.  To sync these resources, ensure to either use `tables: ["*"]` or `tables: ["aws_resource_share_associations"]`
* **aws:** Moved `aws_ram_resource_share_permissions` to be a relation of `aws_ram_resource_shares` instead of a top level table. To sync `aws_ram_resource_share_permissions` you'll need to ensure to sync `aws_ram_resource_shares` too by specifying either `tables: ["*"]` or `tables: ["aws_ram_resource_shares"]`

### Bug Fixes

* **aws:** Fix ram resources ([#4636](https://github.com/cloudquery/cloudquery/issues/4636)) ([2609f3e](https://github.com/cloudquery/cloudquery/commit/2609f3e868369f3a2af8fa6c7e5bdc00a036974a))
* **aws:** Moved `aws_ram_resource_share_permissions` to be a relation of `aws_ram_resource_shares` instead of a top level table. To sync `aws_ram_resource_share_permissions` you'll need to ensure to sync `aws_ram_resource_shares` too by specifying either `tables: ["*"]` or `tables: ["aws_ram_resource_shares"]` ([2609f3e](https://github.com/cloudquery/cloudquery/commit/2609f3e868369f3a2af8fa6c7e5bdc00a036974a))
* **aws:** Removed `aws_resource_share_associated_principals` and `aws_resource_share_associated_resources` relations of `aws_ram_resource_shares` and created a new top level table to hold this data `aws_resource_share_associations`.  To sync these resources, ensure to either use `tables: ["*"]` or `tables: ["aws_resource_share_associations"]` ([2609f3e](https://github.com/cloudquery/cloudquery/commit/2609f3e868369f3a2af8fa6c7e5bdc00a036974a))

## [6.2.0](https://github.com/cloudquery/cloudquery/compare/plugins-source-aws-v6.1.0...plugins-source-aws-v6.2.0) (2022-11-16)


### Features

* **aws:** Add Support for Eventbridge scheduler ([#4693](https://github.com/cloudquery/cloudquery/issues/4693)) ([a4ef661](https://github.com/cloudquery/cloudquery/commit/a4ef661c23e95e44df57a380cd8dda8bdd7900ee))
* **aws:** More SSM resources ([#4381](https://github.com/cloudquery/cloudquery/issues/4381)) ([42b0c70](https://github.com/cloudquery/cloudquery/commit/42b0c70c2246977596b32b164e0958d62688acd9))

## [6.1.0](https://github.com/cloudquery/cloudquery/compare/plugins-source-aws-v6.0.1...plugins-source-aws-v6.1.0) (2022-11-16)


### Features

* **aws:** Add Support for ECS Scale In Task protection ([#4688](https://github.com/cloudquery/cloudquery/issues/4688)) ([44e8830](https://github.com/cloudquery/cloudquery/commit/44e88305ef0430e873e3bcacc6f6349ad7c336a7))


### Bug Fixes

* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/route53 to v1.24.0 ([#4694](https://github.com/cloudquery/cloudquery/issues/4694)) ([a6d1fec](https://github.com/cloudquery/cloudquery/commit/a6d1fecba8e56d38c1b92bd6599298d2d0f1efc3))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/workspaces to v1.26.0 ([#4695](https://github.com/cloudquery/cloudquery/issues/4695)) ([ea07f09](https://github.com/cloudquery/cloudquery/commit/ea07f0923b1f5541c543f17e1b3074838420a682))

## [6.0.1](https://github.com/cloudquery/cloudquery/compare/plugins-source-aws-v6.0.0...plugins-source-aws-v6.0.1) (2022-11-15)


### Bug Fixes

* **deps:** Update plugin-sdk for aws to v1.5.3 ([#4640](https://github.com/cloudquery/cloudquery/issues/4640)) ([29f6adc](https://github.com/cloudquery/cloudquery/commit/29f6adc387dc7003803b8830c3f683ae916360d4))
* Update endpoints ([#4686](https://github.com/cloudquery/cloudquery/issues/4686)) ([a33c13e](https://github.com/cloudquery/cloudquery/commit/a33c13e12cb6e3473c4ee16dd01e7454c0ba8705))

## [6.0.0](https://github.com/cloudquery/cloudquery/compare/plugins-source-aws-v5.2.0...plugins-source-aws-v6.0.0) (2022-11-15)


### ⚠ BREAKING CHANGES

* **aws:** Renamed `aws_apprunner_auto_scaling_configuration` to `aws_apprunner_auto_scaling_configurations`
* **aws:** Renamed `aws_apprunner_vpc_connector` to `aws_apprunner_vpc_connectors`
* **aws:** Renamed `aws_apprunner_vpc_ingress_connection` to `aws_apprunner_vpc_ingress_connections`
* **aws:** Renamed `aws_ec2_regional_config` to `aws_ec2_regional_configs`
* **aws:** Renamed `aws_lightsail_disk_snapshot` to `aws_lightsail_disk_snapshots`
* **aws:** Renamed `aws_xray_encryption_config` to `aws_xray_encryption_configs`

### Bug Fixes

* **aws:** Ensure all AWS table names are plural ([944f1b1](https://github.com/cloudquery/cloudquery/commit/944f1b12991965c2de5dc58c6c504ae7775e368d))

## [5.2.0](https://github.com/cloudquery/cloudquery/compare/plugins-source-aws-v5.1.2...plugins-source-aws-v5.2.0) (2022-11-15)


### Features

* Add MSK and codegen from AWS Client structs ([#3967](https://github.com/cloudquery/cloudquery/issues/3967)) ([eb602bb](https://github.com/cloudquery/cloudquery/commit/eb602bb45cb325f75b4e1fc22727b25ed7db5f33))

## [5.1.2](https://github.com/cloudquery/cloudquery/compare/plugins-source-aws-v5.1.1...plugins-source-aws-v5.1.2) (2022-11-15)


### Bug Fixes

* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/rds to v1.29.0 ([#4612](https://github.com/cloudquery/cloudquery/issues/4612)) ([4bcf227](https://github.com/cloudquery/cloudquery/commit/4bcf227dc16956164a22db53caee5fac62c05e50))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/xray to v1.14.0 ([#4613](https://github.com/cloudquery/cloudquery/issues/4613)) ([e2c7e64](https://github.com/cloudquery/cloudquery/commit/e2c7e64c93a52867eab903a964d71110b42498e3))

## [5.1.1](https://github.com/cloudquery/cloudquery/compare/plugins-source-aws-v5.1.0...plugins-source-aws-v5.1.1) (2022-11-14)


### Bug Fixes

* **aws:** RDS Version Errors ([#4431](https://github.com/cloudquery/cloudquery/issues/4431)) ([c9aaf79](https://github.com/cloudquery/cloudquery/commit/c9aaf7976f20d1275f0859fdb61258bac5676536))

## [5.1.0](https://github.com/cloudquery/cloudquery/compare/plugins-source-aws-v5.0.0...plugins-source-aws-v5.1.0) (2022-11-14)


### Features

* **aws:** Implement Timestream resources ([#4553](https://github.com/cloudquery/cloudquery/issues/4553)) ([5a2ac0f](https://github.com/cloudquery/cloudquery/commit/5a2ac0f1fe2b36bf1356722c77ff7d79b445842d))


### Bug Fixes

* **deps:** Update plugin-sdk for aws to v1.5.2 ([#4555](https://github.com/cloudquery/cloudquery/issues/4555)) ([c5c1fc1](https://github.com/cloudquery/cloudquery/commit/c5c1fc1df7d4f4f4ab92e8a47b140075103affc8))

## [5.0.0](https://github.com/cloudquery/cloudquery/compare/plugins-source-aws-v4.17.1...plugins-source-aws-v5.0.0) (2022-11-14)


### ⚠ BREAKING CHANGES

* change `aws_rds_engine_versions`  to be regional. If you've previously synced this table you'll need to drop it for the PK change (migration) to succeed.

### Bug Fixes

* change `aws_rds_engine_versions`  to be regional. If you've previously synced this table you'll need to drop it for the PK change (migration) to succeed. ([71926dd](https://github.com/cloudquery/cloudquery/commit/71926dd2923be6dbea0550103c05a3011fa7ba9f))

## [4.17.1](https://github.com/cloudquery/cloudquery/compare/plugins-source-aws-v4.17.0...plugins-source-aws-v4.17.1) (2022-11-14)


### Bug Fixes

* **deps:** Update plugin-sdk for aws to v1.5.1 ([#4496](https://github.com/cloudquery/cloudquery/issues/4496)) ([62e1c11](https://github.com/cloudquery/cloudquery/commit/62e1c11709522796a44d7b52c883cd67c163b5b3))

## [4.17.0](https://github.com/cloudquery/cloudquery/compare/plugins-source-aws-v4.16.0...plugins-source-aws-v4.17.0) (2022-11-13)


### Features

* **aws:** Add support for FSX File Caches ([#4091](https://github.com/cloudquery/cloudquery/issues/4091)) ([5b6597a](https://github.com/cloudquery/cloudquery/commit/5b6597afe6cde7a4c2f8d50c9008e71c131284e2)), closes [#2605](https://github.com/cloudquery/cloudquery/issues/2605)


### Bug Fixes

* **aws:** Remove NotFound checks in PreResourceResolvers ([#4376](https://github.com/cloudquery/cloudquery/issues/4376)) ([01a173b](https://github.com/cloudquery/cloudquery/commit/01a173b786eeac8c8addb1e2642f262eff11f397)), closes [#4374](https://github.com/cloudquery/cloudquery/issues/4374)
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/ecs to v1.19.0 ([#4452](https://github.com/cloudquery/cloudquery/issues/4452)) ([21aff52](https://github.com/cloudquery/cloudquery/commit/21aff52cf7157c2e346d7b454cd596234cb181ab))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/elasticsearchservice to v1.17.0 ([#4453](https://github.com/cloudquery/cloudquery/issues/4453)) ([0b6a342](https://github.com/cloudquery/cloudquery/commit/0b6a342da93fb5084b1aab73462fe9c1cb084348))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/iot to v1.31.0 ([#4454](https://github.com/cloudquery/cloudquery/issues/4454)) ([85d3fc9](https://github.com/cloudquery/cloudquery/commit/85d3fc9b1480d274abc51329e775d975b455d73a))

## [4.16.0](https://github.com/cloudquery/cloudquery/compare/plugins-source-aws-v4.15.1...plugins-source-aws-v4.16.0) (2022-11-13)


### Features

* **aws:** Debug logging to include body ([#3964](https://github.com/cloudquery/cloudquery/issues/3964)) ([089ba2f](https://github.com/cloudquery/cloudquery/commit/089ba2fc89df2d41930aa5d96dda85a96b4897dd))
* **aws:** Quicksight resources ([#4116](https://github.com/cloudquery/cloudquery/issues/4116)) ([6190de2](https://github.com/cloudquery/cloudquery/commit/6190de2f272c8fff9cf22adcfaa2ccb4acab2e4a))


### Bug Fixes

* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/emr to v1.20.14 ([#4367](https://github.com/cloudquery/cloudquery/issues/4367)) ([a36d3e2](https://github.com/cloudquery/cloudquery/commit/a36d3e2927387a100c8881eb6cfb343112007372))

## [4.15.1](https://github.com/cloudquery/cloudquery/compare/plugins-source-aws-v4.15.0...plugins-source-aws-v4.15.1) (2022-11-13)


### Bug Fixes

* **deps:** Update module github.com/aws/aws-sdk-go-v2/feature/s3/manager to v1.11.39 ([#4433](https://github.com/cloudquery/cloudquery/issues/4433)) ([48a8d73](https://github.com/cloudquery/cloudquery/commit/48a8d732d748c2626ae8f568590209454d45b0b7))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/acm to v1.16.1 ([#4450](https://github.com/cloudquery/cloudquery/issues/4450)) ([f1f22be](https://github.com/cloudquery/cloudquery/commit/f1f22be23782c153440a404882870f22ca0906df))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/ec2 to v1.70.0 ([#4451](https://github.com/cloudquery/cloudquery/issues/4451)) ([ffc1ced](https://github.com/cloudquery/cloudquery/commit/ffc1ced3b3a31702d687a56dd37d4d9ef02c0a53))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/glue to v1.34.1 ([#4434](https://github.com/cloudquery/cloudquery/issues/4434)) ([a71ab65](https://github.com/cloudquery/cloudquery/commit/a71ab65ee5b5f2413bbf9da76591da8e3f25f7e4))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/kms to v1.18.16 ([#4435](https://github.com/cloudquery/cloudquery/issues/4435)) ([e1976c4](https://github.com/cloudquery/cloudquery/commit/e1976c40cb061a126e7a0e84ae4884d8510ea038))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/lambda to v1.24.10 ([#4436](https://github.com/cloudquery/cloudquery/issues/4436)) ([bab76a5](https://github.com/cloudquery/cloudquery/commit/bab76a552a5bdc87fcbbace1e1180102c5bb899f))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/lightsail to v1.24.0 ([#4455](https://github.com/cloudquery/cloudquery/issues/4455)) ([5afbde7](https://github.com/cloudquery/cloudquery/commit/5afbde7bf41af535494fc576b83c1ff9342c94f8))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/neptune to v1.18.1 ([#4437](https://github.com/cloudquery/cloudquery/issues/4437)) ([53589da](https://github.com/cloudquery/cloudquery/commit/53589da9759df49354821ed96ce2f1bcd96c7de1))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/rds to v1.28.1 ([#4438](https://github.com/cloudquery/cloudquery/issues/4438)) ([474ebdc](https://github.com/cloudquery/cloudquery/commit/474ebdc3492af185aa14e21dfc75242e76212f74))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/redshift to v1.26.14 ([#4439](https://github.com/cloudquery/cloudquery/issues/4439)) ([11ed59b](https://github.com/cloudquery/cloudquery/commit/11ed59bd4d34c10ba77aa90881dc01a53fa828f6))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/resourcegroups to v1.12.21 ([#4440](https://github.com/cloudquery/cloudquery/issues/4440)) ([a8023ab](https://github.com/cloudquery/cloudquery/commit/a8023ab9639acf7543878a629ee90a57d2542714))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/route53 to v1.23.0 ([#4457](https://github.com/cloudquery/cloudquery/issues/4457)) ([2be62d6](https://github.com/cloudquery/cloudquery/commit/2be62d66abea037c89b402e2793e2e33e8f418b1))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/s3control to v1.25.1 ([#4441](https://github.com/cloudquery/cloudquery/issues/4441)) ([4a5f30a](https://github.com/cloudquery/cloudquery/commit/4a5f30a0f2ac68177cea8beddd76b7b606712b59))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/secretsmanager to v1.16.5 ([#4442](https://github.com/cloudquery/cloudquery/issues/4442)) ([0a43de8](https://github.com/cloudquery/cloudquery/commit/0a43de80693541878e02ac047385996edca44454))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/sesv2 to v1.15.1 ([#4443](https://github.com/cloudquery/cloudquery/issues/4443)) ([6a19df3](https://github.com/cloudquery/cloudquery/commit/6a19df32743ca7c46db8972722ad47c7e6bcd52c))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/sns to v1.18.4 ([#4448](https://github.com/cloudquery/cloudquery/issues/4448)) ([7049020](https://github.com/cloudquery/cloudquery/commit/704902067f930c337c06717b60c4520e70b3f240))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/sqs to v1.19.13 ([#4449](https://github.com/cloudquery/cloudquery/issues/4449)) ([10188f1](https://github.com/cloudquery/cloudquery/commit/10188f1ee4475550b3fe157b37ed8f79cc56387f))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/ssm to v1.32.1 ([#4444](https://github.com/cloudquery/cloudquery/issues/4444)) ([d01a68d](https://github.com/cloudquery/cloudquery/commit/d01a68d07f86a5ad258d749d59a0b2670beeaab1))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/transfer to v1.23.3 ([#4445](https://github.com/cloudquery/cloudquery/issues/4445)) ([6345e79](https://github.com/cloudquery/cloudquery/commit/6345e79c9ae373133fdaf2ec5d8ba9005088b2ee))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/xray to v1.13.22 ([#4446](https://github.com/cloudquery/cloudquery/issues/4446)) ([4f6af8a](https://github.com/cloudquery/cloudquery/commit/4f6af8a2ae85234dfbc034e62f4921f4ac7c68ec))
* Update endpoints ([#4429](https://github.com/cloudquery/cloudquery/issues/4429)) ([630688d](https://github.com/cloudquery/cloudquery/commit/630688df16065408935475e8b2b741dff5d8fa94))

## [4.15.0](https://github.com/cloudquery/cloudquery/compare/plugins-source-aws-v4.14.0...plugins-source-aws-v4.15.0) (2022-11-11)


### Features

* Increase MaxResults for appstream resources ([#4377](https://github.com/cloudquery/cloudquery/issues/4377)) ([17f5dcd](https://github.com/cloudquery/cloudquery/commit/17f5dcd81473b96c354492cdf339bdf5ea892ac1))


### Bug Fixes

* **deps:** Update module github.com/aws/aws-sdk-go-v2/credentials to v1.12.24 ([#4344](https://github.com/cloudquery/cloudquery/issues/4344)) ([90d3661](https://github.com/cloudquery/cloudquery/commit/90d3661b1adc0421b4abb99aabe3aa25a15715e1))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/feature/s3/manager to v1.11.38 ([#4345](https://github.com/cloudquery/cloudquery/issues/4345)) ([3191ffb](https://github.com/cloudquery/cloudquery/commit/3191ffb95e0e01ac1a2d14f4ceb44ff07e756e86))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/accessanalyzer to v1.17.1 ([#4346](https://github.com/cloudquery/cloudquery/issues/4346)) ([3ee406b](https://github.com/cloudquery/cloudquery/commit/3ee406b9390cc5a6cc1596db7cbe7a6f8c390c3e))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/apigateway to v1.15.23 ([#4347](https://github.com/cloudquery/cloudquery/issues/4347)) ([06ea319](https://github.com/cloudquery/cloudquery/commit/06ea319f51844aeaf63fea5c150734eb5358efc6))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/apigatewayv2 to v1.12.21 ([#4348](https://github.com/cloudquery/cloudquery/issues/4348)) ([3c94cd2](https://github.com/cloudquery/cloudquery/commit/3c94cd263c1505501d92d53807cbe118dbc169db))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/applicationautoscaling to v1.15.21 ([#4349](https://github.com/cloudquery/cloudquery/issues/4349)) ([4e6e897](https://github.com/cloudquery/cloudquery/commit/4e6e89733d9a6e9fbdb36231b4ee86e8983a7a16))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/athena to v1.19.1 ([#4350](https://github.com/cloudquery/cloudquery/issues/4350)) ([e1b4df6](https://github.com/cloudquery/cloudquery/commit/e1b4df61afb4760ad10320eaf07b394aa205747d))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/autoscaling to v1.24.1 ([#4351](https://github.com/cloudquery/cloudquery/issues/4351)) ([0ff4842](https://github.com/cloudquery/cloudquery/commit/0ff4842d2206b32a356a6d5346b84b6f098ce931))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/cloudformation to v1.23.1 ([#4352](https://github.com/cloudquery/cloudquery/issues/4352)) ([b9fb504](https://github.com/cloudquery/cloudquery/commit/b9fb5041e37bbfa8cdf1525d48147590664cc9ad))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/cloudtrail to v1.20.1 ([#4353](https://github.com/cloudquery/cloudquery/issues/4353)) ([d7ad670](https://github.com/cloudquery/cloudquery/commit/d7ad670b118884883bec2fb941cb2db9a81cc5b6))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/cloudwatch to v1.21.9 ([#4354](https://github.com/cloudquery/cloudquery/issues/4354)) ([0285a02](https://github.com/cloudquery/cloudquery/commit/0285a0239cc4d084cffee858e0fd3c10c41bdb24))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs to v1.16.2 ([#4355](https://github.com/cloudquery/cloudquery/issues/4355)) ([c6a67b2](https://github.com/cloudquery/cloudquery/commit/c6a67b27e2efca06f790859e14a3be963ec2fc36))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/configservice to v1.27.3 ([#4356](https://github.com/cloudquery/cloudquery/issues/4356)) ([51b219c](https://github.com/cloudquery/cloudquery/commit/51b219c49d682762c9c516e8f6c51a368b131ff4))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/databasemigrationservice to v1.21.15 ([#4357](https://github.com/cloudquery/cloudquery/issues/4357)) ([528845c](https://github.com/cloudquery/cloudquery/commit/528845c2b8f726d68ba169d8ba2bcf91feddfe42))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/directconnect to v1.17.21 ([#4358](https://github.com/cloudquery/cloudquery/issues/4358)) ([0aa1cd1](https://github.com/cloudquery/cloudquery/commit/0aa1cd11a57f107b667cfadf3e1e17b3a7619248))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/docdb to v1.19.14 ([#4359](https://github.com/cloudquery/cloudquery/issues/4359)) ([80801b9](https://github.com/cloudquery/cloudquery/commit/80801b9636096295813855b24b960508bae5ddc6))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/dynamodb to v1.17.4 ([#4360](https://github.com/cloudquery/cloudquery/issues/4360)) ([c8daacb](https://github.com/cloudquery/cloudquery/commit/c8daacb8dc776a7d0b1aef263dbe8176b7dd26cc))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/ecr to v1.17.21 ([#4361](https://github.com/cloudquery/cloudquery/issues/4361)) ([c1171d3](https://github.com/cloudquery/cloudquery/commit/c1171d31f84585e8ecb95d52e2130051ff67ce54))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/efs to v1.17.19 ([#4363](https://github.com/cloudquery/cloudquery/issues/4363)) ([05b3dab](https://github.com/cloudquery/cloudquery/commit/05b3dab678df9279476c546c45a2cc25d4a6214d))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/elasticache to v1.23.1 ([#4364](https://github.com/cloudquery/cloudquery/issues/4364)) ([1cd89e7](https://github.com/cloudquery/cloudquery/commit/1cd89e71a60b0cca3d1761bfd892973c47d38040))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/elasticloadbalancing to v1.14.21 ([#4365](https://github.com/cloudquery/cloudquery/issues/4365)) ([b7591b7](https://github.com/cloudquery/cloudquery/commit/b7591b7e39142b01ce6880a7f7a1a801a9cd5691))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2 to v1.18.23 ([#4366](https://github.com/cloudquery/cloudquery/issues/4366)) ([005cc7d](https://github.com/cloudquery/cloudquery/commit/005cc7d89272b01b13ef79183a13c95fea6f34da))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/eventbridge to v1.16.18 ([#4368](https://github.com/cloudquery/cloudquery/issues/4368)) ([e5b5fa8](https://github.com/cloudquery/cloudquery/commit/e5b5fa82ab4c66dd4fa03d6ddfe56368bdf121f4))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/kinesis to v1.15.22 ([#4369](https://github.com/cloudquery/cloudquery/issues/4369)) ([443daad](https://github.com/cloudquery/cloudquery/commit/443daad7cbedf23b77ba7bb7c4a1602681b1c015))
* **deps:** Update plugin-sdk for aws to v1.5.0 ([#4385](https://github.com/cloudquery/cloudquery/issues/4385)) ([0c23dfd](https://github.com/cloudquery/cloudquery/commit/0c23dfde40a1155d7f5a8b9ba20ffab5ca2ffbf3))
* Fix links in Grafana compliance dashboards ([#4338](https://github.com/cloudquery/cloudquery/issues/4338)) ([e71ba56](https://github.com/cloudquery/cloudquery/commit/e71ba567fdd21ae9cf059023795c6765d1766848))
* Update endpoints ([#4335](https://github.com/cloudquery/cloudquery/issues/4335)) ([6917760](https://github.com/cloudquery/cloudquery/commit/69177607bfd5fe1a486a741c9c2d130338dc045b))
* Update some descriptions ([#4371](https://github.com/cloudquery/cloudquery/issues/4371)) ([e5bccf8](https://github.com/cloudquery/cloudquery/commit/e5bccf8652bc265d92c8689bf953a334ab0fe3d5))

## [4.14.0](https://github.com/cloudquery/cloudquery/compare/plugins-source-aws-v4.13.1...plugins-source-aws-v4.14.0) (2022-11-10)


### Features

* Parallelize Athena subresources ([#4118](https://github.com/cloudquery/cloudquery/issues/4118)) ([35111c4](https://github.com/cloudquery/cloudquery/commit/35111c415083b4d9e9c8f01074bc7d9159b50b0a))


### Bug Fixes

* **deps:** Update plugin-sdk for aws to v1.4.0 ([#4226](https://github.com/cloudquery/cloudquery/issues/4226)) ([69238d4](https://github.com/cloudquery/cloudquery/commit/69238d496cb202e2296ede762e38002b1b6a2fb0))
* **deps:** Update plugin-sdk for aws to v1.4.1 ([#4288](https://github.com/cloudquery/cloudquery/issues/4288)) ([570ea0b](https://github.com/cloudquery/cloudquery/commit/570ea0b6d6dd81b08222cc7767aad75df770ad1c))

## [4.13.1](https://github.com/cloudquery/cloudquery/compare/plugins-source-aws-v4.13.0...plugins-source-aws-v4.13.1) (2022-11-10)


### Bug Fixes

* **deps:** Update plugin-sdk for aws to v1.3.2 ([#4193](https://github.com/cloudquery/cloudquery/issues/4193)) ([21c7e3e](https://github.com/cloudquery/cloudquery/commit/21c7e3e8d38f0240e21d8708f402d5afd0800e71))

## [4.13.0](https://github.com/cloudquery/cloudquery/compare/plugins-source-aws-v4.12.0...plugins-source-aws-v4.13.0) (2022-11-10)


### Features

* **aws:** Support RDS Database Versions ([#4121](https://github.com/cloudquery/cloudquery/issues/4121)) ([3b16c05](https://github.com/cloudquery/cloudquery/commit/3b16c05ed72036f124a3a3a737ec9e61668a497c))


### Bug Fixes

* **deps:** Update plugin-sdk for aws to v1.3.1 ([#4144](https://github.com/cloudquery/cloudquery/issues/4144)) ([72d0c45](https://github.com/cloudquery/cloudquery/commit/72d0c4542b9ae99b006663bce7475a676b3a9ba5))

## [4.12.0](https://github.com/cloudquery/cloudquery/compare/plugins-source-aws-v4.11.3...plugins-source-aws-v4.12.0) (2022-11-10)


### Features

* **aws:** Add AWS RAM resources ([#3961](https://github.com/cloudquery/cloudquery/issues/3961)) ([0c16add](https://github.com/cloudquery/cloudquery/commit/0c16addf0534147f5e23cf34d1f9286c681e4aa4))

## [4.11.3](https://github.com/cloudquery/cloudquery/compare/plugins-source-aws-v4.11.2...plugins-source-aws-v4.11.3) (2022-11-09)


### Bug Fixes

* **deps:** Update plugin-sdk for aws to v1.3.0 ([#4067](https://github.com/cloudquery/cloudquery/issues/4067)) ([baaa101](https://github.com/cloudquery/cloudquery/commit/baaa1014199acd361e38142c5615cfde2c28c7bc))

## [4.11.2](https://github.com/cloudquery/cloudquery/compare/plugins-source-aws-v4.11.1...plugins-source-aws-v4.11.2) (2022-11-09)


### Bug Fixes

* **deps:** Update plugin-sdk for aws to v1.2.0 ([#4036](https://github.com/cloudquery/cloudquery/issues/4036)) ([4ee4dcb](https://github.com/cloudquery/cloudquery/commit/4ee4dcbf34d0ed3c4e148936d978c79a68e47a28))

## [4.11.1](https://github.com/cloudquery/cloudquery/compare/plugins-source-aws-v4.11.0...plugins-source-aws-v4.11.1) (2022-11-09)


### Bug Fixes

* **deps:** Update plugin-sdk for aws to v1.1.1 ([#3983](https://github.com/cloudquery/cloudquery/issues/3983)) ([7a4184e](https://github.com/cloudquery/cloudquery/commit/7a4184e213f51d9e18e66d7949895fcb47e581d4))
* **deps:** Update plugin-sdk for csv to v1.1.0 ([#3918](https://github.com/cloudquery/cloudquery/issues/3918)) ([f1acd68](https://github.com/cloudquery/cloudquery/commit/f1acd688fcd90011cc9be1be2285e3fe9369e341))

## [4.11.0](https://github.com/cloudquery/cloudquery/compare/plugins-source-aws-v4.10.0...plugins-source-aws-v4.11.0) (2022-11-09)


### Features

* **aws:** Support AppRunner Tags ([#3968](https://github.com/cloudquery/cloudquery/issues/3968)) ([7c4b633](https://github.com/cloudquery/cloudquery/commit/7c4b6336fa56ef0e34aeabf1a164dc4011828059))

## [4.10.0](https://github.com/cloudquery/cloudquery/compare/plugins-source-aws-v4.9.0...plugins-source-aws-v4.10.0) (2022-11-09)


### Features

* **aws:** Update AWS packages ([#3965](https://github.com/cloudquery/cloudquery/issues/3965)) ([139899d](https://github.com/cloudquery/cloudquery/commit/139899db078490b0cdd9cacea4d0f894960e900e))


### Bug Fixes

* **codegen:** Fix issues related to missing codegen for AWS ([#3954](https://github.com/cloudquery/cloudquery/issues/3954)) ([d485853](https://github.com/cloudquery/cloudquery/commit/d485853a143db091ebe65d71803da4e5b8933d70))
* Update endpoints ([#3966](https://github.com/cloudquery/cloudquery/issues/3966)) ([31f19b4](https://github.com/cloudquery/cloudquery/commit/31f19b4689196bcde25fc125131cf52625fd5133))

## [4.9.0](https://github.com/cloudquery/cloudquery/compare/plugins-source-aws-v4.8.0...plugins-source-aws-v4.9.0) (2022-11-08)


### Features

* **aws:** ECR Image Scan Findings ([#3607](https://github.com/cloudquery/cloudquery/issues/3607)) ([8c53348](https://github.com/cloudquery/cloudquery/commit/8c533482643cd150c20c8ec5e9c77e0c85304fdf))


### Bug Fixes

* **deps:** Update plugin-sdk for aws to v1.1.0 ([#3914](https://github.com/cloudquery/cloudquery/issues/3914)) ([ff566a2](https://github.com/cloudquery/cloudquery/commit/ff566a29baaf0106d850c60d9c2f444d2663d2cf))

## [4.8.0](https://github.com/cloudquery/cloudquery/compare/plugins-source-aws-v4.7.6...plugins-source-aws-v4.8.0) (2022-11-08)


### Features

* **aws:** ECR Repository Policy ([#3844](https://github.com/cloudquery/cloudquery/issues/3844)) ([18956f4](https://github.com/cloudquery/cloudquery/commit/18956f45a525235002af013ef0544757360c245f))

## [4.7.6](https://github.com/cloudquery/cloudquery/compare/plugins-source-aws-v4.7.5...plugins-source-aws-v4.7.6) (2022-11-08)


### Bug Fixes

* **aws:** Elasticache Subnet Groups ([#3835](https://github.com/cloudquery/cloudquery/issues/3835)) ([0c6ee6f](https://github.com/cloudquery/cloudquery/commit/0c6ee6fdea0e668731af33903ec0609083862d3e))
* **deps:** Update dependency cloudquery/cloudquery to v1.6.6 ([#3830](https://github.com/cloudquery/cloudquery/issues/3830)) ([2b30af3](https://github.com/cloudquery/cloudquery/commit/2b30af3b6269e827d4744748c898046330648521))
* **deps:** Update plugin-sdk for aws to v1.0.3 ([#3845](https://github.com/cloudquery/cloudquery/issues/3845)) ([e2e042d](https://github.com/cloudquery/cloudquery/commit/e2e042d4f3c0caa1919cca52589f1a8837fc183d))
* **deps:** Upgrade plugin-sdk to v1.0.4 for plugins ([#3889](https://github.com/cloudquery/cloudquery/issues/3889)) ([6767243](https://github.com/cloudquery/cloudquery/commit/6767243ec70bfae7a4c457bf4b5edf013c54c392))
* Update endpoints ([#3840](https://github.com/cloudquery/cloudquery/issues/3840)) ([5e18993](https://github.com/cloudquery/cloudquery/commit/5e18993f1e64b8c96298421139bc8b6382c4698e))

## [4.7.5](https://github.com/cloudquery/cloudquery/compare/plugins-source-aws-v4.7.4...plugins-source-aws-v4.7.5) (2022-11-07)


### Bug Fixes

* **deps:** Update plugin-sdk for aws to v1 ([#3774](https://github.com/cloudquery/cloudquery/issues/3774)) ([491d7f5](https://github.com/cloudquery/cloudquery/commit/491d7f5ff78c9ce62c5339077fb82aac8dc6a7a8))

## [4.7.4](https://github.com/cloudquery/cloudquery/compare/plugins-source-aws-v4.7.3...plugins-source-aws-v4.7.4) (2022-11-07)


### Bug Fixes

* ServiceQuotas resource performance improvement ([#3783](https://github.com/cloudquery/cloudquery/issues/3783)) ([805d897](https://github.com/cloudquery/cloudquery/commit/805d897d4c2ed78ff4c9b5b86a0d4e70e5d89587))

## [4.7.3](https://github.com/cloudquery/cloudquery/compare/plugins-source-aws-v4.7.2...plugins-source-aws-v4.7.3) (2022-11-07)


### Bug Fixes

* **deps:** Update SDK to v0.13.23 ([#3743](https://github.com/cloudquery/cloudquery/issues/3743)) ([d1a1820](https://github.com/cloudquery/cloudquery/commit/d1a1820f5192d7a18d405b7fdb02b6afa65f009b))

## [4.7.2](https://github.com/cloudquery/cloudquery/compare/plugins-source-aws-v4.7.1...plugins-source-aws-v4.7.2) (2022-11-06)


### Bug Fixes

* **deps:** Update plugin-sdk for aws to v0.13.22 ([#3677](https://github.com/cloudquery/cloudquery/issues/3677)) ([11effa2](https://github.com/cloudquery/cloudquery/commit/11effa22ddd4b12dec1291ea940df4824f8a58e1))

## [4.7.1](https://github.com/cloudquery/cloudquery/compare/plugins-source-aws-v4.7.0...plugins-source-aws-v4.7.1) (2022-11-06)


### Bug Fixes

* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/sagemaker to v1.54.0 ([#3675](https://github.com/cloudquery/cloudquery/issues/3675)) ([59b244d](https://github.com/cloudquery/cloudquery/commit/59b244da7f2ac447b308ad655449a802908fbe34))

## [4.7.0](https://github.com/cloudquery/cloudquery/compare/plugins-source-aws-v4.6.1...plugins-source-aws-v4.7.0) (2022-11-06)


### Features

* Appstream resources ([#3567](https://github.com/cloudquery/cloudquery/issues/3567)) ([084a1d3](https://github.com/cloudquery/cloudquery/commit/084a1d3db2209a768db8b2fb7034fc85f537a617))

## [4.6.1](https://github.com/cloudquery/cloudquery/compare/plugins-source-aws-v4.6.0...plugins-source-aws-v4.6.1) (2022-11-06)


### Bug Fixes

* **deps:** Update plugin-sdk for aws to v0.13.21 ([#3629](https://github.com/cloudquery/cloudquery/issues/3629)) ([5a239b2](https://github.com/cloudquery/cloudquery/commit/5a239b25e29160918f712b9db131242ec302c82d))

## [4.6.0](https://github.com/cloudquery/cloudquery/compare/plugins-source-aws-v4.5.0...plugins-source-aws-v4.6.0) (2022-11-06)


### Features

* AppRunner add support for Connections ([#3602](https://github.com/cloudquery/cloudquery/issues/3602)) ([b6c17a2](https://github.com/cloudquery/cloudquery/commit/b6c17a299bc52d7cf78893803ed1028554c90b26))
* **aws:** AppRunner Resources: VPC Connector, VPC Ingress Connection And Autotscaling ([#3450](https://github.com/cloudquery/cloudquery/issues/3450)) ([f5cd42c](https://github.com/cloudquery/cloudquery/commit/f5cd42c308f99dc62d3e476316dd6a4b69a5a8b7))
* **aws:** AppRunner support ObservabilityConfiguration ([#3603](https://github.com/cloudquery/cloudquery/issues/3603)) ([b93a66d](https://github.com/cloudquery/cloudquery/commit/b93a66d702a927c9dbd3936b7a237006584fb293))
* **aws:** Support identitystore and ssoadmin ([#3005](https://github.com/cloudquery/cloudquery/issues/3005)) ([afa463d](https://github.com/cloudquery/cloudquery/commit/afa463d07fd288a0aa937740c5d39f7884482bd4))


### Bug Fixes

* **aws:** Elasticache Engine Versions PK ([#3562](https://github.com/cloudquery/cloudquery/issues/3562)) ([59a7400](https://github.com/cloudquery/cloudquery/commit/59a740069527946d92ed448288122118d247551d)), closes [#3561](https://github.com/cloudquery/cloudquery/issues/3561)
* **deps:** Update plugin-sdk for aws to v0.13.20 ([#3569](https://github.com/cloudquery/cloudquery/issues/3569)) ([3876311](https://github.com/cloudquery/cloudquery/commit/38763114e431f44c593e6a139eefec1fa586c45b))
* Fix bug in s3_cross_region_replication policy ([#3565](https://github.com/cloudquery/cloudquery/issues/3565)) ([515a7d0](https://github.com/cloudquery/cloudquery/commit/515a7d0ed1d83aa393bf74e60a12cc806a6684b4))
* Fix documentation ([#3608](https://github.com/cloudquery/cloudquery/issues/3608)) ([ea14f06](https://github.com/cloudquery/cloudquery/commit/ea14f06e86928db81e13da9d17169e3f8ef5a3af))
* Update endpoints ([#3605](https://github.com/cloudquery/cloudquery/issues/3605)) ([20b9f4f](https://github.com/cloudquery/cloudquery/commit/20b9f4f9f8ee79fe44ee26ae7f934006be7d96bd))

## [4.5.0](https://github.com/cloudquery/cloudquery/compare/plugins-source-aws-v4.4.0...plugins-source-aws-v4.5.0) (2022-11-03)


### Features

* **aws:** KMS Key Grants, better key fetch ([#3441](https://github.com/cloudquery/cloudquery/issues/3441)) ([98575f0](https://github.com/cloudquery/cloudquery/commit/98575f0551e363e18764da854e8e911a8389112c)), closes [#1906](https://github.com/cloudquery/cloudquery/issues/1906)

## [4.4.0](https://github.com/cloudquery/cloudquery/compare/plugins-source-aws-v4.3.0...plugins-source-aws-v4.4.0) (2022-11-03)


### Features

* Added docdb resources `global_clusters`, `event_subscriptions`, `events_categories`, `pending_maintenance_actions` ([#3277](https://github.com/cloudquery/cloudquery/issues/3277)) ([d14058e](https://github.com/cloudquery/cloudquery/commit/d14058edd602249fd34515ffe79d40a9a7b3d783))

## [4.3.0](https://github.com/cloudquery/cloudquery/compare/plugins-source-aws-v4.2.1...plugins-source-aws-v4.3.0) (2022-11-03)


### Features

* **aws:** Add Support for Service Quotas ([#3489](https://github.com/cloudquery/cloudquery/issues/3489)) ([d9d0dac](https://github.com/cloudquery/cloudquery/commit/d9d0dac94f94f1316d124dcf27d9f73fdddfb729))
* **aws:** Managed Workflows for Apache Airflow (MWAA) environments ([#3431](https://github.com/cloudquery/cloudquery/issues/3431)) ([a8a08c3](https://github.com/cloudquery/cloudquery/commit/a8a08c35fad099f80f09327d23c67fdd468545fa)), closes [#2300](https://github.com/cloudquery/cloudquery/issues/2300)


### Bug Fixes

* **aws-code-gen:** Prefix service name when filtering relations from table list ([#3546](https://github.com/cloudquery/cloudquery/issues/3546)) ([aefaa9a](https://github.com/cloudquery/cloudquery/commit/aefaa9affb5914123ba49818df6ddca17ce8da27))

## [4.2.1](https://github.com/cloudquery/cloudquery/compare/plugins-source-aws-v4.2.0...plugins-source-aws-v4.2.1) (2022-11-03)


### Bug Fixes

* **deps:** Update plugin-sdk for aws to v0.13.19 ([#3500](https://github.com/cloudquery/cloudquery/issues/3500)) ([18ddc2d](https://github.com/cloudquery/cloudquery/commit/18ddc2deaa88d7719341fb3284620840fa48539a))

## [4.2.0](https://github.com/cloudquery/cloudquery/compare/plugins-source-aws-v4.1.0...plugins-source-aws-v4.2.0) (2022-11-03)


### Features

* **aws:** SES ConfigurationSet, ConfigurationSetEventDestination, ContactList, EmailIdentities ([#3475](https://github.com/cloudquery/cloudquery/issues/3475)) ([fbe562b](https://github.com/cloudquery/cloudquery/commit/fbe562b5aaaf0f8ad83bce06f561b8171062d366))

## [4.1.0](https://github.com/cloudquery/cloudquery/compare/plugins-source-aws-v4.0.0...plugins-source-aws-v4.1.0) (2022-11-02)


### Features

* **aws:** Add AppRunner Operations and Custom Domain Support ([#3448](https://github.com/cloudquery/cloudquery/issues/3448)) ([7bc1282](https://github.com/cloudquery/cloudquery/commit/7bc12821cafd53b01c8e750573f9dad2a5071930))

## [4.0.0](https://github.com/cloudquery/cloudquery/compare/plugins-source-aws-v3.8.0...plugins-source-aws-v4.0.0) (2022-11-02)


### ⚠ BREAKING CHANGES

* **aws:** Unified tag structure (#3330)

### Features

* Add eventbridge resources ([#3160](https://github.com/cloudquery/cloudquery/issues/3160)) ([67d3a35](https://github.com/cloudquery/cloudquery/commit/67d3a35329175a7dbeef2ec2e11cf0a20d859fa3))
* **aws:** Fraud Detector support ([#3076](https://github.com/cloudquery/cloudquery/issues/3076)) ([f0e309a](https://github.com/cloudquery/cloudquery/commit/f0e309aeb294521a48f174255447be8d9385998a))
* Migrate cli, plugins and destinations to new type system ([#3323](https://github.com/cloudquery/cloudquery/issues/3323)) ([f265a94](https://github.com/cloudquery/cloudquery/commit/f265a94448ad55c968b26ba8a19681bc81086c11))
* Update AWS Services (new fields) ([#3324](https://github.com/cloudquery/cloudquery/issues/3324)) ([0b65803](https://github.com/cloudquery/cloudquery/commit/0b65803c99797c9e2a5f33c6ac70bd57df14d7ef))


### Bug Fixes

* Add id to aws_cloudfront_cache_policies, region to aws_ec2_transit_gateways ([#3444](https://github.com/cloudquery/cloudquery/issues/3444)) ([41362e2](https://github.com/cloudquery/cloudquery/commit/41362e2d7428a4ee641edb254e3cf7a3c4e8bbdc))
* **aws:** Unified tag structure ([#3330](https://github.com/cloudquery/cloudquery/issues/3330)) ([c9c1e4c](https://github.com/cloudquery/cloudquery/commit/c9c1e4cf0d3c5175799a21829afefd712f322aad))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/credentials to v1.12.23 ([#3378](https://github.com/cloudquery/cloudquery/issues/3378)) ([c33bf73](https://github.com/cloudquery/cloudquery/commit/c33bf73b514f632d0a1c6581720b70d258e5bff1))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/feature/s3/manager to v1.11.37 ([#3379](https://github.com/cloudquery/cloudquery/issues/3379)) ([3f1a71d](https://github.com/cloudquery/cloudquery/commit/3f1a71dc665e3c3714f442cd170f0fdeeda724c8))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/acm to v1.15.2 ([#3380](https://github.com/cloudquery/cloudquery/issues/3380)) ([b4329ee](https://github.com/cloudquery/cloudquery/commit/b4329ee9d5597e942e6dcde01b26a15e368cd7e2))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/apigateway to v1.15.22 ([#3381](https://github.com/cloudquery/cloudquery/issues/3381)) ([5816480](https://github.com/cloudquery/cloudquery/commit/581648021f0a15dfad46691877df1ccbf4698e89))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/apigatewayv2 to v1.12.20 ([#3385](https://github.com/cloudquery/cloudquery/issues/3385)) ([2dd39a6](https://github.com/cloudquery/cloudquery/commit/2dd39a6f30974c9e2c53455d5ccdfbb10cf26112))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/applicationautoscaling to v1.15.20 ([#3386](https://github.com/cloudquery/cloudquery/issues/3386)) ([0702717](https://github.com/cloudquery/cloudquery/commit/0702717c387ef1508ccd2a25283d082313fa6229))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/appsync to v1.15.12 ([#3387](https://github.com/cloudquery/cloudquery/issues/3387)) ([6bb936a](https://github.com/cloudquery/cloudquery/commit/6bb936a450ff8e4a2646d40c60cd4a42a5449664))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/athena to v1.18.12 ([#3388](https://github.com/cloudquery/cloudquery/issues/3388)) ([e95f748](https://github.com/cloudquery/cloudquery/commit/e95f748534f4436ae59a0c50cef8cd31e732f7d4))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/autoscaling to v1.23.18 ([#3389](https://github.com/cloudquery/cloudquery/issues/3389)) ([5af977f](https://github.com/cloudquery/cloudquery/commit/5af977fb4f04ee9d3d3821e089650400d2e3556f))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/backup to v1.17.11 ([#3390](https://github.com/cloudquery/cloudquery/issues/3390)) ([1ecaab7](https://github.com/cloudquery/cloudquery/commit/1ecaab7a4e4ea4bccd0c491ef5908b8a40fc7c79))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/cloudfront to v1.20.7 ([#3391](https://github.com/cloudquery/cloudquery/issues/3391)) ([46465d1](https://github.com/cloudquery/cloudquery/commit/46465d15f40556e73cc03b0a7e9674c8892daf36))
* **deps:** Update plugin-sdk for aws to v0.13.17 ([#3399](https://github.com/cloudquery/cloudquery/issues/3399)) ([f2cd266](https://github.com/cloudquery/cloudquery/commit/f2cd2660ba0339eb4d2631e6e59215ea32906028))
* **deps:** Update plugin-sdk for aws to v0.13.18 ([#3409](https://github.com/cloudquery/cloudquery/issues/3409)) ([92fa576](https://github.com/cloudquery/cloudquery/commit/92fa576382488584341a7ec5e4447741ac053cc5))
* Filtering for DocDB + Neptune ([#3271](https://github.com/cloudquery/cloudquery/issues/3271)) ([8080c6e](https://github.com/cloudquery/cloudquery/commit/8080c6e036c99badb9dbbc355d357000d8c9b903))
* Update endpoints ([#3368](https://github.com/cloudquery/cloudquery/issues/3368)) ([b78c6b1](https://github.com/cloudquery/cloudquery/commit/b78c6b1e42052bcc8de64a734994d28bb84b01d3))

## [3.8.0](https://github.com/cloudquery/cloudquery/compare/plugins-source-aws-v3.7.0...plugins-source-aws-v3.8.0) (2022-10-31)


### Features

* Update all plugins to SDK with metrics and DFS scheduler ([#3286](https://github.com/cloudquery/cloudquery/issues/3286)) ([a35b8e8](https://github.com/cloudquery/cloudquery/commit/a35b8e89d625287a9b9406ff18cfac78ffdb1241))

## [3.7.0](https://github.com/cloudquery/cloudquery/compare/plugins-source-aws-v3.6.0...plugins-source-aws-v3.7.0) (2022-10-30)


### Features

* **aws:** AWS ServiceCatalog ([#3275](https://github.com/cloudquery/cloudquery/issues/3275)) ([d6f955c](https://github.com/cloudquery/cloudquery/commit/d6f955c8838b9cc78f627172fda77b223527a904)), closes [#1410](https://github.com/cloudquery/cloudquery/issues/1410)

## [3.6.0](https://github.com/cloudquery/cloudquery/compare/plugins-source-aws-v3.5.3...plugins-source-aws-v3.6.0) (2022-10-28)


### Features

* Document DB certificates, cluster_parameter_groups, engine_versions, instances, subnet_groups ([#3173](https://github.com/cloudquery/cloudquery/issues/3173)) ([c12306f](https://github.com/cloudquery/cloudquery/commit/c12306f8c2bdee5a7545364f9e85732d40b9c5bb))


### Bug Fixes

* Update endpoints ([#3272](https://github.com/cloudquery/cloudquery/issues/3272)) ([5851a5e](https://github.com/cloudquery/cloudquery/commit/5851a5e68cb7fe316b301010e46049a2b3a43492))
* Update endpoints ([#3281](https://github.com/cloudquery/cloudquery/issues/3281)) ([6a31581](https://github.com/cloudquery/cloudquery/commit/6a31581339218a5bb23a4f0ca9554c62093ddcfc))

## [3.5.3](https://github.com/cloudquery/cloudquery/compare/plugins-source-aws-v3.5.2...plugins-source-aws-v3.5.3) (2022-10-27)


### Bug Fixes

* **aws:** ARN formats ([#3205](https://github.com/cloudquery/cloudquery/issues/3205)) ([8cf8889](https://github.com/cloudquery/cloudquery/commit/8cf88895e5d8d0bb122a26090ce3a41bd5f5cd1a))
* **deps:** Update plugin-sdk for aws to v0.13.14 ([#3211](https://github.com/cloudquery/cloudquery/issues/3211)) ([e7f1c86](https://github.com/cloudquery/cloudquery/commit/e7f1c861ba1af643763563a0d158d7b4ab9c92f3))

## [3.5.2](https://github.com/cloudquery/cloudquery/compare/plugins-source-aws-v3.5.1...plugins-source-aws-v3.5.2) (2022-10-25)


### Bug Fixes

* Update endpoints ([#3199](https://github.com/cloudquery/cloudquery/issues/3199)) ([ed6bc26](https://github.com/cloudquery/cloudquery/commit/ed6bc265281e9ed147b04d5d930cb3a06df73672))

## [3.5.1](https://github.com/cloudquery/cloudquery/compare/plugins-source-aws-v3.5.0...plugins-source-aws-v3.5.1) (2022-10-25)


### Bug Fixes

* Update endpoints ([#3186](https://github.com/cloudquery/cloudquery/issues/3186)) ([48f95b0](https://github.com/cloudquery/cloudquery/commit/48f95b03ac00f575ec9d7b30baebba462a28e1fe))

## [3.5.0](https://github.com/cloudquery/cloudquery/compare/plugins-source-aws-v3.4.3...plugins-source-aws-v3.5.0) (2022-10-23)


### Features

* Added aws docdb clusters ([#3073](https://github.com/cloudquery/cloudquery/issues/3073)) ([51ac6d7](https://github.com/cloudquery/cloudquery/commit/51ac6d7a7f38b84b190e312f5489f350ce9ae002))


### Bug Fixes

* Add warning logs when accounts cannot be instantiated (AWS) ([#3084](https://github.com/cloudquery/cloudquery/issues/3084)) ([33cbb2a](https://github.com/cloudquery/cloudquery/commit/33cbb2a96237342dd57a208e6973c14f6654811a))
* Update endpoints ([#3165](https://github.com/cloudquery/cloudquery/issues/3165)) ([7ab530b](https://github.com/cloudquery/cloudquery/commit/7ab530b5a5789ac1688e0790c0ba2b9c9a9e713d))

## [3.4.3](https://github.com/cloudquery/cloudquery/compare/plugins-source-aws-v3.4.2...plugins-source-aws-v3.4.3) (2022-10-20)


### Bug Fixes

* **deps:** Update plugin-sdk for aws to v0.13.12 ([#3093](https://github.com/cloudquery/cloudquery/issues/3093)) ([9a97682](https://github.com/cloudquery/cloudquery/commit/9a97682252f13c7995fb0e747fbc7208e16a84c0))

## [3.4.2](https://github.com/cloudquery/cloudquery/compare/plugins-source-aws-v3.4.1...plugins-source-aws-v3.4.2) (2022-10-19)


### Bug Fixes

* Update endpoints ([#3071](https://github.com/cloudquery/cloudquery/issues/3071)) ([0fe3a62](https://github.com/cloudquery/cloudquery/commit/0fe3a626efb12c8d423083db3f33a694d221894d))

## [3.4.1](https://github.com/cloudquery/cloudquery/compare/plugins-source-aws-v3.4.0...plugins-source-aws-v3.4.1) (2022-10-19)


### Bug Fixes

* **deps:** Update plugin-sdk for aws to v0.13.11 ([#3011](https://github.com/cloudquery/cloudquery/issues/3011)) ([e0d0e59](https://github.com/cloudquery/cloudquery/commit/e0d0e592c0b88bbb8f5923679fdd9d4a22dc3879))

## [3.4.0](https://github.com/cloudquery/cloudquery/compare/plugins-source-aws-v3.3.0...plugins-source-aws-v3.4.0) (2022-10-19)


### Features

* **aws:** Add Support for App Runner Services ([#2997](https://github.com/cloudquery/cloudquery/issues/2997)) ([b22b06f](https://github.com/cloudquery/cloudquery/commit/b22b06f5508367d022057b12918df0928536d00d))


### Bug Fixes

* Update endpoints ([#3003](https://github.com/cloudquery/cloudquery/issues/3003)) ([776d849](https://github.com/cloudquery/cloudquery/commit/776d849c3594f4646575ba44b3d764e3b13db1c2))

## [3.3.0](https://github.com/cloudquery/cloudquery/compare/plugins-source-aws-v3.2.1...plugins-source-aws-v3.3.0) (2022-10-18)


### Features

* Add Support for Neptune ([#2923](https://github.com/cloudquery/cloudquery/issues/2923)) ([728f54c](https://github.com/cloudquery/cloudquery/commit/728f54c92f25a110a0be20fc0b64270abfdae8ed))

## [3.2.1](https://github.com/cloudquery/cloudquery/compare/plugins-source-aws-v3.2.0...plugins-source-aws-v3.2.1) (2022-10-18)


### Bug Fixes

* EC2 ARN ([#2987](https://github.com/cloudquery/cloudquery/issues/2987)) ([0231d4b](https://github.com/cloudquery/cloudquery/commit/0231d4be3a13703f3b05bc7c2fd9e7d535f62c98))

## [3.2.0](https://github.com/cloudquery/cloudquery/compare/plugins-source-aws-v3.1.0...plugins-source-aws-v3.2.0) (2022-10-18)


### Features

* **aws:** EMR Fetch only running clusters ([#2918](https://github.com/cloudquery/cloudquery/issues/2918)) ([d2f19f5](https://github.com/cloudquery/cloudquery/commit/d2f19f59218132f73612cca9b0328d01cdf60f67))

## [3.1.0](https://github.com/cloudquery/cloudquery/compare/plugins-source-aws-v3.0.1...plugins-source-aws-v3.1.0) (2022-10-18)


### Features

* Add URLs to docs for more AWS resources ([#2729](https://github.com/cloudquery/cloudquery/issues/2729)) ([665aea2](https://github.com/cloudquery/cloudquery/commit/665aea252fc893d5d8f3f5234b67876009144b47))


### Bug Fixes

* **deps:** Update plugin-sdk for aws to v0.13.9 ([#2926](https://github.com/cloudquery/cloudquery/issues/2926)) ([1fe9a43](https://github.com/cloudquery/cloudquery/commit/1fe9a43ac58555ee75e6f26a402e22853532ab7f))
* Improve IAM Policy sync performance ([#2826](https://github.com/cloudquery/cloudquery/issues/2826)) ([13d4689](https://github.com/cloudquery/cloudquery/commit/13d46898c8ea98540e3618666256af92daac9537))

## [3.0.1](https://github.com/cloudquery/cloudquery/compare/plugins-source-aws-v3.0.0...plugins-source-aws-v3.0.1) (2022-10-16)


### Bug Fixes

* Add codegen for AWS regions ([#2701](https://github.com/cloudquery/cloudquery/issues/2701)) ([9b1b2ca](https://github.com/cloudquery/cloudquery/commit/9b1b2ca6cd1be1f7a6a0f0ab834eaad2215a2559))

## [3.0.0](https://github.com/cloudquery/cloudquery/compare/plugins-source-aws-v2.5.2...plugins-source-aws-v3.0.0) (2022-10-16)


### ⚠ BREAKING CHANGES

* More tag cleanup (#2837)

### Features

* **aws:** CQ Policies to Validate IMDSv2 Usage and Enforcement across EC2, Lightsail, and AMIs ([#2807](https://github.com/cloudquery/cloudquery/issues/2807)) ([31d4f79](https://github.com/cloudquery/cloudquery/commit/31d4f797fef81f93fc325f54015e8f601ea39450))
* **aws:** Glacier resources ([#2703](https://github.com/cloudquery/cloudquery/issues/2703)) ([34e9f61](https://github.com/cloudquery/cloudquery/commit/34e9f61b992fb3ab0bd8b0c46b6c3d52104a49c3)), closes [#2583](https://github.com/cloudquery/cloudquery/issues/2583)


### Bug Fixes

* **aws:** Clean up tag fields ([#2776](https://github.com/cloudquery/cloudquery/issues/2776)) ([d071a0e](https://github.com/cloudquery/cloudquery/commit/d071a0e82ccd92c86ef84f68231415153067b5e5))
* **deps:** Update plugin-sdk for aws to v0.13.8 ([#2848](https://github.com/cloudquery/cloudquery/issues/2848)) ([14e2571](https://github.com/cloudquery/cloudquery/commit/14e2571b65d37d797e8db521f7960ec6bfd3a6f9))
* More tag cleanup ([#2837](https://github.com/cloudquery/cloudquery/issues/2837)) ([e781491](https://github.com/cloudquery/cloudquery/commit/e7814914b9a2e72af1d74427d86182e278acce92)), closes [#2836](https://github.com/cloudquery/cloudquery/issues/2836)
* Update endpoints ([#2831](https://github.com/cloudquery/cloudquery/issues/2831)) ([a3b289c](https://github.com/cloudquery/cloudquery/commit/a3b289c1d2200dc2cf70bfc248c021e6a3c6c714))
* Update endpoints ([#2890](https://github.com/cloudquery/cloudquery/issues/2890)) ([2d3b885](https://github.com/cloudquery/cloudquery/commit/2d3b885c306c12588f6253f1f326847ee82017cf))

## [2.5.2](https://github.com/cloudquery/cloudquery/compare/plugins-source-aws-v2.5.1...plugins-source-aws-v2.5.2) (2022-10-13)


### Bug Fixes

* **deps:** Update plugin-sdk for aws to v0.13.7 ([#2778](https://github.com/cloudquery/cloudquery/issues/2778)) ([b0d9e2b](https://github.com/cloudquery/cloudquery/commit/b0d9e2bb91e4ff617712a066e82ce7a67b22dc80))

## [2.5.1](https://github.com/cloudquery/cloudquery/compare/plugins-source-aws-v2.5.0...plugins-source-aws-v2.5.1) (2022-10-12)


### Bug Fixes

* Improperly Configured Service Clients ([#2749](https://github.com/cloudquery/cloudquery/issues/2749)) ([b22a269](https://github.com/cloudquery/cloudquery/commit/b22a269b77a72636bbef0d910631c2305008a416))

## [2.5.0](https://github.com/cloudquery/cloudquery/compare/plugins-source-aws-v2.4.0...plugins-source-aws-v2.5.0) (2022-10-12)


### Features

* Add Support for ECR Public Repositories ([#2645](https://github.com/cloudquery/cloudquery/issues/2645)) ([0e5d974](https://github.com/cloudquery/cloudquery/commit/0e5d9747f32348829eaeffa1fb2b134334f45e44))
* **aws:** Add Column for WAFRegional for ListResourcesForWebACL ([#2648](https://github.com/cloudquery/cloudquery/issues/2648)) ([5c7dfed](https://github.com/cloudquery/cloudquery/commit/5c7dfed99d2b1421209fd39c4ddcf93440fd652f))


### Bug Fixes

* **deps:** Update plugin-sdk for aws to v0.13.5 ([#2660](https://github.com/cloudquery/cloudquery/issues/2660)) ([748a0b3](https://github.com/cloudquery/cloudquery/commit/748a0b3f2cd8429696b38daa386bd0ca32fc3fdf))
* **deps:** Update plugin-sdk for aws to v0.13.6 ([#2717](https://github.com/cloudquery/cloudquery/issues/2717)) ([0fba29f](https://github.com/cloudquery/cloudquery/commit/0fba29faa002b8786febb66ec1a2adaf6c666cc8))
* **deps:** Update plugin-sdk for azure to v0.13.5 ([#2591](https://github.com/cloudquery/cloudquery/issues/2591)) ([c36f60a](https://github.com/cloudquery/cloudquery/commit/c36f60a36fc20823f471ced3ba1726d778bfcda2))
* **deps:** Update plugin-sdk for cloudflare to v0.13.5 ([#2593](https://github.com/cloudquery/cloudquery/issues/2593)) ([ed96887](https://github.com/cloudquery/cloudquery/commit/ed968873a7310daca0dff9fafc94394cca9801e4))
* **deps:** Update plugin-sdk for digitalocean to v0.13.5 ([#2594](https://github.com/cloudquery/cloudquery/issues/2594)) ([5570015](https://github.com/cloudquery/cloudquery/commit/55700155cf1afdbe7e2dd8cc9ae5477a992c1306))
* **deps:** Update plugin-sdk for gcp to v0.13.5 ([#2595](https://github.com/cloudquery/cloudquery/issues/2595)) ([ec17c48](https://github.com/cloudquery/cloudquery/commit/ec17c48959bcdab3d1aed763beee1d0bf37a589e))
* Fix multiplexer for AWS Organizations ([#2727](https://github.com/cloudquery/cloudquery/issues/2727)) ([21bb98e](https://github.com/cloudquery/cloudquery/commit/21bb98e79d1203a5a59a44099c8610c8bf680bbc))

## [2.4.0](https://github.com/cloudquery/cloudquery/compare/plugins-source-aws-v2.3.0...plugins-source-aws-v2.4.0) (2022-10-12)


### Features

* Add API URL for many AWS resources ([#2686](https://github.com/cloudquery/cloudquery/issues/2686)) ([dba6ffa](https://github.com/cloudquery/cloudquery/commit/dba6ffaabe422d5748e066f7b0581c0dca0aa804))
* Add Support for ECR Registries ([#2602](https://github.com/cloudquery/cloudquery/issues/2602)) ([a1c3f4f](https://github.com/cloudquery/cloudquery/commit/a1c3f4f2d90a848273f34064f03beae396ba45f9))

## [2.3.0](https://github.com/cloudquery/cloudquery/compare/plugins-source-aws-v2.2.0...plugins-source-aws-v2.3.0) (2022-10-12)


### Features

* **aws:** Use PreResourceResolver when necessary, remove ListAndDetailResolver ([#2460](https://github.com/cloudquery/cloudquery/issues/2460)) ([340f614](https://github.com/cloudquery/cloudquery/commit/340f61448a64b2aa8cf7f79cfcde2e671f4fd436))

## [2.2.0](https://github.com/cloudquery/cloudquery/compare/plugins-source-aws-v2.1.0...plugins-source-aws-v2.2.0) (2022-10-12)


### Features

* Add support for EC2 reserved instances ([#1912](https://github.com/cloudquery/cloudquery/issues/1912)) ([42120a2](https://github.com/cloudquery/cloudquery/commit/42120a2207cb1889536bdad4211a383751abdd5f))
* Add Support for HSMv2 Clusters and Backups ([#2530](https://github.com/cloudquery/cloudquery/issues/2530)) ([b448bc6](https://github.com/cloudquery/cloudquery/commit/b448bc6cebdf7ddbcdcd9fedb6ebb5fa427821f2))
* Add Support For KMS Aliases ([#2528](https://github.com/cloudquery/cloudquery/issues/2528)) ([2e422dc](https://github.com/cloudquery/cloudquery/commit/2e422dc411836da33fb8c4b6f1a6238de2fee469))


### Bug Fixes

* **aws:** Migrate Grafana dashboards ([#2621](https://github.com/cloudquery/cloudquery/issues/2621)) ([438e439](https://github.com/cloudquery/cloudquery/commit/438e4392c8acd89a6469150636ede39aa9af9304))
* **docs:** Add more info on adding Resource ([#2603](https://github.com/cloudquery/cloudquery/issues/2603)) ([9a29bbd](https://github.com/cloudquery/cloudquery/commit/9a29bbd8f850370f157f29816cbdc11ae3d25e3d))
* Update AWS plugin to SDK v0.13.5 ([#2661](https://github.com/cloudquery/cloudquery/issues/2661)) ([a835034](https://github.com/cloudquery/cloudquery/commit/a8350347e0562a858fc86c1b5c398225f8993642))

## [2.1.0](https://github.com/cloudquery/cloudquery/compare/plugins-source-aws-v2.0.1...plugins-source-aws-v2.1.0) (2022-10-09)


### Features

* **aws:** Use PreResourceResolver to resolve list/describe resources ([#2461](https://github.com/cloudquery/cloudquery/issues/2461)) ([f31ece8](https://github.com/cloudquery/cloudquery/commit/f31ece870bcd2d43c12ec06df711325bee1ad43e))


### Bug Fixes

* **deps:** Update plugin-sdk for aws to v0.12.10 ([#2544](https://github.com/cloudquery/cloudquery/issues/2544)) ([4e4fdb6](https://github.com/cloudquery/cloudquery/commit/4e4fdb6d8f6ae82a1287ead4b02f8f4bd1dc843e))
* **deps:** Update plugin-sdk for aws to v0.12.8 ([#2495](https://github.com/cloudquery/cloudquery/issues/2495)) ([ddb163e](https://github.com/cloudquery/cloudquery/commit/ddb163e18eb802c70dd0bdd6b6beaee31c977bf5))
* **deps:** Update plugin-sdk for aws to v0.12.9 ([#2509](https://github.com/cloudquery/cloudquery/issues/2509)) ([cda0307](https://github.com/cloudquery/cloudquery/commit/cda0307f67c928c4add1afaf0159516fef676f6a))
* Update endpoints ([#2490](https://github.com/cloudquery/cloudquery/issues/2490)) ([624e7a8](https://github.com/cloudquery/cloudquery/commit/624e7a867089e1fb711d6d98ccca118832496206))
* Update endpoints ([#2529](https://github.com/cloudquery/cloudquery/issues/2529)) ([f170e50](https://github.com/cloudquery/cloudquery/commit/f170e50a7deba4761f5b12bee7b92245dd088f0c))

## [2.0.1](https://github.com/cloudquery/cloudquery/compare/plugins-source-aws-v2.0.0...plugins-source-aws-v2.0.1) (2022-10-06)


### Bug Fixes

* **aws:** Re-add Cloudtrail tags ([#2479](https://github.com/cloudquery/cloudquery/issues/2479)) ([c2857d5](https://github.com/cloudquery/cloudquery/commit/c2857d5ec9b2c5ee8f892a5cfeb3ff685ac9fce2))
* **aws:** Re-add Route53 hosted zone tags ([#2480](https://github.com/cloudquery/cloudquery/issues/2480)) ([ebb9eb6](https://github.com/cloudquery/cloudquery/commit/ebb9eb6c2a64dbb01cec2511602b220a88f11185))
* **aws:** Regen glue jobs ([#2476](https://github.com/cloudquery/cloudquery/issues/2476)) ([c787928](https://github.com/cloudquery/cloudquery/commit/c78792861ba655d7165db6652391415df03eade9))
* Fixed crashes when json field is nil ([#2486](https://github.com/cloudquery/cloudquery/issues/2486)) ([7dcda4f](https://github.com/cloudquery/cloudquery/commit/7dcda4f16f18d4b6868a859930d27e48ee30a056))

## [2.0.0](https://github.com/cloudquery/cloudquery/compare/plugins-source-aws-v1.0.1...plugins-source-aws-v2.0.0) (2022-10-06)


### ⚠ BREAKING CHANGES

* Update all AWS Client Versions (#2410)
* Store some JSON columns as jsonb not as a text (#2463)

### Features

* Add new fields to Lightsail and EC2 ([#2403](https://github.com/cloudquery/cloudquery/issues/2403)) ([7e3fc83](https://github.com/cloudquery/cloudquery/commit/7e3fc83c1c0366e3e901f856e8e3e009340cb138))
* **aws:** EMR Clusters to use PreResourceResolver ([#2411](https://github.com/cloudquery/cloudquery/issues/2411)) ([b01e6ca](https://github.com/cloudquery/cloudquery/commit/b01e6caa0bb4d67824fb0ad0091839552d6cef4f))
* **aws:** Update cis 1.5.0 policy ([#1615](https://github.com/cloudquery/cloudquery/issues/1615)) ([0f64196](https://github.com/cloudquery/cloudquery/commit/0f641966ab704cf82fa0bdb1f3cbf18ca015f40f))


### Bug Fixes

* AWS codegen: fix ec2 naming ([#2473](https://github.com/cloudquery/cloudquery/issues/2473)) ([27dd558](https://github.com/cloudquery/cloudquery/commit/27dd55885c5e97251b27f98b6ca3978f3a5cc838))
* **deps:** Update plugin-sdk for aws to v0.12.3 ([#2383](https://github.com/cloudquery/cloudquery/issues/2383)) ([0ed4d5f](https://github.com/cloudquery/cloudquery/commit/0ed4d5f658ea5f89cf097a50047341751335aece))
* **deps:** Update plugin-sdk for aws to v0.12.4 ([#2394](https://github.com/cloudquery/cloudquery/issues/2394)) ([d8c9657](https://github.com/cloudquery/cloudquery/commit/d8c965750d714e143fa65d83a2017b4a648607f7))
* **deps:** Update plugin-sdk for aws to v0.12.6 ([#2416](https://github.com/cloudquery/cloudquery/issues/2416)) ([0539c03](https://github.com/cloudquery/cloudquery/commit/0539c034fb0e2b63224762e2babb3e9e8634f4ab))
* **deps:** Update plugin-sdk for aws to v0.12.7 ([#2445](https://github.com/cloudquery/cloudquery/issues/2445)) ([6e39611](https://github.com/cloudquery/cloudquery/commit/6e396115346232fcdd5fa51d03e1a87da914b504))
* Store some JSON columns as jsonb not as a text ([#2463](https://github.com/cloudquery/cloudquery/issues/2463)) ([5da8a1f](https://github.com/cloudquery/cloudquery/commit/5da8a1ffc9ccdba92b2212f0e108cd4aa15dcafb))
* Update all AWS Client Versions ([#2410](https://github.com/cloudquery/cloudquery/issues/2410)) ([ca34f1c](https://github.com/cloudquery/cloudquery/commit/ca34f1c3ba9570aec84a42b41de3e4b6a18b3cfe))
* Update endpoints ([#2414](https://github.com/cloudquery/cloudquery/issues/2414)) ([62d8802](https://github.com/cloudquery/cloudquery/commit/62d88027fa3eb0e6e06c556bf0ff02b9ec5ff5a0))

## [1.0.1](https://github.com/cloudquery/cloudquery/compare/plugins-source-aws-v1.0.0...plugins-source-aws-v1.0.1) (2022-10-04)


### Bug Fixes

* **aws:** Fix throttle errors ([#2364](https://github.com/cloudquery/cloudquery/issues/2364)) ([6993b2a](https://github.com/cloudquery/cloudquery/commit/6993b2a318baa8585073912f82b7d3603fc4d6a1))

## [1.0.0](https://github.com/cloudquery/cloudquery/compare/plugins/source/aws/v0.13.24...plugins-source-aws-v1.0.0) (2022-10-04)


### ⚠ BREAKING CHANGES

* [Official v1 release](https://www.cloudquery.io/blog/cloudquery-v1-release)

### Features

* [Official v1 release](https://www.cloudquery.io/blog/cloudquery-v1-release)

## [0.15.4-pre.0](https://github.com/cloudquery/cloudquery/compare/plugins-source-aws-v0.15.3-pre.0...plugins-source-aws-v0.15.4-pre.0) (2022-10-04)


### Bug Fixes

* Update endpoints ([#2297](https://github.com/cloudquery/cloudquery/issues/2297)) ([9a5e2c8](https://github.com/cloudquery/cloudquery/commit/9a5e2c84e01f082d03fa655f257afb4a11487444))

## [0.15.3-pre.0](https://github.com/cloudquery/cloudquery/compare/plugins-source-aws-v0.15.2-pre.0...plugins-source-aws-v0.15.3-pre.0) (2022-10-03)


### Bug Fixes

* AWS policies ([#1911](https://github.com/cloudquery/cloudquery/issues/1911)) ([23f1792](https://github.com/cloudquery/cloudquery/commit/23f17927215e854e3220364894e0b7ee5bfb0416))
* **deps:** Update plugin-sdk for aws to v0.11.6 ([#2251](https://github.com/cloudquery/cloudquery/issues/2251)) ([1c60152](https://github.com/cloudquery/cloudquery/commit/1c6015200795cc4e3a768132733e178aff257687))

## [0.15.2-pre.0](https://github.com/cloudquery/cloudquery/compare/plugins-source-aws-v0.15.1-pre.0...plugins-source-aws-v0.15.2-pre.0) (2022-10-03)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/plugin-sdk to v0.11.5 ([#2226](https://github.com/cloudquery/cloudquery/issues/2226)) ([63d7bea](https://github.com/cloudquery/cloudquery/commit/63d7bea93af36f464833bca79770a5d4005f5020))

## [0.15.1-pre.0](https://github.com/cloudquery/cloudquery/compare/plugins-source-aws-v0.15.0-pre.0...plugins-source-aws-v0.15.1-pre.0) (2022-10-02)


### Bug Fixes

* Add ARNs to resources that were missing them ([#2099](https://github.com/cloudquery/cloudquery/issues/2099)) ([122b5c8](https://github.com/cloudquery/cloudquery/commit/122b5c8f53e3bc7411f752fe2e38ab0823aeef94))
* **deps:** Update github.com/gocarina/gocsv digest to ad3251f ([#2178](https://github.com/cloudquery/cloudquery/issues/2178)) ([94e0e83](https://github.com/cloudquery/cloudquery/commit/94e0e8374a02de71d915fce80f739a3da72c1045))
* **deps:** Update golang.org/x/sync digest to 8fcdb60 ([#2170](https://github.com/cloudquery/cloudquery/issues/2170)) ([4cd4259](https://github.com/cloudquery/cloudquery/commit/4cd4259c4e0dc01422824de69ec85415494ea62b))
* **deps:** Update module github.com/aws/aws-sdk-go to v1.44.109 ([#2174](https://github.com/cloudquery/cloudquery/issues/2174)) ([e16760c](https://github.com/cloudquery/cloudquery/commit/e16760c3fb7397bfc77a51b5c29cec41e64d27ca))
* **deps:** Update module github.com/aws/aws-sdk-go-v2 to v1.16.16 ([#2179](https://github.com/cloudquery/cloudquery/issues/2179)) ([de378c0](https://github.com/cloudquery/cloudquery/commit/de378c0f183130caa56c1520a79d1a1c187b2941))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/waf to v1.11.17 ([#2180](https://github.com/cloudquery/cloudquery/issues/2180)) ([7f28bd4](https://github.com/cloudquery/cloudquery/commit/7f28bd465423ca522ac41c1ebb93e14e4ab8b1f5))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/wafregional to v1.12.18 ([#2181](https://github.com/cloudquery/cloudquery/issues/2181)) ([f294e59](https://github.com/cloudquery/cloudquery/commit/f294e596a7eb5125ceb5084cebfc745fde091888))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/xray to v1.13.19 ([#2182](https://github.com/cloudquery/cloudquery/issues/2182)) ([cb54af0](https://github.com/cloudquery/cloudquery/commit/cb54af0d3279e74008bfce7a882138ec689b16d6))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v0.10.2 ([#2048](https://github.com/cloudquery/cloudquery/issues/2048)) ([e407991](https://github.com/cloudquery/cloudquery/commit/e4079914772d8191639b9935aa5970b8e27b082f))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v0.11.0 ([#2135](https://github.com/cloudquery/cloudquery/issues/2135)) ([1729467](https://github.com/cloudquery/cloudquery/commit/1729467b2119555e18b15d73c91cd501ccf7ecb8))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v0.11.2 ([#2162](https://github.com/cloudquery/cloudquery/issues/2162)) ([5701aa5](https://github.com/cloudquery/cloudquery/commit/5701aa5b0a8d04e9e99e3efe6e27d5f7ff29b216))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v0.11.4 ([#2209](https://github.com/cloudquery/cloudquery/issues/2209)) ([1131665](https://github.com/cloudquery/cloudquery/commit/113166541731a755a9cf138ce6635da37b4710a0))
* Generate EC2 resources that were not being generated ([#2124](https://github.com/cloudquery/cloudquery/issues/2124)) ([87347f9](https://github.com/cloudquery/cloudquery/commit/87347f91fca70fb371f11e55e211973e82d9238f))
* Re-add AWS EC2 Instance state_transition_reason_time ([#2087](https://github.com/cloudquery/cloudquery/issues/2087)) ([3cf59df](https://github.com/cloudquery/cloudquery/commit/3cf59df518af3f792ee2915e813d8c679481ac31))
* Update endpoints ([#2102](https://github.com/cloudquery/cloudquery/issues/2102)) ([eb9fc3d](https://github.com/cloudquery/cloudquery/commit/eb9fc3dce9a91086c241230a9b6e5a1e296ebc60))
* Update endpoints ([#2155](https://github.com/cloudquery/cloudquery/issues/2155)) ([6ecc34f](https://github.com/cloudquery/cloudquery/commit/6ecc34fb3b13cf725cf95553fdeadde4dfefed15))
* Use custom resolver for Id fields ([#2117](https://github.com/cloudquery/cloudquery/issues/2117)) ([ae289a6](https://github.com/cloudquery/cloudquery/commit/ae289a6d249c784d4b200a4728a832780de1a4a1))
* Use ParentResourceFieldResolver instead of ParentPathResolver ([#2126](https://github.com/cloudquery/cloudquery/issues/2126)) ([9f06971](https://github.com/cloudquery/cloudquery/commit/9f06971719a59ec430c5f8037df3f613f96b361b))

## [0.15.0-pre.0](https://github.com/cloudquery/cloudquery/compare/plugins-source-aws-v0.14.2-pre.0...plugins-source-aws-v0.15.0-pre.0) (2022-09-26)


### Features

* Add first draft of AWS v2 migration guide ([#1992](https://github.com/cloudquery/cloudquery/issues/1992)) ([6acae93](https://github.com/cloudquery/cloudquery/commit/6acae93bcce9b25e5fc8523550d12cb911cd29d4))
* Provide a decoded policy document field inside aws_iam_policies.policy_version_list ([#2020](https://github.com/cloudquery/cloudquery/issues/2020)) ([e7c51e1](https://github.com/cloudquery/cloudquery/commit/e7c51e170d60dd2a0876fa6b7ce4035d8c32e17d))


### Bug Fixes

* Add missing fields to aws_iam_policies ([#2005](https://github.com/cloudquery/cloudquery/issues/2005)) ([24a22cb](https://github.com/cloudquery/cloudquery/commit/24a22cb9b031eb30a88a27429f12bf02839ccbd5))
* Autofilling aws plugin default config values ([#1935](https://github.com/cloudquery/cloudquery/issues/1935)) ([ddb98a1](https://github.com/cloudquery/cloudquery/commit/ddb98a1881b6b024202e17213e72bf78c7ceb2fd))
* AWS EBS Snapshots attributes column type ([#2075](https://github.com/cloudquery/cloudquery/issues/2075)) ([30ca062](https://github.com/cloudquery/cloudquery/commit/30ca062d904e1701c16395885ccbbc42b6aad253))
* AWS EBS Snapshots attributes column type (take two) ([#2077](https://github.com/cloudquery/cloudquery/issues/2077)) ([899771b](https://github.com/cloudquery/cloudquery/commit/899771b198a9e06cec39223174481308dc48eda1))
* AWS Policy for RDS public accessibility ([#2060](https://github.com/cloudquery/cloudquery/issues/2060)) ([9cde8a0](https://github.com/cloudquery/cloudquery/commit/9cde8a087b8aa8260caf3564cf362031da3fd9d1))
* AWS redshift cluster parameters ([#2063](https://github.com/cloudquery/cloudquery/issues/2063)) ([9b9ff9d](https://github.com/cloudquery/cloudquery/commit/9b9ff9d1145bdb2c527718d50b6203de5cabae91)), closes [#1979](https://github.com/cloudquery/cloudquery/issues/1979)
* **deps:** Update module github.com/cloudquery/plugin-sdk to v0.7.13 ([#1954](https://github.com/cloudquery/cloudquery/issues/1954)) ([2ee4718](https://github.com/cloudquery/cloudquery/commit/2ee4718d3b84defd43218d1958bc669396aafe32))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v0.8.0 ([#1997](https://github.com/cloudquery/cloudquery/issues/1997)) ([4fa40da](https://github.com/cloudquery/cloudquery/commit/4fa40da04b427f864d2dc11f133e5c83e53ce4b6))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v0.8.1 ([#2024](https://github.com/cloudquery/cloudquery/issues/2024)) ([8f88de4](https://github.com/cloudquery/cloudquery/commit/8f88de4b4eaeabae7369ba309e765a252392ee8c))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v0.8.2 ([#2044](https://github.com/cloudquery/cloudquery/issues/2044)) ([9b69b46](https://github.com/cloudquery/cloudquery/commit/9b69b468536521b20b77ec1fc180fc85aeeba376))
* Fix applicationautoscaling multiplexing (v2) ([#2009](https://github.com/cloudquery/cloudquery/issues/2009)) ([4ea6026](https://github.com/cloudquery/cloudquery/commit/4ea6026d2637d0ac31589125f6d05cf5abd6eb9d))
* Fix columns of aws_iam_openid_connect_identity_providers (v2) ([#2001](https://github.com/cloudquery/cloudquery/issues/2001)) ([d378672](https://github.com/cloudquery/cloudquery/commit/d3786728c2d8cda2beddc13254552f1df4b6b5f3))
* Remove underscores in ec_2, s_3 and others ([#1998](https://github.com/cloudquery/cloudquery/issues/1998)) ([0df193f](https://github.com/cloudquery/cloudquery/commit/0df193faae90d6c191d0849ecb0616825203233d))
* Update endpoints ([#2019](https://github.com/cloudquery/cloudquery/issues/2019)) ([0e83552](https://github.com/cloudquery/cloudquery/commit/0e83552e2fdb76181adc399b16c58ad46a9f09e8))
* Update endpoints ([#2037](https://github.com/cloudquery/cloudquery/issues/2037)) ([4719ca4](https://github.com/cloudquery/cloudquery/commit/4719ca4a00dbfa512be2b4bce78ce69a16c9baa0))
* Update endpoints ([#2080](https://github.com/cloudquery/cloudquery/issues/2080)) ([889bef2](https://github.com/cloudquery/cloudquery/commit/889bef23fb01b9b849cff6eb781f0e5864ef5285))

## [0.14.2-pre.0](https://github.com/cloudquery/cloudquery/compare/plugins/source/aws/v0.14.1-pre.0...plugins/source/aws/v0.14.2-pre.0) (2022-09-22)


### Bug Fixes

* Add config_exists field back to aws_s3_accounts (v2) ([#1933](https://github.com/cloudquery/cloudquery/issues/1933)) ([744dfc0](https://github.com/cloudquery/cloudquery/commit/744dfc07ff1da357a4ae4008e8220b7e3e869480))
* Add logging_configuration field to aws_wafv2_web_acls ([#1934](https://github.com/cloudquery/cloudquery/issues/1934)) ([374fa4f](https://github.com/cloudquery/cloudquery/commit/374fa4fbfb720afc2385232ac1ef95df5e38b9c9))
* **aws:** Fix cloudtrail status column not defined ([#1892](https://github.com/cloudquery/cloudquery/issues/1892)) ([e480a56](https://github.com/cloudquery/cloudquery/commit/e480a56d85c3524e5bc635a127cc8f66c92f1d08))
* Fix AWS JSON issues found by new SDK check (v2) ([#1931](https://github.com/cloudquery/cloudquery/issues/1931)) ([b92d76e](https://github.com/cloudquery/cloudquery/commit/b92d76e2418e80d35ebd851a3df313ff47d2535b))
* Fix some issues with AWS S3 resource (v2) ([#1932](https://github.com/cloudquery/cloudquery/issues/1932)) ([ee6e311](https://github.com/cloudquery/cloudquery/commit/ee6e3117d49ae00db2fc3943e4004cd56e6b0b59))
* Update endpoints ([#1941](https://github.com/cloudquery/cloudquery/issues/1941)) ([0c615be](https://github.com/cloudquery/cloudquery/commit/0c615be0e9539b5a6e03125df89a5c6c4b7d816a))

## [0.14.1-pre.0](https://github.com/cloudquery/cloudquery/compare/plugins/source/aws/v0.14.0-pre.0...plugins/source/aws/v0.14.1-pre.0) (2022-09-21)


### Bug Fixes

* Fix some AWS resources with postResolvers ([#1918](https://github.com/cloudquery/cloudquery/issues/1918)) ([7cd8fc9](https://github.com/cloudquery/cloudquery/commit/7cd8fc9ee062b269d4b7e0d03ada8e37dd70ff3a))

## [0.14.0-pre.0](https://github.com/cloudquery/cloudquery/compare/plugins/source/aws-v0.13.24-pre.0...plugins/source/aws/v0.14.0-pre.0) (2022-09-21)


### ⚠ BREAKING CHANGES

* Migrate AWS plugin to v2 (#1774)

### Features

* Add AWS SSM Parameters resource ([#1222](https://github.com/cloudquery/cloudquery/issues/1222)) ([3fde704](https://github.com/cloudquery/cloudquery/commit/3fde7040afaac294ec1cfecad76dafeaa74c9288))
* Add cq-gen config for apigateway ([#1541](https://github.com/cloudquery/cloudquery/issues/1541)) ([15bb1a6](https://github.com/cloudquery/cloudquery/commit/15bb1a6a1c0c2ec0f5362b66b1e099a92ef29a31))
* Add cq-gen hcl and a new field to sqs queues ([#1453](https://github.com/cloudquery/cloudquery/issues/1453)) ([a9584fd](https://github.com/cloudquery/cloudquery/commit/a9584fd2e925a0216db749e07d3673dcf65597a2))
* Add elasticache resources ([#1327](https://github.com/cloudquery/cloudquery/issues/1327)) ([1e5ef30](https://github.com/cloudquery/cloudquery/commit/1e5ef3060ce3ccf788eaea7aff3462e3fcbd0d27))
* Add fsx data repo associations ([#1280](https://github.com/cloudquery/cloudquery/issues/1280)) ([8b02ce3](https://github.com/cloudquery/cloudquery/commit/8b02ce323967f13b890f757691a05bb788cdabe6))
* Add fsx data repo tasks ([#1279](https://github.com/cloudquery/cloudquery/issues/1279)) ([e5774fa](https://github.com/cloudquery/cloudquery/commit/e5774faeefeec66997229d56753eadce4ddc6fcd))
* Add fsx filesystems ([#1277](https://github.com/cloudquery/cloudquery/issues/1277)) ([2c46e9e](https://github.com/cloudquery/cloudquery/commit/2c46e9e6468609d4286b05869f2628c89f56de14))
* Add fsx snapshots ([#1278](https://github.com/cloudquery/cloudquery/issues/1278)) ([750d878](https://github.com/cloudquery/cloudquery/commit/750d87832ab910e8b461dd53677050cc36997277))
* Add fsx storage virtual machines ([#1296](https://github.com/cloudquery/cloudquery/issues/1296)) ([b4f335d](https://github.com/cloudquery/cloudquery/commit/b4f335d3fae84b38062d05dd98a30cfc1ed1a0bd))
* Add fsx volumes ([#1322](https://github.com/cloudquery/cloudquery/issues/1322)) ([9031692](https://github.com/cloudquery/cloudquery/commit/90316928cda66a74874d9d700a1b21c8d994942b))
* Add IAM policy tags ([#1433](https://github.com/cloudquery/cloudquery/issues/1433)) ([70d8365](https://github.com/cloudquery/cloudquery/commit/70d836532de8fb16717c9d66aa921f5b93426faa))
* Add tags for Glue Databases ([#1326](https://github.com/cloudquery/cloudquery/issues/1326)) ([2e083e7](https://github.com/cloudquery/cloudquery/commit/2e083e724dc88263443f93c3153d33e5f97326db))
* Add transfer servers ([#1284](https://github.com/cloudquery/cloudquery/issues/1284)) ([a3bf2bf](https://github.com/cloudquery/cloudquery/commit/a3bf2bfd0c1fc59621bd0815f11b87a2eeb05bae))
* Add website, docs and blog to our main repo ([#1159](https://github.com/cloudquery/cloudquery/issues/1159)) ([dd69948](https://github.com/cloudquery/cloudquery/commit/dd69948feced004497f127d284f2604de0354a1f))
* AWS v2: Comment out max_retries and max_backoff in example config ([#1836](https://github.com/cloudquery/cloudquery/issues/1836)) ([388b7ec](https://github.com/cloudquery/cloudquery/commit/388b7ecc4763c50899c825652cf043cc389af366))
* AWS v2: Split users and credential report users into separate resources ([#1835](https://github.com/cloudquery/cloudquery/issues/1835)) ([d5b772e](https://github.com/cloudquery/cloudquery/commit/d5b772e78ec6a520df6760bcd6fa17df257fbb05))
* Extend sns subscription data ([#1424](https://github.com/cloudquery/cloudquery/issues/1424)) ([63887e7](https://github.com/cloudquery/cloudquery/commit/63887e7cdf1c3b559c9cc201ea02bb1b849ed0f9))
* Implement EC2 Key Pairs ([#1403](https://github.com/cloudquery/cloudquery/issues/1403)) ([#1325](https://github.com/cloudquery/cloudquery/issues/1325)) ([b9d5b74](https://github.com/cloudquery/cloudquery/commit/b9d5b74038f55934684e5623d54a392ea3da2224))
* Migrate AWS plugin to v2 ([#1774](https://github.com/cloudquery/cloudquery/issues/1774)) ([ab4483e](https://github.com/cloudquery/cloudquery/commit/ab4483e32a2f6e7b06c4ca15e01f407cc5c7e158))
* Remove global region ([#1883](https://github.com/cloudquery/cloudquery/issues/1883)) ([99e7089](https://github.com/cloudquery/cloudquery/commit/99e708984c93db4ae68b95f57ee72967b044d8a1))


### Bug Fixes

* **aws:** Nil pointer in config unmarshal ([#1859](https://github.com/cloudquery/cloudquery/issues/1859)) ([64ea261](https://github.com/cloudquery/cloudquery/commit/64ea26112e7455b6363f056051f9c15e664f78ae))
* **deps:** Update github.com/gocarina/gocsv digest to 71f3a5c ([#1660](https://github.com/cloudquery/cloudquery/issues/1660)) ([cf26fcc](https://github.com/cloudquery/cloudquery/commit/cf26fccc87d076ab624df89f658028b7c001467a))
* **deps:** Update golang.org/x/sync digest to 7fc1605 ([#1652](https://github.com/cloudquery/cloudquery/issues/1652)) ([daafae1](https://github.com/cloudquery/cloudquery/commit/daafae1c60c14c90b70c3338a8ff6dc25ba84290))
* **deps:** Update module github.com/cloudquery/cq-gen to v0.0.10 ([#1474](https://github.com/cloudquery/cloudquery/issues/1474)) ([b142e13](https://github.com/cloudquery/cloudquery/commit/b142e135172b1eed1abb2cbec85054ea7f66199d))
* **deps:** Update module github.com/cloudquery/cq-gen to v0.0.11 ([#1491](https://github.com/cloudquery/cloudquery/issues/1491)) ([5140bef](https://github.com/cloudquery/cloudquery/commit/5140bef4aa7c50a97a604db1e92df75ead2893fc))
* **deps:** Update module github.com/cloudquery/cq-gen to v0.0.12 ([#1503](https://github.com/cloudquery/cloudquery/issues/1503)) ([a740719](https://github.com/cloudquery/cloudquery/commit/a7407199c9617784a1834b9d0c42788e03301de5))
* **deps:** Update module github.com/cloudquery/cq-gen to v0.0.9 ([#1286](https://github.com/cloudquery/cloudquery/issues/1286)) ([67ac422](https://github.com/cloudquery/cloudquery/commit/67ac422f392387e674cb70386e612befa5b455f0))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v0.6.2 ([#1838](https://github.com/cloudquery/cloudquery/issues/1838)) ([5b16c59](https://github.com/cloudquery/cloudquery/commit/5b16c59dd415cf0a775dbc38cd62c99b97f04ea5))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v0.6.3 ([#1858](https://github.com/cloudquery/cloudquery/issues/1858)) ([9e3ace7](https://github.com/cloudquery/cloudquery/commit/9e3ace775da2d600968ef4275e9e0013d4dfd825))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v0.6.4 ([#1862](https://github.com/cloudquery/cloudquery/issues/1862)) ([5d141cf](https://github.com/cloudquery/cloudquery/commit/5d141cf6006e26cf240ddf295dda53c16f7386a4))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v0.7.1 ([#1865](https://github.com/cloudquery/cloudquery/issues/1865)) ([474bb70](https://github.com/cloudquery/cloudquery/commit/474bb7081b6e9b6ffc5ac949ed3a664f92083c82))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v0.7.2 ([#1872](https://github.com/cloudquery/cloudquery/issues/1872)) ([49ed26d](https://github.com/cloudquery/cloudquery/commit/49ed26d231c91ac1b5b00cc55d3d0a8a5a6306f7))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v0.7.3 ([#1886](https://github.com/cloudquery/cloudquery/issues/1886)) ([7435d59](https://github.com/cloudquery/cloudquery/commit/7435d593e51ca829d3a328eebc9517e9cb2a4ef0))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v0.7.4 ([#1889](https://github.com/cloudquery/cloudquery/issues/1889)) ([63a5362](https://github.com/cloudquery/cloudquery/commit/63a5362995aa680b291f2411d01e776e884896d4))
* Docs and trigger release ([2028f79](https://github.com/cloudquery/cloudquery/commit/2028f79f3638ba25bab94ceffd65f4cf1ce1d34e))
* **ec2:** Add ARN to key pair ([#1339](https://github.com/cloudquery/cloudquery/issues/1339)) ([189e2c0](https://github.com/cloudquery/cloudquery/commit/189e2c0469c2726cd676c99a051b8a1dd652164b))
* ECS Tags ([#1515](https://github.com/cloudquery/cloudquery/issues/1515)) ([e076217](https://github.com/cloudquery/cloudquery/commit/e076217a5cdddb7dfac35e8a9b9ea94d95abb160))
* Fix date parsing in Credential Report CSVs (v2) ([#1902](https://github.com/cloudquery/cloudquery/issues/1902)) ([00f71ed](https://github.com/cloudquery/cloudquery/commit/00f71ed92fd79c1a05bbb15e6c12510abe1d660b))
* Handle SNS Subscriptions in Pending State ([#1705](https://github.com/cloudquery/cloudquery/issues/1705)) ([6412087](https://github.com/cloudquery/cloudquery/commit/6412087872cd5281b3c889fc4428460108a06088))
* IAM roles ([#1566](https://github.com/cloudquery/cloudquery/issues/1566)) ([09ef545](https://github.com/cloudquery/cloudquery/commit/09ef545a0c9530128bd6960cb54a8dd35e8311da))
* IAM Users ([#1567](https://github.com/cloudquery/cloudquery/issues/1567)) ([1e319da](https://github.com/cloudquery/cloudquery/commit/1e319da8c3f9660305480b36b7c69cf4fa474ba9))
* IAM users PK should be account_id + id, not ARN ([#1876](https://github.com/cloudquery/cloudquery/issues/1876)) ([20096aa](https://github.com/cloudquery/cloudquery/commit/20096aab0f6761f7b53efaff9ad819f64f9260e9))
* Regenerate iam users + some fixes ([#1888](https://github.com/cloudquery/cloudquery/issues/1888)) ([0361ff7](https://github.com/cloudquery/cloudquery/commit/0361ff73f14119d242cbe00c4f4200416ff2246e))
* Small fixes to example.yml and typos in cloudtrail errors ([#1877](https://github.com/cloudquery/cloudquery/issues/1877)) ([c2754fa](https://github.com/cloudquery/cloudquery/commit/c2754faf63e22a59aa976237dfe71bf3b52acb66))
* Tags in users.go ([#1708](https://github.com/cloudquery/cloudquery/issues/1708)) ([ce34eed](https://github.com/cloudquery/cloudquery/commit/ce34eedc1975925d4da9563b99f58d9c29a00eee))
* Update endpoints ([#1273](https://github.com/cloudquery/cloudquery/issues/1273)) ([186a840](https://github.com/cloudquery/cloudquery/commit/186a840bf5702d9845c368a411dd6effd13ee4da))
* Update endpoints ([#1432](https://github.com/cloudquery/cloudquery/issues/1432)) ([4a3c861](https://github.com/cloudquery/cloudquery/commit/4a3c8615ee7a799a5c79b44c6cb55fb2c24591dc))
* Update endpoints ([#1514](https://github.com/cloudquery/cloudquery/issues/1514)) ([eafea83](https://github.com/cloudquery/cloudquery/commit/eafea83f31528e5244eef0eaf621d9e7b9c2b1cf))
* Update endpoints ([#1539](https://github.com/cloudquery/cloudquery/issues/1539)) ([d63fc1d](https://github.com/cloudquery/cloudquery/commit/d63fc1d5c303295a7549de04a7683d104956fa76))
* Update endpoints ([#1563](https://github.com/cloudquery/cloudquery/issues/1563)) ([373ae23](https://github.com/cloudquery/cloudquery/commit/373ae23a3c12da0ce1a5e49c59886f0e29221db0))
* Update endpoints ([#1607](https://github.com/cloudquery/cloudquery/issues/1607)) ([29e5910](https://github.com/cloudquery/cloudquery/commit/29e591085df5fa889905c206d11fd2f0b5d8167b))
* Update endpoints ([#1627](https://github.com/cloudquery/cloudquery/issues/1627)) ([1ec2fef](https://github.com/cloudquery/cloudquery/commit/1ec2fef2adf383f74a9d933c8b2e48d4f4fdc919))
* Update endpoints ([#1703](https://github.com/cloudquery/cloudquery/issues/1703)) ([b001114](https://github.com/cloudquery/cloudquery/commit/b001114f5c1c09ef7782bf75b54f451b5c76d1a4))
* Update endpoints ([#1709](https://github.com/cloudquery/cloudquery/issues/1709)) ([739c188](https://github.com/cloudquery/cloudquery/commit/739c18865c7cfc824cfd5f8535078ffed19f8678))
* Update endpoints ([#1830](https://github.com/cloudquery/cloudquery/issues/1830)) ([bf05794](https://github.com/cloudquery/cloudquery/commit/bf05794e82896925c0da41c31d9ea53b8be00861))
* Update endpoints ([#1880](https://github.com/cloudquery/cloudquery/issues/1880)) ([a64f5aa](https://github.com/cloudquery/cloudquery/commit/a64f5aa329c1f1f3e580b2682f363e7046087491))
* Update endpoints ([#1904](https://github.com/cloudquery/cloudquery/issues/1904)) ([af84989](https://github.com/cloudquery/cloudquery/commit/af849890b672506fc2a6fa504d67a0f024d15826))

## [0.13.24](https://github.com/cloudquery/cloudquery/compare/plugins/source/aws/v0.13.23...plugins/source/aws/v0.13.24) (2022-09-04)


### Bug Fixes

* Handle SNS Subscriptions in Pending State ([#1705](https://github.com/cloudquery/cloudquery/issues/1705)) ([6412087](https://github.com/cloudquery/cloudquery/commit/6412087872cd5281b3c889fc4428460108a06088))
* Tags in users.go ([#1708](https://github.com/cloudquery/cloudquery/issues/1708)) ([ce34eed](https://github.com/cloudquery/cloudquery/commit/ce34eedc1975925d4da9563b99f58d9c29a00eee))
* Update endpoints ([#1703](https://github.com/cloudquery/cloudquery/issues/1703)) ([b001114](https://github.com/cloudquery/cloudquery/commit/b001114f5c1c09ef7782bf75b54f451b5c76d1a4))
* Update endpoints ([#1709](https://github.com/cloudquery/cloudquery/issues/1709)) ([739c188](https://github.com/cloudquery/cloudquery/commit/739c18865c7cfc824cfd5f8535078ffed19f8678))

## [0.13.23](https://github.com/cloudquery/cloudquery/compare/plugins/source/aws/v0.13.22...plugins/source/aws/v0.13.23) (2022-09-01)


### Bug Fixes

* **deps:** Update golang.org/x/sync digest to 7fc1605 ([#1652](https://github.com/cloudquery/cloudquery/issues/1652)) ([daafae1](https://github.com/cloudquery/cloudquery/commit/daafae1c60c14c90b70c3338a8ff6dc25ba84290))

## [0.13.22](https://github.com/cloudquery/cloudquery/compare/plugins/source/aws/v0.13.21...plugins/source/aws/v0.13.22) (2022-08-31)


### Bug Fixes

* Update endpoints ([#1627](https://github.com/cloudquery/cloudquery/issues/1627)) ([1ec2fef](https://github.com/cloudquery/cloudquery/commit/1ec2fef2adf383f74a9d933c8b2e48d4f4fdc919))

## [0.13.21](https://github.com/cloudquery/cloudquery/compare/plugins/source/aws/v0.13.20...plugins/source/aws/v0.13.21) (2022-08-29)


### Bug Fixes

* Update endpoints ([#1607](https://github.com/cloudquery/cloudquery/issues/1607)) ([29e5910](https://github.com/cloudquery/cloudquery/commit/29e591085df5fa889905c206d11fd2f0b5d8167b))

## [0.13.20](https://github.com/cloudquery/cloudquery/compare/plugins/source/aws/v0.13.19...plugins/source/aws/v0.13.20) (2022-08-29)


### Bug Fixes

* Docs and trigger release ([2028f79](https://github.com/cloudquery/cloudquery/commit/2028f79f3638ba25bab94ceffd65f4cf1ce1d34e))

## [0.13.19](https://github.com/cloudquery/cloudquery/compare/plugins/source/aws/v0.13.18...plugins/source/aws/v0.13.19) (2022-08-29)


### Bug Fixes

* IAM roles ([#1566](https://github.com/cloudquery/cloudquery/issues/1566)) ([09ef545](https://github.com/cloudquery/cloudquery/commit/09ef545a0c9530128bd6960cb54a8dd35e8311da))
* IAM Users ([#1567](https://github.com/cloudquery/cloudquery/issues/1567)) ([1e319da](https://github.com/cloudquery/cloudquery/commit/1e319da8c3f9660305480b36b7c69cf4fa474ba9))

## [0.13.18](https://github.com/cloudquery/cloudquery/compare/plugins/source/aws/v0.13.17...plugins/source/aws/v0.13.18) (2022-08-28)


### Bug Fixes

* Update endpoints ([#1563](https://github.com/cloudquery/cloudquery/issues/1563)) ([373ae23](https://github.com/cloudquery/cloudquery/commit/373ae23a3c12da0ce1a5e49c59886f0e29221db0))

## [0.13.17](https://github.com/cloudquery/cloudquery/compare/plugins/source/aws/v0.13.16...plugins/source/aws/v0.13.17) (2022-08-24)


### Features

* Add cq-gen config for apigateway ([#1541](https://github.com/cloudquery/cloudquery/issues/1541)) ([15bb1a6](https://github.com/cloudquery/cloudquery/commit/15bb1a6a1c0c2ec0f5362b66b1e099a92ef29a31))


### Bug Fixes

* Update endpoints ([#1539](https://github.com/cloudquery/cloudquery/issues/1539)) ([d63fc1d](https://github.com/cloudquery/cloudquery/commit/d63fc1d5c303295a7549de04a7683d104956fa76))

## [0.13.16](https://github.com/cloudquery/cloudquery/compare/plugins/source/aws/v0.13.15...plugins/source/aws/v0.13.16) (2022-08-21)


### Features

* Add fsx volumes ([#1322](https://github.com/cloudquery/cloudquery/issues/1322)) ([9031692](https://github.com/cloudquery/cloudquery/commit/90316928cda66a74874d9d700a1b21c8d994942b))

## [0.13.15](https://github.com/cloudquery/cloudquery/compare/plugins/source/aws/v0.13.14...plugins/source/aws/v0.13.15) (2022-08-21)


### Features

* Add cq-gen hcl and a new field to sqs queues ([#1453](https://github.com/cloudquery/cloudquery/issues/1453)) ([a9584fd](https://github.com/cloudquery/cloudquery/commit/a9584fd2e925a0216db749e07d3673dcf65597a2))
* Add fsx data repo associations ([#1280](https://github.com/cloudquery/cloudquery/issues/1280)) ([8b02ce3](https://github.com/cloudquery/cloudquery/commit/8b02ce323967f13b890f757691a05bb788cdabe6))
* Add fsx data repo tasks ([#1279](https://github.com/cloudquery/cloudquery/issues/1279)) ([e5774fa](https://github.com/cloudquery/cloudquery/commit/e5774faeefeec66997229d56753eadce4ddc6fcd))
* Add fsx storage virtual machines ([#1296](https://github.com/cloudquery/cloudquery/issues/1296)) ([b4f335d](https://github.com/cloudquery/cloudquery/commit/b4f335d3fae84b38062d05dd98a30cfc1ed1a0bd))


### Bug Fixes

* **deps:** Update module github.com/cloudquery/cq-gen to v0.0.11 ([#1491](https://github.com/cloudquery/cloudquery/issues/1491)) ([5140bef](https://github.com/cloudquery/cloudquery/commit/5140bef4aa7c50a97a604db1e92df75ead2893fc))
* **deps:** Update module github.com/cloudquery/cq-gen to v0.0.12 ([#1503](https://github.com/cloudquery/cloudquery/issues/1503)) ([a740719](https://github.com/cloudquery/cloudquery/commit/a7407199c9617784a1834b9d0c42788e03301de5))
* ECS Tags ([#1515](https://github.com/cloudquery/cloudquery/issues/1515)) ([e076217](https://github.com/cloudquery/cloudquery/commit/e076217a5cdddb7dfac35e8a9b9ea94d95abb160))
* Update endpoints ([#1514](https://github.com/cloudquery/cloudquery/issues/1514)) ([eafea83](https://github.com/cloudquery/cloudquery/commit/eafea83f31528e5244eef0eaf621d9e7b9c2b1cf))

## [0.13.14](https://github.com/cloudquery/cloudquery/compare/plugins/source/aws/v0.13.13...plugins/source/aws/v0.13.14) (2022-08-18)


### Features

* Add fsx snapshots ([#1278](https://github.com/cloudquery/cloudquery/issues/1278)) ([750d878](https://github.com/cloudquery/cloudquery/commit/750d87832ab910e8b461dd53677050cc36997277))
* Add transfer servers ([#1284](https://github.com/cloudquery/cloudquery/issues/1284)) ([a3bf2bf](https://github.com/cloudquery/cloudquery/commit/a3bf2bfd0c1fc59621bd0815f11b87a2eeb05bae))
* Extend sns subscription data ([#1424](https://github.com/cloudquery/cloudquery/issues/1424)) ([63887e7](https://github.com/cloudquery/cloudquery/commit/63887e7cdf1c3b559c9cc201ea02bb1b849ed0f9))


### Bug Fixes

* **deps:** Update module github.com/cloudquery/cq-gen to v0.0.10 ([#1474](https://github.com/cloudquery/cloudquery/issues/1474)) ([b142e13](https://github.com/cloudquery/cloudquery/commit/b142e135172b1eed1abb2cbec85054ea7f66199d))

## [0.13.13](https://github.com/cloudquery/cloudquery/compare/plugins/source/aws/v0.13.12...plugins/source/aws/v0.13.13) (2022-08-17)


### Features

* Add AWS SSM Parameters resource ([#1222](https://github.com/cloudquery/cloudquery/issues/1222)) ([3fde704](https://github.com/cloudquery/cloudquery/commit/3fde7040afaac294ec1cfecad76dafeaa74c9288))
* Add fsx filesystems ([#1277](https://github.com/cloudquery/cloudquery/issues/1277)) ([2c46e9e](https://github.com/cloudquery/cloudquery/commit/2c46e9e6468609d4286b05869f2628c89f56de14))
* Add IAM policy tags ([#1433](https://github.com/cloudquery/cloudquery/issues/1433)) ([70d8365](https://github.com/cloudquery/cloudquery/commit/70d836532de8fb16717c9d66aa921f5b93426faa))
* Add tags for Glue Databases ([#1326](https://github.com/cloudquery/cloudquery/issues/1326)) ([2e083e7](https://github.com/cloudquery/cloudquery/commit/2e083e724dc88263443f93c3153d33e5f97326db))
* Add website, docs and blog to our main repo ([#1159](https://github.com/cloudquery/cloudquery/issues/1159)) ([dd69948](https://github.com/cloudquery/cloudquery/commit/dd69948feced004497f127d284f2604de0354a1f))


### Bug Fixes

* Update endpoints ([#1432](https://github.com/cloudquery/cloudquery/issues/1432)) ([4a3c861](https://github.com/cloudquery/cloudquery/commit/4a3c8615ee7a799a5c79b44c6cb55fb2c24591dc))

## [0.13.12](https://github.com/cloudquery/cloudquery/compare/plugins/source/aws/v0.13.11...plugins/source/aws/v0.13.12) (2022-08-16)


### Features

* Add elasticache resources ([#1327](https://github.com/cloudquery/cloudquery/issues/1327)) ([1e5ef30](https://github.com/cloudquery/cloudquery/commit/1e5ef3060ce3ccf788eaea7aff3462e3fcbd0d27))

## [0.13.11](https://github.com/cloudquery/cloudquery/compare/plugins/source/aws/v0.13.10...plugins/source/aws/v0.13.11) (2022-08-16)


### Bug Fixes

* **ec2:** Add ARN to key pair ([#1339](https://github.com/cloudquery/cloudquery/issues/1339)) ([189e2c0](https://github.com/cloudquery/cloudquery/commit/189e2c0469c2726cd676c99a051b8a1dd652164b))

## [0.13.10](https://github.com/cloudquery/cloudquery/compare/plugins/source/aws/v0.13.9...plugins/source/aws/v0.13.10) (2022-08-16)


### Features

* Implement EC2 Key Pairs ([#1403](https://github.com/cloudquery/cloudquery/issues/1403)) ([#1325](https://github.com/cloudquery/cloudquery/issues/1325)) ([b9d5b74](https://github.com/cloudquery/cloudquery/commit/b9d5b74038f55934684e5623d54a392ea3da2224))

## [0.13.9](https://github.com/cloudquery/cloudquery/compare/plugins/source/aws-v0.13.8...plugins/source/aws/v0.13.9) (2022-08-15)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/cq-gen to v0.0.9 ([#1286](https://github.com/cloudquery/cloudquery/issues/1286)) ([67ac422](https://github.com/cloudquery/cloudquery/commit/67ac422f392387e674cb70386e612befa5b455f0))
* Update endpoints ([#1273](https://github.com/cloudquery/cloudquery/issues/1273)) ([186a840](https://github.com/cloudquery/cloudquery/commit/186a840bf5702d9845c368a411dd6effd13ee4da))

## [0.13.8](https://github.com/cloudquery/cq-provider-aws/compare/v0.13.7...v0.13.8) (2022-08-12)


### Features

* Add support for EventBridge ([#1407](https://github.com/cloudquery/cq-provider-aws/issues/1407)) ([2d6d138](https://github.com/cloudquery/cq-provider-aws/commit/2d6d138327995c16587800774c24bea6f0de7e98))
* Add support for Inspector Classic & Inspector V2 findings ([#1305](https://github.com/cloudquery/cq-provider-aws/issues/1305)) ([#1412](https://github.com/cloudquery/cq-provider-aws/issues/1412)) ([1f1ae96](https://github.com/cloudquery/cq-provider-aws/commit/1f1ae9631f2087016a70d14b67216a34e0f38dfe))
* Kinesis Firehose Support ([#1359](https://github.com/cloudquery/cq-provider-aws/issues/1359)) ([4324f6b](https://github.com/cloudquery/cq-provider-aws/commit/4324f6b2b09399cc86ebd55788730421b2e298a8))


### Bug Fixes

* Update endpoints ([#1418](https://github.com/cloudquery/cq-provider-aws/issues/1418)) ([b9f8ece](https://github.com/cloudquery/cq-provider-aws/commit/b9f8ecede9f4892b6f600f9149ca31ea16ec32e8))

## [0.13.7](https://github.com/cloudquery/cq-provider-aws/compare/v0.13.6...v0.13.7) (2022-08-11)


### Features

* Add 'elasticache.clusters' resource ([#1400](https://github.com/cloudquery/cq-provider-aws/issues/1400)) ([d27b0d9](https://github.com/cloudquery/cq-provider-aws/commit/d27b0d93b59dad9bac18b72d1123450fb84e5a75))
* Add support for appsync apis ([#1393](https://github.com/cloudquery/cq-provider-aws/issues/1393)) ([2197701](https://github.com/cloudquery/cq-provider-aws/commit/21977016f62cb43cbc158e7d08d4afe6ee222c10))
* Add support for resource groups ([#1396](https://github.com/cloudquery/cq-provider-aws/issues/1396)) ([dc6aeab](https://github.com/cloudquery/cq-provider-aws/commit/dc6aeab7ce42c8ec699075a6d1d2a221a6590733))
* Added glue table indexes ([#1377](https://github.com/cloudquery/cq-provider-aws/issues/1377)) ([b008f1b](https://github.com/cloudquery/cq-provider-aws/commit/b008f1ba8324e921e05095d130dfd24b8dc3042e))
* Implement Glue registries resource ([#1334](https://github.com/cloudquery/cq-provider-aws/issues/1334)) ([5e20e88](https://github.com/cloudquery/cq-provider-aws/commit/5e20e88f30ffe4597c00e1a219f1239b8a8cdb13))
* More Resource Simplification ([#1399](https://github.com/cloudquery/cq-provider-aws/issues/1399)) ([902c8e0](https://github.com/cloudquery/cq-provider-aws/commit/902c8e082be63678e46b8dcd5b3ed3e090f1e847))
* Remove regional override ([#1276](https://github.com/cloudquery/cq-provider-aws/issues/1276)) ([7f8025a](https://github.com/cloudquery/cq-provider-aws/commit/7f8025a1b355bbc3f554fe8e098286cd8b568324))
* Simplify Resources ([#1385](https://github.com/cloudquery/cq-provider-aws/issues/1385)) ([1f7eab8](https://github.com/cloudquery/cq-provider-aws/commit/1f7eab8804692b9f69df3ac3ff776e89dc5f08a0))


### Bug Fixes

* **build:** Don't filter paths and enforce //check-for-changes on new cq-gen config files ([#1401](https://github.com/cloudquery/cq-provider-aws/issues/1401)) ([17e40a3](https://github.com/cloudquery/cq-provider-aws/commit/17e40a3a7abce210a89c8abf28493d7c9b8e9471))
* **build:** Fix name of script in Github workflow ([#1405](https://github.com/cloudquery/cq-provider-aws/issues/1405)) ([b4ed653](https://github.com/cloudquery/cq-provider-aws/commit/b4ed6531f4d2936d025c291d617e95fe636f3398))
* Update endpoints ([#1402](https://github.com/cloudquery/cq-provider-aws/issues/1402)) ([bb81947](https://github.com/cloudquery/cq-provider-aws/commit/bb819477a1c81c0b351e06544ff2b892b782c0c0))

## [0.13.6](https://github.com/cloudquery/cq-provider-aws/compare/v0.13.5...v0.13.6) (2022-08-07)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/cq-provider-sdk to v0.14.7 ([#1394](https://github.com/cloudquery/cq-provider-aws/issues/1394)) ([328c7cc](https://github.com/cloudquery/cq-provider-aws/commit/328c7ccc048e76aa1374b13c5ff7c972b248ecf8))

## [0.13.5](https://github.com/cloudquery/cq-provider-aws/compare/v0.13.4...v0.13.5) (2022-08-05)


### Features

* Add CloudWatch Logs Log Group Resource ([#1317](https://github.com/cloudquery/cq-provider-aws/issues/1317)) ([5ae109e](https://github.com/cloudquery/cq-provider-aws/commit/5ae109e1afce58c41d7de7b7271c804fe3f91201))
* Add tags for Cloudwatch alarms ([#1374](https://github.com/cloudquery/cq-provider-aws/issues/1374)) ([07bd11b](https://github.com/cloudquery/cq-provider-aws/commit/07bd11bd4cf393340ccfdb433765414867537ce5))
* Added glue classifiers ([#1389](https://github.com/cloudquery/cq-provider-aws/issues/1389)) ([d7722e5](https://github.com/cloudquery/cq-provider-aws/commit/d7722e502ce8102dc5122a7b24f0ebf2fe91c8b5))
* Added glue connections ([#1388](https://github.com/cloudquery/cq-provider-aws/issues/1388)) ([516721e](https://github.com/cloudquery/cq-provider-aws/commit/516721eda3f5c10da165a4eb6621b61ed3db6213))
* Added glue crawlers ([#1363](https://github.com/cloudquery/cq-provider-aws/issues/1363)) ([177e690](https://github.com/cloudquery/cq-provider-aws/commit/177e69073600ebd23d1bad84547a338977cbabae))
* Added glue data catalog encryption settings ([#1356](https://github.com/cloudquery/cq-provider-aws/issues/1356)) ([a6c6246](https://github.com/cloudquery/cq-provider-aws/commit/a6c6246e24c3b72d5588b03f4cb4b7fb427ba62b))
* Added glue dev endpoints ([#1361](https://github.com/cloudquery/cq-provider-aws/issues/1361)) ([2cbec07](https://github.com/cloudquery/cq-provider-aws/commit/2cbec07282f09866438472976d0fec2df76b3443))
* Added glue ml transforms ([#1365](https://github.com/cloudquery/cq-provider-aws/issues/1365)) ([baffddf](https://github.com/cloudquery/cq-provider-aws/commit/baffddf6b6089e517213bcd282ce9452850a2991))
* Added glue security configurations ([#1382](https://github.com/cloudquery/cq-provider-aws/issues/1382)) ([4d1cf3e](https://github.com/cloudquery/cq-provider-aws/commit/4d1cf3e49483cd47f5e8a42aa4bc91b90e2b5794))
* Implement Glue Triggers resource ([#1322](https://github.com/cloudquery/cq-provider-aws/issues/1322)) ([8261048](https://github.com/cloudquery/cq-provider-aws/commit/82610486e3190db9a599fe61b2c9ab77ea3d3e74))
* Update EC2 Services ([#1383](https://github.com/cloudquery/cq-provider-aws/issues/1383)) ([cc33980](https://github.com/cloudquery/cq-provider-aws/commit/cc33980c893ba43c7ebe283e1a88b3f6cb11583c))


### Bug Fixes

* **deps:** Update module github.com/cloudquery/cq-gen to v0.0.8 ([#1392](https://github.com/cloudquery/cq-provider-aws/issues/1392)) ([90d8fa5](https://github.com/cloudquery/cq-provider-aws/commit/90d8fa52d71c90b228d539d19f06089fda68a02c))
* **deps:** Update tubone24/update_release digest to 2146f15 ([#1349](https://github.com/cloudquery/cq-provider-aws/issues/1349)) ([37dee52](https://github.com/cloudquery/cq-provider-aws/commit/37dee5271bee28643e97566c61bbc1c8c19ffe82))
* Update endpoints ([#1375](https://github.com/cloudquery/cq-provider-aws/issues/1375)) ([0b20ba8](https://github.com/cloudquery/cq-provider-aws/commit/0b20ba86c11e75a65c5d1a1703484456c8a02dd8))
* Update endpoints ([#1386](https://github.com/cloudquery/cq-provider-aws/issues/1386)) ([9af45f7](https://github.com/cloudquery/cq-provider-aws/commit/9af45f700c95df17b218672bd8c3bac35b869c73))

## [0.13.4](https://github.com/cloudquery/cq-provider-aws/compare/v0.13.3...v0.13.4) (2022-08-02)


### Features

* Add Kinesis Data Stream support ([#1348](https://github.com/cloudquery/cq-provider-aws/issues/1348)) ([767bfab](https://github.com/cloudquery/cq-provider-aws/commit/767bfaba0d382d6971023bdd13535d0a6cd95ec6))
* Add Tags for ECR Repo ([#1369](https://github.com/cloudquery/cq-provider-aws/issues/1369)) ([3b31598](https://github.com/cloudquery/cq-provider-aws/commit/3b31598782e450b04c31a95938b2df9906828adf))
* Added glue databases and tables ([#1345](https://github.com/cloudquery/cq-provider-aws/issues/1345)) ([0284a37](https://github.com/cloudquery/cq-provider-aws/commit/0284a37e7ebafbdcf15140a1ccdedde09f0bb13b))
* Added glue jobs ([#1352](https://github.com/cloudquery/cq-provider-aws/issues/1352)) ([562a6b3](https://github.com/cloudquery/cq-provider-aws/commit/562a6b334ee077f5fab76cdd3321ea7ddfec3f91))
* Column Resolvers    ([#1301](https://github.com/cloudquery/cq-provider-aws/issues/1301)) ([9b2dbed](https://github.com/cloudquery/cq-provider-aws/commit/9b2dbed888f44b4430c66051ce9e30d9ecea7673))


### Bug Fixes

* **deps:** Update module github.com/cloudquery/cq-gen to v0.0.7 ([#1362](https://github.com/cloudquery/cq-provider-aws/issues/1362)) ([3060854](https://github.com/cloudquery/cq-provider-aws/commit/3060854773cca20f7b18980b02193bb15273649e))
* **deps:** Update module github.com/hashicorp/go-hclog to v1.2.2 ([#1350](https://github.com/cloudquery/cq-provider-aws/issues/1350)) ([82ec301](https://github.com/cloudquery/cq-provider-aws/commit/82ec301d9eed9fe00812d7e8d0d89b3ff753faaa))
* Update endpoints ([#1347](https://github.com/cloudquery/cq-provider-aws/issues/1347)) ([3191f3e](https://github.com/cloudquery/cq-provider-aws/commit/3191f3e08f2c1142e5a79a76ec6e5b0a0da9f30c))

## [0.13.3](https://github.com/cloudquery/cq-provider-aws/compare/v0.13.2...v0.13.3) (2022-07-28)


### Features

* Add Glue Workflows resource ([#1310](https://github.com/cloudquery/cq-provider-aws/issues/1310)) ([bd45348](https://github.com/cloudquery/cq-provider-aws/commit/bd453484478bc86cac4ca27c1d9938a0730bfe6a))


### Bug Fixes

* **deps:** Update module github.com/cloudquery/cq-provider-sdk to v0.14.5 ([#1315](https://github.com/cloudquery/cq-provider-aws/issues/1315)) ([dcd3f17](https://github.com/cloudquery/cq-provider-aws/commit/dcd3f1757d91f22c758700504288df36c013d8c2))
* **deps:** Update module github.com/cloudquery/cq-provider-sdk to v0.14.6 ([#1332](https://github.com/cloudquery/cq-provider-aws/issues/1332)) ([cdb9d0a](https://github.com/cloudquery/cq-provider-aws/commit/cdb9d0ae049f050a2952acace7076820a268ca33))
* Properly Handle Error in data_catalogs ([#1326](https://github.com/cloudquery/cq-provider-aws/issues/1326)) ([26a8339](https://github.com/cloudquery/cq-provider-aws/commit/26a83395ed007ed8a888aea45b6037b25e0af38f))
* Update endpoints ([#1335](https://github.com/cloudquery/cq-provider-aws/issues/1335)) ([b493edc](https://github.com/cloudquery/cq-provider-aws/commit/b493edc6ebcf1dfb2ad5a48582cf1a4964f6bc97))

## [0.13.2](https://github.com/cloudquery/cq-provider-aws/compare/v0.13.1...v0.13.2) (2022-07-27)


### Features

* Add support for EC2 instance types ([#1278](https://github.com/cloudquery/cq-provider-aws/issues/1278)) ([b49ae24](https://github.com/cloudquery/cq-provider-aws/commit/b49ae24b4ce831c727c092bf84b6556cdc00e8a7))
* Added lightsail container_services ([#1295](https://github.com/cloudquery/cq-provider-aws/issues/1295)) ([ed3e028](https://github.com/cloudquery/cq-provider-aws/commit/ed3e02819d169bc02b3702cea3c67a91089175eb))
* Added lightsail distributions ([#1294](https://github.com/cloudquery/cq-provider-aws/issues/1294)) ([6e0c06f](https://github.com/cloudquery/cq-provider-aws/commit/6e0c06fe10b295e4d2851fd7191c57b2b9318518))
* Added lightsail instances relations ([#1266](https://github.com/cloudquery/cq-provider-aws/issues/1266)) ([d6ecaae](https://github.com/cloudquery/cq-provider-aws/commit/d6ecaae27c529e735b209354df907626f81485cb))


### Bug Fixes

* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/codebuild to v1.19.8 ([#1282](https://github.com/cloudquery/cq-provider-aws/issues/1282)) ([109656f](https://github.com/cloudquery/cq-provider-aws/commit/109656feed7060f9727be6738f98898b453eb3da))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/codepipeline to v1.13.8 ([#1283](https://github.com/cloudquery/cq-provider-aws/issues/1283)) ([9615bf3](https://github.com/cloudquery/cq-provider-aws/commit/9615bf3b14b35d00d5013299ac6e6c4d2a7a316e))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/cognitoidentity to v1.13.8 ([#1284](https://github.com/cloudquery/cq-provider-aws/issues/1284)) ([b616009](https://github.com/cloudquery/cq-provider-aws/commit/b6160096c28f9da4c442bacf474a43cff93939a3))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider to v1.17.3 ([#1285](https://github.com/cloudquery/cq-provider-aws/issues/1285)) ([69875b9](https://github.com/cloudquery/cq-provider-aws/commit/69875b97b19b53ec1ac258dd0526ba7a093e1495))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/dax to v1.11.8 ([#1286](https://github.com/cloudquery/cq-provider-aws/issues/1286)) ([5a4b29f](https://github.com/cloudquery/cq-provider-aws/commit/5a4b29f0c054374b1349246fad91c3ca2d02a854))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/directconnect to v1.17.8 ([#1287](https://github.com/cloudquery/cq-provider-aws/issues/1287)) ([02f4d0c](https://github.com/cloudquery/cq-provider-aws/commit/02f4d0c746a54efb7c9b4f6feffaef5f18e9db72))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/dynamodb to v1.15.9 ([#1288](https://github.com/cloudquery/cq-provider-aws/issues/1288)) ([af3b414](https://github.com/cloudquery/cq-provider-aws/commit/af3b4147309df858647170fabbd3446fd246dbf6))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/ecr to v1.17.8 ([#1289](https://github.com/cloudquery/cq-provider-aws/issues/1289)) ([d3fa5d0](https://github.com/cloudquery/cq-provider-aws/commit/d3fa5d0d92084c14cef03bb1b8355c0eaaf2803e))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/ecs to v1.18.11 ([#1290](https://github.com/cloudquery/cq-provider-aws/issues/1290)) ([7180a9d](https://github.com/cloudquery/cq-provider-aws/commit/7180a9da4f5b45952958ca0aca358dca0cfa4303))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/efs to v1.17.6 ([#1291](https://github.com/cloudquery/cq-provider-aws/issues/1291)) ([8b50d24](https://github.com/cloudquery/cq-provider-aws/commit/8b50d2415c70c474ebb0c0b01ae619a672bb97e1))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/eks to v1.21.4 ([#1292](https://github.com/cloudquery/cq-provider-aws/issues/1292)) ([4b20c72](https://github.com/cloudquery/cq-provider-aws/commit/4b20c72da8d3553f9a6cc392a83d319892013a61))
* **deps:** Update tubone24/update_release digest to e8cd303 ([#1281](https://github.com/cloudquery/cq-provider-aws/issues/1281)) ([409ea75](https://github.com/cloudquery/cq-provider-aws/commit/409ea75cf8954eb0623c5d9bbab13f500e654b9d))
* Update endpoints ([#1271](https://github.com/cloudquery/cq-provider-aws/issues/1271)) ([4738faa](https://github.com/cloudquery/cq-provider-aws/commit/4738faa6e0b315552afa833aa28307f93d2c2dc9))

## [0.13.1](https://github.com/cloudquery/cq-provider-aws/compare/v0.13.0...v0.13.1) (2022-07-22)


### Features

* Added lightsail alarms ([#1242](https://github.com/cloudquery/cq-provider-aws/issues/1242)) ([19e3476](https://github.com/cloudquery/cq-provider-aws/commit/19e3476ae032765acb50a7f934a3da311fa9fadc))
* Added lightsail certificates ([#1245](https://github.com/cloudquery/cq-provider-aws/issues/1245)) ([0ee77fc](https://github.com/cloudquery/cq-provider-aws/commit/0ee77fcbe0124a21d4eb6ae2ca49954f195a840b))
* Added lightsail database snapshots ([#1263](https://github.com/cloudquery/cq-provider-aws/issues/1263)) ([6749225](https://github.com/cloudquery/cq-provider-aws/commit/6749225829e3262d83e9ba44d15a0f055dc0f8e2))
* Added lightsail databases ([#1251](https://github.com/cloudquery/cq-provider-aws/issues/1251)) ([72c2702](https://github.com/cloudquery/cq-provider-aws/commit/72c2702bcd0046f5354e20b9d6d953eb063b1bda))
* Added lightsail disks ([#1240](https://github.com/cloudquery/cq-provider-aws/issues/1240)) ([fab23e0](https://github.com/cloudquery/cq-provider-aws/commit/fab23e0737078c6cd5543aef0a4da8b8e90895d7))
* Added lightsail load balancers ([#1254](https://github.com/cloudquery/cq-provider-aws/issues/1254)) ([382bb7f](https://github.com/cloudquery/cq-provider-aws/commit/382bb7fc96d99b7f82eefbe241a4fa1ed54d7ee6))
* Added lightsail static ips ([#1248](https://github.com/cloudquery/cq-provider-aws/issues/1248)) ([62a85f7](https://github.com/cloudquery/cq-provider-aws/commit/62a85f7ec13bd169a47ed939599820c82f9a2cf4))
* Remove non standard List/Detail implementations ([#1237](https://github.com/cloudquery/cq-provider-aws/issues/1237)) ([004a544](https://github.com/cloudquery/cq-provider-aws/commit/004a5444741d3e7d6fa1020449ce3ab6be6e5213))


### Bug Fixes

* Correctly use pagination on EC2 instance list response ([#1270](https://github.com/cloudquery/cq-provider-aws/issues/1270)) ([aebcc3c](https://github.com/cloudquery/cq-provider-aws/commit/aebcc3cb483bcbc81626de51681513ae3fd8580b))
* Default execution_time for main policy ([#1264](https://github.com/cloudquery/cq-provider-aws/issues/1264)) ([f8f0590](https://github.com/cloudquery/cq-provider-aws/commit/f8f0590c842f1040915359b30110a6d134622001))
* **deps:** Update module github.com/cloudquery/cq-provider-sdk to v0.14.3 ([#1265](https://github.com/cloudquery/cq-provider-aws/issues/1265)) ([4ebc1d9](https://github.com/cloudquery/cq-provider-aws/commit/4ebc1d9df187c6deabe800eeddd77dc26f7ea8dc))
* **deps:** Update module github.com/cloudquery/cq-provider-sdk to v0.14.4 ([#1268](https://github.com/cloudquery/cq-provider-aws/issues/1268)) ([1bd574d](https://github.com/cloudquery/cq-provider-aws/commit/1bd574d15070e0e6e5fa65e73a4c76c5e45c5532))
* Ignore CF distributions NoSuchResource on ListTags ([#1238](https://github.com/cloudquery/cq-provider-aws/issues/1238)) ([01efd1a](https://github.com/cloudquery/cq-provider-aws/commit/01efd1af29e56e5a8c6aae9b87ab63868451c426))
* Ignore some not founds in lambda functions ([#1252](https://github.com/cloudquery/cq-provider-aws/issues/1252)) ([865e1c6](https://github.com/cloudquery/cq-provider-aws/commit/865e1c66ee8058252af69d602ad76cbb5e0038e5))
* Lightsail alarms adjusted ([#1260](https://github.com/cloudquery/cq-provider-aws/issues/1260)) ([6f1e3a0](https://github.com/cloudquery/cq-provider-aws/commit/6f1e3a01fe119adf0e5e537a78bebed5ddc4c292))

## [0.13.0](https://github.com/cloudquery/cq-provider-aws/compare/v0.12.29...v0.13.0) (2022-07-21)


### ⚠ BREAKING CHANGES

* Update SDK to v0.14.1 (#1239)

### Features

* Add asset inventory and public endpoint dashboards ([#1257](https://github.com/cloudquery/cq-provider-aws/issues/1257)) ([5164b9f](https://github.com/cloudquery/cq-provider-aws/commit/5164b9fdafd1a8979450d1352d048fd0caf37a07))
* Add compliance dashboard ([#1255](https://github.com/cloudquery/cq-provider-aws/issues/1255)) ([8d3e0a1](https://github.com/cloudquery/cq-provider-aws/commit/8d3e0a1c3a8be810c8ad60f7853816ce2d8b1893))
* Added lightsail buckets ([#1097](https://github.com/cloudquery/cq-provider-aws/issues/1097)) ([74b216a](https://github.com/cloudquery/cq-provider-aws/commit/74b216a9471a29c295c16628f56a94d0d2419a4a))
* Policies ([#1220](https://github.com/cloudquery/cq-provider-aws/issues/1220)) ([8a2cb92](https://github.com/cloudquery/cq-provider-aws/commit/8a2cb9247dac084449f3fe293c670e633f309e22))


### Bug Fixes

* Another bucket not found error ([#1247](https://github.com/cloudquery/cq-provider-aws/issues/1247)) ([5216cd0](https://github.com/cloudquery/cq-provider-aws/commit/5216cd0950fc7efadce737faa6069a9cfe5a06ce))
* Check For Nil  ([#1223](https://github.com/cloudquery/cq-provider-aws/issues/1223)) ([bb2c120](https://github.com/cloudquery/cq-provider-aws/commit/bb2c120cae1b4f1239984faa3078656f5e66cee2))
* Classify DNS errors as user ([#1190](https://github.com/cloudquery/cq-provider-aws/issues/1190)) ([c509dae](https://github.com/cloudquery/cq-provider-aws/commit/c509daea1f1130dc4b92d76345f75d04686ffbd5))
* **deps:** Update github.com/gocarina/gocsv digest to 8b2118d ([#1202](https://github.com/cloudquery/cq-provider-aws/issues/1202)) ([f9e9ff2](https://github.com/cloudquery/cq-provider-aws/commit/f9e9ff222975d86a145444b6218857076518adfa))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/feature/s3/manager to v1.11.20 ([#1226](https://github.com/cloudquery/cq-provider-aws/issues/1226)) ([9bc0008](https://github.com/cloudquery/cq-provider-aws/commit/9bc00084a0d12ca89c167f8eabe5d41fcb71e973))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/apigatewayv2 to v1.12.8 ([#1227](https://github.com/cloudquery/cq-provider-aws/issues/1227)) ([604f7e5](https://github.com/cloudquery/cq-provider-aws/commit/604f7e56c53d4269f37dcd5c4ad024bdefec5b61))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/applicationautoscaling to v1.15.8 ([#1228](https://github.com/cloudquery/cq-provider-aws/issues/1228)) ([a7de9c8](https://github.com/cloudquery/cq-provider-aws/commit/a7de9c89b042a4e3a6236a20afd052c70bb15cf8))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/autoscaling to v1.23.5 ([#1229](https://github.com/cloudquery/cq-provider-aws/issues/1229)) ([164871f](https://github.com/cloudquery/cq-provider-aws/commit/164871f8d63a757adb0b4f8f742ac002dea9189c))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/cloudfront to v1.18.4 ([#1230](https://github.com/cloudquery/cq-provider-aws/issues/1230)) ([3e39351](https://github.com/cloudquery/cq-provider-aws/commit/3e3935123d0cf09f4695ce7cec050d432c0a3818))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/cloudtrail to v1.16.4 ([#1231](https://github.com/cloudquery/cq-provider-aws/issues/1231)) ([93e81ec](https://github.com/cloudquery/cq-provider-aws/commit/93e81ecaac0a7168507e727f8811e20c0dc774ad))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/cloudwatch to v1.18.6 ([#1232](https://github.com/cloudquery/cq-provider-aws/issues/1232)) ([f01a61d](https://github.com/cloudquery/cq-provider-aws/commit/f01a61d6cac94c141095516837851644199c250b))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs to v1.15.10 ([#1233](https://github.com/cloudquery/cq-provider-aws/issues/1233)) ([11f5e3b](https://github.com/cloudquery/cq-provider-aws/commit/11f5e3b27059c3723efbf65e7b56015f821c4438))
* **deps:** Update module github.com/cloudquery/cq-gen to v0.0.6 ([#1244](https://github.com/cloudquery/cq-provider-aws/issues/1244)) ([c3aea9d](https://github.com/cloudquery/cq-provider-aws/commit/c3aea9d4406b32a58cddde82f9688bf2508de0cc))
* **deps:** Update module github.com/cloudquery/cq-provider-sdk to v0.14.2 ([#1241](https://github.com/cloudquery/cq-provider-aws/issues/1241)) ([195e048](https://github.com/cloudquery/cq-provider-aws/commit/195e048b1c80f170122809ed957aa9861d318b1a))
* **deps:** Update module github.com/cloudquery/faker/v3 to v3.7.7 ([#1219](https://github.com/cloudquery/cq-provider-aws/issues/1219)) ([e0f76bc](https://github.com/cloudquery/cq-provider-aws/commit/e0f76bc1fe00e54eb5e5ee1da06ead32b3b54202))
* **deps:** Update myrotvorets/set-commit-status-action digest to 85c3f9a ([#1224](https://github.com/cloudquery/cq-provider-aws/issues/1224)) ([1bf2e56](https://github.com/cloudquery/cq-provider-aws/commit/1bf2e56f7e5562e3f5ab674fdcb931126a145942))
* **deps:** Update tubone24/update_release digest to e5b78c8 ([#1225](https://github.com/cloudquery/cq-provider-aws/issues/1225)) ([1de1217](https://github.com/cloudquery/cq-provider-aws/commit/1de12173c94c9a1f367307ace736de0a9498a368))
* Update endpoints ([#1221](https://github.com/cloudquery/cq-provider-aws/issues/1221)) ([7b86dd9](https://github.com/cloudquery/cq-provider-aws/commit/7b86dd9fb1999737d6cb4fb66a238db45a45b60d))
* Update endpoints ([#1222](https://github.com/cloudquery/cq-provider-aws/issues/1222)) ([538e821](https://github.com/cloudquery/cq-provider-aws/commit/538e821752a462172d35dde66fcefe42ac7f3da5))
* Update endpoints ([#1236](https://github.com/cloudquery/cq-provider-aws/issues/1236)) ([2683bd8](https://github.com/cloudquery/cq-provider-aws/commit/2683bd80422d581caa7881cf86348f838980ec6b))
* Update endpoints ([#1253](https://github.com/cloudquery/cq-provider-aws/issues/1253)) ([1a7a8f3](https://github.com/cloudquery/cq-provider-aws/commit/1a7a8f37897ffc95b3c5f9be22931fd207f4a217))


### Miscellaneous Chores

* Update SDK to v0.14.1 ([#1239](https://github.com/cloudquery/cq-provider-aws/issues/1239)) ([2dc2f89](https://github.com/cloudquery/cq-provider-aws/commit/2dc2f890f190b977e365745370e9fb4d52516d7d))

## [0.12.29](https://github.com/cloudquery/cq-provider-aws/compare/v0.12.28...v0.12.29) (2022-07-13)


### Bug Fixes

* Choose correct region for S3 ([#1216](https://github.com/cloudquery/cq-provider-aws/issues/1216)) ([e75f91b](https://github.com/cloudquery/cq-provider-aws/commit/e75f91bd41be4bf6122e2672e7d09539ab115cd0))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/config to v1.15.14 ([#1204](https://github.com/cloudquery/cq-provider-aws/issues/1204)) ([6441700](https://github.com/cloudquery/cq-provider-aws/commit/64417008b0588e72541a50209fbbe964ba1aaa2c))
* Improve Error/Diag handling for nested resources ([#1214](https://github.com/cloudquery/cq-provider-aws/issues/1214)) ([9a55267](https://github.com/cloudquery/cq-provider-aws/commit/9a55267fb951152a5b4d1d6e78b0f86661cf85fb))
* Update endpoints ([#1212](https://github.com/cloudquery/cq-provider-aws/issues/1212)) ([83edf07](https://github.com/cloudquery/cq-provider-aws/commit/83edf07e23ab7e5d196b1723a3b4225e9e2a4624))
* Update endpoints ([#1215](https://github.com/cloudquery/cq-provider-aws/issues/1215)) ([67f595a](https://github.com/cloudquery/cq-provider-aws/commit/67f595a34a0e6fcd7021a3d1de0756e0832f878e))

## [0.12.28](https://github.com/cloudquery/cq-provider-aws/compare/v0.12.27...v0.12.28) (2022-07-11)


### Features

* Generic list and detail ([#1000](https://github.com/cloudquery/cq-provider-aws/issues/1000)) ([16217c8](https://github.com/cloudquery/cq-provider-aws/commit/16217c824ac66196af63ba5b28d55a4e3f3cf4a5))
* Partial lambda Function  Fetch ([#1194](https://github.com/cloudquery/cq-provider-aws/issues/1194)) ([f757824](https://github.com/cloudquery/cq-provider-aws/commit/f75782469a5d805f320c3467b3db489a5713ee40))


### Bug Fixes

* **deps:** Update module github.com/aws/aws-sdk-go-v2 to v1.16.7 ([#1203](https://github.com/cloudquery/cq-provider-aws/issues/1203)) ([35fde37](https://github.com/cloudquery/cq-provider-aws/commit/35fde37ec9007999b56ae24682a681fe7247e6f8))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/credentials to v1.12.8 ([#1205](https://github.com/cloudquery/cq-provider-aws/issues/1205)) ([103b548](https://github.com/cloudquery/cq-provider-aws/commit/103b548618cac9e1480698860f6a712f925d1d71))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/feature/s3/manager to v1.11.19 ([#1206](https://github.com/cloudquery/cq-provider-aws/issues/1206)) ([825e55f](https://github.com/cloudquery/cq-provider-aws/commit/825e55f842c0f0f07740242c1130c6f95ef36ed8))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/acm to v1.14.8 ([#1208](https://github.com/cloudquery/cq-provider-aws/issues/1208)) ([e0a3b4a](https://github.com/cloudquery/cq-provider-aws/commit/e0a3b4aba11637bfa39af86e0d0486d29d47b86c))
* **deps:** Update module github.com/cloudquery/cq-provider-sdk to v0.13.5 ([#1193](https://github.com/cloudquery/cq-provider-aws/issues/1193)) ([0add41c](https://github.com/cloudquery/cq-provider-aws/commit/0add41c0e0d11169ebb12071a056ced017975a62))
* **deps:** Update myrotvorets/set-commit-status-action digest to 987babf ([#1145](https://github.com/cloudquery/cq-provider-aws/issues/1145)) ([bf69648](https://github.com/cloudquery/cq-provider-aws/commit/bf69648c21b51200d0c9d9671b22b9bf8bc2c140))
* **deps:** Update tubone24/update_release digest to 246880c ([#1146](https://github.com/cloudquery/cq-provider-aws/issues/1146)) ([44d9c76](https://github.com/cloudquery/cq-provider-aws/commit/44d9c7656a75016b5308114215e50af4dce6cfeb))
* Handle invalid/malformed token ([#1210](https://github.com/cloudquery/cq-provider-aws/issues/1210)) ([5131326](https://github.com/cloudquery/cq-provider-aws/commit/51313269f5180433d44f1602b7158610bfa3198f))
* Ignore AWS backup tags for not found resource ([#1211](https://github.com/cloudquery/cq-provider-aws/issues/1211)) ([ac11c91](https://github.com/cloudquery/cq-provider-aws/commit/ac11c91698177fe64a1e91840e8bd7299d4d262b))
* Ignore not found errors in apigateway ([#1192](https://github.com/cloudquery/cq-provider-aws/issues/1192)) ([e963155](https://github.com/cloudquery/cq-provider-aws/commit/e963155da7c3b18bfc3bf8483ec317df46cb6abf))
* Ignore SimpleQueueService.NonExistentQueue ([#1173](https://github.com/cloudquery/cq-provider-aws/issues/1173)) ([7452701](https://github.com/cloudquery/cq-provider-aws/commit/745270176d1bbf15e3de3da0598b3983bf48c81e))
* Ignore some not founds in autoscaling groups ([#1181](https://github.com/cloudquery/cq-provider-aws/issues/1181)) ([33c2f11](https://github.com/cloudquery/cq-provider-aws/commit/33c2f1127fda56a9d4599710bd7a52f9e0eda741))
* Ignore TargetGroupNotFound in target_groups ([#1174](https://github.com/cloudquery/cq-provider-aws/issues/1174)) ([1a16deb](https://github.com/cloudquery/cq-provider-aws/commit/1a16deb7aa8a068a8845dacd46398961cde8fc24))
* Panic in cloudtrail ([#1189](https://github.com/cloudquery/cq-provider-aws/issues/1189)) ([0d98a38](https://github.com/cloudquery/cq-provider-aws/commit/0d98a385bc3ad6efe3a8b0579968d8871ecee58e))

## [0.12.27](https://github.com/cloudquery/cq-provider-aws/compare/v0.12.26...v0.12.27) (2022-07-07)


### Features

* Account specific default_region ([#1177](https://github.com/cloudquery/cq-provider-aws/issues/1177)) ([ac0e91b](https://github.com/cloudquery/cq-provider-aws/commit/ac0e91b494a5ec3eb9df705c726b8e12d3a970cc))
* use 'unavailable' for resources without region ([#1183](https://github.com/cloudquery/cq-provider-aws/issues/1183)) ([e7dac0f](https://github.com/cloudquery/cq-provider-aws/commit/e7dac0f42224464583dc1491a296b5e6589dcacc))


### Bug Fixes

* Classify ExpiredTokenException as ACCESS err ([#1171](https://github.com/cloudquery/cq-provider-aws/issues/1171)) ([3e36d75](https://github.com/cloudquery/cq-provider-aws/commit/3e36d759ad6c85a19658606cb43ddae3cd32ca24))
* Ec2 Panic.go ([#1185](https://github.com/cloudquery/cq-provider-aws/issues/1185)) ([dc9db0d](https://github.com/cloudquery/cq-provider-aws/commit/dc9db0d7acfadafd418402752b6862c120367a79))
* Fix Tests ([#1178](https://github.com/cloudquery/cq-provider-aws/issues/1178)) ([8c4d01d](https://github.com/cloudquery/cq-provider-aws/commit/8c4d01d8e8e791b587504d9c953ba6633d099050))
* Ignore rds db snapshot not found ([#1182](https://github.com/cloudquery/cq-provider-aws/issues/1182)) ([0619b4e](https://github.com/cloudquery/cq-provider-aws/commit/0619b4ee41a1afb7041f33d3623c36bdc636ecc7))
* Ignore s3 "bucket not found" error ([#1172](https://github.com/cloudquery/cq-provider-aws/issues/1172)) ([bdb3332](https://github.com/cloudquery/cq-provider-aws/commit/bdb333257c734caec17c64459c193c2a296a8da8))
* Ignore S3 Buckets Errors ([#1175](https://github.com/cloudquery/cq-provider-aws/issues/1175)) ([2907699](https://github.com/cloudquery/cq-provider-aws/commit/2907699a2b193762fbc278535f4afb332b114ad6))

## [0.12.26](https://github.com/cloudquery/cq-provider-aws/compare/v0.12.25...v0.12.26) (2022-07-06)


### Features

* Add support for AWS SES Template ([#1133](https://github.com/cloudquery/cq-provider-aws/issues/1133)) ([59081a8](https://github.com/cloudquery/cq-provider-aws/commit/59081a8f3e834c592dbe788ec5ab0adee3247668))
* Add support for Lightsail Instances ([#1138](https://github.com/cloudquery/cq-provider-aws/issues/1138)) ([bcfb724](https://github.com/cloudquery/cq-provider-aws/commit/bcfb72497535c14a2f98e38d15dfbf094264a2c8))


### Bug Fixes

* **deps:** Update module github.com/cloudquery/cq-provider-sdk to v0.13.4 ([#1164](https://github.com/cloudquery/cq-provider-aws/issues/1164)) ([7684ef4](https://github.com/cloudquery/cq-provider-aws/commit/7684ef425f9be77838005bcc01a7a3761efe3f8b))
* Do not fail if ec2 image is not found ([#1170](https://github.com/cloudquery/cq-provider-aws/issues/1170)) ([a0d6104](https://github.com/cloudquery/cq-provider-aws/commit/a0d6104b3ded675a7b693dbba1dae08e413959e2))
* Ignore few NotFound responses in elbv2 and apigateway ([#1161](https://github.com/cloudquery/cq-provider-aws/issues/1161)) ([24622bb](https://github.com/cloudquery/cq-provider-aws/commit/24622bbe017c98b34bb3fe40ebfc98f949d925bd))
* Update endpoints ([#1169](https://github.com/cloudquery/cq-provider-aws/issues/1169)) ([425b8cf](https://github.com/cloudquery/cq-provider-aws/commit/425b8cf7d0164e82c8c58f4bf51bb03aa214a9e7))

## [0.12.25](https://github.com/cloudquery/cq-provider-aws/compare/v0.12.24...v0.12.25) (2022-07-04)


### Bug Fixes

* **deps:** Update github.com/gocarina/gocsv digest to 72f2e84 ([#1151](https://github.com/cloudquery/cq-provider-aws/issues/1151)) ([1c870e8](https://github.com/cloudquery/cq-provider-aws/commit/1c870e80279c72ca3f768aaaeaf7e9ab84f87cfe))
* **deps:** Update module github.com/aws/aws-sdk-go-v2 to v1.16.6 ([#1152](https://github.com/cloudquery/cq-provider-aws/issues/1152)) ([4ac2583](https://github.com/cloudquery/cq-provider-aws/commit/4ac2583ebe3a074a7e61bec50c22bb059f0668f2))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/accessanalyzer to v1.15.7 ([#1153](https://github.com/cloudquery/cq-provider-aws/issues/1153)) ([9ed2343](https://github.com/cloudquery/cq-provider-aws/commit/9ed234382f41d08c734fd56e54ca735d3bd299ec))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/codepipeline to v1.13.7 ([#1154](https://github.com/cloudquery/cq-provider-aws/issues/1154)) ([0a34f8e](https://github.com/cloudquery/cq-provider-aws/commit/0a34f8e4660077ce0cb5e70a3155d794c3b0309f))
* **deps:** Update module github.com/cloudquery/cq-provider-sdk to v0.13.3 ([#1156](https://github.com/cloudquery/cq-provider-aws/issues/1156)) ([d4b0701](https://github.com/cloudquery/cq-provider-aws/commit/d4b070185414560f183a1de95d25c5f1cf536d13))
* **docs:** Update documentation about adding new resources ([#1136](https://github.com/cloudquery/cq-provider-aws/issues/1136)) ([fe5a5ad](https://github.com/cloudquery/cq-provider-aws/commit/fe5a5ad8dabd9cd533e2899bbb769c9c4f3c1a8c))
* **tests:** Update Lightsail terraform for Integration testing of Instances ([#1137](https://github.com/cloudquery/cq-provider-aws/issues/1137)) ([488a003](https://github.com/cloudquery/cq-provider-aws/commit/488a003a2d74faa991d5343fcfc15dd7760a1cb5))

## [0.12.24](https://github.com/cloudquery/cq-provider-aws/compare/v0.12.23...v0.12.24) (2022-07-03)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/cq-provider-sdk to v0.13.2 ([#1141](https://github.com/cloudquery/cq-provider-aws/issues/1141)) ([110e0d2](https://github.com/cloudquery/cq-provider-aws/commit/110e0d2f635b5f3ba7bdb5e4395f416edd1accbb))

## [0.12.23](https://github.com/cloudquery/cq-provider-aws/compare/v0.12.22...v0.12.23) (2022-07-03)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/cq-provider-sdk to v0.13.1 ([#1127](https://github.com/cloudquery/cq-provider-aws/issues/1127)) ([527a614](https://github.com/cloudquery/cq-provider-aws/commit/527a614d903be767d4a6acc9f50f4801f7af5c71))
* Diagnostic blocking Use of Orgs  ([#1134](https://github.com/cloudquery/cq-provider-aws/issues/1134)) ([ca6f745](https://github.com/cloudquery/cq-provider-aws/commit/ca6f745bd95df401cb7dc93a926f3b171322aba8))
* **docs:** Update instructions for adding new resources and add install-tools command ([#1128](https://github.com/cloudquery/cq-provider-aws/issues/1128)) ([29ac7d3](https://github.com/cloudquery/cq-provider-aws/commit/29ac7d38d8d1ec2425a67ae6e364993a191b1096))
* Typo in example config ([#1132](https://github.com/cloudquery/cq-provider-aws/issues/1132)) ([2cdebb7](https://github.com/cloudquery/cq-provider-aws/commit/2cdebb78d361cc2b7bbbd6408bc1bb6f49612151)), closes [#1131](https://github.com/cloudquery/cq-provider-aws/issues/1131)

## [0.12.22](https://github.com/cloudquery/cq-provider-aws/compare/v0.12.21...v0.12.22) (2022-06-29)


### Bug Fixes

* Add yml test_policy configuration ([#1120](https://github.com/cloudquery/cq-provider-aws/issues/1120)) ([0c19e0a](https://github.com/cloudquery/cq-provider-aws/commit/0c19e0ad177133cde8e7e2727cae88c1774d1c7f))
* Fix typo ([#1122](https://github.com/cloudquery/cq-provider-aws/issues/1122)) ([2b929ab](https://github.com/cloudquery/cq-provider-aws/commit/2b929abb49d0e739ed9b4425ac512561b588ca52))
* Update endpoints ([#1126](https://github.com/cloudquery/cq-provider-aws/issues/1126)) ([4ef62d0](https://github.com/cloudquery/cq-provider-aws/commit/4ef62d050fb7fe6bb5e1a1e28392184a5255c3a5))

## [0.12.21](https://github.com/cloudquery/cq-provider-aws/compare/v0.12.20...v0.12.21) (2022-06-27)


### Features

* Add image deprecation_time ([#1099](https://github.com/cloudquery/cq-provider-aws/issues/1099)) ([55762fd](https://github.com/cloudquery/cq-provider-aws/commit/55762fdadd1b5ce9d471cece40ae0f64b5c9b3fb))


### Bug Fixes

* Docs to Yaml ([#1117](https://github.com/cloudquery/cq-provider-aws/issues/1117)) ([5145461](https://github.com/cloudquery/cq-provider-aws/commit/51454613191428c93c2472a6d12f431e42f2572c))
* Improve Errors AWS Errors ([#1100](https://github.com/cloudquery/cq-provider-aws/issues/1100)) ([1897dbc](https://github.com/cloudquery/cq-provider-aws/commit/1897dbc4b9d4a4a00b10d8319bde3bc441c3acbc))
* Small Fixes  ([#1089](https://github.com/cloudquery/cq-provider-aws/issues/1089)) ([f6fec91](https://github.com/cloudquery/cq-provider-aws/commit/f6fec9110e742a154e94f25d65080e24bb7b5a58))

## [0.12.20](https://github.com/cloudquery/cq-provider-aws/compare/v0.12.19...v0.12.20) (2022-06-27)


### Bug Fixes

* **deps:** fix(deps): Update module github.com/cloudquery/cq-provider-sdk to v0.12.5 ([#1113](https://github.com/cloudquery/cq-provider-aws/issues/1113)) ([f80c663](https://github.com/cloudquery/cq-provider-aws/commit/f80c66316c1e773d21cc2320ac34b61e2931eee3))

## [0.12.19](https://github.com/cloudquery/cq-provider-aws/compare/v0.12.18...v0.12.19) (2022-06-27)


### Bug Fixes

* **deps:** fix(deps): Update module github.com/cloudquery/cq-provider-sdk to v0.12.4 ([#1111](https://github.com/cloudquery/cq-provider-aws/issues/1111)) ([0dba643](https://github.com/cloudquery/cq-provider-aws/commit/0dba643fa8776940644fac936dacf8923b42393e))

## [0.12.18](https://github.com/cloudquery/cq-provider-aws/compare/v0.12.17...v0.12.18) (2022-06-26)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/cq-provider-sdk to v0.12.3 ([#1106](https://github.com/cloudquery/cq-provider-aws/issues/1106)) ([80c2ec4](https://github.com/cloudquery/cq-provider-aws/commit/80c2ec4f5de162d6469b9778e5218ce954839586))

## [0.12.17](https://github.com/cloudquery/cq-provider-aws/compare/v0.12.16...v0.12.17) (2022-06-26)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/cq-provider-sdk to v0.12.2 ([#1101](https://github.com/cloudquery/cq-provider-aws/issues/1101)) ([5a694bb](https://github.com/cloudquery/cq-provider-aws/commit/5a694bbd8e161f3dd69490db47593fe8748a7ca0))
* Update endpoints ([#1104](https://github.com/cloudquery/cq-provider-aws/issues/1104)) ([0779d72](https://github.com/cloudquery/cq-provider-aws/commit/0779d7286cf94db6cce08681c125fc04a2185b15))
* YAML Include accounts block ([#1103](https://github.com/cloudquery/cq-provider-aws/issues/1103)) ([6829f71](https://github.com/cloudquery/cq-provider-aws/commit/6829f7132c92c2ff39ae049c6f77daa93db48255))

## [0.12.16](https://github.com/cloudquery/cq-provider-aws/compare/v0.12.15...v0.12.16) (2022-06-22)


### Features

* YAML config support ([#1067](https://github.com/cloudquery/cq-provider-aws/issues/1067)) ([24d2722](https://github.com/cloudquery/cq-provider-aws/commit/24d27225fc6cf8494865ca06eb46be37a2b0c23f))


### Bug Fixes

* **deps:** Upgrade AWS ECR SDK ([#1079](https://github.com/cloudquery/cq-provider-aws/issues/1079)) ([7de5bda](https://github.com/cloudquery/cq-provider-aws/commit/7de5bdacfe58433883448d24bff9f31ba2e17dc7))
* Classify to many open files ([#1064](https://github.com/cloudquery/cq-provider-aws/issues/1064)) ([89b1684](https://github.com/cloudquery/cq-provider-aws/commit/89b1684652e61aeb58fd3d2dec7faf9e2aea6a12))
* **deps:** Update module github.com/cloudquery/cq-provider-sdk to v0.12.1 ([#1083](https://github.com/cloudquery/cq-provider-aws/issues/1083)) ([6a4dc83](https://github.com/cloudquery/cq-provider-aws/commit/6a4dc8348463e8068ca699808edb88818b5ee4fa))
* List buckets only us-east-1 ([#1088](https://github.com/cloudquery/cq-provider-aws/issues/1088)) ([b3d7476](https://github.com/cloudquery/cq-provider-aws/commit/b3d74768d4d8cec5e4fcdb53794cb788d2815270))
* Rename S3 Resources ([#1082](https://github.com/cloudquery/cq-provider-aws/issues/1082)) ([757b9aa](https://github.com/cloudquery/cq-provider-aws/commit/757b9aa900f7715b6f9741f74c62fc8fe4b4f539))

## [0.12.16-rc2](https://github.com/cloudquery/cq-provider-aws/compare/v0.12.16-rc1...v0.12.16-rc2) (2022-06-22)


### Bug Fixes

* Classify to many open files ([#1064](https://github.com/cloudquery/cq-provider-aws/issues/1064)) ([89b1684](https://github.com/cloudquery/cq-provider-aws/commit/89b1684652e61aeb58fd3d2dec7faf9e2aea6a12))
* **deps:** Update module github.com/cloudquery/cq-provider-sdk to v0.12.1 ([#1083](https://github.com/cloudquery/cq-provider-aws/issues/1083)) ([6a4dc83](https://github.com/cloudquery/cq-provider-aws/commit/6a4dc8348463e8068ca699808edb88818b5ee4fa))
* List buckets only us-east-1 ([#1088](https://github.com/cloudquery/cq-provider-aws/issues/1088)) ([b3d7476](https://github.com/cloudquery/cq-provider-aws/commit/b3d74768d4d8cec5e4fcdb53794cb788d2815270))
* Rename S3 Resources ([#1082](https://github.com/cloudquery/cq-provider-aws/issues/1082)) ([757b9aa](https://github.com/cloudquery/cq-provider-aws/commit/757b9aa900f7715b6f9741f74c62fc8fe4b4f539))


### Miscellaneous Chores

* Release 0.12.16-rc2 ([#1090](https://github.com/cloudquery/cq-provider-aws/issues/1090)) ([1dca452](https://github.com/cloudquery/cq-provider-aws/commit/1dca4522023c52f55b8dd4c8165a9d4c27b8e1b8))
* Release v0.12.16 ([#1086](https://github.com/cloudquery/cq-provider-aws/issues/1086)) ([885ff24](https://github.com/cloudquery/cq-provider-aws/commit/885ff24a94a31920e2dda925ef495b9ab30cd360))

## [0.12.16-rc1](https://github.com/cloudquery/cq-provider-aws/compare/v0.12.15...v0.12.16-rc1) (2022-06-21)


### Features

* YAML config support ([#1067](https://github.com/cloudquery/cq-provider-aws/issues/1067)) ([24d2722](https://github.com/cloudquery/cq-provider-aws/commit/24d27225fc6cf8494865ca06eb46be37a2b0c23f))


### Bug Fixes

* **deps:** Upgrade AWS ECR SDK ([#1079](https://github.com/cloudquery/cq-provider-aws/issues/1079)) ([7de5bda](https://github.com/cloudquery/cq-provider-aws/commit/7de5bdacfe58433883448d24bff9f31ba2e17dc7))


### Miscellaneous Chores

* Release v0.12.16-rc1 ([#1084](https://github.com/cloudquery/cq-provider-aws/issues/1084)) ([745eb88](https://github.com/cloudquery/cq-provider-aws/commit/745eb884941be331b295b4961fc84ddfc68022f5))

## [0.12.15](https://github.com/cloudquery/cq-provider-aws/compare/v0.12.14...v0.12.15) (2022-06-20)


### Bug Fixes

* Change PK in aws_ec2_vpc_endpoint_services from (account_id, id) ([#1077](https://github.com/cloudquery/cq-provider-aws/issues/1077)) ([fbbfdc5](https://github.com/cloudquery/cq-provider-aws/commit/fbbfdc5ea8e07c2c407f547f18f7de3a101ac958))
* **deps:** Update module github.com/cloudquery/cq-provider-sdk to v0.11.4 ([#1073](https://github.com/cloudquery/cq-provider-aws/issues/1073)) ([d58c24b](https://github.com/cloudquery/cq-provider-aws/commit/d58c24b5da42198e1d2520dcc8abceac635801bd))
* Panic in ecs task definitions ([#1076](https://github.com/cloudquery/cq-provider-aws/issues/1076)) ([ecdd07e](https://github.com/cloudquery/cq-provider-aws/commit/ecdd07ebaeec4659c8afb82036d6a8047632efdc))
* Redact IPv6 addresses in "dial tcp" errors ([#1075](https://github.com/cloudquery/cq-provider-aws/issues/1075)) ([fcd04d5](https://github.com/cloudquery/cq-provider-aws/commit/fcd04d58b862c4c1a1c4bdb0299b9f90a5a01ffb))

## [0.12.14](https://github.com/cloudquery/cq-provider-aws/compare/v0.12.13...v0.12.14) (2022-06-20)


### Bug Fixes

* Classify credential errors as USER type ([#1056](https://github.com/cloudquery/cq-provider-aws/issues/1056)) ([e04e493](https://github.com/cloudquery/cq-provider-aws/commit/e04e4939d4d7398fa30212dbef384b8b8165de94))
* Redact separate error messages separately ([#1071](https://github.com/cloudquery/cq-provider-aws/issues/1071)) ([e46371b](https://github.com/cloudquery/cq-provider-aws/commit/e46371be7b6522d42465db994cf97f49475da2f0))

## [0.12.13](https://github.com/cloudquery/cq-provider-aws/compare/v0.12.12...v0.12.13) (2022-06-19)


### Bug Fixes

* Update endpoints ([#1065](https://github.com/cloudquery/cq-provider-aws/issues/1065)) ([b95d6b5](https://github.com/cloudquery/cq-provider-aws/commit/b95d6b54465639a7bd087d73680b707a1cf17ab6))
* Update endpoints ([#1068](https://github.com/cloudquery/cq-provider-aws/issues/1068)) ([a8cf7e8](https://github.com/cloudquery/cq-provider-aws/commit/a8cf7e89fe144b03ac359e31e3b770fa35bc41ae))

## [0.12.12](https://github.com/cloudquery/cq-provider-aws/compare/v0.12.11...v0.12.12) (2022-06-15)


### Features

* Add VPC Endpoint Services and Configurations ([#1029](https://github.com/cloudquery/cq-provider-aws/issues/1029)) ([668ea91](https://github.com/cloudquery/cq-provider-aws/commit/668ea91620f85994606935d3b7a8f171a7d8e8a7))


### Bug Fixes

* **deps:** Update module github.com/cloudquery/cq-provider-sdk to v0.11.2 ([#1062](https://github.com/cloudquery/cq-provider-aws/issues/1062)) ([5b2bc76](https://github.com/cloudquery/cq-provider-aws/commit/5b2bc764d31377483c251e90f59ef526d926b556))
* **deps:** Update module github.com/cloudquery/cq-provider-sdk to v0.11.3 ([#1063](https://github.com/cloudquery/cq-provider-aws/issues/1063)) ([b81b84c](https://github.com/cloudquery/cq-provider-aws/commit/b81b84c77416acfa2a380d83027c0a077ae99c73))
* Resolvers Returning Early ([#1059](https://github.com/cloudquery/cq-provider-aws/issues/1059)) ([449aefc](https://github.com/cloudquery/cq-provider-aws/commit/449aefcb9d660689b5fe22c51cecaee520b584ac))

## [0.12.11](https://github.com/cloudquery/cq-provider-aws/compare/v0.12.10...v0.12.11) (2022-06-14)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/cq-provider-sdk to v0.11.1 ([#1053](https://github.com/cloudquery/cq-provider-aws/issues/1053)) ([a48cf77](https://github.com/cloudquery/cq-provider-aws/commit/a48cf77726731a823a4407a3e752d1aff857cca2))
* Explicitly Ignore EC2 Classic EIPs ([#1055](https://github.com/cloudquery/cq-provider-aws/issues/1055)) ([28feadf](https://github.com/cloudquery/cq-provider-aws/commit/28feadf5a664478075cb004230bba7f331805e04))
* Possible use of a wrong region in s3 buckets. ([#1052](https://github.com/cloudquery/cq-provider-aws/issues/1052)) ([e54b46a](https://github.com/cloudquery/cq-provider-aws/commit/e54b46ac8c42b440504006d8236787ace724320a))

## [0.12.10](https://github.com/cloudquery/cq-provider-aws/compare/v0.12.9...v0.12.10) (2022-06-13)


### Bug Fixes

* Added `not found` errors handlers in athena resources ([#1021](https://github.com/cloudquery/cq-provider-aws/issues/1021)) ([f18abef](https://github.com/cloudquery/cq-provider-aws/commit/f18abef774d73b9ba7a6f910f1c3f40b5f3029e1))
* Added `not found` errors handlers in athena resources ([#1021](https://github.com/cloudquery/cq-provider-aws/issues/1021)) ([297fa6c](https://github.com/cloudquery/cq-provider-aws/commit/297fa6ca8db0f327193e9c69862adb55cae6824b))
* change from IgnoreAccessDeniedServiceDisabled to IgnoreCommonErrors ([#1033](https://github.com/cloudquery/cq-provider-aws/issues/1033)) ([1b98229](https://github.com/cloudquery/cq-provider-aws/commit/1b982292b2f3359e8364b2fb6937a0992472e1d0))
* Redshift Panic when accessing nested attributes ([8bae50b](https://github.com/cloudquery/cq-provider-aws/commit/8bae50b4f9dab65e8753f00594df0998a638aa02))
* Update endpoints ([#1032](https://github.com/cloudquery/cq-provider-aws/issues/1032)) ([bd6159a](https://github.com/cloudquery/cq-provider-aws/commit/bd6159a4a34f795c402fad46d7aa5cbb663dec40))

## [0.12.9](https://github.com/cloudquery/cq-provider-aws/compare/v0.12.8...v0.12.9) (2022-06-10)


### Bug Fixes

* Added FailedResourceAccessException to error classifier ([#1028](https://github.com/cloudquery/cq-provider-aws/issues/1028)) ([f43fe15](https://github.com/cloudquery/cq-provider-aws/commit/f43fe15f502936b99a331842ba5e07824a904e62))
* Continue fetching on incorrect account permissions ([#1030](https://github.com/cloudquery/cq-provider-aws/issues/1030)) ([71008d2](https://github.com/cloudquery/cq-provider-aws/commit/71008d2854ab8d5f8248a38a916eaf1554d4391e))
* Improve IAM Report Error Handling ([#1009](https://github.com/cloudquery/cq-provider-aws/issues/1009)) ([1c77a63](https://github.com/cloudquery/cq-provider-aws/commit/1c77a63d868376b050476276c799e3f8b88f6de5))

## [0.12.8](https://github.com/cloudquery/cq-provider-aws/compare/v0.12.7...v0.12.8) (2022-06-08)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/cq-provider-sdk to v0.11.0 ([#1023](https://github.com/cloudquery/cq-provider-aws/issues/1023)) ([ae18dbb](https://github.com/cloudquery/cq-provider-aws/commit/ae18dbbdc3511d6533ed01ae37643fb79415a22a))
* Ignore listener certs for gateway listeners ([#1005](https://github.com/cloudquery/cq-provider-aws/issues/1005)) ([6553c8c](https://github.com/cloudquery/cq-provider-aws/commit/6553c8c7febdaa99af5fff9662aafab013561f62))
* Lambda function Tags null ([#1016](https://github.com/cloudquery/cq-provider-aws/issues/1016)) ([35721ba](https://github.com/cloudquery/cq-provider-aws/commit/35721ba79fddf132b4d78cfd07dabbc395e5f2e7))
* Panic in fetchRdsInstanceDbSubnetGroupSubnets ([#1020](https://github.com/cloudquery/cq-provider-aws/issues/1020)) ([0aa25cf](https://github.com/cloudquery/cq-provider-aws/commit/0aa25cfb47eb3271387dcd16fd88a18d3aa9b5f4))
* Panic in IsInvalidParameterValueError ([#1019](https://github.com/cloudquery/cq-provider-aws/issues/1019)) ([32df59e](https://github.com/cloudquery/cq-provider-aws/commit/32df59e1e15a6ffcc9c173f86c2b9fbea6d4237a))
* Update endpoints ([#1025](https://github.com/cloudquery/cq-provider-aws/issues/1025)) ([bf3e6ec](https://github.com/cloudquery/cq-provider-aws/commit/bf3e6ec3f660f3a180bd7a05e73c1f53cffa8c41))

## [0.12.7](https://github.com/cloudquery/cq-provider-aws/compare/v0.12.6...v0.12.7) (2022-06-07)


### Bug Fixes

* AWS configuration error level owerwrite removed ([#999](https://github.com/cloudquery/cq-provider-aws/issues/999)) ([32d60e6](https://github.com/cloudquery/cq-provider-aws/commit/32d60e69a030c6b90ac9c5b4caa6a329c952ac7d))
* **deps:** Update module github.com/cloudquery/cq-provider-sdk to v0.10.10 ([#1014](https://github.com/cloudquery/cq-provider-aws/issues/1014)) ([2398536](https://github.com/cloudquery/cq-provider-aws/commit/23985362a12a2050457243b4a90413bee2325f38))
* **deps:** Update module github.com/cloudquery/cq-provider-sdk to v0.10.11 ([#1015](https://github.com/cloudquery/cq-provider-aws/issues/1015)) ([0c3cf3f](https://github.com/cloudquery/cq-provider-aws/commit/0c3cf3f501293e8632ad8a5d0e82a2a0b6856c83))
* **deps:** Update module github.com/cloudquery/cq-provider-sdk to v0.10.9 ([#1012](https://github.com/cloudquery/cq-provider-aws/issues/1012)) ([f566007](https://github.com/cloudquery/cq-provider-aws/commit/f566007556de1d4cc2b048d7fc38c603dedb048e))
* Not Supported region added to error classifier ([#1006](https://github.com/cloudquery/cq-provider-aws/issues/1006)) ([be79739](https://github.com/cloudquery/cq-provider-aws/commit/be7973907251f089dcc988ae0d094a3e2416a5fe))
* Update endpoints ([#1017](https://github.com/cloudquery/cq-provider-aws/issues/1017)) ([729f230](https://github.com/cloudquery/cq-provider-aws/commit/729f23097a0c40659dff2aca154d82dcec409136))

## [0.12.6](https://github.com/cloudquery/cq-provider-aws/compare/v0.12.5...v0.12.6) (2022-06-07)


### Features

* Added Target Health Descriptions to Target Groups ([#996](https://github.com/cloudquery/cq-provider-aws/issues/996)) ([d1ffc37](https://github.com/cloudquery/cq-provider-aws/commit/d1ffc372bd5715c5756b206b3d1a885e0dd5c636))


### Bug Fixes

* **deps:** Update module github.com/cloudquery/cq-provider-sdk to v0.10.8 ([#1011](https://github.com/cloudquery/cq-provider-aws/issues/1011)) ([0d8f687](https://github.com/cloudquery/cq-provider-aws/commit/0d8f687579c78ba1a11b3bc8ceb415ab6aa9d872))
* Update endpoints ([#1008](https://github.com/cloudquery/cq-provider-aws/issues/1008)) ([6bfa91a](https://github.com/cloudquery/cq-provider-aws/commit/6bfa91ae63f5707ee17e135244731157882703ae))

## [0.12.5](https://github.com/cloudquery/cq-provider-aws/compare/v0.12.4...v0.12.5) (2022-06-06)


### Bug Fixes

* Wrap provider errors ([#989](https://github.com/cloudquery/cq-provider-aws/issues/989)) ([53d391b](https://github.com/cloudquery/cq-provider-aws/commit/53d391b821f2a1340ad04aa50a9a8fbe1744ac11))

### [0.12.4](https://github.com/cloudquery/cq-provider-aws/compare/v0.12.3...v0.12.4) (2022-06-01)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/cq-gen to v0.0.5 ([#984](https://github.com/cloudquery/cq-provider-aws/issues/984)) ([9c7fd19](https://github.com/cloudquery/cq-provider-aws/commit/9c7fd19dd741d5b5e4334c1a8ad32a6fff2a51a6))
* **deps:** Update module github.com/cloudquery/cq-provider-sdk to v0.10.7 ([#985](https://github.com/cloudquery/cq-provider-aws/issues/985)) ([6fe6cc0](https://github.com/cloudquery/cq-provider-aws/commit/6fe6cc087489b7eb29dc78e66549148fb46865c1))
* Fixed Auth config diagnostics level ([#987](https://github.com/cloudquery/cq-provider-aws/issues/987)) ([9e8be90](https://github.com/cloudquery/cq-provider-aws/commit/9e8be90fefc224a9f05f205e4e3e3850fabd130f))
* Ignore NotFound on all aws resoruces ([#982](https://github.com/cloudquery/cq-provider-aws/issues/982)) ([a68e885](https://github.com/cloudquery/cq-provider-aws/commit/a68e8859735c61d13c961d38acc26135ff354b0d))

### [0.12.3](https://github.com/cloudquery/cq-provider-aws/compare/v0.12.2...v0.12.3) (2022-06-01)


### Bug Fixes

* Bucket missing region ([#978](https://github.com/cloudquery/cq-provider-aws/issues/978)) ([1467b6e](https://github.com/cloudquery/cq-provider-aws/commit/1467b6e590c646d222da2c0a2bb0c7962f6d7f12))
* Handle panic on Waf logging configuration not found ([#970](https://github.com/cloudquery/cq-provider-aws/issues/970)) ([8ffb3e6](https://github.com/cloudquery/cq-provider-aws/commit/8ffb3e6803090f03c5e736b3f6b4fce941bcd457))

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
