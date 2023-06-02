# Changelog

## [2.1.0](https://github.com/cloudquery/cloudquery/compare/plugins-destination-mysql-v2.0.2...plugins-destination-mysql-v2.1.0) (2023-06-01)


### Features

* **deps:** Upgrade to Apache Arrow v13 (latest `cqmain`) ([#10605](https://github.com/cloudquery/cloudquery/issues/10605)) ([a55da3d](https://github.com/cloudquery/cloudquery/commit/a55da3dbefafdc68a6bda2d5f1d334d12dd97b97))
* **mysql:** Migrate to SDK V3 native arrow ([#10867](https://github.com/cloudquery/cloudquery/issues/10867)) ([f28fc65](https://github.com/cloudquery/cloudquery/commit/f28fc6575adaba5a898d17bf35a2ba168c8767b0))
* **mysql:** Update to SDK v3.6.3 ([#10943](https://github.com/cloudquery/cloudquery/issues/10943)) ([927e8fc](https://github.com/cloudquery/cloudquery/commit/927e8fc4dd3d4ee12d4a77b26acc886e0d8a55e2))


### Bug Fixes

* **deps:** Update github.com/apache/arrow/go/v13 digest to e07e22c ([#11151](https://github.com/cloudquery/cloudquery/issues/11151)) ([5083cf7](https://github.com/cloudquery/cloudquery/commit/5083cf720f0ae98e07448ba2ae1116048e2d3a90))
* **deps:** Update golang.org/x/exp digest to 2e198f4 ([#11155](https://github.com/cloudquery/cloudquery/issues/11155)) ([c46c62b](https://github.com/cloudquery/cloudquery/commit/c46c62b68692f527485d7f4b84265abc5dc1142c))
* **deps:** Update google.golang.org/genproto digest to e85fd2c ([#11156](https://github.com/cloudquery/cloudquery/issues/11156)) ([dbe7e92](https://github.com/cloudquery/cloudquery/commit/dbe7e9293d693a6821570e0e0b80202a936b6d3c))
* **deps:** Update module github.com/cloudquery/plugin-pb-go to v1.0.8 ([#10798](https://github.com/cloudquery/cloudquery/issues/10798)) ([27ff430](https://github.com/cloudquery/cloudquery/commit/27ff430527932d59a4d488a6767547eda8853940))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v3 to v3.6.7 ([#11043](https://github.com/cloudquery/cloudquery/issues/11043)) ([3c6d885](https://github.com/cloudquery/cloudquery/commit/3c6d885c3d201b0b39cbc1406c6e54a57ec5ed5f))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v3 to v3.7.0 ([#11113](https://github.com/cloudquery/cloudquery/issues/11113)) ([487bf87](https://github.com/cloudquery/cloudquery/commit/487bf871afe360cb8d9d592dfea48837d6e7cf27))

## [2.0.2](https://github.com/cloudquery/cloudquery/compare/plugins-destination-mysql-v2.0.1...plugins-destination-mysql-v2.0.2) (2023-05-02)


### Bug Fixes

* **deps:** Update github.com/apache/arrow/go/v12 digest to 0ea1a10 ([#10461](https://github.com/cloudquery/cloudquery/issues/10461)) ([022709f](https://github.com/cloudquery/cloudquery/commit/022709f710cc6d95aee60260d6f58991698bbf42))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v2 to v2.5.1 ([#10448](https://github.com/cloudquery/cloudquery/issues/10448)) ([cc85b93](https://github.com/cloudquery/cloudquery/commit/cc85b939fe945939caf72f8c08095e1e744b9ee8))
* **mysql:** Account for no PK in append mode migration ([#10413](https://github.com/cloudquery/cloudquery/issues/10413)) ([7c7e0d4](https://github.com/cloudquery/cloudquery/commit/7c7e0d40e088db5a7026bf6ef8e88df7e3773dbe))

## [2.0.1](https://github.com/cloudquery/cloudquery/compare/plugins-destination-mysql-v2.0.0...plugins-destination-mysql-v2.0.1) (2023-04-25)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/plugin-sdk/v2 to v2.3.8 ([#10213](https://github.com/cloudquery/cloudquery/issues/10213)) ([f358666](https://github.com/cloudquery/cloudquery/commit/f35866611cd206c37e6e9f9ad3329561e4cb32af))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v2 to v2.4.0 ([#10278](https://github.com/cloudquery/cloudquery/issues/10278)) ([a0a713e](https://github.com/cloudquery/cloudquery/commit/a0a713e8490b970b9d8bfaa1b50e01f43ff51c36))

## [2.0.0](https://github.com/cloudquery/cloudquery/compare/plugins-destination-mysql-v1.0.5...plugins-destination-mysql-v2.0.0) (2023-04-20)


### ⚠ BREAKING CHANGES

* This release introduces an internal change to our type system to use [Apache Arrow](https://arrow.apache.org/). This should not have any visible breaking changes, however due to the size of the change we are introducing it under a major version bump to communicate that it might have some bugs that we weren't able to catch during our internal tests. If you encounter an issue during the upgrade, please submit a [bug report](https://github.com/cloudquery/cloudquery/issues/new/choose).
* Timestamp fields type changed from `datetime` to `datetime(6)` to avoid losing fractional data
* Mac Address and Inet column type changed from `nvarchar(255)` to `text` as we cannot assume `nvarchar(255)` columns represent valid Mac Address or Inet
* Float fields type changed from `float` to `double` for increased precision

### Features

* Float fields type changed from `float` to `double` for increased precision  ([c5d3508](https://github.com/cloudquery/cloudquery/commit/c5d3508ddc5b95579e059f9532e2e64453b5ed86))
* Mac Address and Inet column type changed from `nvarchar(255)` to `text` as we cannot assume `nvarchar(255)` columns represent valid Mac Address or Inet  ([c5d3508](https://github.com/cloudquery/cloudquery/commit/c5d3508ddc5b95579e059f9532e2e64453b5ed86))
* Timestamp fields type changed from `datetime` to `datetime(6)` to avoid losing fractional data ([c5d3508](https://github.com/cloudquery/cloudquery/commit/c5d3508ddc5b95579e059f9532e2e64453b5ed86))
* Update to plugin SDK v2  ([c5d3508](https://github.com/cloudquery/cloudquery/commit/c5d3508ddc5b95579e059f9532e2e64453b5ed86))
* Update to use [Apache Arrow](https://arrow.apache.org/) type system ([c5d3508](https://github.com/cloudquery/cloudquery/commit/c5d3508ddc5b95579e059f9532e2e64453b5ed86))


### Bug Fixes

* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.45.0 ([#9863](https://github.com/cloudquery/cloudquery/issues/9863)) ([2799d62](https://github.com/cloudquery/cloudquery/commit/2799d62518283ac304beecda9478f8f2db43cdc5))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v2 to v2.3.6 ([#10208](https://github.com/cloudquery/cloudquery/issues/10208)) ([91c80a7](https://github.com/cloudquery/cloudquery/commit/91c80a795b46480447cfaef67c4db721a31e3206))
* Update to SDK v2.3.7, remove release calls ([#10209](https://github.com/cloudquery/cloudquery/issues/10209)) ([5442837](https://github.com/cloudquery/cloudquery/commit/544283754bda58ba5c053b2bc55a97de0f408e96))

## [1.0.5](https://github.com/cloudquery/cloudquery/compare/plugins-destination-mysql-v1.0.4...plugins-destination-mysql-v1.0.5) (2023-04-04)


### Bug Fixes

* **deps:** Update golang.org/x/exp digest to 10a5072 ([#9587](https://github.com/cloudquery/cloudquery/issues/9587)) ([31f913f](https://github.com/cloudquery/cloudquery/commit/31f913f8e3538a2ba41b089bb11eae78aaf42ab2))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.44.1 ([#9520](https://github.com/cloudquery/cloudquery/issues/9520)) ([202c31b](https://github.com/cloudquery/cloudquery/commit/202c31b2788c3df35b5df7d07fdc750f92e7bb23))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.44.2 ([#9661](https://github.com/cloudquery/cloudquery/issues/9661)) ([a27dc84](https://github.com/cloudquery/cloudquery/commit/a27dc84a9b67b68b5b75b04dd3afe13e2c556082))
* **deps:** Update module github.com/mattn/go-isatty to v0.0.18 ([#9609](https://github.com/cloudquery/cloudquery/issues/9609)) ([5b2908e](https://github.com/cloudquery/cloudquery/commit/5b2908e8260c6e48f8c5fd6b8bd6c772f0c779d1))

## [1.0.4](https://github.com/cloudquery/cloudquery/compare/plugins-destination-mysql-v1.0.3...plugins-destination-mysql-v1.0.4) (2023-03-21)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.44.0 ([#9167](https://github.com/cloudquery/cloudquery/issues/9167)) ([49d6477](https://github.com/cloudquery/cloudquery/commit/49d647730a85ea6fae51e97194ba61c0625d1331))

## [1.0.3](https://github.com/cloudquery/cloudquery/compare/plugins-destination-mysql-v1.0.2...plugins-destination-mysql-v1.0.3) (2023-03-14)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.43.0 ([#8949](https://github.com/cloudquery/cloudquery/issues/8949)) ([31dfc63](https://github.com/cloudquery/cloudquery/commit/31dfc634850b699ba7bb7876399270a7367d6c7e))

## [1.0.2](https://github.com/cloudquery/cloudquery/compare/plugins-destination-mysql-v1.0.1...plugins-destination-mysql-v1.0.2) (2023-03-13)


### Bug Fixes

* **mysql-migrate:** Properly set `NotNull` when a column is a primary Key ([#8924](https://github.com/cloudquery/cloudquery/issues/8924)) ([8c1ffed](https://github.com/cloudquery/cloudquery/commit/8c1ffedbaa5de4b9575aa2b9d5d6e7b760534f43))

## [1.0.1](https://github.com/cloudquery/cloudquery/compare/plugins-destination-mysql-v1.0.0...plugins-destination-mysql-v1.0.1) (2023-03-07)


### Bug Fixes

* **deps:** Update golang.org/x/exp digest to c95f2b4 ([#8560](https://github.com/cloudquery/cloudquery/issues/8560)) ([9c3bd5b](https://github.com/cloudquery/cloudquery/commit/9c3bd5b68f9741a360fde6c54bf3f5f3efe06d9e))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.41.0 ([#8682](https://github.com/cloudquery/cloudquery/issues/8682)) ([ea9d065](https://github.com/cloudquery/cloudquery/commit/ea9d065ae9f77c6dd990570974630ae6ac3f153e))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.42.0 ([#8725](https://github.com/cloudquery/cloudquery/issues/8725)) ([b83b277](https://github.com/cloudquery/cloudquery/commit/b83b277a2421d1caf46a26c3229041b27a3da148))
* **deps:** Update module github.com/stretchr/testify to v1.8.2 ([#8599](https://github.com/cloudquery/cloudquery/issues/8599)) ([2ec8086](https://github.com/cloudquery/cloudquery/commit/2ec808677328410cc96c97a693ef65022d314c32))

## 1.0.0 (2023-02-27)


### Features

* Add MySQL destination plugin ([#8438](https://github.com/cloudquery/cloudquery/issues/8438)) ([e85b29f](https://github.com/cloudquery/cloudquery/commit/e85b29fcab574b574a78105c7ed6ed5ad1c4ff0d))
