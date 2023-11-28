# Changelog

## [1.0.8](https://github.com/cloudquery/cloudquery/compare/plugins-source-square-v1.0.7...plugins-source-square-v1.0.8) (2023-11-23)


### Bug Fixes

* Add missing tables for Docker plugins, remove on-demand rendering of tables ([#15433](https://github.com/cloudquery/cloudquery/issues/15433)) ([2286412](https://github.com/cloudquery/cloudquery/commit/22864120467b98d9284c533e58a7c9e09f8790ab))

## [1.0.7](https://github.com/cloudquery/cloudquery/compare/plugins-source-square-v1.0.6...plugins-source-square-v1.0.7) (2023-11-16)


### Bug Fixes

* **deps:** Update dependency cloudquery-plugin-sdk to v0.1.8 ([#15213](https://github.com/cloudquery/cloudquery/issues/15213)) ([02c07e0](https://github.com/cloudquery/cloudquery/commit/02c07e0f67f41cb62f1da2e84305f1e28b823cb4))
* **deps:** Update dependency pyarrow to v14 [SECURITY] ([#15213](https://github.com/cloudquery/cloudquery/issues/15213)) ([02c07e0](https://github.com/cloudquery/cloudquery/commit/02c07e0f67f41cb62f1da2e84305f1e28b823cb4))
* **deps:** Update module github.com/docker/docker to v24 [SECURITY] ([#15060](https://github.com/cloudquery/cloudquery/issues/15060)) ([41acd0e](https://github.com/cloudquery/cloudquery/commit/41acd0e4ac63221e90cca89a7137a8685692267d))

## [1.0.6](https://github.com/cloudquery/cloudquery/compare/plugins-source-square-v1.0.5...plugins-source-square-v1.0.6) (2023-10-24)


### Bug Fixes

* **deps:** Update dependency cloudquery-plugin-sdk to v0.1.7 ([#14869](https://github.com/cloudquery/cloudquery/issues/14869)) ([1c83cbe](https://github.com/cloudquery/cloudquery/commit/1c83cbea76fc37b00457f7dc4b13b80b066b10cf))

## [1.0.5](https://github.com/cloudquery/cloudquery/compare/plugins-source-square-v1.0.4...plugins-source-square-v1.0.5) (2023-10-23)


### Bug Fixes

* **deps:** Update dependency cloudquery-plugin-sdk to v0.1.6 ([#14509](https://github.com/cloudquery/cloudquery/issues/14509)) ([c8cb05e](https://github.com/cloudquery/cloudquery/commit/c8cb05eaee332a84743799debc6b3954a8aa718e))

## [1.0.4](https://github.com/cloudquery/cloudquery/compare/plugins-source-square-v1.0.3...plugins-source-square-v1.0.4) (2023-10-10)


### Bug Fixes

* **deps:** Update dependency cloudquery-plugin-sdk to v0.1.5 ([#14478](https://github.com/cloudquery/cloudquery/issues/14478)) ([83f686a](https://github.com/cloudquery/cloudquery/commit/83f686a85a7ed419f482d596ca0c90c1ef908646))

## [1.0.3](https://github.com/cloudquery/cloudquery/compare/plugins-source-square-v1.0.2...plugins-source-square-v1.0.3) (2023-10-04)


### Bug Fixes

* **deps:** Update dependency cloudquery-plugin-sdk to v0.1.4 ([#14268](https://github.com/cloudquery/cloudquery/issues/14268)) ([8f4c911](https://github.com/cloudquery/cloudquery/commit/8f4c91104d43862df4bb5a775624d9a4a5130b34))

## [1.0.2](https://github.com/cloudquery/cloudquery/compare/plugins-source-square-v1.0.1...plugins-source-square-v1.0.2) (2023-09-27)


### Bug Fixes

* Tidy go.mod ([#14061](https://github.com/cloudquery/cloudquery/issues/14061)) ([11bd971](https://github.com/cloudquery/cloudquery/commit/11bd971f6a0089c92e47af6be24f552b2d920f21))

## [1.0.1](https://github.com/cloudquery/cloudquery/compare/plugins-source-square-v1.0.0...plugins-source-square-v1.0.1) (2023-09-05)


### Bug Fixes

* **deps:** Update dependency cloudquery-plugin-sdk to v0.1.2 ([#13571](https://github.com/cloudquery/cloudquery/issues/13571)) ([de71388](https://github.com/cloudquery/cloudquery/commit/de713889f9ccdbb963839d37edc122ff0ca7518e))
* **deps:** Update dependency cloudquery-plugin-sdk to v0.1.3 ([#13583](https://github.com/cloudquery/cloudquery/issues/13583)) ([ca673c1](https://github.com/cloudquery/cloudquery/commit/ca673c16ffa38eaab303f9502823696c85cd4d61))
* **deps:** Update github.com/99designs/go-keychain digest to 9cf53c8 ([#13561](https://github.com/cloudquery/cloudquery/issues/13561)) ([a170256](https://github.com/cloudquery/cloudquery/commit/a17025657e92b017fe3c8bd37abfaa2354e6e818))

## [1.0.0](https://github.com/cloudquery/cloudquery/compare/plugins-source-square-v0.1.2...plugins-source-square-v1.0.0) (2023-08-14)


### ⚠ BREAKING CHANGES

* Update to latest Python SDK ([#13051](https://github.com/cloudquery/cloudquery/issues/13051)). This removes the `doc` command from the plugin, but the CLI command `cloudquery tables` can be used instead.

### Bug Fixes

* Update to latest Python SDK ([#13051](https://github.com/cloudquery/cloudquery/issues/13051)) ([a8254d6](https://github.com/cloudquery/cloudquery/commit/a8254d6233b15fb02b1987250bb0021bc904f507))

## [0.1.2](https://github.com/cloudquery/cloudquery/compare/plugins-source-square-v0.1.1...plugins-source-square-v0.1.2) (2023-08-09)


### Bug Fixes

* **deps:** Update dependency cloudquery-plugin-sdk to v0.0.11 ([#12863](https://github.com/cloudquery/cloudquery/issues/12863)) ([d6f063d](https://github.com/cloudquery/cloudquery/commit/d6f063d67d65652a494d1bb9d28f6c5115f58a90))

## [0.1.1](https://github.com/cloudquery/cloudquery/compare/plugins-source-square-v0.1.0...plugins-source-square-v0.1.1) (2023-08-08)


### This Release has the Following Changes to Tables
- Table `square_bookings` was added
- Table `square_disputes` was added
- Table `square_invoices` was added
- Table `square_locations` was added
- Table `square_merchants` was added
- Table `square_payments` was added
- Table `square_payouts` was added
- Table `square_refunds` was added

### Bug Fixes

* **deps:** Update dependency cloudquery-plugin-sdk to v0.0.10 ([#12733](https://github.com/cloudquery/cloudquery/issues/12733)) ([472ca1e](https://github.com/cloudquery/cloudquery/commit/472ca1eb903da6a922a4fdd9891917c47346e0bb))
* **deps:** Update dependency cloudquery-plugin-sdk to v0.0.9 ([#12729](https://github.com/cloudquery/cloudquery/issues/12729)) ([81d3dda](https://github.com/cloudquery/cloudquery/commit/81d3dda753af406e6fd458d1ce26f10bbdff2146))

## 0.1.0 (2023-08-02)


### Features

* **resources:** Add square plugin in python ([#12565](https://github.com/cloudquery/cloudquery/issues/12565)) ([f1938f3](https://github.com/cloudquery/cloudquery/commit/f1938f3bf96042315d966a6292205edecc7f0ac5))
