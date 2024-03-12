# Changelog

## [1.3.3](https://github.com/cloudquery/cloudquery/compare/plugins-source-typeform-v1.3.2...plugins-source-typeform-v1.3.3) (2024-03-12)


### Bug Fixes

* **deps:** Update dependency cloudquery-plugin-sdk to v0.1.16 ([#17105](https://github.com/cloudquery/cloudquery/issues/17105)) ([cb24442](https://github.com/cloudquery/cloudquery/commit/cb24442f1fa14bedd1a5cb119b16807b60eac750))

## [1.3.2](https://github.com/cloudquery/cloudquery/compare/plugins-source-typeform-v1.3.1...plugins-source-typeform-v1.3.2) (2024-03-05)


### Bug Fixes

* **deps:** Update dependency cloudquery-plugin-sdk to v0.1.15 ([#16991](https://github.com/cloudquery/cloudquery/issues/16991)) ([9cbfb42](https://github.com/cloudquery/cloudquery/commit/9cbfb4209797a74b812fa9547e81f84076f1dc68))

## [1.3.1](https://github.com/cloudquery/cloudquery/compare/plugins-source-typeform-v1.3.0...plugins-source-typeform-v1.3.1) (2024-02-27)


### Bug Fixes

* Use env var replacement in example configuration ([#16848](https://github.com/cloudquery/cloudquery/issues/16848)) ([fc790d0](https://github.com/cloudquery/cloudquery/commit/fc790d0c8fab70e06a1082263329d0adb96dd1e3))

## [1.3.0](https://github.com/cloudquery/cloudquery/compare/plugins-source-typeform-v1.2.3...plugins-source-typeform-v1.3.0) (2024-02-13)


### Features

* Add JSON Schema to `typeform` source plugin ([#16590](https://github.com/cloudquery/cloudquery/issues/16590)) ([1af2d81](https://github.com/cloudquery/cloudquery/commit/1af2d81be54894ea76e62c7db991b6ac13e68949))


### Bug Fixes

* **deps:** Update dependency cloudquery-plugin-sdk to v0.1.14 ([#16578](https://github.com/cloudquery/cloudquery/issues/16578)) ([7fe2c3b](https://github.com/cloudquery/cloudquery/commit/7fe2c3b5b04f314dec35bb980c19feb53925265f))

## [1.2.3](https://github.com/cloudquery/cloudquery/compare/plugins-source-typeform-v1.2.2...plugins-source-typeform-v1.2.3) (2024-02-01)


### Bug Fixes

* **deps:** Update dependency cloudquery-plugin-sdk to v0.1.13 ([#16462](https://github.com/cloudquery/cloudquery/issues/16462)) ([d76eef1](https://github.com/cloudquery/cloudquery/commit/d76eef19cd2e5aa76de6101d2d5ff00f0e54cedc))

## [1.2.2](https://github.com/cloudquery/cloudquery/compare/plugins-source-typeform-v1.2.1...plugins-source-typeform-v1.2.2) (2024-01-15)


### Bug Fixes

* **deps:** Update dependency cloudquery-plugin-sdk to v0.1.12 ([#16007](https://github.com/cloudquery/cloudquery/issues/16007)) ([f337f94](https://github.com/cloudquery/cloudquery/commit/f337f944aeba52e1fba06cc4ebcd25cccafbf16f))

## [1.2.1](https://github.com/cloudquery/cloudquery/compare/plugins-source-typeform-v1.2.0...plugins-source-typeform-v1.2.1) (2024-01-02)


### Bug Fixes

* **deps:** Update dependency cloudquery-plugin-sdk to v0.1.11 ([#15940](https://github.com/cloudquery/cloudquery/issues/15940)) ([f6ae69a](https://github.com/cloudquery/cloudquery/commit/f6ae69a9719db6207da570863e41745b59c726d5))

## [1.2.0](https://github.com/cloudquery/cloudquery/compare/plugins-source-typeform-v1.1.2...plugins-source-typeform-v1.2.0) (2023-12-29)


### Features

* Add package command ([#15726](https://github.com/cloudquery/cloudquery/issues/15726)) ([2d36b9b](https://github.com/cloudquery/cloudquery/commit/2d36b9b4759966db95baa92ce03c76b4e4d1e5d2))


### Bug Fixes

* **deps:** Update dependency cloudquery-plugin-sdk to v0.1.10 ([#15857](https://github.com/cloudquery/cloudquery/issues/15857)) ([3ab55b8](https://github.com/cloudquery/cloudquery/commit/3ab55b8de769c4c7bf9b979a183dd23125994255))
* **deps:** Update dependency cloudquery-plugin-sdk to v0.1.9 ([#15720](https://github.com/cloudquery/cloudquery/issues/15720)) ([697b330](https://github.com/cloudquery/cloudquery/commit/697b330ac7d59435424791cbbb1b544e7870b463))

## [1.1.2](https://github.com/cloudquery/cloudquery/compare/plugins-source-typeform-v1.1.1...plugins-source-typeform-v1.1.2) (2023-11-23)


### Bug Fixes

* Add missing tables for Docker plugins, remove on-demand rendering of tables ([#15433](https://github.com/cloudquery/cloudquery/issues/15433)) ([2286412](https://github.com/cloudquery/cloudquery/commit/22864120467b98d9284c533e58a7c9e09f8790ab))

## [1.1.1](https://github.com/cloudquery/cloudquery/compare/plugins-source-typeform-v1.1.0...plugins-source-typeform-v1.1.1) (2023-11-16)


### Bug Fixes

* **deps:** Update dependency cloudquery-plugin-sdk to v0.1.8 ([#15213](https://github.com/cloudquery/cloudquery/issues/15213)) ([02c07e0](https://github.com/cloudquery/cloudquery/commit/02c07e0f67f41cb62f1da2e84305f1e28b823cb4))
* **deps:** Update dependency pyarrow to v14 [SECURITY] ([#15213](https://github.com/cloudquery/cloudquery/issues/15213)) ([02c07e0](https://github.com/cloudquery/cloudquery/commit/02c07e0f67f41cb62f1da2e84305f1e28b823cb4))
* **deps:** Update module github.com/docker/docker to v24 [SECURITY] ([#15060](https://github.com/cloudquery/cloudquery/issues/15060)) ([41acd0e](https://github.com/cloudquery/cloudquery/commit/41acd0e4ac63221e90cca89a7137a8685692267d))

## [1.1.0](https://github.com/cloudquery/cloudquery/compare/plugins-source-typeform-v1.0.7...plugins-source-typeform-v1.1.0) (2023-10-27)


### This Release has the Following Changes to Tables
- Table `typeform_form_responses`: column added with name `tags` and type `json`

### Features

* Add tags column to typeform plugin ([#15019](https://github.com/cloudquery/cloudquery/issues/15019)) ([17066b6](https://github.com/cloudquery/cloudquery/commit/17066b67fb4d914e9a2a7aa2088c915a772380a0))

## [1.0.7](https://github.com/cloudquery/cloudquery/compare/plugins-source-typeform-v1.0.6...plugins-source-typeform-v1.0.7) (2023-10-24)


### Bug Fixes

* **deps:** Update dependency cloudquery-plugin-sdk to v0.1.7 ([#14869](https://github.com/cloudquery/cloudquery/issues/14869)) ([1c83cbe](https://github.com/cloudquery/cloudquery/commit/1c83cbea76fc37b00457f7dc4b13b80b066b10cf))

## [1.0.6](https://github.com/cloudquery/cloudquery/compare/plugins-source-typeform-v1.0.5...plugins-source-typeform-v1.0.6) (2023-10-23)


### Bug Fixes

* **deps:** Update dependency cloudquery-plugin-sdk to v0.1.6 ([#14509](https://github.com/cloudquery/cloudquery/issues/14509)) ([c8cb05e](https://github.com/cloudquery/cloudquery/commit/c8cb05eaee332a84743799debc6b3954a8aa718e))

## [1.0.5](https://github.com/cloudquery/cloudquery/compare/plugins-source-typeform-v1.0.4...plugins-source-typeform-v1.0.5) (2023-10-10)


### Bug Fixes

* **deps:** Update dependency cloudquery-plugin-sdk to v0.1.5 ([#14478](https://github.com/cloudquery/cloudquery/issues/14478)) ([83f686a](https://github.com/cloudquery/cloudquery/commit/83f686a85a7ed419f482d596ca0c90c1ef908646))

## [1.0.4](https://github.com/cloudquery/cloudquery/compare/plugins-source-typeform-v1.0.3...plugins-source-typeform-v1.0.4) (2023-10-04)


### Bug Fixes

* **deps:** Update dependency cloudquery-plugin-sdk to v0.1.4 ([#14268](https://github.com/cloudquery/cloudquery/issues/14268)) ([8f4c911](https://github.com/cloudquery/cloudquery/commit/8f4c91104d43862df4bb5a775624d9a4a5130b34))

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
