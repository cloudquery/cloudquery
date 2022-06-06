# Changelog

All notable changes to this provider will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

<!-- 
## Unreleased
### 🚀 Added
### :gear: Changed
### :spider: Fixed
### 💥 Breaking Changes

-->

## [0.8.7](https://github.com/cloudquery/cq-provider-gcp/compare/v0.8.6...v0.8.7) (2022-06-06)


### Features

* Rollback hit only enabled APIs ([#321](https://github.com/cloudquery/cq-provider-gcp/issues/321)) ([86fc301](https://github.com/cloudquery/cq-provider-gcp/commit/86fc3011b8c6db737107f6ae9eaf0b597dda3d96))

## [0.8.6](https://github.com/cloudquery/cq-provider-gcp/compare/v0.8.5...v0.8.6) (2022-06-06)


### Features

* Hit only enabled APIs if permissions allow ([#317](https://github.com/cloudquery/cq-provider-gcp/issues/317)) ([7a48703](https://github.com/cloudquery/cq-provider-gcp/commit/7a48703fae18e571f109e8155fadd7d5a4950087))


### Bug Fixes

* Adjusted bigquery error severity when api is disabled ([#314](https://github.com/cloudquery/cq-provider-gcp/issues/314)) ([bac820e](https://github.com/cloudquery/cq-provider-gcp/commit/bac820e84e5e60e41771bc070549748dd8cd11ed))

## [0.8.5](https://github.com/cloudquery/cq-provider-gcp/compare/v0.8.4...v0.8.5) (2022-06-03)


### Features

* Hit only enabled/disabled apis ([#309](https://github.com/cloudquery/cq-provider-gcp/issues/309)) ([b7151f7](https://github.com/cloudquery/cq-provider-gcp/commit/b7151f78b5b0be6ea13063f2181d3c549a4a9122))


### Bug Fixes

* Remove unneeded type assertion ([#310](https://github.com/cloudquery/cq-provider-gcp/issues/310)) ([9980d06](https://github.com/cloudquery/cq-provider-gcp/commit/9980d0699bcab19d90d7207105e6405648768e2e))
* Wrap provider errors ([#313](https://github.com/cloudquery/cq-provider-gcp/issues/313)) ([953490f](https://github.com/cloudquery/cq-provider-gcp/commit/953490fa76c07e41026ae038c41f9786eebc6011))

### [0.8.4](https://github.com/cloudquery/cq-provider-gcp/compare/v0.8.3...v0.8.4) (2022-06-01)


### Features

* Convert id TypeBigInt to TypeString ([#305](https://github.com/cloudquery/cq-provider-gcp/issues/305)) ([f739796](https://github.com/cloudquery/cq-provider-gcp/commit/f739796df3347c3aacb95422ecfbdc06c89708fd))


### Bug Fixes

* **deps:** Update module github.com/cloudquery/cq-provider-sdk to v0.10.7 ([#308](https://github.com/cloudquery/cq-provider-gcp/issues/308)) ([05ff866](https://github.com/cloudquery/cq-provider-gcp/commit/05ff8662689a992eac0f3785b80fa1a6b60a163c))

### [0.8.3](https://github.com/cloudquery/cq-provider-gcp/compare/v0.8.2...v0.8.3) (2022-06-01)


### Features

* compute_url_maps ID from BigInt to String ([#298](https://github.com/cloudquery/cq-provider-gcp/issues/298)) ([3f2c8d7](https://github.com/cloudquery/cq-provider-gcp/commit/3f2c8d746c9362b6f82d8f14dc24252ecd97f711))

### [0.8.2](https://github.com/cloudquery/cq-provider-gcp/compare/v0.8.1...v0.8.2) (2022-06-01)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/cq-provider-sdk to v0.10.6 ([#300](https://github.com/cloudquery/cq-provider-gcp/issues/300)) ([1f529f2](https://github.com/cloudquery/cq-provider-gcp/commit/1f529f28737f201ca9ad862d1a7116400a2e4519))

### [0.8.1](https://github.com/cloudquery/cq-provider-gcp/compare/v0.8.0...v0.8.1) (2022-05-31)


### Features

* Add GKE clusters ([#267](https://github.com/cloudquery/cq-provider-gcp/issues/267)) ([4edf94c](https://github.com/cloudquery/cq-provider-gcp/commit/4edf94c42f0eaaa5425993185f6ec53899cd88c3))


### Bug Fixes

* **deps:** Update module github.com/cloudquery/cq-provider-sdk to v0.10.2 ([#286](https://github.com/cloudquery/cq-provider-gcp/issues/286)) ([2261fdf](https://github.com/cloudquery/cq-provider-gcp/commit/2261fdf070d1fa068f7502778c82f99b013d7f8a))
* **deps:** Update module github.com/cloudquery/cq-provider-sdk to v0.10.3 ([#287](https://github.com/cloudquery/cq-provider-gcp/issues/287)) ([c145063](https://github.com/cloudquery/cq-provider-gcp/commit/c14506324fa6957792f4fc9bfe3b514af758b392))
* **deps:** Update module github.com/cloudquery/cq-provider-sdk to v0.10.4 ([#290](https://github.com/cloudquery/cq-provider-gcp/issues/290)) ([ce4e586](https://github.com/cloudquery/cq-provider-gcp/commit/ce4e586b14db526c758669f61c2d202d7e6b4187))
* **deps:** Update module github.com/cloudquery/cq-provider-sdk to v0.10.5 ([#296](https://github.com/cloudquery/cq-provider-gcp/issues/296)) ([3e2e5b4](https://github.com/cloudquery/cq-provider-gcp/commit/3e2e5b49c5125bbc708c2adf6e8248edbb593ed7))
* Remove relation tables PK ([#265](https://github.com/cloudquery/cq-provider-gcp/issues/265)) ([802e532](https://github.com/cloudquery/cq-provider-gcp/commit/802e532563fd90a8a930dd6234fc6ede8034ad1c))

## [0.8.0](https://github.com/cloudquery/cq-provider-gcp/compare/v0.7.4...v0.8.0) (2022-05-24)


### ⚠ BREAKING CHANGES

* Remove migrations (#271)

### Features

* Remove migrations ([#271](https://github.com/cloudquery/cq-provider-gcp/issues/271)) ([cf0ed91](https://github.com/cloudquery/cq-provider-gcp/commit/cf0ed91ab0a7a2124da467b131c333851010b617))


### Bug Fixes

* **deps:** Update module github.com/cloudquery/cq-provider-sdk to v0.10.1 ([#281](https://github.com/cloudquery/cq-provider-gcp/issues/281)) ([6419145](https://github.com/cloudquery/cq-provider-gcp/commit/6419145b95cf19659df20d914cd621ccc71cf6c1))
* Temporary remove serviceusage ([#284](https://github.com/cloudquery/cq-provider-gcp/issues/284)) ([7d7adbd](https://github.com/cloudquery/cq-provider-gcp/commit/7d7adbd6b955ca56f2bd61584804209816934fef))

### [0.7.4](https://github.com/cloudquery/cq-provider-gcp/compare/v0.7.3...v0.7.4) (2022-05-23)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/cq-provider-sdk to v0.9.5 ([#277](https://github.com/cloudquery/cq-provider-gcp/issues/277)) ([892ec48](https://github.com/cloudquery/cq-provider-gcp/commit/892ec486477eb6642e71a15da2ff913653f564b1))
* Ignore not found in gcp services ([#263](https://github.com/cloudquery/cq-provider-gcp/issues/263)) ([4850734](https://github.com/cloudquery/cq-provider-gcp/commit/48507349ebc3efe7229c38dfa179c0fa678a8edf))

### [0.7.3](https://github.com/cloudquery/cq-provider-gcp/compare/v0.7.2...v0.7.3) (2022-05-17)


### Features

* Added InstanceGroups ([#186](https://github.com/cloudquery/cq-provider-gcp/issues/186)) ([7c049cd](https://github.com/cloudquery/cq-provider-gcp/commit/7c049cda3cae79f3b1f04a6a733d7e23f254c24c))
* Added serviceusage_services ([#249](https://github.com/cloudquery/cq-provider-gcp/issues/249)) ([df90b8f](https://github.com/cloudquery/cq-provider-gcp/commit/df90b8f759dc64117788a3265904afbb44331a63))


### Bug Fixes

* **deps:** Update module github.com/cloudquery/cq-provider-sdk to v0.9.4 ([#261](https://github.com/cloudquery/cq-provider-gcp/issues/261)) ([dbd1398](https://github.com/cloudquery/cq-provider-gcp/commit/dbd139805aabd0700346fc5a0c1cbd2313a79103))

### [0.7.2](https://github.com/cloudquery/cq-provider-gcp/compare/v0.7.1...v0.7.2) (2022-05-10)


### Features

* Billing resources ([#239](https://github.com/cloudquery/cq-provider-gcp/issues/239)) ([2ba023f](https://github.com/cloudquery/cq-provider-gcp/commit/2ba023fe5c5156e12994ad3e8d11bc7fd605caa8))


### Bug Fixes

* **deps:** Update module github.com/cloudquery/cq-provider-sdk to v0.9.3 ([#232](https://github.com/cloudquery/cq-provider-gcp/issues/232)) ([9815dd6](https://github.com/cloudquery/cq-provider-gcp/commit/9815dd68d56892c09e4ab11bb689f244a2d17b6e))

## [v0.5.1] - 2022-01-03
###### SDK Version: 0.6.1
### :spider: Fixed
* Fixed issues with disabled services [#84](https://github.com/cloudquery/cq-provider-gcp/pull/84)
### :gear: Changed
* Updated to SDK version [v0.6.1](https://github.com/cloudquery/cq-provider-sdk/blob/v0.6.1/CHANGELOG.md#v061---2022-01-03)


## [v0.5.0] - 2021-11-21
###### SDK Version: 0.5.0

### :rocket: Added
* Add support for error classifier in GCP. [#78](https://github.com/cloudquery/cq-provider-gcp/issues/78)

### :gear: Changed
* Upgraded to SDK Version [v0.5.0](https://github.com/cloudquery/cq-provider-sdk/blob/main/CHANGELOG.md#v050---2021-10-21)

## [v0.4.8] - 2021-10-07
###### SDK Version: v0.4.9

### :gear: Changed
* Upgraded to SDK Version [v0.4.9](https://github.com/cloudquery/cq-provider-sdk/blob/main/CHANGELOG.md)
* Dropped column `post_key_revocation_action_type` from `gcp_compute_instances`


## [v0.4.7] - 2021-10-03
###### SDK Version: v0.4.7

### :rocket: Added
* added migration tests for improved stability [#67](https://github.com/cloudquery/cq-provider-gcp/pull/67)

### :spider: Fixed
* Fixed issues in integration tests [#69](https://github.com/cloudquery/cq-provider-gcp/pull/69) [#70](https://github.com/cloudquery/cq-provider-gcp/pull/70)

### :gear: Changed
* Upgraded to SDK Version [v0.4.7](https://github.com/cloudquery/cq-provider-sdk/blob/main/CHANGELOG.md#v047---2021-09-23)

## [v0.4.6] - 2021-09-14
###### SDK Version: v0.4.4

### :spider: Fixed
* Fixed forbidden error in gcp buckets iam policy fetch [#59](https://github.com/cloudquery/cq-provider-gcp/pull/59)

### :gear: Changed
* Upgraded to SDK Version [v0.4.4](https://github.com/cloudquery/cq-provider-sdk/blob/main/CHANGELOG.md#v044---2021-09-13)

### :rocket: Added
* added e2e intergration tests for improved stability [#51](https://github.com/cloudquery/cq-provider-gcp/pull/51)

## [v0.4.5] - 2021-09-09
###### SDK Version: 0.4.3

### :spider: Fixed
Embed migrations [#58](https://github.com/cloudquery/cq-provider-gcp/pull/58)

## [v0.4.4] - 2021-09-09
###### SDK Version: 0.4.3

### :rocket: Added
* added support for urlmaps resources [#47](https://github.com/cloudquery/cq-provider-gcp/issues/47)

### :spider: Fixed
* iam.service_accounts duplicate primary key fixed. [#53](https://github.com/cloudquery/cq-provider-gcp/pull/53)

### :gear: Changed
* Upgraded to SDK Version [0.4.3](https://github.com/cloudquery/cq-provider-sdk/blob/main/CHANGELOG.md)

## [v0.4.3] - 2021-08-13
###### SDK Version: 0.3.2

### :spider: Fixed
* Add Ignore error to all resources. [#45](https://github.com/cloudquery/cq-provider-gcp/pull/45)

## [v0.4.2] - 2021-08-12
###### SDK Version: 0.3.2

### :spider: Fixed
* Fixed common errors in gcp provider. [#43](https://github.com/cloudquery/cq-provider-gcp/pull/43)

## [v0.4.1] - 2021-08-11
###### SDK Version: 0.3.2
### :gear: Changed
* Upgraded to SDK Version [0.3.2](https://github.com/cloudquery/cq-provider-sdk/blob/main/CHANGELOG.md#v032---2020-08-11)

## [v0.4.0] - 2021-07-27
###### SDK Version: 0.3.0

### :rocket: Added

* Added a changelog :)
* Added support for passing credentials in `config.hcl` [#35](https://github.com/cloudquery/cq-provider-gcp/pull/35) [#36](https://github.com/cloudquery/cq-provider-gcp/pull/36) 

### :gear: Changed
* Upgraded to SDK Version [0.3.0](https://github.com/cloudquery/cq-provider-sdk/blob/main/CHANGELOG.md)
* **Breaking Change**: default CloudQuery "id" from `id` to `cq_id` [#41](https://github.com/cloudquery/cq-provider-sdk/pull/41)


### :spider: Fixed
* Fixed Domain registrations bad path. [#38](https://github.com/cloudquery/cq-provider-gcp/pull/38)

## [v0.3.6] - 2021-07-15
###### SDK Version: 0.2.8

Base version at which changelog was introduced.

### Supported Resources
- bigquery.datasets
- cloudfunctions.functions
- compute.addresses
- compute.autoscalers
- compute.backend_services
- compute.disk_types
- compute.disks
- compute.firewalls
- compute.forwarding_rules
- compute.images
- compute.instances
- compute.interconnects
- compute.urlmaps
- compute.networks
- compute.projects
- compute.ssl_certificates
- compute.ssl_policies
- compute.subnetworks
- compute.target_https_proxies
- compute.target_ssl_proxies
- compute.vpn_gateways
- crm.projects
- dns.managed_zones
- dns.policies
- domains.registrations
- iam.project_roles
- iam.service_accounts
- kms.keys
- logging.metrics
- logging.sinks
- monitoring.alert_policies
- resource_manager.folders
- resource_manager.projects
- sql.instances
- storage.buckets
