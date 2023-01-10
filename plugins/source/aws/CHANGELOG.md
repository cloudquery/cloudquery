# Changelog

All notable changes to this provider will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).


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
