# Changelog

## [4.0.2](https://github.com/cloudquery/cloudquery/compare/plugins-destination-duckdb-v4.0.1...plugins-destination-duckdb-v4.0.2) (2023-06-13)


### Bug Fixes

* Don't update `UNIQUE` columns (DuckDB contraint) ([#11516](https://github.com/cloudquery/cloudquery/issues/11516)) ([697a9d9](https://github.com/cloudquery/cloudquery/commit/697a9d9a19fe05e5f93d34e45497fbaf3dab1b00))

## [4.0.1](https://github.com/cloudquery/cloudquery/compare/plugins-destination-duckdb-v4.0.0...plugins-destination-duckdb-v4.0.1) (2023-06-13)


### Bug Fixes

* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to 0f7bd3b ([#11412](https://github.com/cloudquery/cloudquery/issues/11412)) ([dd1e2e8](https://github.com/cloudquery/cloudquery/commit/dd1e2e892d95515fd7332339262abaefd2a256c5))
* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to 7f6aaff ([#11432](https://github.com/cloudquery/cloudquery/issues/11432)) ([55dfebc](https://github.com/cloudquery/cloudquery/commit/55dfebc064608fb47caaf3b8e68c8002de8a7dc3))
* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to 8f72077 ([#11395](https://github.com/cloudquery/cloudquery/issues/11395)) ([d91fc5c](https://github.com/cloudquery/cloudquery/commit/d91fc5ce24f64c29fff6988b19ec2c2775cc379b))
* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to 90670b8 ([#11279](https://github.com/cloudquery/cloudquery/issues/11279)) ([a6cdc91](https://github.com/cloudquery/cloudquery/commit/a6cdc912e4b38a3faf798c5147a986ffe2539643))
* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to b359e74 ([#11405](https://github.com/cloudquery/cloudquery/issues/11405)) ([5d92765](https://github.com/cloudquery/cloudquery/commit/5d927659bd4f7c445a0e312487f1655ffb9a60f6))
* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to d8eacf8 ([#11449](https://github.com/cloudquery/cloudquery/issues/11449)) ([742dafd](https://github.com/cloudquery/cloudquery/commit/742dafd5bf5cdc8facb94fda5de1d84c88897cbd))
* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to e258cfb ([#11391](https://github.com/cloudquery/cloudquery/issues/11391)) ([eacbe9a](https://github.com/cloudquery/cloudquery/commit/eacbe9ad3ea16d88f27c4593fa2774574ac8fe4e))

## [4.0.0](https://github.com/cloudquery/cloudquery/compare/plugins-destination-duckdb-v3.0.0...plugins-destination-duckdb-v4.0.0) (2023-06-06)


### ⚠ BREAKING CHANGES

* **types:** Support writing Apache Arrow nested types:
    * Structs as DuckDB structs
    * Maps as DuckDB maps

### Features

* **types:** Proper support for nested types ([#11196](https://github.com/cloudquery/cloudquery/issues/11196)) ([7c6a3e2](https://github.com/cloudquery/cloudquery/commit/7c6a3e2d24632ffd0ac3eacd6bd65b89394ebe0f))


### Bug Fixes

* **deps:** Update `github.com/cloudquery/plugin-sdk/v3` to `v3.10.0` ([#11116](https://github.com/cloudquery/cloudquery/issues/11116)) ([bba7c4e](https://github.com/cloudquery/cloudquery/commit/bba7c4ef9368741ed00e5c04bdc2bc9a1de9a521))
* **deps:** Update github.com/apache/arrow/go/v13 digest to e07e22c ([#11151](https://github.com/cloudquery/cloudquery/issues/11151)) ([5083cf7](https://github.com/cloudquery/cloudquery/commit/5083cf720f0ae98e07448ba2ae1116048e2d3a90))
* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to 88d5dc2 ([#11226](https://github.com/cloudquery/cloudquery/issues/11226)) ([9f306bc](https://github.com/cloudquery/cloudquery/commit/9f306bcaf3833b4611f0df5c50277be43aa19cbb))
* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to a7aad4c ([#11184](https://github.com/cloudquery/cloudquery/issues/11184)) ([8a0822e](https://github.com/cloudquery/cloudquery/commit/8a0822e31fc0eef99de2cdd2bd6d7e4c8b4131bf))
* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to c67fb39 ([#11169](https://github.com/cloudquery/cloudquery/issues/11169)) ([dcb0f92](https://github.com/cloudquery/cloudquery/commit/dcb0f9296a770a5cc2eb6bffd6b1ee30fbccb5dc))
* **deps:** Update golang.org/x/exp digest to 2e198f4 ([#11155](https://github.com/cloudquery/cloudquery/issues/11155)) ([c46c62b](https://github.com/cloudquery/cloudquery/commit/c46c62b68692f527485d7f4b84265abc5dc1142c))
* **deps:** Update google.golang.org/genproto digest to e85fd2c ([#11156](https://github.com/cloudquery/cloudquery/issues/11156)) ([dbe7e92](https://github.com/cloudquery/cloudquery/commit/dbe7e9293d693a6821570e0e0b80202a936b6d3c))
* **deps:** Update module github.com/cloudquery/plugin-pb-go to v1.0.9 ([#11240](https://github.com/cloudquery/cloudquery/issues/11240)) ([f92cd4b](https://github.com/cloudquery/cloudquery/commit/f92cd4bfe3c3d0088964d52ab9cd01ca4cf622e1))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v3 to v3.10.3 ([#11150](https://github.com/cloudquery/cloudquery/issues/11150)) ([dc00994](https://github.com/cloudquery/cloudquery/commit/dc00994e32936af7e9893c93561d0f9df225a929))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v3 to v3.10.4 ([#11244](https://github.com/cloudquery/cloudquery/issues/11244)) ([8fceef6](https://github.com/cloudquery/cloudquery/commit/8fceef6f9041e173923555d8ff221cfe83b424c2))
* **race:** Fix issue caused by [Over-Eager Unique Constraint Checking](https://duckdb.org/docs/sql/indexes#over-eager-unique-constraint-checking) in DuckDB overwrite ([#11215](https://github.com/cloudquery/cloudquery/issues/11215)) ([c0b9f0a](https://github.com/cloudquery/cloudquery/commit/c0b9f0af37b70b7a386ae4b1c20e52794db87681))

## [3.0.0](https://github.com/cloudquery/cloudquery/compare/plugins-destination-duckdb-v2.0.1...plugins-destination-duckdb-v3.0.0) (2023-05-24)


### ⚠ BREAKING CHANGES

* **duckdb:** Move DuckDB write to Parquet ([#10874](https://github.com/cloudquery/cloudquery/issues/10874))

### Features

* **deps:** Upgrade to Apache Arrow v13 (latest `cqmain`) ([#10605](https://github.com/cloudquery/cloudquery/issues/10605)) ([a55da3d](https://github.com/cloudquery/cloudquery/commit/a55da3dbefafdc68a6bda2d5f1d334d12dd97b97))
* **duckdb:** Migrate to SDK V3 ([#10874](https://github.com/cloudquery/cloudquery/issues/10874)) ([84e6631](https://github.com/cloudquery/cloudquery/commit/84e663193b5cecdeb56f9a5debcd4ff59e1c49bb))


### Bug Fixes

* **deps:** Update module github.com/cloudquery/plugin-pb-go to v1.0.8 ([#10798](https://github.com/cloudquery/cloudquery/issues/10798)) ([27ff430](https://github.com/cloudquery/cloudquery/commit/27ff430527932d59a4d488a6767547eda8853940))

## [2.0.1](https://github.com/cloudquery/cloudquery/compare/plugins-destination-duckdb-v2.0.0...plugins-destination-duckdb-v2.0.1) (2023-05-02)


### Bug Fixes

* **deps:** Update github.com/apache/arrow/go/v12 digest to 0ea1a10 ([#10461](https://github.com/cloudquery/cloudquery/issues/10461)) ([022709f](https://github.com/cloudquery/cloudquery/commit/022709f710cc6d95aee60260d6f58991698bbf42))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v2 to v2.5.0 ([#10390](https://github.com/cloudquery/cloudquery/issues/10390)) ([f706688](https://github.com/cloudquery/cloudquery/commit/f706688b2f5b8393d09d57020d31fb1d280f0dbd))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v2 to v2.5.1 ([#10448](https://github.com/cloudquery/cloudquery/issues/10448)) ([cc85b93](https://github.com/cloudquery/cloudquery/commit/cc85b939fe945939caf72f8c08095e1e744b9ee8))

## [2.0.0](https://github.com/cloudquery/cloudquery/compare/plugins-destination-duckdb-v1.0.6...plugins-destination-duckdb-v2.0.0) (2023-04-26)


### ⚠ BREAKING CHANGES

* This release introduces an internal change to our type system to use [Apache Arrow](https://arrow.apache.org/). This should not have any visible breaking changes, however due to the size of the change we are introducing it under a major version bump to communicate that it might have some bugs that we weren't able to catch during our internal tests. If you encounter an issue during the upgrade, please submit a [bug report](https://github.com/cloudquery/cloudquery/issues/new/choose).

### Features

* Update to use [Apache Arrow](https://arrow.apache.org/) type system ([e38eae6](https://github.com/cloudquery/cloudquery/commit/e38eae6bffbdd34f5959ff3cd7124b789ed2dd26))

## [1.0.6](https://github.com/cloudquery/cloudquery/compare/plugins-destination-duckdb-v1.0.5...plugins-destination-duckdb-v1.0.6) (2023-04-25)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.45.0 ([#9863](https://github.com/cloudquery/cloudquery/issues/9863)) ([2799d62](https://github.com/cloudquery/cloudquery/commit/2799d62518283ac304beecda9478f8f2db43cdc5))

## [1.0.5](https://github.com/cloudquery/cloudquery/compare/plugins-destination-duckdb-v1.0.4...plugins-destination-duckdb-v1.0.5) (2023-04-09)


### Bug Fixes

* **duckdb:** Fix multiple pks in delete from ([#9775](https://github.com/cloudquery/cloudquery/issues/9775)) ([706217f](https://github.com/cloudquery/cloudquery/commit/706217fef50703ee228f4df20d95eb55d934fb86))

## [1.0.4](https://github.com/cloudquery/cloudquery/compare/plugins-destination-duckdb-v1.0.3...plugins-destination-duckdb-v1.0.4) (2023-04-09)


### Bug Fixes

* **duckdb:** Support multiple PKs ([#9772](https://github.com/cloudquery/cloudquery/issues/9772)) ([d94d87d](https://github.com/cloudquery/cloudquery/commit/d94d87da16ab6df52dd3705c5f6c60cd151a26fc))

## [1.0.3](https://github.com/cloudquery/cloudquery/compare/plugins-destination-duckdb-v1.0.2...plugins-destination-duckdb-v1.0.3) (2023-04-04)


### Bug Fixes

* **deps:** Update ghcr.io/gythialy/golang-cross Docker tag to v1.20.2 ([#9599](https://github.com/cloudquery/cloudquery/issues/9599)) ([46ce2dc](https://github.com/cloudquery/cloudquery/commit/46ce2dc3165e87f4017608c883c7e4ede3d8b19d))
* **deps:** Update golang.org/x/exp digest to 10a5072 ([#9587](https://github.com/cloudquery/cloudquery/issues/9587)) ([31f913f](https://github.com/cloudquery/cloudquery/commit/31f913f8e3538a2ba41b089bb11eae78aaf42ab2))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.44.1 ([#9520](https://github.com/cloudquery/cloudquery/issues/9520)) ([202c31b](https://github.com/cloudquery/cloudquery/commit/202c31b2788c3df35b5df7d07fdc750f92e7bb23))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.44.2 ([#9661](https://github.com/cloudquery/cloudquery/issues/9661)) ([a27dc84](https://github.com/cloudquery/cloudquery/commit/a27dc84a9b67b68b5b75b04dd3afe13e2c556082))
* **deps:** Update module github.com/marcboeker/go-duckdb to v1.2.2 ([#9608](https://github.com/cloudquery/cloudquery/issues/9608)) ([4dfbc9e](https://github.com/cloudquery/cloudquery/commit/4dfbc9e5e9b892c053b878fa60568045459a17d1))
* **deps:** Update module github.com/mattn/go-isatty to v0.0.18 ([#9609](https://github.com/cloudquery/cloudquery/issues/9609)) ([5b2908e](https://github.com/cloudquery/cloudquery/commit/5b2908e8260c6e48f8c5fd6b8bd6c772f0c779d1))

## [1.0.2](https://github.com/cloudquery/cloudquery/compare/plugins-destination-duckdb-v1.0.1...plugins-destination-duckdb-v1.0.2) (2023-03-21)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.44.0 ([#9167](https://github.com/cloudquery/cloudquery/issues/9167)) ([49d6477](https://github.com/cloudquery/cloudquery/commit/49d647730a85ea6fae51e97194ba61c0625d1331))

## [1.0.1](https://github.com/cloudquery/cloudquery/compare/plugins-destination-duckdb-v1.0.0...plugins-destination-duckdb-v1.0.1) (2023-03-14)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.43.0 ([#8949](https://github.com/cloudquery/cloudquery/issues/8949)) ([31dfc63](https://github.com/cloudquery/cloudquery/commit/31dfc634850b699ba7bb7876399270a7367d6c7e))

## 1.0.0 (2023-03-09)


### Features

* DuckDB destination ([#8758](https://github.com/cloudquery/cloudquery/issues/8758)) ([2ed9c37](https://github.com/cloudquery/cloudquery/commit/2ed9c37708d7df595ce633d3b13099c6074086c6))
