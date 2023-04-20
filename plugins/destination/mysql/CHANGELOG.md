# Changelog

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
