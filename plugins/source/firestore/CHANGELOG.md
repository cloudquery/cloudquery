# Changelog

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
