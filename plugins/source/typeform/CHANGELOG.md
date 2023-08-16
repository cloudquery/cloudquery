# Changelog

## [1.0.0](https://github.com/cloudquery/cloudquery/compare/plugins-source-typeform-v0.1.1...plugins-source-typeform-v1.0.0) (2023-08-16)


### This Release has the Following Changes to Tables
- Table `typeform_form_responses`: column type changed from `extension<json<JSONType>>` to `json` for `answers` (:warning: breaking)
- Table `typeform_form_responses`: column type changed from `extension<json<JSONType>>` to `json` for `calculated` (:warning: breaking)
- Table `typeform_form_responses`: column type changed from `extension<json<JSONType>>` to `json` for `hidden` (:warning: breaking)
- Table `typeform_form_responses`: column type changed from `extension<json<JSONType>>` to `json` for `metadata` (:warning: breaking)
- Table `typeform_form_responses`: column type changed from `extension<json<JSONType>>` to `json` for `variables` (:warning: breaking)
- Table `typeform_form_responses`: column type changed from `string` to `utf8` for `form_id` (:warning: breaking)
- Table `typeform_form_responses`: column type changed from `string` to `utf8` for `landing_id` (:warning: breaking)
- Table `typeform_form_responses`: column type changed from `string` to `utf8` for `response_id` (:warning: breaking)
- Table `typeform_form_responses`: column type changed from `string` to `utf8` for `token` (:warning: breaking)
- Table `typeform_forms`: column type changed from `extension<json<JSONType>>` to `json` for `_links` (:warning: breaking)
- Table `typeform_forms`: column type changed from `extension<json<JSONType>>` to `json` for `self` (:warning: breaking)
- Table `typeform_forms`: column type changed from `extension<json<JSONType>>` to `json` for `settings` (:warning: breaking)
- Table `typeform_forms`: column type changed from `extension<json<JSONType>>` to `json` for `theme` (:warning: breaking)
- Table `typeform_forms`: column type changed from `string` to `utf8` for `id` (:warning: breaking)
- Table `typeform_forms`: column type changed from `string` to `utf8` for `title` (:warning: breaking)
- Table `typeform_forms`: column type changed from `string` to `utf8` for `type` (:warning: breaking)

### ⚠ BREAKING CHANGES

* Upgrade SDK to v0.1.1 ([#13055](https://github.com/cloudquery/cloudquery/issues/13055))

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
