# Changelog

## [2.0.0](https://github.com/cloudquery/cloudquery/compare/plugins-destination-meilisearch-v1.1.0...plugins-destination-meilisearch-v2.0.0) (2023-04-25)


### ⚠ BREAKING CHANGES

* This release introduces an internal change to our type system to use [Apache Arrow](https://arrow.apache.org/). This should not have any visible breaking changes, however due to the size of the change we are introducing it under a major version bump to communicate that it might have some bugs that we weren't able to catch during our internal tests. If you encounter an issue during the upgrade, please submit a [bug report](https://github.com/cloudquery/cloudquery/issues/new/choose).

### Features

* Update to use [Apache Arrow](https://arrow.apache.org/) type system ([859212b](https://github.com/cloudquery/cloudquery/commit/859212bd9a5908617d6ee26597049a18b6d136c1))


### Bug Fixes

* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.45.0 ([#9863](https://github.com/cloudquery/cloudquery/issues/9863)) ([2799d62](https://github.com/cloudquery/cloudquery/commit/2799d62518283ac304beecda9478f8f2db43cdc5))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v2 to v2.3.8 ([#10213](https://github.com/cloudquery/cloudquery/issues/10213)) ([f358666](https://github.com/cloudquery/cloudquery/commit/f35866611cd206c37e6e9f9ad3329561e4cb32af))
* **deps:** Update module github.com/cloudquery/plugin-sdk/v2 to v2.4.0 ([#10278](https://github.com/cloudquery/cloudquery/issues/10278)) ([a0a713e](https://github.com/cloudquery/cloudquery/commit/a0a713e8490b970b9d8bfaa1b50e01f43ff51c36))

## [1.1.0](https://github.com/cloudquery/cloudquery/compare/plugins-destination-meilisearch-v1.0.0...plugins-destination-meilisearch-v1.1.0) (2023-04-04)


### Features

* **meilisearch:** Use Meilisearch v1.1 ([#9640](https://github.com/cloudquery/cloudquery/issues/9640)) ([ff4edb0](https://github.com/cloudquery/cloudquery/commit/ff4edb0e827aab9704503cd21a7dc099a67ae381))


### Bug Fixes

* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.44.1 ([#9520](https://github.com/cloudquery/cloudquery/issues/9520)) ([202c31b](https://github.com/cloudquery/cloudquery/commit/202c31b2788c3df35b5df7d07fdc750f92e7bb23))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.44.2 ([#9661](https://github.com/cloudquery/cloudquery/issues/9661)) ([a27dc84](https://github.com/cloudquery/cloudquery/commit/a27dc84a9b67b68b5b75b04dd3afe13e2c556082))

## 1.0.0 (2023-03-29)


### Features

* **meilisearch:** Meilisearch destination plugin ([#9404](https://github.com/cloudquery/cloudquery/issues/9404)) ([40ccf7b](https://github.com/cloudquery/cloudquery/commit/40ccf7bc2febf4a4a1526c912f959481d97bdb6b)), closes [#8857](https://github.com/cloudquery/cloudquery/issues/8857)
