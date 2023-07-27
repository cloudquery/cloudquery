# Changelog

## [3.3.2](https://github.com/cloudquery/cloudquery/compare/plugins-destination-clickhouse-v3.3.1...plugins-destination-clickhouse-v3.3.2) (2023-07-25)


### Bug Fixes

* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to 10df4b9 ([#12443](https://github.com/cloudquery/cloudquery/issues/12443)) ([e385283](https://github.com/cloudquery/cloudquery/commit/e38528309f862f37bc7e278f9b69cf92d5aa5bd5))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v4 to v4.2.3 ([#12307](https://github.com/cloudquery/cloudquery/issues/12307)) ([8f14e4d](https://github.com/cloudquery/cloudquery/commit/8f14e4de7bf4d4c833f501135ea0610916a42f8b))
* **migration:** Make it clear that migration can be done manually and not only via `migrate_mode: forced` ([#12390](https://github.com/cloudquery/cloudquery/issues/12390)) ([33d39cf](https://github.com/cloudquery/cloudquery/commit/33d39cfa87243660241518e23fd5d845ce56d9da))

## [3.3.1](https://github.com/cloudquery/cloudquery/compare/plugins-destination-clickhouse-v3.3.0...plugins-destination-clickhouse-v3.3.1) (2023-07-18)


### Bug Fixes

* **clickhouse:** Update to SDK v4.2.1 ([#12155](https://github.com/cloudquery/cloudquery/issues/12155)) ([bde9102](https://github.com/cloudquery/cloudquery/commit/bde91021801587138918d137d715558f0a3a2d5f))
* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to 8e2219b ([#12220](https://github.com/cloudquery/cloudquery/issues/12220)) ([24e8fb5](https://github.com/cloudquery/cloudquery/commit/24e8fb588740896fe11a660e8b80231e741b753c))

## [3.3.0](https://github.com/cloudquery/cloudquery/compare/plugins-destination-clickhouse-v3.2.1...plugins-destination-clickhouse-v3.3.0) (2023-07-14)


### Features

* **clickhouse:** Migrate to SDK v4 ([#11794](https://github.com/cloudquery/cloudquery/issues/11794)) ([4bb2d8a](https://github.com/cloudquery/cloudquery/commit/4bb2d8a5327b0e0bc521770e580d28a5f6121d6f)), closes [#11747](https://github.com/cloudquery/cloudquery/issues/11747)
* Use chunks while preparing batch ([#11658](https://github.com/cloudquery/cloudquery/issues/11658)) ([8294e7d](https://github.com/cloudquery/cloudquery/commit/8294e7d28afe6b10ff208757244be3adab30e6fb))


### Bug Fixes

* **deps:** Update github.com/apache/arrow/go/v13 digest to 5a06b2e ([#11857](https://github.com/cloudquery/cloudquery/issues/11857)) ([43c2f5f](https://github.com/cloudquery/cloudquery/commit/43c2f5f3a893e5286f67c4943a9d1bc2736e2aeb))
* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to 0656028 ([#11739](https://github.com/cloudquery/cloudquery/issues/11739)) ([7a6ad49](https://github.com/cloudquery/cloudquery/commit/7a6ad49e8402d51e914d6fdc444956c89db91ad3))
* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to 0a52533 ([#12091](https://github.com/cloudquery/cloudquery/issues/12091)) ([927cefa](https://github.com/cloudquery/cloudquery/commit/927cefa943ec3969a2ec39b628bc1eba545a2108))
* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to 1e68c51 ([#11637](https://github.com/cloudquery/cloudquery/issues/11637)) ([46043bc](https://github.com/cloudquery/cloudquery/commit/46043bce410f86ba42390a6b190f9232fc2a1ded))
* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to 43638cb ([#11672](https://github.com/cloudquery/cloudquery/issues/11672)) ([3c60bbb](https://github.com/cloudquery/cloudquery/commit/3c60bbbb0233b17f934583766938780745145864))
* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to 4d76231 ([#11532](https://github.com/cloudquery/cloudquery/issues/11532)) ([6f04233](https://github.com/cloudquery/cloudquery/commit/6f042333acbd2506f7800ccb89a8c5cbfb7ad8d4))
* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to 8366a22 ([#11717](https://github.com/cloudquery/cloudquery/issues/11717)) ([8eeff5b](https://github.com/cloudquery/cloudquery/commit/8eeff5b17486d72845f830b99983f950fee7f5a0))
* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to 95d3199 ([#11708](https://github.com/cloudquery/cloudquery/issues/11708)) ([03f214f](https://github.com/cloudquery/cloudquery/commit/03f214f3dfd719b74ce9eb698ba255a8cf7528c7))
* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to a2a76eb ([#12104](https://github.com/cloudquery/cloudquery/issues/12104)) ([311f474](https://github.com/cloudquery/cloudquery/commit/311f4749af2491a606f29483190717a5fe238da6))
* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to b0832be ([#11651](https://github.com/cloudquery/cloudquery/issues/11651)) ([71e8c29](https://github.com/cloudquery/cloudquery/commit/71e8c29624494a3e1cd104e46266a610ce57c83c))
* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to d864719 ([#11611](https://github.com/cloudquery/cloudquery/issues/11611)) ([557a290](https://github.com/cloudquery/cloudquery/commit/557a2903af272b8e2e4c9eebb36e39cd8a41a805))
* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to df3b664 ([#11882](https://github.com/cloudquery/cloudquery/issues/11882)) ([9635b22](https://github.com/cloudquery/cloudquery/commit/9635b22b10a2cd9ca0f91819cffb7f4ba75dc2d9))
* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to f060192 ([#11730](https://github.com/cloudquery/cloudquery/issues/11730)) ([c7019c2](https://github.com/cloudquery/cloudquery/commit/c7019c26c311f29b66c90fc5d461a0daf71d191c))
* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to f0dffc6 ([#11689](https://github.com/cloudquery/cloudquery/issues/11689)) ([18ac0e9](https://github.com/cloudquery/cloudquery/commit/18ac0e9dbef31d06701f1f13d263ad840ac60c5e))
* **deps:** Update module github.com/cloudquery/plugin-pb-go to v1.1.0 ([#11665](https://github.com/cloudquery/cloudquery/issues/11665)) ([d8947c9](https://github.com/cloudquery/cloudquery/commit/d8947c9efa6ab8bf3952ad9d929e8ed81f2dea55))
* **deps:** Update module github.com/cloudquery/plugin-pb-go to v1.2.0 ([#11720](https://github.com/cloudquery/cloudquery/issues/11720)) ([7ef521d](https://github.com/cloudquery/cloudquery/commit/7ef521db1423c6f0de197b08c73adf22c896f999))
* **deps:** Update module github.com/cloudquery/plugin-pb-go to v1.2.1 ([#11722](https://github.com/cloudquery/cloudquery/issues/11722)) ([309be72](https://github.com/cloudquery/cloudquery/commit/309be7276d7de157013c281b6fb3934513898b3f))
* **deps:** Update module github.com/cloudquery/plugin-pb-go to v1.3.3 ([#11726](https://github.com/cloudquery/cloudquery/issues/11726)) ([f0ca611](https://github.com/cloudquery/cloudquery/commit/f0ca61195014bde707761a15efa27a92955b59db))
* **deps:** Update module github.com/cloudquery/plugin-pb-go to v1.3.4 ([#11753](https://github.com/cloudquery/cloudquery/issues/11753)) ([cd4fe1c](https://github.com/cloudquery/cloudquery/commit/cd4fe1c54f85f8511252bebd5671361618ddb0d3))
* **deps:** Update module github.com/cloudquery/plugin-pb-go to v1.5.0 ([#11850](https://github.com/cloudquery/cloudquery/issues/11850)) ([3255857](https://github.com/cloudquery/cloudquery/commit/3255857938bf16862d52491f5c2a8a0fa53faef0))
* **deps:** Update module github.com/cloudquery/plugin-pb-go to v1.7.0 ([#12166](https://github.com/cloudquery/cloudquery/issues/12166)) ([94390dd](https://github.com/cloudquery/cloudquery/commit/94390dde19d0c37fee9d035219d62f6ae7edb127))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v3 to v3.10.6 ([#11473](https://github.com/cloudquery/cloudquery/issues/11473)) ([7272133](https://github.com/cloudquery/cloudquery/commit/72721336632e127dd37de4541f2f503bf4f73fb6))
* Use `configtype.Duration` ([#11940](https://github.com/cloudquery/cloudquery/issues/11940)) ([9876bcb](https://github.com/cloudquery/cloudquery/commit/9876bcb0adb3b0ad5c664783d842f729a09c30c4))

## [3.2.1](https://github.com/cloudquery/cloudquery/compare/plugins-destination-clickhouse-v3.2.0...plugins-destination-clickhouse-v3.2.1) (2023-06-13)


### Bug Fixes

* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to 0f7bd3b ([#11412](https://github.com/cloudquery/cloudquery/issues/11412)) ([dd1e2e8](https://github.com/cloudquery/cloudquery/commit/dd1e2e892d95515fd7332339262abaefd2a256c5))
* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to 7f6aaff ([#11432](https://github.com/cloudquery/cloudquery/issues/11432)) ([55dfebc](https://github.com/cloudquery/cloudquery/commit/55dfebc064608fb47caaf3b8e68c8002de8a7dc3))
* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to 8f72077 ([#11395](https://github.com/cloudquery/cloudquery/issues/11395)) ([d91fc5c](https://github.com/cloudquery/cloudquery/commit/d91fc5ce24f64c29fff6988b19ec2c2775cc379b))
* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to 90670b8 ([#11279](https://github.com/cloudquery/cloudquery/issues/11279)) ([a6cdc91](https://github.com/cloudquery/cloudquery/commit/a6cdc912e4b38a3faf798c5147a986ffe2539643))
* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to b359e74 ([#11405](https://github.com/cloudquery/cloudquery/issues/11405)) ([5d92765](https://github.com/cloudquery/cloudquery/commit/5d927659bd4f7c445a0e312487f1655ffb9a60f6))
* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to d8eacf8 ([#11449](https://github.com/cloudquery/cloudquery/issues/11449)) ([742dafd](https://github.com/cloudquery/cloudquery/commit/742dafd5bf5cdc8facb94fda5de1d84c88897cbd))
* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to e258cfb ([#11391](https://github.com/cloudquery/cloudquery/issues/11391)) ([eacbe9a](https://github.com/cloudquery/cloudquery/commit/eacbe9ad3ea16d88f27c4593fa2774574ac8fe4e))

## [3.2.0](https://github.com/cloudquery/cloudquery/compare/plugins-destination-clickhouse-v3.1.1...plugins-destination-clickhouse-v3.2.0) (2023-06-06)


### Features

* **nested:** Simplify constructing nested ClickHouse value ([#11205](https://github.com/cloudquery/cloudquery/issues/11205)) ([dd6b275](https://github.com/cloudquery/cloudquery/commit/dd6b2751c723714eeb69c2b152266dfbc07152d8))


### Bug Fixes

* **deps:** Update `github.com/cloudquery/plugin-sdk/v3` to `v3.8.1` ([#11078](https://github.com/cloudquery/cloudquery/issues/11078)) ([0a1764f](https://github.com/cloudquery/cloudquery/commit/0a1764ff0cceb650d914ad5cb0a34e66d33baf3d))
* **deps:** Update github.com/apache/arrow/go/v13 digest to e07e22c ([#11151](https://github.com/cloudquery/cloudquery/issues/11151)) ([5083cf7](https://github.com/cloudquery/cloudquery/commit/5083cf720f0ae98e07448ba2ae1116048e2d3a90))
* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to 20b0de9 ([#11199](https://github.com/cloudquery/cloudquery/issues/11199)) ([dc3565d](https://github.com/cloudquery/cloudquery/commit/dc3565d3fd6a640d9d10b4fd3a7fe6009a9d02a5))
* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to 88d5dc2 ([#11226](https://github.com/cloudquery/cloudquery/issues/11226)) ([9f306bc](https://github.com/cloudquery/cloudquery/commit/9f306bcaf3833b4611f0df5c50277be43aa19cbb))
* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to a7aad4c ([#11184](https://github.com/cloudquery/cloudquery/issues/11184)) ([8a0822e](https://github.com/cloudquery/cloudquery/commit/8a0822e31fc0eef99de2cdd2bd6d7e4c8b4131bf))
* **deps:** Update github.com/cloudquery/arrow/go/v13 digest to c67fb39 ([#11169](https://github.com/cloudquery/cloudquery/issues/11169)) ([dcb0f92](https://github.com/cloudquery/cloudquery/commit/dcb0f9296a770a5cc2eb6bffd6b1ee30fbccb5dc))
* **deps:** Update golang.org/x/exp digest to 2e198f4 ([#11155](https://github.com/cloudquery/cloudquery/issues/11155)) ([c46c62b](https://github.com/cloudquery/cloudquery/commit/c46c62b68692f527485d7f4b84265abc5dc1142c))
* **deps:** Update google.golang.org/genproto digest to e85fd2c ([#11156](https://github.com/cloudquery/cloudquery/issues/11156)) ([dbe7e92](https://github.com/cloudquery/cloudquery/commit/dbe7e9293d693a6821570e0e0b80202a936b6d3c))
* **deps:** Update module github.com/cloudquery/plugin-pb-go to v1.0.9 ([#11240](https://github.com/cloudquery/cloudquery/issues/11240)) ([f92cd4b](https://github.com/cloudquery/cloudquery/commit/f92cd4bfe3c3d0088964d52ab9cd01ca4cf622e1))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v3 to v3.10.4 ([#11244](https://github.com/cloudquery/cloudquery/issues/11244)) ([8fceef6](https://github.com/cloudquery/cloudquery/commit/8fceef6f9041e173923555d8ff221cfe83b424c2))

## [3.1.1](https://github.com/cloudquery/cloudquery/compare/plugins-destination-clickhouse-v3.1.0...plugins-destination-clickhouse-v3.1.1) (2023-05-18)


### Bug Fixes

* Upgrade to plugin-sdk v3.5.2 (Fixes delete-stale for incremental tables) ([#10854](https://github.com/cloudquery/cloudquery/issues/10854)) ([96c17c7](https://github.com/cloudquery/cloudquery/commit/96c17c7ea4a2a455cc3fa52728d818c87e0cff33))

## [3.1.0](https://github.com/cloudquery/cloudquery/compare/plugins-destination-clickhouse-v3.0.0...plugins-destination-clickhouse-v3.1.0) (2023-05-18)


### Features

* **clickhouse:** Migrate to `github.com/cloudquery/plugin-sdk/v3` ([#10807](https://github.com/cloudquery/cloudquery/issues/10807)) ([e0edde7](https://github.com/cloudquery/cloudquery/commit/e0edde7e86956c5fce12cf47bbbe9394d4043652)), closes [#10714](https://github.com/cloudquery/cloudquery/issues/10714)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/plugin-pb-go to v1.0.8 ([#10798](https://github.com/cloudquery/cloudquery/issues/10798)) ([27ff430](https://github.com/cloudquery/cloudquery/commit/27ff430527932d59a4d488a6767547eda8853940))

## [3.0.0](https://github.com/cloudquery/cloudquery/compare/plugins-destination-clickhouse-v2.0.1...plugins-destination-clickhouse-v3.0.0) (2023-05-09)


### ⚠ BREAKING CHANGES

* This change enables [`allow_nullable_key`](https://clickhouse.com/docs/en/operations/settings/settings#allow-nullable-key) for tables ([#10284](https://github.com/cloudquery/cloudquery/issues/10284)).

### Features

* Allow nullable columns in primary keys ([#10284](https://github.com/cloudquery/cloudquery/issues/10284)) ([26e7d0e](https://github.com/cloudquery/cloudquery/commit/26e7d0e9deb5eb2709e3df605a008e8dba4d6552))
* **deps:** Upgrade to Apache Arrow v13 (latest `cqmain`) ([#10605](https://github.com/cloudquery/cloudquery/issues/10605)) ([a55da3d](https://github.com/cloudquery/cloudquery/commit/a55da3dbefafdc68a6bda2d5f1d334d12dd97b97))
* Update to use [Apache Arrow](https://arrow.apache.org/) type system ([#10284](https://github.com/cloudquery/cloudquery/issues/10284)) ([26e7d0e](https://github.com/cloudquery/cloudquery/commit/26e7d0e9deb5eb2709e3df605a008e8dba4d6552))

## [2.0.1](https://github.com/cloudquery/cloudquery/compare/plugins-destination-clickhouse-v2.0.0...plugins-destination-clickhouse-v2.0.1) (2023-04-25)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.45.0 ([#9863](https://github.com/cloudquery/cloudquery/issues/9863)) ([2799d62](https://github.com/cloudquery/cloudquery/commit/2799d62518283ac304beecda9478f8f2db43cdc5))

## [2.0.0](https://github.com/cloudquery/cloudquery/compare/plugins-destination-clickhouse-v1.3.3...plugins-destination-clickhouse-v2.0.0) (2023-04-04)


### ⚠ BREAKING CHANGES

* **clickhouse:** Stop reading `ca_cert` value as file path. See [file variable substitution](/docs/advanced-topics/environment-variable-substitution#file-variable-substitution-example) for how to read this value from a file.

### Features

* **clickhouse:** Read only plain `ca_cert` value ([#9495](https://github.com/cloudquery/cloudquery/issues/9495)) ([dcffd50](https://github.com/cloudquery/cloudquery/commit/dcffd506b847ec3634c05fdd4e841764f3434b91))


### Bug Fixes

* **clickhouse:** Check `ca_cert` append result ([#9505](https://github.com/cloudquery/cloudquery/issues/9505)) ([eea1b11](https://github.com/cloudquery/cloudquery/commit/eea1b11151560be38c5413e839c372d5c6eb64a4))
* **deps:** Update golang.org/x/exp digest to 10a5072 ([#9587](https://github.com/cloudquery/cloudquery/issues/9587)) ([31f913f](https://github.com/cloudquery/cloudquery/commit/31f913f8e3538a2ba41b089bb11eae78aaf42ab2))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.44.1 ([#9520](https://github.com/cloudquery/cloudquery/issues/9520)) ([202c31b](https://github.com/cloudquery/cloudquery/commit/202c31b2788c3df35b5df7d07fdc750f92e7bb23))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.44.2 ([#9661](https://github.com/cloudquery/cloudquery/issues/9661)) ([a27dc84](https://github.com/cloudquery/cloudquery/commit/a27dc84a9b67b68b5b75b04dd3afe13e2c556082))
* **deps:** Update module github.com/mattn/go-isatty to v0.0.18 ([#9609](https://github.com/cloudquery/cloudquery/issues/9609)) ([5b2908e](https://github.com/cloudquery/cloudquery/commit/5b2908e8260c6e48f8c5fd6b8bd6c772f0c779d1))
* Ignore primary key option when migrating tables ([3a0c68b](https://github.com/cloudquery/cloudquery/commit/3a0c68b59b8b15b3b7b7fa3bb7584b0ad9c5782b))

## [1.3.3](https://github.com/cloudquery/cloudquery/compare/plugins-destination-clickhouse-v1.3.2...plugins-destination-clickhouse-v1.3.3) (2023-03-23)


### Bug Fixes

* **clickhouse:** Add `ON CLUSTER` for `DROP TABLE` statement ([#9377](https://github.com/cloudquery/cloudquery/issues/9377)) ([76a74ff](https://github.com/cloudquery/cloudquery/commit/76a74ffb3479bc7c086b33020d665a28bdf75db5))

## [1.3.2](https://github.com/cloudquery/cloudquery/compare/plugins-destination-clickhouse-v1.3.1...plugins-destination-clickhouse-v1.3.2) (2023-03-21)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.44.0 ([#9167](https://github.com/cloudquery/cloudquery/issues/9167)) ([49d6477](https://github.com/cloudquery/cloudquery/commit/49d647730a85ea6fae51e97194ba61c0625d1331))

## [1.3.1](https://github.com/cloudquery/cloudquery/compare/plugins-destination-clickhouse-v1.3.0...plugins-destination-clickhouse-v1.3.1) (2023-03-14)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.43.0 ([#8949](https://github.com/cloudquery/cloudquery/issues/8949)) ([31dfc63](https://github.com/cloudquery/cloudquery/commit/31dfc634850b699ba7bb7876399270a7367d6c7e))

## [1.3.0](https://github.com/cloudquery/cloudquery/compare/plugins-destination-clickhouse-v1.2.1...plugins-destination-clickhouse-v1.3.0) (2023-03-10)


### Features

* **clickhouse:** Add table engine option ([#8844](https://github.com/cloudquery/cloudquery/issues/8844)) ([447b29c](https://github.com/cloudquery/cloudquery/commit/447b29c172129a2c4a24fb81053a54fb9c6d8103)), closes [#8667](https://github.com/cloudquery/cloudquery/issues/8667)

## [1.2.1](https://github.com/cloudquery/cloudquery/compare/plugins-destination-clickhouse-v1.2.0...plugins-destination-clickhouse-v1.2.1) (2023-03-07)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.41.0 ([#8682](https://github.com/cloudquery/cloudquery/issues/8682)) ([ea9d065](https://github.com/cloudquery/cloudquery/commit/ea9d065ae9f77c6dd990570974630ae6ac3f153e))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.42.0 ([#8725](https://github.com/cloudquery/cloudquery/issues/8725)) ([b83b277](https://github.com/cloudquery/cloudquery/commit/b83b277a2421d1caf46a26c3229041b27a3da148))

## [1.2.0](https://github.com/cloudquery/cloudquery/compare/plugins-destination-clickhouse-v1.1.3...plugins-destination-clickhouse-v1.2.0) (2023-03-01)


### Features

* **clickhouse:** Support distributed DDL ([#8663](https://github.com/cloudquery/cloudquery/issues/8663)) ([c46705f](https://github.com/cloudquery/cloudquery/commit/c46705f02cec99fd573ed1d1721921c58d1f4cab)), closes [#8654](https://github.com/cloudquery/cloudquery/issues/8654)


### Bug Fixes

* **deps:** Update golang.org/x/exp digest to c95f2b4 ([#8560](https://github.com/cloudquery/cloudquery/issues/8560)) ([9c3bd5b](https://github.com/cloudquery/cloudquery/commit/9c3bd5b68f9741a360fde6c54bf3f5f3efe06d9e))
* **deps:** Update module github.com/andybalholm/brotli to v1.0.5 ([#8570](https://github.com/cloudquery/cloudquery/issues/8570)) ([1251c4d](https://github.com/cloudquery/cloudquery/commit/1251c4dd228cee7d34af4e9ec8df1e9ccfb41e3e))
* **deps:** Update module github.com/ClickHouse/ch-go to v0.53.0 ([#8652](https://github.com/cloudquery/cloudquery/issues/8652)) ([a016609](https://github.com/cloudquery/cloudquery/commit/a0166095c8b57330d1ba292848f7df3c09728032))
* **deps:** Update module github.com/ClickHouse/clickhouse-go/v2 to v2.6.5 ([#8568](https://github.com/cloudquery/cloudquery/issues/8568)) ([d553b70](https://github.com/cloudquery/cloudquery/commit/d553b700a05bb0c0d8a59f74f454b0c46371a6b7))
* **deps:** Update module github.com/stretchr/testify to v1.8.2 ([#8599](https://github.com/cloudquery/cloudquery/issues/8599)) ([2ec8086](https://github.com/cloudquery/cloudquery/commit/2ec808677328410cc96c97a693ef65022d314c32))

## [1.1.3](https://github.com/cloudquery/cloudquery/compare/plugins-destination-clickhouse-v1.1.2...plugins-destination-clickhouse-v1.1.3) (2023-02-28)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.40.0 ([#8401](https://github.com/cloudquery/cloudquery/issues/8401)) ([4cf36d6](https://github.com/cloudquery/cloudquery/commit/4cf36d68684f37c0407332930766c1ba60807a93))

## [1.1.2](https://github.com/cloudquery/cloudquery/compare/plugins-destination-clickhouse-v1.1.1...plugins-destination-clickhouse-v1.1.2) (2023-02-23)


### Bug Fixes

* **clickhouse:** Bump minimum ClickHouse version ([#8406](https://github.com/cloudquery/cloudquery/issues/8406)) ([a5890b2](https://github.com/cloudquery/cloudquery/commit/a5890b2d304b06b3460612c52980ac60dfcf6058))

## [1.1.1](https://github.com/cloudquery/cloudquery/compare/plugins-destination-clickhouse-v1.1.0...plugins-destination-clickhouse-v1.1.1) (2023-02-23)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.39.0 ([#8344](https://github.com/cloudquery/cloudquery/issues/8344)) ([9c57544](https://github.com/cloudquery/cloudquery/commit/9c57544d06f9a774adcc659bcabd2518a905bdaa))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.39.1 ([#8371](https://github.com/cloudquery/cloudquery/issues/8371)) ([e3274c1](https://github.com/cloudquery/cloudquery/commit/e3274c109739bc107387627d340a713470c3a3c1))
* **migration:** Don't add trailing comma to create table ([#8402](https://github.com/cloudquery/cloudquery/issues/8402)) ([1182b21](https://github.com/cloudquery/cloudquery/commit/1182b2194cd87d9f61c35aac3acd5ba25ec352da))

## [1.1.0](https://github.com/cloudquery/cloudquery/compare/plugins-destination-clickhouse-v1.0.0...plugins-destination-clickhouse-v1.1.0) (2023-02-21)


### Features

* **deps:** Update ClickHouse plugin-sdk to v1.38.2 ([#8255](https://github.com/cloudquery/cloudquery/issues/8255)) ([ddbc004](https://github.com/cloudquery/cloudquery/commit/ddbc004eb7e65c756929afd84758d756aba7549b))


### Bug Fixes

* **deps:** Update module golang.org/x/net to v0.7.0 [SECURITY] ([#8176](https://github.com/cloudquery/cloudquery/issues/8176)) ([fc4cef8](https://github.com/cloudquery/cloudquery/commit/fc4cef86dce4ca76ca8397e897ab744e48975834))

## 1.0.0 (2023-02-14)


### Features

* **clickhouse:** ClickHouse destination ([#6982](https://github.com/cloudquery/cloudquery/issues/6982)) ([09411e4](https://github.com/cloudquery/cloudquery/commit/09411e45c01609b986f0cc6f42554096ae5558dc)), closes [#6254](https://github.com/cloudquery/cloudquery/issues/6254)
