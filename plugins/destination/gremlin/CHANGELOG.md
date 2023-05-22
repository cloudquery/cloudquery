# Changelog

## [2.1.0](https://github.com/cloudquery/cloudquery/compare/plugins-destination-gremlin-v2.0.2...plugins-destination-gremlin-v2.1.0) (2023-05-18)


### Features

* **deps:** Upgrade to Apache Arrow v13 (latest `cqmain`) ([#10605](https://github.com/cloudquery/cloudquery/issues/10605)) ([a55da3d](https://github.com/cloudquery/cloudquery/commit/a55da3dbefafdc68a6bda2d5f1d334d12dd97b97))
* **gremlin:** Add `complete_types` option ([#10846](https://github.com/cloudquery/cloudquery/issues/10846)) ([43166e1](https://github.com/cloudquery/cloudquery/commit/43166e15788a6591df73e4ae23d6129556339bd2)), closes [#10839](https://github.com/cloudquery/cloudquery/issues/10839)
* **gremlin:** Upgrade to SDK v3 ([#10821](https://github.com/cloudquery/cloudquery/issues/10821)) ([27fc337](https://github.com/cloudquery/cloudquery/commit/27fc33799aecfb4d33b5c21add30cd59e0a472ba))


### Bug Fixes

* **deps:** Update module github.com/aws/aws-sdk-go-v2/config to v1.18.25 ([#10786](https://github.com/cloudquery/cloudquery/issues/10786)) ([caca1a4](https://github.com/cloudquery/cloudquery/commit/caca1a41e298c06afb6f474b8fd911c4544a2eec))
* **deps:** Update module github.com/cloudquery/plugin-pb-go to v1.0.8 ([#10798](https://github.com/cloudquery/cloudquery/issues/10798)) ([27ff430](https://github.com/cloudquery/cloudquery/commit/27ff430527932d59a4d488a6767547eda8853940))
* **gremlin:** Upgrade gremlingo to official 3.6.3 ([#10824](https://github.com/cloudquery/cloudquery/issues/10824)) ([656f2a2](https://github.com/cloudquery/cloudquery/commit/656f2a297e3612b29757a9c9de6b7a7afb563684))

## [2.0.2](https://github.com/cloudquery/cloudquery/compare/plugins-destination-gremlin-v2.0.1...plugins-destination-gremlin-v2.0.2) (2023-05-09)


### Bug Fixes

* **deps:** Update module github.com/aws/aws-sdk-go-v2/config to v1.18.23 ([#10576](https://github.com/cloudquery/cloudquery/issues/10576)) ([eeb13d5](https://github.com/cloudquery/cloudquery/commit/eeb13d5b1b6b6fcb32764c8711bfbb79da35f9a8))

## [2.0.1](https://github.com/cloudquery/cloudquery/compare/plugins-destination-gremlin-v2.0.0...plugins-destination-gremlin-v2.0.1) (2023-05-02)


### Bug Fixes

* **deps:** Update github.com/apache/arrow/go/v12 digest to 0ea1a10 ([#10461](https://github.com/cloudquery/cloudquery/issues/10461)) ([022709f](https://github.com/cloudquery/cloudquery/commit/022709f710cc6d95aee60260d6f58991698bbf42))
* **deps:** Update github.com/apache/tinkerpop/gremlin-go/v3 digest to ca452a5 ([#10462](https://github.com/cloudquery/cloudquery/issues/10462)) ([6aa93eb](https://github.com/cloudquery/cloudquery/commit/6aa93ebc0cdcafd55f528298e1d7ae55730b7f8f))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v2 to v2.5.0 ([#10390](https://github.com/cloudquery/cloudquery/issues/10390)) ([f706688](https://github.com/cloudquery/cloudquery/commit/f706688b2f5b8393d09d57020d31fb1d280f0dbd))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v2 to v2.5.1 ([#10448](https://github.com/cloudquery/cloudquery/issues/10448)) ([cc85b93](https://github.com/cloudquery/cloudquery/commit/cc85b939fe945939caf72f8c08095e1e744b9ee8))

## [2.0.0](https://github.com/cloudquery/cloudquery/compare/plugins-destination-gremlin-v1.1.0...plugins-destination-gremlin-v2.0.0) (2023-04-25)


### ⚠ BREAKING CHANGES

* This release introduces an internal change to our type system to use [Apache Arrow](https://arrow.apache.org/). This should not have any visible breaking changes, however due to the size of the change we are introducing it under a major version bump to communicate that it might have some bugs that we weren't able to catch during our internal tests. If you encounter an issue during the upgrade, please submit a [bug report](https://github.com/cloudquery/cloudquery/issues/new/choose).

### Features

* Update to use [Apache Arrow](https://arrow.apache.org/) type system ([82b258c](https://github.com/cloudquery/cloudquery/commit/82b258c8300187bb88e0988fb74e26540d41624b))


### Bug Fixes

* **deps:** Update module github.com/aws/aws-sdk-go-v2/config to v1.18.21 ([#10127](https://github.com/cloudquery/cloudquery/issues/10127)) ([3bcde69](https://github.com/cloudquery/cloudquery/commit/3bcde697c5f927fa4eab52ea4293f1f7724812d1))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.45.0 ([#9863](https://github.com/cloudquery/cloudquery/issues/9863)) ([2799d62](https://github.com/cloudquery/cloudquery/commit/2799d62518283ac304beecda9478f8f2db43cdc5))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v2 to v2.4.0 ([#10278](https://github.com/cloudquery/cloudquery/issues/10278)) ([a0a713e](https://github.com/cloudquery/cloudquery/commit/a0a713e8490b970b9d8bfaa1b50e01f43ff51c36))

## [1.1.0](https://github.com/cloudquery/cloudquery/compare/plugins-destination-gremlin-v1.0.3...plugins-destination-gremlin-v1.1.0) (2023-04-12)


### Features

* **gremlin:** Refresh IAM auth as necessary ([#9106](https://github.com/cloudquery/cloudquery/issues/9106)) ([423e397](https://github.com/cloudquery/cloudquery/commit/423e39748da6f1263c60858c28cb1a5c51889309))


### Bug Fixes

* **deps:** Update module github.com/aws/aws-sdk-go-v2 to v1.17.8 ([#9781](https://github.com/cloudquery/cloudquery/issues/9781)) ([69bb790](https://github.com/cloudquery/cloudquery/commit/69bb790afbeac9ff01a41e71c8f631fb60fe64d1))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/config to v1.18.20 ([#9782](https://github.com/cloudquery/cloudquery/issues/9782)) ([1febd5b](https://github.com/cloudquery/cloudquery/commit/1febd5bbd944459a2fcbe380eb90385ecccfb079))

## [1.0.3](https://github.com/cloudquery/cloudquery/compare/plugins-destination-gremlin-v1.0.2...plugins-destination-gremlin-v1.0.3) (2023-04-04)


### Bug Fixes

* **deps:** Update golang.org/x/exp digest to 10a5072 ([#9587](https://github.com/cloudquery/cloudquery/issues/9587)) ([31f913f](https://github.com/cloudquery/cloudquery/commit/31f913f8e3538a2ba41b089bb11eae78aaf42ab2))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.44.1 ([#9520](https://github.com/cloudquery/cloudquery/issues/9520)) ([202c31b](https://github.com/cloudquery/cloudquery/commit/202c31b2788c3df35b5df7d07fdc750f92e7bb23))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.44.2 ([#9661](https://github.com/cloudquery/cloudquery/issues/9661)) ([a27dc84](https://github.com/cloudquery/cloudquery/commit/a27dc84a9b67b68b5b75b04dd3afe13e2c556082))
* **deps:** Update module github.com/mattn/go-isatty to v0.0.18 ([#9609](https://github.com/cloudquery/cloudquery/issues/9609)) ([5b2908e](https://github.com/cloudquery/cloudquery/commit/5b2908e8260c6e48f8c5fd6b8bd6c772f0c779d1))

## [1.0.2](https://github.com/cloudquery/cloudquery/compare/plugins-destination-gremlin-v1.0.1...plugins-destination-gremlin-v1.0.2) (2023-03-28)


### Bug Fixes

* **deps:** Update module github.com/aws/aws-sdk-go-v2 to v1.17.7 ([#9425](https://github.com/cloudquery/cloudquery/issues/9425)) ([c8a4ab1](https://github.com/cloudquery/cloudquery/commit/c8a4ab1aaf52a1ae68f816b26b6bf7c47910501e))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/config to v1.18.19 ([#9426](https://github.com/cloudquery/cloudquery/issues/9426)) ([2017697](https://github.com/cloudquery/cloudquery/commit/2017697a59970f61c79e713054e8d3e4e482c453))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/credentials to v1.13.18 ([#9427](https://github.com/cloudquery/cloudquery/issues/9427)) ([b2ef029](https://github.com/cloudquery/cloudquery/commit/b2ef0292574d3fa03b7cba8d8a6d25031210079a))

## [1.0.1](https://github.com/cloudquery/cloudquery/compare/plugins-destination-gremlin-v1.0.0...plugins-destination-gremlin-v1.0.1) (2023-03-21)


### Bug Fixes

* **deps:** Update module github.com/aws/aws-sdk-go-v2/config to v1.18.18 ([#9227](https://github.com/cloudquery/cloudquery/issues/9227)) ([f630ecc](https://github.com/cloudquery/cloudquery/commit/f630ecc28c19e8388626c823954dca9f561e3920))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.44.0 ([#9167](https://github.com/cloudquery/cloudquery/issues/9167)) ([49d6477](https://github.com/cloudquery/cloudquery/commit/49d647730a85ea6fae51e97194ba61c0625d1331))

## 1.0.0 (2023-03-15)


### Features

* Gremlin destination ([#8842](https://github.com/cloudquery/cloudquery/issues/8842)) ([50696a4](https://github.com/cloudquery/cloudquery/commit/50696a4fdaa1c47b5c4ba9ae60dab6a680af79bd))


### Bug Fixes

* **deps:** Update module github.com/aws/aws-sdk-go-v2/internal/configsources to v1.1.30 ([#8887](https://github.com/cloudquery/cloudquery/issues/8887)) ([95e7ee8](https://github.com/cloudquery/cloudquery/commit/95e7ee8f2f5615ebe7e8938173212849261b7ef6))
* **deps:** Update module github.com/aws/aws-sdk-go-v2/internal/endpoints/v2 to v2.4.24 ([#8888](https://github.com/cloudquery/cloudquery/issues/8888)) ([ebe3698](https://github.com/cloudquery/cloudquery/commit/ebe36989f04fa53be7684e6b00ddd0cca8a794f7))
