# Changelog

## [2.0.0](https://github.com/cloudquery/cloudquery/compare/plugins-destination-bigquery-v1.3.2...plugins-destination-bigquery-v2.0.0) (2022-12-29)


### ⚠ BREAKING CHANGES

* **bigquery-spec:** Move `batch_size` from the plugin spec to the top level spec ([#6092](https://github.com/cloudquery/cloudquery/issues/6092))

### Bug Fixes

* **bigquery-spec:** Move `batch_size` from the plugin spec to the top level spec ([#6092](https://github.com/cloudquery/cloudquery/issues/6092)) ([a1d706e](https://github.com/cloudquery/cloudquery/commit/a1d706e839d3c81d5cb2cbb5971bdbe05a25288e))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.16.0 ([#6098](https://github.com/cloudquery/cloudquery/issues/6098)) ([7bacdf3](https://github.com/cloudquery/cloudquery/commit/7bacdf3364716eab08fa1a84ae4047b42edeee7e))

## [1.3.2](https://github.com/cloudquery/cloudquery/compare/plugins-destination-bigquery-v1.3.1...plugins-destination-bigquery-v1.3.2) (2022-12-28)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.15.0 ([#6071](https://github.com/cloudquery/cloudquery/issues/6071)) ([684b525](https://github.com/cloudquery/cloudquery/commit/684b525aaa285fcae70dd87af56679c1205adebe))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.15.1 ([#6079](https://github.com/cloudquery/cloudquery/issues/6079)) ([650659c](https://github.com/cloudquery/cloudquery/commit/650659c3c6766df571868e2ec3a2007cb76696eb))

## [1.3.1](https://github.com/cloudquery/cloudquery/compare/plugins-destination-bigquery-v1.3.0...plugins-destination-bigquery-v1.3.1) (2022-12-28)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.14.0 ([#6025](https://github.com/cloudquery/cloudquery/issues/6025)) ([35b2cfc](https://github.com/cloudquery/cloudquery/commit/35b2cfc7fc7bcdaceb7ee674e3a17f0f5673b366))

## [1.3.0](https://github.com/cloudquery/cloudquery/compare/plugins-destination-bigquery-v1.2.1...plugins-destination-bigquery-v1.3.0) (2022-12-23)


### Features

* **destinations:** Migrate to managed batching SDK ([#5805](https://github.com/cloudquery/cloudquery/issues/5805)) ([2f130c1](https://github.com/cloudquery/cloudquery/commit/2f130c12c6e83ccd8a2d036ab5c47b55e2fb5280))

## [1.2.1](https://github.com/cloudquery/cloudquery/compare/plugins-destination-bigquery-v1.2.0...plugins-destination-bigquery-v1.2.1) (2022-12-20)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.12.6 ([#5790](https://github.com/cloudquery/cloudquery/issues/5790)) ([8e2663c](https://github.com/cloudquery/cloudquery/commit/8e2663c17c3347afd5e53f665462adc3e709c96c))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.12.7 ([#5797](https://github.com/cloudquery/cloudquery/issues/5797)) ([15da529](https://github.com/cloudquery/cloudquery/commit/15da5294786fa2656228ca5bbc48ef1fc44e486b))

## [1.2.0](https://github.com/cloudquery/cloudquery/compare/plugins-destination-bigquery-v1.1.2...plugins-destination-bigquery-v1.2.0) (2022-12-16)


### Features

* **bigquery:** Add configurable batch size ([#5681](https://github.com/cloudquery/cloudquery/issues/5681)) ([17e110a](https://github.com/cloudquery/cloudquery/commit/17e110a54c2842eb4f89188f6e97b11624c3f5f6))


### Bug Fixes

* **bigquery:** BQ batch size ([#5684](https://github.com/cloudquery/cloudquery/issues/5684)) ([5288b47](https://github.com/cloudquery/cloudquery/commit/5288b47c71ca197cba773dc9adc5409a865aa19f))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.12.2 ([#5583](https://github.com/cloudquery/cloudquery/issues/5583)) ([d721c4e](https://github.com/cloudquery/cloudquery/commit/d721c4e06b8a97b5373215aca0e4ed64942ac489))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.12.3 ([#5639](https://github.com/cloudquery/cloudquery/issues/5639)) ([6452d0e](https://github.com/cloudquery/cloudquery/commit/6452d0ed5a44abad9d7530af6e79cde6504d0c4c))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.12.4 ([#5649](https://github.com/cloudquery/cloudquery/issues/5649)) ([b4aa889](https://github.com/cloudquery/cloudquery/commit/b4aa889e396db3b0887d1684e4bc07da6050af43))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.12.5 ([#5661](https://github.com/cloudquery/cloudquery/issues/5661)) ([b354b8a](https://github.com/cloudquery/cloudquery/commit/b354b8a3683fa2bc918c1002afac487427d65a5f))

## [1.1.2](https://github.com/cloudquery/cloudquery/compare/plugins-destination-bigquery-v1.1.1...plugins-destination-bigquery-v1.1.2) (2022-12-13)


### Bug Fixes

* BigQuery destination: wait for table migrations to complete ([#5544](https://github.com/cloudquery/cloudquery/issues/5544)) ([712ee39](https://github.com/cloudquery/cloudquery/commit/712ee399368a629e0d86809d117e081321567480))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.12.0 ([#5539](https://github.com/cloudquery/cloudquery/issues/5539)) ([fb71293](https://github.com/cloudquery/cloudquery/commit/fb71293d5cfe1b2ef32ba83d604ac3c48e662bce))
* Fix deadlock in BigQuery destination if error occurs during write ([#5550](https://github.com/cloudquery/cloudquery/issues/5550)) ([e087095](https://github.com/cloudquery/cloudquery/commit/e087095f3f4c32401ca4b7d18e8599b6f589924f))

## [1.1.1](https://github.com/cloudquery/cloudquery/compare/plugins-destination-bigquery-v1.1.0...plugins-destination-bigquery-v1.1.1) (2022-12-09)


### Bug Fixes

* Fix BigQuery JSON credential loading and allow users to set dataset location in config ([#5527](https://github.com/cloudquery/cloudquery/issues/5527)) ([446e486](https://github.com/cloudquery/cloudquery/commit/446e48648ea7d6e4136e10bec8e260a5efa76c0b))

## [1.1.0](https://github.com/cloudquery/cloudquery/compare/plugins-destination-bigquery-v1.0.1...plugins-destination-bigquery-v1.1.0) (2022-12-08)


### Features

* Support service account keys in BigQuery destination ([#5508](https://github.com/cloudquery/cloudquery/issues/5508)) ([5059ecc](https://github.com/cloudquery/cloudquery/commit/5059ecca76b1e186024ac4d582c2ef83f7c12e51))


### Bug Fixes

* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.11.1 ([#5458](https://github.com/cloudquery/cloudquery/issues/5458)) ([58b7432](https://github.com/cloudquery/cloudquery/commit/58b74321cd253c9a843c8c103f324abb93952195))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.11.2 ([#5497](https://github.com/cloudquery/cloudquery/issues/5497)) ([c1876cf](https://github.com/cloudquery/cloudquery/commit/c1876cf793b43d825a25fb3c9ba4996e4b09964f))

## [1.0.1](https://github.com/cloudquery/cloudquery/compare/plugins-destination-bigquery-v1.0.0...plugins-destination-bigquery-v1.0.1) (2022-12-06)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.11.0 ([#5416](https://github.com/cloudquery/cloudquery/issues/5416)) ([2e7ca35](https://github.com/cloudquery/cloudquery/commit/2e7ca35922fdb14fd717f582aaaa9693dae2ef4c))

## 1.0.0 (2022-12-02)


### Features

* BigQuery destination plugin ([#5102](https://github.com/cloudquery/cloudquery/issues/5102)) ([865fe06](https://github.com/cloudquery/cloudquery/commit/865fe067a769f98eb6b4bcb651541dde4ec39851))
