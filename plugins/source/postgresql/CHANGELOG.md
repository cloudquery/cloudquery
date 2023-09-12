# Changelog

## [3.0.4](https://github.com/cloudquery/cloudquery/compare/plugins-source-postgresql-v3.0.3...plugins-source-postgresql-v3.0.4) (2023-09-12)


### Bug Fixes

* **deps:** Update github.com/cloudquery/arrow/go/v14 digest to 483f6b2 ([#13780](https://github.com/cloudquery/cloudquery/issues/13780)) ([8d31b44](https://github.com/cloudquery/cloudquery/commit/8d31b44f787f42d47f186cdcc4a5739a3a370a5f))

## [3.0.3](https://github.com/cloudquery/cloudquery/compare/plugins-source-postgresql-v3.0.2...plugins-source-postgresql-v3.0.3) (2023-09-05)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.7.0 ([#13625](https://github.com/cloudquery/cloudquery/issues/13625)) ([bb5463f](https://github.com/cloudquery/cloudquery/commit/bb5463fb5919f50f1327eecae884b2ab99fb8b34))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.7.1 ([#13713](https://github.com/cloudquery/cloudquery/issues/13713)) ([73004dc](https://github.com/cloudquery/cloudquery/commit/73004dcabd05bf474d8b5960b8c747a894b98560))

## [3.0.2](https://github.com/cloudquery/cloudquery/compare/plugins-source-postgresql-v3.0.1...plugins-source-postgresql-v3.0.2) (2023-09-05)


### Bug Fixes

* **deps:** Update github.com/99designs/go-keychain digest to 9cf53c8 ([#13561](https://github.com/cloudquery/cloudquery/issues/13561)) ([a170256](https://github.com/cloudquery/cloudquery/commit/a17025657e92b017fe3c8bd37abfaa2354e6e818))
* **deps:** Update github.com/apache/arrow/go/v14 digest to a526ba6 ([#13562](https://github.com/cloudquery/cloudquery/issues/13562)) ([248672b](https://github.com/cloudquery/cloudquery/commit/248672beb020828cde1cb608d5c1ed6d656c777b))
* **deps:** Update github.com/cloudquery/arrow/go/v14 digest to cd3d411 ([#13598](https://github.com/cloudquery/cloudquery/issues/13598)) ([f22bfa6](https://github.com/cloudquery/cloudquery/commit/f22bfa6b2d4fd0caeacf0726ccd307db38f8860c))
* **deps:** Update github.com/jackc/pglogrepl digest to 9ed16cb ([#13564](https://github.com/cloudquery/cloudquery/issues/13564)) ([05af8c7](https://github.com/cloudquery/cloudquery/commit/05af8c7622c2aff1d402b2361c96f1a01c631ed0))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.6.0 ([#13492](https://github.com/cloudquery/cloudquery/issues/13492)) ([c305876](https://github.com/cloudquery/cloudquery/commit/c305876e3d92944aa6c1a26547f786fdc5b50e23))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.6.4 ([#13519](https://github.com/cloudquery/cloudquery/issues/13519)) ([9d25165](https://github.com/cloudquery/cloudquery/commit/9d25165820703844c6de96688d939aa5033608ae))
* Duplicate columns when multiple contraints ([#13612](https://github.com/cloudquery/cloudquery/issues/13612)) ([a9806c5](https://github.com/cloudquery/cloudquery/commit/a9806c5449252326f3357ebf1560510633b3ab29)), closes [#13459](https://github.com/cloudquery/cloudquery/issues/13459)

## [3.0.1](https://github.com/cloudquery/cloudquery/compare/plugins-source-postgresql-v3.0.0...plugins-source-postgresql-v3.0.1) (2023-08-29)


### Bug Fixes

* **deps:** Update `github.com/cloudquery/arrow/go/v13` to `github.com/apache/arrow/go/v14` ([#13341](https://github.com/cloudquery/cloudquery/issues/13341)) ([feb8f87](https://github.com/cloudquery/cloudquery/commit/feb8f87d8d761eb9c49ce84329ad0397f730a918))
* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to 5b83d4f ([#13203](https://github.com/cloudquery/cloudquery/issues/13203)) ([b0a4b8c](https://github.com/cloudquery/cloudquery/commit/b0a4b8ccf7c429bf5a6ed88866865212015b68e4))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.5.1 ([#13195](https://github.com/cloudquery/cloudquery/issues/13195)) ([a184c37](https://github.com/cloudquery/cloudquery/commit/a184c3786ad49df8564344773e9b96f617ef87a1))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.5.2 ([#13262](https://github.com/cloudquery/cloudquery/issues/13262)) ([5c55aa3](https://github.com/cloudquery/cloudquery/commit/5c55aa35282786375e8ce9493b2a4878e0fb27bc))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.5.5 ([#13285](https://github.com/cloudquery/cloudquery/issues/13285)) ([e076abd](https://github.com/cloudquery/cloudquery/commit/e076abd9d67813a29ced0c1b7b1664fd728b9ba8))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.5.6 ([#13345](https://github.com/cloudquery/cloudquery/issues/13345)) ([a995a05](https://github.com/cloudquery/cloudquery/commit/a995a0598a209e0fe3ba09f4ced2a052dc14b67a))
* Document limitations of Postgres CDC syncing ([#13193](https://github.com/cloudquery/cloudquery/issues/13193)) ([92a9577](https://github.com/cloudquery/cloudquery/commit/92a9577a005da927428f07f10d0f377b83901211))

## [3.0.0](https://github.com/cloudquery/cloudquery/compare/plugins-source-postgresql-v2.0.6...plugins-source-postgresql-v3.0.0) (2023-08-15)


### ⚠ BREAKING CHANGES

* Don't mark columns unique if they're part of a composite constraint ([#13134](https://github.com/cloudquery/cloudquery/issues/13134))

### Features

* Allow multiple rows in sync tables per sent record ([#13137](https://github.com/cloudquery/cloudquery/issues/13137)) ([5086346](https://github.com/cloudquery/cloudquery/commit/50863463e3cc72eab8fd72311fb3f312b996cb3c)), closes [#13110](https://github.com/cloudquery/cloudquery/issues/13110)
* Simplify value transformation ([#13129](https://github.com/cloudquery/cloudquery/issues/13129)) ([09c18c3](https://github.com/cloudquery/cloudquery/commit/09c18c3ad6979e96bf732f49f1f1a090768390ff))


### Bug Fixes

* Don't mark columns unique if they're part of a composite constraint ([#13134](https://github.com/cloudquery/cloudquery/issues/13134)) ([b56d267](https://github.com/cloudquery/cloudquery/commit/b56d267c97d5503641d31cd8f4424171fdf47a55)), closes [#13131](https://github.com/cloudquery/cloudquery/issues/13131)
* Handle differences in `format_type` implementation between PostgreSQL & CockroachDB ([#13112](https://github.com/cloudquery/cloudquery/issues/13112)) ([8c03d9e](https://github.com/cloudquery/cloudquery/commit/8c03d9ea469b6ef284dcbededac372c8af972299))

## [2.0.6](https://github.com/cloudquery/cloudquery/compare/plugins-source-postgresql-v2.0.5...plugins-source-postgresql-v2.0.6) (2023-08-15)


### Bug Fixes

* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to e9683e1 ([#13015](https://github.com/cloudquery/cloudquery/issues/13015)) ([6557696](https://github.com/cloudquery/cloudquery/commit/65576966d3bd14297499a5b85d3b4fc2c7918df3))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.4.0 ([#12850](https://github.com/cloudquery/cloudquery/issues/12850)) ([0861200](https://github.com/cloudquery/cloudquery/commit/086120054b45213947e95be954ba6164b9cf6587))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.5.0 ([#13068](https://github.com/cloudquery/cloudquery/issues/13068)) ([7bb0e4b](https://github.com/cloudquery/cloudquery/commit/7bb0e4ba654971726e16a6a501393e3831170307))

## [2.0.5](https://github.com/cloudquery/cloudquery/compare/plugins-source-postgresql-v2.0.4...plugins-source-postgresql-v2.0.5) (2023-08-08)


### Bug Fixes

* **deps:** Update github.com/apache/arrow/go/v13 digest to 112f949 ([#12659](https://github.com/cloudquery/cloudquery/issues/12659)) ([48d73a9](https://github.com/cloudquery/cloudquery/commit/48d73a93e678994f43171c363f5a75c29547b0b9))
* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to 3452eb0 ([#12595](https://github.com/cloudquery/cloudquery/issues/12595)) ([c1c0949](https://github.com/cloudquery/cloudquery/commit/c1c09490b17f2e64435e05d745890cdb8b22310d))
* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to f53878d ([#12778](https://github.com/cloudquery/cloudquery/issues/12778)) ([6f5d58e](https://github.com/cloudquery/cloudquery/commit/6f5d58e3b84d3c76b1d1a3d6c5a488f77995a057))
* **deps:** Update github.com/jackc/pglogrepl digest to 38e8a4e ([#12668](https://github.com/cloudquery/cloudquery/issues/12668)) ([f94bba6](https://github.com/cloudquery/cloudquery/commit/f94bba629f2f3d7ea84dc4d3190c83d810aa1f2d))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.2.4 ([#12718](https://github.com/cloudquery/cloudquery/issues/12718)) ([f059a15](https://github.com/cloudquery/cloudquery/commit/f059a159a2ee406ab2b0a33792c244cd217025a6))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.2.5 ([#12731](https://github.com/cloudquery/cloudquery/issues/12731)) ([d267239](https://github.com/cloudquery/cloudquery/commit/d267239aa3aca5f94bd36a8db1ec0d9f7dc0865f))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.2.6 ([#12799](https://github.com/cloudquery/cloudquery/issues/12799)) ([fb0e0d7](https://github.com/cloudquery/cloudquery/commit/fb0e0d75ab010f421c834e58d93676de76fcb423))

## [2.0.4](https://github.com/cloudquery/cloudquery/compare/plugins-source-postgresql-v2.0.3...plugins-source-postgresql-v2.0.4) (2023-07-26)


### Bug Fixes

* Support syncing tables with a single column ([#12533](https://github.com/cloudquery/cloudquery/issues/12533)) ([1106a24](https://github.com/cloudquery/cloudquery/commit/1106a24e4fc8899772678ad9e0fa1d275f28d0d5))

## [2.0.3](https://github.com/cloudquery/cloudquery/compare/plugins-source-postgresql-v2.0.2...plugins-source-postgresql-v2.0.3) (2023-07-25)


### Bug Fixes

* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to 10df4b9 ([#12443](https://github.com/cloudquery/cloudquery/issues/12443)) ([e385283](https://github.com/cloudquery/cloudquery/commit/e38528309f862f37bc7e278f9b69cf92d5aa5bd5))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.2.3 ([#12307](https://github.com/cloudquery/cloudquery/issues/12307)) ([8f14e4d](https://github.com/cloudquery/cloudquery/commit/8f14e4de7bf4d4c833f501135ea0610916a42f8b))
* Don't create duplicate columns for unique constraints ([#12518](https://github.com/cloudquery/cloudquery/issues/12518)) ([99e7ad9](https://github.com/cloudquery/cloudquery/commit/99e7ad983fb71f009e5c5c3ae461c01b95bae3d4))

## [2.0.2](https://github.com/cloudquery/cloudquery/compare/plugins-source-postgresql-v2.0.1...plugins-source-postgresql-v2.0.2) (2023-07-18)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.2.0 ([#12256](https://github.com/cloudquery/cloudquery/issues/12256)) ([eaec331](https://github.com/cloudquery/cloudquery/commit/eaec33165345ad51fdb6ddbffbf8a1199ebd6384))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.2.1 ([#12272](https://github.com/cloudquery/cloudquery/issues/12272)) ([557ca69](https://github.com/cloudquery/cloudquery/commit/557ca69a7dee9dabb80e6afb6f41f205fd8a80d8))

## [2.0.1](https://github.com/cloudquery/cloudquery/compare/plugins-source-postgresql-v2.0.0...plugins-source-postgresql-v2.0.1) (2023-07-17)


### Bug Fixes

* Fail if can't read current_schema or current_database ([#12238](https://github.com/cloudquery/cloudquery/issues/12238)) ([8a39440](https://github.com/cloudquery/cloudquery/commit/8a39440beed178a3df8e67a2747d19bd962a84c3))

## [2.0.0](https://github.com/cloudquery/cloudquery/compare/plugins-source-postgresql-v1.2.1...plugins-source-postgresql-v2.0.0) (2023-07-17)


### ⚠ BREAKING CHANGES

* To enable CDC in this version you'll need to use the `cdc_id` configuration string property, instead of the `cdc` boolean one. Please refer to the [docs](https://www.cloudquery.io/docs/plugins/sources/postgresql/overview) for more information
* Upgrades the postgresql source plugin to use plugin-sdk v4. This version does not contain any user-facing breaking changes, but because it is now using CloudQuery gRPC protocol v3, it does require use of a destination plugin that also supports protocol v3. All recent destination plugin versions support this.
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

* To enable CDC in this version you'll need to use the `cdc_id` configuration string property, instead of the `cdc` boolean one. Please refer to the [docs](https://www.cloudquery.io/docs/plugins/sources/postgresql/overview) for more information ([d2f9cc7](https://github.com/cloudquery/cloudquery/commit/d2f9cc72ab19f143982f52e7d8dece5a39415ef9))
* Update to use [Apache Arrow](https://arrow.apache.org/) type system ([#11212](https://github.com/cloudquery/cloudquery/issues/11212)) ([0cf8f24](https://github.com/cloudquery/cloudquery/commit/0cf8f247e0141978f70dbc5c225288037c3370a2))
* Upgrades the postgresql source plugin to use plugin-sdk v4. This version does not contain any user-facing breaking changes, but because it is now using CloudQuery gRPC protocol v3, it does require use of a destination plugin that also supports protocol v3. All recent destination plugin versions support this. ([d2f9cc7](https://github.com/cloudquery/cloudquery/commit/d2f9cc72ab19f143982f52e7d8dece5a39415ef9))


### Bug Fixes

* **deps:** Update github.com/apache/arrow/go/v13 digest to 5a06b2e ([#11857](https://github.com/cloudquery/cloudquery/issues/11857)) ([43c2f5f](https://github.com/cloudquery/cloudquery/commit/43c2f5f3a893e5286f67c4943a9d1bc2736e2aeb))
* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to 0656028 ([#11739](https://github.com/cloudquery/cloudquery/issues/11739)) ([7a6ad49](https://github.com/cloudquery/cloudquery/commit/7a6ad49e8402d51e914d6fdc444956c89db91ad3))
* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to 0a52533 ([#12091](https://github.com/cloudquery/cloudquery/issues/12091)) ([927cefa](https://github.com/cloudquery/cloudquery/commit/927cefa943ec3969a2ec39b628bc1eba545a2108))
* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to 1e68c51 ([#11637](https://github.com/cloudquery/cloudquery/issues/11637)) ([46043bc](https://github.com/cloudquery/cloudquery/commit/46043bce410f86ba42390a6b190f9232fc2a1ded))
* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to 43638cb ([#11672](https://github.com/cloudquery/cloudquery/issues/11672)) ([3c60bbb](https://github.com/cloudquery/cloudquery/commit/3c60bbbb0233b17f934583766938780745145864))
* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to 4d76231 ([#11532](https://github.com/cloudquery/cloudquery/issues/11532)) ([6f04233](https://github.com/cloudquery/cloudquery/commit/6f042333acbd2506f7800ccb89a8c5cbfb7ad8d4))
* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to 8366a22 ([#11717](https://github.com/cloudquery/cloudquery/issues/11717)) ([8eeff5b](https://github.com/cloudquery/cloudquery/commit/8eeff5b17486d72845f830b99983f950fee7f5a0))
* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to 8e2219b ([#12220](https://github.com/cloudquery/cloudquery/issues/12220)) ([24e8fb5](https://github.com/cloudquery/cloudquery/commit/24e8fb588740896fe11a660e8b80231e741b753c))
* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to 95d3199 ([#11708](https://github.com/cloudquery/cloudquery/issues/11708)) ([03f214f](https://github.com/cloudquery/cloudquery/commit/03f214f3dfd719b74ce9eb698ba255a8cf7528c7))
* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to a2a76eb ([#12104](https://github.com/cloudquery/cloudquery/issues/12104)) ([311f474](https://github.com/cloudquery/cloudquery/commit/311f4749af2491a606f29483190717a5fe238da6))
* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to b0832be ([#11651](https://github.com/cloudquery/cloudquery/issues/11651)) ([71e8c29](https://github.com/cloudquery/cloudquery/commit/71e8c29624494a3e1cd104e46266a610ce57c83c))
* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to d864719 ([#11611](https://github.com/cloudquery/cloudquery/issues/11611)) ([557a290](https://github.com/cloudquery/cloudquery/commit/557a2903af272b8e2e4c9eebb36e39cd8a41a805))
* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to df3b664 ([#11882](https://github.com/cloudquery/cloudquery/issues/11882)) ([9635b22](https://github.com/cloudquery/cloudquery/commit/9635b22b10a2cd9ca0f91819cffb7f4ba75dc2d9))
* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to f060192 ([#11730](https://github.com/cloudquery/cloudquery/issues/11730)) ([c7019c2](https://github.com/cloudquery/cloudquery/commit/c7019c26c311f29b66c90fc5d461a0daf71d191c))
* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to f0dffc6 ([#11689](https://github.com/cloudquery/cloudquery/issues/11689)) ([18ac0e9](https://github.com/cloudquery/cloudquery/commit/18ac0e9dbef31d06701f1f13d263ad840ac60c5e))
* **deps:** Update github.com/jackc/pglogrepl digest to 5fd22a6 ([#11867](https://github.com/cloudquery/cloudquery/issues/11867)) ([f30e1bf](https://github.com/cloudquery/cloudquery/commit/f30e1bf71048f243b0af335c63aa343a9fb3b4aa))
* **deps:** Update module github.com/cloudquery/plugin-pb-go to v1.1.0 ([#11665](https://github.com/cloudquery/cloudquery/issues/11665)) ([d8947c9](https://github.com/cloudquery/cloudquery/commit/d8947c9efa6ab8bf3952ad9d929e8ed81f2dea55))
* **deps:** Update module github.com/cloudquery/plugin-pb-go to v1.2.0 ([#11720](https://github.com/cloudquery/cloudquery/issues/11720)) ([7ef521d](https://github.com/cloudquery/cloudquery/commit/7ef521db1423c6f0de197b08c73adf22c896f999))
* **deps:** Update module github.com/cloudquery/plugin-pb-go to v1.2.1 ([#11722](https://github.com/cloudquery/cloudquery/issues/11722)) ([309be72](https://github.com/cloudquery/cloudquery/commit/309be7276d7de157013c281b6fb3934513898b3f))
* **deps:** Update module github.com/cloudquery/plugin-pb-go to v1.3.3 ([#11726](https://github.com/cloudquery/cloudquery/issues/11726)) ([f0ca611](https://github.com/cloudquery/cloudquery/commit/f0ca61195014bde707761a15efa27a92955b59db))
* **deps:** Update module github.com/cloudquery/plugin-pb-go to v1.3.4 ([#11753](https://github.com/cloudquery/cloudquery/issues/11753)) ([cd4fe1c](https://github.com/cloudquery/cloudquery/commit/cd4fe1c54f85f8511252bebd5671361618ddb0d3))
* **deps:** Update module github.com/cloudquery/plugin-pb-go to v1.5.0 ([#11850](https://github.com/cloudquery/cloudquery/issues/11850)) ([3255857](https://github.com/cloudquery/cloudquery/commit/3255857938bf16862d52491f5c2a8a0fa53faef0))
* **deps:** Update module github.com/cloudquery/plugin-pb-go to v1.6.0 ([#11916](https://github.com/cloudquery/cloudquery/issues/11916)) ([421e752](https://github.com/cloudquery/cloudquery/commit/421e7529360965175c8d156ff006d2b703ee9da2))
* **deps:** Update module github.com/cloudquery/plugin-pb-go to v1.7.0 ([#12166](https://github.com/cloudquery/cloudquery/issues/12166)) ([94390dd](https://github.com/cloudquery/cloudquery/commit/94390dde19d0c37fee9d035219d62f6ae7edb127))

## [1.2.1](https://github.com/cloudquery/cloudquery/compare/plugins-source-postgresql-v1.2.0...plugins-source-postgresql-v1.2.1) (2023-06-13)


### Bug Fixes

* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to 0f7bd3b ([#11412](https://github.com/cloudquery/cloudquery/issues/11412)) ([dd1e2e8](https://github.com/cloudquery/cloudquery/commit/dd1e2e892d95515fd7332339262abaefd2a256c5))
* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to 7f6aaff ([#11432](https://github.com/cloudquery/cloudquery/issues/11432)) ([55dfebc](https://github.com/cloudquery/cloudquery/commit/55dfebc064608fb47caaf3b8e68c8002de8a7dc3))
* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to 8f72077 ([#11395](https://github.com/cloudquery/cloudquery/issues/11395)) ([d91fc5c](https://github.com/cloudquery/cloudquery/commit/d91fc5ce24f64c29fff6988b19ec2c2775cc379b))
* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to 90670b8 ([#11279](https://github.com/cloudquery/cloudquery/issues/11279)) ([a6cdc91](https://github.com/cloudquery/cloudquery/commit/a6cdc912e4b38a3faf798c5147a986ffe2539643))
* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to b359e74 ([#11405](https://github.com/cloudquery/cloudquery/issues/11405)) ([5d92765](https://github.com/cloudquery/cloudquery/commit/5d927659bd4f7c445a0e312487f1655ffb9a60f6))
* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to d8eacf8 ([#11449](https://github.com/cloudquery/cloudquery/issues/11449)) ([742dafd](https://github.com/cloudquery/cloudquery/commit/742dafd5bf5cdc8facb94fda5de1d84c88897cbd))
* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to e258cfb ([#11391](https://github.com/cloudquery/cloudquery/issues/11391)) ([eacbe9a](https://github.com/cloudquery/cloudquery/commit/eacbe9ad3ea16d88f27c4593fa2774574ac8fe4e))

## [1.2.0](https://github.com/cloudquery/cloudquery/compare/plugins-source-postgresql-v1.1.1...plugins-source-postgresql-v1.2.0) (2023-06-06)


### Features

* **deps:** Upgrade to Apache Arrow v13 (latest `cqmain`) ([#10605](https://github.com/cloudquery/cloudquery/issues/10605)) ([a55da3d](https://github.com/cloudquery/cloudquery/commit/a55da3dbefafdc68a6bda2d5f1d334d12dd97b97))


### Bug Fixes

* **cdc:** Allow dashes in source names when using Postgres CDC ([#11219](https://github.com/cloudquery/cloudquery/issues/11219)) ([1b7f65b](https://github.com/cloudquery/cloudquery/commit/1b7f65bca0f04cf507430b4f629e7ef3c15add39))
* **deps:** Update github.com/apache/arrow/go/v13 digest to e07e22c ([#11151](https://github.com/cloudquery/cloudquery/issues/11151)) ([5083cf7](https://github.com/cloudquery/cloudquery/commit/5083cf720f0ae98e07448ba2ae1116048e2d3a90))
* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to 20b0de9 ([#11199](https://github.com/cloudquery/cloudquery/issues/11199)) ([dc3565d](https://github.com/cloudquery/cloudquery/commit/dc3565d3fd6a640d9d10b4fd3a7fe6009a9d02a5))
* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to 88d5dc2 ([#11226](https://github.com/cloudquery/cloudquery/issues/11226)) ([9f306bc](https://github.com/cloudquery/cloudquery/commit/9f306bcaf3833b4611f0df5c50277be43aa19cbb))
* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to a7aad4c ([#11184](https://github.com/cloudquery/cloudquery/issues/11184)) ([8a0822e](https://github.com/cloudquery/cloudquery/commit/8a0822e31fc0eef99de2cdd2bd6d7e4c8b4131bf))
* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to c67fb39 ([#11169](https://github.com/cloudquery/cloudquery/issues/11169)) ([dcb0f92](https://github.com/cloudquery/cloudquery/commit/dcb0f9296a770a5cc2eb6bffd6b1ee30fbccb5dc))
* **deps:** Update github.com/jackc/pglogrepl digest to 0c5b98f ([#11154](https://github.com/cloudquery/cloudquery/issues/11154)) ([b01b8e5](https://github.com/cloudquery/cloudquery/commit/b01b8e56191ce488283cefe5bfb98b37437ceea7))
* **deps:** Update golang.org/x/exp digest to 2e198f4 ([#11155](https://github.com/cloudquery/cloudquery/issues/11155)) ([c46c62b](https://github.com/cloudquery/cloudquery/commit/c46c62b68692f527485d7f4b84265abc5dc1142c))
* **deps:** Update google.golang.org/genproto digest to e85fd2c ([#11156](https://github.com/cloudquery/cloudquery/issues/11156)) ([dbe7e92](https://github.com/cloudquery/cloudquery/commit/dbe7e9293d693a6821570e0e0b80202a936b6d3c))
* **deps:** Update module github.com/cloudquery/plugin-pb-go to v1.0.8 ([#10798](https://github.com/cloudquery/cloudquery/issues/10798)) ([27ff430](https://github.com/cloudquery/cloudquery/commit/27ff430527932d59a4d488a6767547eda8853940))
* **deps:** Update module github.com/cloudquery/plugin-pb-go to v1.0.9 ([#11240](https://github.com/cloudquery/cloudquery/issues/11240)) ([f92cd4b](https://github.com/cloudquery/cloudquery/commit/f92cd4bfe3c3d0088964d52ab9cd01ca4cf622e1))

## [1.1.1](https://github.com/cloudquery/cloudquery/compare/plugins-source-postgresql-v1.1.0...plugins-source-postgresql-v1.1.1) (2023-05-02)


### Bug Fixes

* **deps:** Update github.com/apache/arrow/go/v12 digest to 0ea1a10 ([#10461](https://github.com/cloudquery/cloudquery/issues/10461)) ([022709f](https://github.com/cloudquery/cloudquery/commit/022709f710cc6d95aee60260d6f58991698bbf42))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v2 to v2.5.0 ([#10390](https://github.com/cloudquery/cloudquery/issues/10390)) ([f706688](https://github.com/cloudquery/cloudquery/commit/f706688b2f5b8393d09d57020d31fb1d280f0dbd))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v2 to v2.5.1 ([#10448](https://github.com/cloudquery/cloudquery/issues/10448)) ([cc85b93](https://github.com/cloudquery/cloudquery/commit/cc85b939fe945939caf72f8c08095e1e744b9ee8))
* **postgresql:** Ensure all table names are escaped ([#10473](https://github.com/cloudquery/cloudquery/issues/10473)) ([62295f2](https://github.com/cloudquery/cloudquery/commit/62295f2979216484d2727037e36d6fe83a06559b))

## [1.1.0](https://github.com/cloudquery/cloudquery/compare/plugins-source-postgresql-v1.0.8...plugins-source-postgresql-v1.1.0) (2023-04-25)


### Features

* **source/postgresql:** Upgrade to `github.com/cloudquery/plugin-sdk/v2` ([#10071](https://github.com/cloudquery/cloudquery/issues/10071)) ([b0b9b05](https://github.com/cloudquery/cloudquery/commit/b0b9b05a3fdcd51ef2d672bc14755082eb8db617)), closes [#10039](https://github.com/cloudquery/cloudquery/issues/10039)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/plugin-sdk/v2 to v2.2.0 ([#10135](https://github.com/cloudquery/cloudquery/issues/10135)) ([cf33b89](https://github.com/cloudquery/cloudquery/commit/cf33b892ead0bb231e3956aa70967de552a21624))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v2 to v2.2.2 ([#10143](https://github.com/cloudquery/cloudquery/issues/10143)) ([8f887e0](https://github.com/cloudquery/cloudquery/commit/8f887e05de2096e8efd1e55863a8cf3c7620ccc3))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v2 to v2.3.0 ([#10163](https://github.com/cloudquery/cloudquery/issues/10163)) ([9a7f214](https://github.com/cloudquery/cloudquery/commit/9a7f21460772200e7a588409ebc7eb19f97b195b))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v2 to v2.3.1 ([#10175](https://github.com/cloudquery/cloudquery/issues/10175)) ([5b53423](https://github.com/cloudquery/cloudquery/commit/5b53423e72672f6c2bfb8ae00cfce1641410443e))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v2 to v2.3.3 ([#10187](https://github.com/cloudquery/cloudquery/issues/10187)) ([b185248](https://github.com/cloudquery/cloudquery/commit/b1852480b6ec8b721d94c72d8435051352f26932))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v2 to v2.3.4 ([#10196](https://github.com/cloudquery/cloudquery/issues/10196)) ([c6d2f59](https://github.com/cloudquery/cloudquery/commit/c6d2f59c7d77177a351cb82ecdc381dec6aad30c))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v2 to v2.3.5 ([#10200](https://github.com/cloudquery/cloudquery/issues/10200)) ([5a33693](https://github.com/cloudquery/cloudquery/commit/5a33693fe29f7068b03d80be1859d6e479c42c0d))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v2 to v2.3.6 ([#10208](https://github.com/cloudquery/cloudquery/issues/10208)) ([91c80a7](https://github.com/cloudquery/cloudquery/commit/91c80a795b46480447cfaef67c4db721a31e3206))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v2 to v2.3.8 ([#10213](https://github.com/cloudquery/cloudquery/issues/10213)) ([f358666](https://github.com/cloudquery/cloudquery/commit/f35866611cd206c37e6e9f9ad3329561e4cb32af))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v2 to v2.4.0 ([#10278](https://github.com/cloudquery/cloudquery/issues/10278)) ([a0a713e](https://github.com/cloudquery/cloudquery/commit/a0a713e8490b970b9d8bfaa1b50e01f43ff51c36))

## [1.0.8](https://github.com/cloudquery/cloudquery/compare/plugins-source-postgresql-v1.0.7...plugins-source-postgresql-v1.0.8) (2023-04-13)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.45.0 ([#9863](https://github.com/cloudquery/cloudquery/issues/9863)) ([2799d62](https://github.com/cloudquery/cloudquery/commit/2799d62518283ac304beecda9478f8f2db43cdc5))

## [1.0.7](https://github.com/cloudquery/cloudquery/compare/plugins-source-postgresql-v1.0.6...plugins-source-postgresql-v1.0.7) (2023-04-04)


### Bug Fixes

* **deps:** Update github.com/jackc/pglogrepl digest to 5ef673a ([#9564](https://github.com/cloudquery/cloudquery/issues/9564)) ([f3c5db1](https://github.com/cloudquery/cloudquery/commit/f3c5db14dda45d70d2889f849a8c98d4d5411b32))
* **deps:** Update github.com/jackc/pgx-zerolog digest to f978528 ([#9565](https://github.com/cloudquery/cloudquery/issues/9565)) ([d226511](https://github.com/cloudquery/cloudquery/commit/d2265111364347db93ef2c23b8c7f2d238282bd3))
* **deps:** Update golang.org/x/exp digest to 10a5072 ([#9587](https://github.com/cloudquery/cloudquery/issues/9587)) ([31f913f](https://github.com/cloudquery/cloudquery/commit/31f913f8e3538a2ba41b089bb11eae78aaf42ab2))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.44.1 ([#9520](https://github.com/cloudquery/cloudquery/issues/9520)) ([202c31b](https://github.com/cloudquery/cloudquery/commit/202c31b2788c3df35b5df7d07fdc750f92e7bb23))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.44.2 ([#9661](https://github.com/cloudquery/cloudquery/issues/9661)) ([a27dc84](https://github.com/cloudquery/cloudquery/commit/a27dc84a9b67b68b5b75b04dd3afe13e2c556082))
* **deps:** Update module github.com/mattn/go-isatty to v0.0.18 ([#9609](https://github.com/cloudquery/cloudquery/issues/9609)) ([5b2908e](https://github.com/cloudquery/cloudquery/commit/5b2908e8260c6e48f8c5fd6b8bd6c772f0c779d1))

## [1.0.6](https://github.com/cloudquery/cloudquery/compare/plugins-source-postgresql-v1.0.5...plugins-source-postgresql-v1.0.6) (2023-03-21)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.44.0 ([#9167](https://github.com/cloudquery/cloudquery/issues/9167)) ([49d6477](https://github.com/cloudquery/cloudquery/commit/49d647730a85ea6fae51e97194ba61c0625d1331))

## [1.0.5](https://github.com/cloudquery/cloudquery/compare/plugins-source-postgresql-v1.0.4...plugins-source-postgresql-v1.0.5) (2023-03-14)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.43.0 ([#8949](https://github.com/cloudquery/cloudquery/issues/8949)) ([31dfc63](https://github.com/cloudquery/cloudquery/commit/31dfc634850b699ba7bb7876399270a7367d6c7e))
* **source-postgres:** Support more Postgres and CockroachDB types ([#8926](https://github.com/cloudquery/cloudquery/issues/8926)) ([9c52479](https://github.com/cloudquery/cloudquery/commit/9c524797463b642949878a0ac9c3f549822b9ca0))

## [1.0.4](https://github.com/cloudquery/cloudquery/compare/plugins-source-postgresql-v1.0.3...plugins-source-postgresql-v1.0.4) (2023-03-08)


### Bug Fixes

* Respect tables and skip_tables config in PostgreSQL source plugin ([#8791](https://github.com/cloudquery/cloudquery/issues/8791)) ([e546566](https://github.com/cloudquery/cloudquery/commit/e5465663ded7f10b6b4ab50a7764d48851675a2b))

## [1.0.3](https://github.com/cloudquery/cloudquery/compare/plugins-source-postgresql-v1.0.2...plugins-source-postgresql-v1.0.3) (2023-03-07)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.41.0 ([#8682](https://github.com/cloudquery/cloudquery/issues/8682)) ([ea9d065](https://github.com/cloudquery/cloudquery/commit/ea9d065ae9f77c6dd990570974630ae6ac3f153e))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.42.0 ([#8725](https://github.com/cloudquery/cloudquery/issues/8725)) ([b83b277](https://github.com/cloudquery/cloudquery/commit/b83b277a2421d1caf46a26c3229041b27a3da148))
* **deps:** Update module github.com/jackc/pgx/v5 to v5.3.1 ([#8572](https://github.com/cloudquery/cloudquery/issues/8572)) ([11f2e3a](https://github.com/cloudquery/cloudquery/commit/11f2e3a8a5ede0524f55606d080838bd4b3f556c))
* **deps:** Update module github.com/stretchr/testify to v1.8.2 ([#8599](https://github.com/cloudquery/cloudquery/issues/8599)) ([2ec8086](https://github.com/cloudquery/cloudquery/commit/2ec808677328410cc96c97a693ef65022d314c32))

## [1.0.2](https://github.com/cloudquery/cloudquery/compare/plugins-source-postgresql-v1.0.1...plugins-source-postgresql-v1.0.2) (2023-02-28)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.39.0 ([#8344](https://github.com/cloudquery/cloudquery/issues/8344)) ([9c57544](https://github.com/cloudquery/cloudquery/commit/9c57544d06f9a774adcc659bcabd2518a905bdaa))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.39.1 ([#8371](https://github.com/cloudquery/cloudquery/issues/8371)) ([e3274c1](https://github.com/cloudquery/cloudquery/commit/e3274c109739bc107387627d340a713470c3a3c1))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.40.0 ([#8401](https://github.com/cloudquery/cloudquery/issues/8401)) ([4cf36d6](https://github.com/cloudquery/cloudquery/commit/4cf36d68684f37c0407332930766c1ba60807a93))

## [1.0.1](https://github.com/cloudquery/cloudquery/compare/plugins-source-postgresql-v1.0.0...plugins-source-postgresql-v1.0.1) (2023-02-21)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.38.2 ([#8156](https://github.com/cloudquery/cloudquery/issues/8156)) ([ac2d2d7](https://github.com/cloudquery/cloudquery/commit/ac2d2d70d5c4bc45fb8734bd4deb8a1e36074f6d))

## 1.0.0 (2023-02-19)


### Features

* **source-pg:** New Postgres source plugin with CDC ([#7975](https://github.com/cloudquery/cloudquery/issues/7975)) ([428917c](https://github.com/cloudquery/cloudquery/commit/428917c74ead3afe12cff9fb5f4830174829dcc2))
