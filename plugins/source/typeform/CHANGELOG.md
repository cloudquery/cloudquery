# Changelog

## [1.0.3](https://github.com/cloudquery/cloudquery/compare/plugins-source-typeform-v1.0.2...plugins-source-typeform-v1.0.3) (2023-09-27)


### Bug Fixes

* Tidy go.mod ([#14061](https://github.com/cloudquery/cloudquery/issues/14061)) ([11bd971](https://github.com/cloudquery/cloudquery/commit/11bd971f6a0089c92e47af6be24f552b2d920f21))

## [1.0.2](https://github.com/cloudquery/cloudquery/compare/plugins-source-typeform-v1.0.1...plugins-source-typeform-v1.0.2) (2023-09-05)


### Bug Fixes

* **deps:** Update dependency cloudquery-plugin-sdk to v0.1.2 ([#13571](https://github.com/cloudquery/cloudquery/issues/13571)) ([de71388](https://github.com/cloudquery/cloudquery/commit/de713889f9ccdbb963839d37edc122ff0ca7518e))
* **deps:** Update dependency cloudquery-plugin-sdk to v0.1.3 ([#13583](https://github.com/cloudquery/cloudquery/issues/13583)) ([ca673c1](https://github.com/cloudquery/cloudquery/commit/ca673c16ffa38eaab303f9502823696c85cd4d61))
* **deps:** Update github.com/99designs/go-keychain digest to 9cf53c8 ([#13561](https://github.com/cloudquery/cloudquery/issues/13561)) ([a170256](https://github.com/cloudquery/cloudquery/commit/a17025657e92b017fe3c8bd37abfaa2354e6e818))

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
