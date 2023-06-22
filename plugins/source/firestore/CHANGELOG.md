# Changelog

## [2.0.1](https://github.com/cloudquery/cloudquery/compare/plugins-source-firestore-v2.0.0...plugins-source-firestore-v2.0.1) (2023-06-06)


### Bug Fixes

* **deps:** Update github.com/apache/arrow/go/v13 digest to e07e22c ([#11151](https://github.com/cloudquery/cloudquery/issues/11151)) ([5083cf7](https://github.com/cloudquery/cloudquery/commit/5083cf720f0ae98e07448ba2ae1116048e2d3a90))
* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to 20b0de9 ([#11199](https://github.com/cloudquery/cloudquery/issues/11199)) ([dc3565d](https://github.com/cloudquery/cloudquery/commit/dc3565d3fd6a640d9d10b4fd3a7fe6009a9d02a5))
* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to 88d5dc2 ([#11226](https://github.com/cloudquery/cloudquery/issues/11226)) ([9f306bc](https://github.com/cloudquery/cloudquery/commit/9f306bcaf3833b4611f0df5c50277be43aa19cbb))
* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to a7aad4c ([#11184](https://github.com/cloudquery/cloudquery/issues/11184)) ([8a0822e](https://github.com/cloudquery/cloudquery/commit/8a0822e31fc0eef99de2cdd2bd6d7e4c8b4131bf))
* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to c67fb39 ([#11169](https://github.com/cloudquery/cloudquery/issues/11169)) ([dcb0f92](https://github.com/cloudquery/cloudquery/commit/dcb0f9296a770a5cc2eb6bffd6b1ee30fbccb5dc))
* **deps:** Update github.com/golang/groupcache digest to 41bb18b ([#11153](https://github.com/cloudquery/cloudquery/issues/11153)) ([da1648a](https://github.com/cloudquery/cloudquery/commit/da1648a93e39e0e744b887a08702375f82856c42))
* **deps:** Update golang.org/x/exp digest to 2e198f4 ([#11155](https://github.com/cloudquery/cloudquery/issues/11155)) ([c46c62b](https://github.com/cloudquery/cloudquery/commit/c46c62b68692f527485d7f4b84265abc5dc1142c))
* **deps:** Update google.golang.org/genproto digest to e85fd2c ([#11156](https://github.com/cloudquery/cloudquery/issues/11156)) ([dbe7e92](https://github.com/cloudquery/cloudquery/commit/dbe7e9293d693a6821570e0e0b80202a936b6d3c))
* **deps:** Update module github.com/cloudquery/plugin-pb-go to v1.0.9 ([#11240](https://github.com/cloudquery/cloudquery/issues/11240)) ([f92cd4b](https://github.com/cloudquery/cloudquery/commit/f92cd4bfe3c3d0088964d52ab9cd01ca4cf622e1))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v3 to v3.10.3 ([#11150](https://github.com/cloudquery/cloudquery/issues/11150)) ([dc00994](https://github.com/cloudquery/cloudquery/commit/dc00994e32936af7e9893c93561d0f9df225a929))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v3 to v3.10.4 ([#11244](https://github.com/cloudquery/cloudquery/issues/11244)) ([8fceef6](https://github.com/cloudquery/cloudquery/commit/8fceef6f9041e173923555d8ff221cfe83b424c2))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v3 to v3.7.0 ([#11113](https://github.com/cloudquery/cloudquery/issues/11113)) ([487bf87](https://github.com/cloudquery/cloudquery/commit/487bf871afe360cb8d9d592dfea48837d6e7cf27))
* **timestamp:** Use `time.Time` in `__created_at` & `__updated_at` resolvers ([#11213](https://github.com/cloudquery/cloudquery/issues/11213)) ([436a5fb](https://github.com/cloudquery/cloudquery/commit/436a5fba286ed4dff5d5b4ecb7c96c024394909d)), closes [#10391](https://github.com/cloudquery/cloudquery/issues/10391)

## [2.0.0](https://github.com/cloudquery/cloudquery/compare/plugins-source-firestore-v1.0.2...plugins-source-firestore-v2.0.0) (2023-05-30)


### ⚠ BREAKING CHANGES

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

* **deps:** Upgrade to Apache Arrow v13 (latest `cqmain`) ([#10605](https://github.com/cloudquery/cloudquery/issues/10605)) ([a55da3d](https://github.com/cloudquery/cloudquery/commit/a55da3dbefafdc68a6bda2d5f1d334d12dd97b97))
* Update to use [Apache Arrow](https://arrow.apache.org/) type system ([#11010](https://github.com/cloudquery/cloudquery/issues/11010)) ([be60546](https://github.com/cloudquery/cloudquery/commit/be60546c61ef03fbc95120baea57c3eca9a4159f))


### Bug Fixes

* **deps:** Update module github.com/cloudquery/plugin-pb-go to v1.0.8 ([#10798](https://github.com/cloudquery/cloudquery/issues/10798)) ([27ff430](https://github.com/cloudquery/cloudquery/commit/27ff430527932d59a4d488a6767547eda8853940))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v3 to v3.6.7 ([#11043](https://github.com/cloudquery/cloudquery/issues/11043)) ([3c6d885](https://github.com/cloudquery/cloudquery/commit/3c6d885c3d201b0b39cbc1406c6e54a57ec5ed5f))

## [1.0.2](https://github.com/cloudquery/cloudquery/compare/plugins-source-firestore-v1.0.1...plugins-source-firestore-v1.0.2) (2023-05-02)


### Bug Fixes

* **deps:** Update github.com/apache/arrow/go/v12 digest to 0ea1a10 ([#10461](https://github.com/cloudquery/cloudquery/issues/10461)) ([022709f](https://github.com/cloudquery/cloudquery/commit/022709f710cc6d95aee60260d6f58991698bbf42))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v2 to v2.5.0 ([#10390](https://github.com/cloudquery/cloudquery/issues/10390)) ([f706688](https://github.com/cloudquery/cloudquery/commit/f706688b2f5b8393d09d57020d31fb1d280f0dbd))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v2 to v2.5.1 ([#10448](https://github.com/cloudquery/cloudquery/issues/10448)) ([cc85b93](https://github.com/cloudquery/cloudquery/commit/cc85b939fe945939caf72f8c08095e1e744b9ee8))

## [1.0.1](https://github.com/cloudquery/cloudquery/compare/plugins-source-firestore-v1.0.0...plugins-source-firestore-v1.0.1) (2023-04-25)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/plugin-sdk/v2 to v2.2.2 ([#10143](https://github.com/cloudquery/cloudquery/issues/10143)) ([8f887e0](https://github.com/cloudquery/cloudquery/commit/8f887e05de2096e8efd1e55863a8cf3c7620ccc3))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v2 to v2.3.0 ([#10163](https://github.com/cloudquery/cloudquery/issues/10163)) ([9a7f214](https://github.com/cloudquery/cloudquery/commit/9a7f21460772200e7a588409ebc7eb19f97b195b))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v2 to v2.3.1 ([#10175](https://github.com/cloudquery/cloudquery/issues/10175)) ([5b53423](https://github.com/cloudquery/cloudquery/commit/5b53423e72672f6c2bfb8ae00cfce1641410443e))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v2 to v2.3.3 ([#10187](https://github.com/cloudquery/cloudquery/issues/10187)) ([b185248](https://github.com/cloudquery/cloudquery/commit/b1852480b6ec8b721d94c72d8435051352f26932))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v2 to v2.3.4 ([#10196](https://github.com/cloudquery/cloudquery/issues/10196)) ([c6d2f59](https://github.com/cloudquery/cloudquery/commit/c6d2f59c7d77177a351cb82ecdc381dec6aad30c))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v2 to v2.3.5 ([#10200](https://github.com/cloudquery/cloudquery/issues/10200)) ([5a33693](https://github.com/cloudquery/cloudquery/commit/5a33693fe29f7068b03d80be1859d6e479c42c0d))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v2 to v2.3.6 ([#10208](https://github.com/cloudquery/cloudquery/issues/10208)) ([91c80a7](https://github.com/cloudquery/cloudquery/commit/91c80a795b46480447cfaef67c4db721a31e3206))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v2 to v2.3.8 ([#10213](https://github.com/cloudquery/cloudquery/issues/10213)) ([f358666](https://github.com/cloudquery/cloudquery/commit/f35866611cd206c37e6e9f9ad3329561e4cb32af))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v2 to v2.4.0 ([#10278](https://github.com/cloudquery/cloudquery/issues/10278)) ([a0a713e](https://github.com/cloudquery/cloudquery/commit/a0a713e8490b970b9d8bfaa1b50e01f43ff51c36))

## 1.0.0 (2023-04-17)


### Features

* **firestore:** Add firestore source ([#9653](https://github.com/cloudquery/cloudquery/issues/9653)) ([0fe7a6c](https://github.com/cloudquery/cloudquery/commit/0fe7a6c332bd72e02367a9b7e7455c6d42aaf765))
* **firestore:** Add Sentry ([#10003](https://github.com/cloudquery/cloudquery/issues/10003)) ([aad54bc](https://github.com/cloudquery/cloudquery/commit/aad54bc0ff0e89aef1976809e2dd747456e78656))
* **firestore:** Upgrade to `github.com/cloudquery/plugin-sdk/v2` ([#9956](https://github.com/cloudquery/cloudquery/issues/9956)) ([44ae1ba](https://github.com/cloudquery/cloudquery/commit/44ae1bacdfc1a20669697cffb99a1e8f7cc92f19)), closes [#9954](https://github.com/cloudquery/cloudquery/issues/9954)
* **website:** Add Firestore ([#10007](https://github.com/cloudquery/cloudquery/issues/10007)) ([9586efc](https://github.com/cloudquery/cloudquery/commit/9586efc136eece57a1cb911fb681dc35cb3305b6))


### Bug Fixes

* **deps:** Update module github.com/cloudquery/plugin-sdk/v2 to v2.2.0 ([#10135](https://github.com/cloudquery/cloudquery/issues/10135)) ([cf33b89](https://github.com/cloudquery/cloudquery/commit/cf33b892ead0bb231e3956aa70967de552a21624))
* **firestore:** Change default version to 'development' ([#10005](https://github.com/cloudquery/cloudquery/issues/10005)) ([8527b66](https://github.com/cloudquery/cloudquery/commit/8527b66c7a1d8e8562daa793bba970b4fa9a2ea0))
