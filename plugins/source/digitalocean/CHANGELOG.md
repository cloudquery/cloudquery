# Changelog

All notable changes to this provider will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [3.0.3](https://github.com/cloudquery/cloudquery/compare/plugins-source-digitalocean-v3.0.2...plugins-source-digitalocean-v3.0.3) (2022-11-23)


### Bug Fixes

* **deps:** Update module github.com/aws/aws-sdk-go-v2/config to v1.18.3 ([#4943](https://github.com/cloudquery/cloudquery/issues/4943)) ([e4aaf3f](https://github.com/cloudquery/cloudquery/commit/e4aaf3f0976a0836301b5de70a8e933c7abb5365))
* **deps:** Update plugin-sdk for digitalocean to v1.7.0 ([#4908](https://github.com/cloudquery/cloudquery/issues/4908)) ([d78f127](https://github.com/cloudquery/cloudquery/commit/d78f1271801a2c1894d2d530306dac85efaf2abe))
* **deps:** Update plugin-sdk for digitalocean to v1.8.0 ([#4971](https://github.com/cloudquery/cloudquery/issues/4971)) ([1af8cae](https://github.com/cloudquery/cloudquery/commit/1af8cae144893ff1ceca1d5d9e7f9e8c7de620c4))

## [3.0.2](https://github.com/cloudquery/cloudquery/compare/plugins-source-digitalocean-v3.0.1...plugins-source-digitalocean-v3.0.2) (2022-11-21)


### Bug Fixes

* **deps:** Update plugin-sdk for digitalocean to v1.6.0 ([#4846](https://github.com/cloudquery/cloudquery/issues/4846)) ([8ef4e8b](https://github.com/cloudquery/cloudquery/commit/8ef4e8b5c8428e13ee2dbcb41d3adb282beb0e33))

## [3.0.1](https://github.com/cloudquery/cloudquery/compare/plugins-source-digitalocean-v3.0.0...plugins-source-digitalocean-v3.0.1) (2022-11-18)


### Bug Fixes

* **deps:** Update module github.com/aws/aws-sdk-go-v2/config to v1.18.1 ([#4736](https://github.com/cloudquery/cloudquery/issues/4736)) ([db70d2a](https://github.com/cloudquery/cloudquery/commit/db70d2a602fc4edfc74ed61fd7d28ada6da6a3af))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/config to v1.18.2 ([#4775](https://github.com/cloudquery/cloudquery/issues/4775)) ([136fb42](https://github.com/cloudquery/cloudquery/commit/136fb4213da150f8f9e4a68019fbe7fe94397370))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/s3 to v1.29.3 ([#4779](https://github.com/cloudquery/cloudquery/issues/4779)) ([9332cfb](https://github.com/cloudquery/cloudquery/commit/9332cfbfbd7ec9b16e849f8ea28d2ae0e0e7508d))

## [3.0.0](https://github.com/cloudquery/cloudquery/compare/plugins-source-digitalocean-v2.2.22...plugins-source-digitalocean-v3.0.0) (2022-11-16)


### ⚠ BREAKING CHANGES

* **digitalocean:** Remove PK from `id` column in `digitalocean_space_cors` table. This is only a breaking change if you did not use spaces, and start using spaces after this version. Before updating you'll need to drop your database.

### Bug Fixes

* **digitalocean:** Remove PK from `id` column in `digitalocean_space_cors` table ([e43ea01](https://github.com/cloudquery/cloudquery/commit/e43ea014da595cb8cb7cdab2d268661c2b47db05))

## [2.2.22](https://github.com/cloudquery/cloudquery/compare/plugins-source-digitalocean-v2.2.21...plugins-source-digitalocean-v2.2.22) (2022-11-15)


### Bug Fixes

* **deps:** Update plugin-sdk for digitalocean to v1.5.3 ([#4645](https://github.com/cloudquery/cloudquery/issues/4645)) ([d3ac14c](https://github.com/cloudquery/cloudquery/commit/d3ac14cc4372333f49099772da193718257905de))

## [2.2.21](https://github.com/cloudquery/cloudquery/compare/plugins-source-digitalocean-v2.2.20...plugins-source-digitalocean-v2.2.21) (2022-11-14)


### Bug Fixes

* **deps:** Update plugin-sdk for digitalocean to v1.5.2 ([#4560](https://github.com/cloudquery/cloudquery/issues/4560)) ([da4bb2d](https://github.com/cloudquery/cloudquery/commit/da4bb2d85af7709208dd8a3d5bbb56bedef657c3))

## [2.2.20](https://github.com/cloudquery/cloudquery/compare/plugins-source-digitalocean-v2.2.19...plugins-source-digitalocean-v2.2.20) (2022-11-14)


### Bug Fixes

* **deps:** Update plugin-sdk for digitalocean to v1.5.1 ([#4501](https://github.com/cloudquery/cloudquery/issues/4501)) ([59fba9d](https://github.com/cloudquery/cloudquery/commit/59fba9d8a5ae9c294b2486be193d550e15f9d5f9))

## [2.2.19](https://github.com/cloudquery/cloudquery/compare/plugins-source-digitalocean-v2.2.18...plugins-source-digitalocean-v2.2.19) (2022-11-13)


### Bug Fixes

* **deps:** Update module github.com/aws/aws-sdk-go-v2/config to v1.18.0 ([#4447](https://github.com/cloudquery/cloudquery/issues/4447)) ([2453e88](https://github.com/cloudquery/cloudquery/commit/2453e880718bccd5ddeee5a046697495ddef23c9))

## [2.2.18](https://github.com/cloudquery/cloudquery/compare/plugins-source-digitalocean-v2.2.17...plugins-source-digitalocean-v2.2.18) (2022-11-13)


### Bug Fixes

* **deps:** Update module github.com/aws/aws-sdk-go-v2/service/s3 to v1.29.2 ([#4458](https://github.com/cloudquery/cloudquery/issues/4458)) ([19cece0](https://github.com/cloudquery/cloudquery/commit/19cece035b31401b1d7786daf5299d638ec86689))

## [2.2.17](https://github.com/cloudquery/cloudquery/compare/plugins-source-digitalocean-v2.2.16...plugins-source-digitalocean-v2.2.17) (2022-11-11)


### Bug Fixes

* **deps:** Update plugin-sdk for digitalocean to v1.5.0 ([#4390](https://github.com/cloudquery/cloudquery/issues/4390)) ([7bfad68](https://github.com/cloudquery/cloudquery/commit/7bfad68e9e8cccf4d545e27d684583cb831da3da))

## [2.2.16](https://github.com/cloudquery/cloudquery/compare/plugins-source-digitalocean-v2.2.15...plugins-source-digitalocean-v2.2.16) (2022-11-10)


### Bug Fixes

* **deps:** Update plugin-sdk for digitalocean to v1.4.1 ([#4293](https://github.com/cloudquery/cloudquery/issues/4293)) ([40e3ab8](https://github.com/cloudquery/cloudquery/commit/40e3ab8aa6838e07bd7350d526400fb9807a972b))

## [2.2.15](https://github.com/cloudquery/cloudquery/compare/plugins-source-digitalocean-v2.2.14...plugins-source-digitalocean-v2.2.15) (2022-11-10)


### Bug Fixes

* **deps:** Update plugin-sdk for digitalocean to v1.4.0 ([#4231](https://github.com/cloudquery/cloudquery/issues/4231)) ([6ed0913](https://github.com/cloudquery/cloudquery/commit/6ed0913856672a409dd21c240587868c665032ce))

## [2.2.14](https://github.com/cloudquery/cloudquery/compare/plugins-source-digitalocean-v2.2.13...plugins-source-digitalocean-v2.2.14) (2022-11-10)


### Bug Fixes

* **deps:** Update plugin-sdk for digitalocean to v1.3.2 ([#4198](https://github.com/cloudquery/cloudquery/issues/4198)) ([bb7c0c9](https://github.com/cloudquery/cloudquery/commit/bb7c0c967479be6af2c1a81d4f53e39735c3ad3e))

## [2.2.13](https://github.com/cloudquery/cloudquery/compare/plugins-source-digitalocean-v2.2.12...plugins-source-digitalocean-v2.2.13) (2022-11-10)


### Bug Fixes

* **deps:** Update plugin-sdk for digitalocean to v1.3.1 ([#4149](https://github.com/cloudquery/cloudquery/issues/4149)) ([38f34d2](https://github.com/cloudquery/cloudquery/commit/38f34d2fab7bbf202f31adbac5c7ee5052e1f003))

## [2.2.12](https://github.com/cloudquery/cloudquery/compare/plugins-source-digitalocean-v2.2.11...plugins-source-digitalocean-v2.2.12) (2022-11-10)


### Bug Fixes

* **deps:** Update module github.com/aws/aws-sdk-go-v2 to v1.17.1 ([#4124](https://github.com/cloudquery/cloudquery/issues/4124)) ([650b3be](https://github.com/cloudquery/cloudquery/commit/650b3beb358e429f3737c407ab012bd379021c29))
* **deps:** Update module github.com/aws/smithy-go to v1.13.4 ([#4123](https://github.com/cloudquery/cloudquery/issues/4123)) ([b4b8372](https://github.com/cloudquery/cloudquery/commit/b4b83721a316bbaa2435004b28b3d6357bc7eb5f))

## [2.2.11](https://github.com/cloudquery/cloudquery/compare/plugins-source-digitalocean-v2.2.10...plugins-source-digitalocean-v2.2.11) (2022-11-09)


### Bug Fixes

* **deps:** Update plugin-sdk for digitalocean to v1.3.0 ([#4073](https://github.com/cloudquery/cloudquery/issues/4073)) ([d9c6e73](https://github.com/cloudquery/cloudquery/commit/d9c6e732dcbe2b46d52106c6ce4fdbd3ace657a1))

## [2.2.10](https://github.com/cloudquery/cloudquery/compare/plugins-source-digitalocean-v2.2.9...plugins-source-digitalocean-v2.2.10) (2022-11-09)


### Bug Fixes

* **deps:** Update plugin-sdk for digitalocean to v1.2.0 ([#4041](https://github.com/cloudquery/cloudquery/issues/4041)) ([1ab7ff0](https://github.com/cloudquery/cloudquery/commit/1ab7ff0187513778f93f515e89861c48364816d0))

## [2.2.9](https://github.com/cloudquery/cloudquery/compare/plugins-source-digitalocean-v2.2.8...plugins-source-digitalocean-v2.2.9) (2022-11-09)


### Bug Fixes

* **deps:** Update plugin-sdk for csv to v1.1.0 ([#3918](https://github.com/cloudquery/cloudquery/issues/3918)) ([f1acd68](https://github.com/cloudquery/cloudquery/commit/f1acd688fcd90011cc9be1be2285e3fe9369e341))

## [2.2.8](https://github.com/cloudquery/cloudquery/compare/plugins-source-digitalocean-v2.2.7...plugins-source-digitalocean-v2.2.8) (2022-11-08)


### Bug Fixes

* **deps:** Update plugin-sdk for digitalocean to v1.1.0 ([#3919](https://github.com/cloudquery/cloudquery/issues/3919)) ([3a2ee23](https://github.com/cloudquery/cloudquery/commit/3a2ee23358187f0a5133855f0031c4046a324de6))

## [2.2.7](https://github.com/cloudquery/cloudquery/compare/plugins-source-digitalocean-v2.2.6...plugins-source-digitalocean-v2.2.7) (2022-11-08)


### Bug Fixes

* **deps:** Update plugin-sdk for digitalocean to v1.0.3 ([#3850](https://github.com/cloudquery/cloudquery/issues/3850)) ([5c96a8e](https://github.com/cloudquery/cloudquery/commit/5c96a8e629fc5b8a4b8326b68b6bcb612c777659))
* **deps:** Upgrade plugin-sdk to v1.0.4 for plugins ([#3889](https://github.com/cloudquery/cloudquery/issues/3889)) ([6767243](https://github.com/cloudquery/cloudquery/commit/6767243ec70bfae7a4c457bf4b5edf013c54c392))

## [2.2.6](https://github.com/cloudquery/cloudquery/compare/plugins-source-digitalocean-v2.2.5...plugins-source-digitalocean-v2.2.6) (2022-11-07)


### Bug Fixes

* **deps:** Update plugin-sdk for digitalocean to v1 ([#3779](https://github.com/cloudquery/cloudquery/issues/3779)) ([e5311b9](https://github.com/cloudquery/cloudquery/commit/e5311b92c73d505210a701cd847e564b329fa058))

## [2.2.5](https://github.com/cloudquery/cloudquery/compare/plugins-source-digitalocean-v2.2.4...plugins-source-digitalocean-v2.2.5) (2022-11-07)


### Bug Fixes

* **deps:** Update SDK to v0.13.23 ([#3740](https://github.com/cloudquery/cloudquery/issues/3740)) ([5e66c6d](https://github.com/cloudquery/cloudquery/commit/5e66c6d5d1a0550e5722369eb90f77a949c1344c))

## [2.2.4](https://github.com/cloudquery/cloudquery/compare/plugins-source-digitalocean-v2.2.3...plugins-source-digitalocean-v2.2.4) (2022-11-06)


### Bug Fixes

* **deps:** Update plugin-sdk for digitalocean to v0.13.22 ([#3682](https://github.com/cloudquery/cloudquery/issues/3682)) ([f7d95d0](https://github.com/cloudquery/cloudquery/commit/f7d95d034923f71675795955b803a5f73a753937))

## [2.2.3](https://github.com/cloudquery/cloudquery/compare/plugins-source-digitalocean-v2.2.2...plugins-source-digitalocean-v2.2.3) (2022-11-06)


### Bug Fixes

* **deps:** Update plugin-sdk for digitalocean to v0.13.21 ([#3634](https://github.com/cloudquery/cloudquery/issues/3634)) ([101eabe](https://github.com/cloudquery/cloudquery/commit/101eabe30b456e4be937934428563844fb4def6f))

## [2.2.2](https://github.com/cloudquery/cloudquery/compare/plugins-source-digitalocean-v2.2.1...plugins-source-digitalocean-v2.2.2) (2022-11-04)


### Bug Fixes

* **deps:** Update plugin-sdk for digitalocean to v0.13.20 ([#3574](https://github.com/cloudquery/cloudquery/issues/3574)) ([2c00e07](https://github.com/cloudquery/cloudquery/commit/2c00e07938ae3b537aaa905cf11184c607de0ca8))

## [2.2.1](https://github.com/cloudquery/cloudquery/compare/plugins-source-digitalocean-v2.2.0...plugins-source-digitalocean-v2.2.1) (2022-11-03)


### Bug Fixes

* **deps:** Update plugin-sdk for digitalocean to v0.13.19 ([#3505](https://github.com/cloudquery/cloudquery/issues/3505)) ([9f3a1a4](https://github.com/cloudquery/cloudquery/commit/9f3a1a4ef51c759a2899802a9669eafae9c139a9))

## [2.2.0](https://github.com/cloudquery/cloudquery/compare/plugins-source-digitalocean-v2.1.0...plugins-source-digitalocean-v2.2.0) (2022-11-01)


### Features

* Migrate cli, plugins and destinations to new type system ([#3323](https://github.com/cloudquery/cloudquery/issues/3323)) ([f265a94](https://github.com/cloudquery/cloudquery/commit/f265a94448ad55c968b26ba8a19681bc81086c11))


### Bug Fixes

* **deps:** Update plugin-sdk for digitalocean to v0.13.17 ([#3403](https://github.com/cloudquery/cloudquery/issues/3403)) ([e2f2255](https://github.com/cloudquery/cloudquery/commit/e2f2255084c527c601ac82b515bfd722b5370952))
* **deps:** Update plugin-sdk for digitalocean to v0.13.18 ([#3412](https://github.com/cloudquery/cloudquery/issues/3412)) ([9d19db1](https://github.com/cloudquery/cloudquery/commit/9d19db1313dfd4f97d6f7f9856d44467da8ac2b6))

## [2.1.0](https://github.com/cloudquery/cloudquery/compare/plugins-source-digitalocean-v2.0.8...plugins-source-digitalocean-v2.1.0) (2022-10-31)


### Features

* Update all plugins to SDK with metrics and DFS scheduler ([#3286](https://github.com/cloudquery/cloudquery/issues/3286)) ([a35b8e8](https://github.com/cloudquery/cloudquery/commit/a35b8e89d625287a9b9406ff18cfac78ffdb1241))

## [2.0.8](https://github.com/cloudquery/cloudquery/compare/plugins-source-digitalocean-v2.0.7...plugins-source-digitalocean-v2.0.8) (2022-10-27)


### Bug Fixes

* **deps:** Update plugin-sdk for digitalocean to v0.13.13 ([#3215](https://github.com/cloudquery/cloudquery/issues/3215)) ([2753d46](https://github.com/cloudquery/cloudquery/commit/2753d467f5fc5e1a253fad4fd0f8945779d171b6))
* **deps:** Update plugin-sdk for digitalocean to v0.13.14 ([#3231](https://github.com/cloudquery/cloudquery/issues/3231)) ([7282748](https://github.com/cloudquery/cloudquery/commit/728274896eb4d12788c39f1ec060f84ca4172eac))

## [2.0.7](https://github.com/cloudquery/cloudquery/compare/plugins-source-digitalocean-v2.0.6...plugins-source-digitalocean-v2.0.7) (2022-10-20)


### Bug Fixes

* **deps:** Update plugin-sdk for digitalocean to v0.13.12 ([#3103](https://github.com/cloudquery/cloudquery/issues/3103)) ([daf0708](https://github.com/cloudquery/cloudquery/commit/daf0708ec904332bce41540f680fbb90b56ebbc6))

## [2.0.6](https://github.com/cloudquery/cloudquery/compare/plugins-source-digitalocean-v2.0.5...plugins-source-digitalocean-v2.0.6) (2022-10-20)


### Bug Fixes

* **deps:** Update plugin-sdk for digitalocean to v0.13.12 ([#3103](https://github.com/cloudquery/cloudquery/issues/3103)) ([daf0708](https://github.com/cloudquery/cloudquery/commit/daf0708ec904332bce41540f680fbb90b56ebbc6))

## [2.0.5](https://github.com/cloudquery/cloudquery/compare/plugins-source-digitalocean-v2.0.4...plugins-source-digitalocean-v2.0.5) (2022-10-19)


### Bug Fixes

* **deps:** Update plugin-sdk to v0.13.11 ([#3030](https://github.com/cloudquery/cloudquery/issues/3030)) ([9909c4a](https://github.com/cloudquery/cloudquery/commit/9909c4a0715a06b7c1d69c9bd23c500ac7b4adc1))

## [2.0.4](https://github.com/cloudquery/cloudquery/compare/plugins-source-digitalocean-v2.0.3...plugins-source-digitalocean-v2.0.4) (2022-10-18)


### Bug Fixes

* **deps:** Update plugin-sdk for digitalocean to v0.13.9 ([#2930](https://github.com/cloudquery/cloudquery/issues/2930)) ([76eb5ba](https://github.com/cloudquery/cloudquery/commit/76eb5ba0407fe6056f5b3bbe456e7fb52667575f))

## [2.0.3](https://github.com/cloudquery/cloudquery/compare/plugins-source-digitalocean-v2.0.2...plugins-source-digitalocean-v2.0.3) (2022-10-14)


### Bug Fixes

* **deps:** Update plugin-sdk for digitalocean to v0.13.8 ([#2852](https://github.com/cloudquery/cloudquery/issues/2852)) ([c48bed4](https://github.com/cloudquery/cloudquery/commit/c48bed4733f25563fef3db81be2e847ace5992f7))

## [2.0.2](https://github.com/cloudquery/cloudquery/compare/plugins-source-digitalocean-v2.0.1...plugins-source-digitalocean-v2.0.2) (2022-10-13)


### Bug Fixes

* **deps:** Update plugin-sdk for digitalocean to v0.13.7 ([#2782](https://github.com/cloudquery/cloudquery/issues/2782)) ([f3e6678](https://github.com/cloudquery/cloudquery/commit/f3e66783421ea8e6e7175d62ad8e75f001a6ddb0))

## [2.0.1](https://github.com/cloudquery/cloudquery/compare/plugins-source-digitalocean-v2.0.0...plugins-source-digitalocean-v2.0.1) (2022-10-12)


### Bug Fixes

* **deps:** Update plugin-sdk for digitalocean to v0.13.6 ([#2721](https://github.com/cloudquery/cloudquery/issues/2721)) ([3068622](https://github.com/cloudquery/cloudquery/commit/3068622eb871218567da12f06ba890c2e0dc71e0))

## [2.0.0](https://github.com/cloudquery/cloudquery/compare/plugins-source-digitalocean-v1.0.1...plugins-source-digitalocean-v2.0.0) (2022-10-12)


### ⚠ BREAKING CHANGES

* Rename certificates `sha_1_fingerprint` to `sha1_fingerprint` (#2629)

### Bug Fixes

* **deps:** Update plugin-sdk for digitalocean to v0.12.10 ([#2548](https://github.com/cloudquery/cloudquery/issues/2548)) ([e26297e](https://github.com/cloudquery/cloudquery/commit/e26297e5ec0f05b467f153802ed7cf83b6c24066))
* Rename certificates `sha_1_fingerprint` to `sha1_fingerprint` ([#2629](https://github.com/cloudquery/cloudquery/issues/2629)) ([6f5aba9](https://github.com/cloudquery/cloudquery/commit/6f5aba94b929810f6f3445dcaaf8947a1c26a997))
* Upgrade source SDK versions to v0.13.5 ([#2610](https://github.com/cloudquery/cloudquery/issues/2610)) ([611868e](https://github.com/cloudquery/cloudquery/commit/611868e7fbb707b524ccc5c04a7ff95fe122ae05))

## [1.0.1](https://github.com/cloudquery/cloudquery/compare/plugins-source-digitalocean-v1.0.0...plugins-source-digitalocean-v1.0.1) (2022-10-09)


### Bug Fixes

* **deps:** Update plugin-sdk for digitalocean to v0.12.3 ([#2356](https://github.com/cloudquery/cloudquery/issues/2356)) ([63a500e](https://github.com/cloudquery/cloudquery/commit/63a500e32820a52940807a2bd040dd3bf82b9e5b))
* **deps:** Update plugin-sdk for digitalocean to v0.12.4 ([#2398](https://github.com/cloudquery/cloudquery/issues/2398)) ([6e70184](https://github.com/cloudquery/cloudquery/commit/6e7018429c69bb65e00c5c089f2457721c5d5617))
* **deps:** Update plugin-sdk for digitalocean to v0.12.5 ([#2420](https://github.com/cloudquery/cloudquery/issues/2420)) ([940e9ab](https://github.com/cloudquery/cloudquery/commit/940e9ab0eef25cd5a32e9eea280b620d8b35b641))
* **deps:** Update plugin-sdk for digitalocean to v0.12.6 ([#2436](https://github.com/cloudquery/cloudquery/issues/2436)) ([379a6a6](https://github.com/cloudquery/cloudquery/commit/379a6a6d718ef2ada97e5991e7d938600fcf85c4))
* **deps:** Update plugin-sdk for digitalocean to v0.12.7 ([#2449](https://github.com/cloudquery/cloudquery/issues/2449)) ([121983d](https://github.com/cloudquery/cloudquery/commit/121983d27d74a2c09ad4e7cca9e9fbe6e06f6b02))
* **deps:** Update plugin-sdk for digitalocean to v0.12.8 ([#2499](https://github.com/cloudquery/cloudquery/issues/2499)) ([e346804](https://github.com/cloudquery/cloudquery/commit/e3468040148fad827f3b24262d6f815532465743))
* **deps:** Update plugin-sdk for digitalocean to v0.12.9 ([#2513](https://github.com/cloudquery/cloudquery/issues/2513)) ([43e2eab](https://github.com/cloudquery/cloudquery/commit/43e2eab158df7be3b52266705cad2b15452ef358))

## [1.0.0](https://github.com/cloudquery/cloudquery/compare/plugins/source/digitalocean/v0.6.5...plugins-source-digitalocean-v1.0.0) (2022-10-04)


### ⚠ BREAKING CHANGES

* [Official v1 release](https://www.cloudquery.io/blog/cloudquery-v1-release)

### Features

* [Official v1 release](https://www.cloudquery.io/blog/cloudquery-v1-release)

## [0.7.4-pre.0](https://github.com/cloudquery/cloudquery/compare/plugins-source-digitalocean-v0.7.3-pre.0...plugins-source-digitalocean-v0.7.4-pre.0) (2022-10-03)


### Bug Fixes

* **deps:** Update plugin-sdk for digitalocean to v0.11.6 ([#2255](https://github.com/cloudquery/cloudquery/issues/2255)) ([c144272](https://github.com/cloudquery/cloudquery/commit/c144272ea6b4026b8eb0fd761934989bb0948d1c))

## [0.7.3-pre.0](https://github.com/cloudquery/cloudquery/compare/plugins-source-digitalocean-v0.7.2-pre.0...plugins-source-digitalocean-v0.7.3-pre.0) (2022-10-03)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/plugin-sdk to v0.11.5 ([#2229](https://github.com/cloudquery/cloudquery/issues/2229)) ([51dcc5d](https://github.com/cloudquery/cloudquery/commit/51dcc5ded3d20862c76cddd6f2b51035d7eef5f2))

## [0.7.2-pre.0](https://github.com/cloudquery/cloudquery/compare/plugins-source-digitalocean-v0.7.1-pre.0...plugins-source-digitalocean-v0.7.2-pre.0) (2022-10-02)


### Bug Fixes

* **deps:** Update module github.com/aws/aws-sdk-go-v2 to v1.16.16 ([#2179](https://github.com/cloudquery/cloudquery/issues/2179)) ([de378c0](https://github.com/cloudquery/cloudquery/commit/de378c0f183130caa56c1520a79d1a1c187b2941))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v0.10.2 ([#2048](https://github.com/cloudquery/cloudquery/issues/2048)) ([e407991](https://github.com/cloudquery/cloudquery/commit/e4079914772d8191639b9935aa5970b8e27b082f))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v0.11.0 ([#2135](https://github.com/cloudquery/cloudquery/issues/2135)) ([1729467](https://github.com/cloudquery/cloudquery/commit/1729467b2119555e18b15d73c91cd501ccf7ecb8))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v0.11.2 ([#2162](https://github.com/cloudquery/cloudquery/issues/2162)) ([5701aa5](https://github.com/cloudquery/cloudquery/commit/5701aa5b0a8d04e9e99e3efe6e27d5f7ff29b216))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v0.11.4 ([#2212](https://github.com/cloudquery/cloudquery/issues/2212)) ([35b54e7](https://github.com/cloudquery/cloudquery/commit/35b54e7eeedff5f2c36daca0e7e4cd79c57ea848))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v0.8.0 ([#1997](https://github.com/cloudquery/cloudquery/issues/1997)) ([4fa40da](https://github.com/cloudquery/cloudquery/commit/4fa40da04b427f864d2dc11f133e5c83e53ce4b6))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v0.8.1 ([#2024](https://github.com/cloudquery/cloudquery/issues/2024)) ([8f88de4](https://github.com/cloudquery/cloudquery/commit/8f88de4b4eaeabae7369ba309e765a252392ee8c))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v0.8.2 ([#2044](https://github.com/cloudquery/cloudquery/issues/2044)) ([9b69b46](https://github.com/cloudquery/cloudquery/commit/9b69b468536521b20b77ec1fc180fc85aeeba376))

## [0.7.1-pre.0](https://github.com/cloudquery/cloudquery/compare/plugins/source/digitalocean/v0.7.0-pre.0...plugins/source/digitalocean/v0.7.1-pre.0) (2022-09-22)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/plugin-sdk to v0.7.12 ([#1916](https://github.com/cloudquery/cloudquery/issues/1916)) ([27d8153](https://github.com/cloudquery/cloudquery/commit/27d81534baaa1312a6bd87294d298dd8b5348a79))

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
