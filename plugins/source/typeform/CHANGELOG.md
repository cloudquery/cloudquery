# Changelog

## [1.0.1](https://github.com/cloudquery/cloudquery/compare/plugins-source-typeform-v1.0.0...plugins-source-typeform-v1.0.1) (2023-08-18)


### Bug Fixes

* Fix JSON logging ([#13185](https://github.com/cloudquery/cloudquery/issues/13185)) ([83e46a0](https://github.com/cloudquery/cloudquery/commit/83e46a070234c75b6d9b74bfad89b7cd370b5c76))

## [1.0.0](https://github.com/cloudquery/cloudquery/compare/plugins-source-typeform-v0.1.1...plugins-source-typeform-v1.0.0) (2023-08-16)


### ⚠ BREAKING CHANGES

* Upgrade SDK to v0.1.1 ([#13055](https://github.com/cloudquery/cloudquery/issues/13055)). This upgrades the Python SDK dependency to v0.1.1, and starts generating documentation using the cloudquery tables command instead of the (now-deprecated) plugin doc command.

*Note*: there are no breaking schema changes in this PR - only different aliases of the types are now used in the docs to be in line with those used by Go plugins. The breaking change refers to the removal of the plugin's doc command brought about by the upgrade of the SDK version. `cloudquery tables` can be used instead.

### Bug Fixes

* Upgrade SDK to v0.1.1 ([#13055](https://github.com/cloudquery/cloudquery/issues/13055)) ([1b841d8](https://github.com/cloudquery/cloudquery/commit/1b841d84637bef7b4707796292bb52bed7fa7a77))

## [0.1.1](https://github.com/cloudquery/cloudquery/compare/plugins-source-typeform-v0.1.0...plugins-source-typeform-v0.1.1) (2023-08-09)


### Bug Fixes

* **deps:** Update dependency cloudquery-plugin-sdk to v0.0.11 ([#12863](https://github.com/cloudquery/cloudquery/issues/12863)) ([d6f063d](https://github.com/cloudquery/cloudquery/commit/d6f063d67d65652a494d1bb9d28f6c5115f58a90))

## 0.1.0 (2023-08-03)


### This Release has the Following Changes to Tables
- Table `typeform_form_responses` was added
- Table `typeform_forms` was added

### Features

* Typeform plugin ([#12732](https://github.com/cloudquery/cloudquery/issues/12732)) ([112b5b5](https://github.com/cloudquery/cloudquery/commit/112b5b503f2787673e7c3b59f8b8c6e29d0b4c4e))
