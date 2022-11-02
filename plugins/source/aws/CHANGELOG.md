# Changelog

All notable changes to this provider will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).


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
