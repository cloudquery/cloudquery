# Changelog

## [2.1.4](https://github.com/cloudquery/cloudquery/compare/plugins-destination-mssql-v2.1.3...plugins-destination-mssql-v2.1.4) (2023-02-21)


### Bug Fixes

* **deps:** Update module golang.org/x/net to v0.7.0 [SECURITY] ([#8176](https://github.com/cloudquery/cloudquery/issues/8176)) ([fc4cef8](https://github.com/cloudquery/cloudquery/commit/fc4cef86dce4ca76ca8397e897ab744e48975834))

## [2.1.3](https://github.com/cloudquery/cloudquery/compare/plugins-destination-mssql-v2.1.2...plugins-destination-mssql-v2.1.3) (2023-02-14)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.36.0 ([#7809](https://github.com/cloudquery/cloudquery/issues/7809)) ([c85a9cb](https://github.com/cloudquery/cloudquery/commit/c85a9cb697477520e94a1fd260c56b89da62fc87))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.37.0 ([#7933](https://github.com/cloudquery/cloudquery/issues/7933)) ([dc9cffb](https://github.com/cloudquery/cloudquery/commit/dc9cffbf37bbc6fae73a20bf47e6bbf17e74d1f9))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.37.1 ([#8008](https://github.com/cloudquery/cloudquery/issues/8008)) ([c47aac0](https://github.com/cloudquery/cloudquery/commit/c47aac0b5e3190a04299713651b97e360043911f))

## [2.1.2](https://github.com/cloudquery/cloudquery/compare/plugins-destination-mssql-v2.1.1...plugins-destination-mssql-v2.1.2) (2023-02-07)


### Bug Fixes

* **deps:** Update golang.org/x/exp digest to f062dba ([#7531](https://github.com/cloudquery/cloudquery/issues/7531)) ([59d5575](https://github.com/cloudquery/cloudquery/commit/59d55758b0951553b8d246d1e78b4e3917ff1976))
* **deps:** Update google.golang.org/genproto digest to 1c01626 ([#7533](https://github.com/cloudquery/cloudquery/issues/7533)) ([c549c27](https://github.com/cloudquery/cloudquery/commit/c549c275077f1cdfb9df0b3f3c129cbf0b150552))
* **deps:** Update module github.com/Azure/azure-sdk-for-go/sdk/azidentity to v1.2.1 ([#7540](https://github.com/cloudquery/cloudquery/issues/7540)) ([3b5c838](https://github.com/cloudquery/cloudquery/commit/3b5c83832064d729ad1097728f7d12aedbbb9400))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.33.0 ([#7595](https://github.com/cloudquery/cloudquery/issues/7595)) ([c5adc75](https://github.com/cloudquery/cloudquery/commit/c5adc750d4b0242563997c04c582f8da27913095))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.33.1 ([#7614](https://github.com/cloudquery/cloudquery/issues/7614)) ([2fe665c](https://github.com/cloudquery/cloudquery/commit/2fe665cdd80d88c5699bb203bd7accd604dfba99))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.34.0 ([#7719](https://github.com/cloudquery/cloudquery/issues/7719)) ([6a33085](https://github.com/cloudquery/cloudquery/commit/6a33085c75adcf2387f7bbb5aa4f7a84ce7e2957))

## [2.1.1](https://github.com/cloudquery/cloudquery/compare/plugins-destination-mssql-v2.1.0...plugins-destination-mssql-v2.1.1) (2023-01-31)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.29.0 ([#7121](https://github.com/cloudquery/cloudquery/issues/7121)) ([b7441c9](https://github.com/cloudquery/cloudquery/commit/b7441c93c274ae3a6009474a2b28f44a172dd6dc))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.30.0 ([#7222](https://github.com/cloudquery/cloudquery/issues/7222)) ([73ca21c](https://github.com/cloudquery/cloudquery/commit/73ca21c4259545f7e949c9d780d8184db475d2ac))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.31.0 ([#7228](https://github.com/cloudquery/cloudquery/issues/7228)) ([36e8549](https://github.com/cloudquery/cloudquery/commit/36e8549f722658d909865723630fad1b2821db62))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.32.0 ([#7334](https://github.com/cloudquery/cloudquery/issues/7334)) ([b684122](https://github.com/cloudquery/cloudquery/commit/b68412222219f9ca160c0753290709d52de7fcd6))

## [2.1.0](https://github.com/cloudquery/cloudquery/compare/plugins-destination-mssql-v2.0.0...plugins-destination-mssql-v2.1.0) (2023-01-24)


### Features

* **mssql:** Use `LEFT JOIN` for overwrite ([#7086](https://github.com/cloudquery/cloudquery/issues/7086)) ([97c7981](https://github.com/cloudquery/cloudquery/commit/97c79816850b62989670916fd93f4f52528ce2b3))

## [2.0.0](https://github.com/cloudquery/cloudquery/compare/plugins-destination-mssql-v1.0.0...plugins-destination-mssql-v2.0.0) (2023-01-23)


### ⚠ BREAKING CHANGES

* **mssql:** Change column types from `varchar` to `nvarchar` to properly store Unicode characters

### Bug Fixes

* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.28.0 ([#7009](https://github.com/cloudquery/cloudquery/issues/7009)) ([12ac005](https://github.com/cloudquery/cloudquery/commit/12ac005428a355d06a5939fbe06a82d49533e662))
* **mssql:** Change column types from `varchar` to `nvarchar` to properly store Unicode characters ([748c7df](https://github.com/cloudquery/cloudquery/commit/748c7df498dfd7ff342c3e380af1ee75f5d1034b))
* **mssql:** Fix JSON transformer to store unescaped data ([748c7df](https://github.com/cloudquery/cloudquery/commit/748c7df498dfd7ff342c3e380af1ee75f5d1034b))

## 1.0.0 (2023-01-19)


### Features

* **mssql:** Microsoft SQL Server destination ([#6417](https://github.com/cloudquery/cloudquery/issues/6417)) ([c71bc88](https://github.com/cloudquery/cloudquery/commit/c71bc887280c9efb05afb98321a611bc33e387b0))
