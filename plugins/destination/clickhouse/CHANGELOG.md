# Changelog

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
