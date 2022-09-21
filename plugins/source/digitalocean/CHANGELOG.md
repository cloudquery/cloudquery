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

## [0.7.0-pre.0](https://github.com/cloudquery/cloudquery/compare/plugins/source/digitalocean-v0.6.5-pre.0...plugins/source/digitalocean/v0.7.0-pre.0) (2022-09-21)


### ⚠ BREAKING CHANGES

* Migrate DigitalOcean plugin to v2 (#1794)

### Features

* Add Sentry DSN ([#1913](https://github.com/cloudquery/cloudquery/issues/1913)) ([5cc036e](https://github.com/cloudquery/cloudquery/commit/5cc036e956cb9dc92832783e15088c8249fe2941))
* Added throttling for digitalocean API calls ([#1546](https://github.com/cloudquery/cloudquery/issues/1546)) ([bb40b59](https://github.com/cloudquery/cloudquery/commit/bb40b5951978918f2b9332063ff251652df55754))
* Migrate DigitalOcean plugin to v2 ([#1794](https://github.com/cloudquery/cloudquery/issues/1794)) ([e556185](https://github.com/cloudquery/cloudquery/commit/e5561853c092adb0ed73139aee18ca3b3671b27d))


### Bug Fixes

* **deps:** Update Terraform tls to v4.0.2 ([#1653](https://github.com/cloudquery/cloudquery/issues/1653)) ([8f3bbeb](https://github.com/cloudquery/cloudquery/commit/8f3bbeba64723c6744ae0f9db747261668f6a087))

## [0.6.5](https://github.com/cloudquery/cloudquery/compare/plugins/source/digitalocean/v0.6.4...plugins/source/digitalocean/v0.6.5) (2022-09-01)


### Bug Fixes

* **deps:** Update Terraform tls to v4.0.2 ([#1653](https://github.com/cloudquery/cloudquery/issues/1653)) ([8f3bbeb](https://github.com/cloudquery/cloudquery/commit/8f3bbeba64723c6744ae0f9db747261668f6a087))


## [0.6.4](https://github.com/cloudquery/cloudquery/compare/plugins/source/digitalocean-v0.6.3...plugins/source/digitalocean/v0.6.4) (2022-08-25)


### Features

* Added throttling for digitalocean API calls ([#1546](https://github.com/cloudquery/cloudquery/issues/1546)) ([bb40b59](https://github.com/cloudquery/cloudquery/commit/bb40b5951978918f2b9332063ff251652df55754))

## [0.6.3](https://github.com/cloudquery/cq-provider-digitalocean/compare/v0.6.2...v0.6.3) (2022-08-07)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/cq-provider-sdk to v0.14.7 ([#221](https://github.com/cloudquery/cq-provider-digitalocean/issues/221)) ([9fa0052](https://github.com/cloudquery/cq-provider-digitalocean/commit/9fa0052d4242b94523fbd2d7fd919e4e5a2c63d2))

## [0.6.2](https://github.com/cloudquery/cq-provider-digitalocean/compare/v0.6.1...v0.6.2) (2022-08-07)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/cq-provider-sdk to v0.14.6 ([#214](https://github.com/cloudquery/cq-provider-digitalocean/issues/214)) ([9ec2448](https://github.com/cloudquery/cq-provider-digitalocean/commit/9ec244814dc73578a5636b4b36d17de00bd105c9))
* **deps:** Update module github.com/hashicorp/go-hclog to v1.2.2 ([#216](https://github.com/cloudquery/cq-provider-digitalocean/issues/216)) ([79da06b](https://github.com/cloudquery/cq-provider-digitalocean/commit/79da06b780ddd0f92fb7a3aa77c83fdd8a22a91f))

## [0.6.1](https://github.com/cloudquery/cq-provider-digitalocean/compare/v0.6.0...v0.6.1) (2022-07-27)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/cq-provider-sdk to v0.14.5 ([#210](https://github.com/cloudquery/cq-provider-digitalocean/issues/210)) ([aed751b](https://github.com/cloudquery/cq-provider-digitalocean/commit/aed751b10d4be793cdd83bda19985ee125701e15))

## [0.6.0](https://github.com/cloudquery/cq-provider-digitalocean/compare/v0.5.24...v0.6.0) (2022-07-27)


### ⚠ BREAKING CHANGES

* Update SDK to v0.14.1 (#203)

### Bug Fixes

* **deps:** Update module github.com/cloudquery/cq-provider-sdk to v0.14.2 ([#201](https://github.com/cloudquery/cq-provider-digitalocean/issues/201)) ([cfdb788](https://github.com/cloudquery/cq-provider-digitalocean/commit/cfdb78818237bb0ba727299cb91e67423878354d))
* **deps:** Update module github.com/cloudquery/cq-provider-sdk to v0.14.3 ([#204](https://github.com/cloudquery/cq-provider-digitalocean/issues/204)) ([d2ce98b](https://github.com/cloudquery/cq-provider-digitalocean/commit/d2ce98bfab8263ecae15cc0a569c758dd25597b3))
* **deps:** Update module github.com/cloudquery/cq-provider-sdk to v0.14.4 ([#205](https://github.com/cloudquery/cq-provider-digitalocean/issues/205)) ([535e160](https://github.com/cloudquery/cq-provider-digitalocean/commit/535e160c7cf5efc7f6db1a5b8d5fefd225440c42))
* **deps:** Update Terraform tls to v4 ([#208](https://github.com/cloudquery/cq-provider-digitalocean/issues/208)) ([5d2d735](https://github.com/cloudquery/cq-provider-digitalocean/commit/5d2d735472525beb7fcde2258c50188d51f23afc))
* **deps:** Update tubone24/update_release digest to 87bc28c ([#182](https://github.com/cloudquery/cq-provider-digitalocean/issues/182)) ([f5b0d8e](https://github.com/cloudquery/cq-provider-digitalocean/commit/f5b0d8e1e828b4618b18750708726842e9c0cfb2))


### Miscellaneous Chores

* Update SDK to v0.14.1 ([#203](https://github.com/cloudquery/cq-provider-digitalocean/issues/203)) ([cfcdc3c](https://github.com/cloudquery/cq-provider-digitalocean/commit/cfcdc3ce7d7d32339b3c5bca6465e7564f7d7fc4))

## [0.5.24](https://github.com/cloudquery/cq-provider-digitalocean/compare/v0.5.23...v0.5.24) (2022-07-11)


### Bug Fixes

* **deps:** Update module github.com/aws/aws-sdk-go-v2 to v1.16.7 ([#198](https://github.com/cloudquery/cq-provider-digitalocean/issues/198)) ([6c55cf9](https://github.com/cloudquery/cq-provider-digitalocean/commit/6c55cf97c5745ed19c819618775da5d7cadccf35))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/s3 to v1.27.1 ([#199](https://github.com/cloudquery/cq-provider-digitalocean/issues/199)) ([394e688](https://github.com/cloudquery/cq-provider-digitalocean/commit/394e6883c87f7685a5c5b03b2652ad3a777de2c3))
* **deps:** Update module github.com/cloudquery/cq-provider-sdk to v0.13.5 ([#196](https://github.com/cloudquery/cq-provider-digitalocean/issues/196)) ([c524dd4](https://github.com/cloudquery/cq-provider-digitalocean/commit/c524dd434b350a539d7392d3f71f5beec86fd1ae))

## [0.5.23](https://github.com/cloudquery/cq-provider-digitalocean/compare/v0.5.22...v0.5.23) (2022-07-05)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/cq-provider-sdk to v0.13.4 ([#193](https://github.com/cloudquery/cq-provider-digitalocean/issues/193)) ([6b1bac5](https://github.com/cloudquery/cq-provider-digitalocean/commit/6b1bac54588e17b3e0d7bf47b96ec0d2a9c05aa3))

## [0.5.22](https://github.com/cloudquery/cq-provider-digitalocean/compare/v0.5.21...v0.5.22) (2022-07-04)


### Bug Fixes

* **deps:** Update module github.com/aws/aws-sdk-go-v2 to v1.16.6 ([#186](https://github.com/cloudquery/cq-provider-digitalocean/issues/186)) ([bdb7777](https://github.com/cloudquery/cq-provider-digitalocean/commit/bdb777740f2d6219e29d637e53cdc1d9f3532670))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/s3 to v1.27.0 ([#188](https://github.com/cloudquery/cq-provider-digitalocean/issues/188)) ([fa68cba](https://github.com/cloudquery/cq-provider-digitalocean/commit/fa68cba288c41d7f5ece1758b557334f5ac10484))
* **deps:** Update module github.com/digitalocean/godo to v1.81.0 ([#185](https://github.com/cloudquery/cq-provider-digitalocean/issues/185)) ([e5efa66](https://github.com/cloudquery/cq-provider-digitalocean/commit/e5efa6680cf6e49261f67d041f555572edd43955))

## [0.5.21](https://github.com/cloudquery/cq-provider-digitalocean/compare/v0.5.20...v0.5.21) (2022-07-04)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/cq-provider-sdk to v0.13.3 ([#190](https://github.com/cloudquery/cq-provider-digitalocean/issues/190)) ([23412d0](https://github.com/cloudquery/cq-provider-digitalocean/commit/23412d0f3596956375ce5336c39b93511f442b2a))
* **deps:** Update Terraform tls to v3.4.0 ([#183](https://github.com/cloudquery/cq-provider-digitalocean/issues/183)) ([3e189fe](https://github.com/cloudquery/cq-provider-digitalocean/commit/3e189fee7d92e4aeb7545805e25ab815ed852181))
* Docs to Yaml ([#172](https://github.com/cloudquery/cq-provider-digitalocean/issues/172)) ([4cf00e6](https://github.com/cloudquery/cq-provider-digitalocean/commit/4cf00e6b0ea3f6162cd3145d5f1a6b5f1fefeb94))

## [0.5.20](https://github.com/cloudquery/cq-provider-digitalocean/compare/v0.5.19...v0.5.20) (2022-07-03)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/cq-provider-sdk to v0.13.2 ([#179](https://github.com/cloudquery/cq-provider-digitalocean/issues/179)) ([f3ead38](https://github.com/cloudquery/cq-provider-digitalocean/commit/f3ead3819d35ae8f269be10e1bfac059df309aa1))

## [0.5.19](https://github.com/cloudquery/cq-provider-digitalocean/compare/v0.5.18...v0.5.19) (2022-07-03)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/cq-provider-sdk to v0.13.1 ([#175](https://github.com/cloudquery/cq-provider-digitalocean/issues/175)) ([f970c4b](https://github.com/cloudquery/cq-provider-digitalocean/commit/f970c4bf9603c7684f367226124636f563a96fb3))

## [0.5.18](https://github.com/cloudquery/cq-provider-digitalocean/compare/v0.5.17...v0.5.18) (2022-06-30)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/cq-provider-sdk to v0.13.0 ([#173](https://github.com/cloudquery/cq-provider-digitalocean/issues/173)) ([7088abf](https://github.com/cloudquery/cq-provider-digitalocean/commit/7088abf414cebd07cfa04473c763702cd31697e5))

## [0.5.17](https://github.com/cloudquery/cq-provider-digitalocean/compare/v0.5.16...v0.5.17) (2022-06-27)


### Features

* Classify error 420 "Too many requests" ([#154](https://github.com/cloudquery/cq-provider-digitalocean/issues/154)) ([780cd66](https://github.com/cloudquery/cq-provider-digitalocean/commit/780cd667a4b657b210bd49b381df8a517fc242e1))

## [0.5.16](https://github.com/cloudquery/cq-provider-digitalocean/compare/v0.5.15...v0.5.16) (2022-06-27)


### Bug Fixes

* **deps:** fix(deps): Update module github.com/cloudquery/cq-provider-sdk to v0.12.5 ([#167](https://github.com/cloudquery/cq-provider-digitalocean/issues/167)) ([5298f7d](https://github.com/cloudquery/cq-provider-digitalocean/commit/5298f7d4c302f181668f319a61c42710b1ab840c))

## [0.5.15](https://github.com/cloudquery/cq-provider-digitalocean/compare/v0.5.14...v0.5.15) (2022-06-27)


### Bug Fixes

* **deps:** fix(deps): Update module github.com/cloudquery/cq-provider-sdk to v0.12.4 ([#166](https://github.com/cloudquery/cq-provider-digitalocean/issues/166)) ([bc9016a](https://github.com/cloudquery/cq-provider-digitalocean/commit/bc9016a61356695856ee31223d7dce652a2e46bd))
* Support floating ip without droplet ([#156](https://github.com/cloudquery/cq-provider-digitalocean/issues/156)) ([c703305](https://github.com/cloudquery/cq-provider-digitalocean/commit/c703305e0030fae986bd7dfaa07c104fd8b62f20))

## [0.5.14](https://github.com/cloudquery/cq-provider-digitalocean/compare/v0.5.13...v0.5.14) (2022-06-26)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/cq-provider-sdk to v0.12.3 ([#161](https://github.com/cloudquery/cq-provider-digitalocean/issues/161)) ([8f075b5](https://github.com/cloudquery/cq-provider-digitalocean/commit/8f075b5039de676810f880a6334acf1aa1762ef6))

## [0.5.13](https://github.com/cloudquery/cq-provider-digitalocean/compare/v0.5.12...v0.5.13) (2022-06-26)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/cq-provider-sdk to v0.12.2 ([#153](https://github.com/cloudquery/cq-provider-digitalocean/issues/153)) ([e9080a1](https://github.com/cloudquery/cq-provider-digitalocean/commit/e9080a1eb0bc4acd90b2004758494178ae175de6))

## [0.5.12](https://github.com/cloudquery/cq-provider-digitalocean/compare/v0.5.11...v0.5.12) (2022-06-22)


### Features

* YAML config support ([#155](https://github.com/cloudquery/cq-provider-digitalocean/issues/155)) ([0507d07](https://github.com/cloudquery/cq-provider-digitalocean/commit/0507d07422ba9e4edcbd9476c350e4b4d0012f36))

## [0.5.11](https://github.com/cloudquery/cq-provider-digitalocean/compare/v0.5.10...v0.5.11) (2022-06-20)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/cq-provider-sdk to v0.11.4 ([#150](https://github.com/cloudquery/cq-provider-digitalocean/issues/150)) ([5398207](https://github.com/cloudquery/cq-provider-digitalocean/commit/5398207fd57e8c5b1599f8ea07111750ead55727))

## [0.5.10](https://github.com/cloudquery/cq-provider-digitalocean/compare/v0.5.9...v0.5.10) (2022-06-16)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/cq-provider-sdk to v0.11.2 ([#145](https://github.com/cloudquery/cq-provider-digitalocean/issues/145)) ([0bb8f91](https://github.com/cloudquery/cq-provider-digitalocean/commit/0bb8f91b8715cd85906efba1b5eb162881e5915e))
* **deps:** Update module github.com/cloudquery/cq-provider-sdk to v0.11.3 ([#147](https://github.com/cloudquery/cq-provider-digitalocean/issues/147)) ([c3d4c48](https://github.com/cloudquery/cq-provider-digitalocean/commit/c3d4c4823e515fa116d85b809124e6d6eef72d3b))
* Registry not found on Registry.Get ([#144](https://github.com/cloudquery/cq-provider-digitalocean/issues/144)) ([74c455a](https://github.com/cloudquery/cq-provider-digitalocean/commit/74c455ac61fa8d77d123f6b86bd8293ca0bc0729))

## [0.5.9](https://github.com/cloudquery/cq-provider-digitalocean/compare/v0.5.8...v0.5.9) (2022-06-14)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/cq-provider-sdk to v0.11.1 ([#142](https://github.com/cloudquery/cq-provider-digitalocean/issues/142)) ([d823327](https://github.com/cloudquery/cq-provider-digitalocean/commit/d82332777d05843d76e362934d5b0a52d9c6f40e))

## [0.5.8](https://github.com/cloudquery/cq-provider-digitalocean/compare/v0.5.7...v0.5.8) (2022-06-13)


### Features

* Add error classifier ([#90](https://github.com/cloudquery/cq-provider-digitalocean/issues/90)) ([8a94c1f](https://github.com/cloudquery/cq-provider-digitalocean/commit/8a94c1f6a5ea6a45312e4ed7bb4248f5e372f837))

## [0.5.7](https://github.com/cloudquery/cq-provider-digitalocean/compare/v0.5.6...v0.5.7) (2022-06-08)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/cq-provider-sdk to v0.11.0 ([#137](https://github.com/cloudquery/cq-provider-digitalocean/issues/137)) ([9a371d6](https://github.com/cloudquery/cq-provider-digitalocean/commit/9a371d6d4cf0ecbe59d1cf862341343177c3dae3))

## [0.5.6](https://github.com/cloudquery/cq-provider-digitalocean/compare/v0.5.5...v0.5.6) (2022-06-07)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/cq-provider-sdk to v0.10.10 ([#135](https://github.com/cloudquery/cq-provider-digitalocean/issues/135)) ([c37926c](https://github.com/cloudquery/cq-provider-digitalocean/commit/c37926ca9d0249312b9342915a7a2f3b08c742f2))
* **deps:** Update module github.com/cloudquery/cq-provider-sdk to v0.10.11 ([#136](https://github.com/cloudquery/cq-provider-digitalocean/issues/136)) ([8943f0f](https://github.com/cloudquery/cq-provider-digitalocean/commit/8943f0ffd95475fe125f618770866d225b6c6735))
* **deps:** Update module github.com/cloudquery/cq-provider-sdk to v0.10.9 ([#133](https://github.com/cloudquery/cq-provider-digitalocean/issues/133)) ([157bfc8](https://github.com/cloudquery/cq-provider-digitalocean/commit/157bfc8e181fad73a58d0aaab8f768645de2926f))

## [0.5.5](https://github.com/cloudquery/cq-provider-digitalocean/compare/v0.5.4...v0.5.5) (2022-06-07)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/cq-provider-sdk to v0.10.8 ([#131](https://github.com/cloudquery/cq-provider-digitalocean/issues/131)) ([8ab4c31](https://github.com/cloudquery/cq-provider-digitalocean/commit/8ab4c310203f06d87473e22c09dcd6b877fe7aaa))

## [0.5.4](https://github.com/cloudquery/cq-provider-digitalocean/compare/v0.5.3...v0.5.4) (2022-06-06)


### Bug Fixes

* Wrap provider errors ([#125](https://github.com/cloudquery/cq-provider-digitalocean/issues/125)) ([2489cb2](https://github.com/cloudquery/cq-provider-digitalocean/commit/2489cb2d3cc775e926a1ae01605204e5743c39a1))

### [0.5.3](https://github.com/cloudquery/cq-provider-digitalocean/compare/v0.5.2...v0.5.3) (2022-06-01)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/cq-provider-sdk to v0.10.7 ([#123](https://github.com/cloudquery/cq-provider-digitalocean/issues/123)) ([e7a64d0](https://github.com/cloudquery/cq-provider-digitalocean/commit/e7a64d027460cd8657300f929e75860d4e2720ba))

### [0.5.2](https://github.com/cloudquery/cq-provider-digitalocean/compare/v0.5.1...v0.5.2) (2022-06-01)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/cq-provider-sdk to v0.10.6 ([#121](https://github.com/cloudquery/cq-provider-digitalocean/issues/121)) ([88eee53](https://github.com/cloudquery/cq-provider-digitalocean/commit/88eee536f02aa97346dffa09dafb48d3442f869f))
* Remove relation tables PK ([#95](https://github.com/cloudquery/cq-provider-digitalocean/issues/95)) ([5b48809](https://github.com/cloudquery/cq-provider-digitalocean/commit/5b4880932508cd0108fa4d87ba7eca10d5bfafd6))

### [0.5.1](https://github.com/cloudquery/cq-provider-digitalocean/compare/v0.5.0...v0.5.1) (2022-05-31)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/cq-provider-sdk to v0.10.2 ([#107](https://github.com/cloudquery/cq-provider-digitalocean/issues/107)) ([17c001b](https://github.com/cloudquery/cq-provider-digitalocean/commit/17c001b786fa9c4f558d0c64b3b8760508e92b32))
* **deps:** Update module github.com/cloudquery/cq-provider-sdk to v0.10.3 ([#109](https://github.com/cloudquery/cq-provider-digitalocean/issues/109)) ([1c14850](https://github.com/cloudquery/cq-provider-digitalocean/commit/1c14850e19c3d8e660f37152eeae4a6e60629768))
* **deps:** Update module github.com/cloudquery/cq-provider-sdk to v0.10.4 ([#112](https://github.com/cloudquery/cq-provider-digitalocean/issues/112)) ([c64d20a](https://github.com/cloudquery/cq-provider-digitalocean/commit/c64d20a843593e8fc5564d1719e14b3043ced2ee))
* **deps:** Update module github.com/cloudquery/cq-provider-sdk to v0.10.5 ([#117](https://github.com/cloudquery/cq-provider-digitalocean/issues/117)) ([376aeab](https://github.com/cloudquery/cq-provider-digitalocean/commit/376aeabadb10b2303bd64c506e597150a07de060))

## [0.5.0](https://github.com/cloudquery/cq-provider-digitalocean/compare/v0.4.4...v0.5.0) (2022-05-24)


### ⚠ BREAKING CHANGES

* Remove migrations (#104)

### Features

* Remove migrations ([#104](https://github.com/cloudquery/cq-provider-digitalocean/issues/104)) ([71b9270](https://github.com/cloudquery/cq-provider-digitalocean/commit/71b927022b6d5b4fc136e863f29fb0107e00d8fd))


### Bug Fixes

* **deps:** Update module github.com/cloudquery/cq-provider-sdk to v0.10.1 ([#103](https://github.com/cloudquery/cq-provider-digitalocean/issues/103)) ([ba59490](https://github.com/cloudquery/cq-provider-digitalocean/commit/ba59490cdb8f969fd3a40c4b234af03ba197e842))

### [0.4.4](https://github.com/cloudquery/cq-provider-digitalocean/compare/v0.4.3...v0.4.4) (2022-05-23)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/cq-provider-sdk to v0.9.5 ([#100](https://github.com/cloudquery/cq-provider-digitalocean/issues/100)) ([99d07ed](https://github.com/cloudquery/cq-provider-digitalocean/commit/99d07edbff267f35a5940771d8d511cb52ae0931))

### [0.4.3](https://github.com/cloudquery/cq-provider-digitalocean/compare/v0.4.2...v0.4.3) (2022-05-17)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/cq-provider-sdk to v0.9.4 ([#92](https://github.com/cloudquery/cq-provider-digitalocean/issues/92)) ([148a6f1](https://github.com/cloudquery/cq-provider-digitalocean/commit/148a6f1dacb9d49c674c03f4f199feb8073880c5))

### [0.4.2](https://github.com/cloudquery/cq-provider-digitalocean/compare/v0.4.1...v0.4.2) (2022-05-09)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/cq-provider-sdk to v0.9.3 ([#69](https://github.com/cloudquery/cq-provider-digitalocean/issues/69)) ([d0de000](https://github.com/cloudquery/cq-provider-digitalocean/commit/d0de000b65fbd6a967edf3e17026addc2432bdd8))

## [v0.2.4] - 2022-01-04
###### SDK Version: 0.6.1
### 💥 Breaking Changes
* Updated to SDK Version [v0.6.1](https://github.com/cloudquery/cq-provider-sdk/blob/main/CHANGELOG.md#v061---2022-01-03) [#27](https://github.com/cloudquery/cq-provider-digitalocean/pull/27)


## [v0.1.0] - 2021-09-12
###### SDK Version: 0.4.3

### :rocket: Added
 - Added Support for the following resources:
    * VPCs
    * CDNs
    * Certificates
    * Spaces
    * Firewalls
    * Registry
    * Alert Policies
    * Floating Ips
    * Databases
    * Load Balancers  
    * Account
    * Droplets
    * Billing History
    * Volumes
    * Regions
    * Sizes
    * Snapshots
    * Projects
    * Keys
    * Domains
    * Balance
    * Images
